package database

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var connections = make(map[string]*Connect)

type Connect struct {
	*websocket.Conn
}

func SetConnection(sessionID string, conn *websocket.Conn) {
	_, ok := connections[sessionID]
	if !ok {
		fmt.Println("Couldn't set connection because it already exists")
		return
	}

	connection := &Connect{conn}
	connections[sessionID] = connection
}

func GetConnection(sessionID string) *Connect {
	conn, ok := connections[sessionID]
	if !ok {
		return nil
	}

	return conn
}

func (conn *Connect) sendMessage(notification *Notification) {

}
