package database

import (
	"MT-GO/services"
	"github.com/goccy/go-json"
	"log"
	"slices"
)

type transfer struct {
	Action string
	Item   string `json:"item"`
	With   string `json:"with"`
	Count  int32  `json:"count"`
}

func (c *Character) TransferItem(moveAction map[string]any) {
	transfer := new(transfer)
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &transfer); err != nil {
		log.Println(err)
		return
	}

	inventoryCache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	toMergeIndex := *inventoryCache.GetIndexOfItemByID(transfer.Item)
	toMerge := &c.Inventory.Items[toMergeIndex]

	mergeWithIndex := *inventoryCache.GetIndexOfItemByID(transfer.With)
	mergeWith := &c.Inventory.Items[mergeWithIndex]

	toMerge.UPD.StackObjectsCount -= transfer.Count
	mergeWith.UPD.StackObjectsCount += transfer.Count
}

type split struct {
	Action    string `json:"Action"`
	SplitItem string `json:"splitItem"`
	NewItem   string `json:"newItem"`
	Container moveTo `json:"container"`
	Count     int32  `json:"count"`
}

func (c *Character) SplitItem(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	split := new(split)
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &split); err != nil {
		log.Println(err)
		return
	}

	invCache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	originalItem := &c.Inventory.Items[*invCache.GetIndexOfItemByID(split.SplitItem)]
	originalItem.UPD.StackObjectsCount -= split.Count

	newItem := &InventoryItem{
		ID:       split.NewItem,
		TPL:      originalItem.TPL,
		UPD:      originalItem.UPD,
		ParentID: split.Container.ID,
		SlotID:   split.Container.Container,
	}

	newItem.UPD.StackObjectsCount = split.Count

	if split.Container.Location != nil {
		newItem.Location = &InventoryItemLocation{
			IsSearched: split.Container.Location.IsSearched,
			X:          split.Container.Location.X,
			Y:          split.Container.Location.Y,
		}
		if split.Container.Location.R == "Vertical" {
			newItem.Location.R = float64(1)
		} else {
			newItem.Location.R = float64(0)
		}

		height, width := MeasurePurchaseForInventoryMapping([]InventoryItem{*newItem})
		itemFlatMap := invCache.CreateFlatMapLookup(height, width, newItem)
		itemFlatMap.Coordinates = invCache.GenerateCoordinatesFromLocation(*itemFlatMap)
		invCache.AddItemToContainer(split.NewItem, itemFlatMap)
	}

	c.Inventory.Items = append(c.Inventory.Items, *newItem)
	invCache.SetSingleInventoryIndex(newItem.ID, int16(len(c.Inventory.Items)-1))

	profileChangesEvent.ProfileChanges[c.ID].Items.Change = append(profileChangesEvent.ProfileChanges[c.ID].Items.Change, *originalItem)
	profileChangesEvent.ProfileChanges[c.ID].Items.New = append(profileChangesEvent.ProfileChanges[c.ID].Items.New, InventoryItem{ID: newItem.ID, TPL: newItem.TPL, UPD: newItem.UPD})
}

type remove struct {
	Action string `json:"Action"`
	ItemId string `json:"item"`
}

func (c *Character) RemoveItem(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	remove := new(remove)
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &remove); err != nil {
		log.Println(err)
		return
	}

	inventoryCache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}
	itemChildren := GetInventoryItemFamilyTreeIDs(c.Inventory.Items, remove.ItemId)

	var itemIndex int16
	toDelete := make([]int16, 0, len(itemChildren))
	for _, itemID := range itemChildren {
		itemIndex = *inventoryCache.GetIndexOfItemByID(itemID)
		toDelete = append(toDelete, itemIndex)
	}

	inventoryCache.ClearItemFromContainer(remove.ItemId)
	c.Inventory.RemoveItemsFromInventoryByIndices(toDelete)
	inventoryCache.SetInventoryIndex(&c.Inventory)

	profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, InventoryItem{ID: remove.ItemId})

}

type applyInventoryChanges struct {
	Action       string
	ChangedItems []any `json:"changedItems"`
}

//TODO: Make ApplyInventoryChanges not look like shit

