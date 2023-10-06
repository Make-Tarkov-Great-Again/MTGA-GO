package database

import (
	"fmt"
	"log"
	"path/filepath"
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

func (c *Character) GetQuestsAvailableToPlayer() []interface{} {
	var output []interface{}

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
					loyaltyLevel := float64(GetTraderByUID(trader).GetTraderLoyaltyLevel(c))
					traderStandings[trader] = &loyaltyLevel
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
					loyaltyLevel := float64(GetTraderByUID(trader).GetTraderLoyaltyLevel(c))
					traderStandings[trader] = &loyaltyLevel
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

func (c *Character) SaveCharacter(sessionID string) {
	characterFilePath := filepath.Join(profilesPath, sessionID, "character.json")

	err := tools.WriteToFile(characterFilePath, c)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Character saved")
}

// QuestAccept updates an existing Accepted quest, or creates and appends new Accepted Quest to cache and Character
func (c *Character) QuestAccept(qid string) *ProfileChangesEvent {

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

	changeEvent := CreateProfileChangesEvent(c)
	if query.Rewards.Start != nil {
		fmt.Println("There are rewards heeyrrrr!")
		fmt.Println(changeEvent.ProfileChanges[c.ID].ID)

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
	changeEvent.ProfileChanges[c.ID].Quests = c.GetQuestsAvailableToPlayer()

	dialogue.SaveDialogue(c.ID)
	c.SaveCharacter(c.ID)
	return changeEvent
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

func (c *Character) ExamineItem(moveAction map[string]interface{}) {
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
			assortItem := GetTraderByUID(examine.FromOwner.ID).GetAssortItemByID(examine.Item)
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

	c.Info.Experience += int(experience)
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

func (c *Character) MoveItemInStash(moveAction map[string]interface{}, profileChangesEvent *ProfileChangesEvent) {
	move := new(move)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &move)
	if err != nil {
		log.Fatalln(err)
	}

	cache := *GetCacheByUID(c.ID).Inventory.GetIndexOfItemByUID(move.Item)
	itemInInventory := &c.Inventory.Items[cache]

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

	itemInInventory.ParentID = &move.To.ID
	itemInInventory.SlotID = &move.To.Container

	profileChangesEvent.ProfileChanges[c.ID].Production = nil
}

type swap struct {
	Action string
	Item   string `json:"item"`
	To     moveTo `json:"to"`
	Item2  string `json:"item2"`
	To2    moveTo `json:"to2"`
}

func (c *Character) SwapItemInStash(moveAction map[string]interface{}, profileChangesEvent *ProfileChangesEvent) {
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

	itemInInventory.ParentID = &swap.To.ID
	itemInInventory.SlotID = &swap.To.Container

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

	itemInInventory.ParentID = &swap.To2.ID
	itemInInventory.SlotID = &swap.To2.Container

	profileChangesEvent.ProfileChanges[c.ID].Production = nil
}

type fold struct {
	Action string
	Item   string `json:"item"`
	Value  bool   `json:"value"`
}

func (c *Character) FoldItem(moveAction map[string]interface{}, profileChangesEvent *ProfileChangesEvent) {
	fold := new(fold)
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

func (c *Character) ReadEncyclopedia(moveAction map[string]interface{}) {
	readEncyclopedia := new(readEncyclopedia)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &readEncyclopedia)
	if err != nil {
		log.Fatalln(err)
	}

	length := int8(len(readEncyclopedia.IDs))
	for i := int8(0); i < length; i++ {
		c.Encyclopedia[readEncyclopedia.IDs[i]] = true
	}
}

type merge struct {
	Action string
	Item   string `json:"item"`
	With   string `json:"with"`
}

func (c *Character) MergeItem(moveAction map[string]interface{}, profileChangesEvent *ProfileChangesEvent) {
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

	*mergeWith.UPD.StackObjectsCount += *toMerge.UPD.StackObjectsCount

	inventoryCache.ClearItemFromContainer(toMerge.ID)
	c.Inventory.RemoveItemFromInventoryByIndex(toMergeIndex)
	inventoryCache.SetInventoryIndex(&c.Inventory)

	profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, &InventoryItem{ID: toMerge.ID})
}

// RemoveItemFromInventoryByIndex takes the existing Inventory.Items and removes an InventoryItem at its index
// by shifting the indexes to the left
func (inv *Inventory) RemoveItemFromInventoryByIndex(index int16) {
	if index < 0 || index >= int16(len(inv.Items)) {
		log.Fatalln("[RemoveItemFromInventoryByIndex] Index out of Range")
	}
	copy(inv.Items[index:], inv.Items[index+1:])
	inv.Items = inv.Items[:len(inv.Items)-1]
}

type transfer struct {
	Action string
	Item   string `json:"item"`
	With   string `json:"with"`
	Count  int32  `json:"count"`
}

func (c *Character) TransferItem(moveAction map[string]interface{}) {
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

	*toMerge.UPD.StackObjectsCount -= transfer.Count
	*mergeWith.UPD.StackObjectsCount += transfer.Count
}

type split struct {
	Action    string `json:"Action"`
	SplitItem string `json:"splitItem"`
	NewItem   string `json:"newItem"`
	Container moveTo `json:"container"`
	Count     int32  `json:"count"`
}

func (c *Character) SplitItem(moveAction map[string]interface{}, profileChangesEvent *ProfileChangesEvent) {
	split := new(split)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &split)
	if err != nil {
		log.Fatalln(err)
	}

	invCache := *GetCacheByUID(c.ID).Inventory

	originalItem := &c.Inventory.Items[*invCache.GetIndexOfItemByUID(split.SplitItem)]
	*originalItem.UPD.StackObjectsCount -= split.Count

	newItem := InventoryItem{
		ID:  split.NewItem,
		TPL: originalItem.TPL,
		UPD: originalItem.UPD,
		Location: &InventoryItemLocation{
			IsSearched: split.Container.Location.IsSearched,
			X:          split.Container.Location.X,
			Y:          split.Container.Location.Y,
		},
		ParentID: &split.Container.ID,
		SlotID:   &split.Container.Container,
	}

	if split.Container.Location.R == "Vertical" {
		newItem.Location.R = float64(1)
	} else {
		newItem.Location.R = float64(0)
	}

	*newItem.UPD.StackObjectsCount = split.Count

	height, width := MeasurePurchaseForInventoryMapping([]*InventoryItem{&newItem})
	if width != 0 {
		width--
	}
	if height != 0 {
		height--
	}

	itemFlatMap := invCache.CreateFlatMapLookup(height, width, &newItem)

	for column := itemFlatMap.StartX; column <= itemFlatMap.EndX; column++ {
		itemFlatMap.Coordinates = append(itemFlatMap.Coordinates, column)
		for row := int16(1); row <= int16(itemFlatMap.Height); row++ {
			coordinate := row*int16(invCache.Stash.Container.Width) + column
			itemFlatMap.Coordinates = append(itemFlatMap.Coordinates, coordinate)
		}
	}

	c.Inventory.Items = append(c.Inventory.Items, newItem)
	invCache.AddItemToContainer(split.NewItem, itemFlatMap)
	invCache.SetInventoryIndex(&c.Inventory)
	profileChangesEvent.ProfileChanges[c.ID].Items.New = append(profileChangesEvent.ProfileChanges[c.ID].Items.New, &newItem)
}

type remove struct {
	Action string `json:"Action"`
	ItemId string `json:"item"`
}

func (c *Character) RemoveItem(moveAction map[string]interface{}, profileChangesEvent *ProfileChangesEvent) {
	remove := new(remove)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &remove)
	if err != nil {
		log.Fatalln(err)
	}

	inventoryCache := GetCacheByUID(c.ID).Inventory

	itemChildren := GetInventoryItemFamilyTreeIDs(c.Inventory.Items, remove.ItemId)
	var itemIndex int16
	for _, itemID := range itemChildren {
		itemIndex = *inventoryCache.GetIndexOfItemByUID(itemID)
		inventoryCache.ClearItemFromContainer(itemID)
		c.Inventory.RemoveItemFromInventoryByIndex(itemIndex)
		profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, &InventoryItem{ID: itemID})
	}
	inventoryCache.SetInventoryIndex(&c.Inventory)
}

