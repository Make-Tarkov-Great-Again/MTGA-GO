package structs

type Globals struct {
	Config               map[string]interface{} `json:"config"`
	BotPresets           [18]interface{}        `json:"bot_presets"`
	BotWeaponScatterings [4]interface{}         `json:"BotWeaponScatterings"`
	ItemPresets          map[string]interface{} `json:"ItemPresets"`
}
