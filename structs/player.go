package structs

type PlayerScavTemplate struct {
	ID                string                 `json:"_id"`
	AID               int                    `json:"aid"`
	Savage            *string                `json:"savage"`
	Info              PlayerInfo             `json:"Info"`
	Customization     PlayerCustomization    `json:"Customization"`
	Health            HealthInfo             `json:"Health"`
	Inventory         InventoryInfo          `json:"Inventory"`
	Skills            PlayerSkills           `json:"Skills"`
	Stats             PlayerStats            `json:"Stats"`
	Encyclopedia      map[string]bool        `json:"Encyclopedia"`
	ConditionCounters ConditionCounters      `json:"ConditionCounters"`
	BackendCounters   map[string]interface{} `json:"BackendCounters"`
	InsuredItems      []interface{}          `json:"InsuredItems"`
	Hideout           interface{}            `json:"Hideout"`
	Bonuses           []interface{}          `json:"Bonuses"`
	Notes             struct {
		Notes [][]interface{} `json:"Notes"`
	} `json:"Notes"`
	Quests       []interface{}     `json:"Quests"`
	RagfairInfo  PlayerRagfairInfo `json:"RagfairInfo"`
	WishList     []interface{}     `json:"WishList"`
	TradersInfo  []interface{}     `json:"TradersInfo"`
	UnlockedInfo struct {
		UnlockedProductionRecipe []interface{} `json:"unlockedProductionRecipe"`
	} `json:"UnlockedInfo"`
}

type PlayerTemplate struct {
	ID                string                 `json:"_id"`
	AID               int                    `json:"aid"`
	Savage            *string                `json:"savage"`
	Info              PlayerInfo             `json:"Info"`
	Customization     PlayerCustomization    `json:"Customization"`
	Health            HealthInfo             `json:"Health"`
	Inventory         InventoryInfo          `json:"Inventory"`
	Skills            PlayerSkills           `json:"Skills"`
	Stats             PlayerStats            `json:"Stats"`
	Encyclopedia      map[string]bool        `json:"Encyclopedia"`
	ConditionCounters ConditionCounters      `json:"ConditionCounters"`
	BackendCounters   map[string]interface{} `json:"BackendCounters"`
	InsuredItems      []InsuredItem          `json:"InsuredItems"`
	Hideout           PlayerHideout          `json:"Hideout"`
	Bonuses           []Bonus                `json:"Bonuses"`
	Notes             struct {
		Notes [][]interface{} `json:"Notes"`
	} `json:"Notes"`
	Quests       []map[string]interface{}     `json:"Quests"`
	RagfairInfo  PlayerRagfairInfo            `json:"RagfairInfo"`
	WishList     []string                     `json:"WishList"`
	TradersInfo  map[string]PlayerTradersInfo `json:"TradersInfo"`
	UnlockedInfo struct {
		UnlockedProductionRecipe []interface{} `json:"unlockedProductionRecipe"`
	} `json:"UnlockedInfo"`
}

type PlayerTradersInfo struct {
	Unlocked bool    `json:"unlocked"`
	Disabled bool    `json:"disabled"`
	SalesSum float32 `json:"salesSum"`
	Standing float32 `json:"standing"`
}

type PlayerRagfairInfo struct {
	Rating          float64       `json:"rating"`
	IsRatingGrowing bool          `json:"isRatingGrowing"`
	Offers          []interface{} `json:"offers"`
}

type PlayerHideoutArea struct {
	Type                  int           `json:"type"`
	Level                 int           `json:"level"`
	Active                bool          `json:"active"`
	PassiveBonusesEnabled bool          `json:"passiveBonusesEnabled"`
	CompleteTime          int           `json:"completeTime"`
	Constructing          bool          `json:"constructing"`
	Slots                 []interface{} `json:"slots"`
	LastRecipe            string        `json:"lastRecipe"`
}
type PlayerHideout struct {
	Production  map[string]interface{} `json:"Production"`
	Areas       []PlayerHideoutArea    `json:"Areas"`
	Improvement map[string]interface{} `json:"Improvement"`
	//Seed        int                    `json:"Seed"`
}

