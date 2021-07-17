package service

import (
	"fmt"
	"net/http"

	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"go.mongodb.org/mongo-driver/bson"
)

type HealthService struct {
	healthRepository *repository.HealthRepository
	log              logger.Logger
}

func NewHealthService() *HealthService {
	return &HealthService{
		healthRepository: repository.NewHealthRepository(),
		log:              logger.Instance(),
	}
}

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

func (s *HealthService) AddHealth(agentID string, health types.Health) types.StandardResponse {
	// TODO: add halth data
	return types.StandardResponse{
		Data:       nil,
		StatusCode: http.StatusOK,
		Success:    true,
	}
}