func (c *Character) ApplyInventoryChanges(moveAction map[string]any) {
	applyInventoryChanges := new(applyInventoryChanges)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &applyInventoryChanges)
	if err != nil {
		log.Println(err)
		return
	}

	cache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range applyInventoryChanges.ChangedItems {
		properties, ok := item.(map[string]any)
		if !ok {
			log.Println("Cannot type assert item from Auto-Sort items slice")
			return
		}

		UID, ok := properties["_id"].(string)
		if !ok {
			log.Println("Cannot type assert item `_id` property from Auto-Sort items slice")
			return
		}
		itemInInventory := &c.Inventory.Items[*cache.GetIndexOfItemByID(UID)]

		parent, ok := properties["parentId"].(string)
		if !ok {
			log.Println("Cannot type assert item `parentId` property from Auto-Sort items slice")
			return
		}
		itemInInventory.ParentID = parent

		slotId, ok := properties["slotId"].(string)
		if !ok {
			log.Println("Cannot type assert item `slotId` property from Auto-Sort items slice")
			return
		}
		itemInInventory.SlotID = slotId

		location, ok := properties["location"].(map[string]any)
		if !ok {
			itemInInventory.Location = nil
			continue
		} else {
			r, ok := location["r"].(string)
			if !ok {
				log.Println("Cannot type assert item.Location `r` property from Auto-Sort items slice")
				return
			}

			if r == "Horizontal" || r == "1" {
				itemInInventory.Location.R = float64(0)
			} else {
				itemInInventory.Location.R = float64(1)
			}

			if x, ok := location["r"].(float64); ok {
				itemInInventory.Location.X = x
			}

			if isSearched, ok := location["isSearched"].(bool); ok {
				itemInInventory.Location.IsSearched = isSearched
			}

			if y, ok := location["r"].(float64); ok {
				itemInInventory.Location.Y = y
			}
		}
	}
}

type buyFromTrader struct {
	Action      string
	Type        string          `json:"type"`
	TID         string          `json:"tid"`
	ItemID      string          `json:"item_id"`
	Count       int32           `json:"count"`
	SchemeID    int8            `json:"scheme_id"`
	SchemeItems []tradingScheme `json:"scheme_items"`
}

type tradingScheme struct {
	ID    string
	Count int32
}

type sellToTrader struct {
	Action string
	Type   string      `json:"type"`
	TID    string      `json:"tid"`
	Items  []soldItems `json:"items"`
	Price  int32       `json:"price"`
}

type soldItems struct {
	ID       string `json:"id"`
	Count    int32  `json:"count"`
	SchemeID int8   `json:"scheme_id"`
}

func (c *Character) TradingConfirm(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	invCache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	switch moveAction["type"] {
	case "buy_from_trader":
		buy := new(buyFromTrader)
		data, _ := json.Marshal(moveAction)
		err := json.Unmarshal(data, &buy)
		if err != nil {
			log.Println(err)
			return
		}

		c.BuyFromTrader(buy, invCache, profileChangesEvent)
	case "sell_to_trader":
		sell := new(sellToTrader)
		data, _ := json.Marshal(moveAction)
		err := json.Unmarshal(data, &sell)
		if err != nil {
			log.Println(err)
			return
		}

		c.SellToTrader(sell, invCache, profileChangesEvent)
	default:
		log.Println("YO! TRADINGCONFIRM.", moveAction["type"], "ISNT SUPPORTED YET HAHAHHAHAHAHAHAHHAHAHHHHHHHHHHHHHAHAHAHAHAHHAHA")
	}
}

