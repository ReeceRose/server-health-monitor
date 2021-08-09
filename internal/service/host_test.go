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
)

//go:generate mockery --dir=../ -r --name IHealthRepository
//go:generate mockery --dir=../ -r --name IHostRepository

type testHostServiceHelper struct {
	hostService   IHostService
	healthService IHealthService
	hostRepo      repository.IHostRepository
	healthRepo    repository.IHealthRepository
	mock          *mock.Mock
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
	healthService := NewHealthService(healthRepo)
	hostService := NewHostService(hostRepo, healthRepo)
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
		healthRepo:    healthRepo,
		mock:          &hostRepo.Mock,
		healthMock:    healthMock,
	}
}

func TestHost_GetHosts_ReturnsExpectedHostData(t *testing.T) {
	helper := getInitializedHostService()
	helper.mock.On("Find", bson.M{}).Return(hostData, nil)

	response := helper.hostService.GetHosts("1")

	data := response.Data.([]types.Host)

	assert.Equal(t, 2, len(data))
	assert.Equal(t, int64(1), data[0].CreateTime)
	assert.Equal(t, uint64(10), data[0].Uptime)
	assert.Equal(t, int64(2), data[1].CreateTime)
	assert.Equal(t, uint64(20), data[1].Uptime)

	helper.mock.AssertExpectations(t)
}

func TestHost_GetHosts_HandlesError(t *testing.T) {
	helper := getInitializedHostService()
	helper.mock.On("Find", bson.M{}).Return(nil, fmt.Errorf("failed to fetch host data"))

	response := helper.hostService.GetHosts("1")

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, response.Data, []types.Host{})
	assert.False(t, response.Success)
	assert.Equal(t, response.Error, "failed to get all hosts data - Request ID: 1")

	helper.mock.AssertExpectations(t)
}

func TestHost_GetHostByID_ReturnsExpectedHostData(t *testing.T) {
	helper := getInitializedHostService()
	helper.mock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{hostData[0]}, nil)

	response := helper.hostService.GetHostByID("1", "1")

	data := response.Data.([]types.Host)

	assert.Equal(t, 1, len(data))
	assert.Equal(t, int64(1), data[0].CreateTime)
	assert.Equal(t, uint64(10), data[0].Uptime)
	assert.Equal(t, "test machine 1", data[0].Hostname)

	helper.mock.AssertExpectations(t)
}

func TestHost_GetHostByID_HandlesDatabaseError(t *testing.T) {
	helper := getInitializedHostService()
	helper.mock.On("Find", bson.M{"agentID": "100"}).Return(nil, fmt.Errorf("failed to fetch host data"))

	response := helper.hostService.GetHostByID("1", "100")

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, response.Data, []types.Host{})
	assert.False(t, response.Success)
	assert.Equal(t, response.Error, "failed to get host data for agent: 100 - Request ID: 1")

	helper.mock.AssertExpectations(t)
}

func TestHost_GetHostByID_HandlesNoHostsError(t *testing.T) {
	helper := getInitializedHostService()
	helper.mock.On("Find", bson.M{"agentID": "100"}).Return(nil, nil)

	response := helper.hostService.GetHostByID("1", "100")

	assert.Equal(t, 204, response.StatusCode)
	assert.Equal(t, response.Data, []types.Host{})
	assert.False(t, response.Success)
	assert.Equal(t, response.Error, "failed to get host data for agent: 100 - Request ID: 1")

	helper.mock.AssertExpectations(t)
}

func TestHost_AddHost_InsertsHost(t *testing.T) {
	helper := getInitializedHostService()
	helper.mock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{}, nil)
	helper.mock.On("Insert", mock.Anything).Return("123", nil)

	response := helper.hostService.AddHost("1", "1", &hostData[0])

	data := response.Data.(*types.Host)

	assert.Equal(t, uint64(10), data.Uptime)
	assert.Equal(t, "test machine 1", data.Hostname)
	assert.Equal(t, "1", data.AgentID)

	helper.mock.AssertExpectations(t)
}

func TestHost_AddHost_HandlesUpdateExistingHostError(t *testing.T) {
	helper := getInitializedHostService()
	helper.mock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{hostData[0]}, nil)
	helper.mock.On("UpdateByID", mock.Anything).Return(nil)
	response := helper.hostService.AddHost("1", "1", &hostData[0])

	assert.Equal(t, 200, response.StatusCode)
	assert.True(t, response.Success)

	helper.mock.AssertExpectations(t)
}

func TestHost_AddHost_HandlesFailedToUpdateExistingHostError(t *testing.T) {
	helper := getInitializedHostService()
	helper.mock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{hostData[0]}, nil)
	helper.mock.On("UpdateByID", mock.Anything).Return(fmt.Errorf("failed to update data"))

	response := helper.hostService.AddHost("1", "1", &hostData[0])

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, response.Data, []types.Host{})
	assert.False(t, response.Success)
	assert.Equal(t, response.Error, "failed to update data for agent: 1 - Request ID 1")

	helper.mock.AssertExpectations(t)
}

func TestHost_AddHost_HandlesFailedToInsertHostError(t *testing.T) {
	helper := getInitializedHostService()
	helper.mock.On("Find", bson.M{"agentID": "1"}).Return([]types.Host{}, nil)
	helper.mock.On("Insert", mock.Anything).Return("", fmt.Errorf("failed to insert data"))

	response := helper.hostService.AddHost("1", "1", &hostData[1])

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, response.Data, []types.Host{})
	assert.False(t, response.Success)
	assert.Equal(t, response.Error, "failed to insert data for agent: 1 - Request ID 1")

	helper.mock.AssertExpectations(t)
}

func TestHost_IsHostOnline_HostIsOnline(t *testing.T) {
	helper := getInitializedHostService()

	helper.healthMock.On("Find", mock.Anything).Return([]types.Health{
		{
			CreateTime: time.Now().UnixNano(),
		},
	}, nil)

	response := helper.hostService.isHostOnline(&hostData[0])

	assert.True(t, response)

	helper.mock.AssertExpectations(t)
	helper.healthMock.AssertExpectations(t)
}

func TestHost_IsHostOnline_HostIsOffline(t *testing.T) {
	helper := getInitializedHostService()

	response := helper.hostService.isHostOnline(&hostData[0])

	assert.True(t, response)

	helper.mock.AssertExpectations(t)
	helper.healthMock.AssertExpectations(t)
}

func TestHost_IsHostOnline_HostIsOfflineWhenErrorOccurs(t *testing.T) {
	hostRepo := new(mocks.IHostRepository)
	healthRepo := new(mocks.IHealthRepository)
	hostService := NewHostService(hostRepo, healthRepo)
	healthMock := &healthRepo.Mock

	healthMock.On("Find", mock.Anything).Return(nil, fmt.Errorf("failed to get data"))

	response := hostService.isHostOnline(&hostData[0])

	assert.False(t, response)

	healthMock.AssertExpectations(t)
}
