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

func GetEdition(version string) *structs.Edition {
	edition, ok := editions[version]
	if !ok {
		return nil
	}
	return edition
}

func setEditions() {
	directories, err := tools.GetDirectoriesFrom(editionsDirPath)
	if err != nil {
		panic(err)
	}

	for _, directory := range directories {

		editionPath := filepath.Join(editionsDirPath, directory)
		files, err := tools.GetFilesFrom(editionPath)
		if err != nil {
			panic(err)
		}
		edition := &structs.Edition{}

		for _, file := range files {

			raw := tools.GetJSONRawMessage(filepath.Join(editionPath, file))
			removeJSON := strings.TrimSuffix(file, ".json")
			if strings.Contains(removeJSON, "character_") {
				template := &structs.PlayerTemplate{}
				err := json.Unmarshal(raw, template)
				if err != nil {
					panic(err)
				}

				name := strings.TrimPrefix(removeJSON, "character_")
				if name == "bear" {
					edition.Bear = template
				} else {
					edition.Usec = template
				}
				continue
			}

			storage := &structs.Storage{}
			err := json.Unmarshal(raw, storage)
			if err != nil {
				panic(err)
			}
		}
		editions[directory] = edition
	}
}
