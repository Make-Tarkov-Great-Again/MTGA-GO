package structs

type Item struct {
	ID       string       `json:"_id"`
	Tpl      string       `json:"_tpl"`
	Upd      ItemUpd      `json:"upd,omitempty"`
	ParentID string       `json:"parentId,omitempty"`
	SlotID   string       `json:"slotId,omitempty"`
	Location ItemLocation `json:"location,omitempty"`
}

type PresetItem struct {
	ID       string  `json:"_id"`
	Tpl      string  `json:"_tpl"`
	Upd      ItemUpd `json:"upd,omitempty"`
	ParentID string  `json:"parentId,omitempty"`
	SlotID   string  `json:"slotId,omitempty"`
}

type ItemLocation struct {
	X          float32 `json:"x,omitempty"`
	Y          float32 `json:"y,omitempty"`
	Z          float32 `json:"z,omitempty"`
	R          float32 `json:"r,omitempty"`
	IsSearched bool    `json:"isSearched,omitempty"`
	Rotation   string  `json:"rotation,omitempty"`
}

type ItemUpd struct {
	Foldable struct {
		Folded bool `json:"Folded,omitempty"`
	} `json:"Foldable,omitempty"`
	Togglable struct {
		On bool `json:"On,omitempty"`
	} `json:"Togglable,omitempty"`
	FireMode struct {
		FireMode string `json:"FireMode,omitempty"`
	} `json:"FireMode,omitempty"`
	StackObjectsCount int `json:"StackObjectsCount,omitempty"`
	Repairable        struct {
		MaxDurability float32 `json:"MaxDurability,omitempty"`
		Durability    float32 `json:"Durability,omitempty"`
	} `json:"Repairable,omitempty"`
	Sight struct {
		ScopesCurrentCalibPointIndexes []int `json:"ScopesCurrentCalibPointIndexes,omitempty"`
		ScopesSelectedModes            []int `json:"ScopesSelectedModes,omitempty"`
		SelectedScope                  int   `json:"SelectedScope,omitempty"`
	} `json:"Sight,omitempty"`
	FoodDrink struct {
		HpPercent int `json:"HpPercent,omitempty"`
	} `json:"FoodDrink,omitempty"`
	Resource   Value `json:"Resource,omitempty"`
	SideEffect Value `json:"SideEffect,omitempty"`
	MedKit     struct {
		HpResource int `json:"HpResource,omitempty"`
	} `json:"MedKit,omitempty"`
	RepairKit struct {
		Resource int `json:"Resource,omitempty"`
	} `json:"RepairKit,omitempty"`
	Key struct {
		NumberOfUsages int `json:"NumberOfUsages,omitempty"`
	} `json:"Key,omitempty"`
	SpawnedInSession bool `json:"SpawnedInSession,omitempty"`
	Dogtag           struct {
		AccountId       string `json:"AccountId,omitempty"`
		ProfileId       string `json:"ProfileId,omitempty"`
		Nickname        string `json:"Nickname,omitempty"`
		Side            string `json:"Side,omitempty"`
		Level           int    `json:"Level,omitempty"`
		Time            string `json:"Time,omitempty"`
		Status          string `json:"Status,omitempty"`
		KillerAccountId string `json:"KillerAccountId,omitempty"`
		KillerProfileId string `json:"KillerProfileId,omitempty"`
		KillerName      string `json:"KillerName,omitempty"`
		WeaponName      string `json:"WeaponName,omitempty"`
	} `json:"Dogtag,omitempty"`
	Light struct {
		IsActive     bool `json:"IsActive,omitempty"`
		SelectedMode int  `json:"SelectedMode,omitempty"`
	} `json:"Light,omitempty"`
	Buff struct {
		Rarity              string `json:"rarity,omitempty"`
		BuffType            string `json:"buffType,omitempty"`
		Value               int    `json:"value,omitempty"`
		ThresholdDurability int    `json:"thresholdDurability,omitempty"`
	} `json:"Buff,omitempty"`
	Map struct {
		Markers []MapMarker `json:"Markers,omitempty"`
	} `json:"Map,omitempty"`
	FaceShield struct {
		Hits int `json:"Hits,omitempty"`
	} `json:"FaceShield,omitempty"`
	Tag struct {
		Color float32 `json:"Color,omitempty"`
		Name  string  `json:"Name,omitempty"`
	} `json:"Tag,omitempty"`
}

type MapMarker struct {
	X float32 `json:"X,omitempty"`
	Y float32 `json:"Y,omitempty"`
}
