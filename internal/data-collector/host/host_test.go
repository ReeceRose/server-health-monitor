package host

import (
	"fmt"
	"testing"

	"github.com/PR-Developers/server-health-monitor/internal/data-collector/host/mocks"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/stretchr/testify/assert"
)

//go:generate mockery --dir=../../ -r --name HostInformation

func TestHost_GetInfo_ReturnsExpectedHostInformation(t *testing.T) {
	wrapper := new(mocks.HostInformation)
	hostInfo := &types.Host{
		Hostname: "test",
		Platform: "ubuntu",
		OS:       "linux",
	}
	wrapper.On("Info").Return(hostInfo, nil)

	hostWrapper = wrapper

	host := GetInfo()

	assert.Equal(t, hostInfo, host)
	wrapper.AssertExpectations(t)
}

func TestHost_GetInfo_HandlesError(t *testing.T) {
	wrapper := new(mocks.HostInformation)
	wrapper.On("Info").Return(nil, fmt.Errorf("platform not supported"))

	hostWrapper = wrapper

	host := GetInfo()

	assert.Nil(t, host)
	wrapper.AssertExpectations(t)
}
