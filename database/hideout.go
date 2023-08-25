package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"path/filepath"
	"strings"
)

var hideout = structs.Hideout{}

// 0: hideout.Areas[index] -> setHideoutAreasIndex()
var areaIndex = map[int]interface{}{}

// 0 (Area): "_id of production": "production", etc. -> setHideoutProductionIndex()
var productionsIndex = map[int]map[string]interface{}{}

func GetHideout() *structs.Hideout {
	return &hideout
}

func setHideout() {

	dynamic := make(map[string]json.RawMessage)
	files, err := tools.GetFilesFrom(hideoutPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		name := strings.TrimSuffix(file, ".json")
		raw := tools.GetJSONRawMessage(filepath.Join(hideoutPath, file))
		dynamic[name] = raw
	}

	jsonData, err := json.Marshal(dynamic)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &hideout)
	if err != nil {
		panic(err)
	}
}
