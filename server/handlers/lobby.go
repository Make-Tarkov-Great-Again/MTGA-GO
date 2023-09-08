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
	body := fmt.Sprintf("%s/getwebsocket/%s", database.GetWebSocketAddress(), services.GetSessionID(r))
	services.ZlibJSONReply(w, body)
}

func LobbySWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SWS")
}

func LobbyNotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NOTIFY for," + r.URL.String())
}
