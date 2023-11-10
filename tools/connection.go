// Package tools contains various tools that are used throughout the srv
package tools

import (
	"net"
	"time"
)

// CheckInternet returns true if the internet is connected
func CheckInternet() bool {
	// Create a channel of boolean values to receive results.
	result := make(chan bool)

	// Start a goroutine that performs an internet connectivity check.
	go func() {
		// The internet connectivity check consists of a race between a DNS lookup to google.com
		// and a timeout of 5 seconds. We use a select statement to perform the race.
		select {
		case <-time.After(5 * time.Second):
			result <- false // Timeout occurred.
		case <-func() chan bool {
			lookupResult := make(chan bool)
			go func() {
				_, err := net.LookupHost("google.com")
				if err != nil {
					lookupResult <- false
				} else {
					lookupResult <- true
				}
			}()
			return lookupResult
		}():
			result <- true // Internet connectivity exists.
		}
	}()

	// Wait for the internet connectivity check to complete, then return the result.
	return <-result
}
