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
	osWrapper.WriteFile(utils.GetVariable(consts.LOG_FILE), nil, 066)
}

func TestLogger_Instance_InitializesLogger(t *testing.T) {
	resetLogger()

	assert.Nil(t, logger)
	log := Instance()
	assert.NotNil(t, log)
	assert.NotNil(t, logger)
	log = Instance()
	assert.NotNil(t, log)
}

func TestLogger_Instance_ReturnsNilWhenFileFails(t *testing.T) {
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

func TestLogger_Info_WritesInfoTag(t *testing.T) {
	resetLogger()

	log := Instance()

	log.Info("generic message")
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))
	assert.Contains(t, string(data), "INFO:")
	osWrapper.Remove(consts.LOG_FILE)
}

func TestLogger_Infof_WritesExpectedValue(t *testing.T) {
	resetLogger()

	log := Instance()

	log.Infof("generic message %d", 1)
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))
	assert.Contains(t, string(data), "INFO:")
	assert.Contains(t, string(data), "generic message 1")
	osWrapper.Remove(consts.LOG_FILE)
}

func TestLogger_Warning_WritesWarningTag(t *testing.T) {
	resetLogger()

	log := Instance()

	log.Warning("generic message")
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))
	assert.Contains(t, string(data), "WARNING:")

	osWrapper.Remove(consts.LOG_FILE)
}

func TestLogger_Warningf_WritesExpectedValue(t *testing.T) {
	resetLogger()

	log := Instance()

	log.Warningf("generic message %d", 2)
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))
	assert.Contains(t, string(data), "WARNING:")
	assert.Contains(t, string(data), "generic message 2")
	osWrapper.Remove(consts.LOG_FILE)
}

func TestLogger_Error_WritesErrorTag(t *testing.T) {
	resetLogger()

	log := Instance()

	log.Error("generic message")
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))

	assert.Contains(t, string(data), "ERROR:")

	osWrapper.Remove(consts.LOG_FILE)
}

func TestLogger_Errorf_WritesExpectedValue(t *testing.T) {
	resetLogger()

	log := Instance()

	log.Errorf("generic message %d", 3)
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))
	assert.Contains(t, string(data), "ERROR:")
	assert.Contains(t, string(data), "generic message 3")
	osWrapper.Remove(consts.LOG_FILE)
}

func TestLogger_Logger_ReturnsLogger(t *testing.T) {
	resetLogger()

	log := Instance()

	logger := log.Logger()
	logger.Println("generic message")
	data, _ := osWrapper.ReadFile(utils.GetVariable(consts.LOG_FILE))

	assert.Contains(t, string(data), "generic message")

	osWrapper.Remove(consts.LOG_FILE)
}
