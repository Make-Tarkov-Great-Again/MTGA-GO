package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"net/http"
)

func MessagingFriendList(w http.ResponseWriter, r *http.Request) {
	friends := database.GetAccountByUID(services.GetSessionID(r)).Friends.Friends
	body := services.ApplyResponseBody(friends)
	services.ZlibJSONReply(w, body)
}

func MessagingDialogList(w http.ResponseWriter, r *http.Request) {
	dialogues := database.GetProfileByUID(services.GetSessionID(r)).Dialogue

	data := []interface{}{}
	for _, dialogue := range dialogues {
		data = append(data, dialogue)
	}

	body := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, body)
	return
}

func MessagingFriendRequestInbox(w http.ResponseWriter, r *http.Request) {
	friends := database.GetAccountByUID(services.GetSessionID(r)).FriendRequestInbox
	body := map[string]interface{}{
		"err":  0,
		"data": friends,
	}
	services.ZlibJSONReply(w, body)
}

func MessagingFriendRequestOutbox(w http.ResponseWriter, r *http.Request) {
	friends := database.GetAccountByUID(services.GetSessionID(r)).FriendRequestOutbox
	body := map[string]interface{}{
		"err":  0,
		"data": friends,
	}
	services.ZlibJSONReply(w, body)
}
