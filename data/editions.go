package data

import (
	"MT-GO/tools"
	"log"
	"path/filepath"

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

		switch file {
		case "storage.json":
			storage := new(EditionStorage)
			if err := json.Unmarshal(raw, storage); err != nil {
				log.Println(err)
				return
			}
			edition.Storage = storage
			continue

		case "character_usec.json":
			template := new(Character)
			if err := json.Unmarshal(raw, template); err != nil {
				log.Println(err)
				return
			}
			edition.Usec = template
			continue

		case "character_bear.json":
			template := new(Character)
			if err := json.Unmarshal(raw, template); err != nil {
				log.Println(err)
				return
			}
			edition.Bear = template
			continue

		default:
			log.Println("huh")
		}
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
