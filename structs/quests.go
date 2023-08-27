package structs

type Quest struct {
	Name       string
	Dialogue   QuestDialogues                    `json:",omitempty"`
	Conditions *QuestAvailabilityConditions      `json:",omitempty"`
	Rewards    QuestRewardAvailabilityConditions `json:",omitempty"`
}

type QuestDialogues struct {
	Description string
	Accepted    string
	Started     string
	Complete    string
	Success     string
	Fail        string
}

type QuestAvailabilityConditions struct {
	AvailableForStart  *QuestConditionTypes `json:"AvailableForStart,omitempty"`
	AvailableForFinish *QuestConditionTypes `json:"AvailableForFinish,omitempty"`
	Fail               *QuestConditionTypes `json:"Fail,omitempty"`
}

type QuestConditionTypes struct {
	Level          *LevelCondition               `json:"Level,omitempty"`
	Quest          map[string]*QuestCondition    `json:"Quest,omitempty"`
	TraderLoyalty  map[string]*LevelCondition    `json:"TraderLoyalty,omitempty"`
	TraderStanding map[string]*LevelCondition    `json:"TraderStanding,omitempty"`
	HandoverItem   map[string]*HandoverCondition `json:"HandoverItem,omitempty"`
	WeaponAssembly map[string]*HandoverCondition `json:"WeaponAssembly,omitempty"`
	FindItem       map[string]*HandoverCondition `json:"FindItem,omitempty"`
	Skills         map[string]*LevelCondition    `json:"Skill,omitempty"`
}

type HandoverCondition struct {
	ItemToHandover string
	Amount         int
}

type QuestCondition struct {
	Status          int
	PreviousQuestID string
}

type LevelCondition struct {
	CompareMethod string
	Level         float64
}

type QuestRewardAvailabilityConditions struct {
	Start   *QuestRewards `json:"Started,omitempty"`
	Success *QuestRewards `json:"Success,omitempty"`
	Fail    *QuestRewards `json:"Fail,omitempty"`
}

type QuestRewards struct {
	Experience            int                                     `json:"Experience,omitempty"`
	Items                 map[string]*QuestRewardItem             `json:"Item,omitempty"`
	AssortmentUnlock      string                                  `json:"AssortmentUnlock,omitempty"`
	TraderStanding        map[string]*float64                     `json:"TraderStanding,omitempty"`
	TraderStandingRestore map[string]*float64                     `json:"TraderStandingRestore,omitempty"`
	TraderUnlock          string                                  `json:"TraderUnlock,omitempty"`
	Skills                map[string]*int                         `json:"Skills,omitempty"`
	ProductionScheme      map[string]*QuestRewardProductionScheme `json:"ProductionScheme,omitempty"`
}

type QuestRewardProductionScheme struct {
	Item         string
	LoyaltyLevel int
	AreaID       int
}

type QuestRewardItem struct {
	FindInRaid bool
	Items      []map[string]interface{}
	Value      int
}

type QuestRewardAssortUnlock struct {
	Items        []map[string]interface{}
	LoyaltyLevel int
}
