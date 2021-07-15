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

func getDefaultForKey(key string) string {
	switch key {
	case consts.API_PORT:
		return "3000"
	case consts.API_URL:
		return "https://localhost:3000/api/v1/"
	case consts.CERT_DIR:
		return "/certs"
	case consts.API_CERT:
		return "localhost.crt"
	case consts.API_KEY:
		return "localhost.key"
	case consts.DB_URI:
		return "mongodb://localhost:27017"
	case consts.DB_USER:
		return ""
	case consts.DB_PASS:
		return ""
	}
	return ""
}
