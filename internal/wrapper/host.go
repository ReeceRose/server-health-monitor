package wrapper

import (
	"github.com/PR-Developers/server-health-monitor/internal/types"

	"github.com/shirou/gopsutil/v3/host"
)

// HostInformation is an interface which provides method signatures for fetching host information
type HostInformation interface {
	Info() (*types.Host, error)
}

// GopsHost is the implementation of gopsutil host
type GopsHost struct {
}

var (
	_ HostInformation = (*GopsHost)(nil)
)

// Info returns host information
func (h *GopsHost) Info() (*types.Host, error) {
	host, err := host.Info()
	return (*types.Host)(host), err
}
