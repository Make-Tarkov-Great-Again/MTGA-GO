package structs

type PlayerTemplate struct {
	ID            string                 `json:"_id"`
	AID           int                    `json:"aid"`
	Savage        string                 `json:"savage"`
	Info          map[string]interface{} `json:"Info"`
	Customization struct {
		Head  string `json:"Head"`
		Body  string `json:"Body"`
		Feet  string `json:"Feet"`
		Hands string `json:"Hands"`
	} `json:"Customization"`
	Health          HealthInfo             `json:"Health"`
	Inventory       InventoryInfo          `json:"Inventory"`
	Skills          map[string]interface{} `json:"Skills"`
	Stats           map[string]interface{} `json:"Stats"`
	Encyclopedia    map[string]bool        `json:"Encyclopedia"`
	BackendCounters map[string]interface{} `json:"BackendCounters"`
	InsuredItems    []interface{}          `json:"InsuredItems"`
	Bonuses         []interface{}          `json:"Bonuses"`
	Notes           struct {
		Notes [][]interface{} `json:"Notes"`
	} `json:"Notes"`
	Quests        []interface{} `json:"Quests"`
	WishList      []string      `json:"WishList"`
	SurvivorClass string        `json:"SurvivorClass"`
}

type InventoryInfo struct {
	Items           []interface{} `json:"items"`
	Equipment       string        `json:"equipment"`
	Stash           string        `json:"stash"`
	SortingTable    string        `json:"sortingTable"`
	QuestRaidItems  string        `json:"questRaidItems"`
	QuestStashItems string        `json:"questStashItems"`
	FastPanel       interface{}   `json:"fastPanel"`
}

type HealthInfo struct {
	Hydration   CurrMaxHealth `json:"Hydration"`
	Energy      CurrMaxHealth `json:"Energy"`
	Temperature CurrMaxHealth `json:"Temperature"`
	BodyParts   BodyParts     `json:"BodyParts"`
	UpdateTime  int           `json:"UpdateTime"`
}

type HealthOf struct {
	Health CurrMaxHealth `json:"Health"`
}

type CurrMaxHealth struct {
	Current float32 `json:"Current"`
	Maximum float32 `json:"Maximum"`
}

type BodyParts struct {
	Head     HealthOf `json:"Head"`
	Chest    HealthOf `json:"Chest"`
	Stomach  HealthOf `json:"Stomach"`
	LeftArm  HealthOf `json:"LeftArm"`
	RightArm HealthOf `json:"RightArm"`
	LeftLeg  HealthOf `json:"LeftLeg"`
	RightLeg HealthOf `json:"RightLeg"`
}
