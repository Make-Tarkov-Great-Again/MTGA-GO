package database

import (
	"MT-GO/tools"
	"fmt"
	"strconv"

	"github.com/goccy/go-json"
)

// var quests map[string]interface{}
var questsQuery = map[string]*Quest{}
var quests = map[string]interface{}{}

func GetQuestsQuery() map[string]*Quest {
	return questsQuery
}

func GetQuestByQID(qid string) interface{} {
	return quests[qid]
}

func GetQuests() map[string]interface{} {
	return quests
}

const (
	ForFinish string = "AvailableForFinish"
	ForStart  string = "AvailableForStart"
	Fail      string = "Fail"
	Started   string = "Started"
	Success   string = "Success"
)

func setQuests() {

	raw := tools.GetJSONRawMessage(questsPath)
	err := json.Unmarshal(raw, &quests)
	if err != nil {
		panic(err)
	}

	dynamic := make(map[string]map[string]interface{})
	err = json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	for k, v := range dynamic {
		var quest = &Quest{}

		quest.Name = v["QuestName"].(string)
		quest.Dialogue = processDialogue(v)

		questConditions, ok := v["conditions"].(map[string]interface{})
		if ok {
			conditions := &QuestAvailabilityConditions{}
			process := processConditions(questConditions)

			data, err := json.Marshal(process)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(data, conditions)
			if err != nil {
				fmt.Println(err)
			}

			empty := QuestAvailabilityConditions{}
			if *conditions == empty {
				quest.Conditions = nil
			} else {
				quest.Conditions = conditions
			}
		} else {
			fmt.Println("Conditions don't exist? Check " + k)
		}

		questRewards, ok := v["rewards"].(map[string]interface{})
		if ok {
			rewards := &QuestRewardAvailabilityConditions{}
			process := processQuestRewards(questRewards)
			data, err := json.Marshal(process)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(data, rewards)
			if err != nil {
				fmt.Println(err)
			}
			quest.Rewards = *rewards

			/* 			empty := &QuestRewardAvailabilityConditions{}
			   			if rewards != empty {
			   			} */

		}

		questsQuery[k] = quest
	}
	//_ = tools.WriteToFile("questsDatabase.json", questsQuery)
}

func processDialogue(quest map[string]interface{}) QuestDialogues {
	dialogues := &QuestDialogues{}

	description, _ := quest["description"].(string)
	dialogues.Description = description

	complete, _ := quest["completePlayerMessage"].(string)
	dialogues.Complete = complete

	fail, _ := quest["failMessageText"].(string)
	dialogues.Fail = fail

	started, _ := quest["startedMessageText"].(string)
	dialogues.Started = started

	success, _ := quest["successMessageText"].(string)
	dialogues.Success = success

	accepted, _ := quest["acceptPlayerMessage"].(string)
	dialogues.Accepted = accepted

	return *dialogues
}

const (
	parent string = "_parent"
	props  string = "_props"
)

func processConditions(conditions map[string]interface{}) map[string]map[string]interface{} {
	output := make(map[string]map[string]interface{})

	processCategory := func(category string, conditionList []interface{}) {
		processedCategory := make(map[string]interface{})

		for _, condition := range conditionList {
			conditionMap, ok := condition.(map[string]interface{})
			if !ok || len(conditionMap) == 0 {
				continue
			}

			name, _ := conditionMap[parent].(string)
			if name == "FindItem" || name == "CounterCreator" ||
				name == "PlaceBeacon" || name == "LeaveItemAtLocation" {
				continue
			}

			process := processCondition(name, conditionMap)
			processedCategory[name] = process

		}

		if len(processedCategory) > 0 {
			output[category] = processedCategory
		}
	}

	fails, _ := conditions[Fail].([]interface{})
	processCategory(Fail, fails)

	starts, _ := conditions[ForStart].([]interface{})
	processCategory(ForStart, starts)

	successes, _ := conditions[ForFinish].([]interface{})
	processCategory(ForFinish, successes)

	return output
}

