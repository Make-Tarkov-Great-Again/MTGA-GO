package structs

type Quest struct {
	QuestName                  string                 `json:"QuestName,omitempty"`
	ID                         string                 `json:"_id,omitempty"`
	AcceptPlayerMessage        string                 `json:"acceptPlayerMessage,omitempty"`
	CanShowNotificationsInGame bool                   `json:"canShowNotificationsInGame,omitempty"`
	ChangeQuestMessageText     string                 `json:"changeQuestMessageText,omitempty"`
	CompletePlayerMessage      string                 `json:"completePlayerMessage,omitempty"`
	Description                string                 `json:"description,omitempty"`
	FailMessageText            string                 `json:"failMessageText,omitempty"`
	Image                      string                 `json:"image,omitempty"`
	InstantComplete            bool                   `json:"instantComplete,omitempty"`
	IsKey                      bool                   `json:"isKey,omitempty"`
	Location                   string                 `json:"location,omitempty"`
	Name                       string                 `json:"name,omitempty"`
	Note                       string                 `json:"note,omitempty"`
	QuestStatus                map[string]interface{} `json:"questStatus,omitempty"`
	Restartable                bool                   `json:"restartable,omitempty"`
	SecretQuest                bool                   `json:"secretQuest,omitempty"`
	Side                       string                 `json:"side,omitempty"`
	StartedMessageText         string                 `json:"startedMessageText,omitempty"`
	SuccessMessageText         string                 `json:"successMessageText,omitempty"`
	TemplateID                 string                 `json:"templateId,omitempty"`
	TraderID                   string                 `json:"traderId,omitempty"`
	Type                       string                 `json:"type,omitempty"`
	Conditions                 struct {
		AvailableForFinish []Condition `json:"AvailableForFinish,omitempty"`
		AvailableForStart  []Condition `json:"AvailableForStart,omitempty"`
		Fail               []Condition `json:"Fail,omitempty"`
	} `json:"conditions,omitempty"`
	Rewards struct {
		Fail    []Reward `json:"Fail,omitempty"`
		Started []Reward `json:"Started,omitempty"`
		Success []Reward `json:"Success,omitempty"`
	} `json:"rewards,omitempty"`
}

type Condition struct {
	Parent        string         `json:"_parent,omitempty"`
	Props         ConditionProps `json:"_props,omitempty"`
	DynamicLocale bool           `json:"dynamicLocale,omitempty"`
}

type ConditionProps struct {
	Counter struct {
		Conditions []CounterCondition `json:"conditions,omitempty"`
		ID         string             `json:"id,omitempty"`
	} `json:"counter,omitempty"`
	DoNotResetIfCounterCompleted bool                  `json:"doNotResetIfCounterCompleted,omitempty"`
	DynamicLocale                bool                  `json:"dynamicLocale,omitempty"`
	ID                           string                `json:"id,omitempty"`
	Index                        int16                 `json:"index,omitempty"`
	OneSessionOnly               bool                  `json:"oneSessionOnly,omitempty"`
	ParentID                     string                `json:"parentId,omitempty"`
	Type                         string                `json:"type,omitempty"`
	Status                       []int8                `json:"status,omitempty"`
	Value                        float32               `json:"value,omitempty"`
	VisibilityConditions         []VisibilityCondition `json:"visibilityConditions,omitempty"`
}

type CounterCondition struct {
	Parent string       `json:"_parent,omitempty"`
	Props  CounterProps `json:"_props,omitempty"`
}

type CounterProps struct {
	ID     string      `json:"id,omitempty"`
	Target interface{} `json:"target,omitempty"` //can be array or string
	Value  float32     `json:"value,omitempty"`
}

type VisibilityCondition struct {
	Parent string          `json:"_parent,omitempty"`
	Props  VisibilityProps `json:"_props,omitempty"`
}

type VisibilityProps struct {
	ID     string `json:"id,omitempty"`
	Target string `json:"target,omitempty"`
}

type RewardItem PresetItem

type Reward struct {
	FindInRaid   bool         `json:"findInRaid,omitempty"`
	ID           string       `json:"id,omitempty"`
	Index        int8         `json:"index,omitempty"`
	Items        []RewardItem `json:"items,omitempty"`
	LoyaltyLevel int8         `json:"loyaltyLevel,omitempty"`
	Target       string       `json:"target,omitempty"`
	TraderID     string       `json:"traderId,omitempty"`
	Type         string       `json:"type,omitempty"`
	Value        float32      `json:"value,omitempty"`
}
