package repository

import (
	"github.com/PR-Developers/server-health-monitor/internal/database"
	"github.com/PR-Developers/server-health-monitor/internal/types"
)

type healthRepository interface {
	Find(query interface{}) ([]types.Health, error)
	// FindOne(where ...interface{}) (types.Health, error)
	Insert(*types.Health) (string, error)
	// Update(value interface{}) ([]types.Health, error)
	// Delete(value interface{}) (types.Health, error)
}

type baseRepository struct {
	db *database.MongoDB
}
