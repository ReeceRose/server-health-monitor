package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (s *hostService) GetHosts(requestID string, includeHealthData bool) types.HostReponse {
	s.log.Info("attemping to get all hosts - Request ID: " + requestID)
	data, err := s.hostRepository.Find(bson.M{})
	if err != nil {
		return types.HostReponse{
			Data:       []types.Host{},
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Sprintf("failed to get all hosts data - Request ID: %s", requestID),
			Success:    false,
		}
	}

	if includeHealthData {
		s.getHealthDataForHosts(requestID, &data)
	}

	s.log.Info("successfully got all hosts data - Request ID: " + requestID)

	return types.HostReponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// GetHostByID returns a host given an id
func (s *hostService) GetHostByID(requestID, agentID string, includeHealthData bool) types.HostReponse {
	s.log.Infof("attemping to get host data for agent: %s - Request ID: %s", agentID, requestID)

	data, err := s.hostRepository.Find(bson.M{"agentID": agentID})
	if err != nil {
		return types.HostReponse{
			Data:       []types.Host{},
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Sprintf("failed to get host data for agent: %s - Request ID: %s", agentID, requestID),
			Success:    false,
		}
	}
	if len(data) == 0 {
		return types.HostReponse{
			Data:       []types.Host{},
			StatusCode: http.StatusNoContent,
			Error:      fmt.Sprintf("failed to get host data for agent: %s - Request ID: %s", agentID, requestID),
			Success:    false,
		}
	}

	if includeHealthData {
		s.getHealthDataForHosts(agentID, &data)
	}

	s.log.Infof("successfully got host data for agent: %s - Request ID: %s", agentID, requestID)

	return types.HostReponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// AddHost inserts new host
func (s *hostService) AddHost(requestID string, agentID string, data *types.Host) types.HostReponse {
	s.log.Infof("attemping to insert host data for agent: %s - Request ID: %s", agentID, requestID)

	now := time.Now().UTC().UnixNano()
	data.AgentID = agentID
	data.UpdateTime = now
	res := s.GetHostByID(requestID, agentID, false)
	if res.Success && len(res.Data) >= 1 {
		first := res.Data[0]
		data.ID = first.ID
		data.CreateTime = first.CreateTime
		err := s.hostRepository.UpdateByID(data)

		if err != nil {
			return types.HostReponse{
				Data:       []types.Host{},
				StatusCode: http.StatusInternalServerError,
				Error:      fmt.Sprintf("failed to update data for agent: %s - Request ID %s", agentID, requestID),
				Success:    false,
			}
		}

		s.log.Infof("successfully updated host data for agent: %s - Request ID: %s", agentID, requestID)
	} else {
		data.ID = primitive.NewObjectID()
		data.CreateTime = now

		_, err := s.hostRepository.Insert(data)

		if err != nil {
			return types.HostReponse{
				Data:       []types.Host{},
				StatusCode: http.StatusInternalServerError,
				Error:      fmt.Sprintf("failed to insert data for agent: %s - Request ID %s", agentID, requestID),
				Success:    false,
			}
		}

		s.log.Infof("successfully inserted host data for agent: %s - Request ID: %s", agentID, requestID)
	}

	return types.HostReponse{
		Data:       []types.Host{*data},
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

func (s *hostService) isHostOnline(requestID string, agentID string) bool {
	res := s.healthService.GetLatestHealthDataByAgentID(requestID, agentID,
		utils.GetMinimumLastHealthPacketTime(time.Now(),
			utils.GetMinutesToIncludeHealthData(),
		),
	)
	return res.Success && len(res.Data) >= 1
}

func (s *hostService) getHealthDataForHosts(requestID string, hosts *[]types.Host) {
	for i := range *hosts {
		(*hosts)[i].Online = s.isHostOnline(requestID, (*hosts)[i].AgentID)
		res := s.healthService.GetLatestHealthDataByAgentID(requestID, (*hosts)[i].AgentID,
			utils.GetMinimumLastHealthPacketTime(time.Now(), utils.GetMinutesToIncludeHealthData()),
		)
		if res.Success {
			(*hosts)[i].Health = res.Data
			if len((*hosts)[i].Health) >= 1 {
				(*hosts)[i].LastConnected = (*hosts)[i].Health[0].CreateTime
			} else {
				// Get the last known health packet
				options := options.Find()
				options.SetSort(bson.D{primitive.E{Key: "createTime", Value: -1}})
				options.SetLimit(1)
				health := s.healthService.GetHealthForAgentWithOptions(requestID, (*hosts)[i].AgentID, options)
				if len(health) >= 1 {
					(*hosts)[i].LastConnected = health[0].CreateTime
				}
			}
		}
	}
}
