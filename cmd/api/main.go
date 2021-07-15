package main

import (
	"github.com/PR-Developers/server-health-monitor/internal/api/server"
	"github.com/PR-Developers/server-health-monitor/internal/database"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
)

func main() {
	log := logger.Instance()
	log.Info("Server Health Monitor API")

	server := server.New()

	database, err := database.Initialize()
	if err != nil {
		panic(err)
	}

	defer database.Disconnect()

	server.Start()
}
