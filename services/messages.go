package services

import (
	"MT-GO/tools"
	"fmt"
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

type NPCmessage struct {
	RecipientID string `json:"recipientId"`
	Sender      int8   `json:"sender"`
	DialogType  int8   `json:"dialogType"`
	Trader      string `json:"trader"`
	TemplateID  string `json:"template"`
}

type NPCmessageWithItems struct {
	*NPCmessage
	Items                          []interface{} `json:"items"`
	ItemsMaxStorageLifetimeSeconds int           `json:"itemsMaxStorageLifetimeSeconds"`
}

const itemsMaxStorageLifetimeSeconds int = 3600

func SendNPCMessage(playerID string, sender string, traderID string, dialogueID string, items []interface{}) {
	message := NPCmessage{
		RecipientID: playerID,
		Sender:      MessageType[sender],
		DialogType:  MessageType["Trader"],
		Trader:      traderID,
		TemplateID:  dialogueID,
	}

	if len(items) == 0 { // send message without items
		fmt.Println()
	} else {
		itemMessage := NPCmessageWithItems{
			NPCmessage:                     &message,
			Items:                          items,
			ItemsMaxStorageLifetimeSeconds: int(tools.GetCurrentTimeInSeconds()) + itemsMaxStorageLifetimeSeconds,
		}
		fmt.Println(itemMessage)
	}
}
