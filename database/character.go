package database

import (
	"log"
	"path/filepath"
	"strings"

	"MT-GO/services"
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

// #region Character getters

func GetCharacterByID(uid string) *Character {
	if profile, ok := profiles[uid]; ok {
		return profile.Character
	}

	log.Println("Profile with UID ", uid, " does not have a character")
	return nil
}

func (c *Character) GetQuestsAvailableToPlayer() ([]any, error) {
	var output []any

	query := GetQuestsQuery()

	cachedQuests, err := GetQuestCacheByID(c.ID)
	if err != nil {
		return nil, err
	}

	characterHasQuests := len(cachedQuests.Index) != 0

	for key, value := range query {

		if services.CheckIfQuestForOtherFaction(c.Info.Side, key) {
			continue
		}

		if strings.Contains(value.Name, "-Event") {
			//TODO: filter events properly
			continue
		}

		if value.Conditions == nil || value.Conditions.AvailableForStart == nil {
			output = append(output, GetQuestByID(key))
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
			output = append(output, GetQuestByID(key))
			continue
		}

		loyaltyCheck := false
		if forStart.TraderLoyalty != nil {
			for trader, loyalty := range forStart.TraderLoyalty {

				traderInfo, ok := c.TradersInfo[trader]
				if !ok || traderInfo.LoyaltyLevel == 0 {
					if data, err := GetTraderByUID(trader); err != nil {
						return nil, err
					} else {
						data.SetTraderLoyaltyLevel(c)
					}
				}

				loyaltyCheck = services.LevelComparisonCheck(
					loyalty.Level,
					float64(traderInfo.LoyaltyLevel),
					loyalty.CompareMethod,
				)
			}

			if !loyaltyCheck {
				continue
			}
		}

		standingCheck := false
		if forStart.TraderStanding != nil {
			for trader, loyalty := range forStart.TraderStanding {

				traderInfo, ok := c.TradersInfo[trader]
				if !ok || traderInfo.LoyaltyLevel == 0 {
					if data, err := GetTraderByUID(trader); err != nil {
						return nil, err
					} else {
						data.SetTraderLoyaltyLevel(c)
					}
				}

				standingCheck = services.LevelComparisonCheck(
					loyalty.Level,
					float64(traderInfo.LoyaltyLevel),
					loyalty.CompareMethod,
				)
			}

			if !standingCheck {
				continue
			}
		}

		if forStart.Quest != nil && characterHasQuests {
			if c.IsPreviousQuestComplete(forStart.Quest, cachedQuests) {
				output = append(output, GetQuestByID(key))
				continue
			}
		}
	}

	return output, nil
}

func (c *Character) IsPreviousQuestComplete(quests map[string]*QuestCondition, cachedQuests *QuestCache) bool {
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
		log.Println("Removed", cleaned, "modded item(s) from your inventory")
		inv.Items = newItems
		return true
	}
	return false

}

func (c *Character) SaveCharacter() {
	characterFilePath := filepath.Join(profilesPath, c.ID, "character.json")

	err := tools.WriteToFile(characterFilePath, c)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Character saved")
}

/*func (c *Character) SaveCharacter(sessionID string) {
	characterFilePath := filepath.Join(profilesPath, sessionID, "character.json")

	err := tools.WriteToFile(characterFilePath, c)
	if err != nil {
			log.Println(err)
			return
		}
	log.Println("Character saved")
}*/

