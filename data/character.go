package data

import (
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"path/filepath"
	"strings"
)

func setCharacter(path string) *Character[map[string]PlayerTradersInfo] {
	output := new(Character[map[string]PlayerTradersInfo])

	data := tools.GetJSONRawMessage(path)
	if err := json.UnmarshalNoEscape(data, output); err != nil {
		log.Println(err)
	}

	return output
}

func GetCharacterByID(uid string) *Character[map[string]PlayerTradersInfo] {
	profile, ok := db.profile[uid]
	if !ok {
		log.Println(characterNotExist, uid)
		return nil
	}
	return profile.Character
}

func GetQuestsAvailableToPlayer(c Character[map[string]PlayerTradersInfo]) ([]any, error) {
	cachedQuests, err := GetQuestCacheByID(c.ID)
	if err != nil && len(c.Quests) != 0 {
		return nil, err
	}

	characterHasQuests := cachedQuests != nil && len(cachedQuests.Index) != 0
	var output []any
	query := GetQuestsQuery()
	for key, value := range query {
		if CheckIfQuestForOtherFaction(c.Info.Side, key) || strings.HasSuffix(value.Name, "-Event") {
			continue
		}

		if value.Conditions.AvailableForStart == nil {
			output = append(output, GetQuestByID(key))
			continue
		}

		forStart := value.Conditions.AvailableForStart

		if forStart.Level != nil {
			if !tools.LevelComparisonCheck(
				forStart.Level.Level,
				c.Info.Level,
				forStart.Level.CompareMethod) {
				continue
			}
		}

		if forStart.Quest == nil && forStart.TraderLoyalty == nil && forStart.TraderStanding == nil {
			output = append(output, GetQuestByID(key))
			continue
		}

		loyaltyCheck := false
		if forStart.TraderLoyalty != nil {
			for trader, loyalty := range forStart.TraderLoyalty {
				traderInfo, ok := c.TradersInfo[trader]
				if !ok {
					log.Fatal("%s doesn't exist in TradersInfo", trader)
				}
				data, err := GetTraderByUID(trader)
				if err != nil {
					return nil, err
				}
				data.SetTraderLoyaltyLevel(&c)

				loyaltyCheck = tools.LevelComparisonCheck(
					loyalty.Level,
					traderInfo.LoyaltyLevel,
					loyalty.CompareMethod,
				)
			}

			if !loyaltyCheck {
				continue
			}
		}

		standingCheck := false
		if forStart.TraderStanding != nil {
			for trader, loyalty := range forStart.TraderStanding {
				traderInfo, ok := c.TradersInfo[trader]
				if !ok {
					log.Fatal("%s doesn't exist in TradersInfo", trader)
				}
				data, err := GetTraderByUID(trader)
				if err != nil {
					return nil, err
				}
				data.SetTraderLoyaltyLevel(&c)

				standingCheck = tools.LevelComparisonCheck(
					loyalty.Level,
					traderInfo.LoyaltyLevel,
					loyalty.CompareMethod,
				)
			}

			if !standingCheck {
				continue
			}
		}

		if forStart.Quest != nil && characterHasQuests {
			if c.IsPreviousQuestComplete(forStart.Quest, cachedQuests) {
				output = append(output, GetQuestByID(key))
				continue
			}
		}
	}

	return output, nil
}

func (c Character[T]) IsPreviousQuestComplete(quests map[string]*QuestCondition, cachedQuests *QuestCache) bool {
	previousQuestCompleted := false
	for _, v := range quests {
		index, ok := cachedQuests.Index[v.PreviousQuestID]
		if !ok {
			continue
		}

		previousQuestCompleted = v.Status == c.Quests[index].Status
	}
	return previousQuestCompleted
}

func (inv *Inventory) CleanInventoryOfDeletedItemMods() bool {
	allItems := GetItems()

	newItems := make([]InventoryItem, 0, len(inv.Items))

	cleaned := 0
	for _, item := range inv.Items {
		if _, ok := allItems[item.TPL]; !ok {
			cleaned++
			continue
		}
		newItems = append(newItems, item)
	}

	if cleaned != 0 {
		log.Println("Removed", cleaned, "modded item(s) from your inventory")
		inv.Items = newItems
		return true
	}
	return false

}

func (c Character[T]) SaveCharacter() error {
	characterFilePath := filepath.Join(profilesPath, c.ID, "character.json")

	if err := tools.WriteToFile(characterFilePath, c); err != nil {
		return fmt.Errorf(characterNotSaved, c.ID, err)
	}
	log.Println("Character saved")
	return nil
}

