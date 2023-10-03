package database

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

var connections = make(map[string]*Connect)

type Connect struct {
	*websocket.Conn
}

func SetConnection(sessionID string, conn *websocket.Conn) {
	_, ok := connections[sessionID]
	if ok {
		fmt.Println("Websocket connection has already been established for sessionID:", sessionID)
		return
	}

	connection := &Connect{conn}
	connections[sessionID] = connection
	fmt.Println("Websocket connection has been established for sessionID:", sessionID)
}

func GetConnection(sessionID string) *Connect {
	conn, ok := connections[sessionID]
	if !ok {
		fmt.Println("Websocket connection has not been established for sessionID:", sessionID, ". Returning nil...")
		return nil
	}

	return conn
}

func (conn *Connect) SendMessage(notification *Notification) {
	err := conn.WriteJSON(notification)
	if err != nil {
		log.Fatalln(err)
	}
}
