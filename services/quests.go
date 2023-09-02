package services

import (
	"MT-GO/database"
	"fmt"
	"strings"
)

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

func checkIfQuestForOtherFaction(side string, qid string) bool {
	if side == "Bear" {
		return usecOnlyQuests[qid]
	} else {
		return bearOnlyQuests[qid]
	}
}

func GetQuestsAvailableToPlayer(character *database.Character) []interface{} {
	output := []interface{}{}

	query := database.GetQuestsQuery()

	cachedQuests := database.GetCache(character.ID).Quests
	characterHasQuests := len(cachedQuests.Index) != 0

	traderStandings := make(map[string]*float64) //temporary

	for key, value := range query {

		if checkIfQuestForOtherFaction(character.Info.Side, key) {
			continue
		}

		if strings.Contains(value.Name, "-Event") {
			fmt.Println("Filter event quests ", value.Name, " properly")
			continue
		}

		if value.Conditions == nil || value.Conditions.AvailableForStart == nil {
			output = append(output, database.GetQuestByQID(key))
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
			output = append(output, database.GetQuestByQID(key))
			continue
		}

		loyaltyCheck := false
		if forStart.TraderLoyalty != nil {
			for trader, loyalty := range forStart.TraderLoyalty {

				if traderStandings[trader] == nil {
					loyaltyLevel := float64(database.GetTraderByID(trader).GetTraderLoyaltyLevel(character))
					traderStandings[trader] = &loyaltyLevel
				}

				loyaltyCheck = levelComparisonCheck(
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
					loyaltyLevel := float64(database.GetTraderByID(trader).GetTraderLoyaltyLevel(character))
					traderStandings[trader] = &loyaltyLevel
				}

				standingCheck = levelComparisonCheck(
					loyalty.Level,
					*traderStandings[trader],
					loyalty.CompareMethod)
			}

			if !standingCheck {
				continue
			}
		}

		if forStart.Quest != nil && characterHasQuests {
			if completedPreviousQuestCheck(forStart.Quest, character) {
				output = append(output, database.GetQuestByQID(key))
				continue
			}
		}
	}

	return output
}

/* type QuestStatus int

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
) */

func completedPreviousQuestCheck(quests map[string]*database.QuestCondition, character *database.Character) bool {
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
