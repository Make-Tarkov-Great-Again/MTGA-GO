package data

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"path/filepath"

	"MT-GO/tools"
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
	Users          []DialogUser    `json:"users,omitempty"`
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
	ID                  string         `json:"_id"`
	UID                 string         `json:"uid"`
	Type                int8           `json:"type"`
	DT                  int32          `json:"dt"`
	UtcDateTime         int32          `json:"UtcDateTime,omitempty"`
	Member              map[string]any `json:"Member,omitempty"`
	Text                string         `json:"text"`
	TemplateID          string         `json:"templateId,omitempty"`
	Items               []any          `json:"items,omitempty"`
	HasRewards          bool           `json:"hasRewards"`
	RewardCollected     bool           `json:"rewardCollected"`
	MaxStorageTime      int32          `json:"maxStorageTime,omitempty"`
	SystemData          string         `json:"systemData,omitempty"`
	ProfileChangeEvents []any          `json:"profileChangeEvents"`
}

type DialogMessageView struct {
	Messages              []DialogMessage  `json:"messages"`
	Profiles              []DialogUserInfo `json:"profiles"`
	HasMessageWithRewards bool             `json:"hasMessageWithRewards"`
}

type DialogueDetails struct {
	RecipientID                    string `json:"recipientId"`
	Sender                         int8   `json:"sender"`
	DialogType                     int8   `json:"dialogType"`
	Trader                         string `json:"trader"`
	TemplateID                     string `json:"template"`
	Items                          []any  `json:"items,omitempty"`
	ItemsMaxStorageLifetimeSeconds int32  `json:"itemsMaxStorageLifetimeSeconds,omitempty"`
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
	DT         int32  `json:"dt"`
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

	if d.Users != nil && len(d.Users) > 0 {
		//TODO: DEAL WITH USERS
		log.Println("No users were created")
	}

	return info
}

func (d *Dialog) CreateDialogListEntry() *DialogueInfo {
	info := &DialogueInfo{
		ID:             d.ID,
		Type:           d.Type,
		AttachmentsNew: d.AttachmentsNew,
		New:            d.New,
		Pinned:         d.Pinned,
	}

	if d.Type == 1 {
		log.Println("NOT DONE YET")
	} else {
		info.Message = d.CreateDialogueInfoMessage()
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

	time := int32(tools.GetCurrentTimeInSeconds())
	for _, message := range d.Messages {
		if (message.DT + message.MaxStorageTime) > time {
			messages = append(messages, message)
		}
	}
	return messages
}

func setDialogue(path string) *Dialogue {
	output := make(Dialogue)

	data := tools.GetJSONRawMessage(path)
	if err := json.Unmarshal(data, &output); err != nil {
		log.Fatalln(err)
	}

	return &output
}

func GetDialogueByID(uid string) (*Dialogue, error) {
	profile, err := GetProfileByUID(uid)
	if err != nil {
		return nil, err
	}

	if profile.Dialogue != nil {
		return profile.Dialogue, nil
	}

	return nil, fmt.Errorf(dialogueNotExist, uid)
}

func (d Dialogue) SaveDialogue(sessionID string) error {
	dialogueFilePath := filepath.Join(profilesPath, sessionID, "dialogue.json")

	err := tools.WriteToFile(dialogueFilePath, d)
	if err != nil {
		return fmt.Errorf(dialogueNotSaved, sessionID, err)
	}
	log.Println("Dialogue saved")
	return nil
}

const (
	dialogueNotSaved string = "Dialogue for %s was not saved: %s"
	dialogueNotExist string = "Dialogue for %s does not exist"
	redeemTime       int32  = 172800 // TODO: remove this and put in config (hours x 3600)
)

func CreateQuestDialogue(playerID string, sender string, traderID string, dialogueID string) (*Dialog, *DialogMessage) {
	contents := &DialogueDetails{
		RecipientID:                    playerID,
		Sender:                         MessageType[sender],
		DialogType:                     MessageType["Trader"],
		Trader:                         traderID,
		TemplateID:                     dialogueID,
		ItemsMaxStorageLifetimeSeconds: redeemTime,
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
		DT:         int32(tools.GetCurrentTimeInSeconds()),
		Text:       "",
		TemplateID: dialogueID,
	}

	return dialog, message
}
