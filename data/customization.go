package data

import (
	"fmt"
	"github.com/alphadose/haxmap"
	"log"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

// #region Customization getters

func GetCustomizations() *haxmap.Map[string, *Customization] {
	return db.customization
}

func GetCustomizationByID(id string) (*Customization, error) {
	customization, ok := db.customization.Get(id)
	if !ok {
		return nil, fmt.Errorf("customization with ID %s does not exist", id)
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
	db.customization = haxmap.New[string, *Customization]() //make(map[string]*Customization)
	raw := tools.GetJSONRawMessage(customizationPath)
	if err := json.UnmarshalNoEscape(raw, &db.customization); err != nil {
		msg := tools.CheckParsingError(raw, err)
		log.Fatalln(msg)
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
