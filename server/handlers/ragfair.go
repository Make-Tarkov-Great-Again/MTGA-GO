package handlers

import (
	"MT-GO/services"
	"fmt"
	"net/http"
)

type ragfairOffers struct {
	Offers           []interface{} `json:"offers"`
	OffersCount      int16         `json:"offersCount"`
	SelectedCategory string        `json:"selectedCategory"`
}

func RagfairFind(w http.ResponseWriter, r *http.Request) {
	output := ragfairOffers{
		Offers:           make([]interface{}, 0),
		OffersCount:      0,
		SelectedCategory: "",
	}

	fmt.Println(routeNotImplemented)
	body := services.ApplyResponseBody(output)
	services.ZlibJSONReply(w, body)
}
