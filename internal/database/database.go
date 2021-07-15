package database

import (
	"context"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
}

type MongoDB struct {
	client  *mongo.Client
	context context.Context
}

func Initialize() (*MongoDB, error) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(utils.GetVariable(consts.DB_URI))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		client:  client,
		context: ctx,
	}, nil
}
