package structs

type Trader struct {
	Index       TraderIndex                  `json:",omitempty"`
	Base        map[string]interface{}       `json:",omitempty"`
	Assort      *Assort                      `json:",omitempty"`
	QuestAssort map[string]map[string]string `json:",omitempty"`
	Suits       []TraderSuits                `json:",omitempty"`
	Dialogue    map[string][]string          `json:",omitempty"`
}

type TraderIndex struct {
	Assort *AssortIndex    `json:",omitempty"`
	Suits  map[string]int8 `json:",omitempty"`
}

type AssortIndex struct {
	Items       map[string]int16
	ParentItems map[string]map[string]int16 `json:",omitempty"`
}

type TraderSuits struct {
	ID           string           `json:"_id"`
	Tid          string           `json:"tid"`
	SuiteID      string           `json:"suiteId"`
	IsActive     bool             `json:"isActive"`
	Requirements SuitRequirements `json:"requirements"`
}
type SuitItemRequirements struct {
	Count          int    `json:"count"`
	Tpl            string `json:"_tpl"`
	OnlyFunctional bool   `json:"onlyFunctional"`
}
type SuitRequirements struct {
	LoyaltyLevel         int8                   `json:"loyaltyLevel"`
	ProfileLevel         int8                   `json:"profileLevel"`
	Standing             int8                   `json:"standing"`
	SkillRequirements    []interface{}          `json:"skillRequirements"`
	QuestRequirements    []string               `json:"questRequirements"`
	SuitItemRequirements []SuitItemRequirements `json:"itemRequirements"`
}

type Assort struct {
	NextResupply    int                    `json:"nextResupply"`
	BarterScheme    map[string][][]*Scheme `json:"barter_scheme"`
	Items           []*AssortItem          `json:"items"`
	LoyalLevelItems map[string]int8        `json:"loyal_level_items"`
}

type AssortItem struct {
	ID       string `json:"_id"`
	Tpl      string `json:"_tpl"`
	ParentID string `json:"parentId"`
	SlotID   string `json:"slotId"`
	Upd      struct {
		BuyRestrictionCurrent interface{} `json:"BuyRestrictionCurrent,omitempty"`
		BuyRestrictionMax     interface{} `json:"BuyRestrictionMax,omitempty"`
		StackObjectsCount     int         `json:"StackObjectsCount,omitempty"`
		UnlimitedCount        bool        `json:"UnlimitedCount,omitempty"`
		FireMode              struct {
			FireMode string `json:"FireMode"`
		} `json:"FireMode,omitempty"`
		Foldable struct {
			Folded bool `json:"Folded,omitempty"`
		} `json:"Foldable,omitempty"`
	} `json:"upd,omitempty"`
}

type Scheme struct {
	Tpl   string  `json:"_tpl"`
	Count float32 `json:"count"`
}
