package database

import (
	"fmt"
	"log"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

// #region Customization getters

var customizations = map[string]*Customization{}

func GetCustomizations() map[string]*Customization {
	return customizations
}

func GetCustomizationByID(id string) (*Customization, error) {
	customization, ok := customizations[id]
	if !ok {
		return nil, fmt.Errorf("Customization with ID", id, "does not exist")
	}
	return customization, nil
}

func CustomizationClone(item string) *Customization {
	input, err := GetCustomizationByID(item)
	if err != nil {
		log.Fatalln(err)
	}

	clone := new(Customization)

	data, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &clone); err != nil {
		log.Fatal(err)
	}

	return clone
}

func (c *Customization) Clone() *Customization {
	clone := new(Customization)

	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &clone); err != nil {
		log.Fatal(err)
	}

	return clone
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
	Proto  string                  `json:"_proto,omitempty"`
	Props  CustomizationProperties `json:"_props"`
}

type CustomizationProperties struct {
	BodyPart            string   `json:"BodyPart,omitempty"`
	Description         string   `json:"Description,omitempty"`
	AvailableAsDefault  bool     `json:"AvailableAsDefault,omitempty"`
	IntegratedArmorVest bool     `json:"IntegratedArmorVest,omitempty"`
	Name                string   `json:"Name,omitempty"`
	Prefab              any      `json:"Prefab,omitempty"`
	ShortName           string   `json:"ShortName,omitempty"`
	Side                []string `json:"Side,omitempty"`
	WatchPosition       *Vector3 `json:"WatchPosition,omitempty"`
	WatchPrefab         *Prefab  `json:"WatchPrefab,omitempty"`
	WatchRotation       *Vector3 `json:"WatchRotation,omitempty"`
	Feet                string   `json:"Feet,omitempty"`
	Body                string   `json:"Body,omitempty"`
	Hands               string   `json:"Hands,omitempty"`
}

// #endregion
