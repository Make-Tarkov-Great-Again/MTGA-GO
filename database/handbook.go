package database

import (
	"MT-GO/tools"
	"fmt"

	"github.com/goccy/go-json"
)

var handbook = Handbook{}
var prices = Prices{}

// #region Handbook getters

func GetHandbook() *Handbook {
	return &handbook
}

func GetPrices() *Prices {
	return &prices
}

func GetPriceByID(id string) *int {
	price, ok := prices[id]
	if !ok {
		fmt.Println("Price of ", id, " not found, returning -1")
		return nil
	}
	return &price
}

// #endregion

// #region Handbook setters

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

// #endregion

// #region Handbook structs

type Handbook struct {
	Categories []HandbookCategory `json:"Categories"`
	Items      []HandbookItem     `json:"Items"`
}

type HandbookCategory struct {
	Id       string `json:"Id"`
	ParentId string `json:"ParentId"`
	Icon     string `json:"Icon"`
	Color    string `json:"Color"`
	Order    string `json:"Order"`
}

type HandbookItem struct {
	Id       string `json:"Id"`
	ParentId string `json:"ParentId"`
	Price    int    `json:"Price"`
}

type Prices map[string]int

// #endregion
