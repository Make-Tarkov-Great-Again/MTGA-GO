package database

import (
	"MT-GO/structs"
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var locations = structs.Locations{}

func GetLocations() *structs.Locations {
	return &locations
}

func setLocations() {
	raw := tools.GetJSONRawMessage(locationsFilePath)
	err := json.Unmarshal(raw, &locations)
	if err != nil {
		panic(err)
	}
}
