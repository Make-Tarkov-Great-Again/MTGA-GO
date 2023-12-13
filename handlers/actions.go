package handlers

import (
	"MT-GO/data"
	"MT-GO/pkg"
	"log"
	"net/http"
)

var actionHandlers = map[string]func(map[string]any, string, *data.ProfileChangesEvent){
	"QuestAccept": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.QuestAccept(moveAction["qid"].(string), id, profileChangeEvent)
	},
	"Examine": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.ExamineItem(moveAction, id)
	},
	"Move": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.MoveItemInStash(moveAction, id, profileChangeEvent)
	},
	"Swap": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.SwapItemInStash(moveAction, id, profileChangeEvent)
	},
	"Fold": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.FoldItem(moveAction, id, profileChangeEvent)
	},
	"Merge": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.MergeItem(moveAction, id, profileChangeEvent)
	},
	"Transfer": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.TransferItem(moveAction, id)
	},
	"Split": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.SplitItem(moveAction, id, profileChangeEvent)
	},
	"ApplyInventoryChanges": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.ApplyInventoryChanges(moveAction, id)
	},
	"ReadEncyclopedia": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.ReadEncyclopedia(moveAction, id)
	},
	"TradingConfirm": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.TradingConfirm(moveAction, id, profileChangeEvent)
	},
	"Remove": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.RemoveItem(moveAction, id, profileChangeEvent)
	},
	"CustomizationBuy": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.CustomizationBuy(moveAction, id)
	},
	"CustomizationWear": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.CustomizationWear(moveAction, id)
	},
	"Bind": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.BindItem(moveAction, id)
	},
	"Tag": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.TagItem(moveAction, id)
	},
	"Toggle": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.ToggleItem(moveAction, id)
	},
	"HideoutUpgrade": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.HideoutUpgrade(moveAction, id, profileChangeEvent)
	},
	"HideoutUpgradeComplete": func(moveAction map[string]any, id string, profileChangeEvent *data.ProfileChangesEvent) {
		pkg.HideoutUpgradeComplete(moveAction, id, profileChangeEvent)
	},
}

const (
	actionLog          string = "[ %d / %d ] Action: %s\n"
	actionNotSupported        = "%s is not supported, sending empty response\n"
)

func MainItemsMoving(w http.ResponseWriter, r *http.Request) {
	body := pkg.GetParsedBody(r).(map[string]any)["data"].([]any)
	length := int8(len(body)) - 1

	sessionID := pkg.GetSessionID(r)
	profileChangeEvent := data.GetProfileChangesEvent(sessionID)

	for i := int8(0); i != length; i++ {
		moveAction := body[i].(map[string]any)
		action := moveAction["Action"].(string)
		log.Printf(actionLog, i, length, action)

		if handler, ok := actionHandlers[action]; ok {
			handler(moveAction, sessionID, &profileChangeEvent)
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
