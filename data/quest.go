package data

import (
	"MT-GO/tools"
	"fmt"
	"github.com/alphadose/haxmap"
	"github.com/goccy/go-json"
	"log"
	"strconv"
	"strings"
)

type Quest struct {
	quests        *haxmap.Map[string, map[string]any] //map[string]map[string]any
	query         *haxmap.Map[string, *Query]         //map[string]*Query
	factionQuests *FactionQuests
	status        *haxmap.Map[int8, string]
}

type FactionQuests struct {
	Bear *haxmap.Map[string, struct{}]
	Usec *haxmap.Map[string, struct{}]
}

// #region Quests getters

func GetQuestsQuery() *haxmap.Map[string, *Query] {
	return db.quest.query
}

func GetQuestFromQueryByID(qid string) *Query {
	query, ok := db.quest.query.Get(qid)
	if !ok {
		log.Println("Quest", qid, "does not exist in quests query")
		return nil
	}
	return query
}

func GetQuestByID(qid string) any {
	quest, ok := db.quest.quests.Get(qid)
	if !ok {
		log.Println("Quest", qid, "does not exist in quests")
		return nil
	}
	return quest
}

// #endregion

// region Quest setters
const (
	_type     string = "type"
	ForFinish string = "AvailableForFinish"
	ForStart  string = "AvailableForStart"
	Fail      string = "Fail"
	Started   string = "Started"
	Success   string = "Success"
)

func setQuests() {
	db.quest = &Quest{
		quests: haxmap.New[string, map[string]any](), //make(map[string]map[string]any),
	}
	raw := tools.GetJSONRawMessage(questsPath)
	if err := json.UnmarshalNoEscape(raw, &db.quest.quests); err != nil {
		msg := tools.CheckParsingError(raw, err)
		log.Fatalln(msg)
	}
}

