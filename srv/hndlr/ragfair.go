package hndlr

import (
	"MT-GO/pkg"
	"log"
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
	body := pkg.ApplyResponseBody(output)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}
