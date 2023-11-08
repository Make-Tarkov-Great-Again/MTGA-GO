package database

import (
	"fmt"
	"sync"
)

const cacheNotExist string = "Cache for %s does not exist"

func GetCacheByID(uid string) (*Cache, error) {
	if profile, ok := profiles[uid]; ok {
		return profile.Cache, nil
	}

	return nil, fmt.Errorf(cacheNotExist, uid)
}

const questCacheNotExist string = "Quest Cache for %s does not exist"

func GetQuestCacheByID(uid string) (*QuestCache, error) {
	cache, err := GetCacheByID(uid)
	if err != nil {
		return nil, err
	}

	if cache.Quests != nil {
		return cache.Quests, nil
	}

	return nil, fmt.Errorf(questCacheNotExist, uid)
}

const traderCacheNotExist string = "Trader Cache for %s does not exist"

func GetTraderCacheByID(uid string) (*TraderCache, error) {
	cache, err := GetCacheByID(uid)
	if err != nil {
		return nil, err
	}

	if cache.Traders != nil {
		return cache.Traders, nil
	}
	return nil, fmt.Errorf(traderCacheNotExist, uid)
}

const inventoryCacheNotExist string = "Inventory Cache for %s does not exist"

func GetInventoryCacheByID(uid string) (*InventoryContainer, error) {
	cache, err := GetCacheByID(uid)
	if err != nil {
		return nil, err
	}

	if cache.Inventory != nil {
		return cache.Inventory, nil
	}
	return nil, fmt.Errorf(inventoryCacheNotExist, uid)
}

func (profile *Profile) SetCache() *Cache {
	var cache *Cache
	if profile.Cache == nil {
		cache = &Cache{
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
				Index:   make(map[string]*AssortIndex),
				Assorts: make(map[string]*Assort),
			},
		}
	} else {
		cache = profile.Cache
	}

	if profile.Character.Info.Nickname != "" {
		var wg sync.WaitGroup

		// Define a function to update the quests map
		updateQuests := func() {
			defer wg.Done()
			for index, quest := range profile.Character.Quests {
				cache.Quests.Index[quest.QID] = int8(index)
			}
		}

		// Define a function to update the common skills map
		updateCommonSkills := func() {
			defer wg.Done()
			for index, commonSkill := range profile.Character.Skills.Common {
				cache.Skills.Common[commonSkill.ID] = int8(index)
			}
		}

		// Define a function to update the hideout areas map
		updateHideoutAreas := func() {
			defer wg.Done()
			for index, area := range profile.Character.Hideout.Areas {
				cache.Hideout.Areas[int8(area.Type)] = int8(index)
			}
		}

		// Start Goroutines for parallel execution
		wg.Add(3)
		go updateQuests()
		go updateCommonSkills()
		go updateHideoutAreas()

		cache.Inventory = SetInventoryContainer(&profile.Character.Inventory)
	}
	return cache
}

type Cache struct {
	Inventory *InventoryContainer
	Skills    *SkillsCache
	Hideout   *HideoutCache
	Quests    *QuestCache
	Traders   *TraderCache
}

type SkillsCache struct {
	Common map[string]int8
}

type HideoutCache struct {
	Areas map[int8]int8
}

type TraderCache struct {
	Index   map[string]*AssortIndex
	Assorts map[string]*Assort
}

type QuestCache struct {
	Index map[string]int8
}
