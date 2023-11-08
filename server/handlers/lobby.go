package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"fmt"
	"log"
	"net/http"
)

func LobbyPushNotifier(w http.ResponseWriter, r *http.Request) {
	log.Println("Push notification")
	body := services.ApplyResponseBody([]any{})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func LobbyGetWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Websocket for, " + r.URL.String())
	sessionID := services.GetSessionID(r)
	storage, err := database.GetStorageByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

	connection := database.GetConnection(sessionID)
	for _, v := range storage.Mailbox {
		connection.SendMessage(v)
	}
	storage.Mailbox = storage.Mailbox[:0]
	storage.SaveStorage(sessionID)

	body := fmt.Sprintf("%s/getwebsocket/%s", database.GetWebSocketAddress(), sessionID)
	services.ZlibJSONReply(w, r.RequestURI, body)
}
