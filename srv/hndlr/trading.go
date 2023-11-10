package hndlr

import (
	"log"
	"net/http"

	"MT-GO/pkg"
)

func TradingCustomizationStorage(w http.ResponseWriter, r *http.Request) {
	suitesStorage, err := pkg.GetSuitesStorage(pkg.GetSessionID(r))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(suitesStorage)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func TradingTraderSettings(w http.ResponseWriter, r *http.Request) {
	traderSettings := pkg.GetTraderSettings()
	body := pkg.ApplyResponseBody(traderSettings)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func TradingClothingOffers(w http.ResponseWriter, r *http.Request) {
	suits, err := pkg.GetTraderSuits(r.URL.Path[30:54])
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(suits)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func TradingTraderAssort(w http.ResponseWriter, r *http.Request) {
	traderAssort, err := pkg.GetTraderAssort(r)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(traderAssort)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}
