// Package routes houses all routes to the client and server
package server

import (
	"MT-GO/server/handlers"
	"net/http"
)

// SetRoutes sets all existing routes
func setRoutes(main *http.ServeMux /* , main *http.ServeMux */) {

	setMainRoutes(main)
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
}