func (c *Character) BuyFromTrader(tradeConfirm *buyFromTrader, invCache *InventoryContainer, profileChangesEvent *ProfileChangesEvent) {
	trader, err := GetTraderByUID(tradeConfirm.TID)
	if err != nil {
		log.Println(err)
		return
	}

	assortItem := trader.GetAssortItemByID(tradeConfirm.ItemID)
	if assortItem == nil {
		log.Println("Item of", tradeConfirm.ItemID, "does not exist in trader assort, killing!")
		return
	}

	inventoryItems := ConvertAssortItemsToInventoryItem(assortItem, &c.Inventory.Stash)
	if len(inventoryItems) == 0 {
		log.Println("Converting Assort Item to Inventory Item failed, killing")
		return
	}

	var stackMaxSize = GetItemByID(inventoryItems[len(inventoryItems)-1].TPL).GetStackMaxSize()
	var stackSlice = services.GetCorrectAmountOfItemsPurchased(tradeConfirm.Count, stackMaxSize)
	// Basically gets the correct amount of items to be created, based on StackSize

	//Create copy-of Character.Inventory.Items for modification in the case of any failures to assign later
	copyOfItems := make([]InventoryItem, 0, len(c.Inventory.Items)+(len(inventoryItems)*len(stackSlice)))
	copyOfItems = append(copyOfItems, c.Inventory.Items...)
	//Create copy-of invCache.Stash.Container for modification in the case of any failures to assign later
	copyOfMap := invCache.Stash.Container

	toAdd := make([]InventoryItem, 0, len(stackSlice))
	toDelete := make(map[string]int16)
	traderRelations := c.TradersInfo[tradeConfirm.TID]

	height, width := MeasurePurchaseForInventoryMapping(inventoryItems)

	for _, stack := range stackSlice {
		var copyOfInventoryItems []InventoryItem
		if len(stackSlice) >= 1 {
			copyOfInventoryItems = AssignNewIDs(inventoryItems)
		} else {
			copyOfInventoryItems = inventoryItems
		}

		mainItem := &copyOfInventoryItems[len(copyOfInventoryItems)-1]

		validLocation := invCache.GetValidLocationForItem(height, width)
		if validLocation == nil {
			log.Println("Item", tradeConfirm.ItemID, "was not purchased because we could not find a position in your inventory!!")
			invCache.Stash.Container = copyOfMap //if failure, assign old map
			return
		}

		if stackMaxSize > 1 {
			mainItem.UPD.StackObjectsCount = stack
		}
		mainItem.Location = &InventoryItemLocation{
			IsSearched: true,
			R:          float64(0),
			X:          float64(validLocation.X),
			Y:          float64(validLocation.Y),
		}

		itemFlatMap := invCache.CreateFlatMapLookup(height, width, mainItem)
		itemFlatMap.Coordinates = validLocation.MapInfo
		invCache.AddItemToContainer(mainItem.ID, itemFlatMap)

		toAdd = append(toAdd, copyOfInventoryItems...)
	}

	for _, scheme := range tradeConfirm.SchemeItems {
		index := invCache.GetIndexOfItemByID(scheme.ID)
		if index == nil {
			log.Println("Index of", scheme.ID, "does not exist in cache, killing!")
			return
		}

		itemInInventory := copyOfItems[*index]

		currency := *services.GetCurrencyByName(trader.Base.Currency)
		if services.IsCurrencyByID(itemInInventory.TPL) {
			traderRelations.SalesSum += float32(scheme.Count)
		} else {
			priceOfItem, err := GetPriceByID(itemInInventory.TPL)
			if err != nil {
				log.Println(err)
				return
			}

			if "RUB" != trader.Base.Currency {
				if conversion, err := ConvertFromRouble(*priceOfItem, currency); err == nil {
					traderRelations.SalesSum += float32(conversion)
				} else {
					log.Println(err)
					return
				}
			} else {
				traderRelations.SalesSum += float32(*priceOfItem)
			}
		}

		if itemInInventory.UPD != nil && itemInInventory.UPD.StackObjectsCount != 0 {
			var remainingBalance = scheme.Count

			if itemInInventory.UPD.StackObjectsCount > remainingBalance {
				itemInInventory.UPD.StackObjectsCount -= remainingBalance

				profileChangesEvent.ProfileChanges[c.ID].Items.Change = append(profileChangesEvent.ProfileChanges[c.ID].Items.Change, itemInInventory)
			} else if itemInInventory.UPD.StackObjectsCount == remainingBalance {
				toDelete[itemInInventory.ID] = *index
			} else {
				remainingBalance -= itemInInventory.UPD.StackObjectsCount

				toDelete[itemInInventory.ID] = *index

				//TODO: Consider creating a look-up cache for mergable Inventory.Items

				var toChange []InventoryItem
				for idx, item := range copyOfItems {
					if _, ok := toDelete[item.ID]; ok || item.TPL != currency {
						continue
					}

					change := item.UPD.StackObjectsCount - remainingBalance
					if change > 0 {
						remainingBalance -= item.UPD.StackObjectsCount
						toDelete[item.ID] = int16(idx)
						continue
					} else if change == 0 {
						remainingBalance -= item.UPD.StackObjectsCount
						toDelete[item.ID] = int16(idx)
						break
					}

					item.UPD.StackObjectsCount = change
					toChange = append(toChange, item)
					break
				}
				if remainingBalance > 0 {
					log.Println("Insufficient funds to purchase item, returning")
					invCache.Stash.Container = copyOfMap
					return
				}

				profileChangesEvent.ProfileChanges[c.ID].Items.Change = append(profileChangesEvent.ProfileChanges[c.ID].Items.Change, toChange...)
			}
		} else {
			toDelete[itemInInventory.ID] = *index
		}
	}

	// Add all items from toAdd to Copy of Inventory.Items
	if len(toAdd) == 0 {
		log.Println("balls")
		return
	}

	copyOfItems = append(copyOfItems, toAdd...)
	profileChangesEvent.ProfileChanges[c.ID].Items.New = append(profileChangesEvent.ProfileChanges[c.ID].Items.New, toAdd...)
	/*	for i := len(inventoryItems) - 1; i < len(toAdd); i += len(inventoryItems) {
		if toAdd[i].Location == nil && toAdd[i].SlotID != "hideout" {
			continue
		}
		profileChangesEvent.ProfileChanges[c.ID].Items.New = append(profileChangesEvent.ProfileChanges[c.ID].Items.New, toAdd[i])

	}*/

	//Assign copy-of Character.Inventory.Items to original Character.Inventory.Items
	c.Inventory.Items = copyOfItems

	if len(toDelete) != 0 {
		indices := make([]int16, 0, len(toDelete))
		for id, idx := range toDelete {
			invCache.ClearItemFromContainer(id)
			indices = append(indices, idx)

			profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, InventoryItem{ID: id})
		}
		c.Inventory.RemoveItemsFromInventoryByIndices(indices)
	}
	invCache.SetInventoryIndex(&c.Inventory)

	profileChangesEvent.ProfileChanges[c.ID].TraderRelations[tradeConfirm.TID] = traderRelations
	c.TradersInfo[tradeConfirm.TID] = traderRelations

	log.Println(len(stackSlice), "of Item", tradeConfirm.ItemID, "purchased!")
}

