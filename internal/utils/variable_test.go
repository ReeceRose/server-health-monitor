package utils

import (
	"os"
	"testing"

	"github.com/PR-Developers/server-health-monitor/internal/consts"

	"github.com/stretchr/testify/assert"
)

// No OS wrapper used here as mocking requires too mich boilerplate for this simple test
// and the changes of os.Setenv/os.Getenv/os.Clearenv not working are minimal

func TestGetVariableReturnsValueFromEnvironmentVariable(t *testing.T) {
	os.Clearenv()
	os.Setenv(consts.API_PORT, "4000")
	os.Setenv(consts.API_URL, "https://api.pr-developers.com/api/v1")
	os.Setenv(consts.CERT_DIR, "/certs-dir")
	os.Setenv(consts.CLIENT_CERT, "ssl.crt")
	os.Setenv(consts.API_CERT, "ssl.crt")
	os.Setenv(consts.API_KEY, "ssl.key")
	os.Setenv(consts.DB_URI, "mongodb://localhost:27017/shm")
	os.Setenv(consts.DB_NAME, "shm")
	os.Setenv(consts.DB_USER, "user")
	os.Setenv(consts.DB_PASS, "pass")
	os.Setenv(consts.LOG_FILE, "shm.log")

	assert.Equal(t, "4000", GetVariable(consts.API_PORT))
	assert.Equal(t, "https://api.pr-developers.com/api/v1", GetVariable(consts.API_URL))
	assert.Equal(t, "/certs-dir", GetVariable(consts.CERT_DIR))
	assert.Equal(t, "ssl.crt", GetVariable(consts.CLIENT_CERT))
	assert.Equal(t, "ssl.crt", GetVariable(consts.API_CERT))
	assert.Equal(t, "ssl.key", GetVariable(consts.API_KEY))
	assert.Equal(t, "mongodb://localhost:27017/shm", GetVariable(consts.DB_URI))
	assert.Equal(t, "shm", GetVariable(consts.DB_NAME))
	assert.Equal(t, "user", GetVariable(consts.DB_USER))
	assert.Equal(t, "pass", GetVariable(consts.DB_PASS))
	assert.Equal(t, "shm.log", GetVariable(consts.LOG_FILE))

	os.Clearenv()
}

func TestGetVariableDefaultsValue(t *testing.T) {
	os.Clearenv()
	assert.Equal(t, "3000", GetVariable(consts.API_PORT))
	assert.Equal(t, "https://localhost:3000/api/v1/", GetVariable(consts.API_URL))
	assert.Equal(t, "certs", GetVariable(consts.CERT_DIR))
	assert.Equal(t, "localhost.crt", GetVariable(consts.CLIENT_CERT))
	assert.Equal(t, "localhost.crt", GetVariable(consts.API_CERT))
	assert.Equal(t, "localhost.key", GetVariable(consts.API_KEY))
	assert.Equal(t, "mongodb://localhost:27017/server-health-monitor", GetVariable(consts.DB_URI))
	assert.Equal(t, "server-health-monitor", GetVariable(consts.DB_NAME))
	assert.Equal(t, "admin", GetVariable(consts.DB_USER))
	assert.Equal(t, "admin", GetVariable(consts.DB_PASS))
	assert.Equal(t, "server-health-monitor.log", GetVariable(consts.LOG_FILE))
	assert.Equal(t, "", GetVariable("test-value"))
}

func TestGetVariableWithArgsReturnsFromArgs(t *testing.T) {

}
