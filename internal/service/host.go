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
func (s *hostService) GetHosts(requestID string) types.StandardResponse {
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

	minutesToIncludeHealthData, err := strconv.Atoi(utils.GetVariable(consts.MINUTES_TO_INCLUDE_HEALTH))
	if err != nil {
		minutesToIncludeHealthData = 5
	}

	for i := range data {
		data[i].Online = s.isHostOnline(requestID, data[i].AgentID) //TODO: refactor so two database reads aren't required
		res := s.healthService.GetLatestHealthDataByAgentID(requestID, data[i].AgentID, minutesToIncludeHealthData)
		data[0].Health = res.Data.([]types.Health)

	}

	s.log.Info("successfully got all hosts data - Request ID: " + requestID)

	return types.StandardResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// GetHostByID returns a host given an id
func (s *hostService) GetHostByID(requestID, agentID string) types.StandardResponse {
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
	res := s.healthService.GetLatestHealthDataByAgentID(requestID, agentID, minutesToIncludeHealthData)
	data[0].Health = res.Data.([]types.Health)

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
	res := s.GetHostByID(requestID, agentID)
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

func (s *hostService) isHostOnline(requestID string, agentID string) bool {
	res := s.healthService.GetLatestHealthDataByAgentID(requestID, agentID, 0)
	return len(res.Data.([]types.Health)) >= 1
}
