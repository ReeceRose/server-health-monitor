package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/service/mocks"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockery --dir=../ -r --name IHealthService
//go:generate mockery --dir=../ -r --name IHostRepository

type testHostServiceHelper struct {
	hostService   IHostService
	healthService IHealthService
	hostRepo      repository.IHostRepository
	hostMock      *mock.Mock
	healthMock    *mock.Mock
}

var (
	hostData []types.Host = []types.Host{
		{
			AgentID:    "1",
			CreateTime: 1,
			Uptime:     10,
			Hostname:   "test machine 1",
			Online:     true,
		},
		{
			AgentID:    "2",
			CreateTime: 2,
			Uptime:     20,
			Hostname:   "test machine 2",
			Online:     false,
		},
	}
)

func getInitializedHostService() testHostServiceHelper {
	hostRepo := new(mocks.IHostRepository)
	healthRepo := new(mocks.IHealthRepository)
	healthService := NewHealthService(healthRepo, hostRepo)
	hostService := NewHostService(hostRepo, healthService)
	healthMock := &healthRepo.Mock
	healthMock.On("Find", mock.Anything).Return([]types.Health{
		{
			CreateTime: time.Now().Add(-time.Hour).UnixNano(),
		},
	}, nil)

	return testHostServiceHelper{
		healthService: healthService,
		hostService:   hostService,
		hostRepo:      hostRepo,
		hostMock:      &hostRepo.Mock,
		healthMock:    healthMock,
	}
}

func TestHost_GetHosts_ReturnsExpectedHostData(t *testing.T) {
	helper := getInitializedHostService()
	helper.hostMock.On("Find", bson.M{}).Return(hostData, nil)

	res := helper.hostService.GetHosts("1", false)

	data := res.Data

	assert.Equal(t, 2, len(data))
	assert.Equal(t, 0, len(data[0].Health))
	assert.Equal(t, int64(1), data[0].CreateTime)
	assert.Equal(t, uint64(10), data[0].Uptime)
	assert.Equal(t, int64(2), data[1].CreateTime)
	assert.Equal(t, uint64(20), data[1].Uptime)

	helper.hostMock.AssertExpectations(t)
}

func TestHost_GetHosts_HandlesError(t *testing.T) {
	helper := getInitializedHostService()
	helper.hostMock.On("Find", bson.M{}).Return(nil, fmt.Errorf("failed to fetch host data"))

	res := helper.hostService.GetHosts("1", false)

	assert.Equal(t, 500, res.StatusCode)
	assert.Equal(t, []types.Host{}, res.Data)
	assert.False(t, res.Success)
	assert.Equal(t, "failed to get all hosts data - Request ID: 1", res.Error)

	helper.hostMock.AssertExpectations(t)
}

func TestHost_GetHosts_IncludesHealthData(t *testing.T) {
	hostRepo := new(mocks.IHostRepository)
	healthRepo := new(mocks.IHealthRepository)
	healthService := NewHealthService(healthRepo, hostRepo)
	hostService := NewHostService(hostRepo, healthService)
	hostMock := &hostRepo.Mock
	healthMock := &healthRepo.Mock

	hostMock.On("Find", bson.M{}).Return(hostData, nil)
	healthMock.On("FindWithFilter", mock.Anything, mock.Anything).Return([]types.Health{
		{
			CreateTime: time.Now().UnixNano(),
			AgentID:    "1",
		},
	}, nil)
	res := hostService.GetHosts("1", true)

	data := res.Data

	assert.Equal(t, 2, len(data))
	assert.Equal(t, 1, len(data[0].Health))
	assert.Equal(t, int64(1), data[0].CreateTime)
	assert.Equal(t, uint64(10), data[0].Uptime)
	assert.Equal(t, int64(2), data[1].CreateTime)
	assert.Equal(t, uint64(20), data[1].Uptime)

	hostMock.AssertExpectations(t)
	healthMock.AssertExpectations(t)
}

