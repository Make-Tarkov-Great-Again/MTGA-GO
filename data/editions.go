package data

import (
	"log"
	"path/filepath"
	"strings"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var editions = make(map[string]*Edition)

// #region Edition getters

func GetEditions() map[string]*Edition {
	return editions
}

func GetEditionByName(version string) *Edition {
	edition, _ := editions[version]
	return edition
}

// #endregion

// #region Edition setters

func setEditions() {
	directories, err := tools.GetDirectoriesFrom(editionsDirPath)
	if err != nil {
		log.Println(err)
	}

	for directory := range directories {
		setEdition(directory, editionsDirPath)
	}
}

func setEdition(directory string, editionsDirPath string) {
	editionPath := filepath.Join(editionsDirPath, directory)
	files, err := tools.GetFilesFrom(editionPath)
	if err != nil {
		log.Println(err)
	}

	edition := new(Edition)

	for file := range files {
		raw := tools.GetJSONRawMessage(filepath.Join(editionPath, file))
		removeJSON := strings.TrimSuffix(file, ".json")

		if strings.Contains(removeJSON, "character_") {
			template := new(Character)
			err := json.Unmarshal(raw, template)
			if err != nil {
				log.Println(err)
			}

			name := strings.TrimPrefix(removeJSON, "character_")
			if name == "bear" {
				edition.Bear = template
			} else {
				edition.Usec = template
			}
			continue
		}

		storage := new(EditionStorage)
		err := json.Unmarshal(raw, storage)
		if err != nil {
			log.Println(err)
		}
		edition.Storage = storage
	}

	editions[directory] = edition
}

// #endregion

// #region Edition structs

type Edition struct {
	Bear    *Character      `json:"bear"`
	Usec    *Character      `json:"usec"`
	Storage *EditionStorage `json:"storage"`
}

type EditionStorage struct {
	Bear []string `json:"bear"`
	Usec []string `json:"usec"`
}

// #endregion
