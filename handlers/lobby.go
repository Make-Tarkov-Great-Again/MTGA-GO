package handlers

import (
	"MT-GO/pkg"
	"log"
	"net/http"
)

func LobbyPushNotifier(w http.ResponseWriter, _ *http.Request) {
	log.Println("Push notification")
	body := pkg.ApplyResponseBody([]any{})
	pkg.SendZlibJSONReply(w, body)
}

func LobbyGetWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Websocket for, " + r.URL.String())
	sessionID := pkg.GetSessionID(r)
	if err := pkg.SendQueuedMessagesToPlayer(sessionID); err != nil {
		log.Println(err)
	}

	body := pkg.GetWebSocket(sessionID)
	pkg.SendZlibJSONReply(w, body)
}
