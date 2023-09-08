package server

import (
	"MT-GO/database"
	"MT-GO/services"
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func logAndDecompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL
		fmt.Println("Incoming [" + r.Method + "] Request URL: [" + r.URL.Path + "] on [" + strings.TrimPrefix(r.Host, "127.0.0.1") + "]")

		if r.Header.Get("Connection") == "Upgrade" && r.Header.Get("Upgrade") == "websocket" {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer conn.Close()

			for {
				messageType, p, err := conn.ReadMessage()
				if err != nil {
					return
				}
				if err := conn.WriteMessage(messageType, p); err != nil {
					return
				}
			}
		} else {
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
		}
	})
}

func startHTTPSServer(serverReady chan<- struct{}, certs *services.Certificate, mux *muxt) {
	mux.initRoutes(mux.mux)

	httpsServer := &http.Server{
		Addr: mux.address,
		TLSConfig: &tls.Config{
			RootCAs:      nil,
			Certificates: []tls.Certificate{certs.Certificate},
		},
		Handler: logAndDecompress(mux.mux),
	}

	fmt.Println("Started " + mux.serverName + " HTTPS server on " + mux.address)
	serverReady <- struct{}{}

	err := httpsServer.ListenAndServeTLS(certs.CertFile, certs.KeyFile)
	if err != nil {
		panic(err)
	}
}

type muxt struct {
	mux        *http.ServeMux
	address    string
	serverName string
	initRoutes func(mux *http.ServeMux)
}

func SetHTTPSServer() {
	srv := database.GetServerConfig()
	/* 	hostnames := []string{
		srv.Ports.Main,
		srv.Ports.Messaging,
		srv.Ports.Trading,
		srv.Ports.Flea,
		srv.Ports.Lobby,
	} */

	cert := services.GetCertificate(srv.IP)
	certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
	if err != nil {
		panic(err)
	}
	cert.Certificate = certs

	fmt.Println()

	muxes := []*muxt{
		{
			mux: http.NewServeMux(), address: database.GetMainIPandPort(),
			serverName: "Main", initRoutes: setMainRoutes, // Embed the route initialization function
		},
		{
			mux: http.NewServeMux(), address: database.GetTradingIPandPort(),
			serverName: "Trading", initRoutes: setTradingRoutes, // Embed the route initialization function
		},
		{
			mux: http.NewServeMux(), address: database.GetMessagingIPandPort(),
			serverName: "Messaging", initRoutes: setMessagingRoutes, // Embed the route initialization function
		},
		{
			mux: http.NewServeMux(), address: database.GetRagFairIPandPort(),
			serverName: "RagFair", initRoutes: setRagfairRoutes, // Embed the route initialization function
		},
		{
			mux: http.NewServeMux(), address: database.GetLobbyIPandPort(),
			serverName: "Lobby", initRoutes: setLobbyRoutes, // Embed the route initialization function
		},
	}

	serverReady := make(chan struct{})

	for _, muxData := range muxes {
		go startHTTPSServer(serverReady, cert, muxData)
	}

	for range muxes {
		<-serverReady
	}

	close(serverReady)
	fmt.Println()
}
