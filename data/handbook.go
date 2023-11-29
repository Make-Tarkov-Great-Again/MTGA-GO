package data

import (
	"fmt"
	"log"
	"math"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

type Template struct {
	index    *TemplateIndex
	handbook *Templates
	prices   map[string]int32
}

type TemplateIndex struct {
	Item       map[string]int16
	Categories map[string][]string
}

// #region Handbook getters

func GetHandbook() *Templates {
	return db.template.handbook
}

// GetPrices Get prices of all items
func GetPrices() map[string]int32 {
	return db.template.prices
}

const priceNotFound string = "price of %s not found"

// GetPriceByID Get item price by ID
func GetPriceByID(id string) (int32, error) {
	price, ok := db.template.prices[id]
	if !ok {
		return -1, fmt.Errorf(priceNotFound, id)
	}
	return price, nil
}

// #endregion

// #region Handbook setters
var handbookCategories = make(map[string][]string)

func setHandbook() {
	db.template = &Template{
		handbook: new(Templates),
	}
	raw := tools.GetJSONRawMessage(handbookPath)
	if err := json.UnmarshalNoEscape(raw, &db.template.handbook); err != nil {
		log.Fatalln(err)
	}
}

func setHandbookIndex() {
	db.template.index = &TemplateIndex{
		Item:       make(map[string]int16),
		Categories: make(map[string][]string),
	}

	for _, v := range db.template.handbook.Categories {
		if _, ok := handbookCategories[v.ParentID]; !ok {
			if v.ParentID == "" {
				if _, ok := handbookCategories[v.ID]; !ok {
					handbookCategories[v.ID] = make([]string, 0)
				}
			} else {
				handbookCategories[v.ParentID] = make([]string, 0)
				handbookCategories[v.ParentID] = append(handbookCategories[v.ParentID], v.ID)
			}
		} else {
			handbookCategories[v.ParentID] = append(handbookCategories[v.ParentID], v.ID)
		}
		if _, ok := handbookCategories[v.ID]; !ok {
			handbookCategories[v.ID] = make([]string, 0)
		}
	}

	db.template.prices = make(map[string]int32)
	for idx, v := range db.template.handbook.Items {
		db.template.index.Item[v.ID] = int16(idx)

		if _, ok := handbookCategories[v.ParentID]; !ok {
			handbookCategories[v.ParentID] = make([]string, 0)
			handbookCategories[v.ParentID] = append(handbookCategories[v.ParentID], v.ID)
		} else {
			handbookCategories[v.ParentID] = append(handbookCategories[v.ParentID], v.ID)
		}

		db.template.prices[v.ID] = v.Price
	}

	for id, v := range handbookCategories {
		if len(v) == 0 {
			delete(handbookCategories, id)
			continue
		}
		if len(v) != cap(v) {
			nv := make([]string, 0, len(v))
			nv = append(nv, v...)
			handbookCategories[id] = nv
		}
	}
}

func ConvertFromRouble(amount int32, currency string) (float64, error) {
	price, err := GetPriceByID(currency)
	if err != nil {
		return -1, err
	}
	return math.Round(float64(amount / price)), nil
}

func ConvertToRouble(amount int32, currency string) float64 {
	price, err := GetPriceByID(currency)
	if err != nil {
		log.Println(err)
	}
	return math.Round(float64(amount * (price)))
}

func (hbi *TemplateItem) SetHandbookItemEntry() {
	db.template.handbook.Items = append(db.template.handbook.Items, *hbi)
	db.template.index.Item[hbi.ID] = int16(len(db.template.handbook.Items) - 1)
}

func SetHandbookItemEntry(entry TemplateItem) {
	db.template.handbook.Items = append(db.template.handbook.Items, entry)
	db.template.index.Item[entry.ID] = int16(len(db.template.handbook.Items) - 1)
}

// #endregion

// #region Handbook structs

type Templates struct {
	Categories []TemplateCategories `json:"Categories"`
	Items      []TemplateItem       `json:"Items"`
}

type TemplateCategories struct {
	ID       string `json:"Id"`
	ParentID string `json:"ParentId"`
	Icon     string `json:"Icon"`
	Color    string `json:"Color"`
	Order    string `json:"Order"`
}

type TemplateItem struct {
	ID       string `json:"Id"`
	ParentID string `json:"ParentId"`
	Price    int32  `json:"Price"`
}

type Prices map[string]int32

// #endregion
