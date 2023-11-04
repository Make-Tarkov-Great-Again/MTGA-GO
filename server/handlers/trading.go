package handlers

import (
	"log"
	"net/http"

	"MT-GO/database"
	"MT-GO/services"
)

func TradingCustomizationStorage(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	storage, err := database.GetStorageByUID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

	suitesStorage := map[string]any{
		"_id":    sessionID,
		"suites": storage.Suites,
	}

	body := services.ApplyResponseBody(suitesStorage)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func TradingTraderSettings(w http.ResponseWriter, r *http.Request) {
	traders := database.GetTraders()
	data := make([]*database.TraderBase, 0, len(traders))

	for _, trader := range traders {
		data = append(data, trader.Base)
	}

	body := services.ApplyResponseBody(&data)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func TradingClothingOffers(w http.ResponseWriter, r *http.Request) {
	traderId := r.URL.Path[30:54] //30:54

	suits := database.GetTraders()[traderId].Suits
	body := services.ApplyResponseBody(suits)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func TradingTraderAssort(w http.ResponseWriter, r *http.Request) {
	tid := r.URL.Path[36:]
	trader, err := database.GetTraderByUID(tid)
	if err != nil {
		log.Fatalln(err)
	}

	character := database.GetCharacterByUID(services.GetSessionID(r))
	var assort = trader.GetStrippedAssort(character)

	body := services.ApplyResponseBody(assort)
	services.ZlibJSONReply(w, r.RequestURI, body)
}