// QuestAccept updates an existing Accepted quest, or creates and appends new Accepted Quest to cache and Character
func (c *Character) QuestAccept(qid string, profileChangesEvent *ProfileChangesEvent) {

	cachedQuests, err := GetQuestCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}
	length := len(cachedQuests.Index)
	time := int(tools.GetCurrentTimeInSeconds())

	query := GetQuestFromQueryByID(qid)

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
		log.Println("There are rewards heeyrrrr!")
		log.Println(profileChangesEvent.ProfileChanges[c.ID].ID)

		// TODO: Apply then Get Quest rewards and then route messages from there
		// Character.ApplyQuestRewardsToCharacter()  applies the given reward
		// Quests.GetQuestReward() returns the given reward
		// CreateNPCMessageWithReward()
	}

	dialogue, err := GetDialogueByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	dialog, message := CreateQuestDialogue(c.ID, "QuestStart", query.Trader, query.Dialogue.Description)
	dialog.New++
	dialog.Messages = append(dialog.Messages, *message)

	(*dialogue)[query.Trader] = dialog

	notification := CreateNotification(message)

	connection := GetConnection(c.ID)
	if connection == nil {
		log.Println("Can't send message to character because connection is nil, storing...")
		storage, err := GetStorageByID(c.ID)
		if err != nil {
			log.Println(err)
			return
		}

		storage.Mailbox = append(storage.Mailbox, notification)
		storage.SaveStorage(c.ID)
	} else {
		connection.SendMessage(notification)
	}

	//TODO: Get new player quests from database now that we've accepted one
	quests, err := c.GetQuestsAvailableToPlayer()
	if err != nil {
		log.Println(err)
		return
	}

	profileChangesEvent.ProfileChanges[c.ID].Quests = quests
	dialogue.SaveDialogue(c.ID)
	c.SaveCharacter()
}

func (c *Character) ApplyQuestRewardsToCharacter(rewards *QuestRewards) {
	log.Println()
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
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &examine); err != nil {
		log.Println(err)
		return
	}

	var item *DatabaseItem
	if examine.FromOwner == nil {
		log.Println("Examining Item from Player Inventory")
		cache, err := GetInventoryCacheByID(c.ID)
		if err != nil {
			log.Println(err)
			return
		}

		if index := cache.GetIndexOfItemByID(examine.Item); index != nil {
			itemInInventory := c.Inventory.Items[*index]
			item = GetItemByID(itemInInventory.TPL)
		} else {
			log.Println("[EXAMINE] Examining Item", examine.Item, " from Player Inventory failed, does not exist!")
			return
		}
	} else {
		switch examine.FromOwner.Type {
		case "Trader":
			data, err := GetTraderByUID(examine.FromOwner.ID)
			if err != nil {
				log.Println(err)
				return
			}

			assortItem := data.GetAssortItemByID(examine.Item)
			item = GetItemByID(assortItem[0].Tpl)

		case "HideoutUpgrade":
		case "HideoutProduction":
		case "ScavCase":
			item = GetItemByID(examine.Item)

		case "RagFair":
		default:
			log.Println("[EXAMINE] FromOwner.Type: ", examine.FromOwner.Type, "is not supported, returning...")
			return
		}
	}

	if item == nil {
		log.Println("[EXAMINE] Examining Item", examine.Item, "failed, does not exist in Item Database")
		return
	}

	c.Encyclopedia[item.ID] = true
	log.Println("[EXAMINE] Encyclopedia entry added for", item.ID)

	//add experience
	experience, ok := item.Props["ExamineExperience"].(float64)
	if !ok {
		log.Println("[EXAMINE] Item", examine.Item, "does not have ExamineExperience property, returning...")
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
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &move); err != nil {
		log.Println(err)
		return
	}

	cache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	index := cache.GetIndexOfItemByID(move.Item)
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
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &swap); err != nil {
		log.Println(err)
		return
	}

	cache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	index := *cache.GetIndexOfItemByID(swap.Item)
	itemInInventory := &c.Inventory.Items[index]

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

	index = cache.Lookup.Forward[swap.Item2]
	itemInInventory = &c.Inventory.Items[index]

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
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &fold); err != nil {
		log.Println(err)
		return
	}

	inventoryCache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	index := inventoryCache.GetIndexOfItemByID(fold.Item)
	if index == nil {
		log.Println("Item", fold.Item, "does not exist in cache!")
		return
	}
	itemInInventory := &c.Inventory.Items[*index]
	if itemInInventory.UPD == nil || itemInInventory.UPD.Foldable == nil {
		log.Println(itemInInventory.ID, "cannot be folded!")
		return
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
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &readEncyclopedia); err != nil {
		log.Println(err)
		return
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
	data, err := json.Marshal(moveAction)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(data, &merge); err != nil {
		log.Println(err)
		return
	}

	inventoryCache, err := GetInventoryCacheByID(c.ID)
	if err != nil {
		log.Println(err)
		return
	}

	toMergeIndex := *inventoryCache.GetIndexOfItemByID(merge.Item)
	toMerge := &c.Inventory.Items[toMergeIndex]

	mergeWithIndex := *inventoryCache.GetIndexOfItemByID(merge.With)
	mergeWith := c.Inventory.Items[mergeWithIndex]

	mergeWith.UPD.StackObjectsCount += toMerge.UPD.StackObjectsCount

	inventoryCache.ClearItemFromContainer(toMerge.ID)
	c.Inventory.RemoveSingleItemFromInventoryByIndex(toMergeIndex)
	inventoryCache.SetInventoryIndex(&c.Inventory)

	profileChangesEvent.ProfileChanges[c.ID].Items.Change = append(profileChangesEvent.ProfileChanges[c.ID].Items.Change, mergeWith)
	profileChangesEvent.ProfileChanges[c.ID].Items.Del = append(profileChangesEvent.ProfileChanges[c.ID].Items.Del, InventoryItem{ID: toMerge.ID})
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
		Notes                [][]any `json:"Notes"`
		TransactionInProcess struct {
			HasCheckChanges bool `json:"HasCheckChanges"`
			HasHandlers     bool `json:"HasHandlers"`
		} `json:"TransactionInProcess,omitempty"`
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
	Unlocked     bool    `json:"unlocked"`
	Disabled     bool    `json:"disabled"`
	SalesSum     float32 `json:"salesSum"`
	Standing     float32 `json:"standing"`
	LoyaltyLevel int8    `json:"loyaltyLevel"`
	NextResupply int     `json:"nextResupply"`
}

