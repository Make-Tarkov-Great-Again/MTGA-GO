// Package routes houses all routes to the client and server
package server

import (
	"MT-GO/server/handlers"
	"MT-GO/services"
	"fmt"
	"net/http"
)

var mainRouteHandlers = map[string]http.HandlerFunc{
	"/getBundleList":                              handlers.GetBundleList,
	"/client/raid/person/killed/showMessage":      handlers.ShowPersonKilledMessage,
	"/client/game/start":                          handlers.MainGameStart,
	"/client/putMetrics":                          handlers.MainPutMetrics,
	"/client/menu/locale/":                        handlers.MainMenuLocale,
	"/client/game/version/validate":               handlers.MainVersionValidate,
	"/client/languages":                           handlers.MainLanguages,
	"/client/game/config":                         handlers.MainGameConfig,
	"/client/items":                               handlers.MainItems,
	"/client/customization":                       handlers.MainCustomization,
	"/client/globals":                             handlers.MainGlobals,
	"/client/settings":                            handlers.MainSettings,
	"/client/game/profile/list":                   handlers.MainProfileList,
	"/client/account/customization":               handlers.MainAccountCustomization,
	"/client/locale/":                             handlers.MainLocale,
	"/client/game/keepalive":                      handlers.MainKeepAlive,
	"/client/game/profile/nickname/reserved":      handlers.MainNicknameReserved,
	"/client/game/profile/nickname/validate":      handlers.MainNicknameValidate,
	"/client/game/profile/create":                 handlers.MainProfileCreate,
	"/client/game/profile/select":                 handlers.MainProfileSelect,
	"/client/profile/status":                      handlers.MainProfileStatus,
	"/client/weather":                             handlers.MainWeather,
	"/client/locations":                           handlers.MainLocations,
	"/client/handbook/templates":                  handlers.MainTemplates,
	"/client/hideout/areas":                       handlers.MainHideoutAreas,
	"/client/hideout/qte/list":                    handlers.MainHideoutQTE,
	"/client/hideout/settings":                    handlers.MainHideoutSettings,
	"/client/hideout/production/recipes":          handlers.MainHideoutRecipes,
	"/client/hideout/production/scavcase/recipes": handlers.MainHideoutScavRecipes,
	"/client/handbook/builds/my/list":             handlers.MainBuildsList,
	"/client/quest/list":                          handlers.MainQuestList,
	"/client/match/group/current":                 handlers.MainCurrentGroup,
	"/client/repeatalbeQuests/activityPeriods":    handlers.MainRepeatableQuests,
	"/client/server/list":                         handlers.MainServerList,
	"/client/checkVersion":                        handlers.MainCheckVersion,
	"/client/game/logout":                         handlers.MainLogoout,
	"/client/items/prices/":                       handlers.MainPrices,

	"/client/notifier/channel/create": handlers.MainChannelCreate,

	//Incoming [GET] Request URL: [/files/quest/icon/59689e1c86f7740d14064725.jpg] on [:8080]
	"/files/": services.ServeFiles,

	"/client/game/profile/items/moving": handlers.MainItemsMoving,
}

func AddMainRoute(route string, handler http.HandlerFunc) {
	_, ok := mainRouteHandlers[route]
	if ok {
		fmt.Println("URL already registered")
		return
	}

	mainRouteHandlers[route] = handler
}

func setMainRoutes(mux *http.ServeMux) {
	for route, handler := range mainRouteHandlers {
		mux.HandleFunc(route, handler)
	}
}

var tradingRouteHandlers = map[string]http.HandlerFunc{
	"/client/trading/api/traderSettings":    handlers.TradingTraderSettings,
	"/client/trading/customization/storage": handlers.TradingCustomizationStorage,
	"/files/":                               services.ServeFiles,
	"/client/trading/customization/":        handlers.TradingClothingOffers,
	"/client/trading/api/getTraderAssort/":  handlers.TradingTraderAssort,
}

func AddTradingRoute(route string, handler http.HandlerFunc) {
	_, ok := tradingRouteHandlers[route]
	if ok {
		fmt.Println("URL already registered")
		return
	}

	tradingRouteHandlers[route] = handler
}

func setTradingRoutes(mux *http.ServeMux) {
	for route, handler := range tradingRouteHandlers {
		mux.HandleFunc(route, handler)
	}
}

func setRagfairRoutes(mux *http.ServeMux) {
	/* 	 "/client/ragfair/offer/findbyid"
	   	 "/client/ragfair/itemMarketPrice"
	   	 "/client/ragfair/find" */
}

var messagingRouteHandlers = map[string]http.HandlerFunc{
	"/client/friend/list":                handlers.MessagingFriendList,
	"/client/mail/dialog/list":           handlers.MessagingDialogList,
	"/client/friend/request/list/inbox":  handlers.MessagingFriendRequestInbox,
	"/client/friend/request/list/outbox": handlers.MessagingFriendRequestOutbox,
	"/client/mail/dialog/info":           handlers.MessagingMailDialogInfo,
	"/client/mail/dialog/view":           handlers.MessagingMailDialogView,
	"/client/mail/dialog/pin":            handlers.MessagingMailDialogPin,
	"/client/mail/dialog/unpin":          handlers.MessagingMailDialogUnpin,
	"/client/mail/dialog/remove":         handlers.MessagingMailDialogRemove,
	"/client/mail/dialog/clear":          handlers.MessagingMailDialogClear,
}

func setMessagingRoutes(mux *http.ServeMux) {
	for route, handler := range messagingRouteHandlers {
		mux.HandleFunc(route, handler)
	}

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
	// "/client/mail/dialog/read"

	// "/client/mail/dialog/group/create"
	// "/client/mail/dialog/group/leave"

	// "/client/mail/dialog/group/owner/change"

	// "/client/mail/dialog/group/users/add"
	// "/client/mail/dialog/group/users/remove"

	// "/client/mail/msg/send"

}

var lobbyRouteHandlers = map[string]http.HandlerFunc{
	"/sws/":                        handlers.LobbySWS,
	"/push/notifier/get/":          handlers.LobbyPushNotifier,
	"/push/notifier/getwebsocket/": handlers.LobbyGetWebSocket,
	"/?last_id":                    handlers.LobbyNotify,
}

func setLobbyRoutes(mux *http.ServeMux) {
	for route, handler := range lobbyRouteHandlers {
		mux.HandleFunc(route, handler)
	}
}

func ServeStaticMux(mux *http.ServeMux) {

	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))
}
