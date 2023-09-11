package database

import (
	"MT-GO/tools"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/goccy/go-json"
)

var traders = map[string]*Trader{}

// #region Trader getters

func GetTraders() map[string]*Trader {
	return traders
}

func GetTraderByID(id string) *Trader {
	trader, ok := traders[id]
	if ok {
		return trader
	}
	return nil
}

func (t *Trader) GetAssortItemByID(id string) []*AssortItem {
	item, ok := t.Index.Assort.Items[id]
	if ok {
		return []*AssortItem{t.Assort.Items[item]}
	}

	parentItems, parentOK := t.Index.Assort.ParentItems[id]
	if !parentOK {
		fmt.Println("Assort Item", id, "does not exist for", t.Base["nickname"])
		return nil
	}

	items := make([]*AssortItem, 0, len(parentItems))
	for _, index := range parentItems {
		items = append(items, t.Assort.Items[index])
	}

	return items
}

func (t *Trader) GetStrippedAssort(character *Character) *Assort {
	traderID := t.Base["_id"].(string)

	cache := GetTraderCacheByUID(character.ID)
	cachedAssort, ok := cache.Assorts[traderID]
	if ok {
		return cachedAssort
	}

	_, ok = cache.LoyaltyLevels[traderID]
	if !ok {
		cache.LoyaltyLevels[traderID] = t.GetTraderLoyaltyLevel(character) // check loyalty level
	}
	loyaltyLevel := cache.LoyaltyLevels[traderID]

	assortIndex := AssortIndex{
		Items:       map[string]int16{},
		ParentItems: map[string]map[string]int16{},
	}

	assort := Assort{}

	// TODO: add quest checks
	loyalLevelItems := make(map[string]int8)
	for loyalID, loyalLevel := range t.Assort.LoyalLevelItems {

		if loyaltyLevel >= loyalLevel {
			loyalLevelItems[loyalID] = loyalLevel
			continue

			/* if t.QuestAssort == nil {
				loyalLevelItems[loyalID] = loyalLevel
				continue
			}

			for _, condition := range t.QuestAssort {
				if len(condition) == 0 {
					continue
				}

				for aid, qid := range condition {


				}
			} */
		}
	}

	assort.Items = make([]*AssortItem, 0, len(t.Assort.Items))
	assort.BarterScheme = make(map[string][][]*Scheme)

	var counter int16 = 0
	for itemID := range loyalLevelItems {
		index, ok := t.Index.Assort.Items[itemID]
		if ok {
			assort.BarterScheme[itemID] = t.Assort.BarterScheme[itemID]

			assortIndex.Items[itemID] = counter
			counter++
			assort.Items = append(assort.Items, t.Assort.Items[index])
		} else {
			family, ok := t.Index.Assort.ParentItems[itemID]
			if ok {
				assort.BarterScheme[itemID] = t.Assort.BarterScheme[itemID]

				assortIndex.ParentItems[itemID] = make(map[string]int16)
				for k, v := range family {
					assortIndex.ParentItems[itemID][k] = counter
					counter++
					assort.Items = append(assort.Items, t.Assort.Items[v])
				}
			}
		}
	}

	assort.NextResupply = SetResupplyTimer()

	cache.Index[traderID] = &assortIndex
	cache.Assorts[traderID] = &assort

	return cache.Assorts[traderID]
}

type ResupplyTimer struct {
	TimerResupplyTime     time.Duration
	ResupplyTimeInSeconds int
	NextResupplyTime      int
	TimerSet              bool
	Profiles              map[string]*Profile
}

var rs = &ResupplyTimer{
	TimerResupplyTime:     0,
	ResupplyTimeInSeconds: 3600, //1 hour
	NextResupplyTime:      0,
	TimerSet:              false,
	Profiles:              nil,
}

func SetResupplyTimer() int {
	if rs.TimerSet {
		return rs.NextResupplyTime
	}

	rs.NextResupplyTime = int(tools.GetCurrentTimeInSeconds()) + rs.ResupplyTimeInSeconds
	rs.TimerResupplyTime = time.Duration(rs.ResupplyTimeInSeconds) * time.Second

	rs.TimerSet = true

	go func() {
		timer := time.NewTimer(rs.TimerResupplyTime)
		for {
			<-timer.C
			rs.NextResupplyTime += rs.ResupplyTimeInSeconds
			rs.Profiles = GetProfiles()

			for _, profile := range rs.Profiles {
				traders := profile.Cache.Traders
				for _, assort := range traders.Assorts {
					assort.NextResupply = rs.NextResupplyTime
				}
			}

			timer.Reset(rs.TimerResupplyTime)
		}
	}()

	return rs.NextResupplyTime
}

