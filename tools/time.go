package tools

import (
	"strconv"
	"time"
)

var now = time.Now()

// GetCurrentTimeInSeconds returns the current time in seconds
func GetCurrentTimeInSeconds() string {
	return strconv.FormatInt(now.Unix(), 10)
}

// TimeInHMSFormat returns the current time in the format HH-MM-SS
func TimeInHMSFormat() string {
	return now.Format("15-04-05")
}
