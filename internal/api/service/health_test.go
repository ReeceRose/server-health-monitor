package service

import (
	"testing"

	"github.com/PR-Developers/server-health-monitor/internal/api/service/mocks"
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"github.com/stretchr/testify/mock"
)

//go:generate mockery --dir=../../ -r --name IHealthRepository

type testServiceHelper struct {
	service *HealthService
	repo    repository.IHealthRepository
	mock    *mock.Mock
}

func getInitializedHealthService() testServiceHelper {
	repo := new(mocks.IHealthRepository)
	// repo.On
	service := NewHealthService(repo)

	return testServiceHelper{
		service: service,
		repo:    repo,
		mock:    &repo.Mock,
	}
}

func TestHealth_GetHealth_ReturnsExpectedNumberOfHealthData(t *testing.T) {
	// helper := getInitializedHealthService()
	// helper.mock.On("")
}
