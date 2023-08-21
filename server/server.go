package server

import (
	"MT-GO/database"
	"MT-GO/services"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var buffer *bytes.Buffer

func logAndDecompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL
		fmt.Println("Incoming [", r.Method, "] Request URL: [", r.URL.Path, "]")

		if r.Header.Get("Content-Type") != "application/json" {
			next.ServeHTTP(w, r)
			return
		}

		buffer = services.ZlibInflate(r)
		if buffer == nil || buffer.Len() == 0 {
			next.ServeHTTP(w, r)
			return
		}

		var parsedData map[string]interface{}
		err := json.Unmarshal(buffer.Bytes(), &parsedData)
		if err != nil {
			panic(err)
		}
		buffer.Reset()

		/// Store the parsed data in the request's context
		ctx := context.WithValue(r.Context(), services.ParsedBodyKey, parsedData)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func SetHTTPSServer() {
	main := http.NewServeMux()

	setRoutes(main)

	cert := services.GetCertificate(database.GetServerConfig().IP, database.GetServerConfig().Hostname)
	certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go startHTTPServer(&wg, main, certs)
	wg.Wait()
}

func startHTTPServer(wg *sync.WaitGroup, handler http.Handler, certs tls.Certificate) {
	address := database.GetIPandPort()

	httpsServer := &http.Server{
		Addr:      address,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
		Handler:   logAndDecompress(handler),
	}
	fmt.Println("Started HTTPS server on " + address)
	wg.Done()
	err := httpsServer.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
