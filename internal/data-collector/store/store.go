package store

import (
	"encoding/json"
	"os"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/wrapper"
	"github.com/google/uuid"
)

// Store is an interface which provides method signatures for storing persistent information
type Store interface {
	// Core
	Get() ([]byte, error)
	Store([]byte) error

	// Custom
	GetAgentInformation() types.AgentInformation
}

var (
	_         Store = (*FileStore)(nil)
	fileStore *FileStore
)

// FileStore is a filestore implementation of a store
type FileStore struct {
	osWrapper wrapper.OperatingSystem
}

// Instance returns the active instance of the file store
func Instance(wrapper wrapper.OperatingSystem) Store {
	if fileStore != nil {
		return fileStore
	}

	fileStore = &FileStore{
		osWrapper: wrapper,
	}
	fileStore.createFileIfNotExists(consts.AGENT_STORE_FILENAME)
	return fileStore
}

// Get reads a JSON file and returns the data
func (s *FileStore) Get() ([]byte, error) {
	return s.osWrapper.ReadFile(consts.AGENT_STORE_FILENAME)
}

// Store writes the desired JSON to a JSON file
func (s *FileStore) Store(data []byte) error {
	return s.osWrapper.WriteFile(consts.AGENT_STORE_FILENAME, data, 0644)
}

// createFileIfNotExists is a handy method which creates a given file if it does not exist
func (s *FileStore) createFileIfNotExists(fileName string) error {
	if _, err := s.osWrapper.Stat(fileName); s.osWrapper.IsNotExist(err) {
		file, err := s.osWrapper.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

// GetAgentInformation pulls agent information out of the current store
func (s *FileStore) GetAgentInformation() types.AgentInformation {
	agentData, err := s.Get()
	if err != nil {
		return types.AgentInformation{}
	}

	var agentInformation types.AgentInformation
	json.Unmarshal(agentData, &agentInformation)
	if agentInformation.ID.String() == "00000000-0000-0000-0000-000000000000" {
		agentInformation = types.AgentInformation{}
		agentInformation.ID = uuid.New()
		data, _ := json.Marshal(agentInformation)
		s.Store(data)
	}
	return agentInformation
}
