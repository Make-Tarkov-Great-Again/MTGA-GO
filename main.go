package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		log.Print("welcome to the pool")
		fmt.Fprintf(response, "Hello my friend, welcome to the pool")
	})

	log.Print("pool is open")
	var err error = http.ListenAndServe(":80", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Print("pool is closed")
	} else if err != nil {
		log.Fatal(err)
	}
}
