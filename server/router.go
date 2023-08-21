// Package routes houses all routes to the client and server
package server

import (
	"MT-GO/server/handlers"
	"net/http"
)

// SetRoutes sets all existing routes
func setRoutes(main *http.ServeMux /* , main *http.ServeMux */) {
	setMainRoutes(main)
	//setTradingRoutes(trading)
	//setRagfairRoutes(ragfair)
	//setMessagingRoutes(messaging)
}

func setMainRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/WebSocketAddress", handlers.GetWebSocketAddress)

	mux.HandleFunc("/getBundleList", handlers.GetBundleList)

	mux.HandleFunc("/client/raid/person/killed/showMessage", handlers.ShowPersonKilledMessage)

	mux.HandleFunc("/client/game/start", handlers.MainGameStart)

	mux.HandleFunc("/client/menu/locale/", handlers.MainMenuLocale)

	mux.HandleFunc("/client/game/version/validate", handlers.MainVersionValidate)

	mux.HandleFunc("/client/languages", handlers.MainLanguages)

	mux.HandleFunc("/client/game/config", handlers.MainGameConfig)

	mux.HandleFunc("/client/items", handlers.MainItems)

	mux.HandleFunc("/client/customization", handlers.MainCustomization)

	mux.HandleFunc("/client/globals", handlers.MainGlobals)

	mux.HandleFunc("/client/settings", handlers.MainSettings)

	mux.HandleFunc("/client/game/profile/list", handlers.MainProfileList)

	mux.HandleFunc("/client/account/customization", handlers.MainAccountCustomization)

	mux.HandleFunc("/client/locale/", handlers.MainLocale)

	mux.HandleFunc("/client/game/keepalive", handlers.MainKeepAlive)

	mux.HandleFunc("/client/game/profile/nickname/reserved", handlers.MainNicknameReserved)

	mux.HandleFunc("/client/game/profile/nickname/validate", handlers.MainNicknameValidate)

	mux.HandleFunc("/client/game/profile/create", handlers.MainProfileCreate)

	mux.HandleFunc("/client/game/profile/select", handlers.MainProfileSelect)

	mux.HandleFunc("/client/profile/status", handlers.MainProfileStatus)

	mux.HandleFunc("/client/weather", handlers.MainWeather)

	mux.HandleFunc("/client/locations", handlers.MainLocations)

	mux.HandleFunc("/client/handbook/templates", handlers.MainTemplates)

	mux.HandleFunc("/client/hideout/areas", handlers.MainHideoutAreas)
	mux.HandleFunc("/client/hideout/qte/list", handlers.MainHideoutQTE)
	mux.HandleFunc("/client/hideout/settings", handlers.MainHideoutSettings)
	mux.HandleFunc("/client/hideout/production/recipes", handlers.MainHideoutRecipes)
	mux.HandleFunc("/client/hideout/production/scavcase/recipes", handlers.MainHideoutScavRecipes)

	mux.HandleFunc("/client/handbook/builds/my/list", handlers.MainBuildsList)

	mux.HandleFunc("/client/notifier/channel/create", handlers.MainChannelCreate)

	mux.HandleFunc("/client/quest/list", handlers.MainQuestList)

	mux.HandleFunc("/client/match/group/current", handlers.MainCurrentGroup)

	mux.HandleFunc("/client/repeatalbeQuests/activityPeriods", handlers.MainRepeatableQuests)

	mux.HandleFunc("/client/server/list", handlers.MainServerList)

	mux.HandleFunc("/client/checkVersion", handlers.MainCheckVersion)

	mux.HandleFunc("/client/game/logout", handlers.MainLogoout)
}

// /client/match/group/current

// /client/server/list

// /client/repeatalbeQuests/activityPeriods

func setTradingRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/trading/api/traderSettings", handlers.TradingTraderSettings)
	mux.HandleFunc("/client/trading/customization/storage", handlers.TradingCustomizationStorage)
	// "/client/trading/customization/storage"
	// "/client/trading/api/getTraderAssort/" + traderId
	// "/client/trading/customization/" + traderId + "/offers"
	// "/client/trading/api/traderSettings"
}

func setRagfairRoutes(mux *http.ServeMux) {
	// "/client/ragfair/offer/findbyid"
	// "/client/ragfair/itemMarketPrice"
	// "/client/ragfair/find"
}

/*
Incoming [ POST ] Request URL: [ /client/trading/customization/storage  */

func setMessagingRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/client/friend/list", handlers.MessagingFriendList)
	// "/client/friend/list"

	mux.HandleFunc("/client/mail/dialog/list", handlers.MessagingDialogList)
	// "/client/mail/dialog/list"

	mux.HandleFunc("/client/friend/request/list/inbox", handlers.MessagingFriendRequestInbox)
	// "/client/friend/request/list/inbox"

	mux.HandleFunc("/client/friend/request/list/outbox", handlers.MessagingFriendRequestOutbox)
	// "/client/friend/request/list/outbox"

	// "/client/friend/delete"

	// "/client/friend/request/accept-all"
	// "/client/friend/request/decline"
	// "/client/friend/request/accept"
	// "/client/friend/request/cancel"
	// "/client/friend/request/send"
	// "/client/friend/request/decline"

	// "/client/friend/ignore/remove"

	// "/client/mail/dialog/getAllAttachments"
	// "/client/mail/dialog/clear"
	// "/client/mail/dialog/remove"
	// "/client/mail/dialog/view"
	// "/client/mail/dialog/pin"
	// "/client/mail/dialog/unpin"
	// "/client/mail/dialog/info"
	// "/client/mail/dialog/read"

	// "/client/mail/dialog/group/create"
	// "/client/mail/dialog/group/leave"

	// "/client/mail/dialog/group/owner/change"

	// "/client/mail/dialog/group/users/add"
	// "/client/mail/dialog/group/users/remove"

	// "/client/mail/msg/send"

}