func TestHost_GetHostByID_ReturnsExpectedHostData(t *testing.T) {
	helper := getInitializedHostService()
	helper.hostMock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{hostData[0]}, nil)

	res := helper.hostService.GetHostByID("1", "1", false)

	data := res.Data

	assert.Equal(t, 1, len(data))
	assert.Equal(t, int64(1), data[0].CreateTime)
	assert.Equal(t, uint64(10), data[0].Uptime)
	assert.Equal(t, "test machine 1", data[0].Hostname)

	helper.hostMock.AssertExpectations(t)
}

func TestHost_GetHostByID_IncludesHealthData(t *testing.T) {
	hostRepo := new(mocks.IHostRepository)
	healthRepo := new(mocks.IHealthRepository)
	healthService := NewHealthService(healthRepo, hostRepo)
	hostService := NewHostService(hostRepo, healthService)
	hostMock := &hostRepo.Mock
	healthMock := &healthRepo.Mock

	hostMock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{hostData[0]}, nil)
	healthMock.On("FindWithFilter", mock.Anything, mock.Anything).Return([]types.Health{
		{
			CreateTime: time.Now().UnixNano(),
			AgentID:    "1",
		},
	}, nil)
	res := hostService.GetHostByID("1", "1", true)

	data := res.Data

	assert.Equal(t, 1, len(data))
	assert.Equal(t, 1, len(data[0].Health))
	assert.Equal(t, int64(1), data[0].CreateTime)
	assert.Equal(t, uint64(10), data[0].Uptime)
	assert.Equal(t, "test machine 1", data[0].Hostname)

	hostMock.AssertExpectations(t)
	healthMock.AssertExpectations(t)
}

func TestHost_GetHostByID_HandlesDatabaseError(t *testing.T) {
	helper := getInitializedHostService()
	helper.hostMock.On("Find", bson.M{"agentID": "100"}).Return(nil, fmt.Errorf("failed to fetch host data"))

	res := helper.hostService.GetHostByID("1", "100", false)

	assert.Equal(t, 500, res.StatusCode)
	assert.Equal(t, []types.Host{}, res.Data)
	assert.False(t, res.Success)
	assert.Equal(t, "failed to get host data for agent: 100 - Request ID: 1", res.Error)

	helper.hostMock.AssertExpectations(t)
}

func TestHost_GetHostByID_HandlesNoHostsError(t *testing.T) {
	helper := getInitializedHostService()
	helper.hostMock.On("Find", bson.M{"agentID": "100"}).Return(nil, nil)

	res := helper.hostService.GetHostByID("1", "100", false)

	assert.Equal(t, 204, res.StatusCode)
	assert.Equal(t, []types.Host{}, res.Data)
	assert.False(t, res.Success)
	assert.Equal(t, "failed to get host data for agent: 100 - Request ID: 1", res.Error)

	helper.hostMock.AssertExpectations(t)
}

func TestHost_AddHost_InsertsHost(t *testing.T) {
	helper := getInitializedHostService()
	helper.hostMock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{}, nil)
	helper.hostMock.On("Insert", mock.Anything).Return("123", nil)

	res := helper.hostService.AddHost("1", "1", &hostData[0])

	data := res.Data[0]

	assert.Equal(t, uint64(10), data.Uptime)
	assert.Equal(t, "test machine 1", data.Hostname)
	assert.Equal(t, "1", data.AgentID)

	helper.hostMock.AssertExpectations(t)
}

func TestHost_AddHost_HandlesUpdateExistingHostError(t *testing.T) {
	helper := getInitializedHostService()
	helper.hostMock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{hostData[0]}, nil)
	helper.hostMock.On("UpdateByID", mock.Anything).Return(nil)
	res := helper.hostService.AddHost("1", "1", &hostData[0])

	assert.Equal(t, 200, res.StatusCode)
	assert.True(t, res.Success)

	helper.hostMock.AssertExpectations(t)
}

