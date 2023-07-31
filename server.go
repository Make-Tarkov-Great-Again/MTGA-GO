package main

import (
	"MT-GO/tools"
	"fmt"
	"net/http"
)

/* func setServers(address string) {

	//setHTTPSServer(address)
	//setHTTPServer(address)
} */

func setHTTPServer(address string) {
	httpMux := http.NewServeMux()
	setRoutes(httpMux)

	fmt.Println("Starting HTTP server on " + address)
	fmt.Println()
	go func() {
		err := http.ListenAndServe(address, nil)
		if err != nil {
			panic(err)
		}
	}()
}

func setHTTPSServer(address string) {
	httpsMux := http.NewServeMux()
	setRoutes(httpsMux)

	cert := "user/cert/cert.pem"
	ok := tools.FileExist(cert)
	if !ok {
		fmt.Println("Certificate not found")
		return
	}

	key := "user/cert/key.pem"
	ok = tools.FileExist(key)
	if !ok {
		fmt.Println("Key not found")
		return
	}

	fmt.Println("Starting HTTPS server on " + address)
	err := http.ListenAndServeTLS(address, cert, key, httpsMux)
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