const (
	characterNotSaved string = "Account for %s was not saved: %s"
	characterNotExist string = "Profile with UID %s does not exist"
)

type Character[T TradersInfo] struct {
	ID                string              `json:"_id"`
	AID               int                 `json:"aid"`
	Savage            *string             `json:"savage"`
	Info              PlayerInfo          `json:"Info"`
	Customization     PlayerCustomization `json:"Customization"`
	Health            HealthInfo          `json:"Health"`
	Inventory         Inventory           `json:"Inventory"`
	Skills            PlayerSkills        `json:"Skills"`
	Stats             PlayerStats         `json:"Stats"`
	Encyclopedia      map[string]bool     `json:"Encyclopedia"`
	ConditionCounters ConditionCounters   `json:"ConditionCounters"`
	BackendCounters   map[string]any      `json:"BackendCounters"`
	InsuredItems      []InsuredItem       `json:"InsuredItems"`
	Hideout           *PlayerHideout      `json:"Hideout"`
	Bonuses           []Bonus             `json:"Bonuses"`
	Notes             Notes               `json:"Notes"`
	Quests            []CharacterQuest    `json:"Quests"`
	RagfairInfo       PlayerRagfairInfo   `json:"RagfairInfo"`
	WishList          []string            `json:"WishList"`
	TradersInfo       T                   `json:"TradersInfo"`
	UnlockedInfo      Unlocked            `json:"UnlockedInfo"`
}

type TradersInfo interface {
	[]any | map[string]PlayerTradersInfo
}

type Unlocked struct {
	UnlockedProductionRecipe []any `json:"unlockedProductionRecipe"`
}

type Notes struct {
	Notes                [][]any            `json:"Notes"`
	TransactionInProcess TransactionProcess `json:"TransactionInProcess,omitempty"`
}

type TransactionProcess struct {
	HasCheckChanges bool `json:"HasCheckChanges"`
	HasHandlers     bool `json:"HasHandlers"`
}

type PlayerTradersInfo struct {
	Unlocked     bool    `json:"unlocked"`
	Disabled     bool    `json:"disabled"`
	SalesSum     float32 `json:"salesSum"`
	Standing     float32 `json:"standing"`
	LoyaltyLevel int8    `json:"loyaltyLevel"`
	NextResupply int     `json:"nextResupply"`
}

type PlayerRagfairInfo struct {
	Rating          float32 `json:"rating"`
	IsRatingGrowing bool    `json:"isRatingGrowing"`
	Offers          []any   `json:"offers"`
}

type PlayerHideoutArea struct {
	AreaType              string `json:"AreaType,omitempty"`
	Type                  int    `json:"type"`
	Level                 int    `json:"level"`
	Active                bool   `json:"active"`
	PassiveBonusesEnabled bool   `json:"passiveBonusesEnabled"`
	CompleteTime          int    `json:"completeTime"`
	Constructing          bool   `json:"constructing"`
	Slots                 []any  `json:"slots"`
	LastRecipe            string `json:"lastRecipe"`
}
type PlayerHideout struct {
	Production  map[string]any      `json:"Production"`
	Areas       []PlayerHideoutArea `json:"Areas"`
	Improvement map[string]any      `json:"Improvement"`
	Seed        int                 `json:"Seed,omitempty"`
}

type ConditionCounters struct {
	Counters []any `json:"Counters"`
}

type PlayerStats struct {
	Eft EftStats `json:"Eft"`
}

type StatCounters struct {
	Items []any `json:"Items"`
}

type Counter struct {
	Items []any
}

type EftStats struct {
	SessionCounters        *Counter       `json:"SessionCounters"`
	OverallCounters        Counter        `json:"OverallCounters"`
	SessionExperienceMult  int            `json:"SessionExperienceMult"`
	ExperienceBonusMult    int            `json:"ExperienceBonusMult"`
	TotalSessionExperience int            `json:"TotalSessionExperience"`
	LastSessionDate        int            `json:"LastSessionDate"`
	Aggressor              map[string]any `json:"Aggressor"`
	DroppedItems           []any          `json:"DroppedItems"`
	FoundInRaidItems       []any          `json:"FoundInRaidItems"`
	Victims                []any          `json:"Victims"`
	CarriedQuestItems      []any          `json:"CarriedQuestItems"`
	DamageHistory          map[string]any `json:"DamageHistory"`
	LastPlayerState        *float32       `json:"LastPlayerState"`
	TotalInGameTime        int            `json:"TotalInGameTime"`
	SurvivorClass          string         `json:"SurvivorClass"`
}

