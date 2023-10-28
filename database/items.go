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
	"USD": "5696686a4bdc2da3298b456a",
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

func ItemClone(item string) *DatabaseItem {
	input := GetItemByUID(item)
	clone := new(DatabaseItem)

	clone.ID = input.ID
	clone.Name = input.Name
	clone.Parent = input.Parent
	clone.Type = input.Type
	clone.Props = input.Props
	clone.Proto = input.Proto

	return clone
}

func (i *DatabaseItem) Clone() *DatabaseItem {
	clone := new(DatabaseItem)

	clone.ID = i.ID
	clone.Name = i.Name
	clone.Parent = i.Parent
	clone.Type = i.Type
	clone.Props = i.Props
	clone.Proto = i.Proto

	return clone
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
	grids, ok := i.Props["Grids"].([]any)
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
	slots, ok := i.Props["Slots"].([]any)
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

func SetNewItem(entry DatabaseItem) {
	items[entry.ID] = &entry
}

func (i *DatabaseItem) GenerateNewUPD() *AssortItemUpd {

	itemUpd := new(AssortItemUpd)
	switch i.Parent {
	case "590c745b86f7743cc433c5f2":
		resource, ok := i.Props["Resource"].(float64)
		if !ok {
			fmt.Println("ballsack")
			return nil
		}
		itemUpd.Resource = new(Resource)
		itemUpd.Resource.Value = int16(resource)
		return itemUpd
	case "5448f3ac4bdc2dce718b4569":
		resource, ok := i.Props["MaxHpResource"].(float64)
		if !ok {
			fmt.Println("ballsack")
			return nil
		}

		itemUpd.MedKit = new(MedicalKit)
		itemUpd.MedKit.HpResource = int(resource)
		return itemUpd

	case "5448e8d04bdc2ddf718b4569", "5448e8d64bdc2dce718b4568":
		{
			resource, ok := i.Props["MaxResource"].(float64)
			if !ok {
				fmt.Println("ballsack")
				return nil
			}

			itemUpd.FoodDrink = new(FoodDrink)
			itemUpd.FoodDrink.HpPercent = int16(resource)
			return itemUpd
		}
	case "5a341c4086f77401f2541505", "5448e5284bdc2dcb718b4567",
		"57bef4c42459772e8d35a53b", "5a341c4686f77469e155819e",
		"5447e1d04bdc2dff2f8b4567", "5448e54d4bdc2dcc718b4568",
		"5448e5724bdc2ddf718b4568":
		{
			maxDurability, ok := i.Props["MaxDurability"].(float64)
			if !ok {
				fmt.Println("ballsack")
				return nil
			}

			durability, ok := i.Props["Durability"].(float64)
			if !ok {
				fmt.Println("ballsack")
				return nil
			}

			itemUpd.Repairable = new(Repairable)
			itemUpd.Repairable.MaxDurability = int(maxDurability)
			itemUpd.Repairable.Durability = int(durability)
			return itemUpd
		}
	case "55818ae44bdc2dde698b456c", "55818ac54bdc2d5b648b456e",
		"55818acf4bdc2dde698b456b", "55818ad54bdc2ddc698b4569",
		"55818add4bdc2d5b648b456f", "55818aeb4bdc2ddc698b456a":
		{

			itemUpd.Sight = new(Sight)
			itemUpd.Sight.ScopesCurrentCalibPointIndexes = []int{0}
			itemUpd.Sight.ScopesSelectedModes = []int{0}
			itemUpd.Sight.SelectedScope = 0

			return itemUpd
		}
	case "5447bee84bdc2dc3278b4569", "5447bedf4bdc2d87278b4568",
		"5447bed64bdc2d97278b4568", "5447b6254bdc2dc3278b4568",
		"5447b6194bdc2d67278b4567", "5447b6094bdc2dc3278b4567",
		"5447b5fc4bdc2d87278b4567", "5447b5f14bdc2d61278b4567",
		"5447b5e04bdc2d62278b4567", "617f1ef5e8b54b0998387733":
		{
			maxDurability, ok := i.Props["MaxDurability"].(float64)
			if !ok {
				fmt.Println("ballsack")
				return nil
			}

			durability, ok := i.Props["Durability"].(float64)
			if !ok {
				fmt.Println("ballsack")
				return nil
			}

			itemUpd.Repairable = new(Repairable)
			itemUpd.Foldable = new(Foldable)
			itemUpd.FireMode = new(FireMode)

			itemUpd.Repairable.MaxDurability = int(maxDurability)
			itemUpd.Repairable.Durability = int(durability)
			itemUpd.Foldable.Folded = false
			itemUpd.FireMode.FireMode = "single"

			return itemUpd
		}

	case "5447b5cf4bdc2d65278b4567":
		{
			maxDurability, ok := i.Props["MaxDurability"].(float64)
			if !ok {
				fmt.Println("ballsack")
				return nil
			}

			durability, ok := i.Props["Durability"].(float64)
			if !ok {
				fmt.Println("ballsack")
				return nil
			}

			itemUpd.Repairable = new(Repairable)
			itemUpd.FireMode = new(FireMode)

			itemUpd.Repairable.MaxDurability = int(maxDurability)
			itemUpd.Repairable.Durability = int(durability)
			itemUpd.FireMode.FireMode = "single"

			return itemUpd
		}
	case "616eb7aea207f41933308f46":
		{
			maxRepairResource, ok := i.Props["MaxRepairResource"].(float64)
			if !ok {
				fmt.Println("ballsack")
				return nil
			}

			itemUpd.RepairKit = new(RepairKit)
			itemUpd.RepairKit.Resource = int16(maxRepairResource)
			return itemUpd
		}
	case "5485a8684bdc2da71d8b4567":
		{
			stackMaxSize, ok := i.Props["StackMaxSize"].(float64)
			if !ok {
				fmt.Println("ballsack")
				return nil
			}

			itemUpd.StackObjectsCount = int(stackMaxSize)
			return itemUpd
		}
	case "55818b084bdc2d5b648b4571", "55818b164bdc2ddc698b456c":
		{
			itemUpd.Light = new(Light)

			itemUpd.Light.IsActive = false
			itemUpd.Light.SelectedMode = 0
			return itemUpd
		}
	default:
		log.Println(i.Parent, "does not require a UPD")
		return nil
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

type DatabaseItemProperties map[string]any

// #endregion
