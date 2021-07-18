package logger

import (
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/PR-Developers/server-health-monitor/internal/logger/mocks"
	"github.com/PR-Developers/server-health-monitor/internal/utils"
	"github.com/PR-Developers/server-health-monitor/internal/wrapper"

	"github.com/stretchr/testify/assert"
)

//go:generate mockery --dir=../ -r --name OperatingSystem

func resetLogger() {
	logger = nil
	osWrapper = &wrapper.DefaultOS{}
}

func TestInstanceInitializesLogger(t *testing.T) {
	resetLogger()

	assert.Nil(t, logger)
	log := Instance()
	assert.NotNil(t, log)
	assert.NotNil(t, logger)
	log = Instance()
	assert.NotNil(t, log)
}

func TestLoggerFailsToOpenFile(t *testing.T) {
	logger = nil
	wrapper := new(mocks.OperatingSystem)

	wrapper.On("OpenFile",
		utils.GetVariable(consts.LOG_FILE), os.O_APPEND|os.O_CREATE|os.O_WRONLY, fs.FileMode(0666),
	).Return(nil, fmt.Errorf("failed to read file"))

	osWrapper = wrapper

	log := Instance()
	assert.Nil(t, log)
	wrapper.AssertExpectations(t)
}

func TestLoggerWritesInfoTag(t *testing.T) {
	resetLogger()

	log := Instance()

	log.Info("generic message")
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))
	assert.Contains(t, string(data), "INFO:")
	osWrapper.Remove(consts.LOG_FILE)
}

func TestLoggerWritesWarningTag(t *testing.T) {
	resetLogger()

	log := Instance()

	log.Warning("generic message")
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))
	assert.Contains(t, string(data), "WARNING:")

	osWrapper.Remove(consts.LOG_FILE)
}

func TestLoggerWritesErrorTag(t *testing.T) {
	resetLogger()

	log := Instance()

	log.Error("generic message")
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))

	assert.Contains(t, string(data), "ERROR:")

	osWrapper.Remove(consts.LOG_FILE)
}
