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

var buffer = &bytes.Buffer{}

func logAndDecompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL
		fmt.Println("Incoming [" + r.Method + "] Request URL: [" + r.URL.Path + "] on [" + r.Host + "]")

		if buffer != nil {
			buffer.Reset()
		}

		if r.Header.Get("Content-Type") != "application/json" {
			next.ServeHTTP(w, r)
			return
		}

		buffer = services.ZlibInflate(r)
		if buffer == nil || buffer.Len() == 0 {
			next.ServeHTTP(w, r)
			return
		}

		var parsedData interface{}
		err := json.Unmarshal(buffer.Bytes(), &parsedData)
		if err != nil {
			panic(err)
		}

		/// Store the parsed data in the request's context
		ctx := context.WithValue(r.Context(), services.ParsedBodyKey, parsedData)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func SetHTTPSServer() {
	main := http.NewServeMux()
	setRoutes(main)

	trading := http.NewServeMux()
	setTradingRoutes(trading)

	messaging := http.NewServeMux()
	setMessagingRoutes(messaging)

	ragfair := http.NewServeMux()
	setRagfairRoutes(ragfair)

	cert := services.GetCertificate(database.GetServerConfig().IP, database.GetServerConfig().Hostname)
	certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(4)
	go startMainHTTPSServer(&wg, main, certs)
	go startMessagingHTTPSServer(&wg, messaging, certs)
	go startRagFairHTTPSServer(&wg, ragfair, certs)
	go startTradingHTTPSServer(&wg, trading, certs)
	wg.Wait()
}

func startMainHTTPSServer(wg *sync.WaitGroup, handler http.Handler, certs tls.Certificate) {
	address := database.GetMainIPandPort()

	httpsServer := &http.Server{
		Addr:      address,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
		Handler:   logAndDecompress(handler),
	}
	fmt.Println("Started Main HTTPS server on " + address)
	wg.Done()
	err := httpsServer.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}

func startTradingHTTPSServer(wg *sync.WaitGroup, handler http.Handler, certs tls.Certificate) {
	address := database.GetTradingIPandPort()

	httpsServer := &http.Server{
		Addr:      address,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
		Handler:   logAndDecompress(handler),
	}
	fmt.Println("Started Trading HTTPS server on " + address)
	wg.Done()
	err := httpsServer.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}

func startMessagingHTTPSServer(wg *sync.WaitGroup, handler http.Handler, certs tls.Certificate) {
	address := database.GetMessagingIPandPort()

	httpsServer := &http.Server{
		Addr:      address,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
		Handler:   logAndDecompress(handler),
	}
	fmt.Println("Started Messaging HTTPS server on " + address)
	wg.Done()
	err := httpsServer.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}

func startRagFairHTTPSServer(wg *sync.WaitGroup, handler http.Handler, certs tls.Certificate) {
	address := database.GetRagFairIPandPort()

	httpsServer := &http.Server{
		Addr:      address,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
		Handler:   logAndDecompress(handler),
	}
	fmt.Println("Started RagFair HTTPS server on " + address)
	wg.Done()
	err := httpsServer.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
