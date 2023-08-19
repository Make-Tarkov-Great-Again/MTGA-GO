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

	mux.HandleFunc("/client/game/start", handlers.ClientGameStart)

	mux.HandleFunc("/client/menu/locale/", handlers.ClientMenuLocale)

	mux.HandleFunc("/client/game/version/validate", handlers.ClientVersionValidate)

	mux.HandleFunc("/client/languages", handlers.ClientLanguages)

	mux.HandleFunc("/client/game/config", handlers.ClientGameConfig)

	mux.HandleFunc("/client/items", handlers.ClientItems)

	mux.HandleFunc("/client/customization", handlers.ClientCustomization)

	mux.HandleFunc("/client/globals", handlers.ClientGlobals)

	mux.HandleFunc("/client/trading/api/traderSettings", handlers.ClientTraderSettings)

	mux.HandleFunc("/client/settings", handlers.ClientSettings)

	mux.HandleFunc("/client/game/profile/list", handlers.ClientProfileList)

	mux.HandleFunc("/client/account/customization", handlers.ClientAccountCustomization)

	mux.HandleFunc("/client/locale/", handlers.ClientLocale)

	mux.HandleFunc("/client/game/keepalive", handlers.KeepAlive)

	mux.HandleFunc("/client/game/profile/nickname/reserved", handlers.NicknameReserved)

	mux.HandleFunc("/client/game/profile/nickname/validate", handlers.NicknameValidate)

	mux.HandleFunc("/client/game/profile/create", handlers.ProfileCreate)
}

func setTradingRoutes(mux *http.ServeMux) {
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

func setMessagingRoutes(mux *http.ServeMux) {
	// "/client/friend/list"

	// "/client/friend/request/list/inbox"
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
	// "/client/mail/dialog/list"
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
