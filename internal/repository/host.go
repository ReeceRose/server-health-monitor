package repository

import (
	"fmt"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/database"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type hostRepository struct {
	*baseRepository
}

var (
	_ IHostRepository = (*hostRepository)(nil)
)

// NewHostRepository returns an instanced host repository
func NewHostRepository() IHostRepository {
	db, _ := database.Instance()

	return &hostRepository{
		baseRepository: &baseRepository{
			db:             db,
			collection:     db.Client().Database(utils.GetVariable(consts.DB_NAME)).Collection(consts.COLLECTION_HOST),
			collectionName: consts.COLLECTION_HOST,
			log:            logger.Instance(),
		},
	}
}

// Find all host data given a certain query
func (r *hostRepository) Find(query interface{}) ([]types.Host, error) {
	return r.FindWithFilter(query, nil)
}

// FindWithFilter returns all host data given a certain query and options
func (r *hostRepository) FindWithFilter(query interface{}, options *options.FindOptions) ([]types.Host, error) {
	cursor, err := r.collection.Find(r.db.Context(), query, options)
	if err != nil {
		msg := fmt.Sprintf("failed to read data from collection: %s with query: %s (%s)", r.collectionName, query, err.Error())
		r.log.Error(msg)
		return nil, fmt.Errorf(msg)
	}

	var data []types.Host
	defer cursor.Close(r.db.Context())
	for cursor.Next(r.db.Context()) {
		var record types.Host
		if err = cursor.Decode(&record); err != nil {
			r.log.Warningf("failed to read record on %s with query: %s", r.collectionName, query)
		}
		data = append(data, record)
	}

	return data, nil
}

// Insert a single host record into the database
func (r *hostRepository) Insert(data *types.Host) (string, error) {
	res, err := r.collection.InsertOne(r.db.Context(), data)
	if err != nil {
		msg := fmt.Sprintf("failed to insert data into collection: %s", r.collectionName)
		r.log.Error(msg)
		return "", fmt.Errorf(msg)
	}
	return fmt.Sprintf("%x", res.InsertedID), nil
}

// Replace an existing host record in the database
func (r *hostRepository) UpdateByID(data *types.Host) error {
	_, err := r.collection.UpdateByID(r.db.Context(), data.ID,
		bson.M{
			"$set": data,
		},
	)
	if err != nil {
		r.log.Error(err.Error())
		msg := "failed to update host data"
		r.log.Error(msg)
		return fmt.Errorf(msg)
	}
	return nil
}
