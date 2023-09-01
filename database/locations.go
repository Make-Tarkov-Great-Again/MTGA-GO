package database

import (
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var locations = Locations{}

func GetLocations() *Locations {
	return &locations
}

func setLocations() {
	raw := tools.GetJSONRawMessage(locationsFilePath)
	err := json.Unmarshal(raw, &locations)
	if err != nil {
		panic(err)
	}
}

type Locations struct {
	Locations map[string]interface{} `json:"locations"`
	Paths     []interface{}          `json:"paths"`
}
