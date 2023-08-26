package services

import (
	"MT-GO/database"
	"MT-GO/structs"
	"fmt"
)

func GetQuestsAvailableToPlayer(sessionID string) []interface{} {
	output := []interface{}{}

	character := database.GetCharacterByUID(sessionID)
	quests := database.GetQuests()     //raw quests
	query := database.GetQuestsQuery() //quest query

	characterQuests := character.Quests
	characterHasQuests := len(characterQuests) != 0

	for key, value := range query {

		if value.Conditions == nil || value.Conditions.AvailableForStart == nil {
			output = append(output, quests[key])
			continue
		}

		forStart := value.Conditions.AvailableForStart

		if forStart.Level != nil {
			if !comparisonCheck(
				value.Conditions.AvailableForStart.Level.Level,
				float64(character.Info.Level),
				value.Conditions.AvailableForStart.Level.CompareMethod) {

				continue
			}
		}

		if forStart.Quest == nil && forStart.TraderLoyalty == nil {
			output = append(output, quests[key])
			continue
		}

		if !characterHasQuests {
			continue
		}

		loyaltyCheck := false
		if forStart.TraderLoyalty != nil {
			for trader, loyalty := range forStart.TraderLoyalty {
				loyaltyCheck = comparisonCheck(
					loyalty.Level,
					float64(character.TradersInfo[trader].Standing),
					loyalty.CompareMethod)
			}

			if !loyaltyCheck {
				continue
			}
		}

		if forStart.Quest != nil {
			if completedPreviousQuestCheck(forStart.Quest, characterQuests) && loyaltyCheck {
				output = append(output, quests[key])
				continue
			}
		}

	}

	return output
}

func completedPreviousQuestCheck(quests map[string]*structs.QuestCondition, characterQuests []interface{}) bool {
	fmt.Println("Adjust completedPreviousQuestCheck to account for PreviousQuestID")
	return false
}

func comparisonCheck(questLevel float64, characterLevel float64, compareMethod string) bool {
	fmt.Println("Compare method of " + compareMethod)

	switch compareMethod {
	case ">=":
		return characterLevel >= questLevel
	case ">":
		return characterLevel > questLevel
	case "<":
		return characterLevel < questLevel
	case "<=":
		return characterLevel <= questLevel
	case "=":
		return characterLevel == questLevel
	default:
		fmt.Println("Unknown comparison method of " + compareMethod)
		return false
	}
}