type applyInventoryChanges struct {
	Action       string
	ChangedItems []interface{} `json:"changedItems"`
}

func (c *Character) ApplyInventoryChanges(moveAction map[string]interface{}) {
	applyInventoryChanges := new(applyInventoryChanges)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &applyInventoryChanges)
	if err != nil {
		log.Fatalln(err)
	}

	cache := *GetCacheByUID(c.ID).Inventory
	for _, item := range applyInventoryChanges.ChangedItems {
		properties, ok := item.(map[string]interface{})
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
		itemInInventory.ParentID = &parent

		slotId, ok := properties["slotId"].(string)
		if !ok {
			log.Fatalln("Cannot type assert item `slotId` property from Auto-Sort items slice")
		}
		itemInInventory.SlotID = &slotId

		location, ok := properties["location"].(map[string]interface{})
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

type tradingConfirm struct {
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

type barter struct {
	Item  *InventoryItem
	Count int32
}

func (c *Character) TradingConfirm(moveAction map[string]interface{}, profileChangesEvent *ProfileChangesEvent) {
	tradeConfirm := new(tradingConfirm)
	data, _ := json.Marshal(moveAction)
	err := json.Unmarshal(data, &tradeConfirm)
	if err != nil {
		log.Fatalln(err)
	}

	//TODO: Make everything purchased NOT free lol

	invCache := GetCacheByUID(c.ID).Inventory

	barters := make([]*barter, 0, len(tradeConfirm.SchemeItems))
	for _, scheme := range tradeConfirm.SchemeItems {
		b := new(barter)

		index := invCache.GetIndexOfItemByUID(scheme.ID)
		if index == nil {
			log.Fatalln("Index of", scheme.ID, "does not exist in cache, killing!")
		}

		b.Item = &c.Inventory.Items[*index]
		b.Count = scheme.Count
		barters = append(barters, b)
	}

	// see if item can fit in inventory first
	assortItem := GetTraderByUID(tradeConfirm.TID).GetAssortItemByID(tradeConfirm.ItemID)
	if assortItem == nil {
		log.Fatalln("Item of", tradeConfirm.ItemID, "does not exist in trader assort, killing!")
	}

	inventoryItems := ConvertAssortItemsToInventoryItem(assortItem, &c.Inventory.Stash)
	if len(inventoryItems) == 0 {
		log.Fatalln("Converting Assort Item to Inventory Item failed, killing")
	}

	mainItem := &inventoryItems[len(inventoryItems)-1]
	(*mainItem).UPD.StackObjectsCount = &tradeConfirm.Count //temporary, should be handled in the conversion probably

	height, width := MeasurePurchaseForInventoryMapping(inventoryItems)
	validLocation := invCache.GetValidLocationForItem(height, width)
	if validLocation == nil {
		return
	}

	(*mainItem).Location = &InventoryItemLocation{
		IsSearched: true,
		R:          float64(0),
		X:          float64(validLocation.X),
		Y:          float64(validLocation.Y),
	}

	for _, invItem := range inventoryItems {
		c.Inventory.Items = append(c.Inventory.Items, *invItem)
		profileChangesEvent.ProfileChanges[c.ID].Items.New = append(profileChangesEvent.ProfileChanges[c.ID].Items.New, invItem)
	}

	itemFlatMap := invCache.CreateFlatMapLookup(height, width, *mainItem)
	itemFlatMap.Coordinates = validLocation.MapInfo

	invCache.AddItemToContainer((*mainItem).ID, itemFlatMap)
	invCache.SetInventoryIndex(&c.Inventory)

	fmt.Println("Item", tradeConfirm.ItemID, "purchased!")
}

// #endregion

// #region Character structs

type Character struct {
	ID                string                 `json:"_id"`
	AID               int                    `json:"aid"`
	Savage            *string                `json:"savage"`
	Info              PlayerInfo             `json:"Info"`
	Customization     PlayerCustomization    `json:"Customization"`
	Health            HealthInfo             `json:"Health"`
	Inventory         Inventory              `json:"Inventory"`
	Skills            PlayerSkills           `json:"Skills"`
	Stats             PlayerStats            `json:"Stats"`
	Encyclopedia      map[string]bool        `json:"Encyclopedia"`
	ConditionCounters ConditionCounters      `json:"ConditionCounters"`
	BackendCounters   map[string]interface{} `json:"BackendCounters"`
	InsuredItems      []InsuredItem          `json:"InsuredItems"`
	Hideout           PlayerHideout          `json:"Hideout"`
	Bonuses           []Bonus                `json:"Bonuses"`
	Notes             struct {
		Notes [][]interface{} `json:"Notes"`
	} `json:"Notes"`
	Quests       []CharacterQuest             `json:"Quests"`
	RagfairInfo  PlayerRagfairInfo            `json:"RagfairInfo"`
	WishList     []string                     `json:"WishList"`
	TradersInfo  map[string]PlayerTradersInfo `json:"TradersInfo"`
	UnlockedInfo struct {
		UnlockedProductionRecipe []interface{} `json:"unlockedProductionRecipe"`
	} `json:"UnlockedInfo"`
}

type PlayerTradersInfo struct {
	Unlocked bool    `json:"unlocked"`
	Disabled bool    `json:"disabled"`
	SalesSum float32 `json:"salesSum"`
	Standing float32 `json:"standing"`
}

type PlayerRagfairInfo struct {
	Rating          float64       `json:"rating"`
	IsRatingGrowing bool          `json:"isRatingGrowing"`
	Offers          []interface{} `json:"offers"`
}

type PlayerHideoutArea struct {
	Type                  int           `json:"type"`
	Level                 int           `json:"level"`
	Active                bool          `json:"active"`
	PassiveBonusesEnabled bool          `json:"passiveBonusesEnabled"`
	CompleteTime          int           `json:"completeTime"`
	Constructing          bool          `json:"constructing"`
	Slots                 []interface{} `json:"slots"`
	LastRecipe            string        `json:"lastRecipe"`
}
type PlayerHideout struct {
	Production  map[string]interface{} `json:"Production"`
	Areas       []PlayerHideoutArea    `json:"Areas"`
	Improvement map[string]interface{} `json:"Improvement"`
	//Seed        int                    `json:"Seed"`
}

type ConditionCounters struct {
	Counters []interface{} `json:"Counters"`
}

type PlayerStats struct {
	Eft EftStats `json:"Eft"`
}

type StatCounters struct {
	Items []interface{} `json:"Items"`
}

type EftStats struct {
	SessionCounters        map[string]interface{} `json:"SessionCounters"`
	OverallCounters        map[string]interface{} `json:"OverallCounters"`
	SessionExperienceMult  int                    `json:"SessionExperienceMult"`
	ExperienceBonusMult    int                    `json:"ExperienceBonusMult"`
	TotalSessionExperience int                    `json:"TotalSessionExperience"`
	LastSessionDate        int                    `json:"LastSessionDate"`
	Aggressor              map[string]interface{} `json:"Aggressor"`
	DroppedItems           []interface{}          `json:"DroppedItems"`
	FoundInRaidItems       []interface{}          `json:"FoundInRaidItems"`
	Victims                []interface{}          `json:"Victims"`
	CarriedQuestItems      []interface{}          `json:"CarriedQuestItems"`
	DamageHistory          map[string]interface{} `json:"DamageHistory"`
	LastPlayerState        *float32               `json:"LastPlayerState"`
	TotalInGameTime        int                    `json:"TotalInGameTime"`
	SurvivorClass          string                 `json:"SurvivorClass"`
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
	Nickname               string                 `json:"Nickname"`
	LowerNickname          string                 `json:"LowerNickname"`
	Side                   string                 `json:"Side"`
	Voice                  string                 `json:"Voice"`
	Level                  int                    `json:"Level"`
	Experience             int                    `json:"Experience"`
	RegistrationDate       int                    `json:"RegistrationDate"`
	GameVersion            string                 `json:"GameVersion"`
	AccountType            int                    `json:"AccountType"`
	MemberCategory         int                    `json:"MemberCategory"`
	LockedMoveCommands     bool                   `json:"lockedMoveCommands"`
	SavageLockTime         int                    `json:"SavageLockTime"`
	LastTimePlayedAsSavage int                    `json:"LastTimePlayedAsSavage"`
	Settings               map[string]interface{} `json:"Settings"`
	NicknameChangeDate     int                    `json:"NicknameChangeDate"`
	NeedWipeOptions        []interface{}          `json:"NeedWipeOptions"`
	LastCompletedWipe      struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedWipe"`
	LastCompletedEvent struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedEvent"`
	BannedState             bool          `json:"BannedState"`
	BannedUntil             int           `json:"BannedUntil"`
	IsStreamerModeAvailable bool          `json:"IsStreamerModeAvailable"`
	SquadInviteRestriction  bool          `json:"SquadInviteRestriction"`
	Bans                    []interface{} `json:"Bans"`
}

type InfoSettings struct {
	Role            string  `json:"Role"`
	BotDifficulty   string  `json:"BotDifficulty"`
	Experience      int     `json:"Experience"`
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
	UpdateTime  int             `json:"UpdateTime"`
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
