// Package routes houses all routes to the client and server
package server

import (
	"MT-GO/handlers"
	"MT-GO/pkg"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var mainRouteHandlers = map[string]http.HandlerFunc{
	"/sp/config/bots/difficulty":  handlers.GetBotDifficulty,
	"/raid/profile/save":          handlers.RaidProfileSave,
	"/sp/airdrop/config":          handlers.AirdropConfig,
	"/files/{main}/{type}/{file}": pkg.ServeFiles,

	"/client/game/start":                          handlers.MainGameStart,
	"/client/menu/locale/{id}":                    handlers.MainMenuLocale,
	"/client/game/version/validate":               handlers.MainVersionValidate,
	"/client/languages":                           handlers.MainLanguages,
	"/client/game/config":                         handlers.MainGameConfig,
	"/client/items":                               handlers.MainItems,
	"/client/customization":                       handlers.MainCustomization,
	"/client/globals":                             handlers.MainGlobals,
	"/client/settings":                            handlers.MainSettings,
	"/client/game/profile/list":                   handlers.MainProfileList,
	"/client/account/customization":               handlers.MainAccountCustomization,
	"/client/locale/{id}":                         handlers.MainLocale,
	"/client/game/keepalive":                      handlers.MainKeepAlive,
	"/client/game/profile/nickname/reserved":      handlers.MainNicknameReserved,
	"/client/game/profile/nickname/validate":      handlers.MainNicknameValidate,
	"/client/game/profile/create":                 handlers.MainProfileCreate,
	"/client/game/profile/select":                 handlers.MainProfileSelect,
	"/client/game/profile/voice":                  handlers.ChangeVoice,
	"/client/profile/status":                      handlers.MainProfileStatus,
	"/client/profile/settings":                    handlers.MainProfileSettings,
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
	"/client/server/list":                         handlers.GetServerList,
	"/client/checkVersion":                        handlers.MainCheckVersion,
	"/client/game/logout":                         handlers.MainLogout,
	"/client/items/prices/{id}":                   handlers.MainPrices,
	"/client/notifier/channel/create":             handlers.MainChannelCreate,
	"/client/game/profile/items/moving":           handlers.MainItemsMoving,
	"/client/match/offline/end":                   handlers.OfflineMatchEnd,
	"/client/match/group/exit_from_menu":          handlers.ExitFromMenu,
	"/client/match/group/invite/cancel-all":       handlers.InviteCancelAll,
	"/client/match/available":                     handlers.MatchAvailable,
	"/client/match/raid/not-ready":                handlers.RaidNotReady,
	"/client/match/raid/ready":                    handlers.RaidReady,
	"/client/match/group/status":                  handlers.GroupStatus,
	"/client/match/group/looking/start":           handlers.LookingForGroupStart,
	"/client/match/group/looking/stop":            handlers.LookingForGroupStop,
	"/client/match/updatePing":                    handlers.MatchUpdatePing,
	"/client/raid/configuration":                  handlers.RaidConfiguration,
	"/client/location/getLocalloot":               handlers.GetLocalLoot,
	"/client/insurance/items/cost":                handlers.InsuranceItemsCost,
	"/client/insurance/items/list/cost":           handlers.InsuranceListCost,
	"/client/game/bot/generate":                   handlers.BotGenerate,

	//.14
	"/client/achievement/statistic": handlers.GetAchievementStats,
	"/client/achievement/list":      handlers.GetAchievements,
	"/client/builds/list":           handlers.MainBuildsList,
	//"/client/handbook/builds/my/list" check if still used
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
	"/client/trading/api/traderSettings":        handlers.TradingTraderSettings,
	"/client/trading/customization/storage":     handlers.TradingCustomizationStorage,
	"/files/{main}/{type}/{file}":               pkg.ServeFiles,
	"/client/trading/customization/{id}/offers": handlers.TradingClothingOffers,
	"/client/trading/api/getTraderAssort/{id}":  handlers.TradingTraderAssort,
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
	"/client/ragfair/find": handlers.RagfairFind,
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
	"/push/notifier/get/{id}":          handlers.LobbyPushNotifier,
	"/push/notifier/getwebsocket/{id}": handlers.LobbyGetWebSocket,
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
