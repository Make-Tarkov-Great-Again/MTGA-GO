package main

import (
	"MT-GO/services"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

var address, sessionID string
var backendURL = "https://%s"
var websocketURL = "wss://%s/socket/%s"

func setHTTPSVariables() {
	sessionID = os.Getenv("SESSIONID")
	address = ip + port
	backendURL = fmt.Sprintf(backendURL, address)
	websocketURL = fmt.Sprintf(websocketURL, address, sessionID)
}

func logRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL
		fmt.Printf("Incoming %s Request URL: %s", r.Method, r.URL)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func setHTTPSServer(ip string, port string, hostname string) {
	setHTTPSVariables()

	httpsMux := http.NewServeMux()
	httpsHandler := logRequestHandler(httpsMux)

	setRoutes(httpsMux)

	cert := services.GetCertificate(ip, hostname)
	certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting HTTPS server on " + address)
	go func() {

		httpsServer := &http.Server{
			Addr:      address,
			TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
			Handler:   httpsHandler,
		}

		err := httpsServer.ListenAndServeTLS("", "")
		if err != nil {
			panic(err)
		}
	}()
}

func setRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/WebSocketAddress", webSocketHandler)
	mux.HandleFunc("/getBundleList", getBundleList)
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	services.ZlibReply(w, websocketURL)
}

func getBundleList(w http.ResponseWriter, r *http.Request) {
	services.ZlibJSONReply(w, []string{})
}
