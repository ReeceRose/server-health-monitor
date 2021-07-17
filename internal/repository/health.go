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

type HealthRepository struct {
	*baseRepository
	collection     *mongo.Collection
	collectionName string // collection.Name() is an alternative but this is a static name so no need to query it
	log            logger.Logger
}

var (
	_ healthRepository = (*HealthRepository)(nil)
)

func NewHealthRepository() *HealthRepository {
	db, _ := database.Instance()
	return &HealthRepository{
		baseRepository: &baseRepository{
			db: db,
		},
		collection:     db.Client.Database(utils.GetVariable(consts.DB_NAME)).Collection(consts.COLLECTION_HEALTH),
		collectionName: consts.COLLECTION_HEALTH,
		log:            logger.Instance(),
	}
}

func (r *HealthRepository) Find(query interface{}) ([]types.Health, error) {
	cursor, err := r.collection.Find(r.db.Context, query)
	if err != nil {
		msg := fmt.Sprintf("failed to read data from collection: %s with query: %s (%s)", r.collectionName, query, err.Error())
		r.log.Error(msg)
		return nil, fmt.Errorf(msg)
	}

	var data []types.Health
	defer cursor.Close(r.db.Context)
	for cursor.Next(r.db.Context) {
		var record types.Health
		if err = cursor.Decode(&record); err != nil {
			r.log.Warning(fmt.Sprintf("failed to read record on %s with query: %s", r.collectionName, query))
		}
		data = append(data, record)
	}

	return data, nil
}

func (r *HealthRepository) Insert(data *types.Health) (string, error) {
	res, err := r.collection.InsertOne(r.db.Context, data)
	if err != nil {
		msg := fmt.Sprintf("failed to insert data into collection: %s with data: %x", r.collectionName, data)
		r.log.Error(msg)
		return "", fmt.Errorf(msg)
	}
	return fmt.Sprintf("%x", res.InsertedID), nil
}
