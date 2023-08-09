package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
)

var customization = make(map[string]*structs.Customization)

func GetCustomization() map[string]*structs.Customization {
	return customization
}

func setCustomization() {
	raw := tools.GetJSONRawMessage(customizationPath)
	err := json.Unmarshal(raw, &customization)
	if err != nil {
		panic(err)
	}
}
