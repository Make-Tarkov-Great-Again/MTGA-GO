package handlers

import (
	"net/http"

	"MT-GO/database"
	"MT-GO/services"
)

func TradingCustomizationStorage(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	suites := database.GetStorageByUID(sessionID).Suites

	storage := map[string]interface{}{
		"_id":    sessionID,
		"suites": suites,
	}

	body := services.ApplyResponseBody(storage)
	services.ZlibJSONReply(w, body)
}

func TradingTraderSettings(w http.ResponseWriter, _ *http.Request) {
	traders := database.GetTraders()
	data := make([]*database.TraderBase, 0, len(traders))

	for _, trader := range traders {
		data = append(data, trader.Base)
	}

	body := services.ApplyResponseBody(&data)
	services.ZlibJSONReply(w, body)
}

func TradingClothingOffers(w http.ResponseWriter, r *http.Request) {
	traderId := r.URL.Path[30:54] //30:54

	suits := database.GetTraders()[traderId].Suits
	body := services.ApplyResponseBody(suits)
	services.ZlibJSONReply(w, body)
}

func TradingTraderAssort(w http.ResponseWriter, r *http.Request) {
	trader := database.GetTraderByUID(r.URL.Path[36:])
	character := database.GetCharacterByUID(services.GetSessionID(r))

	var assort = trader.GetStrippedAssort(character)
	body := services.ApplyResponseBody(assort)
	services.ZlibJSONReply(w, body)
}
