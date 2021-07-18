package wrapper

import (
	"github.com/PR-Developers/server-health-monitor/internal/types"

	"github.com/shirou/gopsutil/v3/host"
)

type GopsHost struct {
}

var (
	_ types.HostInformation = (*GopsHost)(nil)
)

func (h *GopsHost) Info() (*types.Host, error) {
	host, err := host.Info()
	return (*types.Host)(host), err
}
