package tools

import (
	"time"
)

// GetCurrentTimeInSeconds returns the current time in seconds
func GetCurrentTimeInSeconds() int64 {
	return time.Now().Unix()
}

// TimeInHMSFormat returns the current time in the format HH-MM-SS
func TimeInHMSFormat() string {
	return time.Now().Format("15-04-05")
}
