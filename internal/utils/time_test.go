package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime_GetMinimumLastHealthPacketTime_ReturnsExpectedTime(t *testing.T) {
	var setTime int64 = 1628042730               // GMT: Wednesday, August 4, 2021 2:05:30 AM
	var previousTime int64 = 1628042610000000000 // GMT: Wednesday, August 4, 2021 2:03:30 AM

	assert.Equal(t, int64(previousTime), GetMinimumLastHealthPacketTime(time.Unix(setTime, 0)))
}
