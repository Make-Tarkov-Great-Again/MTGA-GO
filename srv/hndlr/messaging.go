package hndlr

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
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingDialogList(w http.ResponseWriter, r *http.Request) {
	dialogList, err := pkg.GetDialogueList(r)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(dialogList)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingFriendRequestInbox(w http.ResponseWriter, r *http.Request) {
	body, err := pkg.GetFriendRequestInbox(r)
	if err != nil {
		log.Println(err)
	}
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingFriendRequestOutbox(w http.ResponseWriter, r *http.Request) {
	body, err := pkg.GetFriendRequestOutbox(r)
	if err != nil {
		log.Println(err)
	}

	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogInfo(w http.ResponseWriter, r *http.Request) {
	dialogInfo, err := pkg.GetMailDialogInfo(r)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(dialogInfo)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogView(w http.ResponseWriter, r *http.Request) {
	dialogView, err := pkg.GetMailDialogView(r)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(dialogView)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogPin(w http.ResponseWriter, r *http.Request) {
	if err := pkg.PinMailDialog(r); err != nil {
		log.Println(err)
	}
	body := pkg.ApplyResponseBody([]struct{}{})
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogUnpin(w http.ResponseWriter, r *http.Request) {
	if err := pkg.UnpinMailDialog(r); err != nil {
		log.Println(err)
	}
	body := pkg.ApplyResponseBody([]struct{}{})
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogRemove(w http.ResponseWriter, r *http.Request) {
	if err := pkg.RemoveMailDialog(r); err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody([]struct{}{})
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogClear(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody([]struct{}{})
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}
