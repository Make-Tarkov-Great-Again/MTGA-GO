package database

import (
	"MT-GO/tools"
	"fmt"
	"log"
)

type Inventory struct {
	Items              []InventoryItem `json:"items"`
	Equipment          string          `json:"equipment"`
	Stash              string          `json:"stash"`
	SortingTable       string          `json:"sortingTable"`
	QuestRaidItems     string          `json:"questRaidItems"`
	QuestStashItems    string          `json:"questStashItems"`
	FastPanel          interface{}     `json:"fastPanel"`
	HideoutAreaStashes interface{}     `json:"hideoutAreaStashes"`
}

type InventoryItem struct {
	ID       string                 `json:"_id"`
	TPL      string                 `json:"_tpl,omitempty"`
	Location *InventoryItemLocation `json:"location,omitempty"`
	ParentID *string                `json:"parentId,omitempty"`
	SlotID   *string                `json:"slotId,omitempty"`
	UPD      *InventoryItemUpd      `json:"upd,omitempty"`
}

type InventoryItemUpd struct {
	StackObjectsCount *int32      `json:"StackObjectsCount,omitempty"`
	FireMode          *FireMode   `json:"FireMode,omitempty"`
	Foldable          *Foldable   `json:"Foldable,omitempty"`
	Repairable        *Repairable `json:"Repairable,omitempty"`
	Sight             *Sight      `json:"Sight,omitempty"`
	MedKit            *MedicalKit `json:"MedKit,omitempty"`
	FoodDrink         *FoodDrink  `json:"FoodDrink,omitempty"`
	RepairKit         *RepairKit  `json:"RepairKit,omitempty"`
	Light             *Light      `json:"Light,omitempty"`
	Resource          *Resource   `json:"Resource,omitempty"`
}

type Resource struct {
	Value int16 `json:"Value"`
}

type Light struct {
	IsActive     bool `json:"IsActive"`
	SelectedMode int8 `json:"SelectedMode"`
}
type RepairKit struct {
	Resource int16 `json:"Resource"`
}
type FoodDrink struct {
	HpPercent int16 `json:"HpPercent"`
}

type MedicalKit struct {
	HpResource int `json:"HpResource"`
}

type Sight struct {
	ScopesCurrentCalibPointIndexes []int `json:"ScopesCurrentCalibPointIndexes"`
	ScopesSelectedModes            []int `json:"ScopesSelectedModes"`
	SelectedScope                  int   `json:"SelectedScope"`
}

type Repairable struct {
	Durability    int `json:"Durability"`
	MaxDurability int `json:"MaxDurability"`
}
type Foldable struct {
	Folded bool `json:"Folded"`
}
type FireMode struct {
	FireMode string `json:"FireMode"`
}

type InventoryItemLocation struct {
	IsSearched bool        `json:"isSearched"`
	R          interface{} `json:"r"`
	X          interface{} `json:"x"`
	Y          interface{} `json:"y"`
}

type InventoryContainer struct {
	Stash  *Stash
	Lookup *Lookup
}

type Lookup struct {
	Forward map[string]int16
	Reverse map[int16]string
}

type Stash struct {
	SlotID    string
	Container Map
}

type Map struct {
	Height  int8
	Width   int8
	Map     []string
	FlatMap map[string]FlatMapLookup
}

type FlatMapLookup struct {
	Width   int8
	Height  int8
	Rotated bool
	StartX  int16
	EndX    int16
	StartY  int16
	EndY    int16
}

func SetInventoryContainer(inventory *Inventory) *InventoryContainer {
	output := &InventoryContainer{}

	output.SetInventoryIndex(inventory)
	output.SetInventoryStash(inventory)

	return output
}

