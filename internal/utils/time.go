package utils

import (
	"strconv"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
)

// GetMinimumLastHealthPacketTime returns a timestamp based on the given time and delay (in seconds). If delay is 0, it will use GetVariable to get the delay.
func GetMinimumLastHealthPacketTime(now time.Time, delay int) int64 {
	if delay == 0 {
		var err error
		delay, err = strconv.Atoi(GetVariable(consts.MINUTES_SINCE_HEALTH_SHOW_OFFLINE))
		if err != nil {
			delay = 2
		}
	}
	return now.UTC().Local().Add(-time.Minute * time.Duration(delay)).UnixNano()
}

func GetMinutesToIncludeHealthData() int {
	minutesToIncludeHealthData, err := strconv.Atoi(GetVariable(consts.MINUTES_TO_INCLUDE_HEALTH))
	if err != nil {
		minutesToIncludeHealthData = 5
	}
	return minutesToIncludeHealthData
}
