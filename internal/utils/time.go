package utils

import (
	"strconv"
	"time"

	"github.com/PR-Developers/server-health-monitor/internal/consts"
)

func GetMinimumLastHealthPacketTime(now time.Time) int64 {
	delay, err := strconv.Atoi(GetVariable(consts.HEALTH_DELAY))
	if err != nil {
		delay = 2
	}
	return now.UTC().Local().Add(-time.Minute * time.Duration(delay)).UnixNano()
}
