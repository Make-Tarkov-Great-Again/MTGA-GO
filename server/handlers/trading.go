package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"net/http"
)

func TradingCustomizationStorage(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)

	suites := database.GetProfileByUID(sessionID).Storage.Suites

	storage := map[string]interface{}{
		"_id":    sessionID,
		"suites": suites,
	}

	body := services.ApplyResponseBody(storage)
	services.ZlibJSONReply(w, body)
}

const traderSettingsRoute string = "/client/trading/api/traderSettings"

func TradingTraderSettings(w http.ResponseWriter, r *http.Request) {
	traders := database.GetTraders()
	data := make([]map[string]interface{}, 0, len(traders))

	for _, trader := range traders {
		data = append(data, trader.Base)
	}

	body := services.ApplyResponseBody(&data)
	services.ZlibJSONReply(w, body)
}
