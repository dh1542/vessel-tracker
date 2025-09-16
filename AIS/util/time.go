package util

import (
	"time"
)

func TimeFromInt32(timestamp int32) time.Time {
	return time.Unix(int64(timestamp), 0).UTC()
}
