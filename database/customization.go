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

type Customization struct {
	ID     string                 `json:"_id"`
	Name   string                 `json:"_name"`
	Parent string                 `json:"_parent"`
	Type   string                 `json:"_type"`
	Proto  string                 `json:"_proto"`
	Props  map[string]interface{} `json:"_props"`
}
