package handlers

import (
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
	fmt.Println("Get Websocket")
	body := services.ApplyResponseBody([]interface{}{})
	services.ZlibJSONReply(w, body)
}

func LobbySlash(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/")
	body := services.ApplyResponseBody([]interface{}{})
	services.ZlibJSONReply(w, body)
}

func LobbySWS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SWS")
}

func LobbyNotify(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NOTIFY")
}
