package data

import (
	"log"
	"slices"

	"MT-GO/tools"
	"github.com/goccy/go-json"
)

// RemoveSingleItemFromInventoryByIndex takes the existing Inventory.Items and removes an InventoryItem at its index
// by shifting the indexes to the left
func (inv *Inventory) RemoveSingleItemFromInventoryByIndex(index int16) {
	if index < 0 || index >= int16(len(inv.Items)) {
		log.Println("[RemoveSingleItemFromInventoryByIndex] Index out of Range")
		return
	}

	copy(inv.Items[index:], inv.Items[index+1:])
	inv.Items = inv.Items[:len(inv.Items)-1]
}

// RemoveItemsFromInventoryByIndices takes the existing Inventory.Items and removes an InventoryItem at its index
// by creating new slice to assign to Inventory.Items
func (inv *Inventory) RemoveItemsFromInventoryByIndices(indices []int16) {
	output := make([]InventoryItem, 0, len(inv.Items))

	for idx, item := range inv.Items {
		if slices.Contains(indices, int16(idx)) {
			continue
		}
		output = append(output, item)
	}
	inv.Items = output
}

func CreateNewItem(TPL string, parent string) *InventoryItem {
	item := new(InventoryItem)

	item.ID = tools.GenerateMongoID()
	item.ParentID = parent
	item.TPL = TPL

	return item
}

// ConvertAssortItemsToInventoryItem converts AssortItem to InventoryItem, also reassigns IDs of all items
// as well as their children; sets parent item to last index
func ConvertAssortItemsToInventoryItem(assortItems []*AssortItem, stashID *string) []InventoryItem {
	convertedIDs := make(map[string]string)
	var parent InventoryItem

	input := make([]InventoryItem, 0, len(assortItems))
	for _, assortItem := range assortItems {
		data, err := json.Marshal(assortItem)
		if err != nil {
			log.Println("Failed to marshal Assort Item, returning empty output")
			return input
		}

		inventoryItem := new(InventoryItem)
		err = json.UnmarshalNoEscape(data, inventoryItem)
		if err != nil {
			log.Println("Failed to unmarshal Assort Item to Inventory Item, returning empty output")
			return input
		}

		newID := tools.GenerateMongoID()
		convertedIDs[inventoryItem.ID] = newID
		inventoryItem.ID = newID

		if inventoryItem.SlotID == "hideout" && inventoryItem.ParentID == "hideout" {
			inventoryItem.ParentID = *stashID

			inventoryItem.UPD.BuyRestrictionMax = 0
			inventoryItem.UPD.StackObjectsCount = 0
			inventoryItem.UPD.UnlimitedCount = false

			parent = *inventoryItem
			continue
		}

		input = append(input, *inventoryItem)
	}

	input = append(input, parent)

	output := make([]InventoryItem, 0, len(assortItems))
	for _, item := range input {
		if CID, ok := convertedIDs[item.ParentID]; !ok {
			continue
		} else {
			item.ParentID = CID
			output = append(output, item)
		}
	}

	output = append(output, parent)
	return output
}

func AssignNewIDs(inventoryItems []InventoryItem) []InventoryItem {
	input := make([]InventoryItem, 0, len(inventoryItems))
	convertedIDs := make(map[string]string)

	for _, inventoryItem := range inventoryItems {
		newID := tools.GenerateMongoID()
		convertedIDs[inventoryItem.ID] = newID
		inventoryItem.ID = newID

		input = append(input, inventoryItem)
	}

	output := make([]InventoryItem, 0, len(inventoryItems))
	var parent InventoryItem
	for _, item := range input {
		if item.SlotID == "hideout" {
			parent = item
		}

		CID, ok := convertedIDs[item.ParentID]
		if !ok {
			continue
		}
		item.ParentID = CID
		output = append(output, item)
	}
	output = append(output, parent)
	return output
}

func GetInventoryItemFamilyTreeIDs(items []InventoryItem, parent string) []string {
	var list []string

	for _, childItem := range items {
		if childItem.ParentID == "" {
			continue
		}

		if childItem.ParentID == parent {
			list = append(list, GetInventoryItemFamilyTreeIDs(items, childItem.ID)...)
		}
	}

	list = append(list, parent) // required
	return list
}

