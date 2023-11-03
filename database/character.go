package database

import (
	"fmt"
	"log"
	"path/filepath"
	"slices"
	"strings"

	"MT-GO/services"
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

// #region Character getters

func GetCharacterByUID(uid string) *Character {
	if profile, ok := profiles[uid]; ok {
		return profile.Character
	}

	fmt.Println("Profile with UID ", uid, " does not have a character")
	return nil
}

func (c *Character) GetQuestsAvailableToPlayer() []any {
	var output []any

	query := GetQuestsQuery()

	cachedQuests := GetQuestCacheByUID(c.ID)
	characterHasQuests := len(cachedQuests.Index) != 0

	traderStandings := make(map[string]*float64) //temporary

	for key, value := range query {

		if services.CheckIfQuestForOtherFaction(c.Info.Side, key) {
			continue
		}

		if strings.Contains(value.Name, "-Event") {
			//TODO: filter events properly
			continue
		}

		if value.Conditions == nil || value.Conditions.AvailableForStart == nil {
			output = append(output, GetQuestByQID(key))
			continue
		}

		forStart := value.Conditions.AvailableForStart

		if forStart.Level != nil {
			if !services.LevelComparisonCheck(
				forStart.Level.Level,
				float64(c.Info.Level),
				forStart.Level.CompareMethod) {

				continue
			}
		}

		if forStart.Quest == nil && forStart.TraderLoyalty == nil && forStart.TraderStanding == nil {
			output = append(output, GetQuestByQID(key))
			continue
		}

		loyaltyCheck := false
		if forStart.TraderLoyalty != nil {
			for trader, loyalty := range forStart.TraderLoyalty {

				if traderStandings[trader] == nil {
					if data, err := GetTraderByUID(trader); err != nil {
						log.Fatalln(err)
					} else {
						loyaltyLevel := float64(data.GetTraderLoyaltyLevel(c))
						traderStandings[trader] = &loyaltyLevel
					}
				}

				loyaltyCheck = services.LevelComparisonCheck(
					loyalty.Level,
					*traderStandings[trader],
					loyalty.CompareMethod)
			}

			if !loyaltyCheck {
				continue
			}
		}

		standingCheck := false
		if forStart.TraderStanding != nil {
			for trader, loyalty := range forStart.TraderStanding {

				if traderStandings[trader] == nil {
					if data, err := GetTraderByUID(trader); err != nil {
						log.Fatalln(err)
					} else {
						loyaltyLevel := float64(data.GetTraderLoyaltyLevel(c))
						traderStandings[trader] = &loyaltyLevel
					}
				}

				standingCheck = services.LevelComparisonCheck(
					loyalty.Level,
					*traderStandings[trader],
					loyalty.CompareMethod)
			}

			if !standingCheck {
				continue
			}
		}

		if forStart.Quest != nil && characterHasQuests {
			if c.CompletedPreviousQuestCheck(forStart.Quest, cachedQuests) {
				output = append(output, GetQuestByQID(key))
				continue
			}
		}
	}

	return output
}

func (c *Character) CompletedPreviousQuestCheck(quests map[string]*QuestCondition, cachedQuests *QuestCache) bool {
	previousQuestCompleted := false
	for _, v := range quests {
		index, ok := cachedQuests.Index[v.PreviousQuestID]
		if !ok {
			continue
		}

		previousQuestCompleted = v.Status == c.Quests[index].Status
	}
	return previousQuestCompleted
}

// #endregion

// #region Character functions

func (inv *Inventory) CleanInventoryOfDeletedItemMods() bool {
	allItems := GetItems()

	newItems := make([]InventoryItem, 0, len(inv.Items))

	cleaned := 0
	for _, item := range inv.Items {
		if _, ok := allItems[item.TPL]; !ok {
			cleaned++
			continue
		}
		newItems = append(newItems, item)
	}

	if cleaned != 0 {
		fmt.Println("Removed", cleaned, "modded item(s) from your inventory")
		inv.Items = newItems
		return true
	}
	return false

}

func (c *Character) SaveCharacter(sessionID string) {
	characterFilePath := filepath.Join(profilesPath, sessionID, "character.json")

	err := tools.WriteToFile(characterFilePath, c)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Character saved")
}

// QuestAccept updates an existing Accepted quest, or creates and appends new Accepted Quest to cache and Character
func (c *Character) QuestAccept(qid string, profileChangesEvent *ProfileChangesEvent) {

	cachedQuests := GetQuestCacheByUID(c.ID)
	length := len(cachedQuests.Index)
	time := int(tools.GetCurrentTimeInSeconds())

	query := GetQuestFromQueryByQID(qid)

	quest, ok := cachedQuests.Index[qid]
	if ok { //if exists, update cache and copy to quest on character
		cachedQuest := c.Quests[quest]

		cachedQuest.Status = "Started"
		cachedQuest.StartTime = time
		cachedQuest.StatusTimers[cachedQuest.Status] = time

	} else {
		quest := &CharacterQuest{
			QID:          qid,
			StartTime:    time,
			Status:       "Started",
			StatusTimers: map[string]int{},
		}

		if query.Conditions.AvailableForStart != nil && query.Conditions.AvailableForStart.Quest != nil {
			startCondition := query.Conditions.AvailableForStart.Quest
			for _, questCondition := range startCondition {
				if questCondition.AvailableAfter > 0 {

					quest.StartTime = 0
					quest.Status = "AvailableAfter"
					quest.AvailableAfter = time + questCondition.AvailableAfter
				}
			}
		}

		cachedQuests.Index[qid] = int8(length)
		c.Quests = append(c.Quests, *quest)
	}

	if query.Rewards.Start != nil {
		fmt.Println("There are rewards heeyrrrr!")
		fmt.Println(profileChangesEvent.ProfileChanges[c.ID].ID)

		// TODO: Apply then Get Quest rewards and then route messages from there
		// Character.ApplyQuestRewardsToCharacter()  applies the given reward
		// Quests.GetQuestReward() returns the given reward
		// CreateNPCMessageWithReward()
	}

	dialogue := *GetDialogueByUID(c.ID)

	dialog, message := CreateQuestDialogue(c.ID, "QuestStart", query.Trader, query.Dialogue.Description)
	dialog.New++
	dialog.Messages = append(dialog.Messages, *message)

	dialogue[query.Trader] = dialog

	notification := CreateNotification(message)

	connection := GetConnection(c.ID)
	if connection == nil {
		fmt.Println("Can't send message to character because connection is nil, storing...")
		storage := GetStorageByUID(c.ID)
		storage.Mailbox = append(storage.Mailbox, notification)
		storage.SaveStorage(c.ID)
	} else {
		connection.SendMessage(notification)
	}

	//TODO: Get new player quests from database now that we've accepted one
	profileChangesEvent.ProfileChanges[c.ID].Quests = c.GetQuestsAvailableToPlayer()

	dialogue.SaveDialogue(c.ID)
	c.SaveCharacter(c.ID)
}

func (c *Character) ApplyQuestRewardsToCharacter(rewards *QuestRewards) {
	fmt.Println()
}

type examine struct {
	Action    string     `json:"Action"`
	Item      string     `json:"item"`
	FromOwner *fromOwner `json:"fromOwner,omitempty"`
}
type fromOwner struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (c *Character) ExamineItem(moveAction map[string]any) {
	examine := new(examine)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &examine)
	if err != nil {
		log.Fatalln(err)
	}

	var item *DatabaseItem
	if examine.FromOwner == nil {
		fmt.Println("Examining Item from Player Inventory")
		if index := GetCacheByUID(c.ID).Inventory.GetIndexOfItemByUID(examine.Item); index != nil {
			itemInInventory := c.Inventory.Items[*index]
			item = GetItemByUID(itemInInventory.TPL)
		} else {
			fmt.Println("[EXAMINE] Examining Item", examine.Item, " from Player Inventory failed, does not exist!")
			return
		}
	} else {
		switch examine.FromOwner.Type {
		case "Trader":
			data, err := GetTraderByUID(examine.FromOwner.ID)
			if err != nil {
				log.Fatalln(err)
			}

			assortItem := data.GetAssortItemByID(examine.Item)
			item = GetItemByUID(assortItem[0].Tpl)

		case "HideoutUpgrade":
		case "HideoutProduction":
		case "ScavCase":
			item = GetItemByUID(examine.Item)

		case "RagFair":
		default:
			fmt.Println("[EXAMINE] FromOwner.Type: ", examine.FromOwner.Type, "is not supported, returning...")
			return
		}
	}

	if item == nil {
		fmt.Println("[EXAMINE] Examining Item", examine.Item, "failed, does not exist in Item Database")
		return
	}

	c.Encyclopedia[item.ID] = true
	fmt.Println("[EXAMINE] Encyclopedia entry added for", item.ID)

	//add experience
	experience, ok := item.Props["ExamineExperience"].(float64)
	if !ok {
		fmt.Println("[EXAMINE] Item", examine.Item, "does not have ExamineExperience property, returning...")
		return
	}

	c.Info.Experience += int32(experience)
}

