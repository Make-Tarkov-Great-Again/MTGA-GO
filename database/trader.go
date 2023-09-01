package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/goccy/go-json"
)

type Trader struct {
	*structs.Trader
}

var traders = map[string]*Trader{}

func GetTraderByID(id string) *Trader {
	trader, ok := traders[id]
	if ok {
		return trader
	}
	return nil
}

func GetTraders() map[string]*Trader {
	return traders
}

func setTraders() {
	directory, err := tools.GetDirectoriesFrom(traderPath)
	if err != nil {
		panic(err)
	}

	for _, dir := range directory {
		trader := &Trader{&structs.Trader{}}

		currentTraderPath := filepath.Join(traderPath, dir)

		basePath := filepath.Join(currentTraderPath, "base.json")
		if tools.FileExist(basePath) {
			trader.Base = processBase(basePath)
		}

		assortPath := filepath.Join(currentTraderPath, "assort.json")
		if tools.FileExist(assortPath) {
			trader.Assort, trader.Index.Assort = processAssort(assortPath)
		}

		questsPath := filepath.Join(currentTraderPath, "questassort.json")
		if tools.FileExist(questsPath) {
			trader.QuestAssort = processQuestAssort(questsPath)
		}

		suitsPath := filepath.Join(currentTraderPath, "suits.json")
		if tools.FileExist(suitsPath) {
			trader.Suits, trader.Index.Suits = processSuits(suitsPath)
		}

		dialoguesPath := filepath.Join(currentTraderPath, "dialogue.json")
		if tools.FileExist(dialoguesPath) {
			trader.Dialogue = processDialogues(dialoguesPath)
		}

		traders[dir] = trader
	}
}

func processBase(basePath string) map[string]interface{} {
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

func processAssort(assortPath string) (*structs.Assort, *structs.AssortIndex) {
	var dynamic map[string]interface{}
	raw := tools.GetJSONRawMessage(assortPath)

	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	assort := &structs.Assort{}

	assort.NextResupply = 1672236024

	items, ok := dynamic["items"].([]interface{})
	if ok {
		assort.Items = make([]*structs.AssortItem, 0, len(items))
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

	index := &structs.AssortIndex{}

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

	items = nil
	index.ParentItems = parentItems
	index.Items = childlessItems

	barterSchemes, ok := dynamic["barter_scheme"].(map[string]interface{})
	if ok {
		assort.BarterScheme = make(map[string][][]*structs.Scheme)
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

func processQuestAssort(questsPath string) map[string]map[string]string {
	quests := make(map[string]map[string]string)
	raw := tools.GetJSONRawMessage(questsPath)

	err := json.Unmarshal(raw, &quests)
	if err != nil {
		panic(err)
	}

	return quests
}

func processDialogues(dialoguesPath string) map[string][]string {
	var dynamic map[string]interface{}
	raw := tools.GetJSONRawMessage(dialoguesPath)

	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}
	raw = nil

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

func processSuits(dialoguesPath string) ([]structs.TraderSuits, map[string]int8) {
	var suits []structs.TraderSuits
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

func (t *Trader) GetAssortItemByID(id string) []*structs.AssortItem {
	item, ok := t.Index.Assort.Items[id]
	if ok {
		return []*structs.AssortItem{t.Assort.Items[item]}
	}

	parentItems, parentOK := t.Index.Assort.ParentItems[id]
	if !parentOK {
		fmt.Println("Assort Item", id, "does not exist for", t.Base["nickname"])
		return nil
	}

	items := make([]*structs.AssortItem, 0, len(parentItems))
	for _, index := range parentItems {
		items = append(items, t.Assort.Items[index])
	}

	return items
}

var index = map[string]*structs.AssortIndex{}
var assorts = map[string]*structs.Assort{}
var loyaltyLevels = map[string]int8{}

func (t *Trader) GetStrippedAssort(character *structs.PlayerTemplate) *structs.Assort {
	traderID := t.Base["_id"].(string)

	cachedAssort, ok := assorts[traderID]
	if ok {
		return cachedAssort
	}

	_, ok = loyaltyLevels[traderID]
	if !ok {
		loyaltyLevels[traderID] = t.GetTraderLoyaltyLevel(character) // check loyalty level
	}
	loyaltyLevel := loyaltyLevels[traderID]

	assortIndex := structs.AssortIndex{
		Items:       map[string]int16{},
		ParentItems: map[string]map[string]int16{},
	}

	assort := structs.Assort{}

	// iterate through  loyal
	loyalLevelItems := make(map[string]int8)
	fmt.Println("Don't forget to adjust for Quests")
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

	assort.Items = make([]*structs.AssortItem, 0, len(t.Assort.Items))
	assort.BarterScheme = make(map[string][][]*structs.Scheme)

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

	assort.NextResupply = int(tools.GetCurrentTimeInSeconds() + 3600)
	index[traderID] = &assortIndex
	assorts[traderID] = &assort

	return assorts[traderID]
}

func (t *Trader) GetTraderLoyaltyLevel(character *structs.PlayerTemplate) int8 {
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
