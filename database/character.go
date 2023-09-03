package database

import (
	"MT-GO/services"
	"MT-GO/tools"
	"fmt"
	"path/filepath"
	"strings"
)

func GetCharacterByUID(uid string) *Character {
	if profile, ok := profiles[uid]; ok {
		return profile.Character
	}

	fmt.Println("Profile with UID ", uid, " does not have a character")
	return nil
}

func (c *Character) GetQuestsAvailableToPlayer() []interface{} {
	output := []interface{}{}

	query := GetQuestsQuery()

	cachedQuests := GetCache(c.ID).Quests
	characterHasQuests := len(cachedQuests.Index) != 0

	traderStandings := make(map[string]*float64) //temporary

	for key, value := range query {

		if services.CheckIfQuestForOtherFaction(c.Info.Side, key) {
			continue
		}

		if strings.Contains(value.Name, "-Event") {
			fmt.Println("Filter event quests ", value.Name, " properly")
			continue
		}

		if value.Conditions == nil || value.Conditions.AvailableForStart == nil {
			output = append(output, GetQuestByQID(key))
			continue
		}

		forStart := value.Conditions.AvailableForStart

		if forStart.Level != nil {
			if !services.LevelComparisonCheck(
				forStart.Level.Level,
				float64(c.Info.Level),
				forStart.Level.CompareMethod) {

				continue
			}
		}

		if forStart.Quest == nil && forStart.TraderLoyalty == nil && forStart.TraderStanding == nil {
			output = append(output, GetQuestByQID(key))
			continue
		}

		loyaltyCheck := false
		if forStart.TraderLoyalty != nil {
			for trader, loyalty := range forStart.TraderLoyalty {

				if traderStandings[trader] == nil {
					loyaltyLevel := float64(GetTraderByID(trader).GetTraderLoyaltyLevel(c))
					traderStandings[trader] = &loyaltyLevel
				}

				loyaltyCheck = services.LevelComparisonCheck(
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
					loyaltyLevel := float64(GetTraderByID(trader).GetTraderLoyaltyLevel(c))
					traderStandings[trader] = &loyaltyLevel
				}

				standingCheck = services.LevelComparisonCheck(
					loyalty.Level,
					*traderStandings[trader],
					loyalty.CompareMethod)
			}

			if !standingCheck {
				continue
			}
		}

		if forStart.Quest != nil && characterHasQuests {
			if CompletedPreviousQuestCheck(forStart.Quest, &cachedQuests) {
				output = append(output, GetQuestByQID(key))
				continue
			}
		}
	}

	return output
}

func (c *Character) Save(sessionID string) {
	characterFilePath := filepath.Join(profilesPath, sessionID, "character.json")

	err := tools.WriteToFile(characterFilePath, c)
	if err != nil {
		panic(err)
	}
	fmt.Println("Character saved")
}

// QuestAccept updates an existing Accepted quest, or creates and appends new Accepted Quest to cache and Character
func (c *Character) QuestAccept(qid string) {

	cachedQuests := GetCache(c.ID).Quests
	length := len(cachedQuests.Index)
	time := int(tools.GetCurrentTimeInSeconds())

	if length != 0 {
		quest, ok := cachedQuests.Index[qid]
		if ok { //if exists, update cache and copy to quest on character
			cachedQuest := cachedQuests.Quests[qid]

			cachedQuest.Status = "Started"
			cachedQuest.StartTime = time
			cachedQuest.StatusTimers[cachedQuest.Status] = time

			c.Quests[quest] = cachedQuest
		}
	}

	quest := &CharacterQuest{
		QID:          qid,
		StartTime:    time,
		Status:       "Started",
		StatusTimers: map[string]int{},
	}

	query := GetQuestFromQueryByQID(qid)
	if query.Conditions.AvailableForStart.Quest != nil {
		startCondition := query.Conditions.AvailableForStart.Quest
		for _, questCondition := range startCondition {
			if questCondition.AvailableAfter > 0 {

				quest.StartTime = 0
				quest.Status = "AvailableAfter"
				quest.AvailableAfter = time + questCondition.AvailableAfter
			}
		}
	}

	cachedQuests.Index[qid] = int8(length)
	cachedQuests.Quests[qid] = *quest
	c.Quests = append(c.Quests, *quest)

	//c.Save(c.ID)
	services.SendNPCMessage(c.ID, "QuestStart", query.Trader, query.Dialogue.Description, []interface{}{})

	//TODO: create dialogue and notification for Quest
	fmt.Println()
}