// GetTraderLoyaltyLevel determines the loyalty level of a trader based on character attributes
func (t *Trader) GetTraderLoyaltyLevel(character *Character) int8 {
	loyaltyLevels := t.Base["loyaltyLevels"].([]interface{})
	traderID := t.Base["_id"].(string)

	_, ok := character.TradersInfo[traderID]
	if !ok {
		return -1
	}

	length := len(loyaltyLevels)
	for index := 0; index < length; index++ {
		loyalty := loyaltyLevels[index].(map[string]interface{})
		if character.Info.Level < int(loyalty["minLevel"].(float64)) ||
			character.TradersInfo[traderID].SalesSum < float32(loyalty["minSalesSum"].(float64)) ||
			character.TradersInfo[traderID].Standing < float32(loyalty["minStanding"].(float64)) {

			return int8(index)
		}
	}

	return int8(length)
}

// #endregion

// #region Trader setters

func setTraders() {
	directory, err := tools.GetDirectoriesFrom(traderPath)
	if err != nil {
		panic(err)
	}

	for _, dir := range directory {
		trader := &Trader{}

		currentTraderPath := filepath.Join(traderPath, dir)

		basePath := filepath.Join(currentTraderPath, "base.json")
		if tools.FileExist(basePath) {
			trader.Base = setTraderBase(basePath)
		}

		assortPath := filepath.Join(currentTraderPath, "assort.json")
		if tools.FileExist(assortPath) {
			trader.Assort, trader.Index.Assort = setTraderAssort(assortPath)
		}

		questsPath := filepath.Join(currentTraderPath, "questassort.json")
		if tools.FileExist(questsPath) {
			trader.QuestAssort = setTraderQuestAssort(questsPath)
		}

		suitsPath := filepath.Join(currentTraderPath, "suits.json")
		if tools.FileExist(suitsPath) {
			trader.Suits, trader.Index.Suits = setTraderSuits(suitsPath)
		}

		dialoguesPath := filepath.Join(currentTraderPath, "dialogue.json")
		if tools.FileExist(dialoguesPath) {
			trader.Dialogue = setTraderDialogues(dialoguesPath)
		}

		traders[dir] = trader
	}
}

func setTraderBase(basePath string) map[string]interface{} {
	base := map[string]interface{}{}

	var dynamic map[string]interface{} //here we fucking go

	raw := tools.GetJSONRawMessage(basePath)
	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	loyaltyLevels := dynamic["loyaltyLevels"].([]interface{})
	length := len(loyaltyLevels)

	for i := 0; i < length; i++ {
		level := loyaltyLevels[i].(map[string]interface{})

		insurancePriceCoef, ok := level["insurance_price_coef"].(string)
		if !ok {
			continue
		}

		level["insurance_price_coef"], err = strconv.Atoi(insurancePriceCoef)
		if err != nil {
			panic(err)
		}
	}

	repair := dynamic["repair"].(map[string]interface{})

	repairQuality, ok := repair["quality"].(string)
	if ok {
		repair["quality"], err = strconv.ParseFloat(repairQuality, 32)
		if err != nil {
			panic(err)
		}
	}

	sanitized, err := json.Marshal(dynamic)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(sanitized, &base)
	if err != nil {
		panic(err)
	}

	return base
}

func setTraderAssort(assortPath string) (*Assort, *AssortIndex) {
	var dynamic map[string]interface{}
	raw := tools.GetJSONRawMessage(assortPath)

	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	assort := &Assort{}

	assort.NextResupply = 1672236024

	items, ok := dynamic["items"].([]interface{})
	if ok {
		assort.Items = make([]*AssortItem, 0, len(items))
		data, err := json.Marshal(items)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &assort.Items)
		if err != nil {
			panic(err)
		}

	} else {
		panic("Items not found")
	}

	index := &AssortIndex{}

	parentItems := make(map[string]map[string]int16)
	childlessItems := make(map[string]int16)

	for index, item := range assort.Items {
		_, ok := childlessItems[item.ID]
		if ok {
			continue
		}

		_, ok = parentItems[item.ID]
		if ok {
			continue
		}

		itemChildren := tools.GetItemFamilyTree(items, item.ID)
		if len(itemChildren) == 1 {
			childlessItems[item.ID] = int16(index)
			continue
		}

		family := make(map[string]int16)
		for _, child := range itemChildren {
			for k, v := range assort.Items {
				if child != v.ID {
					continue
				}

				family[child] = int16(k)
				break
			}
		}
		parentItems[item.ID] = family
	}

	index.ParentItems = parentItems
	index.Items = childlessItems

	barterSchemes, ok := dynamic["barter_scheme"].(map[string]interface{})
	if ok {
		assort.BarterScheme = make(map[string][][]*Scheme)
		data, err := json.Marshal(barterSchemes)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &assort.BarterScheme)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Barter scheme not found")
	}

	loyalLevelItems, ok := dynamic["loyal_level_items"].(map[string]interface{})
	if ok {
		assort.LoyalLevelItems = map[string]int8{}
		for key, item := range loyalLevelItems {
			assort.LoyalLevelItems[key] = int8(item.(float64))
		}
	}

	data, err := json.Marshal(loyalLevelItems)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &assort.LoyalLevelItems)
	if err != nil {
		panic(err)
	}

	return assort, index
}