func setQuestLookup() {
	db.quest.query = haxmap.New[string, *Query]()
	db.quest.status = haxmap.New[int8, string]()
	db.quest.factionQuests = &FactionQuests{
		Bear: haxmap.New[string, struct{}](uintptr(4)),
		Usec: haxmap.New[string, struct{}](uintptr(4)),
	}
	questStatus := map[int8]string{
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
	bearOnlyQuests := []string{
		"6179b5b06e9dd54ac275e409",
		"5e381b0286f77420e3417a74",
		"5e4d4ac186f774264f758336",
		"639136d68ba6894d155e77cf",
	}
	usecOnlyQuests := []string{
		"6179b5eabca27a099552e052",
		"5e383a6386f77465910ce1f3",
		"5e4d515e86f77438b2195244",
		"639282134ed9512be67647ed",
	}

	done := make(chan struct{})

	go func() {
		for key, value := range questStatus {
			db.quest.status.Set(key, value)
		}
		for _, key := range bearOnlyQuests {
			db.quest.factionQuests.Bear.Set(key, struct{}{})
		}
		for _, key := range usecOnlyQuests {
			db.quest.factionQuests.Usec.Set(key, struct{}{})
		}
		done <- struct{}{}
	}()

	go func() {
		db.quest.quests.ForEach(func(k string, v map[string]any) bool {
			done2 := make(chan struct{})
			query := &Query{
				Name:   v["QuestName"].(string),
				Trader: v["traderId"].(string),
			}

			go func() {
				query.Dialogue = setQuestDialogue(v)
				done2 <- struct{}{}
			}()

			go func() {
				if questConditions, ok := v["conditions"].(map[string]any); ok {
					query.Conditions = setQuestConditions(questConditions)
				}
				done2 <- struct{}{}
			}()

			go func() {
				if questRewards, _ := v["rewards"].(map[string]any); questRewards != nil {
					query.Rewards = setQuestRewards(questRewards)
				}
				done2 <- struct{}{}
			}()

			for i := int8(0); i < 3; i++ {
				<-done2
			}
			db.quest.query.Set(k, query)
			return true
		})
		done <- struct{}{}
	}()

	for i := int8(0); i < 2; i++ {
		<-done
	}
}

func setQuestDialogue(quest map[string]any) QuestDialogues {
	return QuestDialogues{
		Description: quest["description"].(string),
		Started:     quest["startedMessageText"].(string),
		Success:     quest["successMessageText"].(string),
		Fail:        quest["failMessageText"].(string),
	}

	//TODO: remove if not needed
	/* 	complete, ok := quest["completePlayerMessage"].(string)
	   	if !ok {
	   		log.Println("quest[`completePlayerMessage`]")
	   	}
	   	dialogues.Complete = complete */

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

	done := make(chan struct{})

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

			conditionType := condition["conditionType"].(string)

			switch conditionType {
			case "Level":
				input.Level = &LevelCondition{
					CompareMethod: condition["compareMethod"].(string),
				}

				float, ok := condition["value"].(float64)
				if ok {
					input.Level.Level = int8(float)
					continue
				}

				i, ok := condition["value"].(int)
				if ok {
					input.Level.Level = int8(i)
					continue
				}
			case "Quest": //TODO sometimes availableAfter isn't available...
				if input.Quest == nil {
					input.Quest = make(map[string]*QuestCondition)
				}

				var conditionStatus float64
				if f, ok := condition["status"].([]any)[0].(float64); ok {
					conditionStatus = f
				}

				if s, ok := condition["status"].([]any)[0].(string); ok {
					value, err := strconv.ParseFloat(s, 64)
					if err != nil {
						log.Fatal(err)
					}
					conditionStatus = value
				}

				status, _ := db.quest.status.Get(int8(conditionStatus))

				input.Quest[condition["id"].(string)] = &QuestCondition{
					Status:          status,
					PreviousQuestID: condition["target"].(string),
				}

				avail, _ := condition["availableAfter"].(float64)
				if ok {
					input.Quest[condition["id"].(string)].AvailableAfter = int(avail)
				}

				continue
			case "TraderLoyalty":
				if input.TraderLoyalty == nil {
					input.TraderLoyalty = make(map[string]*LevelCondition)
				}

				levelCondition := &LevelCondition{
					CompareMethod: condition["compareMethod"].(string),
				}

				float, ok := condition["value"].(float64)
				if ok {
					levelCondition.Level = int8(float)
					input.TraderLoyalty[condition["target"].(string)] = levelCondition
					continue
				}

				i, ok := condition["value"].(int)
				if ok {
					levelCondition.Level = int8(i)
					input.TraderLoyalty[condition["target"].(string)] = levelCondition
					continue
				}
			case "TraderStanding":
				if input.TraderStanding == nil {
					input.TraderStanding = make(map[string]*LevelCondition)
				}

				levelCondition := &LevelCondition{
					CompareMethod: condition["compareMethod"].(string),
				}

				float, ok := condition["value"].(float64)
				if ok {
					levelCondition.Level = int8(float)
					input.TraderStanding[condition["target"].(string)] = levelCondition
					continue
				}

				i, ok := condition["value"].(int)
				if ok {
					levelCondition.Level = int8(i)
					input.TraderStanding[condition["target"].(string)] = levelCondition
					continue
				}
			case "HandoverItem":
				if input.HandoverItem == nil {
					input.HandoverItem = make(map[string]*HandoverCondition)
				}
				handover := &HandoverCondition{
					ItemToHandover: condition["target"].([]any)[0].(string),
				}

				isFloat, ok := condition["value"].(float64)
				if ok {
					handover.Amount = isFloat
					input.HandoverItem[condition["id"].(string)] = handover
					continue
				}

				isInt, ok := condition["value"].(int)
				if ok {
					handover.Amount = float64(isInt)
					input.HandoverItem[condition["id"].(string)] = handover
					continue
				}
			case "WeaponAssembly":
				if input.WeaponAssembly == nil {
					input.WeaponAssembly = make(map[string]*HandoverCondition)
				}
				handover := &HandoverCondition{
					ItemToHandover: condition["target"].([]any)[0].(string),
				}

				isFloat, ok := condition["value"].(float64)
				if ok {
					handover.Amount = isFloat
					input.WeaponAssembly[condition["id"].(string)] = handover
					continue
				}

				isInt, ok := condition["value"].(int)
				if ok {
					handover.Amount = float64(isInt)
					input.WeaponAssembly[condition["id"].(string)] = handover
					continue
				}
			case "FindItem":
				if input.FindItem == nil {
					input.FindItem = make(map[string]*HandoverCondition)
				}
				handover := &HandoverCondition{
					ItemToHandover: condition["target"].([]any)[0].(string),
				}

				isFloat, ok := condition["value"].(float64)
				if ok {
					handover.Amount = isFloat
					input.FindItem[condition["id"].(string)] = handover
					continue
				}

				isInt, ok := condition["value"].(int)
				if ok {
					handover.Amount = float64(isInt)
					input.FindItem[condition["id"].(string)] = handover
					continue
				}
			case "Skill":
				if input.Skills == nil {
					input.Skills = make(map[string]*LevelCondition)
				}
				levelCondition := &LevelCondition{
					CompareMethod: condition["compareMethod"].(string),
				}

				float, ok := condition["value"].(float64)
				if ok {
					levelCondition.Level = int8(float)
					input.Skills[condition["target"].(string)] = levelCondition
					continue
				}

				i, ok := condition["value"].(int)
				if ok {
					levelCondition.Level = int8(i)
					input.Skills[condition["target"].(string)] = levelCondition
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
		done <- struct{}{}
	}()
	go func() {
		processCategory(ForStart)
		done <- struct{}{}
	}()
	go func() {
		processCategory(ForFinish)
		done <- struct{}{}
	}()

	for i := int8(0); i < 3; i++ {
		<-done
	}

	return output
}

func setQuestRewards(rewards map[string]any) QuestRewardAvailabilityConditions {
	done := make(chan struct{})
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
				if float, ok := reward["value"].(float64); ok {
					input.Experience = int(float)
					continue
				}

				if isString, ok := reward["value"].(string); ok {
					value, err := strconv.Atoi(isString)
					if err != nil {
						log.Fatal(err)
					}
					input.Experience = value
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

				if float, ok := reward["value"].(float64); ok {
					questRewardItem.Value = int(float)
					input.Items[reward["target"].(string)] = questRewardItem
					continue
				}

				if isString, ok := reward["value"].(string); ok {
					value, err := strconv.Atoi(isString)
					if err != nil {
						log.Fatal(err)
					}
					questRewardItem.Value = value
				}

				if isInt, ok := reward["value"].(int); ok {
					questRewardItem.Value = isInt
				}

				input.Items[reward["target"].(string)] = questRewardItem
				continue
			case "AssortmentUnlock":
				input.AssortmentUnlock = reward["target"].(string)
				continue
			case "TraderStanding":
				if input.TraderStanding == nil {
					input.TraderStanding = make(map[string]float64)
				}

				if float, ok := reward["value"].(float64); ok {
					input.TraderStanding[reward["target"].(string)] = float
					continue
				}

				if i, ok := reward["value"].(int); ok {
					input.TraderStanding[reward["target"].(string)] = float64(i)
					continue
				}

				if s, ok := reward["value"].(string); ok {
					value, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
					if err != nil {
						log.Fatal(err)
					}
					input.TraderStanding[reward["target"].(string)] = value
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

				if float, ok := reward["value"].(float64); ok {
					input.Skills[reward["target"].(string)] = int(float)
					continue
				}
				if i, ok := reward["value"].(int); ok {
					input.Skills[reward["target"].(string)] = i
					continue
				}
				if s, ok := reward["value"].(string); ok {
					value, err := strconv.Atoi(s)
					if err != nil {
						log.Fatal(err)
					}
					input.Skills[reward["target"].(string)] = value
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

				if f, ok := reward["traderId"].(float64); ok {
					scheme.AreaID = int(f)
				}

				if s, ok := reward["traderId"].(string); ok {
					value, err := strconv.Atoi(s)
					if err != nil {
						log.Fatal(err)
					}
					scheme.AreaID = value
				}

				if i, ok := reward["traderId"].(int); ok {
					scheme.AreaID = i
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
		done <- struct{}{}
	}()
	go func() {
		processCategory(Started)
		done <- struct{}{}
	}()
	go func() {
		processCategory(Success)
		done <- struct{}{}
	}()

	for i := int8(0); i < 3; i++ {
		<-done
	}
	return output
}

// #endregion

// #region Quest functions

func CheckIfQuestForOtherFaction(side string, qid string) bool {
	if side == "Bear" {
		_, ok := db.quest.factionQuests.Usec.Get(qid)
		return ok
	}
	_, ok := db.quest.factionQuests.Bear.Get(qid)
	return ok
}

// #endregion

// #region Quest structs

type Query struct {
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
	Level         int8
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
