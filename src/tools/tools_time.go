package tools

import (
	"time"
)

func GetZeroClock(t time.Time) time.Time {
	zeroClock := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return zeroClock
}
