package database

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"
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
	Width       int8
	Height      int8
	StartX      int16
	EndX        int16
	Coordinates []int16
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

		itemFlatMap := *ic.CreateFlatMapLookup(&inventory.Items, &itemInInventory)

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
}

func (ic *InventoryContainer) CreateFlatMapLookup(inventoryItems *[]InventoryItem, itemInInventory *InventoryItem) *FlatMapLookup {
	output := &FlatMapLookup{}

	height, width := ic.MeasureItemForInventoryMapping(*inventoryItems, itemInInventory.ID)
	if height == -1 && width == -1 {
		return nil
	}

	if width != 0 {
		width--
	}
	if height != 0 {
		height--
	}

	if itemInInventory.Location.R.(float64) == 1 {
		output.Height = width
		output.Width = height
	} else {
		output.Height = height
		output.Width = width
	}

	row := int16(itemInInventory.Location.Y.(float64)) * int16(ic.Stash.Container.Width)
	output.StartX = int16(itemInInventory.Location.X.(float64)) + row
	output.EndX = output.StartX + int16(output.Width)

	return output
}

//TODO: Check consistency of this code
// i might need to check if oldX = newX then check heights

// ResetItemSizeInContainer resets item size in InventoryContainer to reflect item size change
func (ic *InventoryContainer) ResetItemSizeInContainer(itemInInventory *InventoryItem, Inventory *Inventory) {
	var stash = *ic.Stash
	var itemFlatMap = stash.Container.FlatMap[itemInInventory.ID]
	var containerMap = &stash.Container.Map
	var stride = int16(stash.Container.Width)

	newItemFlatMap := *ic.CreateFlatMapLookup(&Inventory.Items, itemInInventory)

	if newItemFlatMap.EndX < itemFlatMap.EndX {
		for column := newItemFlatMap.EndX + 1; column <= itemFlatMap.EndX; column++ {
			(*containerMap)[column] = ""

			if newItemFlatMap.Height < itemFlatMap.Height {
				for row := int16(newItemFlatMap.Height) + 1; row <= int16(itemFlatMap.Height); row++ {
					var coordinate = row*stride + itemFlatMap.EndX
					(*containerMap)[coordinate] = ""
				}
			}
		}
	} else if newItemFlatMap.EndX > itemFlatMap.EndX {
		for column := itemFlatMap.EndX + 1; column <= newItemFlatMap.EndX; column++ {
			(*containerMap)[column] = itemInInventory.ID

			if newItemFlatMap.Height > itemFlatMap.Height {
				for row := int16(itemFlatMap.Height) + 1; row <= int16(newItemFlatMap.Height); row++ {
					var coordinate = row*stride + itemFlatMap.EndX
					(*containerMap)[coordinate] = itemInInventory.ID
				}
			}
		}
	}

	stash.Container.FlatMap[itemInInventory.ID] = newItemFlatMap
	ic.SetInventoryIndex(Inventory)
}

// ClearItemFromContainer wipes item, based on the UID, from the cached InventoryContainer
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

func (ic *InventoryContainer) GetValidPositionForItem(itemInInventory *InventoryItem, Inventory *Inventory) []int16 {
	var output []int16

	//var itemID string
	//var containerMap = &ic.Stash.Container.Map
	//var containerFlatMap = &ic.Stash.Container.FlatMap
	//var stride = int16(ic.Stash.Container.Width)

	height, width := ic.MeasureItemForInventoryMapping(Inventory.Items, itemInInventory.ID)
	if height == -1 && width == -1 {
		return nil
	}

	if width != 0 {
		width--
	}
	if height != 0 {
		height--
	}

	return output
}

func ConvertAssortItemsToInventoryItem(assortItems []*AssortItem) []*InventoryItem {
	output := make([]*InventoryItem, 0, len(assortItems))
	for _, assortItem := range assortItems {
		data, err := json.Marshal(assortItem)
		if err != nil {
			log.Println("Failed to marshal Assort Item, returning empty output")
			return output
		}

		inventoryItem := new(InventoryItem)
		err = json.Unmarshal(data, inventoryItem)
		if err != nil {
			log.Println("Failed to unmarshal Assort Item to Inventory Item, returning empty output")
			return output
		}
		output = append(output, inventoryItem)
	}
	return output
}

// TODO: Consider refactoring AddItemToContainer

