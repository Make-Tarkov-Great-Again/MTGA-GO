package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func MessagingFriendList(w http.ResponseWriter, r *http.Request) {
	friends := database.GetAccountByUID(services.GetSessionID(r)).Friends
	body := services.ApplyResponseBody(friends)
	services.ZlibJSONReply(w, body)
}

func MessagingDialogList(w http.ResponseWriter, r *http.Request) {
	dialogues := database.GetDialogueByUID(services.GetSessionID(r))

	data := make([]*database.DialogueInfo, 0, len(*dialogues))
	for _, dialogue := range *dialogues {
		dialog := dialogue.CreateDialogListEntry()

		data = append(data, dialog)
	}

	body := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, body)
}

func MessagingFriendRequestInbox(w http.ResponseWriter, r *http.Request) {
	friends := database.GetAccountByUID(services.GetSessionID(r)).FriendRequestInbox
	body := &FriendRequestMailbox{
		Data: friends,
	}
	services.ZlibJSONReply(w, body)
}

func MessagingFriendRequestOutbox(w http.ResponseWriter, r *http.Request) {
	friends := database.GetAccountByUID(services.GetSessionID(r)).FriendRequestOutbox
	body := &FriendRequestMailbox{
		Data: friends,
	}
	services.ZlibJSONReply(w, body)
}

func MessagingMailDialogInfo(w http.ResponseWriter, r *http.Request) {
	parsedData := services.GetParsedBody(r)
	dialogId, _ := parsedData.(map[string]interface{})["dialogId"].(string)

	dialogues := *database.GetDialogueByUID(services.GetSessionID(r))
	dialog, ok := dialogues[dialogId]
	if !ok {
		fmt.Println("Dialogue does not exist, check ID:", dialogId, ". We crash!")
		return
	}

	dialogInfo := dialog.CreateQuestDialogueInfo()

	body := services.ApplyResponseBody(dialogInfo)
	services.ZlibJSONReply(w, body)
}

func MessagingMailDialogView(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)

	request := new(DialogView)
	rBody, _ := json.Marshal(services.GetParsedBody(r))
	err := json.Unmarshal(rBody, request)
	if err != nil {
		fmt.Println("Data invalid in MessagingMailDialogView")
	}

	data := new(database.DialogMessageView)

	dialogues := *database.GetDialogueByUID(sessionID)
	dialog, ok := dialogues[request.DialogID]
	if !ok {
		fmt.Println("Dialogue does not exist, check ID:", request.DialogID, ".")
		data.Messages = make([]database.DialogMessage, 0)
		data.Profiles = make([]database.DialogUserInfo, 0)
		data.HasMessageWithRewards = false

		body := services.ApplyResponseBody(data)
		services.ZlibJSONReply(w, body)
		return
	}

	switch request.Type {
	case 2:
		data.Messages = dialog.Messages
		data.Profiles = make([]database.DialogUserInfo, 0)
		data.HasMessageWithRewards = dialog.HasMessagesWithRewards()
	case 1:
	case 6:
		fmt.Println("WE HAVENT GOTTEN HERE YET BUDDY")

		data.Messages = dialog.Messages
		//TODO: handle profiles
		data.Profiles = make([]database.DialogUserInfo, 0)
		data.HasMessageWithRewards = dialog.HasMessagesWithRewards()
	default:
		fmt.Println("Request Type:", request.Type, "unsupported at the moment!")
	}

	body := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, body)
	if dialog.New != 0 {
		dialog.New = 0
		dialog.AttachmentsNew = dialog.GetUnreadMessagesWithAttachments()
		dialogues.SaveDialogue(sessionID)
	}
}

func MessagingMailDialogPin(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	dialogId := services.GetParsedBody(r).(map[string]interface{})["dialogId"].(string)

	dialogues := *database.GetDialogueByUID(sessionID)
	dialog, ok := dialogues[dialogId]
	if !ok {
		fmt.Println("Dialogue does not exist, check ID:", dialogId, ". We crash!")
		return
	}

	dialog.Pinned = true
	dialogues.SaveDialogue(sessionID)

	body := services.ApplyResponseBody([]struct{}{})
	services.ZlibJSONReply(w, body)
}

func MessagingMailDialogUnpin(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)

	parsedData := services.GetParsedBody(r)
	dialogId, _ := parsedData.(map[string]interface{})["dialogId"].(string)

	dialogues := *database.GetDialogueByUID(sessionID)
	dialog, ok := dialogues[dialogId]
	if !ok {
		fmt.Println("Dialogue does not exist, check ID:", dialogId, ". We crash!")
		return
	}

	dialog.Pinned = false
	dialogues.SaveDialogue(sessionID)

	body := services.ApplyResponseBody([]struct{}{})
	services.ZlibJSONReply(w, body)
}

func MessagingMailDialogRemove(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)

	parsedData := services.GetParsedBody(r)
	dialogId, _ := parsedData.(map[string]interface{})["dialogId"].(string)

	dialogues := *database.GetDialogueByUID(sessionID)
	delete(dialogues, dialogId)

	dialogues.SaveDialogue(sessionID)

	body := services.ApplyResponseBody([]struct{}{})
	services.ZlibJSONReply(w, body)
}

func MessagingMailDialogClear(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody([]struct{}{})
	services.ZlibJSONReply(w, body)
}