// MeasurePurchaseForInventoryMapping is the same as MeasureItemForInventoryMapping except it's exclusively
// used for Trader/RagFair purchases; returns correct height and width based on items given
func MeasurePurchaseForInventoryMapping(items []InventoryItem) (int8, int8) {
	parentItem := items[len(items)-1]
	itemInDatabase, err := GetItemByID(parentItem.TPL) //parent
	if err != nil {
		log.Println(err)
		return -1, -1
	}

	height, width := itemInDatabase.GetItemSize()             //get parent as starting point
	if itemInDatabase.Parent == "5448e53e4bdc2d60728b4567" || //backpack
		itemInDatabase.Parent == "566168634bdc2d144c8b456c" || //searchableItem
		itemInDatabase.Parent == "5795f317245977243854e041" { //simpleContainer
		return height, width
	}

	var parentFolded = parentItem.UPD != nil && parentItem.UPD.Foldable != nil && parentItem.UPD.Foldable.Folded

	canFold, foldablePropertyExists := itemInDatabase.Props["Foldable"].(bool)
	foldedSlotID, foldedSlotPropertyExists := itemInDatabase.Props["FoldedSlot"].(string)

	if (foldablePropertyExists && canFold) && foldedSlotPropertyExists && parentFolded {
		sizeReduceRight, ok := itemInDatabase.Props["SizeReduceRight"].(float64)
		if !ok {
			log.Println("Could not type assert itemInDatabase.Props.SizeReduceRight of UID", parentItem.ID)
			return -1, -1
		}
		width -= int8(sizeReduceRight)
	}

	if len(items) == 1 {
		return height, width
	}

	sizes := &sizes{}

	var childFolded bool
	for _, item := range items {
		childFolded = item.UPD != nil && item.UPD.Foldable != nil && item.UPD.Foldable.Folded
		if parentFolded || childFolded {
			continue
		} else if (foldablePropertyExists && canFold) &&
			item.SlotID == foldedSlotID &&
			(parentFolded || childFolded) {
			continue
		}

		item, err := GetItemByID(item.TPL)
		if err != nil {
			log.Println(err)
			return -1, -1
		}
		item.GetItemForcedSize(sizes)
	}

	height += sizes.SizeUp + sizes.SizeDown + sizes.ForcedDown + sizes.ForcedUp
	width += sizes.SizeLeft + sizes.SizeRight + sizes.ForcedRight + sizes.ForcedLeft

	return height, width
}

type Inventory struct {
	Items              []InventoryItem   `json:"items"`
	Equipment          string            `json:"equipment"`
	Stash              string            `json:"stash,omitempty"`
	SortingTable       string            `json:"sortingTable,omitempty"`
	QuestRaidItems     string            `json:"questRaidItems,omitempty"`
	QuestStashItems    string            `json:"questStashItems,omitempty"`
	FastPanel          map[string]string `json:"fastPanel"`
	HideoutAreaStashes map[string]string `json:"hideoutAreaStashes"`
}

type InventoryItem struct {
	ID       string      `json:"_id"`
	TPL      string      `json:"_tpl"`
	ParentID string      `json:"parentId,omitempty"`
	SlotID   string      `json:"slotId,omitempty"`
	Location any         `json:"location,omitempty"` // this can also be an int, wow
	UPD      *ItemUpdate `json:"upd,omitempty"`
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
	Tag               *Tag        `json:"Tag,omitempty"`
	Togglable         *Toggle     `json:"Togglable,omitempty"`
}

type ItemUpdate struct {
	StackObjectsCount     int32            `json:"StackObjectsCount,omitempty"`
	FireMode              *FireMode        `json:"FireMode,omitempty"`
	Foldable              *Foldable        `json:"Foldable,omitempty"`
	Repairable            *Repairable      `json:"Repairable,omitempty"`
	Sight                 *Sight           `json:"Sight,omitempty"`
	MedKit                *MedicalKit      `json:"MedKit,omitempty"`
	FoodDrink             *FoodDrink       `json:"FoodDrink,omitempty"`
	RepairKit             *RepairKit       `json:"RepairKit,omitempty"`
	Light                 *Light           `json:"Light,omitempty"`
	Resource              *Resource        `json:"Resource,omitempty"`
	Tag                   *Tag             `json:"Tag,omitempty"`
	Togglable             *Toggle          `json:"Togglable,omitempty"`
	RecodableComponent    *RecodeComponent `json:"RecodableComponent,omitempty"`
	BuyRestrictionCurrent int16            `json:"BuyRestrictionCurrent,omitempty"`
	BuyRestrictionMax     int16            `json:"BuyRestrictionMax,omitempty"`
	UnlimitedCount        bool             `json:"UnlimitedCount,omitempty"`
}

type RecodeComponent struct {
	IsEncoded bool `json:"IsEncoded"`
}

type Toggle struct {
	On bool `json:"On"`
}

type Tag struct {
	Name  string
	Color string
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
	IsSearched bool `json:"isSearched"`
	R          any  `json:"r"`
	X          any  `json:"x"`
	Y          any  `json:"y"`
}
