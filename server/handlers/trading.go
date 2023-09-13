package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"net/http"
	"strings"
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

func TradingTraderSettings(w http.ResponseWriter, r *http.Request) {
	traders := database.GetTraders()
	data := make([]map[string]interface{}, 0, len(traders))

	for _, trader := range traders {
		data = append(data, trader.Base)
	}

	body := services.ApplyResponseBody(&data)
	services.ZlibJSONReply(w, body)
}

const (
	customizationPrefix string = "/client/trading/customization/"
	customizationSuffix string = "/offers"
)

func TradingClothingOffers(w http.ResponseWriter, r *http.Request) {
	traderId := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, customizationPrefix), customizationSuffix)

	suits := database.GetTraders()[traderId].Suits
	body := services.ApplyResponseBody(suits)
	services.ZlibJSONReply(w, body)
}

const assort string = "/client/trading/api/getTraderAssort/"

func TradingTraderAssort(w http.ResponseWriter, r *http.Request) {
	trader := database.GetTraderByUID(strings.TrimPrefix(r.URL.Path, assort))
	character := database.GetCharacterByUID(services.GetSessionID(r))

	var assort = trader.GetStrippedAssort(character)
	body := services.ApplyResponseBody(assort)
	services.ZlibJSONReply(w, body)
}
