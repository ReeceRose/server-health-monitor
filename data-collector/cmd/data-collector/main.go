package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/PR-Developers/server-health-monitor/data-collector/internal/client"
	"github.com/PR-Developers/server-health-monitor/data-collector/internal/host"
)

type Host struct {
	Hostname string
}

func main() {
	fmt.Println("Server Health Monitor - Data Collector Tool")

	host := host.GetInformation()
	// TODO: Read from command line
	var delay time.Duration = 30 // delay in seconds
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(host)
	client := client.NewClient()

	for {
		// Collect new data

		// Make request
		// TODO: better logging
		fmt.Println("Sending new health data")
		_, statusCode, _ := client.Post("health", payload)
		fmt.Printf("Sent health data and got a status code of %v\n", statusCode)
		// Delay for X seconds
		time.Sleep(time.Second * delay)
	}
}
