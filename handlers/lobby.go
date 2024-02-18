package handlers

import (
	"MT-GO/pkg"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func LobbyPushNotifier(w http.ResponseWriter, r *http.Request) {
	log.Println("Push notification")

	sessionID, err := pkg.GetSessionID(r)
	if err != nil {
		log.Println(err)
		return
	}

	if chi.URLParam(r, "id") != sessionID {
		log.Fatalln("AAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	}
	body := pkg.ApplyResponseBody([]any{})
	pkg.SendZlibJSONReply(w, body)
}

func LobbyGetWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Websocket for, " + r.URL.String())
	sessionID, err := pkg.GetSessionID(r)
	if err != nil {
		log.Println(err)
		return
	}
	if chi.URLParam(r, "id") != sessionID {
		log.Fatalln("AAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	}
	if err := pkg.SendQueuedMessagesToPlayer(sessionID); err != nil {
		log.Println(err)
	}

	body := pkg.GetWebSocket(sessionID)
	pkg.SendZlibJSONReply(w, body)
}
