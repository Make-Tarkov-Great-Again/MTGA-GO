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
	"strings"
	"sync"
)

func logAndDecompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL
		fmt.Println("Incoming [" + r.Method + "] Request URL: [" + r.URL.Path + "] on [" + strings.TrimPrefix(r.Host, "127.0.0.1") + "]")

		buffer := &bytes.Buffer{}

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

func startHTTPSServer(wg *sync.WaitGroup, certs tls.Certificate, mux *http.ServeMux, address string, serverName string) {
	httpsServer := &http.Server{
		Addr:      address,
		TLSConfig: &tls.Config{Certificates: []tls.Certificate{certs}},
		Handler:   logAndDecompress(mux),
	}
	fmt.Println("Started " + serverName + " HTTPS server on " + address)
	wg.Done()
	err := httpsServer.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}

func SetHTTPSServer() {
	cert := services.GetCertificate(database.GetServerConfig().IP, database.GetServerConfig().Hostname)
	certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	wg := sync.WaitGroup{}

	muxes := []struct {
		mux        *http.ServeMux
		address    string
		serverName string
	}{
		{http.NewServeMux(), database.GetMainIPandPort(), "Main"},
		{http.NewServeMux(), database.GetTradingIPandPort(), "Trading"},
		{http.NewServeMux(), database.GetMessagingIPandPort(), "Messaging"},
		{http.NewServeMux(), database.GetRagFairIPandPort(), "RagFair"},
	}

	for _, muxData := range muxes {
		wg.Add(1)

		switch muxData.serverName {
		case "Main":
			setMainRoutes(muxData.mux)
		case "Trading":
			setTradingRoutes(muxData.mux)
		case "Messaging":
			setMessagingRoutes(muxData.mux)
		case "RagFair":
			setRagfairRoutes(muxData.mux)
		}

		go startHTTPSServer(&wg, certs, muxData.mux, muxData.address, muxData.serverName)
	}

	wg.Wait()
	fmt.Println()
}