// AddItemToContainer adds item, based on the UID, to the cached InventoryContainer
func (ic *InventoryContainer) AddItemToContainer(UID string, Inventory *Inventory) {
	ic.SetInventoryIndex(Inventory)

	var stash = *ic.Stash
	var itemFlatMap = ic.CreateFlatMapLookup(&Inventory.Items, &Inventory.Items[*ic.GetIndexOfItemByUID(UID)])
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

func GetInventoryItemFamilyTreeIDs(items []InventoryItem, parent string) []string {
	var list []string

	for _, childItem := range items {
		if childItem.ParentID == nil {
			continue
		}

		if *childItem.ParentID == parent {
			list = append(list, GetInventoryItemFamilyTreeIDs(items, childItem.ID)...)
		}
	}

	list = append(list, parent) // required
	return list
}

func GetAllChildItemsInInventory(items []*InventoryItem, parentID string) []InventoryItem {

	// Create a map to store parentIDs and their child objects
	parentReference := make(map[string][]InventoryItem)

	// Loop through inventory items to populate the parentReference map
	for _, thisItem := range items {
		if thisItem.ID != "" && thisItem.ParentID != nil {
			if _, exists := parentReference[*thisItem.ParentID]; !exists {
				parentReference[*thisItem.ParentID] = []InventoryItem{*thisItem}
			} else {
				parentReference[*thisItem.ParentID] = append(parentReference[*thisItem.ParentID], *thisItem)
			}
		}
	}

	if children, exists := parentReference[parentID]; !exists {
		return nil
	} else {
		returnArray := make([]InventoryItem, len(children))
		copy(returnArray, children)

		for _, child := range returnArray {
			if grandchildren, exists := parentReference[child.ID]; exists {
				returnArray = append(returnArray, grandchildren...)
			}
		}

		return returnArray
	}
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

// MeasureItemForInventoryMapping gets the correct size of an item within the Character.Inventory for setting in
// Stash.Container
func (ic *InventoryContainer) MeasureItemForInventoryMapping(items []InventoryItem, parent string) (int8, int8) {
	index := ic.Lookup.Forward[parent]
	itemInInventory := items[index]

	itemInDatabase := GetItemByUID(itemInInventory.TPL) //parent
	height, width := itemInDatabase.GetItemSize()       //get parent as starting point

	if itemInDatabase.Parent == "5448e53e4bdc2d60728b4567" || //backpack
		itemInDatabase.Parent == "566168634bdc2d144c8b456c" || //searchableItem
		itemInDatabase.Parent == "5795f317245977243854e041" { //simpleContainer
		return height, width
	}

	var parentFolded = itemInInventory.UPD != nil && itemInInventory.UPD.Foldable != nil && itemInInventory.UPD.Foldable.Folded

	canFold, foldablePropertyExists := itemInDatabase.Props["Foldable"].(bool)
	foldedSlotID, foldedSlotPropertyExists := itemInDatabase.Props["FoldedSlot"].(string)

	if (foldablePropertyExists && canFold) && foldedSlotPropertyExists && parentFolded {
		sizeReduceRight, ok := itemInDatabase.Props["SizeReduceRight"].(float64)
		if !ok {
			log.Fatalln("Could not type assert itemInDatabase.Props.SizeReduceRight of UID", itemInInventory.ID)
		}
		width -= int8(sizeReduceRight)
	}

	family := GetInventoryItemFamilyTreeIDs(items, parent)
	length := len(family) - 1

	if length == 1 {
		return height, width
	}

	var member string
	sizes := &sizes{}

	var childFolded bool
	for i := 0; i < length; i++ {
		member = family[i]
		index = ic.Lookup.Forward[member]
		itemInInventory = items[index]

		childFolded = itemInInventory.UPD != nil && itemInInventory.UPD.Foldable != nil && itemInInventory.UPD.Foldable.Folded
		if parentFolded || childFolded {
			continue
		} else if (foldablePropertyExists && canFold) &&
			*itemInInventory.SlotID == foldedSlotID &&
			(parentFolded || childFolded) {
			continue
		}

		GetItemByUID(itemInInventory.TPL).GetItemForcedSize(sizes)
	}

	height += sizes.SizeUp + sizes.SizeDown + sizes.ForcedDown + sizes.ForcedUp
	width += sizes.SizeLeft + sizes.SizeRight + sizes.ForcedRight + sizes.ForcedLeft

	return height, width
}

// SetInventoryIndex set/reset InventoryContainer.Lookup for fast Inventory.Items lookup
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
