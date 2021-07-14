package service

import "github.com/PR-Developers/server-health-monitor/internal/types"

type HealthService struct {
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

// TODO: return all health data not just host info.
func (s *HealthService) GetHealth() *types.Host {
	// TODO: return actual health data
	return &types.Host{}
}

// TODO: return all health data not just host info.
func (s *HealthService) GetHealthByServerId() *types.Host {
	// TODO: return actual health data
	return &types.Host{}
}

func (s *HealthService) AddHealth() bool {
	// TODO: add halth data
	return true
}
