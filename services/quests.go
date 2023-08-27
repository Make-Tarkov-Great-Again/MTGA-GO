package services

import (
	"MT-GO/database"
	"MT-GO/structs"
	"fmt"
	"strings"
)

var bearOnlyQuests = map[string]struct{}{
	"6179b5eabca27a099552e052": {},
	"5e383a6386f77465910ce1f3": {},
	"5e4d515e86f77438b2195244": {},
	"639282134ed9512be67647ed": {},
}

var usecOnlyQuests = map[string]struct{}{
	"6179b5eabca27a099552e052": {},
	"5e383a6386f77465910ce1f3": {},
	"5e4d515e86f77438b2195244": {},
	"639282134ed9512be67647ed": {},
}

func checkIfQuestForOtherFaction(side string, qid string) bool {
	if side == "Bear" {
		_, ok := usecOnlyQuests[qid]
		return ok
	} else {
		_, ok := bearOnlyQuests[qid]
		return ok
	}
}

func GetQuestsAvailableToPlayer(sessionID string) []interface{} {
	output := []interface{}{}

	character := database.GetCharacterByUID(sessionID)
	quests := database.GetQuests()     //raw quests
	query := database.GetQuestsQuery() //quest query

	characterHasQuests := len(character.Quests) != 0

	traderStandings := make(map[string]*int) //temporary

	for key, value := range query {

		if checkIfQuestForOtherFaction(character.Info.Side, key) {
			continue
		}

		if strings.Contains(value.Name, "-Event") {
			fmt.Println("Filter event quests ", value.Name, " properly")
			continue
		}

		if value.Conditions == nil || value.Conditions.AvailableForStart == nil {
			output = append(output, quests[key])
			continue
		}

		forStart := value.Conditions.AvailableForStart

		if forStart.Level != nil {
			if !levelComparisonCheck(
				forStart.Level.Level,
				float64(character.Info.Level),
				forStart.Level.CompareMethod) {

				continue
			}
		}

		if forStart.Quest == nil && forStart.TraderLoyalty == nil && forStart.TraderStanding == nil {
			output = append(output, quests[key])
			continue
		}

		loyaltyCheck := false
		if forStart.TraderLoyalty != nil {
			for trader, loyalty := range forStart.TraderLoyalty {

				if traderStandings[trader] == nil {
					loyaltyLevel := GetTraderLoyaltyLevel(trader, character)
					traderStandings[trader] = &loyaltyLevel
				}

				loyaltyCheck = levelComparisonCheck(
					loyalty.Level,
					float64(*traderStandings[trader]),
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
					loyaltyLevel := GetTraderLoyaltyLevel(trader, character)
					traderStandings[trader] = &loyaltyLevel
				}

				standingCheck = levelComparisonCheck(
					loyalty.Level,
					float64(*traderStandings[trader]),
					loyalty.CompareMethod)
			}

			if !standingCheck {
				continue
			}
		}

		if forStart.Quest != nil && characterHasQuests {
			if completedPreviousQuestCheck(forStart.Quest, character) {
				output = append(output, quests[key])
				continue
			}
		}
	}

	return output
}

type QuestStatus int

const (
	Locked             QuestStatus = 0
	AvailableForStart  QuestStatus = 1
	Started            QuestStatus = 2
	AvailableForFinish QuestStatus = 3
	Success            QuestStatus = 4
	Fail               QuestStatus = 5
	FailRestartable    QuestStatus = 6
	MarkedAsFailed     QuestStatus = 7
	Expired            QuestStatus = 8
	AvailableAfter     QuestStatus = 9
)

func completedPreviousQuestCheck(quests map[string]*structs.QuestCondition, character *structs.PlayerTemplate) bool {
	previousQuestCompleted := false

	for _, v := range quests {
		for _, quest := range character.Quests {
			qid, ok := quest["qid"].(string)
			if !ok || qid != v.PreviousQuestID {
				continue
			}

			previousQuestCompleted = v.Status == quest["status"].(int)
		}
	}
	return previousQuestCompleted
}

func levelComparisonCheck(requiredLevel float64, currentLevel float64, compareMethod string) bool {
	switch compareMethod {
	case ">=":
		return currentLevel >= requiredLevel
	case ">":
		return currentLevel > requiredLevel
	case "<":
		return currentLevel < requiredLevel
	case "<=":
		return currentLevel <= requiredLevel
	case "=":
		return currentLevel == requiredLevel
	default:
		fmt.Println("Unknown comparison method of " + compareMethod)
		return false
	}
}
