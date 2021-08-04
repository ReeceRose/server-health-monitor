package repository

import (
	"github.com/PR-Developers/server-health-monitor/internal/database"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"go.mongodb.org/mongo-driver/mongo"
)

// IHealthRepository is an interface which provides method signatures for a health repository
type IHealthRepository interface {
	Find(query interface{}) ([]types.Health, error)
	// FindOne(where ...interface{}) (types.Health, error)
	Insert(data *types.Health) (string, error)
	// Update(value interface{}) ([]types.Health, error)
	// Delete(value interface{}) (types.Health, error)
}

// IHostRepository is an interface which provides method signatures for a host repository
type IHostRepository interface {
	Find(query interface{}) ([]types.Host, error)
	Insert(data *types.Host) (string, error)
	UpdateByID(data *types.Host) error
}

type baseRepository struct {
	db             database.Database
	collection     *mongo.Collection
	collectionName string // collection.Name() is an alternative but this is a static name so no need to query it
	log            logger.Logger
}
