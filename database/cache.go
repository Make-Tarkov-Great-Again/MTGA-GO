package database

import "fmt"

func GetCacheByUID(uid string) *Cache {
	if profile, ok := profiles[uid]; ok {
		return profile.Cache
	}

	fmt.Println("Profile with UID ", uid, " does not have cache")
	return nil
}

func GetQuestCacheByUID(uid string) *QuestCache {
	if profile, ok := profiles[uid]; ok {
		return profile.Cache.Quests
	}
	return nil
}

func GetTraderCacheByUID(uid string) *TraderCache {
	if profile, ok := profiles[uid]; ok {
		return profile.Cache.Traders
	}
	return nil
}

func (p *Profile) setCache() *Cache {
	cache := &Cache{
		Skills: &SkillsCache{
			Common: make(map[string]int8),
		},
		Hideout: &HideoutCache{
			Areas: make(map[int8]int8),
		},
		Quests: &QuestCache{
			Index: make(map[string]int8),
		},
		Traders: &TraderCache{
			Index:         make(map[string]*AssortIndex),
			Assorts:       make(map[string]*Assort),
			LoyaltyLevels: make(map[string]int8),
		},
	}

	for index, quest := range p.Character.Quests {
		cache.Quests.Index[quest.QID] = int8(index)
	}

	for index, commonSkill := range p.Character.Skills.Common {
		cache.Skills.Common[commonSkill.ID] = int8(index)
	}

	for index, area := range p.Character.Hideout.Areas {
		cache.Hideout.Areas[int8(area.Type)] = int8(index)
	}

	return cache
}

type Cache struct {
	Skills  *SkillsCache
	Hideout *HideoutCache
	Quests  *QuestCache
	Traders *TraderCache
}

type SkillsCache struct {
	Common map[string]int8
}

type HideoutCache struct {
	Areas map[int8]int8
}

type TraderCache struct {
	Index         map[string]*AssortIndex
	Assorts       map[string]*Assort
	LoyaltyLevels map[string]int8
}

type QuestCache struct {
	Index map[string]int8
}