func (c *Character) SellToTrader(tradeConfirm *sellToTrader, invCache *InventoryContainer, profileChangesEvent *ProfileChangesEvent) {
	trader, err := GetTraderByUID(tradeConfirm.TID)
	if err != nil {
		log.Println(err)
		return
	}

	saleCurrency := *services.GetCurrencyByName(trader.Base.Currency)

	remainingBalance := tradeConfirm.Price
	stackMaxSize := GetItemByID(saleCurrency).GetStackMaxSize()

	cache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}
	copyOfMap := invCache.Stash.Container
	copyOfItems := make([]InventoryItem, 0, len(c.Inventory.Items))
	copyOfItems = append(copyOfItems, c.Inventory.Items...)

	toDelete := make(map[string]int16)
	for _, item := range tradeConfirm.Items {
		index := *cache.GetIndexOfItemByID(item.ID)
		toDelete[item.ID] = index
	}

	toChange := make([]InventoryItem, 0)
	for _, item := range copyOfItems {
		if remainingBalance == 0 {
			break
		}

		if item.TPL != saleCurrency || item.UPD.StackObjectsCount == stackMaxSize {
			continue
		}

		if item.UPD.StackObjectsCount+remainingBalance > stackMaxSize {
			remainingBalance -= stackMaxSize - item.UPD.StackObjectsCount
			item.UPD.StackObjectsCount = stackMaxSize

			toChange = append(toChange, item)
			continue
		} else {
			item.UPD.StackObjectsCount += remainingBalance
			remainingBalance = 0
			toChange = append(toChange, item)
			break
		}
	}

	if remainingBalance != 0 {
		var toAdd []InventoryItem

		//log.Println("If a new stack isn't made, we cry")

		stackSlice := services.GetCorrectAmountOfItemsPurchased(remainingBalance, stackMaxSize)
		item := []InventoryItem{*CreateNewItem(saleCurrency, c.Inventory.Stash)}
		// since it's one item, just get the height and width once
		height, width := MeasurePurchaseForInventoryMapping(item)

		for _, stack := range stackSlice {
			mainItem := AssignNewIDs(item)[0]

			validLocation := invCache.GetValidLocationForItem(height, width)
			if validLocation == nil {
				log.Println("Item", mainItem.ID, "was not created because we could not find a position in your inventory!")
				invCache.Stash.Container = copyOfMap //if failure, assign old map
				return
			}

			mainItem.UPD.StackObjectsCount = stack
			mainItem.Location = &InventoryItemLocation{
				IsSearched: true,
				R:          float64(0),
				X:          float64(validLocation.X),
				Y:          float64(validLocation.Y),
			}

			itemFlatMap := invCache.CreateFlatMapLookup(height, width, &mainItem)
			itemFlatMap.Coordinates = validLocation.MapInfo
			invCache.AddItemToContainer(mainItem.ID, itemFlatMap)

			toAdd = append(toAdd, mainItem)
		}

		copyOfItems = append(copyOfItems, toAdd...)
		profileChangesEvent.ProfileChanges[c.ID].Items.New = append(profileChangesEvent.ProfileChanges[c.ID].Items.New, toAdd...)
	}

	profileChangesEvent.ProfileChanges[c.ID].Items.Change = append(profileChangesEvent.ProfileChanges[c.ID].Items.Change, toChange...)
	c.Inventory.Items = copyOfItems

	if len(toDelete) != 0 {
		indices := make([]int16, 0, len(toDelete))
		for id, idx := range toDelete {
			invCache.ClearItemFromContainer(id)
			indices = append(indices, idx)
			if _, ok := toDelete[c.Inventory.Items[idx].ParentID]; ok {
				continue
			}
			profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, InventoryItem{ID: id})
		}
		c.Inventory.RemoveItemsFromInventoryByIndices(indices)
	}
	invCache.SetInventoryIndex(&c.Inventory)

	traderRelations := c.TradersInfo[tradeConfirm.TID]
	traderRelations.SalesSum += float32(tradeConfirm.Price)

	profileChangesEvent.ProfileChanges[c.ID].TraderRelations[tradeConfirm.TID] = traderRelations
	c.TradersInfo[tradeConfirm.TID] = traderRelations
}

