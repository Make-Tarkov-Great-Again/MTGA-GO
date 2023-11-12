package data

import (
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
)

var questsQuery = make(map[string]*Quest)
var quests = make(map[string]map[string]any)

// #region Quests getters

func GetQuestsQuery() map[string]*Quest {
	return questsQuery
}

func GetQuestFromQueryByID(qid string) *Quest {
	if query, ok := questsQuery[qid]; !ok {
		log.Println("Quest", qid, "does not exist in quests query")
		return nil
	} else {
		return query
	}
}

func GetQuestByID(qid string) any {
	if quest, ok := quests[qid]; !ok {
		log.Println("Quest", qid, "does not exist in quests")
		return nil
	} else {
		return quest
	}
}

// #endregion

// region Quest setters
const (
	parent    string = "_parent"
	props     string = "_props"
	_type     string = "type"
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
		log.Println(err)
		return
	}
}

func IndexQuests() {
	for k, v := range quests {
		done := make(chan bool)
		quest := &Quest{
			Name:   v["QuestName"].(string),
			Trader: v["traderId"].(string),
		}

		go func() {
			quest.Dialogue = setQuestDialogue(v)
			done <- true
		}()

		go func() {
			if questConditions, ok := v["conditions"].(map[string]any); ok {
				quest.Conditions = setQuestConditions(questConditions)
			}
			done <- true
		}()

		go func() {
			if questRewards, _ := v["rewards"].(map[string]any); questRewards != nil {
				quest.Rewards = setQuestRewards(questRewards)
			}
			done <- true
		}()

		for i := int8(0); i < 3; i++ {
			<-done
		}
		questsQuery[k] = quest
	}
}

func setQuestDialogue(quest map[string]any) QuestDialogues {
	return QuestDialogues{
		Description: quest["description"].(string),
		Started:     quest["startedMessageText"].(string),
		Success:     quest["successMessageText"].(string),
		Fail:        quest["failMessageText"].(string),
	}
	/*dialogues := new(QuestDialogues)

		description, ok := quest["description"].(string)
		if !ok {
			log.Println("quest[`description`]")
		}
	dialogues.Description = description*/

	//TODO: remove if not needed
	/* 	complete, ok := quest["completePlayerMessage"].(string)
	   	if !ok {
	   		log.Println("quest[`completePlayerMessage`]")
	   	}
	   	dialogues.Complete = complete */

	/*fail, ok := quest["failMessageText"].(string)
	if !ok {
		log.Println("quest[`failMessageText`]")
	}
	dialogues.Fail = fail

	started, ok := quest["startedMessageText"].(string)
	if !ok {
		log.Println("quest[`startedMessageText`]")
	}
	dialogues.Started = started

	success, ok := quest["successMessageText"].(string)
	if !ok {
		log.Println("quest[`successMessageText`]")
	}
	dialogues.Success = success*/

	//TODO: remove if not needed
	/* 	accepted, ok := quest["acceptPlayerMessage"].(string)
	   	if !ok {
	   		log.Println("quest[`acceptPlayerMessage`]")
	   	}
	   	dialogues.Accepted = accepted

	return *dialogues*/
}

