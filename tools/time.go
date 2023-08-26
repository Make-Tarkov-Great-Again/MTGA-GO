package tools

import (
	"time"
)

var now = time.Now()

// GetCurrentTimeInSeconds returns the current time in seconds
func GetCurrentTimeInSeconds() int64 {
	return now.Unix()
}

// TimeInHMSFormat returns the current time in the format HH-MM-SS
func TimeInHMSFormat() string {
	return now.Format("15-04-05")
}
