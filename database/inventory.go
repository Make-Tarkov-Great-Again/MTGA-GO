package database

import (
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
	Stash  Stash
	Lookup Lookup
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
	ic.Stash = Stash{}
	stash := &ic.Stash

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

	stride := int16(stash.Container.Width)
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

		stash.Container.FlatMap[itemInInventory.ID] = itemFlatMap

		if height == 0 && width == 0 {
			if stash.Container.Map[itemFlatMap.StartX] != "" {
				log.Fatalln("Flat Map Index of", itemFlatMap.StartX, "is trying to be filled by", itemInInventory.ID, "but is occupied by", stash.Container.Map[itemFlatMap.StartX])
			}

			stash.Container.Map[itemFlatMap.StartX] = itemInInventory.ID
			continue
		}

		for column := itemFlatMap.StartX; column <= itemFlatMap.EndX; column++ {
			if stash.Container.Map[column] != "" {
				log.Fatalln("Flat Map Index of X position", column, "is trying to be filled by", itemInInventory.ID, "but is occupied by", stash.Container.Map[column])
			}
			stash.Container.Map[column] = itemInInventory.ID

			for row := int16(1); row <= int16(itemFlatMap.Height); row++ {
				coordinate := row*stride + column
				if stash.Container.Map[coordinate] != "" {
					log.Fatalln("Flat Map Index of Y position", row, "is trying to be filled by", itemInInventory.ID, "but is occupied by", stash.Container.Map[coordinate])
				}
				stash.Container.Map[coordinate] = itemInInventory.ID
			}
		}
	}
	// TODO: Remove this
	//_ = tools.WriteToFile("/1darray.json", stash.Container.Map)
}

// ClearItemFromStash wipes item, based on the UID, from the container cached map, flatmap, and lookups
func (ic *InventoryContainer) ClearItemFromStash(UID string) {
	stash := &ic.Stash
	itemFlatMap := stash.Container.FlatMap[UID]
	stride := int16(stash.Container.Width)

	for column := itemFlatMap.StartX; column <= itemFlatMap.EndX; column++ {
		if stash.Container.Map[column] != UID {
			log.Fatalln("Flat Map Index of X position", column, "is trying to be emptied of", UID, "but is occupied by", stash.Container.Map[column])
		}
		stash.Container.Map[column] = ""

		for row := int16(1); row <= int16(itemFlatMap.Height); row++ {
			coordinate := row*stride + column
			if stash.Container.Map[coordinate] != UID {
				log.Fatalln("Flat Map Index of Y position", row, "is trying to be emptied of", UID, "but is occupied by", stash.Container.Map[coordinate])
			}
			stash.Container.Map[coordinate] = ""
		}
	}

	delete(ic.Lookup.Reverse, ic.Lookup.Forward[UID])
	delete(ic.Lookup.Forward, UID)
	delete(stash.Container.FlatMap, UID)
}

func GetInventoryItemFamilyTree(items []InventoryItem, parent string) []string {
	var list []string

	for _, childitem := range items {
		if childitem.ParentID == nil {
			continue
		}

		if *childitem.ParentID == parent {
			list = append(list, GetInventoryItemFamilyTree(items, childitem.ID)...)
		}
	}

	list = append(list, parent) // required
	return list
}

func (ic *InventoryContainer) GetSizeInInventory(items []InventoryItem, parent string) (int8, int8) {
	family := GetInventoryItemFamilyTree(items, parent)
	length := len(family)
	index := ic.Lookup.Forward[family[length-1]]
	itemID := items[index].TPL

	height, width := GetItemByUID(itemID).GetItemSize() //get parent as starting point

	if length == 1 {
		return height, width
	}

	for i := 1; i < length-1; i++ {
		index = ic.Lookup.Forward[family[i]]
		itemID = items[index].TPL
		forcedHeight, forcedWidth := GetItemByUID(itemID).GetItemForcedSize()
		height += forcedHeight
		width += forcedWidth
	}

	return height, width
}

func (ic *InventoryContainer) SetInventoryIndex(inventory *Inventory) {
	ic.Lookup = Lookup{
		Forward: make(map[string]int16),
		Reverse: make(map[int16]string),
	}

	index := ic.Lookup
	for idx, item := range inventory.Items {
		pos := int16(idx)

		index.Forward[item.ID] = pos
		index.Reverse[pos] = item.ID
	}
}
