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
