package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"path/filepath"
	"strings"
)

var hideout = structs.Hideout{}

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
