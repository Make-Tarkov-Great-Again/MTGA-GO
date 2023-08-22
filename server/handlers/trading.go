package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"fmt"
	"net/http"
	"strings"
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

const prod string = "https://prod.escapefromtarkov.com/%s"

func TradingFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	icon := strings.Split(r.RequestURI, "/")
	fmt.Println(icon)

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
	traderId := strings.TrimPrefix(r.URL.Path, assort)

	assort := database.GetTraders()[traderId].Assort
	body := services.ApplyResponseBody(assort)
	services.ZlibJSONReply(w, body)
	fmt.Println("You need to add proper transactions")
}
