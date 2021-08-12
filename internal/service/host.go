package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type hostService struct {
	hostRepository repository.IHostRepository
	healthService  IHealthService
	log            logger.Logger
}

var (
	_ IHostService = (*hostService)(nil)
)

// NewHostService returns an instanced host service
func NewHostService(hostRepository repository.IHostRepository, healthService IHealthService) IHostService {
	return &hostService{
		hostRepository: hostRepository,
		healthService:  healthService,
		log:            logger.Instance(),
	}
}

// GetHosts returns all hosts
func (s *hostService) GetHosts(requestID string, includeHealthData bool) types.StandardResponse {
	s.log.Info("attemping to get all hosts - Request ID: " + requestID)
	data, err := s.hostRepository.Find(bson.M{})
	if err != nil {
		return types.StandardResponse{
			Error:      "failed to get all hosts data - Request ID: " + requestID,
			StatusCode: http.StatusInternalServerError,
			Data:       []types.Host{},
			Success:    false,
		}
	}
	if includeHealthData {
		minutesToIncludeHealthData, err := strconv.Atoi(utils.GetVariable(consts.MINUTES_TO_INCLUDE_HEALTH))
		if err != nil {
			minutesToIncludeHealthData = 5
		}

		for i := range data {
			data[i].Online = s.isHostOnline(requestID, data[i].AgentID) //TODO: refactor so two database reads aren't required
			data[i].Health = s.healthService.GetLatestHealthDataByAgentID(requestID, data[i].AgentID,
				utils.GetMinimumLastHealthPacketTime(time.Now(), minutesToIncludeHealthData),
			)
			if len(data[i].Health) >= 1 {
				data[i].LastConnected = data[i].Health[0].CreateTime
			}
		}
	}

	s.log.Info("successfully got all hosts data - Request ID: " + requestID)

	return types.StandardResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// GetHostByID returns a host given an id
func (s *hostService) GetHostByID(requestID, agentID string, includeHealthData bool) types.StandardResponse {
	s.log.Infof("attemping to get host data for agent: %s - Request ID: %s", agentID, requestID)

	data, err := s.hostRepository.Find(bson.M{"agentID": agentID})
	if err != nil {
		return types.StandardResponse{
			Error:      fmt.Sprintf("failed to get host data for agent: %s - Request ID: %s", agentID, requestID),
			StatusCode: http.StatusInternalServerError,
			Data:       []types.Host{},
			Success:    false,
		}
	}

	if len(data) == 0 {
		return types.StandardResponse{
			Error:      fmt.Sprintf("failed to get host data for agent: %s - Request ID: %s", agentID, requestID),
			StatusCode: http.StatusNoContent,
			Data:       []types.Host{},
			Success:    false,
		}
	}

	minutesToIncludeHealthData, err := strconv.Atoi(utils.GetVariable(consts.MINUTES_TO_INCLUDE_HEALTH))
	if err != nil {
		minutesToIncludeHealthData = 5
	}

	data[0].Online = s.isHostOnline(requestID, agentID) // TODO: refactor to remove two database reads
	data[0].Health = s.healthService.GetLatestHealthDataByAgentID(requestID, agentID,
		utils.GetMinimumLastHealthPacketTime(time.Now(), minutesToIncludeHealthData),
	)

	if len(data[0].Health) >= 1 {
		data[0].LastConnected = data[0].Health[0].CreateTime
	}

	s.log.Infof("successfully got host data for agent: %s - Request ID: %s", agentID, requestID)

	return types.StandardResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// AddHost inserts new host
func (s *hostService) AddHost(requestID string, agentID string, data *types.Host) types.StandardResponse {
	s.log.Infof("attemping to insert host data for agent: %s - Request ID: %s", agentID, requestID)

	now := time.Now().UTC().UnixNano()
	data.AgentID = agentID
	data.UpdateTime = now
	res := s.GetHostByID(requestID, agentID, false)
	if original := res.Data.([]types.Host); len(original) >= 1 {
		data.ID = original[0].ID
		data.CreateTime = original[0].CreateTime
		err := s.hostRepository.UpdateByID(data)

		if err != nil {
			return types.StandardResponse{
				Error:      fmt.Sprintf("failed to update data for agent: %s - Request ID %s", agentID, requestID),
				StatusCode: http.StatusInternalServerError,
				Data:       []types.Host{},
				Success:    false,
			}
		}

		s.log.Infof("successfully updated host data for agent: %s - Request ID: %s", agentID, requestID)
	} else {
		data.ID = primitive.NewObjectID()
		data.CreateTime = now

		_, err := s.hostRepository.Insert(data)

		if err != nil {
			return types.StandardResponse{
				Error:      fmt.Sprintf("failed to insert data for agent: %s - Request ID %s", agentID, requestID),
				StatusCode: http.StatusInternalServerError,
				Data:       []types.Host{},
				Success:    false,
			}
		}

		s.log.Infof("successfully inserted host data for agent: %s - Request ID: %s", agentID, requestID)
	}

	return types.StandardResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

func (s *hostService) GetLatestHealthDataForAgents(requestID string, lastCheck int64) types.StandardResponse {
	data := s.GetHosts(requestID, false)
	if !data.Success {
		return types.StandardResponse{
			Error:      fmt.Sprintf("failed to get hosts - Request ID: %s", requestID),
			StatusCode: http.StatusInternalServerError,
			Data:       []types.Health{},
			Success:    false,
		}
	}

	hosts := data.Data.([]types.Host)
	for i, host := range hosts {
		hosts[i].Health = s.healthService.GetLatestHealthDataByAgentID(requestID, host.AgentID, lastCheck)
	}

	return types.StandardResponse{
		Data:       hosts,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

func (s *hostService) isHostOnline(requestID string, agentID string) bool {
	return len(s.healthService.GetLatestHealthDataByAgentID(requestID, agentID, 0)) >= 1
}
