package server

import (
	"MT-GO/data"
	"MT-GO/pkg"
	"crypto/tls"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type muxt struct {
	address    string
	serverName string
	initRoutes func(mux *chi.Mux)
}

func Start() {
	muxers := []*muxt{
		{
			address:    data.GetMainIPandPort(),
			serverName: "Main", initRoutes: loadMainRoutes,
		},
		{
			address:    data.GetTradingIPandPort(),
			serverName: "Trading", initRoutes: loadTradingRoutes,
		},
		{
			address:    data.GetMessagingIPandPort(),
			serverName: "Messaging", initRoutes: loadMessagingRoutes,
		},
		{
			address:    data.GetRagFairIPandPort(),
			serverName: "RagFair", initRoutes: loadRagfairRoutes,
		},
		{
			address:    data.GetLobbyIPandPort(),
			serverName: "Lobby", initRoutes: loadLobbyRoutes,
		},
	}

	serverReady := make(chan struct{})
	serverConfig := data.GetServerConfig()
	pkg.SetDownloadLocal(serverConfig.DownloadImageFiles)
	pkg.SetChannelTemplate()
	pkg.SetGameConfig()

	if serverConfig.Secure {
		cert := GetCertificate(serverConfig.IP)
		certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
		if err != nil {
			log.Fatalf("Error loading X.509 key pair: %v", err)
		}
		cert.Certificate = certs

		for _, mux := range muxers {
			go startSecure(serverReady, cert, mux)
		}
	} else {
		for _, mux := range muxers {
			go startInsecure(serverReady, mux)
		}
	}

	for range muxers {
		<-serverReady
	}
	close(serverReady)
}

func startInsecure(serverReady chan<- struct{}, mux *muxt) {
	r := chi.NewRouter()

	r.Use(middleware.URLFormat)
	r.Use(logRoute)
	r.Use(handleWebSocketUpgrade)
	r.Use(decompress)
	mux.initRoutes(r)

	server := http.Server{
		Addr:              mux.address,
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
		Handler:           r,
	}

	fmt.Println("Started " + mux.serverName + " HTTP server on " + mux.address)
	serverReady <- struct{}{}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func startSecure(serverReady chan<- struct{}, certs *Certificate, mux *muxt) {
	r := chi.NewRouter()

	r.Use(middleware.URLFormat)
	r.Use(logRoute)
	r.Use(handleWebSocketUpgrade)
	r.Use(decompress)

	mux.initRoutes(r)

	httpsServer := &http.Server{
		Addr:      mux.address,
		ConnState: CW.OnStateChange,
		TLSConfig: &tls.Config{
			RootCAs:      nil,
			Certificates: []tls.Certificate{certs.Certificate},
			MinVersion:   tls.VersionTLS12,
		},
		ReadTimeout:       time.Second * 5,
		ReadHeaderTimeout: time.Second * 5,
		Handler:           r,
	}

	fmt.Println("Started " + mux.serverName + " HTTPS server on " + mux.address)
	serverReady <- struct{}{}

	err := httpsServer.ListenAndServeTLS(certs.CertFile, certs.KeyFile)
	if err != nil {
		log.Fatalln(err)
	}
}
