package database

import (
	"log"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var locations = Locations{}

// #region Location getters

func GetLocations() *Locations {
	return &locations
}

// #endregion

// #region Location setters

func setLocations() {
	raw := tools.GetJSONRawMessage(locationsFilePath)
	err := json.Unmarshal(raw, &locations)
	if err != nil {
		log.Fatalln(err)
	}
}

// #endregion

// #region Location structs

type Locations struct {
	Locations map[string]interface{} `json:"locations"`
	Paths     []interface{}          `json:"paths"`
}

// #endregion
