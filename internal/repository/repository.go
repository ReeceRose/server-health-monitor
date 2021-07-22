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
	Insert(*types.Health) (string, error)
	// Update(value interface{}) ([]types.Health, error)
	// Delete(value interface{}) (types.Health, error)
}

type baseRepository struct {
	db             database.Database
	collection     *mongo.Collection
	collectionName string // collection.Name() is an alternative but this is a static name so no need to query it
	log            logger.Logger
}
