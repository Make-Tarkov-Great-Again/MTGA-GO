package pkg

import (
	"MT-GO/data"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"net/http"
)

const systemSenderID string = "59e7125688a45068a6249071"
const itemsMaxStorageLifetimeSeconds int = 3600

func SendQueuedMessagesToPlayer(sessionID string) error {
	storage, err := data.GetStorageByID(sessionID)
	if err != nil {
		return err
	}

	connection := data.GetConnection(sessionID)
	for _, v := range storage.Mailbox {
		connection.SendMessage(v)
	}
	storage.Mailbox = storage.Mailbox[:0]
	storage.SaveStorage(sessionID)
	return nil
}

func GetFriendsList(r *http.Request) (*FriendsList, error) {
	friends, err := data.GetFriendsByID(GetSessionID(r))
	if err != nil {
		return nil, err
	}

	return &FriendsList{
		Friends:      friends.Friends,
		Ignore:       friends.Ignore,
		InIgnoreList: friends.InIgnoreList,
	}, nil
}

func GetDialogueList(r *http.Request) ([]*data.DialogueInfo, error) {
	dialogues, err := data.GetDialogueByID(GetSessionID(r))
	if err != nil {
		return nil, err
	}

	output := make([]*data.DialogueInfo, 0, len(*dialogues))
	for _, dialogue := range *dialogues {
		dialog := dialogue.CreateDialogListEntry()

		output = append(output, dialog)
	}
	return output, nil
}

func GetFriendRequestInbox(r *http.Request) (*FriendRequestMailbox, error) {
	friends, err := data.GetFriendsByID(GetSessionID(r))
	if err != nil {
		return nil, err
	}

	return &FriendRequestMailbox{
		Data: friends.FriendRequestInbox,
	}, nil
}

func GetFriendRequestOutbox(r *http.Request) (*FriendRequestMailbox, error) {
	friends, err := data.GetFriendsByID(GetSessionID(r))
	if err != nil {
		return nil, err
	}

	return &FriendRequestMailbox{
		Data: friends.FriendRequestOutbox,
	}, nil
}

const dialogNotExist string = "Dialogue for %s does not exist"

func GetMailDialogInfo(r *http.Request) (*data.DialogueInfo, error) {
	dialogId, _ := GetParsedBody(r).(map[string]any)["dialogId"].(string)

	dialogues, err := data.GetDialogueByID(GetSessionID(r))
	if err != nil {
		return nil, err
	}

	dialog, ok := (*dialogues)[dialogId]
	if !ok {
		return nil, fmt.Errorf(dialogNotExist, dialogId)
	}

	output := dialog.CreateQuestDialogueInfo()
	return output, nil
}

func GetMailDialogView(r *http.Request) (*data.DialogMessageView, error) {
	request := new(DialogView)
	input, _ := json.Marshal(GetParsedBody(r))
	if err := json.Unmarshal(input, request); err != nil {
		return nil, err
	}

	sessionID := GetSessionID(r)
	output := new(data.DialogMessageView)

	dialogues, err := data.GetDialogueByID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

	dialog, ok := (*dialogues)[request.DialogID]
	if !ok {
		log.Println("Dialogue does not exist, check ID:", request.DialogID, ".")
		output.Messages = make([]data.DialogMessage, 0)
		output.Profiles = make([]data.DialogUserInfo, 0)
		output.HasMessageWithRewards = false

		return output, nil
	}

	switch request.Type {
	case 2:
		output.Messages = dialog.Messages
		output.Profiles = make([]data.DialogUserInfo, 0)
		output.HasMessageWithRewards = dialog.HasMessagesWithRewards()
	case 1:
	case 6:
		log.Println("WE HAVENT GOTTEN HERE YET BUDDY")

		output.Messages = dialog.Messages
		//TODO: handle profiles
		output.Profiles = make([]data.DialogUserInfo, 0)
		output.HasMessageWithRewards = dialog.HasMessagesWithRewards()
	default:
		log.Println("Request Type:", request.Type, "unsupported at the moment!")
	}

	if dialog.New != 0 {
		dialog.New = 0
		dialog.AttachmentsNew = dialog.GetUnreadMessagesWithAttachments()
		dialogues.SaveDialogue(sessionID)
	}

	return output, nil
}

func PinMailDialog(r *http.Request) error {
	sessionID := GetSessionID(r)
	dialogId := GetParsedBody(r).(map[string]any)["dialogId"].(string)

	dialogues, err := data.GetDialogueByID(sessionID)
	if err != nil {
		return err
	}

	dialog, ok := (*dialogues)[dialogId]
	if !ok {
		return fmt.Errorf(dialogNotExist, dialogId)
	}

	dialog.Pinned = true
	dialogues.SaveDialogue(sessionID)
	return nil
}

func UnpinMailDialog(r *http.Request) error {
	sessionID := GetSessionID(r)
	dialogId := GetParsedBody(r).(map[string]any)["dialogId"].(string)

	dialogues, err := data.GetDialogueByID(sessionID)
	if err != nil {
		return err
	}

	dialog, ok := (*dialogues)[dialogId]
	if !ok {
		return fmt.Errorf(dialogNotExist, dialogId)
	}

	dialog.Pinned = false
	dialogues.SaveDialogue(sessionID)
	return nil
}

func RemoveMailDialog(r *http.Request) error {
	sessionID := GetSessionID(r)

	parsedData := GetParsedBody(r)
	dialogId, _ := parsedData.(map[string]any)["dialogId"].(string)

	dialogues, err := data.GetDialogueByID(sessionID)
	if err != nil {
		return err
	}

	delete(*dialogues, dialogId)

	dialogues.SaveDialogue(sessionID)
	return nil
}

type FriendsList struct {
	Friends      []data.FriendRequest
	Ignore       []string
	InIgnoreList []string
}

type FriendRequestMailbox struct {
	Err  int   `json:"err"`
	Data []any `json:"data"`
}

type DialogView struct {
	Type     int8   `json:"type"`
	DialogID string `json:"dialogId"`
	Limit    int8   `json:"limit"`
	Time     int64  `json:"time"`
}
