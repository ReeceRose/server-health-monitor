package host

import (
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/types"

	"github.com/shirou/gopsutil/v3/host"
)

// GetInfo returns all host information
func GetInfo() *types.Host {
	host, err := host.Info()
	if err != nil {
		logger.Instance().Error(err.Error())
		return nil
	}
	return (*types.Host)(host)
}
