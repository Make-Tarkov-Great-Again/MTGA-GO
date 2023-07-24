package structs

type PlayerTemplate struct {
	ID              string                    `json:"_id,omitempty"`
	AID             int                       `json:"aid,omitempty"`
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
		Notes map[string]Note `json:"Notes,omitempty"`
	} `json:"Notes,omitempty"`
	Quests        []Quest  `json:"Quests,omitempty"`
	WishList      []string `json:"WishList,omitempty"`
	SurvivorClass string   `json:"SurvivorClass,omitempty"`
}

type Quest struct {
	Qid                 string            `json:"qid,omitempty"`
	StartTime           string            `json:"startTime,omitempty"`
	Status              int               `json:"status,omitempty"`
	StatusTimers        map[string]string `json:"statusTimers,omitempty"`
	CompletedConditions []string          `json:"completedConditions,omitempty"`
	AvailableAfter      int               `json:"availableAfter,omitempty"`
}

type Note struct {
	Time int    `json:"Time,omitempty"`
	Text string `json:"Text,omitempty"`
}

type Bonus struct {
	Type       string `json:"type,omitempty"`
	TemplateID string `json:"templateId,omitempty"`
	Value      int    `json:"value,omitempty"`
	Passive    bool   `json:"passive,omitempty"`
	Production bool   `json:"production,omitempty"`
	Visible    bool   `json:"visible,omitempty"`
}

type PlayerHideout struct {
	Production   map[string]HideoutProduction  `json:"Production,omitempty"`
	Areas        []HideoutArea                 `json:"Areas,omitempty"`
	Improvements map[string]HideoutImprovement `json:"Improvements,omitempty"`
}

type HideoutArea struct {
	Type                  int               `json:"type,omitempty"`
	Level                 int               `json:"level,omitempty"`
	Active                bool              `json:"active,omitempty"`
	PassiveBonusesEnabled bool              `json:"passiveBonusesEnabled,omitempty"`
	CompleteTime          int               `json:"completeTime,omitempty"`
	Constructing          bool              `json:"constructing,omitempty"`
	Slots                 []HideoutAreaSlot `json:"slots,omitempty"`
	LastRecipe            string            `json:"lastRecipe,omitempty"`
}

type HideoutAreaSlot struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type HideoutImprovement struct {
	Completed                bool   `json:"completed,omitempty"`
	ImproveCompleteTimestamp string `json:"improveCompleteTimestamp,omitempty"`
}

type HideoutProduction struct {
	Progress       int
	InProgress     bool
	RecipeId       string
	Products       []string
	SkipTime       int
	ProductionTime int
	StartTimestamp int
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
	DamageType int    `json:"DamageType,omitempty"`
	Side       int    `json:"Side,omitempty"`
	Role       int    `json:"Role,omitempty"`
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
	SessionExperienceMult  int           `json:"SessionExperienceMult,omitempty"`
	ExperienceBonusMult    int           `json:"ExperienceBonusMult,omitempty"`
	TotalSessionExperience int           `json:"TotalSessionExperience,omitempty"`
	LastSessionDate        int           `json:"LastSessionDate,omitempty"`
	Aggressor              interface{}   `json:"Aggressor,omitempty"`
	DroppedItems           []interface{} `json:"DroppedItems,omitempty"`
	FoundInRaidItems       []interface{} `json:"FoundInRaidItems,omitempty"`
	Victims                []interface{} `json:"Victims,omitempty"`
	CarriedQuestItems      []interface{} `json:"CarriedQuestItems,omitempty"`
	DamageHistory          DamageHistory `json:"DamageHistory,omitempty"`
	DeathCause             DeathCause    `json:"DeathCause,omitempty"`
	LastPlayerState        interface{}   `json:"LastPlayerState,omitempty"`
	TotalInGameTime        int           `json:"TotalInGameTime,omitempty"`
	SurvivorClass          string        `json:"SurvivorClass,omitempty"`
}

type PlayerSkill struct {
	ID                        string `json:"Id,omitempty"`
	Progress                  int    `json:"Progress,omitempty"`
	PointsEarnedDuringSession int    `json:"PointsEarnedDuringSession,omitempty"`
	LastAccess                int    `json:"LastAccess,omitempty"`
}

type PlayerSkills struct {
	Common    []PlayerSkill `json:"Common,omitempty"`
	Mastering []PlayerSkill `json:"Mastering,omitempty"`
	Points    int           `json:"Points,omitempty"`
}

type PlayerInventory struct {
	Items           []Item            `json:"items,omitempty"`
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
	Level                   int            `json:"Level,omitempty"`
	Experience              int            `json:"Experience,omitempty"`
	RegistrationDate        int64          `json:"RegistrationDate,omitempty"`
	GameVersion             string         `json:"GameVersion,omitempty"`
	AccountType             int            `json:"AccountType,omitempty"`
	MemberCategory          int            `json:"MemberCategory,omitempty"`
	LockedMoveCommands      bool           `json:"lockedMoveCommands,omitempty"`
	SavageLockTime          int64          `json:"SavageLockTime,omitempty"`
	LastTimePlayedAsSavage  int64          `json:"LastTimePlayedAsSavage,omitempty"`
	Settings                PlayerSettings `json:"Settings,omitempty"`
	NicknameChangeDate      int64          `json:"NicknameChangeDate,omitempty"`
	NeedWipeOptions         []string       `json:"NeedWipeOptions,omitempty"`
	LastCompletedWipe       Event          `json:"lastCompletedWipe,omitempty"`
	LastCompletedEvent      Event          `json:"lastCompletedEvent,omitempty"`
	BannedState             bool           `json:"BannedState,omitempty"`
	BannedUntil             int64          `json:"BannedUntil,omitempty"`
	IsStreamerModeAvailable bool           `json:"IsStreamerModeAvailable,omitempty"`
	Bans                    []string       `json:"Bans,omitempty"`
}

type PlayerSettings struct {
	Role            string  `json:"Role,omitempty"`
	BotDifficulty   string  `json:"BotDifficulty,omitempty"`
	Experience      int     `json:"Experience,omitempty"`
	StandingForKill float64 `json:"StandingForKill,omitempty"`
	AggressorBonus  float64 `json:"AggressorBonus,omitempty"`
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
	UpdateTime  int64                     `json:"UpdateTime,omitempty"`
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
