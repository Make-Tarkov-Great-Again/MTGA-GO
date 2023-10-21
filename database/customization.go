package database

import (
	"log"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

// #region Customization getters

var customizations = map[string]*Customization{}

func GetCustomizations() map[string]*Customization {
	return customizations
}

func GetCustomization(id string) *Customization {
	customization, ok := customizations[id]
	if !ok {
		log.Println("Customization with ID", id, "does not exist, returning nil!")
		return nil
	}
	return customization
}

// #endregion

// #region Customization setters

func setCustomization() {
	raw := tools.GetJSONRawMessage(customizationPath)
	err := json.Unmarshal(raw, &customizations)
	if err != nil {
		log.Fatalln("Set Customization:", err)
	}
}

// #endregion

// #region Customization structs

type Customization struct {
	ID     string                  `json:"_id"`
	Name   string                  `json:"_name"`
	Parent string                  `json:"_parent"`
	Type   string                  `json:"_type"`
	Proto  *string                 `json:"_proto,omitempty"`
	Props  CustomizationProperties `json:"_props"`
}

type CustomizationProperties struct {
	BodyPart            *string     `json:"BodyPart,omitempty"`
	Description         *string     `json:"Description,omitempty"`
	AvailableAsDefault  *bool       `json:"AvailableAsDefault,omitempty"`
	IntegratedArmorVest *bool       `json:"IntegratedArmorVest,omitempty"`
	Name                *string     `json:"Name,omitempty"`
	Prefab              interface{} `json:"Prefab,omitempty"`
	ShortName           *string     `json:"ShortName,omitempty"`
	Side                *[]string   `json:"Side,omitempty"`
	WatchPosition       *XYZ        `json:"WatchPosition,omitempty"`
	WatchPrefab         *Prefab     `json:"WatchPrefab,omitempty"`
	WatchRotation       *XYZ        `json:"WatchRotation,omitempty"`
	Feet                *string     `json:"Feet,omitempty"`
	Body                *string     `json:"Body,omitempty"`
	Hands               *string     `json:"Hands,omitempty"`
}

// #endregion
