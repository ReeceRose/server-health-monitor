package service

import "github.com/PR-Developers/server-health-monitor/internal/types"

// IHealthService is an interface which provides method signatures for a health service
type IHealthService interface {
	GetHealth(requestID string) types.StandardResponse
	GetHealthByAgentID(requestID, agentID string) types.StandardResponse
	AddHealth(requestID string, agentID string, data *types.Health) types.StandardResponse
	GetLatestHealthDataByAgentID(requestID string, agentID string, delay int) types.StandardResponse
}

// IHostService is an interface which provides method signatures for a host service
type IHostService interface {
	GetHosts(requestID string) types.StandardResponse
	GetHostByID(requestID, agentID string) types.StandardResponse
	AddHost(requestID string, agentID string, data *types.Host) types.StandardResponse
	isHostOnline(requestID string, agentID string) bool
}
