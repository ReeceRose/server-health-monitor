package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type healthService struct {
	healthRepository repository.IHealthRepository
	log              logger.Logger
}

var (
	_ IHealthService = (*healthService)(nil)
)

// NewHealthService returns an instanced health service
func NewHealthService(repository repository.IHealthRepository) IHealthService {
	return &healthService{
		healthRepository: repository,
		log:              logger.Instance(),
	}
}

// GetHealth returns all health data
func (s *healthService) GetHealth(requestID string) types.StandardResponse {
	s.log.Info("attemping to get all health - Request ID: " + requestID)
	data, err := s.healthRepository.Find(bson.M{})
	if err != nil {
		return types.StandardResponse{
			Error:      "failed to get all health data - Request ID: " + requestID,
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
			Success:    false,
		}
	}

	s.log.Info("successfully got all health data - Request ID: " + requestID)

	return types.StandardResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// GetHealthByAgentID returns all health data for a given agent
func (s *healthService) GetHealthByAgentID(requestID, agentID string) types.StandardResponse {
	s.log.Infof("attemping to get health data for agent: %s - Request ID: %s", agentID, requestID)

	data, err := s.healthRepository.Find(bson.M{"agentID": agentID})
	if err != nil {
		return types.StandardResponse{
			Error:      fmt.Sprintf("failed to get data for agent: %s - Request ID: %s", agentID, requestID),
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
			Success:    false,
		}
	}

	s.log.Infof("successfully got health data for agent: %s - Request ID: %s", agentID, requestID)

	return types.StandardResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// AddHealth inserts new health data for a given agent
func (s *healthService) AddHealth(requestID string, agentID string, data *types.Health) types.StandardResponse {
	s.log.Infof("attemping to insert health data for agent: %s - Request ID: %s", agentID, requestID)

	data.AgentID = agentID
	data.ID = primitive.NewObjectID()
	now := time.Now().UTC().UnixNano()
	data.CreateTime = now

	_, err := s.healthRepository.Insert(data)

	if err != nil {
		return types.StandardResponse{
			Error:      fmt.Sprintf("failed to insert data for agent: %s - Request ID %s", agentID, requestID),
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
			Success:    false,
		}
	}

	s.log.Infof("successfully inserted health data for agent: %s - Request ID: %s", agentID, requestID)

	return types.StandardResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}
