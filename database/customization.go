package database

import (
	"MT-GO/tools"
	"encoding/json"
)

var customization map[string]interface{}

func GetCustomization() map[string]interface{} {
	return customization
}

func setCustomization() {
	raw := tools.GetJSONRawMessage(customizationPath)
	err := json.Unmarshal(raw, &customization)
	if err != nil {
		panic(err)
	}
}