type move struct {
	Action string
	Item   string `json:"item"`
	To     moveTo `json:"to"`
}

type moveTo struct {
	ID        string          `json:"id"`
	Container string          `json:"container"`
	Location  *moveToLocation `json:"location,omitempty"`
}

type moveToLocation struct {
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	R          string  `json:"r"`
	IsSearched bool    `json:"isSearched"`
}

func (c *Character) MoveItemInStash(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	move := new(move)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &move)
	if err != nil {
		log.Fatalln(err)
	}

	cache := *GetCacheByUID(c.ID).Inventory
	index := cache.GetIndexOfItemByUID(move.Item)
	itemInInventory := &c.Inventory.Items[*index]

	if move.To.Location != nil {
		moveToLocation := move.To.Location
		var rotation float64 = 0
		if moveToLocation.R == "Vertical" || moveToLocation.R == "1" {
			rotation++
		}

		if itemInInventory.Location == nil {
			itemInInventory.Location = new(InventoryItemLocation)
		}

		itemInInventory.Location.Y = moveToLocation.Y
		itemInInventory.Location.X = moveToLocation.X
		itemInInventory.Location.R = rotation
		itemInInventory.Location.IsSearched = moveToLocation.IsSearched
	} else {
		itemInInventory.Location = nil
	}

	itemInInventory.ParentID = move.To.ID
	itemInInventory.SlotID = move.To.Container

	if itemInInventory.Location != nil {
		cache.UpdateItemFlatMapLookup([]InventoryItem{*itemInInventory})
	}
	if itemInInventory.SlotID != "hideout" {
		cache.ClearItemFromContainerMap(move.Item)
	} else {
		cache.AddItemFromContainerMap(move.Item)
	}
	//cache.SetInventoryIndex(&c.Inventory)

	profileChangesEvent.ProfileChanges[c.ID].Production = nil
}

