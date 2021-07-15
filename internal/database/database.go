package database

import (
	"context"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
}

type MongoDB struct {
	client     *mongo.Client
	context    context.Context
	cancelFunc context.CancelFunc
}

var (
	database *MongoDB
)

func Initialize() (*MongoDB, error) {
	if database != nil {
		return database, nil
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*15)
	clientOptions := options.Client().ApplyURI(utils.GetVariable(consts.DB_URI)).SetAuth(options.Credential{
		Username: utils.GetVariable(consts.DB_USER),
		Password: utils.GetVariable(consts.DB_PASS),
	})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		cancelFunc()
		return nil, err
	}

	database = &MongoDB{
		client:     client,
		context:    ctx,
		cancelFunc: cancelFunc,
	}
	return database, nil
}

func (db *MongoDB) Disconnect() error {
	db.cancelFunc()
	return database.client.Disconnect(db.context)
}
