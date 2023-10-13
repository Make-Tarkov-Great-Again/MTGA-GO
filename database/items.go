package database

import (
	"fmt"
	"log"
	"math"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var items map[string]*DatabaseItem

// #region Item getters

func GetItems() map[string]*DatabaseItem {
	return items
}

func GetItemByUID(uid string) *DatabaseItem {
	item, ok := items[uid]
	if !ok {
		fmt.Println("Item:", uid, " not found in database")
		return nil
	}
	return item
}

var currencyName = map[string]string{
	"RUB": "5449016a4bdc2d6f028b456f",
	"EUR": "569668774bdc2da2298b4568",
	"DOL": "5696686a4bdc2da3298b456a",
}

var currencyByID = map[string]struct{}{
	"5449016a4bdc2d6f028b456f": {}, //RUB
	"569668774bdc2da2298b4568": {}, //EUR
	"5696686a4bdc2da3298b456a": {}, //DOL
}

func IsCurrencyByUID(UID string) bool {
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

func ConvertToRoubles(amount int32, currency string) float64 {
	price := *GetPriceByID(currency)
	return math.Round(float64(amount * price))
}

// GetItemPrice Gets item... price...
func (i *DatabaseItem) GetItemPrice() *int32 {
	return GetPriceByID(i.ID)
}

// GetItemSize Get item height and width
func (i *DatabaseItem) GetItemSize() (int8, int8) {
	height, ok := i.Props["Height"].(float64)
	if !ok {
		fmt.Println("Item:", i.ID, " does not have a height")
		return -1, -1
	}

	width, ok := i.Props["Width"].(float64)
	if !ok {
		fmt.Println("Item:", i.ID, " does not have a width")
		return -1, -1
	}

	return int8(height), int8(width)
}

func (i *DatabaseItem) GetItemForcedSize(sizes *sizes) {
	extraSize, ok := i.Props["ExtraSizeForceAdd"].(bool)
	if !ok {
		return
	}

	extraSizeDown := int8(i.Props["ExtraSizeDown"].(float64))
	extraSizeUp := int8(i.Props["ExtraSizeUp"].(float64))
	extraSizeLeft := int8(i.Props["ExtraSizeLeft"].(float64))
	extraSizeRight := int8(i.Props["ExtraSizeRight"].(float64))

	if extraSize {
		sizes.ForcedDown += extraSizeDown
		sizes.ForcedUp += extraSizeUp
		sizes.ForcedLeft += extraSizeLeft
		sizes.ForcedRight += extraSizeRight
	} else {
		sizes.SizeUp = max(sizes.SizeUp, extraSizeUp)
		sizes.SizeDown = max(sizes.SizeDown, extraSizeDown)
		sizes.SizeLeft = max(sizes.SizeLeft, extraSizeLeft)
		sizes.SizeRight = max(sizes.SizeRight, extraSizeRight)
	}
}

type Grid struct {
	Name   string    `json:"_name"`
	ID     string    `json:"_id"`
	Parent string    `json:"_parent"`
	Props  GridProps `json:"_props"`
	Proto  string    `json:"_proto"`
}
type GridFilters struct {
	Filter         []string `json:"Filter"`
	ExcludedFilter []string `json:"ExcludedFilter"`
}
type GridProps struct {
	Filters        []GridFilters `json:"filters"`
	CellsH         int8          `json:"cellsH"`
	CellsV         int8          `json:"cellsV"`
	MinCount       int8          `json:"minCount"`
	MaxCount       int8          `json:"maxCount"`
	MaxWeight      int8          `json:"maxWeight"`
	IsSortingTable bool          `json:"isSortingTable"`
}

// GetItemGrids Get the grid property from the item if it exists
func (i *DatabaseItem) GetItemGrids() map[string]*Grid {
	grids, ok := i.Props["Grids"].([]interface{})
	if !ok || len(grids) == 0 {
		fmt.Println("Item:", i.ID, " does not have Grid property")
		return nil
	}

	output := make(map[string]*Grid)
	for _, g := range grids {
		grid := new(Grid)
		data, err := json.Marshal(g)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(data, grid)
		if err != nil {
			log.Fatalln(err)
		}
		output[grid.Name] = grid
	}

	return output
}

type Slot struct {
	Name                  string    `json:"_name"`
	ID                    string    `json:"_id"`
	Parent                string    `json:"_parent"`
	Props                 SlotProps `json:"_props"`
	Required              bool      `json:"_required"`
	MergeSlotWithChildren bool      `json:"_mergeSlotWithChildren"`
	Proto                 string    `json:"_proto"`
}
type SlotFilters struct {
	Filter []string `json:"Filter"`
}
type SlotProps struct {
	Filters []SlotFilters `json:"filters"`
}

// GetItemSlots Get the slot property from the item if it exists
func (i *DatabaseItem) GetItemSlots() map[string]*Slot {
	slots, ok := i.Props["Slots"].([]interface{})
	if !ok || len(slots) == 0 {
		fmt.Println("Item:", i.ID, " does not have Grid property")
		return nil
	}

	output := make(map[string]*Slot)
	for _, s := range slots {
		slot := new(Slot)
		data, err := json.Marshal(s)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(data, slot)
		if err != nil {
			log.Fatalln(err)
		}
		output[slot.Name] = slot
	}

	return output
}

func (i *DatabaseItem) GetStackMaxSize() *int32 {
	if size, ok := i.Props["StackMaxSize"].(float64); ok {
		value := int32(size)
		return &value
	}
	return nil
}

// #endregion

// #region Item setters

func setItems() {
	raw := tools.GetJSONRawMessage(itemsPath)
	err := json.Unmarshal(raw, &items)
	if err != nil {
		log.Fatalln(err)
	}
}

// #endregion

// #region Item structs

type DatabaseItem struct {
	ID     string                 `json:"_id"`
	Name   string                 `json:"_name"`
	Parent string                 `json:"_parent"`
	Type   string                 `json:"_type"`
	Props  DatabaseItemProperties `json:"_props"`
	Proto  string                 `json:"_proto"`
}

type DatabaseItemProperties map[string]interface{}

// #endregion