type swap struct {
	Action string
	Item   string `json:"item"`
	To     moveTo `json:"to"`
	Item2  string `json:"item2"`
	To2    moveTo `json:"to2"`
}

func (c *Character) SwapItemInStash(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	swap := new(swap)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &swap)
	if err != nil {
		log.Fatalln(err)
	}

	cache := *GetCacheByUID(c.ID).Inventory.GetIndexOfItemByUID(swap.Item)
	itemInInventory := &c.Inventory.Items[cache]

	if swap.To.Location != nil {
		moveToLocation := swap.To.Location
		var rotation float64 = 0
		if moveToLocation.R == "Vertical" || moveToLocation.R == "1" {
			rotation++
		}

		if itemInInventory.Location == nil {
			itemInInventory.Location = new(InventoryItemLocation)
		}

		itemInInventory.Location.Y = moveToLocation.Y
		itemInInventory.Location.X = moveToLocation.X
		itemInInventory.Location.R = rotation
		itemInInventory.Location.IsSearched = moveToLocation.IsSearched
	} else {
		itemInInventory.Location = nil
	}

	itemInInventory.ParentID = swap.To.ID
	itemInInventory.SlotID = swap.To.Container

	cache = GetCacheByUID(c.ID).Inventory.Lookup.Forward[swap.Item2]
	itemInInventory = &c.Inventory.Items[cache]

	if swap.To2.Location != nil {
		moveToLocation := swap.To2.Location
		var rotation float64 = 0
		if moveToLocation.R == "Vertical" || moveToLocation.R == "1" {
			rotation++
		}

		if itemInInventory.Location == nil {
			itemInInventory.Location = new(InventoryItemLocation)
		}

		itemInInventory.Location.Y = moveToLocation.Y
		itemInInventory.Location.X = moveToLocation.X
		itemInInventory.Location.R = rotation
		itemInInventory.Location.IsSearched = moveToLocation.IsSearched
	} else {
		itemInInventory.Location = nil
	}

	itemInInventory.ParentID = swap.To2.ID
	itemInInventory.SlotID = swap.To2.Container

	profileChangesEvent.ProfileChanges[c.ID].Production = nil
}

type foldItem struct {
	Action string
	Item   string `json:"item"`
	Value  bool   `json:"value"`
}

func (c *Character) FoldItem(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	fold := new(foldItem)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &fold)
	if err != nil {
		log.Fatalln(err)
	}

	inventoryCache := GetCacheByUID(c.ID).Inventory
	index := inventoryCache.GetIndexOfItemByUID(fold.Item)
	if index == nil {
		log.Fatalln("Item", fold.Item, "does not exist in cache!")
	}
	itemInInventory := &c.Inventory.Items[*index]
	if itemInInventory.UPD == nil || itemInInventory.UPD.Foldable == nil {
		log.Fatalln(itemInInventory.ID, "cannot be folded!")
	}

	itemInInventory.UPD.Foldable.Folded = fold.Value

	inventoryCache.ResetItemSizeInContainer(itemInInventory, &c.Inventory)
	profileChangesEvent.ProfileChanges[c.ID].Production = nil
}

type readEncyclopedia struct {
	Action string   `json:"Action"`
	IDs    []string `json:"ids"`
}