func TestHost_AddHost_HandlesFailedToUpdateExistingHostError(t *testing.T) {
	helper := getInitializedHostService()
	helper.hostMock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{hostData[0]}, nil)
	helper.hostMock.On("UpdateByID", mock.Anything).Return(fmt.Errorf("failed to update data"))

	res := helper.hostService.AddHost("1", "1", &hostData[0])

	assert.Equal(t, 500, res.StatusCode)
	assert.Equal(t, []types.Host{}, res.Data)
	assert.False(t, res.Success)
	assert.Equal(t, "failed to update data for agent: 1 - Request ID 1", res.Error)

	helper.hostMock.AssertExpectations(t)
}

func TestHost_AddHost_HandlesFailedToInsertHostError(t *testing.T) {
	helper := getInitializedHostService()
	helper.hostMock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{}, nil)
	helper.hostMock.On("Insert", mock.Anything).Return("", fmt.Errorf("failed to insert data"))

	res := helper.hostService.AddHost("1", "1", &hostData[1])

	assert.Equal(t, 500, res.StatusCode)
	assert.Equal(t, []types.Host{}, res.Data)
	assert.False(t, res.Success)
	assert.Equal(t, "failed to insert data for agent: 1 - Request ID 1", res.Error)

	helper.hostMock.AssertExpectations(t)
}

func TestHost_IsHostOnline_HostIsOnline(t *testing.T) {
	hostRepo := new(mocks.IHostRepository)
	healthRepo := new(mocks.IHealthRepository)
	healthService := NewHealthService(healthRepo, hostRepo)
	hostService := NewHostService(hostRepo, healthService)
	healthMock := &healthRepo.Mock

	healthMock.On("FindWithFilter", mock.Anything, mock.Anything).Return([]types.Health{
		{
			CreateTime: time.Now().UnixNano(),
		},
	}, nil)

	res := hostService.isHostOnline("1", hostData[0].AgentID)

	assert.True(t, res)

	healthMock.AssertExpectations(t)
}

func TestHost_IsHostOnline_HostIsOffline(t *testing.T) {
	hostRepo := new(mocks.IHostRepository)
	healthRepo := new(mocks.IHealthRepository)
	healthService := NewHealthService(healthRepo, hostRepo)
	hostService := NewHostService(hostRepo, healthService)
	healthMock := &healthRepo.Mock
	healthMock.On("FindWithFilter", mock.Anything, mock.Anything).Return([]types.Health{}, nil)

	res := hostService.isHostOnline("1", hostData[0].AgentID)

	assert.False(t, res)

	healthMock.AssertExpectations(t)
}

func TestHost_GetHealthDataForHosts_GetsDataForActiveHost(t *testing.T) {
	helper := getInitializedHostService()
	helper.healthMock.On("FindWithFilter", mock.Anything, mock.Anything).Return([]types.Health{
		{
			CreateTime: 123,
			AgentID:    "1",
		},
	}, nil)
	hosts := hostData

	helper.hostService.getHealthDataForHosts("1", &hosts)

	assert.Equal(t, 1, len(hosts[0].Health))
	assert.NotEqual(t, 0, hosts[0].LastConnected)
}

func TestHost_GetHealthDataForHosts_GetsDataForInactiveHost(t *testing.T) {
	hostRepo := new(mocks.IHostRepository)
	healthRepo := new(mocks.IHealthRepository)
	healthService := NewHealthService(healthRepo, hostRepo)
	hostService := NewHostService(hostRepo, healthService)
	healthMock := &healthRepo.Mock

	options := options.Find()
	options.SetSort(bson.D{primitive.E{Key: "createTime", Value: -1}})
	options.SetLimit(1)

	healthMock.On("FindWithFilter", bson.M{"agentID": "1"}, options).Return([]types.Health{
		{
			CreateTime: 123,
			AgentID:    "1",
		},
	}, nil)
	healthMock.On("FindWithFilter", mock.Anything, mock.Anything).Return([]types.Health{}, nil)
	hosts := hostData

	hostService.getHealthDataForHosts("1", &hosts)

	assert.Equal(t, 0, len(hosts[0].Health))
	assert.Equal(t, int64(123), hosts[0].LastConnected)
}
