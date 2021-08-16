package utils

import (
	"os"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
)

// GetVariable returns a value given a key.
// GetVariable first tries to read from environment variables and will default
// to preset values
func GetVariable(key string) string {
	return GetVariableWithArgs(key, "")
}

// GetVariableWithArgs returns a value given a key.
// GetVariableWithArgs first tries to read from command line arguments,
// environment variables, and will default to preset values
func GetVariableWithArgs(key string, args string) string {
	// TODO: read from arguments
	if args != "" {
	}
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return getDefaultForKey(key)
}

// getDefaultForKey is a handy method to get the default values if
// not present in arguments or environment variables
func getDefaultForKey(key string) string {
	switch key {
	case consts.API_PORT:
		return "3000"
	case consts.API_URL:
		return "https://localhost:3000/api/v1/"
	case consts.WS_URL:
		return "wss://localhost:3000/ws/v1/"
	case consts.CERT_DIR:
		return "certs"
	case consts.CLIENT_CERT:
		return "localhost.crt"
	case consts.API_CERT:
		return "localhost.crt"
	case consts.API_KEY:
		return "localhost.key"
	case consts.DB_URI:
		return "mongodb://localhost:27017/" + GetVariable(consts.DB_NAME)
	case consts.DB_NAME:
		return "server-health-monitor"
	case consts.DB_USER:
		return "admin"
	case consts.DB_PASS:
		return "admin"
	case consts.LOG_FILE:
		return "server-health-monitor.log"
	case consts.MINUTES_SINCE_HEALTH_SHOW_OFFLINE:
		return "5"
	case consts.DC_HEALTH_DELAY:
		return "30"
	case consts.MINUTES_TO_INCLUDE_HEALTH:
		return "5"
	case consts.DATA_WEBSOCKET_DELAY:
		return "30"
	}
	return ""
}
