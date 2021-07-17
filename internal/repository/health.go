package repository

import (
	"fmt"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/database"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type HealthRepository interface {
	Find(query interface{}) ([]types.Health, error)
	// FindOne(where ...interface{}) (types.Health, error)
	// Insert(value interface{}) ([]types.Health, error)
	// Update(value interface{}) ([]types.Health, error)
	// Delete(value interface{}) (types.Health, error)
}

type healthRepository struct {
	*baseRepository
	collection     *mongo.Collection
	collectionName string // collection.Name() is an alternative but this is a static name so no need to query it
}

var (
	_ HealthRepository = (*healthRepository)(nil)
)

func NewHealthRepository() *healthRepository {
	db, _ := database.Instance()
	return &healthRepository{
		baseRepository: &baseRepository{
			db: db,
		},
		collection:     db.Client.Database(utils.GetVariable(consts.DB_NAME)).Collection(consts.COLLECTION_HEALTH),
		collectionName: consts.COLLECTION_HEALTH,
	}
}

func (r *healthRepository) Find(query interface{}) ([]types.Health, error) {
	log := logger.Instance()
	cursor, err := r.collection.Find(r.db.Context, query)
	if err != nil {
		msg := fmt.Sprintf("failed to read data from collection: %s with query: %s (%s)", r.collectionName, query, err.Error())
		log.Error(msg)
		return nil, fmt.Errorf(msg)
	}

	var data []types.Health
	defer cursor.Close(r.db.Context)
	for cursor.Next(r.db.Context) {
		var record types.Health
		if err = cursor.Decode(&record); err != nil {
			log.Warning(fmt.Sprintf("failed to read record on %s with query: %s", r.collectionName, query))
		}
		data = append(data, record)
	}

	return data, nil
}
