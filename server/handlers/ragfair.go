package handlers

import (
	"MT-GO/services"
	"fmt"
	"net/http"
)

type ragfairOffers struct {
	Offers           []any  `json:"offers"`
	OffersCount      int16  `json:"offersCount"`
	SelectedCategory string `json:"selectedCategory"`
}

func RagfairFind(w http.ResponseWriter, r *http.Request) {
	output := ragfairOffers{
		Offers:           make([]any, 0),
		OffersCount:      0,
		SelectedCategory: "",
	}

	log.Println(routeNotImplemented)
	body := services.ApplyResponseBody(output)
	services.ZlibJSONReply(w, r.RequestURI, body)
}
