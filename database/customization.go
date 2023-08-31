package database

import (
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var customizations map[string]interface{}

func GetCustomizations() map[string]interface{} {
	return customizations
}

func GetCustomization(id string) map[string]interface{} {
	customization, ok := customizations[id]
	if !ok {
		return nil
	}
	return customization.(map[string]interface{})
}

func setCustomization() {
	raw := tools.GetJSONRawMessage(customizationPath)
	err := json.Unmarshal(raw, &customizations)
	if err != nil {
		panic(err)
	}
}