func (ic *InventoryContainer) SetInventoryStash(inventory *Inventory) {
	var stash *Stash
	if ic.Stash == nil {
		ic.Stash = &Stash{}
		stash = ic.Stash

		item := GetItemByUID(inventory.Items[ic.Lookup.Forward[inventory.Stash]].TPL)
		grids := item.GetItemGrids()

		for key, value := range grids {
			stash.SlotID = key

			height := value.Props.CellsV
			width := value.Props.CellsH

			arraySize := int(height) * int(width)

			stash.Container = Map{
				Height:  height,
				Width:   width,
				Map:     make([]string, arraySize),
				FlatMap: make(map[string]FlatMapLookup),
			}
		}
	} else {
		stash = ic.Stash
		stash.Container.Map = stash.Container.Map[0:]
		stash.Container.FlatMap = make(map[string]FlatMapLookup)
	}

	var containerMap = &stash.Container.Map
	var containerFlatMap = &stash.Container.FlatMap
	var stride = int16(stash.Container.Width)
	var itemID string

	for index := range ic.Lookup.Reverse {
		itemInInventory := inventory.Items[index]
		if itemInInventory.ParentID == nil ||
			*itemInInventory.ParentID != inventory.Stash ||
			itemInInventory.SlotID == nil ||
			*itemInInventory.SlotID != "hideout" ||
			itemInInventory.Location == nil {
			continue
		}

		itemFlatMap := FlatMapLookup{}

		height, width := ic.GetSizeInInventory(inventory.Items, itemInInventory.ID)
		if height == -1 && width == -1 {
			continue
		}

		// TODO: See if this would be better off in GetSizeInInventory() function
		// answer is NO
		if width != 0 {
			width--
		}
		if height != 0 {
			height--
		}

		if itemInInventory.Location.R.(float64) == 1 {
			itemFlatMap.Height = width
			itemFlatMap.Width = height
		} else {
			itemFlatMap.Height = height
			itemFlatMap.Width = width
		}

		row := int16(itemInInventory.Location.Y.(float64)) * stride
		itemFlatMap.StartX = int16(itemInInventory.Location.X.(float64)) + row
		itemFlatMap.EndX = itemFlatMap.StartX + int16(itemFlatMap.Width)

		(*containerFlatMap)[itemInInventory.ID] = itemFlatMap

		if itemFlatMap.Height == 0 && itemFlatMap.Width == 0 {
			if itemID = (*containerMap)[itemFlatMap.StartX]; itemID != "" {
				log.Fatalln("Flat Map Index of", itemFlatMap.StartX, "is trying to be filled by", itemInInventory.ID, "but is occupied by", stash.Container.Map[itemFlatMap.StartX])
			}

			(*containerMap)[itemFlatMap.StartX] = itemInInventory.ID
			continue
		}

		for column := itemFlatMap.StartX; column <= itemFlatMap.EndX; column++ {
			if itemID = (*containerMap)[column]; itemID != "" {
				log.Fatalln("Flat Map Index of X position", column, "is trying to be filled by", itemInInventory.ID, "but is occupied by", stash.Container.Map[column])
			}
			(*containerMap)[column] = itemInInventory.ID

			for row := int16(1); row <= int16(itemFlatMap.Height); row++ {
				var coordinate = row*stride + column
				if itemID = (*containerMap)[coordinate]; itemID != "" {
					log.Fatalln("Flat Map Index of Y position", row, "is trying to be filled by", itemInInventory.ID, "but is occupied by", stash.Container.Map[coordinate])
				}
				(*containerMap)[coordinate] = itemInInventory.ID
			}
		}
	}
	// TODO: Remove this
	_ = tools.WriteToFile("/1darray.json", stash.Container.Map)
}

// ClearItemFromContainer wipes item, based on the UID, from the Lookup, Map and FlatMap
func (ic *InventoryContainer) ClearItemFromContainer(UID string) {
	var stash = *ic.Stash
	var itemFlatMap = stash.Container.FlatMap[UID]
	var containerMap = &stash.Container.Map

	var stride = int16(stash.Container.Width)
	var itemID string

	for column := itemFlatMap.StartX; column <= itemFlatMap.EndX; column++ {
		itemID = (*containerMap)[column]
		if itemID != UID {
			log.Fatalln("Flat Map Index of X position", column, "is trying to be emptied of", UID, "but is occupied by", stash.Container.Map[column])
		}
		(*containerMap)[column] = ""

		for row := int16(1); row <= int16(itemFlatMap.Height); row++ {
			var coordinate = row*stride + column
			itemID = (*containerMap)[coordinate]
			if itemID != UID {
				log.Fatalln("Flat Map Index of Y position", row, "is trying to be emptied of", UID, "but is occupied by", stash.Container.Map[coordinate])
			}
			(*containerMap)[coordinate] = ""
		}
	}

	delete(ic.Lookup.Reverse, ic.Lookup.Forward[UID])
	delete(ic.Lookup.Forward, UID)
	delete(stash.Container.FlatMap, UID)
}

// TODO: Consider refactoring AddItemToContainer

