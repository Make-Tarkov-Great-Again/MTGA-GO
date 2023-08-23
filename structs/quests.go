package structs

type Quests struct {
	Quests map[string]Quest
}

type Quest struct {
	Dialogue   map[string]string
	Conditions QuestConditions
	Rewards    QuestRewardsConditions
}

type QuestRewardsConditions struct {
	Start   []map[string]interface{}
	Success []map[string]interface{}
	Fail    []map[string]interface{}
}

type QuestRewards struct {
	Experience       QuestRewardExperience
	Item             map[string]QuestRewardItem
	AssortmentUnlock map[string]QuestRewardAssortUnlock
}

type QuestRewardItem struct {
	FindInRaid bool
	Items      []map[string]interface{}
	Target     string
	Value      string
}

type QuestRewardAssortUnlock struct {
	Items        []map[string]interface{}
	LoyaltyLevel int
	Target       string
	TraderID     string
}

type QuestRewardTraderStanding struct {
	Target string
	Value  string
}

type QuestRewardExperience struct {
	Value string
}

type QuestAvailabilityConditions struct {
	Start  map[string]QuestConditions
	Finish map[string]QuestConditions
	Fail   map[string]QuestConditions
}

type QuestConditions struct {
	Level         LevelCondition
	Quest         map[string]QuestCondition
	TraderLoyalty map[string]LevelCondition
	HandoverItem  map[string]HandoverCondition
}

type HandoverCondition struct {
	FindItemInRaid      bool
	ItemToHandover      string
	Amount              int
	MaxDurabilityOfItem float32
	MinDurabilityOfItem float32
}

type QuestCondition struct {
	QuestID         string
	Status          int
	PreviousQuestID string
}

type LevelCondition struct {
	CompareMethod string
	Level         string
}