func processCondition(name string, conditions map[string]interface{}) interface{} {

	output := make(map[string]interface{})
	props, _ := conditions[props].(map[string]interface{})

	switch name {
	case "Level":
		levelCondition := &LevelCondition{}
		compare, _ := props["compareMethod"].(string)
		levelCondition.CompareMethod = compare

		float, ok := props["value"].(float64)
		if ok {
			levelCondition.Level = float
			return levelCondition
		}

		str, ok := props["value"].(string)
		if ok {
			level, _ := strconv.ParseFloat(str, 64)
			levelCondition.Level = level
			return levelCondition
		}

	case "Quest":
		condition := &QuestCondition{}
		questID, _ := props["id"].(string)

		previousQuestID, _ := props["target"].(string)
		condition.PreviousQuestID = previousQuestID

		value, _ := props["status"].([]interface{})[0].(float64)
		condition.Status = int(value)

		output[questID] = condition
		return output

	case "TraderLoyalty", "TraderStanding":
		traderID, _ := props["target"].(string)

		levelCondition := &LevelCondition{}
		compare, _ := props["compareMethod"].(string)
		levelCondition.CompareMethod = compare

		float, ok := props["value"].(float64)
		if ok {
			levelCondition.Level = float
			output[traderID] = levelCondition
			return output
		}

		str, ok := props["value"].(string)
		if ok {
			level, _ := strconv.ParseFloat(str, 64)
			levelCondition.Level = level
			output[traderID] = levelCondition
			return output
		}

		return output

	case "HandoverItem", "WeaponAssembly", "FindItem":

		handover := &HandoverCondition{}

		handoverID, _ := props["id"].(string)

		itemID, _ := props["target"].([]interface{})[0].(string)
		handover.ItemToHandover = itemID

		ifString, _ := props["value"].(string)
		value, _ := strconv.Atoi(ifString)
		handover.Amount = value

		output[handoverID] = handover
		return output

	case "Skill":
		skillName, _ := props["target"].(string)
		levelCondition := &LevelCondition{}

		float, ok := props["value"].(float64)
		if ok {
			levelCondition.Level = float
		}

		str, ok := props["value"].(string)
		if ok {
			level, _ := strconv.ParseFloat(str, 64)
			levelCondition.Level = level
		}

		compare, _ := props["compareMethod"].(string)
		levelCondition.CompareMethod = compare

		output[skillName] = levelCondition
		return output

	default:
		fmt.Println(name + " condition, probably not needed")
	}
	return output
}

const (
	_type string = "type"
)

func processQuestRewards(rewards map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	fails, ok := rewards[Fail].([]interface{})
	if ok && len(fails) != 0 {
		succ := make(map[string]interface{})

		for _, fail := range fails {
			if len(fail.(map[string]interface{})) == 0 {
				continue
			}
			reward := fail.(map[string]interface{})

			name := reward[_type].(string)
			succ[name] = processQuestReward(name, reward)
		}
		output[Fail] = succ
	}

	starts, ok := rewards[Started].([]interface{})
	if ok && len(starts) != 0 {
		succ := make(map[string]interface{})

		for _, start := range starts {
			if len(start.(map[string]interface{})) == 0 {
				continue
			}
			reward := start.(map[string]interface{})

			name := reward[_type].(string)
			succ[name] = processQuestReward(name, reward)
		}
		output[Started] = succ
	}

	successes, ok := rewards[Success].([]interface{})
	if ok && len(successes) != 0 {
		succ := make(map[string]interface{})

		for _, success := range successes {
			if len(success.(map[string]interface{})) == 0 {
				continue
			}

			reward := success.(map[string]interface{})

			name := reward[_type].(string)
			succ[name] = processQuestReward(name, reward)
		}
		output[Success] = succ
	}
	return output
}

