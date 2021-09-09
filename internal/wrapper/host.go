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
	return &types.Host{
		Hostname:             host.Hostname,
		Uptime:               host.Uptime,
		BootTime:             host.BootTime,
		Procs:                host.Procs,
		OS:                   host.OS,
		Platform:             host.Platform,
		PlatformFamily:       host.PlatformFamily,
		PlatformVersion:      host.PlatformVersion,
		KernelVersion:        host.KernelVersion,
		KernelArch:           host.KernelArch,
		VirtualizationRole:   host.VirtualizationRole,
		VirtualizationSystem: host.VirtualizationSystem,
		HostID:               host.HostID,
	}, err
}
