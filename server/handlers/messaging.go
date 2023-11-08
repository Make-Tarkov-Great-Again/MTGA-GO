package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"encoding/json"
	"log"
	"net/http"
)

type friendsList struct {
	Friends      []database.FriendRequest
	Ignore       []string
	InIgnoreList []string
}

func MessagingFriendList(w http.ResponseWriter, r *http.Request) {
	friends, err := database.GetFriendsByID(services.GetSessionID(r))
	if err != nil {
		log.Fatalln(err)
	}

	body := services.ApplyResponseBody(friendsList{
		Friends:      friends.Friends,
		Ignore:       friends.Ignore,
		InIgnoreList: friends.InIgnoreList,
	})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingDialogList(w http.ResponseWriter, r *http.Request) {
	dialogues, err := database.GetDialogueByID(services.GetSessionID(r))
	if err != nil {
		log.Fatalln(err)
	}

	data := make([]*database.DialogueInfo, 0, len(*dialogues))
	for _, dialogue := range *dialogues {
		dialog := dialogue.CreateDialogListEntry()

		data = append(data, dialog)
	}

	body := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingFriendRequestInbox(w http.ResponseWriter, r *http.Request) {
	friends, err := database.GetFriendsByID(services.GetSessionID(r))
	if err != nil {
		log.Fatalln(err)
	}

	body := &FriendRequestMailbox{
		Data: friends.FriendRequestInbox,
	}
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingFriendRequestOutbox(w http.ResponseWriter, r *http.Request) {
	friends, err := database.GetFriendsByID(services.GetSessionID(r))
	if err != nil {
		log.Fatalln(err)
	}

	body := &FriendRequestMailbox{
		Data: friends.FriendRequestOutbox,
	}
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogInfo(w http.ResponseWriter, r *http.Request) {
	parsedData := services.GetParsedBody(r)
	dialogId, _ := parsedData.(map[string]any)["dialogId"].(string)

	dialogues, err := database.GetDialogueByID(services.GetSessionID(r))
	if err != nil {
		log.Fatalln(err)
	}

	dialog, ok := (*dialogues)[dialogId]
	if !ok {
		log.Println("Dialogue does not exist, check ID:", dialogId, ". We crash!")
		return
	}

	dialogInfo := dialog.CreateQuestDialogueInfo()

	body := services.ApplyResponseBody(dialogInfo)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogView(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)

	request := new(DialogView)
	rBody, _ := json.Marshal(services.GetParsedBody(r))
	err := json.Unmarshal(rBody, request)
	if err != nil {
		log.Println("Data invalid in MessagingMailDialogView")
	}

	data := new(database.DialogMessageView)

	dialogues, err := database.GetDialogueByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

	dialog, ok := (*dialogues)[request.DialogID]
	if !ok {
		log.Println("Dialogue does not exist, check ID:", request.DialogID, ".")
		data.Messages = make([]database.DialogMessage, 0)
		data.Profiles = make([]database.DialogUserInfo, 0)
		data.HasMessageWithRewards = false

		body := services.ApplyResponseBody(data)
		services.ZlibJSONReply(w, r.RequestURI, body)
		return
	}

	switch request.Type {
	case 2:
		data.Messages = dialog.Messages
		data.Profiles = make([]database.DialogUserInfo, 0)
		data.HasMessageWithRewards = dialog.HasMessagesWithRewards()
	case 1:
	case 6:
		log.Println("WE HAVENT GOTTEN HERE YET BUDDY")

		data.Messages = dialog.Messages
		//TODO: handle profiles
		data.Profiles = make([]database.DialogUserInfo, 0)
		data.HasMessageWithRewards = dialog.HasMessagesWithRewards()
	default:
		log.Println("Request Type:", request.Type, "unsupported at the moment!")
	}

	body := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, r.RequestURI, body)
	if dialog.New != 0 {
		dialog.New = 0
		dialog.AttachmentsNew = dialog.GetUnreadMessagesWithAttachments()
		dialogues.SaveDialogue(sessionID)
	}
}

func MessagingMailDialogPin(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	dialogId := services.GetParsedBody(r).(map[string]any)["dialogId"].(string)

	dialogues, err := database.GetDialogueByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

	dialog, ok := (*dialogues)[dialogId]
	if !ok {
		log.Println("Dialogue does not exist, check ID:", dialogId, ". We crash!")
		return
	}

	dialog.Pinned = true
	dialogues.SaveDialogue(sessionID)

	body := services.ApplyResponseBody([]struct{}{})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogUnpin(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)

	parsedData := services.GetParsedBody(r)
	dialogId, _ := parsedData.(map[string]any)["dialogId"].(string)

	dialogues, err := database.GetDialogueByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

	dialog, ok := (*dialogues)[dialogId]
	if !ok {
		log.Println("Dialogue does not exist, check ID:", dialogId, ". We crash!")
		return
	}

	dialog.Pinned = false
	dialogues.SaveDialogue(sessionID)

	body := services.ApplyResponseBody([]struct{}{})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogRemove(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)

	parsedData := services.GetParsedBody(r)
	dialogId, _ := parsedData.(map[string]any)["dialogId"].(string)

	dialogues, err := database.GetDialogueByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

	delete(*dialogues, dialogId)

	dialogues.SaveDialogue(sessionID)

	body := services.ApplyResponseBody([]struct{}{})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MessagingMailDialogClear(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody([]struct{}{})
	services.ZlibJSONReply(w, r.RequestURI, body)
}
