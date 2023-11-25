package data

import (
	"fmt"
	"log"
	"math"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var handbook *Handbook
var handbookIndex = map[string]int16{}
var prices = make(map[string]int32)

type templates struct {
	index    map[string]int16
	handbook *Handbook
	prices   map[string]int32
}

// #region Handbook getters

func GetHandbook() *Handbook {
	return handbook
}

// GetPrices Get prices of all items
func GetPrices() map[string]int32 {
	return prices
}

const priceNotFound string = "price of %s not found"

// GetPriceByID Get item price by ID
func GetPriceByID(id string) (*int32, error) {
	price, ok := prices[id]
	if !ok {
		return nil, fmt.Errorf(priceNotFound, id)
	}
	return &price, nil
}

// #endregion

// #region Handbook setters
var handbookCategories = make(map[string][]string)

func setHandbook() {
	handbook = new(Handbook)
	raw := tools.GetJSONRawMessage(handbookPath)
	if err := json.UnmarshalNoEscape(raw, &handbook); err != nil {
		log.Fatalln(err)
	}

	for _, v := range handbook.Categories {
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

	for idx, v := range handbook.Items {
		handbookIndex[v.ID] = int16(idx)

		if _, ok := handbookCategories[v.ParentID]; !ok {
			handbookCategories[v.ParentID] = make([]string, 0)
			handbookCategories[v.ParentID] = append(handbookCategories[v.ParentID], v.ID)
		} else {
			handbookCategories[v.ParentID] = append(handbookCategories[v.ParentID], v.ID)
		}

		prices[v.ID] = v.Price
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
	fmt.Println()
}

func ConvertFromRouble(amount int32, currency string) (float64, error) {
	price, err := GetPriceByID(currency)
	if err != nil {
		return -1, err
	}
	return math.Round(float64(amount / *price)), nil
}

func ConvertToRouble(amount int32, currency string) float64 {
	price, err := GetPriceByID(currency)
	if err != nil {
		log.Println(err)
	}
	return math.Round(float64(amount * (*price)))
}

func (hbi *HandbookItem) SetHandbookItemEntry() {
	handbook.Items = append(handbook.Items, *hbi)
	handbookIndex[hbi.ID] = int16(len(handbook.Items) - 1)
}

func SetHandbookItemEntry(entry HandbookItem) {
	handbook.Items = append(handbook.Items, entry)
	handbookIndex[entry.ID] = int16(len(handbook.Items) - 1)
}

// #endregion

// #region Handbook structs

type Handbook struct {
	Categories []HandbookCategory `json:"Categories"`
	Items      []HandbookItem     `json:"Items"`
}

type HandbookCategory struct {
	ID       string `json:"Id"`
	ParentID string `json:"ParentId"`
	Icon     string `json:"Icon"`
	Color    string `json:"Color"`
	Order    string `json:"Order"`
}

type HandbookItem struct {
	ID       string `json:"Id"`
	ParentID string `json:"ParentId"`
	Price    int32  `json:"Price"`
}

type Prices map[string]int32

// #endregion