type buyCustomization struct {
	Action string           `json:"Action"`
	Offer  string           `json:"offer"`
	Items  []map[string]any `json:"items"`
}

func (c *Character) CustomizationBuy(moveAction map[string]any) {
	customizationBuy := new(buyCustomization)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &customizationBuy)
	if err != nil {
		log.Println(err)
		return
	}

	trader, err := GetTraderByName("Ragman")
	if err != nil {
		log.Println(err)
		return
	}
	suitsIndex := trader.Index.Suits[customizationBuy.Offer]
	suitID := trader.Suits[suitsIndex].SuiteID

	storage, err := GetStorageByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	if !slices.Contains(storage.Suites, suitID) {
		//TODO: Pay for suite before appending to profile
		if len(customizationBuy.Items) == 0 {
			storage.Suites = append(storage.Suites, suitID)
			storage.SaveStorage(c.ID)
			return
		}
		log.Println("Cannot purchase clothing because I haven't implemented it yet lol")
		return
	}
	log.Println("Clothing was already purchased")
}

type wearCustomization struct {
	Action string   `json:"Action"`
	Suites []string `json:"suites"`
}

const (
	lowerParentID = "5cd944d01388ce000a659df9"
	upperParentID = "5cd944ca1388ce03a44dc2a4"
)