func (c *Character) ReadEncyclopedia(moveAction map[string]any) {
	readEncyclopedia := new(readEncyclopedia)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &readEncyclopedia)
	if err != nil {
		log.Fatalln(err)
	}

	for _, id := range readEncyclopedia.IDs {
		c.Encyclopedia[id] = true
	}
}

type merge struct {
	Action string
	Item   string `json:"item"`
	With   string `json:"with"`
}

func (c *Character) MergeItem(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	merge := new(merge)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &merge)
	if err != nil {
		log.Fatalln(err)
	}

	inventoryCache := GetCacheByUID(c.ID).Inventory

	toMergeIndex := *inventoryCache.GetIndexOfItemByUID(merge.Item)
	toMerge := &c.Inventory.Items[toMergeIndex]

	mergeWithIndex := *inventoryCache.GetIndexOfItemByUID(merge.With)
	mergeWith := &c.Inventory.Items[mergeWithIndex]

	mergeWith.UPD.StackObjectsCount += toMerge.UPD.StackObjectsCount

	inventoryCache.ClearItemFromContainer(toMerge.ID)
	c.Inventory.RemoveSingleItemFromInventoryByIndex(toMergeIndex)
	inventoryCache.SetInventoryIndex(&c.Inventory)

	profileChangesEvent.ProfileChanges[c.ID].Items.Change = append(profileChangesEvent.ProfileChanges[c.ID].Items.Change, mergeWith)
	profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, &InventoryItem{ID: toMerge.ID})
}

// RemoveSingleItemFromInventoryByIndex takes the existing Inventory.Items and removes an InventoryItem at its index
// by shifting the indexes to the left
func (inv *Inventory) RemoveSingleItemFromInventoryByIndex(index int16) {
	if index < 0 || index >= int16(len(inv.Items)) {
		log.Fatalln("[RemoveSingleItemFromInventoryByIndex] Index out of Range")
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

type transfer struct {
	Action string
	Item   string `json:"item"`
	With   string `json:"with"`
	Count  int32  `json:"count"`
}

func (c *Character) TransferItem(moveAction map[string]any) {
	transfer := new(transfer)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &transfer)
	if err != nil {
		log.Fatalln(err)
	}

	inventoryCache := GetCacheByUID(c.ID).Inventory

	toMergeIndex := *inventoryCache.GetIndexOfItemByUID(transfer.Item)
	toMerge := &c.Inventory.Items[toMergeIndex]

	mergeWithIndex := *inventoryCache.GetIndexOfItemByUID(transfer.With)
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
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &split)
	if err != nil {
		log.Fatalln(err)
	}

	invCache := *GetCacheByUID(c.ID).Inventory

	originalItem := &c.Inventory.Items[*invCache.GetIndexOfItemByUID(split.SplitItem)]
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

	profileChangesEvent.ProfileChanges[c.ID].Items.Change = append(profileChangesEvent.ProfileChanges[c.ID].Items.Change, originalItem)
	profileChangesEvent.ProfileChanges[c.ID].Items.New = append(profileChangesEvent.ProfileChanges[c.ID].Items.New, &InventoryItem{ID: newItem.ID, TPL: newItem.TPL, UPD: newItem.UPD})
}

type remove struct {
	Action string `json:"Action"`
	ItemId string `json:"item"`
}

