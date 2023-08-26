package structs

type Profile struct {
	Account   *Account
	Character *PlayerTemplate
	Storage   *Storage
	Dialogue  map[string]interface{}
}

type Storage struct {
	//ID        string                 `json:"_id"`
	Suites    []string      `json:"suites"`
	Builds    Builds        `json:"builds"`
	Insurance []interface{} `json:"insurance"`
	Mailbox   []interface{} `json:"mailbox"`
}

type Builds struct {
	EquipmentBuilds []EquipmentBuild `json:"equipmentBuilds"`
	WeaponBuilds    []interface{}    `json:"weaponBuilds"`
}

type EquipmentBuild struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Root      string        `json:"root"`
	Items     []interface{} `json:"items"`
	Type      string        `json:"type"`
	FastPanel []interface{} `json:"fastPanel"`
	BuildType string        `json:"buildType"`
}
