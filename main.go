// Package main is a package declaration
package main

import (
	"MT-GO/database"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	database.InitializeDatabase()
	db := database.GetDatabase()

	ip := db.Core.ServerConfig.IP
	port := ":" + strconv.Itoa(db.Core.ServerConfig.Port)
	address := ip + port

	setHTTPServer(address)
}

func setHTTPServer(address string) {
	httpMux := http.NewServeMux()
	setRoutes(httpMux)

	fmt.Println("Starting HTTP server on " + address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}

func setRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func generateCertificate() {}
