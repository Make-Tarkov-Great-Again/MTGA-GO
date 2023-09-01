package database

import (
	"MT-GO/tools"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
)

var editions = make(map[string]*Edition)

func GetEditions() map[string]*Edition {
	return editions
}

func GetEdition(version string) *Edition {
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
		edition := &Edition{}

		for _, file := range files {

			raw := tools.GetJSONRawMessage(filepath.Join(editionPath, file))
			removeJSON := strings.TrimSuffix(file, ".json")
			if strings.Contains(removeJSON, "character_") {
				template := &Character{}
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

			storage := &EditionStorage{}
			err := json.Unmarshal(raw, storage)
			if err != nil {
				panic(err)
			}
			edition.Storage = storage
		}
		editions[directory] = edition
	}
}

type Edition struct {
	Bear    *Character      `json:"bear"`
	Usec    *Character      `json:"usec"`
	Storage *EditionStorage `json:"storage"`
}

type EditionStorage struct {
	Bear []string `json:"bear"`
	Usec []string `json:"usec"`
}
