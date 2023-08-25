package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// var quests map[string]interface{}
var quests = map[string]*structs.Quest{}

func GetQuests() map[string]*structs.Quest {
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

	dynamic := make(map[string]map[string]interface{})
	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	for k, v := range dynamic {
		var quest = &structs.Quest{}

		quest.Dialogue = processDialogue(v)

		questConditions, ok := v["conditions"].(map[string]interface{})
		if ok {
			conditions := &structs.QuestAvailabilityConditions{}
			process := processConditions(questConditions)

			data, err := json.Marshal(process)
			if err != nil {
				fmt.Println(err)
			}
			err = json.Unmarshal(data, conditions)
			if err != nil {
				fmt.Println(err)
			}

			empty := structs.QuestAvailabilityConditions{}
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
			rewards := &structs.QuestRewardAvailabilityConditions{}
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

			/* 			empty := &structs.QuestRewardAvailabilityConditions{}
			   			if rewards != empty {
			   			} */

		}

		quests[k] = quest
	}

	_ = tools.WriteToFile("questsDatabase.json", quests)
}

func processDialogue(quest map[string]interface{}) structs.QuestDialogues {
	dialogues := &structs.QuestDialogues{}

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
		levelCondition := &structs.LevelCondition{}
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
		condition := &structs.QuestCondition{}
		questID, _ := props["id"].(string)

		previousQuestID, _ := props["target"].(string)
		condition.PreviousQuestID = previousQuestID

		value, _ := props["status"].([]interface{})[0].(float64)
		condition.Status = int(value)

		output[questID] = condition
		return output

	case "TraderLoyalty", "TraderStanding":
		traderID, _ := props["target"].(string)

		levelCondition := &structs.LevelCondition{}
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

		handover := &structs.HandoverCondition{}

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
		levelCondition := &structs.LevelCondition{}

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
		questRewardItem := &structs.QuestRewardItem{}

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
		scheme := &structs.QuestRewardProductionScheme{}

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

func _setQuests() {

	raw := tools.GetJSONRawMessage(questsPath)

	dynamic := make(map[string]interface{})
	err := json.Unmarshal(raw, &dynamic)
	if err != nil {
		panic(err)
	}

	for _, v := range dynamic {
		quest := v.(map[string]interface{})

		conditions, ok := quest["conditions"].(map[string]interface{})
		if !ok {
			panic("quests.conditions is not a map")
		}

		for _, v := range conditions {
			conditionType := v.([]interface{})

			for _, v := range conditionType {
				condition := v.(map[string]interface{})

				properties, ok := condition[props].(map[string]interface{})
				if !ok {
					continue
				}

				value, ok := properties["value"].(string)
				if ok {
					properties["value"], err = strconv.ParseFloat(value, 32)
					if err != nil {
						panic(err)
					}
				}

				counter, ok := properties["counter"].(map[string]interface{})
				if !ok {
					continue
				}

				conditions := counter["conditions"].([]interface{})

				for _, v := range conditions {
					condition := v.(map[string]interface{})

					properties, ok := condition[props].(map[string]interface{})
					if !ok {
						continue
					}

					value, ok := properties["value"].(string)
					if !ok {
						continue
					}

					properties["value"], err = strconv.ParseFloat(value, 32)
					if err != nil {
						panic(err)
					}
				}

			}
		}

		rewards, ok := quest["rewards"].(map[string]interface{})
		if !ok {
			panic("quests.rewards is not a map")
		}

		for _, v := range rewards {
			rewardType := v.([]interface{})

			for _, v := range rewardType {
				reward := v.(map[string]interface{})

				value, ok := reward["value"].(string)
				if ok {
					value = strings.TrimSpace(value)
					reward["value"], err = strconv.ParseFloat(value, 32)
					if err != nil {
						panic(err)
					}
				}
			}
		}

	}

	jsonData, err := json.Marshal(dynamic)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &quests)
	if err != nil {
		panic(err)
	}
}
