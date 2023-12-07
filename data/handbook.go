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
}

type TemplateIndex struct {
	Item       *HandbookItemIndex
	Categories *HandbookCategoryIndex
}

type HandbookCategoryIndex struct {
	Index map[string]int16
	Main  map[string][]string
	Sub   map[string][]string
}

type HandbookItemIndex struct {
	Prices map[string]int32
	Index  map[string]int16
}

// #region Handbook getters

func GetHandbook() *Templates {
	return db.template.handbook
}

// GetPrices Get prices of all items
func GetPrices() map[string]int32 {
	return db.template.index.Item.Prices
}

const priceNotFound string = "price of %s not found"

// GetPriceByID Get item price by ID
func GetPriceByID(id string) (int32, error) {
	price, ok := db.template.index.Item.Prices[id]
	if !ok {
		return -1, fmt.Errorf(priceNotFound, id)
	}
	return price, nil
}

// #endregion

// #region Handbook setters

func setHandbook() {
	db.template = &Template{
		handbook: new(Templates),
	}
	raw := tools.GetJSONRawMessage(handbookPath)
	if err := json.UnmarshalNoEscape(raw, &db.template.handbook); err != nil {
		log.Fatalln(err)
	}
}

func HasGetMainHandbookCategory(id string) ([]string, error) {
	categories, ok := db.template.index.Categories.Main[id]
	if !ok {
		return nil, fmt.Errorf("sub category %s does not exist", id)
	}

	if _, ok := db.template.index.Categories.Main[categories[0]]; !ok {
		return categories, nil
	}

	output := make([]string, 0)
	for _, c := range categories {
		category, ok := db.template.index.Categories.Main[c]
		if !ok {
			continue
		}
		output = append(output, category...)
	}

	if len(output) != 0 {
		return output, nil
	}

	return nil, fmt.Errorf("main category %s does not exist", id)
}

func HasGetHandbookSubCategory(id string) ([]string, error) {
	categories, ok := db.template.index.Categories.Sub[id]
	if !ok {
		return nil, fmt.Errorf("sub category %s does not exist", id)
	}

	if _, ok := db.template.index.Categories.Sub[categories[0]]; !ok {
		return categories, nil
	}

	output := make([]string, 0)
	for _, c := range categories {
		category, ok := db.template.index.Categories.Sub[c]
		if !ok {
			continue
		}
		output = append(output, category...)
	}

	if len(output) != 0 {
		return output, nil
	}

	return nil, fmt.Errorf("sub category %s does not exist", id)
}

func setHandbookIndex() {
	db.template.index = &TemplateIndex{
		Item: &HandbookItemIndex{
			Prices: make(map[string]int32),
			Index:  make(map[string]int16),
		},
		Categories: &HandbookCategoryIndex{
			Index: make(map[string]int16),
			Main:  make(map[string][]string),
			Sub:   make(map[string][]string),
		},
	}

	temp := make(map[string][]string)
	for idx, category := range db.template.handbook.Categories {
		db.template.index.Categories.Index[category.ID] = int16(idx)

		if _, ok := db.template.index.Categories.Main[category.ID]; !ok && category.ParentID == "" {
			db.template.index.Categories.Main[category.ID] = make([]string, 0)
			continue
		}

		if _, ok := temp[category.ParentID]; !ok {
			temp[category.ParentID] = make([]string, 0)
			temp[category.ParentID] = append(temp[category.ParentID], category.ID)
			continue
		}
		temp[category.ParentID] = append(temp[category.ParentID], category.ID)
	}

	for key, value := range temp {
		if _, ok := db.template.index.Categories.Main[key]; ok {
			db.template.index.Categories.Main[key] = value
			continue
		}

		if _, ok := db.template.index.Categories.Sub[key]; ok {
			db.template.index.Categories.Sub[key] = value
			continue
		}

		db.template.index.Categories.Main[key] = value
	}

	for key, value := range db.template.index.Categories.Main {
		if len(value) == 0 {
			db.template.index.Categories.Sub[key] = value
			delete(db.template.index.Categories.Main, key)
		}
	}

	temp = make(map[string][]string)
	for idx, item := range db.template.handbook.Items {
		db.template.index.Item.Index[item.ID] = int16(idx)
		db.template.index.Item.Prices[item.ID] = item.Price

		if _, ok := temp[item.ParentID]; !ok {
			temp[item.ParentID] = make([]string, 0)
			temp[item.ParentID] = append(temp[item.ParentID], item.ID)
			continue
		}
		temp[item.ParentID] = append(temp[item.ParentID], item.ID)
	}

	for key, value := range temp {
		if _, ok := db.template.index.Categories.Sub[key]; ok {
			db.template.index.Categories.Sub[key] = make([]string, 0, len(value))
			db.template.index.Categories.Sub[key] = append(db.template.index.Categories.Sub[key], value...)
			continue
		}
		db.template.index.Categories.Sub[key] = value
	}

	//_ = tools.WriteToFile("faggot.json", db.template.index)
	//fmt.Println()
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
	db.template.index.Item.Index[hbi.ID] = int16(len(db.template.handbook.Items) - 1)
}

func SetHandbookItemEntry(entry TemplateItem) {
	db.template.handbook.Items = append(db.template.handbook.Items, entry)
	db.template.index.Item.Index[entry.ID] = int16(len(db.template.handbook.Items) - 1)
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
