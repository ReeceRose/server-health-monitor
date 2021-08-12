package service

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type healthService struct {
	healthRepository repository.IHealthRepository
	hostRepository   repository.IHostRepository
	log              logger.Logger
}

var (
	_ IHealthService = (*healthService)(nil)
)

// NewHealthService returns an instanced health service
func NewHealthService(healthRepository repository.IHealthRepository, hostRepository repository.IHostRepository) IHealthService {
	return &healthService{
		healthRepository: healthRepository,
		hostRepository:   hostRepository,
		log:              logger.Instance(),
	}
}

// GetHealth returns all health data
func (s *healthService) GetHealth(requestID string) types.HealthReponse {
	s.log.Info("attemping to get all health - Request ID: " + requestID)
	data, err := s.healthRepository.Find(bson.M{})
	if err != nil {
		return types.HealthReponse{
			Data:       []types.Health{},
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Sprintf("failed to get all health data - Request ID: %s", requestID),
			Success:    false,
		}
	}

	s.log.Info("successfully got all health data - Request ID: " + requestID)

	return types.HealthReponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// GetHealthByAgentID returns all health data for a given agent
func (s *healthService) GetHealthByAgentID(requestID, agentID string) types.HealthReponse {
	s.log.Infof("attemping to get health data for agent: %s - Request ID: %s", agentID, requestID)

	data, err := s.healthRepository.Find(bson.M{"agentID": agentID})
	if err != nil {
		return types.HealthReponse{
			Data:       []types.Health{},
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Sprintf("failed to get data for agent: %s - Request ID: %s", agentID, requestID),
			Success:    false,
		}
	}

	s.log.Infof("successfully got health data for agent: %s - Request ID: %s", agentID, requestID)

	return types.HealthReponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// AddHealth inserts new health data for a given agent
func (s *healthService) AddHealth(requestID string, agentID string, data *types.Health) types.HealthReponse {
	s.log.Infof("attemping to insert health data for agent: %s - Request ID: %s", agentID, requestID)

	data.AgentID = agentID
	data.ID = primitive.NewObjectID()
	now := time.Now().UTC().UnixNano()
	data.CreateTime = now

	_, err := s.healthRepository.Insert(data)

	if err != nil {
		return types.HealthReponse{
			Data:       []types.Health{},
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Sprintf("failed to insert data for agent: %s - Request ID %s", agentID, requestID),
			Success:    false,
		}
	}

	s.log.Infof("successfully inserted health data for agent: %s - Request ID: %s", agentID, requestID)

	return types.HealthReponse{
		Data:       []types.Health{*data},
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// TODO: get all health for all agents since the last timestamp

// GetLatestHealthDataByAgentID returns the latest health data since a givrm time for a given agent
func (s *healthService) GetLatestHealthDataByAgentID(requestID string, agentID string, time int64) types.HealthReponse {
	s.log.Infof("attemping to get health data for agent: %s - Request ID: %s", agentID, requestID)

	data, err := s.healthRepository.Find(bson.M{
		"agentID": bson.M{"$eq": agentID},
		"$and": []bson.M{
			{
				"createTime": bson.M{"$gt": time},
			},
		},
	})

	if err != nil {
		return types.HealthReponse{
			Data:       []types.Health{},
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Sprintf("failed to get health data for agent: %s - Request ID: %s", agentID, requestID),
			Success:    false,
		}
	}

	s.log.Infof("successfully got health data for agent: %s - Request ID: %s", agentID, requestID)

	sort.Slice(data, func(i, j int) bool {
		return data[i].CreateTime > data[j].CreateTime
	})

	return types.HealthReponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// GetLatestHealthDataForAgents returns all health data for all agents since a given time for an agent
func (s *healthService) GetLatestHealthDataForAgents(requestID string, since int64) types.HostReponse {
	data, err := s.hostRepository.Find(bson.M{})

	if err != nil {
		return types.HostReponse{
			Data:       []types.Host{},
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Sprintf("failed to get hosts - Request ID: %s", requestID),
			Success:    false,
		}
	}

	for i, host := range data {
		res := s.GetLatestHealthDataByAgentID(requestID, host.AgentID, since)
		if res.Success {
			data[i].Health = res.Data
		}
	}

	return types.HostReponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}
