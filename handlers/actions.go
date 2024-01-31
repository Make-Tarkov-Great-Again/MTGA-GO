package handlers

import (
	"MT-GO/data"
	"MT-GO/pkg"
	"log"
	"net/http"
)

var actionHandlers = map[string]func(map[string]any, string, *data.ProfileChangesEvent){
	"QuestAccept": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.QuestAccept(moveAction["qid"].(string), sessionID, profileChangeEvent)
	},
	"Examine": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.ExamineItem(moveAction, sessionID)
	},
	"Move": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.MoveItemInStash(moveAction, sessionID, profileChangeEvent)
	},
	"Swap": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.SwapItemInStash(moveAction, sessionID, profileChangeEvent)
	},
	"Fold": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.FoldItem(moveAction, sessionID, profileChangeEvent)
	},
	"Merge": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.MergeItem(moveAction, sessionID, profileChangeEvent)
	},
	"Transfer": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.TransferItem(moveAction, sessionID)
	},
	"Split": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.SplitItem(moveAction, sessionID, profileChangeEvent)
	},
	"ApplyInventoryChanges": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.ApplyInventoryChanges(moveAction, sessionID)
	},
	"ReadEncyclopedia": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.ReadEncyclopedia(moveAction, sessionID)
	},
	"TradingConfirm": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.TradingConfirm(moveAction, sessionID, profileChangeEvent)
	},
	"Remove": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.RemoveItem(moveAction, sessionID, profileChangeEvent)
	},
	"CustomizationBuy": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.CustomizationBuy(moveAction, sessionID)
	},
	"CustomizationWear": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.CustomizationWear(moveAction, sessionID)
	},
	"Bind": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.BindItem(moveAction, sessionID)
	},
	"Tag": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.TagItem(moveAction, sessionID)
	},
	"Toggle": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.ToggleItem(moveAction, sessionID)
	},
	"HideoutUpgrade": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.HideoutUpgrade(moveAction, sessionID, profileChangeEvent)
	},
	"HideoutUpgradeComplete": func(moveAction map[string]any, sessionID string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.HideoutUpgradeComplete(moveAction, sessionID, profileChangeEvent)
	},
}

const (
	actionLog          string = "[ %d / %d ] Action: %s\n"
	actionNotSupported string = "%s is not supported, sending empty response\n"
)

func MainItemsMoving(w http.ResponseWriter, r *http.Request) {
	body := pkg.GetParsedBody(r).(map[string]any)["data"].([]any)
	length := int8(len(body))

	sessionID := pkg.GetSessionID(r)
	profileChangeEvent := data.GetProfileChangesEvent(sessionID)

	for i := int8(0); i != length; i++ {
		moveAction := body[i].(map[string]any)
		action := moveAction["Action"].(string)
		log.Printf(actionLog, i, length-1, action)

		if handler, ok := actionHandlers[action]; ok {
			handler(moveAction, sessionID, profileChangeEvent)
		} else {
			log.Printf(actionNotSupported, action)
		}
	}

	character, err := data.GetCharacterByID(sessionID)
	if err != nil {
		log.Fatal(err)
	}

	if err := character.SaveCharacter(); err != nil {
		log.Fatalln(err)
	}
	pkg.SendZlibJSONReply(w, pkg.ApplyResponseBody(profileChangeEvent))
}
