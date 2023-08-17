package server

import (
	"MT-GO/database"
	"MT-GO/services"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func decompressAndParseJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL
		fmt.Println("Incoming [", r.Method, "] Request URL: [", r.URL.Path, "]")
		services.DecompressInZLIBRFC1950(next, w, r)
	})
}

func addContentTypeParser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if containsMethod(method) {
			var body []byte
			buffer := bytes.NewBuffer(body)
			_, err := io.Copy(buffer, r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}
			r.Body = io.NopCloser(buffer)
		}

		next.ServeHTTP(w, r)
	})
}

var methods = []string{"POST", "PUT", "PATCH", "DELETE"}

func containsMethod(method string) bool {
	for _, m := range methods {
		if m == method {
			return true
		}
	}
	return false
}

func SetHTTPSServer(ip string, port string, hostname string) {
	main := http.NewServeMux()

	setRoutes(main)

	cert := services.GetCertificate(ip, hostname)
	certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
	if err != nil {
		panic(err)
	}

	go startHTTPServer(main, certs)
}

func startHTTPServer(handler http.Handler, certs tls.Certificate) {
	address := database.GetIPandPort()
	fmt.Println("Starting HTTPS server on " + address)

	httpsServer := &http.Server{
		Addr:      address,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
		Handler:   addContentTypeParser(decompressAndParseJSONHandler(handler)),
	}

	err := httpsServer.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
