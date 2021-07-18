package store

import (
	"encoding/json"
	"os"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/google/uuid"
)

// Store is used to persist machine information
type Store interface {
	Get() ([]byte, error)
	Store([]byte) error
}

var (
	_         Store = (*FileStore)(nil)
	fileStore *FileStore
)

type FileStore struct {
}

// Instance returns the active instance of the file store
func Instance() *FileStore {
	if fileStore != nil {
		return fileStore
	}
	fileStore = &FileStore{}
	return fileStore
}

// Get reads a JSON file and returns the data
func (s *FileStore) Get() ([]byte, error) {
	s.createFileIfNotExists(consts.AGENT_STORE_FILENAME)
	file, err := os.ReadFile(consts.AGENT_STORE_FILENAME)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Store writes the desired JSON to a JSON file
func (s *FileStore) Store(data []byte) error {
	s.createFileIfNotExists(consts.AGENT_STORE_FILENAME)
	return os.WriteFile(consts.AGENT_STORE_FILENAME, data, 0644)
}

// createFileIfNotExists is a handy method which creates a given file if it does not exist
func (s *FileStore) createFileIfNotExists(fileName string) error {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
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