func (c *Character) CustomizationWear(moveAction map[string]any) {
	customizationWear := new(wearCustomization)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &customizationWear)
	if err != nil {
		log.Println(err)
		return
	}

	for _, SID := range customizationWear.Suites {
		customization, err := GetCustomizationByID(SID)
		if err != nil {
			log.Println(err)
			return
		}

		parentID := customization.Parent

		if parentID == lowerParentID {
			c.Customization.Feet = customization.Props.Feet
			continue
		}

		if parentID == upperParentID {
			c.Customization.Body = customization.Props.Body
			c.Customization.Hands = customization.Props.Hands
			continue
		}
	}
}

type hideoutUpgrade struct {
	Action    string
	AreaType  int8            `json:"areaType"`
	Items     []tradingScheme `json:"items"`
	TimeStamp float64         `json:"timeStamp"`
}

func (c *Character) HideoutUpgrade(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	log.Println("HideoutUpgrade")
	upgrade := new(hideoutUpgrade)
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &upgrade); err != nil {
		log.Println(err)
		return
	}

	hideoutArea := GetHideoutAreaByAreaType(upgrade.AreaType)

	log.Println(hideoutArea)
}

type bindItem struct {
	Action string
	Item   string `json:"item"`
	Index  string `json:"index"`
}

func (c *Character) BindItem(moveAction map[string]any) {
	bind := new(bindItem)
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &bind); err != nil {
		log.Println(err)
		return
	}

	if _, ok := c.Inventory.FastPanel[bind.Index]; !ok {
		c.Inventory.FastPanel[bind.Index] = bind.Item
	} else {
		if c.Inventory.FastPanel[bind.Index] == bind.Item {
			c.Inventory.FastPanel[bind.Index] = ""
		} else {
			c.Inventory.FastPanel[bind.Index] = bind.Item
		}
	}
}

type tagItem struct {
	Action   string
	Item     string `json:"item"`
	TagName  string `json:"TagName"`
	TagColor string `json:"TagColor"`
}

func (c *Character) TagItem(moveAction map[string]any) {
	tag := new(tagItem)
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &tag); err != nil {
		log.Println(err)
		return
	}

	cache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	index := *cache.GetIndexOfItemByID(tag.Item)
	if c.Inventory.Items[index].UPD == nil {
		c.Inventory.Items[index].UPD = new(ItemUpdate)
		c.Inventory.Items[index].UPD.Tag = new(Tag)

		c.Inventory.Items[index].UPD.Tag.Color = tag.TagColor
		c.Inventory.Items[index].UPD.Tag.Name = tag.TagName

	} else if c.Inventory.Items[index].UPD.Tag == nil {
		c.Inventory.Items[index].UPD.Tag = new(Tag)

		c.Inventory.Items[index].UPD.Tag.Color = tag.TagColor
		c.Inventory.Items[index].UPD.Tag.Name = tag.TagName

	} else {
		c.Inventory.Items[index].UPD.Tag.Color = tag.TagColor
		c.Inventory.Items[index].UPD.Tag.Name = tag.TagName
	}

}

type toggleItem struct {
	Action string
	Item   string `json:"item"`
	Value  bool   `json:"value"`
}

func (c *Character) ToggleItem(moveAction map[string]any) {
	toggle := new(toggleItem)
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &toggle); err != nil {
		log.Println(err)
		return
	}

	cache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	index := *cache.GetIndexOfItemByID(toggle.Item)
	if c.Inventory.Items[index].UPD == nil {
		c.Inventory.Items[index].UPD = new(ItemUpdate)
		c.Inventory.Items[index].UPD.Togglable = new(Toggle)
		c.Inventory.Items[index].UPD.Togglable.On = toggle.Value

	} else if c.Inventory.Items[index].UPD.Togglable == nil {
		c.Inventory.Items[index].UPD.Togglable = new(Toggle)
		c.Inventory.Items[index].UPD.Togglable.On = toggle.Value

	} else {
		c.Inventory.Items[index].UPD.Togglable.On = toggle.Value
	}
}

type hideoutUpgradeComplete struct {
	Action    string
	AreaType  int8    `json:"areaType"`
	TimeStamp float64 `json:"timeStamp"`
}

func (c *Character) HideoutUpgradeComplete(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	log.Println("HideoutUpgradeComplete")
	upgradeComplete := new(hideoutUpgradeComplete)
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &upgradeComplete); err != nil {
		log.Println(err)
		return
	}

	log.Println(upgradeComplete)
}
