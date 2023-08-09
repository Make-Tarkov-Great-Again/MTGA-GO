package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"path/filepath"
	"strings"
)

var editions = make(map[string]*structs.Edition)

func GetEditions() map[string]*structs.Edition {
	return editions
}

func setEditions() {
	directories, err := tools.GetDirectoriesFrom(editionsDirPath)
	if err != nil {
		panic(err)
	}

	for _, directory := range directories {
		edition := structs.Edition{}

		editionPath := filepath.Join(editionsDirPath, directory)
		files, err := tools.GetFilesFrom(editionPath)
		if err != nil {
			panic(err)
		}

		dynamic := make(map[string]interface{})
		for _, file := range files {
			raw := tools.GetJSONRawMessage(filepath.Join(editionPath, file))
			removeJSON := strings.TrimSuffix(file, ".json")
			name := strings.TrimPrefix(removeJSON, "character_")

			dynamic[name] = raw
		}

		jsonData, err := json.Marshal(dynamic)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(jsonData, &edition)
		if err != nil {
			panic(err)
		}

		editions[directory] = &edition
	}
}
