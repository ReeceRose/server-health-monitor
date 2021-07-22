package service

import "github.com/PR-Developers/server-health-monitor/internal/types"

// IHealthService is an interface which provides method signatures for a health service
type IHealthService interface {
	GetHealth(requestID string) types.StandardResponse
	GetHealthByAgentID(requestID, agentID string) types.StandardResponse
	AddHealth(requestID string, agentID string, data *types.Health) types.StandardResponse
}
