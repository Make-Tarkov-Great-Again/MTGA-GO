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
	mux.HandleFunc("/getBundleList", handlers.GetBundleList)

	mux.HandleFunc("/client/raid/person/killed/showMessage", handlers.ShowPersonKilledMessage)

	mux.HandleFunc("/client/game/start", handlers.ClientGameStart)

	mux.HandleFunc("/client/menu/locale/", handlers.ClientMenuLocale)

	mux.HandleFunc("/client/game/version/validate", handlers.ClientGameVersionValidate)

	mux.HandleFunc("/client/languages", handlers.ClientLanguages)
}
