package main

import (
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/server"
)

func main() {
	log := logger.Instance()
	log.Info("Server Health Monitor API")

	server := server.New()

	server.Start()
}
