// Package routes houses all routes to the client and server
package server

import (
	"MT-GO/server/handlers"
	"MT-GO/services"
	"log"
	"net/http"
)

var mainRouteHandlers = map[string]http.HandlerFunc{
	"/getBrandName":              handlers.GetBrandName,
	"/sp/config/bots/difficulty": handlers.GetBotDifficulty,
	"/getBundleList":             handlers.GetBundleList,

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
	"/client/game/logout":                         handlers.MainLogout,
	"/client/items/prices/":                       handlers.MainPrices,
	"/client/notifier/channel/create":             handlers.MainChannelCreate,

	"/files/": services.ServeFiles,

	"/client/game/profile/items/moving": handlers.MainItemsMoving,

	"/client/match/offline/end": handlers.OfflineMatchEnd,
	//"/client/match/available": handlers.MatchAvailable,
	//"/client/match/updatePing": handlers.MatchUpdatePing,
	//"/client/match/exit": handlers.MatchExit,
	//"/client/match/join": handlers.MatchJoin,
	"/client/match/group/exit_from_menu":    handlers.ExitFromMenu,
	"/client/match/group/invite/cancel-all": handlers.InviteCancelAll,
	"/client/match/available":               handlers.MatchAvailable,
	"/client/match/raid/not-ready":          handlers.RaidNotReady,
	"/client/match/raid/ready":              handlers.RaidReady,
	"/client/match/group/status":            handlers.GroupStatus,
	"/client/match/group/looking/start":     handlers.LookingForGroupStart,
	"/client/match/group/looking/stop":      handlers.LookingForGroupStop,
	//"/client/match/group/invite/send": handlers.GroupInviteSend,
	//"/client/match/group/invite/accept": handlers.GroupInviteAccept,
	//"/client/match/group/invite/cancel": handlers.GroupInviteCancel,
	//"/client/match/group/transfer": handlers.GroupTransfer,
	//"/client/match/group/leave": handlers.GroupLeave,
	//"/client/match/group/delete": handlers.GroupDelete,
	//"/client/match/group/create": handlers.GroupCreate,
	//"/client/match/group/player/remove": handlers.GroupPlayerRemove,
	//"/client/match/group/start_game": handlers.GroupStartGame,

	//"/client/raid/createFriendlyAI": handlers.CreateFriendlyAI,
	//"/client/raid/person/killed": handlers.PersonKilled,
	"/client/raid/configuration": handlers.RaidConfiguration,

	"/client/location/getLocalloot":     handlers.GetLocalLoot,
	"/client/insurance/items/list/cost": handlers.InsuranceListCost,

	"/client/game/bot/generate": handlers.BotGenerate,

	"/raid/profile/save": handlers.RaidProfileSave,
	"/sp/airdrop/config": handlers.AirdropConfig,
}

func AddMainRoute(route string, handler http.HandlerFunc) {
	_, ok := mainRouteHandlers[route]
	if ok {
		log.Println("URL already registered")
		return
	}

	mainRouteHandlers[route] = handler
}

func OverrideMainRoute(route string, handler http.HandlerFunc) {
	if _, ok := mainRouteHandlers[route]; !ok {
		log.Println("URL doesn't exist")
		return
	}

	log.Println("URL override for", route, "registered")
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
		log.Println("URL already registered")
		return
	}

	tradingRouteHandlers[route] = handler
}

func OverrideTradingRoute(route string, handler http.HandlerFunc) {
	if _, ok := tradingRouteHandlers[route]; !ok {
		log.Println("URL doesn't exist")
		return
	}

	log.Println("URL override for", route, "registered")
	tradingRouteHandlers[route] = handler
}

func setTradingRoutes(mux *http.ServeMux) {
	for route, handler := range tradingRouteHandlers {
		mux.HandleFunc(route, handler)
	}
}

var ragfairRouteHandlers = map[string]http.HandlerFunc{
	//"/client/ragfair/offer/findbyid"
	//"/client/ragfair/itemMarketPrice"
	"/client/ragfair/find": handlers.RagfairFind,
}

func setRagfairRoutes(mux *http.ServeMux) {
	for route, handler := range ragfairRouteHandlers {
		mux.HandleFunc(route, handler)
	}
}

func OverrideRagfairRoute(route string, handler http.HandlerFunc) {
	if _, ok := ragfairRouteHandlers[route]; !ok {
		log.Println("URL doesn't exist")
		return
	}

	log.Println("URL override for", route, "registered")
	ragfairRouteHandlers[route] = handler
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

func OverrideMessagingRoute(route string, handler http.HandlerFunc) {
	if _, ok := messagingRouteHandlers[route]; !ok {
		log.Println("URL doesn't exist")
		return
	}

	log.Println("URL override for", route, "registered")
	messagingRouteHandlers[route] = handler
}

var lobbyRouteHandlers = map[string]http.HandlerFunc{
	"/push/notifier/get/":          handlers.LobbyPushNotifier,
	"/push/notifier/getwebsocket/": handlers.LobbyGetWebSocket,
}

func setLobbyRoutes(mux *http.ServeMux) {
	for route, handler := range lobbyRouteHandlers {
		mux.HandleFunc(route, handler)
	}
}

func OverrideLobbyRoute(route string, handler http.HandlerFunc) {
	if _, ok := lobbyRouteHandlers[route]; !ok {
		log.Println("URL doesn't exist")
		return
	}

	log.Println("URL override for", route, "registered")
	lobbyRouteHandlers[route] = handler
}
