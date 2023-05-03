package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello my friend, welcome to the pool")
	})

	var err error = http.ListenAndServe(":80", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("pool is closed")
	} else if err != nil {
		fmt.Printf("someone pissed in the pool: %s", err)
	}
}
