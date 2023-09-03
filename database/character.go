package database

import (
	"MT-GO/tools"
	"fmt"
	"path/filepath"
)

func GetCharacterByUID(uid string) *Character {
	if profile, ok := profiles[uid]; ok {
		return profile.Character
	}

	fmt.Println("Profile with UID ", uid, " does not have a character")
	return nil
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

	startCondition := GetQuestFromQueryByQID(qid).Conditions.AvailableForStart
	if startCondition.Quest != nil {
		for _, questCondition := range startCondition.Quest {
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

	c.Save(c.ID)

	//TODO: create dialogue and notification for Quest
	fmt.Println(startCondition)
}
