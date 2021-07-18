package service

import (
	"fmt"
	"net/http"

	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HealthService struct {
	healthRepository *repository.HealthRepository
	log              logger.Logger
}

// NewHealthService returns an instanced health service
func NewHealthService() *HealthService {
	return &HealthService{
		healthRepository: repository.NewHealthRepository(),
		log:              logger.Instance(),
	}
}

// GetHealth returns all health data for a given agent
func (s *HealthService) GetHealth(agentID string) types.StandardResponse {
	s.log.Info("attemping to get health data for agent: " + agentID)

	data, err := s.healthRepository.Find(bson.M{"agentID": agentID})
	if err != nil {
		return types.StandardResponse{
			Error:      fmt.Sprintf("failed to get data for server ID: %s", agentID),
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
			Success:    false,
		}
	}

	s.log.Info("successfully got health data for agent: " + agentID)

	return types.StandardResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}

// AddHealth inserts new health data for a given agent
func (s *HealthService) AddHealth(agentID string, data *types.Health) types.StandardResponse {
	s.log.Info("attemping to insert health data for agent: " + agentID)

	data.AgentID = agentID
	data.ID = primitive.NewObjectID()

	_, err := s.healthRepository.Insert(data)

	if err != nil {
		return types.StandardResponse{
			Error:      fmt.Sprintf("failed to insert data for server ID: %s", agentID),
			StatusCode: http.StatusInternalServerError,
			Data:       nil,
			Success:    false,
		}
	}

	s.log.Info("successfully inserted health data for agent: " + agentID)

	return types.StandardResponse{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}
