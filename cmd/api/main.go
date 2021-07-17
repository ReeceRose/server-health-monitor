package main

import (
	"fmt"

	"github.com/PR-Developers/server-health-monitor/internal/api/server"
	"github.com/PR-Developers/server-health-monitor/internal/database"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	log := logger.Instance()
	log.Info("Server Health Monitor API")

	server := server.New()

	database, err := database.Instance()
	defer database.Disconnect()

	// NOTE: This is temporary and just an example
	repository := repository.NewHealthRepository()

	data, err := repository.Find(bson.M{})

	if err != nil {
		panic(err)
	}
	for _, d := range data {
		fmt.Println(d.Host.Hostname)
	}

	server.Start()
}
