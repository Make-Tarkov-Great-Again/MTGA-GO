package database

import (
	"fmt"
	"log"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var handbook = Handbook{}
var prices = make(map[string]int32)

// #region Handbook getters

func GetHandbook() *Handbook {
	return &handbook
}

// GetPrices Get prices of all items
func GetPrices() *map[string]int32 {
	return &prices
}

// GetPriceByID Get item price by ID
func GetPriceByID(id string) *int32 {
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
		log.Fatalln(err)
	}

	for _, v := range handbook.Items {
		prices[v.Id] = v.Price
	}
}

func SetHandbookItemEntry(entry HandbookItem) {
	handbook.Items = append(handbook.Items, entry)
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
	Price    int32  `json:"Price"`
}

type Prices map[string]int32

// #endregion
