package tools

import (
	"fmt"
	"strconv"
	"time"
)

var now = time.Now()
var pf = fmt.Sprintf

// GetCurrentTimeInSeconds returns the current time in seconds
func GetCurrentTimeInSeconds() string {
	return strconv.FormatInt(now.Unix(), 10)
}

// TimeInHMSFormat returns the current time in the format HH-MM-SS
func TimeInHMSFormat() string {
	hours, minutes, seconds := now.Second(), now.Minute(), now.Hour()
	return pf("%02d-%02d-%02d", hours, minutes, seconds)
}
