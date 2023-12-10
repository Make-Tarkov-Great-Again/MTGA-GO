package data

import (
	"log"

	"github.com/gorilla/websocket"
)

type Connect struct {
	*websocket.Conn
}

func DeleteConnection(sessionID string) {
	if _, ok := db.cache.server.websocket.GetAndDel(sessionID); ok {
		log.Println("Connection deleted")
		return
	}
	log.Println("Connection does not exist")
}

func SetConnection(sessionID string, conn *websocket.Conn) {
	if _, ok := db.cache.server.websocket.GetOrSet(sessionID, &Connect{conn}); ok {
		log.Println("Websocket connection has already been established for sessionID:", sessionID)
		return
	}
	log.Println("Websocket connection has been established for sessionID:", sessionID)
}

func GetConnection(sessionID string) *Connect {
	conn, ok := db.cache.server.websocket.Get(sessionID)
	if !ok {
		log.Println("Websocket connection has not been established for sessionID:", sessionID, ". Returning nil...")
		return nil
	}

	return conn
}

func (conn *Connect) SendMessage(notification *Notification) error {
	if err := conn.WriteJSON(notification); err != nil {
		return err
	}
	return nil
}