type PlayerRagfairInfo struct {
	Rating          float32 `json:"rating"`
	IsRatingGrowing bool    `json:"isRatingGrowing"`
	Offers          []any   `json:"offers"`
}

type PlayerHideoutArea struct {
	AreaType              string `json:"AreaType,omitempty"`
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
	Seed        int                 `json:"Seed,omitempty"`
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
	ID              string `json:"id"`
	Type            string `json:"type"`
	TemplateID      string `json:"templateId"`
	IsPositive      bool   `json:"IsPositive,omitempty"`
	LocalizationKey string `json:"LocalizationKey,omitempty"`
	Passive         bool   `json:"passive,omitempty"`
	Production      bool   `json:"production,omitempty"`
	Value           int    `json:"value,omitempty"`
	Visible         bool   `json:"visible,omitempty"`
}

type HealthInfo struct {
	Hydration   CurrMaxHealth   `json:"Hydration"`
	Energy      CurrMaxHealth   `json:"Energy"`
	Temperature CurrMaxHealth   `json:"Temperature"`
	Poison      CurrMaxHealth   `json:"Poison,omitempty"`
	BodyParts   BodyPartsHealth `json:"BodyParts"`
	UpdateTime  int32           `json:"UpdateTime"`
}

type HealthOf struct {
	Effects map[string]any `json:"Effects,omitempty"`
	Health  CurrMaxHealth  `json:"Health"`
}

type CurrMaxHealth struct {
	Current                      float32 `json:"Current"`
	Maximum                      float32 `json:"Maximum"`
	Minimum                      float32 `json:"Minimum,omitempty"`
	OverDamageReceivedMultiplier float32 `json:"OverDamageReceivedMultiplier,omitempty"`
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
