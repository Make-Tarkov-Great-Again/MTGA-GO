package data

import (
	"fmt"
	"github.com/alphadose/haxmap"
	"log"
	"math"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

type Template struct {
	index     *TemplateIndex
	handbook  *Templates
	blacklist *haxmap.Map[string, string]
}

type TemplateIndex struct {
	Item       *HandbookItemIndex
	Categories *HandbookCategoryIndex
}

type HandbookCategoryIndex struct {
	Index *haxmap.Map[string, int16]    //map[string]int16
	Main  *haxmap.Map[string, []string] //map[string][]string
	Sub   *haxmap.Map[string, []string] //map[string][]string
}

type HandbookItemIndex struct {
	Prices *haxmap.Map[string, int32] //map[string]int32
	Index  *haxmap.Map[string, int16] //map[string]int16
}

// #region Handbook getters

func GetHandbook() *Templates {
	return db.template.handbook
}

var currencyName = map[string]string{
	"RUB": "5449016a4bdc2d6f028b456f",
	"EUR": "569668774bdc2da2298b4568",
	"USD": "5696686a4bdc2da3298b456a",
}

var currencyByID = map[string]*struct{}{
	"5449016a4bdc2d6f028b456f": nil, //RUB
	"569668774bdc2da2298b4568": nil, //EUR
	"5696686a4bdc2da3298b456a": nil, //USD
}

func IsCurrencyByID(UID string) bool {
	_, ok := currencyByID[UID]
	return ok
}

func GetCurrencyByName(name string) *string {
	currency, ok := currencyName[name]
	if ok {
		return &currency
	}
	return nil
}

// GetPrices Get prices of all items
func GetPrices() *haxmap.Map[string, int32] {
	return db.template.index.Item.Prices
}

const priceNotFound string = "price of %s not found"

// GetPriceByID Get item price by ID
func GetPriceByID(id string) (int32, error) {
	price, ok := db.template.index.Item.Prices.Get(id)
	if !ok {
		return 999999999, fmt.Errorf(priceNotFound, id)
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
	categories, ok := db.template.index.Categories.Main.Get(id)
	if !ok {
		return nil, fmt.Errorf("sub category %s does not exist", id)
	}

	if _, ok := db.template.index.Categories.Main.Get(categories[0]); !ok {
		return categories, nil
	}

	output := make([]string, 0)
	for _, c := range categories {
		category, ok := db.template.index.Categories.Main.Get(c)
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
	categories, ok := db.template.index.Categories.Sub.Get(id)
	if !ok {
		return nil, fmt.Errorf("sub category %s does not exist", id)
	}

	if _, ok := db.template.index.Categories.Sub.Get(categories[0]); !ok {
		return categories, nil
	}

	output := make([]string, 0)
	for _, c := range categories {
		category, ok := db.template.index.Categories.Sub.Get(c)
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
			Prices: haxmap.New[string, int32](), //make(map[string]int32),
			Index:  haxmap.New[string, int16](), //make(map[string]int16),
		},
		Categories: &HandbookCategoryIndex{
			Index: haxmap.New[string, int16](),    //make(map[string]int16),
			Main:  haxmap.New[string, []string](), //make(map[string][]string),
			Sub:   haxmap.New[string, []string](), //make(map[string][]string),
		},
	}

	temp := make(map[string][]string)
	for idx, category := range db.template.handbook.Categories {
		db.template.index.Categories.Index.Set(category.ID, int16(idx))

		if _, ok := db.template.index.Categories.Main.Get(category.ID); !ok && category.ParentID == "" {
			db.template.index.Categories.Main.Set(category.ID, make([]string, 0))
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
		if _, ok := db.template.index.Categories.Main.Get(key); ok {
			db.template.index.Categories.Main.Set(key, value)
			continue
		}

		if _, ok := db.template.index.Categories.Sub.Get(key); ok {
			db.template.index.Categories.Sub.Set(key, value)
			continue
		}

		db.template.index.Categories.Main.Set(key, value)
	}

	db.template.index.Categories.Main.ForEach(func(key string, value []string) bool {
		if len(value) == 0 {
			db.template.index.Categories.Sub.Set(key, value)
			db.template.index.Categories.Main.Del(key)
		}
		return true
	})

	temp = make(map[string][]string)
	for idx, item := range db.template.handbook.Items {
		db.template.index.Item.Index.Set(item.ID, int16(idx))
		db.template.index.Item.Prices.Set(item.ID, item.Price)

		if _, ok := temp[item.ParentID]; !ok {
			temp[item.ParentID] = make([]string, 0)
			temp[item.ParentID] = append(temp[item.ParentID], item.ID)
			continue
		}
		temp[item.ParentID] = append(temp[item.ParentID], item.ID)
	}

	for key, value := range temp {
		if _, ok := db.template.index.Categories.Sub.Get(key); ok {
			sub, _ := db.template.index.Categories.Sub.GetOrSet(key, make([]string, 0, len(value)))
			sub = append(sub, value...)
			continue
		}
		db.template.index.Categories.Sub.Set(key, value)
	}

	setItemBlacklist()
}

func setItemBlacklist() {
	db.template.blacklist = haxmap.New[string, string]()
	db.item.ForEach(func(key string, item *DatabaseItem) bool {
		if _, ok := db.template.index.Item.Index.Get(key); !ok {
			if item.Type == "Node" {
				db.template.blacklist.Set(key, "node")
			} else {
				db.template.blacklist.Set(key, "item")
			}
		}
		return true
	})
}

func IsItemBlacklist(id string) (string, bool) {
	value, ok := db.template.blacklist.Get(id)
	return value, ok
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
	db.template.index.Item.Index.Set(hbi.ID, int16(len(db.template.handbook.Items)-1))
}

func SetHandbookItemEntry(entry TemplateItem) {
	db.template.handbook.Items = append(db.template.handbook.Items, entry)
	db.template.index.Item.Index.Set(entry.ID, int16(len(db.template.handbook.Items)-1))
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