func setQuestConditions(conditions map[string]any) QuestAvailabilityConditions {
	output := QuestAvailabilityConditions{
		AvailableForStart:  nil,
		AvailableForFinish: nil,
		Fail:               nil,
	}

	done := make(chan bool)

	processCategory := func(category string) {
		conditionList, ok := conditions[category].([]any)
		if !ok || len(conditionList) == 0 {
			return
		}

		var input *QuestConditionTypes
		switch category {
		case Fail:
			output.Fail = new(QuestConditionTypes)
			input = output.Fail
		case ForStart:
			output.AvailableForStart = new(QuestConditionTypes)
			input = output.AvailableForStart
		case ForFinish:
			output.AvailableForFinish = new(QuestConditionTypes)
			input = output.AvailableForFinish
		}

		for _, c := range conditionList {
			condition, ok := c.(map[string]any)
			if !ok || len(condition) == 0 {
				continue
			}

			props := condition[props].(map[string]any)
			name := condition[parent].(string)

			switch name {
			case "Level":
				input.Level = &LevelCondition{
					CompareMethod: props["compareMethod"].(string),
				}

				float, ok := props["value"].(float64)
				if ok {
					input.Level.Level = float
					continue
				}

				i, ok := props["value"].(int)
				if ok {
					input.Level.Level = float64(i)
					continue
				}
			case "Quest": //TODO sometimes availableAfter isn't available...
				if input.Quest == nil {
					input.Quest = make(map[string]*QuestCondition)
				}

				input.Quest[props["id"].(string)] = &QuestCondition{
					Status:          questStatus[int8(props["status"].([]any)[0].(float64))],
					PreviousQuestID: props["target"].(string),
				}

				avail, _ := props["availableAfter"].(float64)
				if ok {
					input.Quest[props["id"].(string)].AvailableAfter = int(avail)
				}

				continue
			case "TraderLoyalty":
				if input.TraderLoyalty == nil {
					input.TraderLoyalty = make(map[string]*LevelCondition)
				}

				levelCondition := &LevelCondition{
					CompareMethod: props["compareMethod"].(string),
				}

				float, ok := props["value"].(float64)
				if ok {
					levelCondition.Level = float
					input.TraderLoyalty[props["target"].(string)] = levelCondition
					continue
				}

				i, ok := props["value"].(int)
				if ok {
					levelCondition.Level = float64(i)
					input.TraderLoyalty[props["target"].(string)] = levelCondition
					continue
				}
			case "TraderStanding":
				if input.TraderStanding == nil {
					input.TraderStanding = make(map[string]*LevelCondition)
				}

				levelCondition := &LevelCondition{
					CompareMethod: props["compareMethod"].(string),
				}

				float, ok := props["value"].(float64)
				if ok {
					levelCondition.Level = float
					input.TraderStanding[props["target"].(string)] = levelCondition
					continue
				}

				i, ok := props["value"].(int)
				if ok {
					levelCondition.Level = float64(i)
					input.TraderStanding[props["target"].(string)] = levelCondition
					continue
				}
			case "HandoverItem":
				if input.HandoverItem == nil {
					input.HandoverItem = make(map[string]*HandoverCondition)
				}
				handover := &HandoverCondition{
					ItemToHandover: props["target"].([]any)[0].(string),
				}

				isFloat, ok := props["value"].(float64)
				if ok {
					handover.Amount = isFloat
					input.HandoverItem[props["id"].(string)] = handover
					continue
				}

				isInt, ok := props["value"].(int)
				if ok {
					handover.Amount = float64(isInt)
					input.HandoverItem[props["id"].(string)] = handover
					continue
				}
			case "WeaponAssembly":
				if input.WeaponAssembly == nil {
					input.WeaponAssembly = make(map[string]*HandoverCondition)
				}
				handover := &HandoverCondition{
					ItemToHandover: props["target"].([]any)[0].(string),
				}

				isFloat, ok := props["value"].(float64)
				if ok {
					handover.Amount = isFloat
					input.WeaponAssembly[props["id"].(string)] = handover
					continue
				}

				isInt, ok := props["value"].(int)
				if ok {
					handover.Amount = float64(isInt)
					input.WeaponAssembly[props["id"].(string)] = handover
					continue
				}
			case "FindItem":
				if input.FindItem == nil {
					input.FindItem = make(map[string]*HandoverCondition)
				}
				handover := &HandoverCondition{
					ItemToHandover: props["target"].([]any)[0].(string),
				}

				isFloat, ok := props["value"].(float64)
				if ok {
					handover.Amount = isFloat
					input.FindItem[props["id"].(string)] = handover
					continue
				}

				isInt, ok := props["value"].(int)
				if ok {
					handover.Amount = float64(isInt)
					input.FindItem[props["id"].(string)] = handover
					continue
				}
			case "Skill":
				if input.Skills == nil {
					input.Skills = make(map[string]*LevelCondition)
				}
				levelCondition := &LevelCondition{
					CompareMethod: props["compareMethod"].(string),
				}

				float, ok := props["value"].(float64)
				if ok {
					levelCondition.Level = float
					input.Skills[props["target"].(string)] = levelCondition
					continue
				}

				i, ok := props["value"].(int)
				if ok {
					levelCondition.Level = float64(i)
					input.Skills[props["target"].(string)] = levelCondition
					continue
				}
			default:
				continue
			}
		}

		if input == new(QuestConditionTypes) {
			input = nil
		}
	}

	go func() {
		processCategory(Fail)
		done <- true
	}()
	go func() {
		processCategory(ForStart)
		done <- true
	}()
	go func() {
		processCategory(ForFinish)
		done <- true
	}()

	for i := int8(0); i < 3; i++ {
		<-done
	}

	return output
}

