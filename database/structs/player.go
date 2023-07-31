package structs

type PlayerTemplate struct {
	ID              string                    `json:"_id,omitempty"`
	AID             int32                     `json:"aid,omitempty"`
	Savage          string                    `json:"savage,omitempty"`
	Info            PlayerInfo                `json:"Info,omitempty"`
	Customization   PlayerCustomization       `json:"Customization,omitempty"`
	Health          PlayerHealth              `json:"Health,omitempty"`
	Skills          PlayerSkills              `json:"Skills,omitempty"`
	Stats           PlayerStats               `json:"Stats,omitempty"`
	Encyclopedia    map[string]bool           `json:"Encyclopedia,omitempty"`
	BackendCounters map[string]BackendCounter `json:"BackendCounters,omitempty"`
	InsuredItems    []InsuredItem             `json:"InsuredItems,omitempty"`
	Bonuses         []Bonus                   `json:"Bonuses,omitempty"`
	Notes           struct {
		Notes []Note `json:"Notes,omitempty"`
	} `json:"Notes,omitempty"`
	Quests        []PlayerQuest `json:"Quests,omitempty"`
	WishList      []string      `json:"WishList,omitempty"`
	SurvivorClass string        `json:"SurvivorClass,omitempty"`
}

type PlayerQuest struct {
	Qid                 string            `json:"qid,omitempty"`
	StartTime           string            `json:"startTime,omitempty"`
	Status              int8              `json:"status,omitempty"`
	StatusTimers        map[string]string `json:"statusTimers,omitempty"`
	CompletedConditions []string          `json:"completedConditions,omitempty"`
	AvailableAfter      int32             `json:"availableAfter,omitempty"`
}

type Note struct {
	Time int32  `json:"Time,omitempty"`
	Text string `json:"Text,omitempty"`
}

type Bonus struct {
	Type       string `json:"type,omitempty"`
	TemplateID string `json:"templateId,omitempty"`
	Value      int32  `json:"value,omitempty"`
	Passive    bool   `json:"passive,omitempty"`
	Production bool   `json:"production,omitempty"`
	Visible    bool   `json:"visible,omitempty"`
}

type PlayerHideout struct {
	Production   map[string]Production         `json:"Production,omitempty"`
	Areas        []Areas                       `json:"Areas,omitempty"`
	Improvements map[string]HideoutImprovement `json:"Improvements,omitempty"`
}

type Areas struct {
	Type                  int8       `json:"type,omitempty"`
	Level                 int8       `json:"level,omitempty"`
	Active                bool       `json:"active,omitempty"`
	PassiveBonusesEnabled bool       `json:"passiveBonusesEnabled,omitempty"`
	CompleteTime          int32      `json:"completeTime,omitempty"`
	Constructing          bool       `json:"constructing,omitempty"`
	Slots                 []AreaSlot `json:"slots,omitempty"`
	LastRecipe            string     `json:"lastRecipe,omitempty"`
}

type AreaSlot struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type HideoutImprovement struct {
	Completed                bool   `json:"completed,omitempty"`
	ImproveCompleteTimestamp string `json:"improveCompleteTimestamp,omitempty"`
}

type Production struct {
	Progress       int32
	InProgress     bool
	RecipeId       string
	Products       []string
	SkipTime       int32
	ProductionTime int32
	StartTimestamp int32
}

type InsuredItem struct {
	Tid    string `json:"tid,omitempty"`
	ItemId string `json:"itemId,omitempty"`
}

type BackendCounter struct {
	Id    string  `json:"id,omitempty"`
	Qid   string  `json:"qid,omitempty"`
	Value float32 `json:"value,omitempty"`
}

type DamageHistory struct {
	LethalDamagePart string                   `json:"LethalDamagePart,omitempty"`
	LethalDamage     interface{}              `json:"LethalDamage,omitempty"`
	BodyParts        map[string][]interface{} `json:"BodyParts,omitempty"`
}

type DeathCause struct {
	DamageType int32  `json:"DamageType,omitempty"`
	Side       int32  `json:"Side,omitempty"`
	Role       int32  `json:"Role,omitempty"`
	WeaponId   string `json:"WeaponId,omitempty"`
}

type OverallCountersItem struct {
	Key   string  `json:"Key,omitempty"`
	Value float32 `json:"Value,omitempty"`
}

