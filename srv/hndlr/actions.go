package hndlr

import (
	"MT-GO/pkg"
	"log"
	"net/http"
)

var actionHandlers = map[string]func(map[string]any, string, *pkg.ProfileChangesEvent){
	"QuestAccept": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.QuestAccept(moveAction["qid"].(string), id, profileChangeEvent)
	},
	"Examine": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.ExamineItem(moveAction, id)
	},
	"Move": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.MoveItemInStash(moveAction, id, profileChangeEvent)
	},
	"Swap": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.SwapItemInStash(moveAction, id, profileChangeEvent)
	},
	"Fold": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.FoldItem(moveAction, id, profileChangeEvent)
	},
	"Merge": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.MergeItem(moveAction, id, profileChangeEvent)
	},
	"Transfer": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.TransferItem(moveAction, id)
	},
	"Split": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.SplitItem(moveAction, id, profileChangeEvent)
	},
	"ApplyInventoryChanges": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.ApplyInventoryChanges(moveAction, id)
	},
	"ReadEncyclopedia": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.ReadEncyclopedia(moveAction, id)
	},
	"TradingConfirm": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.TradingConfirm(moveAction, id, profileChangeEvent)
	},
	"Remove": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.RemoveItem(moveAction, id, profileChangeEvent)
	},
	"CustomizationBuy": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.CustomizationBuy(moveAction, id)
	},
	"CustomizationWear": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.CustomizationWear(moveAction, id)
	},
	"Bind": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.BindItem(moveAction, id)
	},
	"Tag": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.TagItem(moveAction, id)
	},
	"Toggle": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.ToggleItem(moveAction, id)
	},
	"HideoutUpgrade": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.HideoutUpgrade(moveAction, id, profileChangeEvent)
	},
	//HideoutUpgradeComplete
	"HideoutUpgradeComplete": func(moveAction map[string]any, id string, profileChangeEvent *pkg.ProfileChangesEvent) {
		pkg.HideoutUpgradeComplete(moveAction, id, profileChangeEvent)
	},
}

func MainItemsMoving(w http.ResponseWriter, r *http.Request) {
	body := pkg.GetParsedBody(r).(map[string]any)["data"].([]any)
	length := int8(len(body)) - 1

	id := pkg.GetSessionID(r)
	profileChangeEvent := pkg.CreateProfileChangesEvent(id)

	for i, move := range body {
		moveAction := move.(map[string]any)
		action := moveAction["Action"].(string)
		log.Println("[", i, "/", length, "] Action: ", action)

		if handler, ok := actionHandlers[action]; ok {
			handler(moveAction, id, profileChangeEvent)
		} else {
			log.Println(action, "is not supported, sending empty response")
		}
	}

	pkg.ZlibJSONReply(w, r.RequestURI, pkg.ApplyResponseBody(profileChangeEvent))
}
