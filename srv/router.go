// Package routes houses all routes to the client and srv
package srv

import (
	"MT-GO/pkg"
	"MT-GO/srv/hndlr"
	"log"
	"net/http"
)

var mainRouteHandlers = map[string]http.HandlerFunc{
	"/getBrandName":              hndlr.GetBrandName,
	"/sp/config/bots/difficulty": hndlr.GetBotDifficulty,
	"/getBundleList":             hndlr.GetBundleList,

	"/client/raid/person/killed/showMessage":      hndlr.ShowPersonKilledMessage,
	"/client/game/start":                          hndlr.MainGameStart,
	"/client/putMetrics":                          hndlr.MainPutMetrics,
	"/client/menu/locale/":                        hndlr.MainMenuLocale,
	"/client/game/version/validate":               hndlr.MainVersionValidate,
	"/client/languages":                           hndlr.MainLanguages,
	"/client/game/config":                         hndlr.MainGameConfig,
	"/client/items":                               hndlr.MainItems,
	"/client/customization":                       hndlr.MainCustomization,
	"/client/globals":                             hndlr.MainGlobals,
	"/client/settings":                            hndlr.MainSettings,
	"/client/game/profile/list":                   hndlr.MainProfileList,
	"/client/account/customization":               hndlr.MainAccountCustomization,
	"/client/locale/":                             hndlr.MainLocale,
	"/client/game/keepalive":                      hndlr.MainKeepAlive,
	"/client/game/profile/nickname/reserved":      hndlr.MainNicknameReserved,
	"/client/game/profile/nickname/validate":      hndlr.MainNicknameValidate,
	"/client/game/profile/create":                 hndlr.MainProfileCreate,
	"/client/game/profile/select":                 hndlr.MainProfileSelect,
	"/client/profile/status":                      hndlr.MainProfileStatus,
	"/client/weather":                             hndlr.MainWeather,
	"/client/locations":                           hndlr.MainLocations,
	"/client/handbook/templates":                  hndlr.MainTemplates,
	"/client/hideout/areas":                       hndlr.MainHideoutAreas,
	"/client/hideout/qte/list":                    hndlr.MainHideoutQTE,
	"/client/hideout/settings":                    hndlr.MainHideoutSettings,
	"/client/hideout/production/recipes":          hndlr.MainHideoutRecipes,
	"/client/hideout/production/scavcase/recipes": hndlr.MainHideoutScavRecipes,
	"/client/handbook/builds/my/list":             hndlr.MainBuildsList,
	"/client/quest/list":                          hndlr.MainQuestList,
	"/client/match/group/current":                 hndlr.MainCurrentGroup,
	"/client/repeatalbeQuests/activityPeriods":    hndlr.MainRepeatableQuests,
	"/client/server/list":                         hndlr.GetServerList,
	"/client/checkVersion":                        hndlr.MainCheckVersion,
	"/client/game/logout":                         hndlr.MainLogout,
	"/client/items/prices/":                       hndlr.MainPrices,
	"/client/notifier/channel/create":             hndlr.MainChannelCreate,

	"/files/": pkg.ServeFiles,

	"/client/game/profile/items/moving": hndlr.MainItemsMoving,

	"/client/match/offline/end": hndlr.OfflineMatchEnd,
	//"/client/match/available": hndlr.MatchAvailable,
	//"/client/match/updatePing": hndlr.MatchUpdatePing,
	//"/client/match/exit": hndlr.MatchExit,
	//"/client/match/join": hndlr.MatchJoin,
	"/client/match/group/exit_from_menu":    hndlr.ExitFromMenu,
	"/client/match/group/invite/cancel-all": hndlr.InviteCancelAll,
	"/client/match/available":               hndlr.MatchAvailable,
	"/client/match/raid/not-ready":          hndlr.RaidNotReady,
	"/client/match/raid/ready":              hndlr.RaidReady,
	"/client/match/group/status":            hndlr.GroupStatus,
	"/client/match/group/looking/start":     hndlr.LookingForGroupStart,
	"/client/match/group/looking/stop":      hndlr.LookingForGroupStop,
	//"/client/match/group/invite/send": hndlr.GroupInviteSend,
	//"/client/match/group/invite/accept": hndlr.GroupInviteAccept,
	//"/client/match/group/invite/cancel": hndlr.GroupInviteCancel,
	//"/client/match/group/transfer": hndlr.GroupTransfer,
	//"/client/match/group/leave": hndlr.GroupLeave,
	//"/client/match/group/delete": hndlr.GroupDelete,
	//"/client/match/group/create": hndlr.GroupCreate,
	//"/client/match/group/player/remove": hndlr.GroupPlayerRemove,
	//"/client/match/group/start_game": hndlr.GroupStartGame,

	//"/client/raid/createFriendlyAI": hndlr.CreateFriendlyAI,
	//"/client/raid/person/killed": hndlr.PersonKilled,
	"/client/raid/configuration": hndlr.RaidConfiguration,

	"/client/location/getLocalloot":     hndlr.GetLocalLoot,
	"/client/insurance/items/list/cost": hndlr.InsuranceListCost,
	"/client/game/bot/generate":         hndlr.BotGenerate,

	"/raid/profile/save": hndlr.RaidProfileSave,
	"/sp/airdrop/config": hndlr.AirdropConfig,
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
	"/client/trading/api/traderSettings":    hndlr.TradingTraderSettings,
	"/client/trading/customization/storage": hndlr.TradingCustomizationStorage,
	"/files/":                               pkg.ServeFiles,
	"/client/trading/customization/":        hndlr.TradingClothingOffers,
	"/client/trading/api/getTraderAssort/":  hndlr.TradingTraderAssort,
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
	"/client/ragfair/find": hndlr.RagfairFind,
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
	"/client/friend/list":                hndlr.MessagingFriendList,
	"/client/mail/dialog/list":           hndlr.MessagingDialogList,
	"/client/friend/request/list/inbox":  hndlr.MessagingFriendRequestInbox,
	"/client/friend/request/list/outbox": hndlr.MessagingFriendRequestOutbox,
	"/client/mail/dialog/info":           hndlr.MessagingMailDialogInfo,
	"/client/mail/dialog/view":           hndlr.MessagingMailDialogView,
	"/client/mail/dialog/pin":            hndlr.MessagingMailDialogPin,
	"/client/mail/dialog/unpin":          hndlr.MessagingMailDialogUnpin,
	"/client/mail/dialog/remove":         hndlr.MessagingMailDialogRemove,
	"/client/mail/dialog/clear":          hndlr.MessagingMailDialogClear,
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
	"/push/notifier/get/":          hndlr.LobbyPushNotifier,
	"/push/notifier/getwebsocket/": hndlr.LobbyGetWebSocket,
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