func processQuestReward(name string, reward map[string]interface{}) interface{} {
	output := make(map[string]interface{})

	switch name {
	case "Experience":
		float, ok := reward["value"].(float64)
		if ok {
			return int(float)
		}

		str, ok := reward["value"].(string)
		if ok {
			exp, _ := strconv.Atoi(str)
			return exp
		}
	case "Item":
		questRewardItem := &QuestRewardItem{}

		itemID, _ := reward["target"].(string)

		idems, _ := reward["items"].([]interface{})
		questRewardItem.Items = make([]map[string]interface{}, 0, len(idems))
		for _, idem := range idems {
			idem := idem.(map[string]interface{})
			questRewardItem.Items = append(questRewardItem.Items, idem)
		}

		float, ok := reward["value"].(float64)
		if ok {
			questRewardItem.Value = int(float)
		}

		str, ok := reward["value"].(string)
		if ok {
			value, _ := strconv.Atoi(str)
			questRewardItem.Value = value
		}

		output[itemID] = questRewardItem
	case "AssortmentUnlock":
		traderID, _ := reward["target"].(string)
		return traderID
	case "TraderStanding", "TraderStandingRestore":

		traderID, _ := reward["target"].(string)

		float, ok := reward["value"].(float64)
		if ok {
			output[traderID] = float32(float)
			return output
		}

		str, ok := reward["value"].(string)
		if ok {
			value, _ := strconv.Atoi(str)
			output[traderID] = float32(value)
			return output
		}

		return output
	case "Skill":
		skillName, _ := reward["target"].(string)

		float, ok := reward["value"].(float64)
		if ok {
			output[skillName] = float32(float)
		}

		str, ok := reward["value"].(string)
		if ok {
			value, _ := strconv.Atoi(str)
			output[skillName] = float32(value)
		}

		return output
	case "ProductionScheme":
		schemeID, _ := reward["target"].(string)
		scheme := &QuestRewardProductionScheme{}

		loyaltyLevel, _ := reward["loyaltyLevel"].(float64)
		scheme.LoyaltyLevel = int(loyaltyLevel)

		ifString, ok := reward["traderId"].(string)
		if ok {
			area, _ := strconv.Atoi(ifString)
			scheme.AreaID = area
		}

		ifInt, ok := reward["traderId"].(int)
		if ok {
			scheme.AreaID = ifInt
		}

		item, ok := reward["items"].([]interface{})[0].(map[string]interface{})["_tpl"].(string)
		if ok {
			scheme.Item = item
		}

		output[schemeID] = scheme
	case "TraderUnlock":
		traderID, _ := reward["target"].(string)
		return traderID
	default:
		fmt.Println(name + " reward")
	}

	return output
}

// #region Quest structs
type Quest struct {
	Name       string
	Dialogue   QuestDialogues                    `json:",omitempty"`
	Conditions *QuestAvailabilityConditions      `json:",omitempty"`
	Rewards    QuestRewardAvailabilityConditions `json:",omitempty"`
}

type QuestDialogues struct {
	Description string
	Accepted    string
	Started     string
	Complete    string
	Success     string
	Fail        string
}

type QuestAvailabilityConditions struct {
	AvailableForStart  *QuestConditionTypes `json:"AvailableForStart,omitempty"`
	AvailableForFinish *QuestConditionTypes `json:"AvailableForFinish,omitempty"`
	Fail               *QuestConditionTypes `json:"Fail,omitempty"`
}

type QuestConditionTypes struct {
	Level          *LevelCondition               `json:"Level,omitempty"`
	Quest          map[string]*QuestCondition    `json:"Quest,omitempty"`
	TraderLoyalty  map[string]*LevelCondition    `json:"TraderLoyalty,omitempty"`
	TraderStanding map[string]*LevelCondition    `json:"TraderStanding,omitempty"`
	HandoverItem   map[string]*HandoverCondition `json:"HandoverItem,omitempty"`
	WeaponAssembly map[string]*HandoverCondition `json:"WeaponAssembly,omitempty"`
	FindItem       map[string]*HandoverCondition `json:"FindItem,omitempty"`
	Skills         map[string]*LevelCondition    `json:"Skill,omitempty"`
}

type HandoverCondition struct {
	ItemToHandover string
	Amount         int
}

type QuestCondition struct {
	Status          int
	PreviousQuestID string
}

type LevelCondition struct {
	CompareMethod string
	Level         float64
}

type QuestRewardAvailabilityConditions struct {
	Start   *QuestRewards `json:"Started,omitempty"`
	Success *QuestRewards `json:"Success,omitempty"`
	Fail    *QuestRewards `json:"Fail,omitempty"`
}

type QuestRewards struct {
	Experience            int                                     `json:"Experience,omitempty"`
	Items                 map[string]*QuestRewardItem             `json:"Item,omitempty"`
	AssortmentUnlock      string                                  `json:"AssortmentUnlock,omitempty"`
	TraderStanding        map[string]*float64                     `json:"TraderStanding,omitempty"`
	TraderStandingRestore map[string]*float64                     `json:"TraderStandingRestore,omitempty"`
	TraderUnlock          string                                  `json:"TraderUnlock,omitempty"`
	Skills                map[string]*int                         `json:"Skills,omitempty"`
	ProductionScheme      map[string]*QuestRewardProductionScheme `json:"ProductionScheme,omitempty"`
}

type QuestRewardProductionScheme struct {
	Item         string
	LoyaltyLevel int
	AreaID       int
}

type QuestRewardItem struct {
	FindInRaid bool
	Items      []map[string]interface{}
	Value      int
}

type QuestRewardAssortUnlock struct {
	Items        []map[string]interface{}
	LoyaltyLevel int
}

// #endregion
