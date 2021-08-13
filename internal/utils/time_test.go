package utils

import (
	"os"
	"testing"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
	"github.com/stretchr/testify/assert"
)

const (
	setTime int64 = 1628042730 // GMT: Wednesday, August 4, 2021 2:05:30 AM
)

func TestTime_GetMinimumLastHealthPacketTime_ReturnsExpectedTimeWhenZeroIsPassed(t *testing.T) {
	assert.Equal(t, int64(1628042430000000000), GetMinimumLastHealthPacketTime(time.Unix(setTime, 0), 0))
}

func TestTime_GetMinimumLastHealthPacketTime_ReturnsExpectedTimeWhenDelayIsPassed(t *testing.T) {
	assert.Equal(t, int64(1628042610000000000), GetMinimumLastHealthPacketTime(time.Unix(setTime, 0), 2))
}

func TestTime_GetMinimumLastHealthPacketTime_ReturnsExpectedTimeWhenFailedToConvertVariable(t *testing.T) {
	os.Setenv(consts.HEALTH_DELAY, "ABC")
	assert.Equal(t, int64(1628042610000000000), GetMinimumLastHealthPacketTime(time.Unix(setTime, 0), 0))
}
