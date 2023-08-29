package structs

type Trader struct {
	Base        map[string]interface{}
	Assort      *Assort
	QuestAssort map[string][]string
	Suits       []map[string]interface{}
	Dialogue    map[string][]string
}

type Assort struct {
	NextResupply    int                    `json:"nextResupply"`
	BarterScheme    map[string][][]*Scheme `json:"barter_scheme"`
	Items           []*AssortItem          `json:"items"`
	LoyalLevelItems map[string]int         `json:"loyal_level_items"`
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
