package webapi

import (
	"net/http"
)

func WebRouteHandlers() map[string]http.HandlerFunc {
	webRouteHandlers := map[string]http.HandlerFunc{

		"/":         handleHome(),
		"/login":    handleLogin(),
		"/register": handleRegister(),

		"/webapi/account/token": handleTokenCreate(),

		// logged only
		"/settings": handleSettings(),

		// logged only
		"/webapi/account/detail": handleAccountDetail(),
		"/webapi/account/save":   handleAccountSave(),
	}
	return webRouteHandlers
}
