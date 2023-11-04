package database

import (
	"fmt"
	"log"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var handbook = Handbook{}
var handbookIndex = map[string]int16{}
var prices = make(map[string]*int32)

// #region Handbook getters

func GetHandbook() *Handbook {
	return &handbook
}

// GetPrices Get prices of all items
func GetPrices() map[string]*int32 {
	return prices
}

const priceNotFound string = "Price of %s not found"

// GetPriceByID Get item price by ID
func GetPriceByID(id string) (*int32, error) {
	price, ok := prices[id]
	if !ok {

		return nil, fmt.Errorf(priceNotFound, id)
	}
	return price, nil
}

// #endregion

// #region Handbook setters

func setHandbook() {
	raw := tools.GetJSONRawMessage(handbookPath)
	err := json.Unmarshal(raw, &handbook)
	if err != nil {
		log.Fatalln(err)
	}

	for idx, v := range handbook.Items {
		handbookIndex[v.Id] = int16(idx)
		prices[v.Id] = &v.Price
	}
}

const handbookItemEntryNotExist string = "Handbook Item for %s entry doesn't exist"

func (i *DatabaseItem) GetHandbookItemEntry() (*HandbookItem, error) {
	idx, ok := handbookIndex[i.ID]
	if !ok {
		return nil, fmt.Errorf(handbookItemEntryNotExist, i.ID)
	}
	return &handbook.Items[idx], nil
}

const couldNotCreateClone string = "Could not create clone of entry, %s"

func (i *DatabaseItem) CloneHandbookItemEntry() (*HandbookItem, error) {
	handbookEntry, err := i.GetHandbookItemEntry()
	if err != nil {
		return nil, fmt.Errorf(couldNotCreateClone, err)
	}
	return &HandbookItem{Id: "", ParentId: handbookEntry.ParentId, Price: 0}, nil
}

func (hbi *HandbookItem) SetHandbookItemEntry() {
	handbook.Items = append(handbook.Items, *hbi)
	handbookIndex[hbi.Id] = int16(len(handbook.Items) - 1)
}

func SetHandbookItemEntry(entry HandbookItem) {
	handbook.Items = append(handbook.Items, entry)
	handbookIndex[entry.Id] = int16(len(handbook.Items) - 1)
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
