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
	ID             string          `json:"_id"`
	Type           int8            `json:"type"`
	Messages       []DialogMessage `json:"messages"`
	Pinned         bool            `json:"pinned"`
	New            int8            `json:"new"`
	AttachmentsNew int8            `json:"attachmentsNew"`
	Users          []DialogUser    `json:"users"`
}

type DialogUser struct {
	ID   string         `json:"_id"`
	Info DialogUserInfo `json:"info"`
}

type DialogUserInfo struct {
	Nickname       string `json:"Nickname"`
	Side           string `json:"Side"`
	Level          int8   `json:"Level"`
	MemberCategory int8   `json:"MemberCategory"`
}

type DialogMessage struct {
	ID                  string                 `json:"_id"`
	UID                 string                 `json:"uid"`
	Type                int8                   `json:"type"`
	DT                  int64                  `json:"dt"`
	UtcDateTime         int64                  `json:"UtcDateTime,omitempty"`
	Member              map[string]interface{} `json:"Member,omitempty"`
	Text                string                 `json:"text"`
	TemplateID          string                 `json:"templateId,omitempty"`
	Items               []interface{}          `json:"items,omitempty"`
	HasRewards          bool                   `json:"hasRewards"`
	RewardCollected     bool                   `json:"rewardCollected"`
	MaxStorageTime      int64                  `json:"maxStorageTime"`
	SystemData          string                 `json:"systemData,omitempty"`
	ProfileChangeEvents []interface{}          `json:"profileChangeEvents,omitempty"`
}

type DialogMessageView struct {
	Messages              []DialogMessage  `json:"messages"`
	Profiles              []DialogUserInfo `json:"profiles"`
	HasMessageWithRewards bool             `json:"hasMessageWithRewards"`
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
	Users          []DialogUserInfo     `json:"Users,omitempty"`
}

type DialogueInfoMessage struct {
	DT         int64  `json:"dt"`
	Type       int8   `json:"type"`
	TemplateID string `json:"templateId"`
	UID        string `json:"uid"`
	Text       string `json:"text,omitempty"`
	SystemData string `json:"systemData,omitempty"`
}

func (d *Dialog) CreateQuestDialogueInfo() *DialogueInfo {
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

func (d *Dialog) CreateDialogueInfoMessage() *DialogueInfoMessage {
	message := d.Messages[len(d.Messages)-1]

	return &DialogueInfoMessage{
		DT:         message.DT,
		Type:       message.Type,
		TemplateID: message.TemplateID,
		UID:        message.UID,
		Text:       message.Text,
	}
}

func (d *Dialog) CreateDialogueUsers() []DialogUserInfo {
	//users := make([]DialogueUsers, 0)
	return nil
}

func (d *Dialog) HasMessagesWithRewards() bool {
	for _, message := range d.Messages {
		if len(message.Items) > 0 {
			return true
		}
	}
	return false
}

func (d *Dialog) GetUnreadMessagesWithAttachments() int8 {
	active := d.GetActiveMessages()
	var attachmentCount int8 = 0

	for _, message := range active {
		if message.HasRewards && !message.RewardCollected {
			attachmentCount++
		}
	}

	return attachmentCount
}

func (d *Dialog) GetActiveMessages() []DialogMessage {
	messages := make([]DialogMessage, 0, len(d.Messages))

	time := tools.GetCurrentTimeInSeconds()
	for _, message := range d.Messages {
		if (message.DT + message.MaxStorageTime) > time {
			messages = append(messages, message)
		}
	}
	return messages
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
		ID:             tools.GenerateMongoID(),
		UID:            traderID,
		Type:           contents.Sender,
		DT:             tools.GetCurrentTimeInSeconds(),
		Text:           "",
		TemplateID:     dialogueID,
		MaxStorageTime: tools.GetCurrentTimeInSeconds() + 3600,
	}

	return dialog, message
}
