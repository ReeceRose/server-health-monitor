package host

import (
	"github.com/PR-Developers/server-health-monitor/data-collector/internal/logger"
	"github.com/shirou/gopsutil/v3/host"
)

func GetInfo() *host.InfoStat {
	host, err := host.Info()
	if err != nil {
		logger.Instance().Error(err.Error())
		return nil
	}
	return host
}