func setTraderQuestAssort(questsPath string) map[string]map[string]string {
	quests := make(map[string]map[string]string)
	raw := tools.GetJSONRawMessage(questsPath)

	err := json.Unmarshal(raw, &quests)
	if err != nil {
		panic(err)
	}

	return quests
}

func setTraderDialogues(dialoguesPath string) map[string][]string {
	var dynamic map[string]interface{}
	raw := tools.GetJSONRawMessage(dialoguesPath)

	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	dialogues := map[string][]string{}
	for k, v := range dynamic {
		v := v.([]interface{})

		length := len(v)
		dialogues[k] = make([]string, 0, len(v))
		if length == 0 {
			continue
		}

		for _, dialogue := range v {
			dialogues[k] = append(dialogues[k], dialogue.(string))
		}
	}

	return dialogues
}

func setTraderSuits(dialoguesPath string) ([]TraderSuits, map[string]int8) {
	var suits []TraderSuits
	raw := tools.GetJSONRawMessage(dialoguesPath)

	err := json.Unmarshal(raw, &suits)
	if err != nil {
		panic(err)
	}

	suitsIndex := make(map[string]int8)
	for index, suit := range suits {
		suitsIndex[suit.SuiteID] = int8(index)
	}

	return suits, suitsIndex
}

// #endregion Trader->Init

// #region Trader structs

type Trader struct {
	Index       TraderIndex                  `json:",omitempty"`
	Base        map[string]interface{}       `json:",omitempty"`
	Assort      *Assort                      `json:",omitempty"`
	QuestAssort map[string]map[string]string `json:",omitempty"`
	Suits       []TraderSuits                `json:",omitempty"`
	Dialogue    map[string][]string          `json:",omitempty"`
}

type TraderIndex struct {
	Assort *AssortIndex    `json:",omitempty"`
	Suits  map[string]int8 `json:",omitempty"`
}

type AssortIndex struct {
	Items       map[string]int16
	ParentItems map[string]map[string]int16 `json:",omitempty"`
}

type TraderSuits struct {
	ID           string           `json:"_id"`
	Tid          string           `json:"tid"`
	SuiteID      string           `json:"suiteId"`
	IsActive     bool             `json:"isActive"`
	Requirements SuitRequirements `json:"requirements"`
}

type SuitItemRequirements struct {
	Count          int    `json:"count"`
	Tpl            string `json:"_tpl"`
	OnlyFunctional bool   `json:"onlyFunctional"`
}

type SuitRequirements struct {
	LoyaltyLevel         int8                   `json:"loyaltyLevel"`
	ProfileLevel         int8                   `json:"profileLevel"`
	Standing             int8                   `json:"standing"`
	SkillRequirements    []interface{}          `json:"skillRequirements"`
	QuestRequirements    []string               `json:"questRequirements"`
	SuitItemRequirements []SuitItemRequirements `json:"itemRequirements"`
}

type Assort struct {
	NextResupply    int                    `json:"nextResupply"`
	BarterScheme    map[string][][]*Scheme `json:"barter_scheme"`
	Items           []*AssortItem          `json:"items"`
	LoyalLevelItems map[string]int8        `json:"loyal_level_items"`
}

type AssortItem struct {
	ID       string        `json:"_id"`
	Tpl      string        `json:"_tpl"`
	ParentID string        `json:"parentId"`
	SlotID   string        `json:"slotId"`
	Upd      AssortItemUpd `json:"upd,omitempty"`
}

type AssortItemUpd struct {
	BuyRestrictionCurrent interface{} `json:"BuyRestrictionCurrent,omitempty"`
	BuyRestrictionMax     interface{} `json:"BuyRestrictionMax,omitempty"`
	StackObjectsCount     int         `json:"StackObjectsCount,omitempty"`
	UnlimitedCount        bool        `json:"UnlimitedCount,omitempty"`
	FireMode              FireMode    `json:"FireMode,omitempty"`
	Foldable              Foldable    `json:"Foldable,omitempty"`
}

type FireMode struct {
	FireMode string `json:"FireMode"`
}
type Foldable struct {
	Folded bool `json:"Folded,omitempty"`
}

type Scheme struct {
	Tpl   string  `json:"_tpl"`
	Count float32 `json:"count"`
}

// #endregion
