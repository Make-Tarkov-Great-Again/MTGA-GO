package database

import (
	"MT-GO/tools"
	"fmt"
	"path/filepath"
)

var MessageType = map[string]int8{
	"User":         1,
	"Trader":       2,
	"Auction":      3,
	"Flea":         4,
	"Admin":        5,
	"Group":        6,
	"System":       7,
	"Insurance":    8,
	"Global":       9,
	"QuestStart":   10,
	"QuestFail":    11,
	"QuestSuccess": 12,
	"GiftMessage":  13,
	"Support":      14,
}

type Dialog struct {
	ID             string           `json:"_id"`
	Type           int8             `json:"type"`
	Messages       []DialogMessage  `json:"messages"`
	Pinned         bool             `json:"pinned"`
	New            int8             `json:"new"`
	AttachmentsNew int8             `json:"attachmentsNew"`
	Users          []DialogUserInfo `json:"users"`
}

type DialogUserInfo struct {
	ID   string            `json:"_id"`
	Info DialogUserDetails `json:"info"`
}

type DialogUserDetails struct {
	Nickname       string `json:"Nickname"`
	Side           string `json:"Side"`
	Level          int8   `json:"Level"`
	MemberCategory int8   `json:"MemberCategory"`
}

type DialogMessage struct {
	ID                  string        `json:"_id"`
	UID                 string        `json:"uid"`
	Type                int8          `json:"type"`
	DT                  int64         `json:"dt"`
	Text                string        `json:"text"`
	TemplateID          string        `json:"templateId,omitempty"`
	HasRewards          bool          `json:"hasRewards"`
	RewardCollected     bool          `json:"rewardCollected"`
	SystemData          string        `json:"systemData,omitempty"`
	ProfileChangeEvents []interface{} `json:"profileChangeEvents,omitempty"`
}

type DialogueDetails struct {
	RecipientID                    string        `json:"recipientId"`
	Sender                         int8          `json:"sender"`
	DialogType                     int8          `json:"dialogType"`
	Trader                         string        `json:"trader"`
	TemplateID                     string        `json:"template"`
	Items                          []interface{} `json:"items,omitempty"`
	ItemsMaxStorageLifetimeSeconds int           `json:"itemsMaxStorageLifetimeSeconds,omitempty"`
}

type DialogueInfo struct {
	ID             string               `json:"_id"`
	Type           int8                 `json:"type"`
	Message        *DialogueInfoMessage `json:"message"`
	AttachmentsNew int8                 `json:"attachmentsNew"`
	New            int8                 `json:"new"`
	Pinned         bool                 `json:"pinned"`
	Users          []DialogueUsers      `json:"Users"`
}

type DialogueInfoMessage struct {
	DT         int64  `json:"dt"`
	Type       int8   `json:"type"`
	TemplateID string `json:"templateId"`
	UID        string `json:"uid"`
	Text       string `json:"text,omitempty"`
	SystemData string `json:"systemData,omitempty"`
}

type DialogueUsers struct {
}

func (d Dialog) CreateQuestDialogueInfo() *DialogueInfo {
	info := &DialogueInfo{
		ID:             d.ID,
		Type:           d.Type,
		Message:        d.CreateDialogueInfoMessage(),
		New:            d.New,
		AttachmentsNew: d.AttachmentsNew,
		Pinned:         d.Pinned,
	}

	if d.Users != nil {
		//TODO: DEAL WITH USERS
		fmt.Println("No users were created")
	}

	return info
}

func (d Dialog) CreateDialogueInfoMessage() *DialogueInfoMessage {
	message := d.Messages[len(d.Messages)-1]

	return &DialogueInfoMessage{
		DT:         message.DT,
		Type:       message.Type,
		TemplateID: message.TemplateID,
		UID:        message.UID,
		Text:       message.Text,
	}
}

func (d Dialog) CreateDialogueUsers() []DialogueUsers {
	//users := make([]DialogueUsers, 0)
	return nil
}

func GetDialogueByUID(uid string) *Dialogue {
	if profile, ok := profiles[uid]; ok {
		return profile.Dialogue
	}

	fmt.Println("Profile with UID ", uid, " does not have dialogue")
	return nil
}

func (d Dialogue) SaveDialogue(sessionID string) {
	dialogueFilePath := filepath.Join(profilesPath, sessionID, "dialogue.json")

	err := tools.WriteToFile(dialogueFilePath, d)
	if err != nil {
		panic(err)
	}
	fmt.Println("Dialogue saved")
}

func CreateQuestDialogue(playerID string, sender string, traderID string, dialogueID string) (*Dialog, *DialogMessage) {
	contents := &DialogueDetails{
		RecipientID: playerID,
		Sender:      MessageType[sender],
		DialogType:  MessageType["Trader"],
		Trader:      traderID,
		TemplateID:  dialogueID,
	}

	dialog := &Dialog{
		ID:             traderID,
		Type:           contents.DialogType,
		Pinned:         false,
		Messages:       []DialogMessage{},
		New:            0,
		AttachmentsNew: 0,
	}

	message := &DialogMessage{
		ID:         tools.GenerateMongoID(),
		UID:        traderID,
		Type:       contents.Sender,
		DT:         tools.GetCurrentTimeInSeconds(),
		Text:       "",
		TemplateID: dialogueID,
	}

	//dialog.Messages = append(dialog.Messages)

	return dialog, message
}
