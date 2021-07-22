package store

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/data-collector/store/mocks"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/wrapper"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

//go:generate mockery --dir=../../ -r --name OperatingSystem

func resetStore() {
	fileStore = nil
}

func TestStoreInstanceInitializesStore(t *testing.T) {
	resetStore()

	assert.Nil(t, fileStore)
	store := Instance(&wrapper.DefaultOS{})
	assert.NotNil(t, store)
	assert.NotNil(t, fileStore)
	store = Instance(&wrapper.DefaultOS{})
	assert.NotNil(t, store)
}

func TestStoreGetReturnsCorrectData(t *testing.T) {
	resetStore()

	os.WriteFile(consts.AGENT_STORE_FILENAME, []byte("test data"), 0644)

	store := Instance(&wrapper.DefaultOS{})
	data, err := store.Get()

	assert.Nil(t, err)
	assert.Equal(t, []byte("test data"), data)

	os.Remove(consts.AGENT_STORE_FILENAME)
}

func TestStoreStoreWritesCorrectData(t *testing.T) {
	resetStore()

	store := Instance(&wrapper.DefaultOS{})
	err := store.Store([]byte("test data"))
	assert.Nil(t, err)

	data, err := os.ReadFile(consts.AGENT_STORE_FILENAME)
	assert.Equal(t, []byte("test data"), data)
	assert.Nil(t, err)

	os.Remove(consts.AGENT_STORE_FILENAME)
}

func TestStoreCreatesFileWhenNotExists(t *testing.T) {
	resetStore()

	file, err := os.Stat(consts.AGENT_STORE_FILENAME)

	assert.Nil(t, file)
	assert.NotNil(t, err)
	assert.True(t, os.IsNotExist(err))

	Instance(&wrapper.DefaultOS{})
	file, err = os.Stat(consts.AGENT_STORE_FILENAME)
	assert.NotNil(t, file)
	assert.Nil(t, err)
}

func TestStoreHandlesFileCreationError(t *testing.T) {
	resetStore()

	wrapper := new(mocks.OperatingSystem)

	osErr := fmt.Errorf("is not exists")
	wrapper.On("Stat", consts.AGENT_STORE_FILENAME).Return(nil, osErr)
	wrapper.On("IsNotExist", osErr).Return(true)
	wrapper.On("OpenFile",
		consts.AGENT_STORE_FILENAME, os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		fs.FileMode(0644),
	).Return(nil, fmt.Errorf("failed to create file"))

	store := FileStore{
		osWrapper: wrapper,
	}

	err := store.createFileIfNotExists(consts.AGENT_STORE_FILENAME)

	assert.NotNil(t, err)
	assert.Equal(t, "failed to create file", err.Error())
	wrapper.AssertExpectations(t)
}

func TestGetAgentInformationGeneratesDataWhenFileEmpty(t *testing.T) {
	resetStore()

	store := Instance(&wrapper.DefaultOS{})
	data := store.GetAgentInformation()

	rawData, err := store.Get()

	assert.NotNil(t, data)
	assert.Contains(t, string(rawData), "ID")
	assert.Nil(t, err)
}

func TestGetAgentInformationReadsIDFromFile(t *testing.T) {
	resetStore()

	agentInformation := types.AgentInformation{}
	agentInformation.ID = uuid.New()

	data, _ := json.Marshal(agentInformation)

	store := Instance(&wrapper.DefaultOS{})
	store.Store(data)

	agent := store.GetAgentInformation()

	assert.NotNil(t, data)
	assert.Equal(t, agentInformation.ID, agent.ID)
}

func TestGetAgentInformationReturnsEmptyAgentWhenError(t *testing.T) {
	resetStore()

	wrapper := new(mocks.OperatingSystem)
	wrapper.On("ReadFile", consts.AGENT_STORE_FILENAME).Return(nil, fmt.Errorf("failed to read data"))

	store := FileStore{
		osWrapper: wrapper,
	}

	agent := store.GetAgentInformation()

	assert.Equal(t, types.AgentInformation{}, agent)
	wrapper.AssertExpectations(t)
}
