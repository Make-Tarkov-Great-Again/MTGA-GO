package server

import (
	"MT-GO/data"
	"MT-GO/pkg"
	"context"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"
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
	defer conn.Close()

	sessionID := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	data.SetConnection(sessionID, conn)

	go func() {
		defer data.DeleteConnection(sessionID)
		defer conn.Close()

		err = conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		if err != nil {
			return
		}
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			err = conn.WriteMessage(messageType, p)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}()
}

const incomingRoute string = "[%s] %s on %s\n"

func logRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(incomingRoute, r.Method, r.URL.Path, strings.TrimPrefix(r.Host, "127.0.0.1"))
		next.ServeHTTP(w, r)
	})
}

func handleWebSocketUpgrade(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if websocket.IsWebSocketUpgrade(r) {
			upgradeToWebsocket(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func decompress(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Length") == "" {
			next.ServeHTTP(w, r)
			return
		}

		buffer, err := pkg.ZlibInflate(r)
		if err != nil {
			log.Println(err)
			return
		}

		if buffer == nil || len(buffer) == 0 {
			next.ServeHTTP(w, r)
			return
		}

		//TODO: Refactor to replace r.Body ((remove CTX))
		var parsedData map[string]any
		if err := json.UnmarshalNoEscape(buffer, &parsedData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), &pkg.ContextKey{}, parsedData)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
