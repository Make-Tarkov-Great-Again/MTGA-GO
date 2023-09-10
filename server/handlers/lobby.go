package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"fmt"
	"net/http"
)

func LobbyPushNotifier(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Push notification")
	body := services.ApplyResponseBody([]interface{}{})
	services.ZlibJSONReply(w, body)
}

func LobbyGetWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Websocket for, " + r.URL.String())
	sessionID := services.GetSessionID(r)
	storage := database.GetStorageByUID(sessionID)

	connection := database.GetConnection(sessionID)
	for _, v := range storage.Mailbox {
		connection.SendMessage(v)
	}
	storage.Mailbox = make([]*database.Notification, 0, len(storage.Mailbox))
	storage.SaveStorage(sessionID)

	body := fmt.Sprintf("%s/getwebsocket/%s", database.GetWebSocketAddress(), sessionID)
	services.ZlibJSONReply(w, body)
}

func LobbySWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SWS")
}

func LobbyNotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NOTIFY for," + r.URL.String())
}
