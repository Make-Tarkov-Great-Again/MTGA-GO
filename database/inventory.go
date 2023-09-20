package database

import (
	"fmt"
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
	ID       string                  `json:"_id"`
	TPL      string                  `json:"_tpl"`
	Location *InventoryItemLocation  `json:"location,omitempty"`
	ParentID *string                 `json:"parentId,omitempty"`
	SlotID   *string                 `json:"slotId,omitempty"`
	UPD      *map[string]interface{} `json:"upd,omitempty"`
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

	for index := range ic.Lookup.Reverse {
		itemInInventory := inventory.Items[index]
		if itemInInventory.SlotID == nil || *itemInInventory.SlotID != "hideout" || itemInInventory.Location == nil {
			continue
		}

		itemFlatMap := FlatMapLookup{}
		height, width := ic.GetSizeInInventory(inventory.Items, itemInInventory.ID)
		itemFlatMap.Height = height
		itemFlatMap.Width = width

		stash.Container.FlatMap[itemInInventory.ID] = itemFlatMap

		//fmt.Println(height, width)
	}

	fmt.Println(grids)
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

	height, width := GetItemByUID(items[ic.Lookup.Forward[family[0]]].TPL).GetItemSize() //get parent as starting point
	length := len(family)
	if length == 1 {
		return height, width
	}

	for i := 1; i < len(family); i++ {
		forcedHeight, forcedWidth := GetItemByUID(items[ic.Lookup.Forward[family[i]]].TPL).GetItemForcedSize()
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

func (c *Character) GetInventoryIndex(container *InventoryContainer) {}

func (c *Character) SetInventoryContainer(container *InventoryContainer) {}
