package main

import (
	certificate "MT-GO/server"
	"fmt"
	"net/http"
	"os"
)

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

func setHTTPSServer(ip string, port string, hostname string) {
	httpsMux := http.NewServeMux()
	setRoutes(httpsMux)

	cert := certificate.GetCertificate(ip, hostname)
	if cert == nil {
		fmt.Print("fucking faggot cert")
	}
	address := ip + port

	fmt.Println("Starting HTTPS server on " + address)
	go func() {
		err := http.ListenAndServeTLS(address, cert.CertFile, cert.KeyFile, httpsMux)
		if err != nil {
			panic(err)
		}
	}()
}

func setRoutes(mux *http.ServeMux) {
	//mux.HandleFunc("/client/WebSocketAddress", webSocketHandler)
}

// SESSIONID variable is used to identify the session used
var SESSIONID = os.Getenv("SESSIONID")

/* func GetServerAddress() string {

}

func webSocketHandler(w http.ResponseWriter, _ *http.Request) {

} */
