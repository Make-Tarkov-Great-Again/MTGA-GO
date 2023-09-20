package database

import (
	"MT-GO/tools"
	"fmt"

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

// Get item... price...
func (i *DatabaseItem) GetItemPrice() *int32 {
	return GetPriceByID(i.ID)
}

// Get item height and width
func (i *DatabaseItem) GetItemSize() (int8, int8) {
	height, ok := i.Props["Height"].(float64)
	if !ok {
		fmt.Println("Item:", i.ID, " does not have a height")
		return 0, 0
	}

	width, ok := i.Props["Width"].(float64)
	if !ok {
		fmt.Println("Item:", i.ID, " does not have a width")
		return 0, 0
	}

	return int8(height), int8(width)
}

func (i *DatabaseItem) GetItemForcedSize() (int8, int8) {
	extraSize, ok := i.Props["ExtraSizeForceAdd"].(bool)
	if !ok || !extraSize {
		return 0, 0
	}

	var height, width int8 = 0, 0

	down, ok := i.Props["ExtraSizeDown"].(float64)
	if ok {
		height += int8(down)
	}

	up, ok := i.Props["ExtraSizeUp"].(float64)
	if ok {
		height += int8(up)
	}

	left, ok := i.Props["ExtraSizeLeft"].(float64)
	if ok {
		width += int8(left)
	}

	right, ok := i.Props["ExtraSizeRight"].(float64)
	if ok {
		width += int8(right)
	}

	return height, width
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

// Get the grid property from the item if it exists
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
			panic(err)
		}
		err = json.Unmarshal(data, grid)
		if err != nil {
			panic(err)
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

// Get the slot property from the item if it exists
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
			panic(err)
		}
		err = json.Unmarshal(data, slot)
		if err != nil {
			panic(err)
		}
		output[slot.Name] = slot
	}

	return output
}

// #endregion

// #region Item setters

func setItems() {
	raw := tools.GetJSONRawMessage(itemsPath)
	err := json.Unmarshal(raw, &items)
	if err != nil {
		panic(err)
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
