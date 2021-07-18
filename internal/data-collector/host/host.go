package host

import (
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/wrapper"
)

var (
	hostWrapper types.HostInformation = &wrapper.GopsHost{}
)

// GetInfo returns all host information
func GetInfo() *types.Host {
	host, err := hostWrapper.Info()
	if err != nil {
		logger.Instance().Error(err.Error())
		return nil
	}
	return host
}
