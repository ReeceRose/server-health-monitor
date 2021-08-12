package service

import "github.com/PR-Developers/server-health-monitor/internal/types"

// IHealthService is an interface which provides method signatures for a health service
type IHealthService interface {
	GetHealth(requestID string) types.HealthReponse
	GetHealthByAgentID(requestID, agentID string) types.HealthReponse
	AddHealth(requestID string, agentID string, data *types.Health) types.HealthReponse
	GetLatestHealthDataByAgentID(requestID string, agentID string, time int64) types.HealthReponse
	GetLatestHealthDataForAgents(requestID string, time int64) types.HostReponse
}

// IHostService is an interface which provides method signatures for a host service
type IHostService interface {
	GetHosts(requestID string, includeHealthData bool) types.HostReponse
	GetHostByID(requestID, agentID string, includeHealthData bool) types.HostReponse
	AddHost(requestID string, agentID string, data *types.Host) types.HostReponse
	isHostOnline(requestID string, agentID string) bool
}
