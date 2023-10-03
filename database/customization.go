package database

import (
	"log"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

// #region Customization getters

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

// #endregion

// #region Customization setters

func setCustomization() {
	raw := tools.GetJSONRawMessage(customizationPath)
	err := json.Unmarshal(raw, &customizations)
	if err != nil {
		log.Fatalln(err)
	}
}

// #endregion

// #region Customization structs

type Customization struct {
	ID     string                 `json:"_id"`
	Name   string                 `json:"_name"`
	Parent string                 `json:"_parent"`
	Type   string                 `json:"_type"`
	Proto  string                 `json:"_proto"`
	Props  map[string]interface{} `json:"_props"`
}

// #endregion