type SkillsCommon struct {
	ID                        string `json:"Id"`
	Progress                  int    `json:"Progress"`
	PointsEarnedDuringSession int    `json:"PointsEarnedDuringSession"`
	LastAccess                int64  `json:"LastAccess"`
}
type SkillsMastering struct {
	ID       string `json:"Id"`
	Progress int    `json:"Progress"`
}

type PlayerSkills struct {
	Common    []SkillsCommon    `json:"Common"`
	Mastering []SkillsMastering `json:"Mastering"`
	Points    int               `json:"Points"`
}

type PlayerInfo struct {
	Nickname               string         `json:"Nickname"`
	LowerNickname          string         `json:"LowerNickname"`
	Side                   string         `json:"Side"`
	Voice                  string         `json:"Voice"`
	Level                  int8           `json:"Level"`
	Experience             int32          `json:"Experience"`
	RegistrationDate       int32          `json:"RegistrationDate"`
	GameVersion            string         `json:"GameVersion"`
	AccountType            int8           `json:"AccountType"`
	MemberCategory         int8           `json:"MemberCategory"`
	LockedMoveCommands     bool           `json:"lockedMoveCommands"`
	SavageLockTime         int32          `json:"SavageLockTime"`
	LastTimePlayedAsSavage int32          `json:"LastTimePlayedAsSavage"`
	Settings               map[string]any `json:"Settings"`
	NicknameChangeDate     int32          `json:"NicknameChangeDate"`
	NeedWipeOptions        []any          `json:"NeedWipeOptions"`
	LastCompletedWipe      struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedWipe"`
	LastCompletedEvent struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedEvent"`
	BannedState             bool  `json:"BannedState"`
	BannedUntil             int32 `json:"BannedUntil"`
	IsStreamerModeAvailable bool  `json:"IsStreamerModeAvailable"`
	SquadInviteRestriction  bool  `json:"SquadInviteRestriction"`
	Bans                    []any `json:"Bans"`
}

type InfoSettings struct {
	Role            string  `json:"Role"`
	BotDifficulty   string  `json:"BotDifficulty"`
	Experience      int32   `json:"Experience"`
	StandingForKill float32 `json:"StandingForKill"`
	AggressorBonus  float32 `json:"AggressorBonus"`
}

type PlayerCustomization struct {
	Head  string `json:"Head"`
	Body  string `json:"Body"`
	Feet  string `json:"Feet"`
	Hands string `json:"Hands"`
}

type InsuredItem struct {
	Tid    string `json:"tid"`
	ItemID string `json:"itemId"`
}

type Bonus struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	TemplateID      string `json:"templateId"`
	IsPositive      bool   `json:"IsPositive,omitempty"`
	LocalizationKey string `json:"LocalizationKey,omitempty"`
	Passive         bool   `json:"passive,omitempty"`
	Production      bool   `json:"production,omitempty"`
	Value           int32  `json:"value,omitempty"`
	Visible         bool   `json:"visible,omitempty"`
}

type HealthInfo struct {
	Hydration   CurrMaxHealth   `json:"Hydration"`
	Energy      CurrMaxHealth   `json:"Energy"`
	Temperature CurrMaxHealth   `json:"Temperature"`
	Poison      CurrMaxHealth   `json:"Poison,omitempty"`
	BodyParts   BodyPartsHealth `json:"BodyParts"`
	UpdateTime  int32           `json:"UpdateTime"`
}

type HealthOf struct {
	Effects map[string]any `json:"Effects,omitempty"`
	Health  CurrMaxHealth  `json:"Health"`
}

type CurrMaxHealth struct {
	Current                      float32 `json:"Current"`
	Maximum                      float32 `json:"Maximum"`
	Minimum                      float32 `json:"Minimum,omitempty"`
	OverDamageReceivedMultiplier float32 `json:"OverDamageReceivedMultiplier,omitempty"`
}

type BodyPartsHealth struct {
	Head     HealthOf `json:"Head"`
	Chest    HealthOf `json:"Chest"`
	Stomach  HealthOf `json:"Stomach"`
	LeftArm  HealthOf `json:"LeftArm"`
	RightArm HealthOf `json:"RightArm"`
	LeftLeg  HealthOf `json:"LeftLeg"`
	RightLeg HealthOf `json:"RightLeg"`
}

type CharacterQuest struct {
	QID                 string         `json:"qid"`
	StartTime           int            `json:"startTime"`
	Status              string         `json:"status"`
	StatusTimers        map[string]int `json:"statusTimers"`
	CompletedConditions []string       `json:"completedConditions,omitempty"`
	AvailableAfter      int            `json:"availableAfter,omitempty"`
}
