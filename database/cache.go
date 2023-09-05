package database

import "fmt"

func GetCacheByUID(uid string) *Cache {
	if profile, ok := profiles[uid]; ok {
		return profile.Cache
	}

	fmt.Println("Profile with UID ", uid, " does not have cache")
	return nil
}

type Cache struct {
	Quests  QuestCache
	Traders TraderCache
}

type TraderCache struct {
	Index         map[string]*AssortIndex
	Assorts       map[string]*Assort
	LoyaltyLevels map[string]int8
}

type QuestCache struct {
	Index  map[string]int8
	Quests map[string]CharacterQuest
}
