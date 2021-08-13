package database

import (
	"context"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database is an interface which provides method signatures for interacting with a database
type Database interface {
	Disconnect() error
	Client() *mongo.Client
	Context() context.Context
}

type mongoDB struct {
	client *mongo.Client
	// context context.Context
}

var (
	database *mongoDB
	_        Database = (*mongoDB)(nil)
)

// Instance returns the active instance of the database
func Instance() (Database, error) {
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

	database = &mongoDB{
		client: client,
	}
	return database, nil
}

// Disconnect is used to disconnect the database at the end of a session
func (db *mongoDB) Disconnect() error {
	return database.client.Disconnect(context.Background())
}

// Client returns the active database client
func (db *mongoDB) Client() *mongo.Client {
	return database.client
}

// Context returns a new context which is used by the database
func (db *mongoDB) Context() context.Context {
	return context.Background()
}
