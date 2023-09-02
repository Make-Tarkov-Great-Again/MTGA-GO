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

func SaveCharacter(sessionID string, character Character) {
	characterFilePath := filepath.Join(profilesPath, sessionID, "character.json")

	err := tools.WriteToFile(characterFilePath, character)
	if err != nil {
		panic(err)
	}
	fmt.Println("Character saved")
}

// QuestAccept updates an existing Accepted quest, or creates a new Accepted Quest
func (c Character) QuestAccept(qid string) {
	//TODO: check quest cache on character
	cache := GetCache(c.ID)
	quest, ok := cache.Quests[qid]
	if ok {
		fmt.Println(quest)
	}
	//quest := &database.CharacterQuest{}
	fmt.Println()
}
