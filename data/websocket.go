package data

import (
	"log"

	"github.com/gorilla/websocket"
)

var connections = make(map[string]*Connect)

type Connect struct {
	*websocket.Conn
}

func DeleteConnection(sessionID string) {
	_, ok := connections[sessionID]
	if ok {
		log.Println("Connection deleted")

		//TODO: Check if they were in raid, if they were they lose their stuff

		delete(connections, sessionID)
		return
	}
	log.Println("Connection does not exist to delete")
	return
}

func SetConnection(sessionID string, conn *websocket.Conn) {
	_, ok := connections[sessionID]
	if ok {
		log.Println("Websocket connection has already been established for sessionID:", sessionID)
		return
	}

	connection := &Connect{conn}
	connections[sessionID] = connection
	log.Println("Websocket connection has been established for sessionID:", sessionID)
}

func GetConnection(sessionID string) *Connect {
	conn, ok := connections[sessionID]
	if !ok {
		log.Println("Websocket connection has not been established for sessionID:", sessionID, ". Returning nil...")
		return nil
	}

	return conn
}

func (conn *Connect) SendMessage(notification *Notification) {
	err := conn.WriteJSON(notification)
	if err != nil {
		log.Println(err)
		return
	}
}