func (c *Character) RemoveItem(moveAction map[string]any, profileChangesEvent *ProfileChangesEvent) {
	remove := new(remove)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &remove)
	if err != nil {
		log.Fatalln(err)
	}

	inventoryCache := GetCacheByUID(c.ID).Inventory
	itemChildren := GetInventoryItemFamilyTreeIDs(c.Inventory.Items, remove.ItemId)

	var itemIndex int16
	toDelete := make([]int16, 0, len(itemChildren))
	for _, itemID := range itemChildren {
		itemIndex = *inventoryCache.GetIndexOfItemByUID(itemID)
		toDelete = append(toDelete, itemIndex)
	}

	inventoryCache.ClearItemFromContainer(remove.ItemId)
	c.Inventory.RemoveItemsFromInventoryByIndices(toDelete)
	inventoryCache.SetInventoryIndex(&c.Inventory)

	profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, &InventoryItem{ID: remove.ItemId})

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
		log.Fatalln(err)
	}

	cache := *GetCacheByUID(c.ID).Inventory
	for _, item := range applyInventoryChanges.ChangedItems {
		properties, ok := item.(map[string]any)
		if !ok {
			log.Fatalln("Cannot type assert item from Auto-Sort items slice")
		}

		UID, ok := properties["_id"].(string)
		if !ok {
			log.Fatalln("Cannot type assert item `_id` property from Auto-Sort items slice")
		}
		itemInInventory := &c.Inventory.Items[*cache.GetIndexOfItemByUID(UID)]

		parent, ok := properties["parentId"].(string)
		if !ok {
			log.Fatalln("Cannot type assert item `parentId` property from Auto-Sort items slice")
		}
		itemInInventory.ParentID = parent

		slotId, ok := properties["slotId"].(string)
		if !ok {
			log.Fatalln("Cannot type assert item `slotId` property from Auto-Sort items slice")
		}
		itemInInventory.SlotID = slotId

		location, ok := properties["location"].(map[string]any)
		if !ok {
			itemInInventory.Location = nil
			continue
		} else {
			r, ok := location["r"].(string)
			if !ok {
				log.Fatalln("Cannot type assert item.Location `r` property from Auto-Sort items slice")
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
	//TODO: Make everything purchased NOT free lol

	invCache := GetCacheByUID(c.ID).Inventory

	switch moveAction["type"] {
	case "buy_from_trader":
		buy := new(buyFromTrader)
		data, _ := json.Marshal(moveAction)
		err := json.Unmarshal(data, &buy)
		if err != nil {
			log.Fatalln(err)
		}

		c.BuyFromTrader(buy, invCache, profileChangesEvent)
	case "sell_to_trader":
		sell := new(sellToTrader)
		data, _ := json.Marshal(moveAction)
		err := json.Unmarshal(data, &sell)
		if err != nil {
			log.Fatalln(err)
		}

		c.SellToTrader(sell, invCache, profileChangesEvent)
	default:
		fmt.Println("YO! TRADINGCONFIRM.", moveAction["type"], "ISNT SUPPORTED YET HAHAHHAHAHAHAHAHHAHAHHHHHHHHHHHHHAHAHAHAHAHHAHA")
	}
}

func (c *Character) BuyFromTrader(tradeConfirm *buyFromTrader, invCache *InventoryContainer, profileChangesEvent *ProfileChangesEvent) {
	trader, err := GetTraderByUID(tradeConfirm.TID)
	if err != nil {
		log.Fatalln(err)
	}

	assortItem := trader.GetAssortItemByID(tradeConfirm.ItemID)
	if assortItem == nil {
		log.Fatalln("Item of", tradeConfirm.ItemID, "does not exist in trader assort, killing!")
	}

	inventoryItems := ConvertAssortItemsToInventoryItem(assortItem, &c.Inventory.Stash)
	if len(inventoryItems) == 0 {
		log.Fatalln("Converting Assort Item to Inventory Item failed, killing")
	}

	var stackMaxSize = *GetItemByUID(inventoryItems[len(inventoryItems)-1].TPL).GetStackMaxSize()
	var stackSlice = services.GetCorrectAmountOfItemsPurchased(tradeConfirm.Count, stackMaxSize)
	// Basically gets the correct amount of items to be created, based on StackSize

	//Create copy-of Character.Inventory.Items for modification in the case of any failures to assign later
	copyOfItems := make([]InventoryItem, 0, len(c.Inventory.Items)+len(inventoryItems))
	copyOfItems = append(copyOfItems, c.Inventory.Items...)
	//Create copy-of invCache.Stash.Container for modification in the case of any failures to assign later
	copyOfMap := invCache.Stash.Container

	toAdd := make([]InventoryItem, 0, len(stackSlice))
	toDelete := make(map[string]int16)
	traderRelations := c.TradersInfo[tradeConfirm.TID]

	height, width := MeasurePurchaseForInventoryMapping(inventoryItems)

	for _, stack := range stackSlice {
		var copyOfInventoryItems []InventoryItem
		if len(stackSlice) != 1 {
			copyOfInventoryItems = AssignNewIDs(inventoryItems)
		} else {
			copyOfInventoryItems = inventoryItems
		}

		mainItem := &copyOfInventoryItems[len(copyOfInventoryItems)-1]

		validLocation := invCache.GetValidLocationForItem(height, width)
		if validLocation == nil {
			fmt.Println("Item", tradeConfirm.ItemID, "was not purchased because we could not find a position in your inventory!!")
			invCache.Stash.Container = copyOfMap //if failure, assign old map
			return
		}

		stackObjectsCount := stack
		mainItem.UPD.StackObjectsCount = stackObjectsCount
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
		index := invCache.GetIndexOfItemByUID(scheme.ID)
		if index == nil {
			log.Fatalln("Index of", scheme.ID, "does not exist in cache, killing!")
		}

		itemInInventory := copyOfItems[*index]

		currency := *GetCurrencyByName(trader.Base.Currency)
		if IsCurrencyByUID(itemInInventory.TPL) {
			traderRelations.SalesSum += float32(scheme.Count)
		} else {
			priceOfItem := *GetPriceByID(itemInInventory.TPL)
			if "RUB" != trader.Base.Currency {
				conversion := ConvertFromRouble(priceOfItem, currency)
				traderRelations.SalesSum += float32(conversion)
			} else {
				traderRelations.SalesSum += float32(priceOfItem)
			}
		}

		if itemInInventory.UPD != nil && itemInInventory.UPD.StackObjectsCount != 0 {
			var remainingBalance = scheme.Count

			if itemInInventory.UPD.StackObjectsCount > remainingBalance {
				itemInInventory.UPD.StackObjectsCount -= remainingBalance

				profileChangesEvent.ProfileChanges[c.ID].Items.Change = append(profileChangesEvent.ProfileChanges[c.ID].Items.Change, &itemInInventory)
			} else if itemInInventory.UPD.StackObjectsCount == remainingBalance {
				toDelete[itemInInventory.ID] = *index
			} else {
				remainingBalance -= itemInInventory.UPD.StackObjectsCount

				toDelete[itemInInventory.ID] = *index

				//TODO: Consider creating a look-up cache for mergable Inventory.Items

				var toChange []*InventoryItem
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
					toChange = append(toChange, &item)
					break
				}
				if remainingBalance > 0 {
					fmt.Println("Insufficient funds to purchase item, returning")
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
	for _, invItem := range toAdd {
		copyOfItems = append(copyOfItems, invItem)
		profileChangeItem := invItem //assign to variable to be pointed too for profileChangeEvents
		profileChangesEvent.ProfileChanges[c.ID].Items.New = append(profileChangesEvent.ProfileChanges[c.ID].Items.New, &profileChangeItem)
	}

	//Assign copy-of Character.Inventory.Items to original Character.Inventory.Items
	c.Inventory.Items = copyOfItems

	if len(toDelete) != 0 {
		indices := make([]int16, 0, len(toDelete))
		for id, idx := range toDelete {
			invCache.ClearItemFromContainer(id)
			indices = append(indices, idx)

			profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, &InventoryItem{ID: id})
		}
		c.Inventory.RemoveItemsFromInventoryByIndices(indices)
	}
	invCache.SetInventoryIndex(&c.Inventory)

	profileChangesEvent.ProfileChanges[c.ID].TraderRelations[tradeConfirm.TID] = traderRelations
	c.TradersInfo[tradeConfirm.TID] = traderRelations

	fmt.Println(len(toAdd), "of Item", tradeConfirm.ItemID, "purchased!")
}

func (c *Character) SellToTrader(tradeConfirm *sellToTrader, invCache *InventoryContainer, profileChangesEvent *ProfileChangesEvent) {
	trader, err := GetTraderByUID(tradeConfirm.TID)
	if err != nil {
		log.Fatalln(err)
	}

	saleCurrency := *GetCurrencyByName(trader.Base.Currency)

	var remainingBalance = tradeConfirm.Price
	stackMaxSize := *GetItemByUID(saleCurrency).GetStackMaxSize()

	cache := GetCacheByUID(c.ID).Inventory
	copyOfMap := invCache.Stash.Container
	copyOfItems := make([]InventoryItem, 0, len(c.Inventory.Items))
	copyOfItems = append(copyOfItems, c.Inventory.Items...)

	toDelete := make(map[string]int16)
	for _, item := range tradeConfirm.Items {
		index := *cache.GetIndexOfItemByUID(item.ID)
		toDelete[item.ID] = index
	}

	toChange := make([]*InventoryItem, 0)
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

			toChange = append(toChange, &item)
			continue
		} else {
			item.UPD.StackObjectsCount += remainingBalance
			remainingBalance = 0
			toChange = append(toChange, &item)
			break
		}
	}

	if remainingBalance != 0 {
		var toAdd []InventoryItem

		fmt.Println("If a new stack isn't made, we cry")

		stackSlice := services.GetCorrectAmountOfItemsPurchased(remainingBalance, stackMaxSize)
		item := []InventoryItem{*CreateNewItem(saleCurrency, c.Inventory.Stash)}
		// since it's one item, just get the height and width once
		height, width := MeasurePurchaseForInventoryMapping(item)

		for _, stack := range stackSlice {
			mainItem := AssignNewIDs(item)[0]

			validLocation := invCache.GetValidLocationForItem(height, width)
			if validLocation == nil {
				fmt.Println("Item", mainItem.ID, "was not created because we could not find a position in your inventory!")
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

		for _, invItem := range toAdd {
			copyOfItems = append(copyOfItems, invItem)
			profileChangeItem := invItem //assign to variable to be pointed too for profileChangeEvents
			profileChangesEvent.ProfileChanges[c.ID].Items.New = append(profileChangesEvent.ProfileChanges[c.ID].Items.New, &profileChangeItem)
		}

	}

	profileChangesEvent.ProfileChanges[c.ID].Items.Change = append(profileChangesEvent.ProfileChanges[c.ID].Items.Change, toChange...)
	c.Inventory.Items = copyOfItems

	if len(toDelete) != 0 {
		indices := make([]int16, len(toDelete))
		for id, idx := range toDelete {
			invCache.ClearItemFromContainer(id)
			indices = append(indices, idx)

			profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, &InventoryItem{ID: id})
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
		log.Fatalln(err)
	}

	trader, err := GetTraderByName("Ragman")
	if err != nil {
		log.Fatalln(err)
	}
	suitsIndex := trader.Index.Suits[customizationBuy.Offer]
	suitID := trader.Suits[suitsIndex].SuiteID

	storage := GetStorageByUID(c.ID)
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
		log.Fatalln(err)
	}

	for _, SID := range customizationWear.Suites {
		customization := GetCustomization(SID)
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
	fmt.Println("HideoutUpgrade")
	upgrade := new(hideoutUpgrade)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &upgrade)
	if err != nil {
		log.Fatalln(err)
	}

	hideoutArea := GetHideoutAreaByAreaType(upgrade.AreaType)

	fmt.Println(hideoutArea)
	fmt.Println()
}

type bindItem struct {
	Action string
	Item   string `json:"item"`
	Index  string `json:"index"`
}

func (c *Character) BindItem(moveAction map[string]any) {
	bind := new(bindItem)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &bind)
	if err != nil {
		log.Fatalln(err)
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
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &tag)
	if err != nil {
		log.Fatalln(err)
	}

	index := *GetCacheByUID(c.ID).Inventory.GetIndexOfItemByUID(tag.Item)
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
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &toggle)
	if err != nil {
		log.Fatalln(err)
	}

	index := *GetCacheByUID(c.ID).Inventory.GetIndexOfItemByUID(toggle.Item)
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
	fmt.Println("HideoutUpgradeComplete")
	upgradeComplete := new(hideoutUpgradeComplete)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &upgradeComplete)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(upgradeComplete)
	fmt.Println()
}

