package main

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/client"
	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/data-collector/host"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/types"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
)

func main() {
	log := logger.Instance()
	log.Info("Server Health Monitor - Data Collector Tool")

	// TODO: pass arguments to GetVariable
	client, err := client.NewClient(utils.GetVariable(consts.API_URL))
	if err != nil {
		panic(err)
	}

	// TODO: Read from command line
	var delay time.Duration = 30 // delay in seconds

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(host.GetInfo())
	log.Info("Sending initial host data")
	_, statusCode, _ := client.Post("host/", payload)
	log.Infof("Sent host data and got a status code of %v", statusCode)

	for {
		// Collect new data

		// Make request
		log.Info("Sending new health data")
		health := types.Health{
			Uptime: host.GetInfo().Uptime,
		}
		json.NewEncoder(payload).Encode(health)
		_, statusCode, _ := client.Post("health/", payload)
		log.Infof("Sent health data and got a status code of %v", statusCode)
		// Delay for X seconds
		time.Sleep(time.Second * delay)
	}
}
