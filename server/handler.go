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