type PlayerStats struct {
	SessionCounters struct {
		Items []interface{} `json:"Items,omitempty"`
	} `json:"SessionCounters,omitempty"`
	OverallCounters struct {
		Items []OverallCountersItem `json:"Items,omitempty"`
	} `json:"OverallCounters,omitempty"`
	SessionExperienceMult  int32         `json:"SessionExperienceMult,omitempty"`
	ExperienceBonusMult    int32         `json:"ExperienceBonusMult,omitempty"`
	TotalSessionExperience int32         `json:"TotalSessionExperience,omitempty"`
	LastSessionDate        int32         `json:"LastSessionDate,omitempty"`
	Aggressor              interface{}   `json:"Aggressor,omitempty"`
	DroppedItems           []interface{} `json:"DroppedItems,omitempty"`
	FoundInRaidItems       []interface{} `json:"FoundInRaidItems,omitempty"`
	Victims                []interface{} `json:"Victims,omitempty"`
	CarriedQuestItems      []interface{} `json:"CarriedQuestItems,omitempty"`
	DamageHistory          DamageHistory `json:"DamageHistory,omitempty"`
	DeathCause             DeathCause    `json:"DeathCause,omitempty"`
	LastPlayerState        interface{}   `json:"LastPlayerState,omitempty"`
	TotalInGameTime        int32         `json:"TotalInGameTime,omitempty"`
	SurvivorClass          string        `json:"SurvivorClass,omitempty"`
}

type PlayerSkill struct {
	ID                        string `json:"Id,omitempty"`
	Progress                  int32  `json:"Progress,omitempty"`
	PointsEarnedDuringSession int32  `json:"PointsEarnedDuringSession,omitempty"`
	LastAccess                int32  `json:"LastAccess,omitempty"`
}

type PlayerSkills struct {
	Common    []PlayerSkill `json:"Common,omitempty"`
	Mastering []PlayerSkill `json:"Mastering,omitempty"`
	Points    int32         `json:"Points,omitempty"`
}

type PlayerInventory struct {
	Items           []InventoryItem   `json:"items,omitempty"`
	Equipment       string            `json:"equipment,omitempty"`
	Stash           string            `json:"stash,omitempty"`
	SortingTable    string            `json:"sortingTable,omitempty"`
	QuestRaidItems  string            `json:"questRaidItems,omitempty"`
	QuestStashItems string            `json:"questStashItems,omitempty"`
	FastPanel       map[string]string `json:"fastPanel,omitempty"`
}

type PlayerInfo struct {
	Nickname                string         `json:"Nickname,omitempty"`
	LowerNickname           string         `json:"LowerNickname,omitempty"`
	Side                    string         `json:"Side,omitempty"`
	Voice                   string         `json:"Voice,omitempty"`
	Level                   int8           `json:"Level,omitempty"`
	Experience              int32          `json:"Experience,omitempty"`
	RegistrationDate        int32          `json:"RegistrationDate,omitempty"`
	GameVersion             string         `json:"GameVersion,omitempty"`
	AccountType             int8           `json:"AccountType,omitempty"`
	MemberCategory          MemberCategory `json:"MemberCategory,omitempty"`
	LockedMoveCommands      bool           `json:"lockedMoveCommands,omitempty"`
	SavageLockTime          int32          `json:"SavageLockTime,omitempty"`
	LastTimePlayedAsSavage  int32          `json:"LastTimePlayedAsSavage,omitempty"`
	Settings                PlayerSettings `json:"Settings,omitempty"`
	NicknameChangeDate      int32          `json:"NicknameChangeDate,omitempty"`
	NeedWipeOptions         []string       `json:"NeedWipeOptions,omitempty"`
	LastCompletedWipe       Event          `json:"lastCompletedWipe,omitempty"`
	LastCompletedEvent      Event          `json:"lastCompletedEvent,omitempty"`
	BannedState             bool           `json:"BannedState,omitempty"`
	BannedUntil             int32          `json:"BannedUntil,omitempty"`
	IsStreamerModeAvailable bool           `json:"IsStreamerModeAvailable,omitempty"`
	Bans                    []string       `json:"Bans,omitempty"`
}

type PlayerSettings struct {
	Role            string  `json:"Role,omitempty"`
	BotDifficulty   string  `json:"BotDifficulty,omitempty"`
	Experience      int32   `json:"Experience,omitempty"`
	StandingForKill float32 `json:"StandingForKill,omitempty"`
	AggressorBonus  float32 `json:"AggressorBonus,omitempty"`
}

type PlayerCustomization struct {
	Body  string `json:"Body,omitempty"`
	Feet  string `json:"Feet,omitempty"`
	Hands string `json:"Hands,omitempty"`
	Head  string `json:"Head,omitempty"`
}

type PlayerHealth struct {
	Hydration   HealthInfo                `json:"Hydration,omitempty"`
	Energy      HealthInfo                `json:"Energy,omitempty"`
	Temperature HealthInfo                `json:"Temperature,omitempty"`
	BodyParts   map[string]BodyPartHealth `json:"BodyParts,omitempty"`
	UpdateTime  int32                     `json:"UpdateTime,omitempty"`
}

type BodyPartHealth struct {
	Health HealthInfo `json:"Health,omitempty"`
}

type HealthInfo struct {
	Current float32 `json:"Current,omitempty"`
	Maximum float32 `json:"Maximum,omitempty"`
}

type Event struct {
	Oid string `json:"$oid,omitempty"`
}
