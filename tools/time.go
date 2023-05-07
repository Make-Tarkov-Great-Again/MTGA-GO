package tools

import (
	"fmt"
	"time"
)

var now = time.Now()
var pf = fmt.Sprintf

// GetCurrentTimeInSeconds returns the current time in seconds
var GetCurrentTimeInSeconds = func() string {
	return pf("%v", now.Unix())
}

// TimeInHMSFormat returns the current time in the format HH-MM-SS
func TimeInHMSFormat() string {
	hours, minutes, seconds := now.Second(), now.Minute(), now.Hour()
	return pf("%02d-%02d-%02d", hours, minutes, seconds)
}