// #endregion

// #region Character structs

type Character struct {
	ID                string              `json:"_id"`
	AID               int                 `json:"aid"`
	Savage            *string             `json:"savage"`
	Info              PlayerInfo          `json:"Info"`
	Customization     PlayerCustomization `json:"Customization"`
	Health            HealthInfo          `json:"Health"`
	Inventory         Inventory           `json:"Inventory"`
	Skills            PlayerSkills        `json:"Skills"`
	Stats             PlayerStats         `json:"Stats"`
	Encyclopedia      map[string]bool     `json:"Encyclopedia"`
	ConditionCounters ConditionCounters   `json:"ConditionCounters"`
	BackendCounters   map[string]any      `json:"BackendCounters"`
	InsuredItems      []InsuredItem       `json:"InsuredItems"`
	Hideout           PlayerHideout       `json:"Hideout"`
	Bonuses           []Bonus             `json:"Bonuses"`
	Notes             struct {
		Notes [][]any `json:"Notes"`
	} `json:"Notes"`
	Quests       []CharacterQuest             `json:"Quests"`
	RagfairInfo  PlayerRagfairInfo            `json:"RagfairInfo"`
	WishList     []string                     `json:"WishList"`
	TradersInfo  map[string]PlayerTradersInfo `json:"TradersInfo"`
	UnlockedInfo struct {
		UnlockedProductionRecipe []any `json:"unlockedProductionRecipe"`
	} `json:"UnlockedInfo"`
}

