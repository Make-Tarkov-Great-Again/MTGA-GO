package handlers

import (
	"MT-GO/pkg"
	"log"
	"net/http"
)

func MessagingFriendList(w http.ResponseWriter, r *http.Request) {
	friendList, err := pkg.GetFriendsList(r)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(friendList)
	pkg.SendZlibJSONReply(w, body)
}

func MessagingDialogList(w http.ResponseWriter, r *http.Request) {
	dialogList, err := pkg.GetDialogueList(r)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(dialogList)
	pkg.SendZlibJSONReply(w, body)
}

func MessagingFriendRequestInbox(w http.ResponseWriter, r *http.Request) {
	body, err := pkg.GetFriendRequestInbox(r)
	if err != nil {
		log.Println(err)
	}
	pkg.SendZlibJSONReply(w, body)
}

func MessagingFriendRequestOutbox(w http.ResponseWriter, r *http.Request) {
	body, err := pkg.GetFriendRequestOutbox(r)
	if err != nil {
		log.Println(err)
	}

	pkg.SendZlibJSONReply(w, body)
}

func MessagingMailDialogInfo(w http.ResponseWriter, r *http.Request) {
	dialogInfo, err := pkg.GetMailDialogInfo(r)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(dialogInfo)
	pkg.SendZlibJSONReply(w, body)
}

func MessagingMailDialogView(w http.ResponseWriter, r *http.Request) {
	dialogView, err := pkg.GetMailDialogView(r)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(dialogView)
	pkg.SendZlibJSONReply(w, body)
}

func MessagingMailDialogPin(w http.ResponseWriter, r *http.Request) {
	if err := pkg.PinMailDialog(r); err != nil {
		log.Println(err)
	}
	body := pkg.ApplyResponseBody([]struct{}{})
	pkg.SendZlibJSONReply(w, body)
}

func MessagingMailDialogUnpin(w http.ResponseWriter, r *http.Request) {
	if err := pkg.UnpinMailDialog(r); err != nil {
		log.Println(err)
	}
	body := pkg.ApplyResponseBody([]struct{}{})
	pkg.SendZlibJSONReply(w, body)
}

func MessagingMailDialogRemove(w http.ResponseWriter, r *http.Request) {
	if err := pkg.RemoveMailDialog(r); err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody([]struct{}{})
	pkg.SendZlibJSONReply(w, body)
}

func MessagingMailDialogClear(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody([]struct{}{})
	pkg.SendZlibJSONReply(w, body)
}
