package database

import (
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
)

var items map[string]*DatabaseItem

// #region Item getters

func GetItems() map[string]*DatabaseItem {
	return items
}

func GetItemByID(uid string) *DatabaseItem {
	item, ok := items[uid]
	if !ok {
		log.Println("Item:", uid, " not found in database")
		return nil
	}
	return item
}

func ItemClone(item string) *DatabaseItem {
	input := GetItemByID(item)
	clone := new(DatabaseItem)

	data, err := json.MarshalNoEscape(input)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &clone); err != nil {
		log.Fatal(err)
	}

	return clone
}

func (i *DatabaseItem) Clone() *DatabaseItem {
	clone := new(DatabaseItem)

	data, err := json.MarshalNoEscape(i)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &clone); err != nil {
		log.Fatal(err)
	}

	return clone
}

// GetItemPrice Gets item... price...
func (i *DatabaseItem) GetItemPrice() (*int32, error) {
	price, err := GetPriceByID(i.ID)
	if err != nil {
		return nil, err
	}
	return price, nil
}

// GetItemSize Get item height and width
func (i *DatabaseItem) GetItemSize() (int8, int8) {
	height, ok := i.Props["Height"].(float64)
	if !ok {
		log.Println("Item:", i.ID, " does not have a height")
		return -1, -1
	}

	width, ok := i.Props["Width"].(float64)
	if !ok {
		log.Println("Item:", i.ID, " does not have a width")
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

// GetItemGrids Get the grid property from the item if it exists
func (i *DatabaseItem) GetItemGrids() map[string]*Grid {
	grids, ok := i.Props["Grids"].([]any)
	if !ok {
		log.Println("Item:", i.ID, " does not have Grid property")
		return nil
	} else if len(grids) == 0 {
		log.Println("Item:", i.ID, " does not have any Grid components")
		return nil
	}

	output := make(map[string]*Grid)
	for _, g := range grids {
		grid := new(Grid)
		data, err := json.MarshalNoEscape(g)
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

// GetItemSlots Get the slot property from the item if it exists
func (i *DatabaseItem) GetItemSlots() map[string]*Slot {
	slots, ok := i.Props["Slots"].([]any)
	if !ok {
		log.Println("Item:", i.ID, " does not have Slot property")
		return nil
	} else if len(slots) == 0 {
		log.Println("Item:", i.ID, " does not have Slot components")
		return nil
	}

	output := make(map[string]*Slot)
	for _, s := range slots {
		slot := new(Slot)
		data, err := json.MarshalNoEscape(s)
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

func (i *DatabaseItem) GetStackMaxSize() int32 {
	size, _ := i.Props["StackMaxSize"].(float64)
	return int32(size)
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

func (i *DatabaseItem) CreateItemUPD() (*ItemUpdate, error) {

	itemUpd := new(ItemUpdate)
	switch i.Parent {
	case "590c745b86f7743cc433c5f2":
		resource, ok := i.Props["Resource"].(float64)
		if !ok {
			return nil, fmt.Errorf("i.Props[\"Resource\"] is not a float64")
		}
		itemUpd.Resource = new(Resource)
		itemUpd.Resource.Value = int16(resource)
		return itemUpd, nil
	case "5448f3ac4bdc2dce718b4569":
		resource, ok := i.Props["MaxHpResource"].(float64)
		if !ok {
			return nil, fmt.Errorf("i.Props[\"MaxHpResource\"] is not a float64")
		}

		itemUpd.MedKit = new(MedicalKit)
		itemUpd.MedKit.HpResource = int(resource)
		return itemUpd, nil

	case "5448e8d04bdc2ddf718b4569", "5448e8d64bdc2dce718b4568":
		{
			resource, ok := i.Props["MaxResource"].(float64)
			if !ok {
				return nil, fmt.Errorf("i.Props[\"MaxResource\"] is not a float64")
			}

			itemUpd.FoodDrink = new(FoodDrink)
			itemUpd.FoodDrink.HpPercent = int16(resource)
			return itemUpd, nil
		}
	case "5a341c4086f77401f2541505", "5448e5284bdc2dcb718b4567",
		"57bef4c42459772e8d35a53b", "5a341c4686f77469e155819e",
		"5447e1d04bdc2dff2f8b4567", "5448e54d4bdc2dcc718b4568",
		"5448e5724bdc2ddf718b4568":
		{
			maxDurability, ok := i.Props["MaxDurability"].(float64)
			if !ok {
				return nil, fmt.Errorf("i.Props[\"MaxDurability\"] is not a float64")
			}

			durability, ok := i.Props["Durability"].(float64)
			if !ok {
				return nil, fmt.Errorf("i.Props[\"Durability\"] is not a float64")
			}

			itemUpd.Repairable = new(Repairable)
			itemUpd.Repairable.MaxDurability = int(maxDurability)
			itemUpd.Repairable.Durability = int(durability)
			return itemUpd, nil
		}
	case "55818ae44bdc2dde698b456c", "55818ac54bdc2d5b648b456e",
		"55818acf4bdc2dde698b456b", "55818ad54bdc2ddc698b4569",
		"55818add4bdc2d5b648b456f", "55818aeb4bdc2ddc698b456a":
		{

			itemUpd.Sight = new(Sight)
			itemUpd.Sight.ScopesCurrentCalibPointIndexes = []int{0}
			itemUpd.Sight.ScopesSelectedModes = []int{0}
			itemUpd.Sight.SelectedScope = 0

			return itemUpd, nil
		}
	case "5447bee84bdc2dc3278b4569", "5447bedf4bdc2d87278b4568",
		"5447bed64bdc2d97278b4568", "5447b6254bdc2dc3278b4568",
		"5447b6194bdc2d67278b4567", "5447b6094bdc2dc3278b4567",
		"5447b5fc4bdc2d87278b4567", "5447b5f14bdc2d61278b4567",
		"5447b5e04bdc2d62278b4567", "617f1ef5e8b54b0998387733":
		{
			maxDurability, ok := i.Props["MaxDurability"].(float64)
			if !ok {
				return nil, fmt.Errorf("i.Props[\"MaxDurability\"] is not a float64")
			}

			durability, ok := i.Props["Durability"].(float64)
			if !ok {
				return nil, fmt.Errorf("i.Props[\"Durability\"] is not a float64")
			}

			itemUpd.Repairable = new(Repairable)
			itemUpd.Foldable = new(Foldable)
			itemUpd.FireMode = new(FireMode)

			itemUpd.Repairable.MaxDurability = int(maxDurability)
			itemUpd.Repairable.Durability = int(durability)
			itemUpd.Foldable.Folded = false
			itemUpd.FireMode.FireMode = "single"

			return itemUpd, nil
		}

	case "5447b5cf4bdc2d65278b4567":
		{
			maxDurability, ok := i.Props["MaxDurability"].(float64)
			if !ok {
				return nil, fmt.Errorf("i.Props[\"MaxDurability\"] is not a float64")
			}

			durability, ok := i.Props["Durability"].(float64)
			if !ok {
				return nil, fmt.Errorf("i.Props[\"Durability\"] is not a float64")
			}

			itemUpd.Repairable = new(Repairable)
			itemUpd.FireMode = new(FireMode)

			itemUpd.Repairable.MaxDurability = int(maxDurability)
			itemUpd.Repairable.Durability = int(durability)
			itemUpd.FireMode.FireMode = "single"

			return itemUpd, nil
		}
	case "616eb7aea207f41933308f46":
		{
			maxRepairResource, ok := i.Props["MaxRepairResource"].(float64)
			if !ok {
				return nil, fmt.Errorf("i.Props[\"MaxRepairResource\"] is not a float64")
			}

			itemUpd.RepairKit = new(RepairKit)
			itemUpd.RepairKit.Resource = int16(maxRepairResource)
			return itemUpd, nil
		}
	case "5485a8684bdc2da71d8b4567":
		{
			stackMaxSize, ok := i.Props["StackMaxSize"].(float64)
			if !ok {
				return nil, fmt.Errorf("i.Props[\"StackMaxSize\"] is not a float64")
			}

			itemUpd.StackObjectsCount = int32(stackMaxSize)
			return itemUpd, nil
		}
	case "55818b084bdc2d5b648b4571", "55818b164bdc2ddc698b456c":
		{
			itemUpd.Light = new(Light)

			itemUpd.Light.IsActive = false
			itemUpd.Light.SelectedMode = 0
			return itemUpd, nil
		}
	default:
		return nil, fmt.Errorf(i.Parent, "does not require a UPD")
	}
}

const weaponBaseClass = "5422acb9af1c889c16000029"

func (i *DatabaseItem) IsWeapon() bool {
	if i.ID == weaponBaseClass {
		log.Println("Why are you looking at the node?")
		return false
	}

	if i.Parent == weaponBaseClass || items[i.Parent].Parent == weaponBaseClass {
		return true
	}

	return false
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
	Shift  int      `json:"Shift,omitempty"`
	Filter []string `json:"Filter"`
}

type SlotProps struct {
	Filters []SlotFilters `json:"filters"`
}

type Cartridges struct {
	Name     string           `json:"_name"`
	Id       string           `json:"_id"`
	Parent   string           `json:"_parent"`
	MaxCount int              `json:"_max_count"`
	Props    CartridgeFilters `json:"_props"`
	Proto    string           `json:"_proto"`
}

type CartridgeFilters struct {
	Filters []CartridgeFilter `json:"filters"`
}

type CartridgeFilter struct {
	Filter []string `json:"Filter"`
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
	MaxWeight      int16         `json:"maxWeight"`
	IsSortingTable bool          `json:"isSortingTable"`
}

// #endregion
