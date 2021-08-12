package service

import "github.com/PR-Developers/server-health-monitor/internal/types"

// IHealthService is an interface which provides method signatures for a health service
type IHealthService interface {
	GetHealth(requestID string) types.StandardResponse
	GetHealthByAgentID(requestID, agentID string) types.StandardResponse
	AddHealth(requestID string, agentID string, data *types.Health) types.StandardResponse
	GetLatestHealthDataByAgentID(requestID string, agentID string, time int64) []types.Health
}

// IHostService is an interface which provides method signatures for a host service
type IHostService interface {
	GetHosts(requestID string, includeHealthData bool) types.StandardResponse
	GetHostByID(requestID, agentID string, includeHealthData bool) types.StandardResponse
	AddHost(requestID string, agentID string, data *types.Host) types.StandardResponse
	GetLatestHealthDataForAgents(requestID string, lastCheck int64) types.StandardResponse
	isHostOnline(requestID string, agentID string) bool
}
