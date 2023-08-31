package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"fmt"

	"github.com/goccy/go-json"
)

var handbook = structs.Handbook{}
var prices = structs.Prices{}

func GetHandbook() *structs.Handbook {
	return &handbook
}

func GetPrices() *structs.Prices {
	return &prices
}

func GetPriceByID(id string) int {
	price, ok := prices[id]
	if !ok {
		fmt.Println("Price of ", id, " not found, returning -1")
		return -1
	}
	return price
}

func setHandbook() {
	raw := tools.GetJSONRawMessage(handbookPath)
	err := json.Unmarshal(raw, &handbook)
	if err != nil {
		panic(err)
	}

	for _, v := range handbook.Items {
		prices[v.Id] = v.Price
	}
	fmt.Println()
}