type PlayerTradersInfo struct {
	Unlocked bool    `json:"unlocked"`
	Disabled bool    `json:"disabled"`
	SalesSum float32 `json:"salesSum"`
	Standing float32 `json:"standing"`
}

type PlayerRagfairInfo struct {
	Rating          float32 `json:"rating"`
	IsRatingGrowing bool    `json:"isRatingGrowing"`
	Offers          []any   `json:"offers"`
}

type PlayerHideoutArea struct {
	Type                  int    `json:"type"`
	Level                 int    `json:"level"`
	Active                bool   `json:"active"`
	PassiveBonusesEnabled bool   `json:"passiveBonusesEnabled"`
	CompleteTime          int    `json:"completeTime"`
	Constructing          bool   `json:"constructing"`
	Slots                 []any  `json:"slots"`
	LastRecipe            string `json:"lastRecipe"`
}
type PlayerHideout struct {
	Production  map[string]any      `json:"Production"`
	Areas       []PlayerHideoutArea `json:"Areas"`
	Improvement map[string]any      `json:"Improvement"`
	//Seed        int                    `json:"Seed"`
}

type ConditionCounters struct {
	Counters []any `json:"Counters"`
}

type PlayerStats struct {
	Eft EftStats `json:"Eft"`
}

type StatCounters struct {
	Items []any `json:"Items"`
}