// AddItemToContainer adds item, based on the UID, to the Lookup, Map and FlatMap
func (ic *InventoryContainer) AddItemToContainer(UID string, Inventory *Inventory) {
	ic.SetInventoryIndex(Inventory)

	var stash = *ic.Stash
	var itemFlatMap = new(FlatMapLookup)

	itemInInventory := Inventory.Items[*ic.GetIndexOfItemByUID(UID)]
	height, width := ic.GetSizeInInventory(Inventory.Items, UID)
	if height == -1 && width == -1 {
		log.Fatalln("Item", UID, "does not have an item size")
	}

	// TODO: See if this would be better off in GetSizeInInventory() function
	if width != 0 {
		width--
	}
	if height != 0 {
		height--
	}

	if itemInInventory.Location.R.(float64) == 1 {
		itemFlatMap.Height = width
		itemFlatMap.Width = height
	} else {
		itemFlatMap.Height = height
		itemFlatMap.Width = width
	}

	row := int16(itemInInventory.Location.Y.(float64)) * int16(ic.Stash.Container.Width)
	itemFlatMap.StartX = int16(itemInInventory.Location.X.(float64)) + row
	itemFlatMap.EndX = itemFlatMap.StartX + int16(itemFlatMap.Width)
	var containerMap = &stash.Container.Map

	var stride = int16(stash.Container.Width)
	var itemID string

	for column := itemFlatMap.StartX; column <= itemFlatMap.EndX; column++ {
		itemID = (*containerMap)[column]
		if itemID != "" {
			log.Fatalln("Flat Map Index of X position", column, "is trying to be filled by", UID, "but is occupied by", stash.Container.Map[column])
		}
		(*containerMap)[column] = UID

		for row := int16(1); row <= int16(itemFlatMap.Height); row++ {
			var coordinate = row*stride + column
			itemID = (*containerMap)[coordinate]
			if itemID != "" {
				log.Fatalln("Flat Map Index of Y position", row, "is trying to be filled by", UID, "but is occupied by", stash.Container.Map[coordinate])
			}
			(*containerMap)[coordinate] = UID
		}
	}

	stash.Container.FlatMap[UID] = *itemFlatMap
}

func GetInventoryItemFamilyTree(items []InventoryItem, parent string) []string {
	var list []string

	for _, childItem := range items {
		if childItem.ParentID == nil {
			continue
		}

		if *childItem.ParentID == parent {
			list = append(list, GetInventoryItemFamilyTree(items, childItem.ID)...)
		}
	}

	list = append(list, parent) // required
	return list
}

type sizes struct {
	ForcedUp    int8
	ForcedDown  int8
	ForcedRight int8
	ForcedLeft  int8
	SizeUp      int8
	SizeDown    int8
	SizeLeft    int8
	SizeRight   int8
}

func (ic *InventoryContainer) GetSizeInInventory(items []InventoryItem, parent string) (int8, int8) {
	family := GetInventoryItemFamilyTree(items, parent)
	length := len(family) - 1
	index := ic.Lookup.Forward[family[length]]
	UID := items[index].TPL

	height, width := GetItemByUID(UID).GetItemSize() //get parent as starting point

	if length == 1 {
		return height, width
	}

	var member string
	sizes := &sizes{}

	for i := 0; i < length; i++ {
		member = family[i]
		index = ic.Lookup.Forward[member]
		UID = items[index].TPL
		GetItemByUID(UID).GetItemForcedSize(sizes)
	}

	height += sizes.SizeUp + sizes.SizeDown + sizes.ForcedDown + sizes.ForcedUp
	width += sizes.SizeLeft + sizes.SizeRight + sizes.ForcedRight + sizes.ForcedLeft

	return height, width
}

func (ic *InventoryContainer) SetInventoryIndex(inventory *Inventory) {
	if ic.Lookup == nil {
		ic.Lookup = &Lookup{
			Forward: make(map[string]int16),
			Reverse: make(map[int16]string),
		}
	}

	var pos int16
	for idx, item := range inventory.Items {
		pos = int16(idx)

		ic.Lookup.Forward[item.ID] = pos
		ic.Lookup.Reverse[pos] = item.ID
	}
}

// GetIndexOfItemByUID retrieves cached index of the item in your
// Inventory by its UID in Lookup.Forward
func (ic *InventoryContainer) GetIndexOfItemByUID(UID string) *int16 {
	index, ok := ic.Lookup.Forward[UID]
	if !ok {
		fmt.Println("Item of UID", UID, "does not exist in cache. Returning -1")
		return nil
	}
	return &index
}
