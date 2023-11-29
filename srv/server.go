package srv

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"MT-GO/data"
	"MT-GO/pkg"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func upgradeToWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(conn)

	sessionID := r.URL.Path[28:] //mongoID is 24 chars
	data.SetConnection(sessionID, conn)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			data.DeleteConnection(sessionID)
			return
		}
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

const incomingRoute string = "[%s] %s on %s\n"

func logAndDecompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request URL
		log.Printf(incomingRoute, r.Method, r.URL.Path, strings.TrimPrefix(r.Host, "127.0.0.1"))

		if websocket.IsWebSocketUpgrade(r) {
			upgradeToWebsocket(w, r)
		} else {
			if r.Header.Get("Content-Length") == "" {
				next.ServeHTTP(w, r)
				return
			}

			buffer := pkg.ZlibInflate(r)
			if buffer == nil || buffer.Len() == 0 {
				next.ServeHTTP(w, r)
				return
			}

			//TODO: Refactor to replace r.Body ((remove CTX))
			var parsedData map[string]any
			if err := json.Unmarshal(buffer.Bytes(), &parsedData); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			//ctx := context.WithValue(r.Context(), pkg.ParsedBodyKey, parsedData)
			r = r.WithContext(context.WithValue(r.Context(), pkg.ParsedBodyKey, parsedData))

			next.ServeHTTP(w, r)
		}
	})
}

var CW = &ConnectionWatcher{}

func startHTTPSServer(serverReady chan<- struct{}, certs *Certificate, mux *muxt) {
	mux.initRoutes(mux.mux)

	httpsServer := &http.Server{
		Addr:      mux.address,
		ConnState: CW.OnStateChange,
		TLSConfig: &tls.Config{
			RootCAs:      nil,
			Certificates: []tls.Certificate{certs.Certificate},
		},
		Handler: logAndDecompress(mux.mux),
	}

	go func() {
		fmt.Println("Started " + mux.serverName + " HTTPS server on " + mux.address)
		serverReady <- struct{}{}

		// Use a separate goroutine to handle graceful shutdown
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		// Wait for the interrupt signal
		<-stop

		fmt.Println("Shutting down " + mux.serverName + " server gracefully...")

		// Create a context with a timeout to allow existing requests to finish
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Shutdown the server
		if err := httpsServer.Shutdown(ctx); err != nil {
			log.Println("Error shutting down server:", err)
		}
	}()

	err := httpsServer.ListenAndServeTLS(certs.CertFile, certs.KeyFile)
	if err != nil {
		log.Fatalln(err)
	}
}

func startHTTPServer(serverReady chan<- struct{}, mux *muxt) {
	mux.initRoutes(mux.mux)

	httpServer := &http.Server{
		Addr:    mux.address,
		Handler: logAndDecompress(mux.mux),
	}

	go func() {
		fmt.Println("Started " + mux.serverName + " HTTP server on " + mux.address)
		serverReady <- struct{}{}

		// Use a separate goroutine to handle graceful shutdown
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		// Wait for the interrupt signal
		<-stop

		fmt.Println("Shutting down " + mux.serverName + " server gracefully...")

		// Create a context with a timeout to allow existing requests to finish
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Shutdown the server
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Println("Error shutting down server:", err)
		}
	}()

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

type muxt struct {
	mux        *http.ServeMux
	address    string
	serverName string
	initRoutes func(mux *http.ServeMux)
}

func SetServer() {
	muxers := []*muxt{
		{
			mux: http.NewServeMux(), address: data.GetMainIPandPort(),
			serverName: "Main", initRoutes: setMainRoutes,
		},
		{
			mux: http.NewServeMux(), address: data.GetTradingIPandPort(),
			serverName: "Trading", initRoutes: setTradingRoutes,
		},
		{
			mux: http.NewServeMux(), address: data.GetMessagingIPandPort(),
			serverName: "Messaging", initRoutes: setMessagingRoutes,
		},
		{
			mux: http.NewServeMux(), address: data.GetRagFairIPandPort(),
			serverName: "RagFair", initRoutes: setRagfairRoutes,
		},
		{
			mux: http.NewServeMux(), address: data.GetLobbyIPandPort(),
			serverName: "Lobby", initRoutes: setLobbyRoutes,
		},
	}

	serverReady := make(chan struct{})
	srv := data.GetServerConfig()

	if srv.Secure {
		cert := GetCertificate(srv.IP)
		certs, err := tls.LoadX509KeyPair(cert.CertFile, cert.KeyFile)
		if err != nil {
			log.Fatalf("Error loading X.509 key pair: %v", err)
		}
		cert.Certificate = certs

		for _, muxData := range muxers {
			go startHTTPSServer(serverReady, cert, muxData)
		}
	} else {
		for _, muxData := range muxers {
			go startHTTPServer(serverReady, muxData)
		}
	}

	for range muxers {
		<-serverReady
	}
	close(serverReady)

	pkg.SetDownloadLocal(srv.DownloadImageFiles)
	pkg.SetChannelTemplate()
	pkg.SetGameConfig()
}

type ConnectionWatcher struct {
	n int64
}

func (cw *ConnectionWatcher) OnStateChange(_ net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew: //Connection open
		cw.Add(1)
	case http.StateHijacked, http.StateClosed: //Connection Closed
		cw.Add(-1)
	case http.StateActive, http.StateIdle:
		return
	default:
		panic("unhandled default case")
	}
}

func (cw *ConnectionWatcher) Count() int {
	return int(atomic.LoadInt64(&cw.n))
}

func (cw *ConnectionWatcher) Add(c int64) {
	atomic.AddInt64(&cw.n, c)
}
