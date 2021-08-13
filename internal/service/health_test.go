package service

import (
	"fmt"
	"testing"

	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/PR-Developers/server-health-monitor/internal/service/mocks"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

//go:generate mockery --dir=../ -r --name IHealthRepository
//go:generate mockery --dir=../ -r --name IHostRepository

type testHealthServiceHelper struct {
	healthService IHealthService
	healthRepo    repository.IHealthRepository
	hostRepo      repository.IHostRepository
	mock          *mock.Mock
}

var (
	healthData []types.Health = []types.Health{
		{
			AgentID:    "1",
			CreateTime: 1,
			Uptime:     10,
		},
		{
			AgentID:    "2",
			CreateTime: 2,
			Uptime:     20,
		},
	}
)

func getInitializedHealthService() testHealthServiceHelper {
	healthRepo := new(mocks.IHealthRepository)
	hostRepo := new(mocks.IHostRepository)
	// repo.On
	healthService := NewHealthService(healthRepo, hostRepo)

	return testHealthServiceHelper{
		healthService: healthService,
		hostRepo:      hostRepo,
		healthRepo:    healthRepo,
		mock:          &healthRepo.Mock,
	}
}

func TestHealth_GetHealth_ReturnsExpectedHealthData(t *testing.T) {
	helper := getInitializedHealthService()
	helper.mock.On("Find", bson.M{}).Return(healthData, nil)

	res := helper.healthService.GetHealth("1")

	data := res.Data

	assert.Equal(t, 2, len(data))
	assert.Equal(t, int64(1), data[0].CreateTime)
	assert.Equal(t, uint64(10), data[0].Uptime)
	assert.Equal(t, int64(2), data[1].CreateTime)
	assert.Equal(t, uint64(20), data[1].Uptime)

	helper.mock.AssertExpectations(t)
}

func TestHealth_GetHealth_HandlesError(t *testing.T) {
	helper := getInitializedHealthService()
	helper.mock.On("Find", bson.M{}).Return(nil, fmt.Errorf("failed to get data from DB"))

	res := helper.healthService.GetHealth("1")

	assert.Equal(t, res.Data, []types.Health{})
	assert.Equal(t, 500, res.StatusCode)
	assert.Equal(t, "failed to get all health data - Request ID: 1", res.Error)
	assert.False(t, res.Success)

	helper.mock.AssertExpectations(t)
}

func TestHealth_GetHealthByAgentId_ReturnsExpectedHealthData(t *testing.T) {
	helper := getInitializedHealthService()
	helper.mock.On("Find", bson.M{"agentID": "1"}).Return([]types.Health{healthData[0]}, nil)

	res := helper.healthService.GetHealthByAgentID("1", "1")

	data := res.Data

	assert.Equal(t, 1, len(data))

	helper.mock.AssertExpectations(t)
}

func TestHealth_GetHealthByAgentId_HandlesError(t *testing.T) {
	helper := getInitializedHealthService()
	helper.mock.On("Find", bson.M{"agentID": "4"}).Return(nil, fmt.Errorf("failed to get data from DB"))

	res := helper.healthService.GetHealthByAgentID("1", "4")

	assert.Equal(t, res.Data, []types.Health{})
	assert.Equal(t, 500, res.StatusCode)
	assert.Equal(t, "failed to get data for agent: 4 - Request ID: 1", res.Error)
	assert.False(t, res.Success)

	helper.mock.AssertExpectations(t)
}

func TestHealth_AddHealth_AddsExpectedHealthData(t *testing.T) {
	helper := getInitializedHealthService()
	helper.mock.On("Insert", &healthData[0]).Return("1234567", nil)

	res := helper.healthService.AddHealth("1", "1", &healthData[0])

	data := res.Data

	assert.True(t, res.Success)
	assert.NotEmpty(t, data[0].ID)

	helper.mock.AssertExpectations(t)
}

func TestHealth_AddHealth_HandlesError(t *testing.T) {
	helper := getInitializedHealthService()
	helper.mock.On("Insert", &healthData[1]).Return("", fmt.Errorf("failed to insert data into DB"))

	res := helper.healthService.AddHealth("1", "2", &healthData[1])

	assert.Equal(t, res.Data, []types.Health{})
	assert.Equal(t, "failed to insert data for agent: 2 - Request ID 1", res.Error)
	assert.False(t, res.Success)

	helper.mock.AssertExpectations(t)
}

func TestHealth_GetLatestHealthDataByAgentID_SortsDataDescending(t *testing.T) {
	healthRepo := new(mocks.IHealthRepository)
	hostRepo := new(mocks.IHostRepository)
	healthService := NewHealthService(healthRepo, hostRepo)
	healthMock := &healthRepo.Mock

	healthMock.On("FindWithFilter", mock.Anything, mock.Anything).Return([]types.Health{
		{
			AgentID:    "1",
			CreateTime: 1,
		},
		{
			AgentID:    "1",
			CreateTime: 500,
		},
		{
			AgentID:    "1",
			CreateTime: 1000,
		},
	}, nil)

	res := healthService.GetLatestHealthDataByAgentID("1", hostData[0].AgentID, 2)
	data := res.Data

	assert.Equal(t, 3, len(data))
	assert.Equal(t, int64(1000), data[0].CreateTime)
	assert.Equal(t, int64(1), data[2].CreateTime)

	healthMock.AssertExpectations(t)
}

func TestHealth_GetLatestHealthDataByAgentID_HandlesError(t *testing.T) {
	healthRepo := new(mocks.IHealthRepository)
	hostRepo := new(mocks.IHostRepository)
	healthService := NewHealthService(healthRepo, hostRepo)
	healthMock := &healthRepo.Mock

	healthMock.On("FindWithFilter", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("failed to get data"))

	res := healthService.GetLatestHealthDataByAgentID("1", hostData[0].AgentID, 2)

	assert.Equal(t, res.Data, []types.Health{})
	assert.Equal(t, "failed to get health data for agent: 1 - Request ID: 1", res.Error)
	assert.False(t, res.Success)

	healthMock.AssertExpectations(t)
}

func TestHealth_GetLatestHealthDataForAgents_ReturnsExpectedData(t *testing.T) {

}

func TestHealth_GetLatestHealthDataForAgents_HandlesError(t *testing.T) {

}

func TestHealth_GetHealthForAgentWithOptions_ReturnsExpectedData(t *testing.T) {

}

func TestHealth_GetHealthForAgentWithOptions_HandlesError(t *testing.T) {

}
