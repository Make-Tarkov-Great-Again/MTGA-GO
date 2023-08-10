package server

import (
	"MT-GO/services"
	"crypto/tls"
	"fmt"
	"net/http"
)

var address, sessionID string
var backendURL = "https://%s"
var websocketURL = "wss://%s/socket/%s"

func setHTTPSVariables(ip string, port string) {
	address = ip + port
	backendURL = fmt.Sprintf(backendURL, address)
}

func GetWebsocketURL() string {
	//SESSIONID := GetSessionID()
	return fmt.Sprintf(websocketURL, address, sessionID)
}

func GetSessionID() string {
	return ""
}

func logRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL
		fmt.Println("Incoming [", r.Method, "] Request URL: [", r.URL.Path, "]")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func SetHTTPSServer(ip string, port string, hostname string) {
	setHTTPSVariables(ip, port)

	main := http.NewServeMux()

	setRoutes(main)

	cert := services.GetCertificate(ip, hostname)
	certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting HTTPS server on " + address)
	go startHTTPServer(address, main, certs)
}

func startHTTPServer(address string, handler http.Handler, certs tls.Certificate) {
	httpsServer := &http.Server{
		Addr:      address,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
		Handler:   logRequestHandler(handler),
	}

	err := httpsServer.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
