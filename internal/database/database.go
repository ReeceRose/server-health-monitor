package database

import (
	"context"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Disconnect() error
}

type MongoDB struct {
	Client  *mongo.Client
	Context context.Context
}

var (
	database *MongoDB
	_        Database = (*MongoDB)(nil)
)

// Instance returns the active instance of the database
func Instance() (*MongoDB, error) {
	if database != nil {
		return database, nil
	}

	clientOptions := options.Client().ApplyURI(utils.GetVariable(consts.DB_URI)).SetAuth(options.Credential{
		Username: utils.GetVariable(consts.DB_USER),
		Password: utils.GetVariable(consts.DB_PASS),
	})

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	database = &MongoDB{
		Client:  client,
		Context: context.Background(),
	}
	return database, nil
}

// Disconnect is used to disconnect the database at the end of a session
func (db *MongoDB) Disconnect() error {
	return database.Client.Disconnect(db.Context)
}
