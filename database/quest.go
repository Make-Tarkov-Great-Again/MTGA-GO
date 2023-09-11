package database

import (
	"MT-GO/tools"
	"fmt"
	"strconv"

	"github.com/goccy/go-json"
)

var questsQuery = map[string]*Quest{}
var quests = map[string]interface{}{}

// #region Quests getters

func GetQuestsQuery() map[string]*Quest {
	return questsQuery
}

func GetQuestFromQueryByQID(qid string) *Quest {
	query, ok := questsQuery[qid]
	if !ok {
		fmt.Println("Quest", qid, "does not exist in quests query")
		return nil
	}
	return query
}

func GetQuestByQID(qid string) interface{} {
	quest, ok := quests[qid]
	if !ok {
		fmt.Println("Quest", qid, "does not exist in quests")
		return nil
	}
	return quest
}

func GetQuests() map[string]interface{} {
	return quests
}

// #endregion

// region Quest setters
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
		quest.Trader = v["traderId"].(string)
		quest.Dialogue = setQuestDialogue(v)

		questConditions, ok := v["conditions"].(map[string]interface{})
		if ok {
			conditions := &QuestAvailabilityConditions{}
			process := setQuestConditions(questConditions)

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
			process := setQuestRewards(questRewards)
			data, err := json.Marshal(process)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(data, rewards)
			if err != nil {
				fmt.Println(err)
			}
			quest.Rewards = *rewards
		}

		questsQuery[k] = quest
	}
	//_ = tools.WriteToFile("questsDatabase.json", questsQuery)
}

func setQuestDialogue(quest map[string]interface{}) QuestDialogues {
	dialogues := new(QuestDialogues)

	description, ok := quest["description"].(string)
	if !ok {
		fmt.Println("quest[`description`]")
	}
	dialogues.Description = description

	//TODO: remove if not needed
	/* 	complete, ok := quest["completePlayerMessage"].(string)
	   	if !ok {
	   		fmt.Println("quest[`completePlayerMessage`]")
	   	}
	   	dialogues.Complete = complete */

	fail, ok := quest["failMessageText"].(string)
	if !ok {
		fmt.Println("quest[`failMessageText`]")
	}
	dialogues.Fail = fail

	started, ok := quest["startedMessageText"].(string)
	if !ok {
		fmt.Println("quest[`startedMessageText`]")
	}
	dialogues.Started = started

	success, ok := quest["successMessageText"].(string)
	if !ok {
		fmt.Println("quest[`successMessageText`]")
	}
	dialogues.Success = success

	//TODO: remove if not needed
	/* 	accepted, ok := quest["acceptPlayerMessage"].(string)
	   	if !ok {
	   		fmt.Println("quest[`acceptPlayerMessage`]")
	   	}
	   	dialogues.Accepted = accepted */

	return *dialogues
}

const (
	parent string = "_parent"
	props  string = "_props"
)

func setQuestConditions(conditions map[string]interface{}) map[string]map[string]interface{} {
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

			process := processQuestCondition(name, conditionMap)
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

var QuestStatus = map[int8]string{
	0: "Locked",
	1: "AvailableForStart",
	2: "Started",
	3: "AvailableForFinish",
	4: "Success",
	5: "Fail",
	6: "FailRestartable",
	7: "MarkedAsFailed",
	8: "Expired",
	9: "AvailableAfter",
}

func processQuestCondition(name string, conditions map[string]interface{}) interface{} {

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

		isFloat, ok := props["availableAfter"].(float64)
		if ok {
			condition.AvailableAfter = int(isFloat)
		}

		previousQuestID, _ := props["target"].(string)
		condition.PreviousQuestID = previousQuestID

		value, _ := props["status"].([]interface{})[0].(float64)
		condition.Status = QuestStatus[int8(value)]

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

func setQuestRewards(rewards map[string]interface{}) map[string]interface{} {
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
			succ[name] = setQuestReward(name, reward)
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
			succ[name] = setQuestReward(name, reward)
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
			succ[name] = setQuestReward(name, reward)
		}
		output[Success] = succ
	}
	return output
}

func setQuestReward(name string, reward map[string]interface{}) interface{} {
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
			output[traderID] = float
			return output
		}

		str, ok := reward["value"].(string)
		if ok {
			value, _ := strconv.ParseFloat(str, 64)
			output[traderID] = value
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

// #endregion

// #region Quest functions

func CompletedPreviousQuestCheck(quests map[string]*QuestCondition, cachedQuests *QuestCache) bool {
	previousQuestCompleted := false
	for _, v := range quests {
		quest, ok := cachedQuests.Quests[v.PreviousQuestID]
		if !ok {
			continue
		}

		previousQuestCompleted = v.Status == quest.Status
	}
	return previousQuestCompleted
}

// #endregion

// #region Quest structs

type Quest struct {
	Name       string
	Trader     string
	Dialogue   QuestDialogues                    `json:",omitempty"`
	Conditions *QuestAvailabilityConditions      `json:",omitempty"`
	Rewards    QuestRewardAvailabilityConditions `json:",omitempty"`
}

type QuestDialogues struct {
	Description string
	//Accepted    string
	Started string
	//Complete    string
	Success string
	Fail    string
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
	Status          string
	PreviousQuestID string
	AvailableAfter  int `json:",omitempty"`
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