var questStatus = map[int8]string{
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

func setQuestRewards(rewards map[string]any) QuestRewardAvailabilityConditions {
	done := make(chan bool)
	output := QuestRewardAvailabilityConditions{
		Start:   nil,
		Success: nil,
		Fail:    nil,
	}

	processCategory := func(label string) {
		category, ok := rewards[label].([]any)
		if !ok || len(category) == 0 {
			return
		}

		var input *QuestRewards
		switch label {
		case Fail:
			output.Fail = new(QuestRewards)
			input = output.Fail
		case Started:
			output.Start = new(QuestRewards)
			input = output.Start
		case Success:
			output.Success = new(QuestRewards)
			input = output.Success
		default:
			log.Println("huh")
		}

		for _, c := range category {
			reward, ok := c.(map[string]any)
			if !ok || len(reward) == 0 {
				continue
			}

			name := reward[_type].(string)

			switch name {
			case "Experience":
				float, ok := reward["value"].(float64)
				if ok {
					input.Experience = int(float)
					continue
				}
				input.Experience = reward["value"].(int)
				continue
			case "Item":
				if input.Items == nil {
					input.Items = make(map[string]QuestRewardItem)
				}
				questRewardItem := QuestRewardItem{}

				items, _ := reward["items"].([]any)
				questRewardItem.Items = make([]map[string]any, 0, len(items))
				for _, idem := range items {
					idem := idem.(map[string]any)
					questRewardItem.Items = append(questRewardItem.Items, idem)
				}

				float, ok := reward["value"].(float64)
				if ok {
					questRewardItem.Value = int(float)
					input.Items[reward["target"].(string)] = questRewardItem
					continue
				}

				questRewardItem.Value = reward["value"].(int)
				input.Items[reward["target"].(string)] = questRewardItem
				continue
			case "AssortmentUnlock":
				input.AssortmentUnlock = reward["target"].(string)
				continue
			case "TraderStanding":
				if input.TraderStanding == nil {
					input.TraderStanding = make(map[string]float64)
				}
				float, ok := reward["value"].(float64)
				if ok {
					input.TraderStanding[reward["target"].(string)] = float
					continue
				}

				i, ok := reward["value"].(int)
				if ok {
					input.TraderStanding[reward["target"].(string)] = float64(i)
					continue
				}

				fmt.Println("TRADERSTANDING NOT FOUND")
				continue
			case "TraderStandingRestore":
				if input.TraderStandingRestore == nil {
					input.TraderStandingRestore = make(map[string]float64)
				}
				float, ok := reward["value"].(float64)
				if ok {
					input.TraderStandingRestore[reward["target"].(string)] = float
					continue
				}

				i, ok := reward["value"].(int)
				if ok {
					input.TraderStandingRestore[reward["target"].(string)] = float64(i)
					continue
				}
				fmt.Println("TRADERSTANDINGRESTORE NOT FOUND")
				continue
			case "TraderUnlock":
				input.TraderUnlock = reward["target"].(string)
				continue
			case "Skill":
				if input.Skills == nil {
					input.Skills = make(map[string]int)
				}
				float, ok := reward["value"].(float64)
				if ok {
					input.Skills[reward["target"].(string)] = int(float)
					continue
				}
				i, ok := reward["value"].(int)
				if ok {
					input.Skills[reward["target"].(string)] = i
					continue
				}
				fmt.Println("SKILL NOT FOUND")
				continue
			case "ProductionScheme":
				if input.ProductionScheme == nil {
					input.ProductionScheme = make(map[string]QuestRewardProductionScheme)
				}
				scheme := QuestRewardProductionScheme{}
				scheme.LoyaltyLevel = int(reward["loyaltyLevel"].(float64))
				scheme.Item = reward["items"].([]any)[0].(map[string]any)["_tpl"].(string)

				ifFloat, ok := reward["traderId"].(float64)
				if ok {
					scheme.AreaID = int(ifFloat)
				} else {
					scheme.AreaID = reward["traderId"].(int)
				}
				/*
					item, ok := reward["items"].([]any)[0].(map[string]any)["_tpl"].(string)
					if ok {
						scheme.Item = item
					}
				*/
				input.ProductionScheme[reward["target"].(string)] = scheme
				continue
			default:
				log.Println("huh")
				continue
			}
		}

	}

	go func() {
		processCategory(Fail)
		done <- true
	}()
	go func() {
		processCategory(Started)
		done <- true
	}()
	go func() {
		processCategory(Success)
		done <- true
	}()

	for i := int8(0); i < 3; i++ {
		<-done
	}
	return output
}

// #endregion

// #region Quest functions
var bearOnlyQuests = map[string]bool{
	"6179b5eabca27a099552e052": true,
	"5e383a6386f77465910ce1f3": true,
	"5e4d515e86f77438b2195244": true,
	"639282134ed9512be67647ed": true,
}

var usecOnlyQuests = map[string]bool{
	"6179b5eabca27a099552e052": true,
	"5e383a6386f77465910ce1f3": true,
	"5e4d515e86f77438b2195244": true,
	"639282134ed9512be67647ed": true,
}

//TODO: Remove hard-coded quests

func CheckIfQuestForOtherFaction(side string, qid string) bool {
	if side == "Bear" {
		return usecOnlyQuests[qid]
	}
	return bearOnlyQuests[qid]
}

// #endregion

// #region Quest structs

type Quest struct {
	Name       string
	Trader     string
	Dialogue   QuestDialogues                    `json:",omitempty"`
	Conditions QuestAvailabilityConditions       `json:",omitempty"`
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
	Amount         float64
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
	Experience            int                                    `json:"Experience,omitempty"`
	Items                 map[string]QuestRewardItem             `json:"Item,omitempty"`
	AssortmentUnlock      string                                 `json:"AssortmentUnlock,omitempty"`
	TraderStanding        map[string]float64                     `json:"TraderStanding,omitempty"`
	TraderStandingRestore map[string]float64                     `json:"TraderStandingRestore,omitempty"`
	TraderUnlock          string                                 `json:"TraderUnlock,omitempty"`
	Skills                map[string]int                         `json:"Skills,omitempty"`
	ProductionScheme      map[string]QuestRewardProductionScheme `json:"ProductionScheme,omitempty"`
}

type QuestRewardProductionScheme struct {
	Item         string
	LoyaltyLevel int
	AreaID       int
}

type QuestRewardItem struct {
	FindInRaid bool
	Items      []map[string]any
	Value      int
}

type QuestRewardAssortUnlock struct {
	Items        []map[string]any
	LoyaltyLevel int
}

// #endregion
