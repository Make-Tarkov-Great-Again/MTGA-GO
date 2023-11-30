// Package routes houses all routes to the client and server
package server

import (
	handler2 "MT-GO/handlers"
	"MT-GO/pkg"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var mainRouteHandlers = map[string]http.HandlerFunc{
	"/getBrandName":              handler2.GetBrandName,
	"/sp/config/bots/difficulty": handler2.GetBotDifficulty,
	"/getBundleList":             handler2.GetBundleList,
	"/raid/profile/save":         handler2.RaidProfileSave,
	"/sp/airdrop/config":         handler2.AirdropConfig,
	"/files/{id}":                pkg.ServeFiles,

	"/client/game/start":                          handler2.MainGameStart,
	"/client/menu/locale/{id}":                    handler2.MainMenuLocale,
	"/client/game/version/validate":               handler2.MainVersionValidate,
	"/client/languages":                           handler2.MainLanguages,
	"/client/game/config":                         handler2.MainGameConfig,
	"/client/items":                               handler2.MainItems,
	"/client/customization":                       handler2.MainCustomization,
	"/client/globals":                             handler2.MainGlobals,
	"/client/settings":                            handler2.MainSettings,
	"/client/game/profile/list":                   handler2.MainProfileList,
	"/client/account/customization":               handler2.MainAccountCustomization,
	"/client/locale/{id}":                         handler2.MainLocale,
	"/client/game/keepalive":                      handler2.MainKeepAlive,
	"/client/game/profile/nickname/reserved":      handler2.MainNicknameReserved,
	"/client/game/profile/nickname/validate":      handler2.MainNicknameValidate,
	"/client/game/profile/create":                 handler2.MainProfileCreate,
	"/client/game/profile/select":                 handler2.MainProfileSelect,
	"/client/game/profile/voice":                  handler2.ChangeVoice,
	"/client/profile/status":                      handler2.MainProfileStatus,
	"/client/profile/settings":                    handler2.MainProfileSettings,
	"/client/weather":                             handler2.MainWeather,
	"/client/locations":                           handler2.MainLocations,
	"/client/handbook/templates":                  handler2.MainTemplates,
	"/client/hideout/areas":                       handler2.MainHideoutAreas,
	"/client/hideout/qte/list":                    handler2.MainHideoutQTE,
	"/client/hideout/settings":                    handler2.MainHideoutSettings,
	"/client/hideout/production/recipes":          handler2.MainHideoutRecipes,
	"/client/hideout/production/scavcase/recipes": handler2.MainHideoutScavRecipes,
	"/client/handbook/builds/my/list":             handler2.MainBuildsList,
	"/client/quest/list":                          handler2.MainQuestList,
	"/client/match/group/current":                 handler2.MainCurrentGroup,
	"/client/repeatalbeQuests/activityPeriods":    handler2.MainRepeatableQuests,
	"/client/server/list":                         handler2.GetServerList,
	"/client/checkVersion":                        handler2.MainCheckVersion,
	"/client/game/logout":                         handler2.MainLogout,
	"/client/items/prices/{id}":                   handler2.MainPrices,
	"/client/notifier/channel/create":             handler2.MainChannelCreate,
	"/client/game/profile/items/moving":           handler2.MainItemsMoving,
	"/client/match/offline/end":                   handler2.OfflineMatchEnd,
	"/client/match/group/exit_from_menu":          handler2.ExitFromMenu,
	"/client/match/group/invite/cancel-all":       handler2.InviteCancelAll,
	"/client/match/available":                     handler2.MatchAvailable,
	"/client/match/raid/not-ready":                handler2.RaidNotReady,
	"/client/match/raid/ready":                    handler2.RaidReady,
	"/client/match/group/status":                  handler2.GroupStatus,
	"/client/match/group/looking/start":           handler2.LookingForGroupStart,
	"/client/match/group/looking/stop":            handler2.LookingForGroupStop,
	"/client/match/updatePing":                    handler2.MatchUpdatePing,
	"/client/raid/configuration":                  handler2.RaidConfiguration,
	"/client/location/getLocalloot":               handler2.GetLocalLoot,
	"/client/insurance/items/cost":                handler2.InsuranceItemsCost,
	"/client/insurance/items/list/cost":           handler2.InsuranceListCost,
	"/client/game/bot/generate":                   handler2.BotGenerate,
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

func loadMainRoutes(mux *chi.Mux) {
	for route, handler := range mainRouteHandlers {
		mux.HandleFunc(route, handler)
	}
}

var tradingRouteHandlers = map[string]http.HandlerFunc{
	"/client/trading/api/traderSettings":       handler2.TradingTraderSettings,
	"/client/trading/customization/storage":    handler2.TradingCustomizationStorage,
	"/files/{file}":                            pkg.ServeFiles,
	"/client/trading/customization/{id}":       handler2.TradingClothingOffers,
	"/client/trading/api/getTraderAssort/{id}": handler2.TradingTraderAssort,
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

func loadTradingRoutes(mux *chi.Mux) {
	for route, handler := range tradingRouteHandlers {
		mux.HandleFunc(route, handler)
	}
}

var ragfairRouteHandlers = map[string]http.HandlerFunc{
	//"/client/ragfair/offer/findbyid"
	//"/client/ragfair/itemMarketPrice"
	"/client/ragfair/find": handler2.RagfairFind,
}

func loadRagfairRoutes(mux *chi.Mux) {
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
	"/client/friend/list":                handler2.MessagingFriendList,
	"/client/mail/dialog/list":           handler2.MessagingDialogList,
	"/client/friend/request/list/inbox":  handler2.MessagingFriendRequestInbox,
	"/client/friend/request/list/outbox": handler2.MessagingFriendRequestOutbox,
	"/client/mail/dialog/info":           handler2.MessagingMailDialogInfo,
	"/client/mail/dialog/view":           handler2.MessagingMailDialogView,
	"/client/mail/dialog/pin":            handler2.MessagingMailDialogPin,
	"/client/mail/dialog/unpin":          handler2.MessagingMailDialogUnpin,
	"/client/mail/dialog/remove":         handler2.MessagingMailDialogRemove,
	"/client/mail/dialog/clear":          handler2.MessagingMailDialogClear,
}

func loadMessagingRoutes(mux *chi.Mux) {
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
	"/push/notifier/get/{id}":          handler2.LobbyPushNotifier,
	"/push/notifier/getwebsocket/{id}": handler2.LobbyGetWebSocket,
}

func loadLobbyRoutes(mux *chi.Mux) {
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
