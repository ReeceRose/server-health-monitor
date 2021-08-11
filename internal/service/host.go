package service

import (
	"fmt"
	"net/http"
	"sort"
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
	hostRepository   repository.IHostRepository
	healthRepository repository.IHealthRepository
	log              logger.Logger
}

var (
	_ IHostService = (*hostService)(nil)
)

// NewHostService returns an instanced host service
func NewHostService(hostRepository repository.IHostRepository, healthRepository repository.IHealthRepository) IHostService {
	return &hostService{
		hostRepository:   hostRepository,
		healthRepository: healthRepository,
		log:              logger.Instance(),
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
		data[i].Online = s.isHostOnline(data[i].AgentID) //TODO: refactor so two database reads aren't required
		data[i].Health = s.getHealthData(data[i].AgentID, minutesToIncludeHealthData)
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

	data[0].Online = s.isHostOnline(agentID) // TODO: refactor to remove two database reads
	data[0].Health = s.getHealthData(agentID, minutesToIncludeHealthData)

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

func (s *hostService) isHostOnline(agentID string) bool {
	return len(s.getHealthData(agentID, 0)) >= 1
}

func (s *hostService) getHealthData(agentID string, delay int) []types.Health {
	res, err := s.healthRepository.Find(bson.M{
		"agentID": bson.M{"$eq": agentID},
		"$and": []bson.M{
			{
				"createTime": bson.M{"$gt": utils.GetMinimumLastHealthPacketTime(time.Now(), delay)},
			},
		},
	})
	if err != nil {
		logger.Instance().Info(err.Error())
		return nil
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].CreateTime > res[j].CreateTime
	})
	return res
}
