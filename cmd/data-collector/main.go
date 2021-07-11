package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/client"
	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/data-collector/host"
	"github.com/PR-Developers/server-health-monitor/internal/logger"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
)

type Host struct {
	Hostname string
}

func main() {
	log := logger.Instance()
	log.Info("Server Health Monitor - Data Collector Tool")

	host := host.GetInfo()
	// TODO: Read from command line
	var delay time.Duration = 30 // delay in seconds
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(host)
	// TODO: pass arguments to GetVariable
	client := client.NewClient(utils.GetVariable(consts.API_URL, ""))

	for {
		// Collect new data

		// Make request
		log.Info("Sending new health data")
		_, statusCode, _ := client.Post("health", payload)
		log.Info(fmt.Sprintf("Sent health data and got a status code of %v", statusCode))
		// Delay for X seconds
		time.Sleep(time.Second * delay)
	}
}
