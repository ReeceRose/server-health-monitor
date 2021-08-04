package utils

import (
	"time"
)

func GetMinimumLastHealthPacketTime(now time.Time) int64 {
	return now.UTC().Local().Add(-time.Minute * 2).UnixNano()
}