type EftStats struct {
	SessionCounters        map[string]any `json:"SessionCounters"`
	OverallCounters        map[string]any `json:"OverallCounters"`
	SessionExperienceMult  int            `json:"SessionExperienceMult"`
	ExperienceBonusMult    int            `json:"ExperienceBonusMult"`
	TotalSessionExperience int            `json:"TotalSessionExperience"`
	LastSessionDate        int            `json:"LastSessionDate"`
	Aggressor              map[string]any `json:"Aggressor"`
	DroppedItems           []any          `json:"DroppedItems"`
	FoundInRaidItems       []any          `json:"FoundInRaidItems"`
	Victims                []any          `json:"Victims"`
	CarriedQuestItems      []any          `json:"CarriedQuestItems"`
	DamageHistory          map[string]any `json:"DamageHistory"`
	LastPlayerState        *float32       `json:"LastPlayerState"`
	TotalInGameTime        int            `json:"TotalInGameTime"`
	SurvivorClass          string         `json:"SurvivorClass"`
}

type SkillsCommon struct {
	ID                        string `json:"Id"`
	Progress                  int    `json:"Progress"`
	PointsEarnedDuringSession int    `json:"PointsEarnedDuringSession"`
	LastAccess                int64  `json:"LastAccess"`
}
type SkillsMastering struct {
	ID       string `json:"Id"`
	Progress int    `json:"Progress"`
}

type PlayerSkills struct {
	Common    []SkillsCommon    `json:"Common"`
	Mastering []SkillsMastering `json:"Mastering"`
	Points    int               `json:"Points"`
}

type PlayerInfo struct {
	Nickname               string         `json:"Nickname"`
	LowerNickname          string         `json:"LowerNickname"`
	Side                   string         `json:"Side"`
	Voice                  string         `json:"Voice"`
	Level                  int8           `json:"Level"`
	Experience             int32          `json:"Experience"`
	RegistrationDate       int32          `json:"RegistrationDate"`
	GameVersion            string         `json:"GameVersion"`
	AccountType            int8           `json:"AccountType"`
	MemberCategory         int8           `json:"MemberCategory"`
	LockedMoveCommands     bool           `json:"lockedMoveCommands"`
	SavageLockTime         int32          `json:"SavageLockTime"`
	LastTimePlayedAsSavage int32          `json:"LastTimePlayedAsSavage"`
	Settings               map[string]any `json:"Settings"`
	NicknameChangeDate     int32          `json:"NicknameChangeDate"`
	NeedWipeOptions        []any          `json:"NeedWipeOptions"`
	LastCompletedWipe      struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedWipe"`
	LastCompletedEvent struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedEvent"`
	BannedState             bool  `json:"BannedState"`
	BannedUntil             int32 `json:"BannedUntil"`
	IsStreamerModeAvailable bool  `json:"IsStreamerModeAvailable"`
	SquadInviteRestriction  bool  `json:"SquadInviteRestriction"`
	Bans                    []any `json:"Bans"`
}

type InfoSettings struct {
	Role            string  `json:"Role"`
	BotDifficulty   string  `json:"BotDifficulty"`
	Experience      int32   `json:"Experience"`
	StandingForKill float32 `json:"StandingForKill"`
	AggressorBonus  float32 `json:"AggressorBonus"`
}

type PlayerCustomization struct {
	Head  string `json:"Head"`
	Body  string `json:"Body"`
	Feet  string `json:"Feet"`
	Hands string `json:"Hands"`
}

type InsuredItem struct {
	Tid    string `json:"tid"`
	ItemID string `json:"itemId"`
}

type Bonus struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	TemplateID string `json:"templateId"`
}

type HealthInfo struct {
	Hydration   CurrMaxHealth   `json:"Hydration"`
	Energy      CurrMaxHealth   `json:"Energy"`
	Temperature CurrMaxHealth   `json:"Temperature"`
	BodyParts   BodyPartsHealth `json:"BodyParts"`
	UpdateTime  int32           `json:"UpdateTime"`
}

type HealthOf struct {
	Health CurrMaxHealth `json:"Health"`
}

type CurrMaxHealth struct {
	Current float32 `json:"Current"`
	Maximum float32 `json:"Maximum"`
}

type BodyPartsHealth struct {
	Head     HealthOf `json:"Head"`
	Chest    HealthOf `json:"Chest"`
	Stomach  HealthOf `json:"Stomach"`
	LeftArm  HealthOf `json:"LeftArm"`
	RightArm HealthOf `json:"RightArm"`
	LeftLeg  HealthOf `json:"LeftLeg"`
	RightLeg HealthOf `json:"RightLeg"`
}

type CharacterQuest struct {
	QID                 string         `json:"qid"`
	StartTime           int            `json:"startTime"`
	Status              string         `json:"status"`
	StatusTimers        map[string]int `json:"statusTimers"`
	CompletedConditions []string       `json:"completedConditions,omitempty"`
	AvailableAfter      int            `json:"availableAfter,omitempty"`
}

// #endregion
