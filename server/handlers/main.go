package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"MT-GO/structs"
	"MT-GO/tools"
	"fmt"
	"net/http"
	"strings"
)

const ROUTE_NOT_IMPLEMENTED = "Route is not implemented yet, using empty values instead"

// GetBundleList returns a list of custom bundles to the client
func GetBundleList(w http.ResponseWriter, _ *http.Request) {
	fmt.Println(ROUTE_NOT_IMPLEMENTED)
	services.ZlibJSONReply(w, []string{})
}

func ShowPersonKilledMessage(w http.ResponseWriter, _ *http.Request) {
	services.ZlibJSONReply(w, "true")
}

func ClientGameStart(w http.ResponseWriter, _ *http.Request) {
	data := map[string]interface{}{
		"utc_time": tools.GetCurrentTimeInSeconds(),
	}

	body := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, body)
}

func ClientMenuLocale(w http.ResponseWriter, r *http.Request) {
	locale := strings.TrimPrefix(r.URL.Path, "/client/menu/locale/")

	body := struct {
		Data   *structs.LocaleMenu
		Err    int
		Errmsg interface{}
	}{
		Data: database.GetLocalesMenuByName(locale),
	}

	services.ZlibJSONReply(w, body)
}

func ClientGameVersionValidate(w http.ResponseWriter, _ *http.Request) {
	body := struct {
		Data   interface{}
		Err    int
		Errmsg interface{}
	}{
		Err: 0,
	}

	services.ZlibJSONReply(w, body)
}

func ClientLanguages(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Data   map[string]string
		Err    int
		Errmsg interface{}
	}{
		Data: database.GetLanguages(),
	}

	services.ZlibJSONReply(w, body)
}
