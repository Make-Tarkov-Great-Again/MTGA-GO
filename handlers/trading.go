package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"

	"MT-GO/pkg"
)

func TradingCustomizationStorage(w http.ResponseWriter, r *http.Request) {
	suitesStorage, err := pkg.GetSuitesStorage(pkg.GetSessionID(r))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(suitesStorage)
	pkg.SendZlibJSONReply(w, body)
}

func TradingTraderSettings(w http.ResponseWriter, _ *http.Request) {
	traderSettings := pkg.GetTraderSettings()
	body := pkg.ApplyResponseBody(traderSettings)
	pkg.SendZlibJSONReply(w, body)
}

func TradingClothingOffers(w http.ResponseWriter, r *http.Request) {
	suits, err := pkg.GetTraderSuits(chi.URLParam(r, "id"))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(suits)
	pkg.SendZlibJSONReply(w, body)
}

func TradingTraderAssort(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	traderAssort, err := pkg.GetTraderAssort(r)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(traderAssort)
	pkg.SendZlibJSONReply(w, body)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Response Time: %v\n", elapsedTime)
}