type ConditionCounters struct {
	Counters []interface{} `json:"Counters"`
}

type PlayerStats struct {
	Eft EftStats `json:"Eft"`
}

type StatCounters struct {
	Items []interface{} `json:"Items"`
}

type EftStats struct {
	SessionCounters        map[string]interface{} `json:"SessionCounters"`
	OverallCounters        map[string]interface{} `json:"OverallCounters"`
	SessionExperienceMult  int                    `json:"SessionExperienceMult"`
	ExperienceBonusMult    int                    `json:"ExperienceBonusMult"`
	TotalSessionExperience int                    `json:"TotalSessionExperience"`
	LastSessionDate        int                    `json:"LastSessionDate"`
	Aggressor              map[string]interface{} `json:"Aggressor"`
	DroppedItems           []interface{}          `json:"DroppedItems"`
	FoundInRaidItems       []interface{}          `json:"FoundInRaidItems"`
	Victims                []interface{}          `json:"Victims"`
	CarriedQuestItems      []interface{}          `json:"CarriedQuestItems"`
	DamageHistory          map[string]interface{} `json:"DamageHistory"`
	LastPlayerState        *float32               `json:"LastPlayerState"`
	TotalInGameTime        int                    `json:"TotalInGameTime"`
	SurvivorClass          string                 `json:"SurvivorClass"`
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
	Nickname               string                 `json:"Nickname"`
	LowerNickname          string                 `json:"LowerNickname"`
	Side                   string                 `json:"Side"`
	Voice                  string                 `json:"Voice"`
	Level                  int                    `json:"Level"`
	Experience             int                    `json:"Experience"`
	RegistrationDate       int                    `json:"RegistrationDate"`
	GameVersion            string                 `json:"GameVersion"`
	AccountType            int                    `json:"AccountType"`
	MemberCategory         int                    `json:"MemberCategory"`
	LockedMoveCommands     bool                   `json:"lockedMoveCommands"`
	SavageLockTime         int                    `json:"SavageLockTime"`
	LastTimePlayedAsSavage int                    `json:"LastTimePlayedAsSavage"`
	Settings               map[string]interface{} `json:"Settings"`
	NicknameChangeDate     int                    `json:"NicknameChangeDate"`
	NeedWipeOptions        []interface{}          `json:"NeedWipeOptions"`
	LastCompletedWipe      struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedWipe"`
	LastCompletedEvent struct {
		Oid string `json:"$oid"`
	} `json:"lastCompletedEvent"`
	BannedState             bool          `json:"BannedState"`
	BannedUntil             int           `json:"BannedUntil"`
	IsStreamerModeAvailable bool          `json:"IsStreamerModeAvailable"`
	SquadInviteRestriction  bool          `json:"SquadInviteRestriction"`
	Bans                    []interface{} `json:"Bans"`
}

type InfoSettings struct {
	Role            string  `json:"Role"`
	BotDifficulty   string  `json:"BotDifficulty"`
	Experience      int     `json:"Experience"`
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
	ID         string `json:"id"`
	Type       string `json:"type"`
	TemplateID string `json:"templateId"`
}

type InventoryInfo struct {
	Items              []interface{} `json:"items"`
	Equipment          string        `json:"equipment"`
	Stash              string        `json:"stash"`
	SortingTable       string        `json:"sortingTable"`
	QuestRaidItems     string        `json:"questRaidItems"`
	QuestStashItems    string        `json:"questStashItems"`
	FastPanel          interface{}   `json:"fastPanel"`
	HideoutAreaStashes interface{}   `json:"hideoutAreaStashes"`
}

type HealthInfo struct {
	Hydration   CurrMaxHealth   `json:"Hydration"`
	Energy      CurrMaxHealth   `json:"Energy"`
	Temperature CurrMaxHealth   `json:"Temperature"`
	BodyParts   BodyPartsHealth `json:"BodyParts"`
	UpdateTime  int             `json:"UpdateTime"`
}

type HealthOf struct {
	Health CurrMaxHealth `json:"Health"`
}

type CurrMaxHealth struct {
	Current float32 `json:"Current"`
	Maximum float32 `json:"Maximum"`
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
