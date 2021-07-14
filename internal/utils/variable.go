package utils

import (
	"os"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
)

// GetVariable returns a value given a key.
// The order in which GetVariable tries to read and pull values from is:
// arguments (from CLI), environemnt variables, and finally default values
func GetVariable(key string, args string) string {
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
		return "http://localhost:3000/api/v1/"
	}
	return ""
}
