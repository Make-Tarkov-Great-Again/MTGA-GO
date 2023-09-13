package database

import (
	"MT-GO/tools"
	"path/filepath"
	"strings"
	"sync"

	"github.com/goccy/go-json"
)

var editions = make(map[string]*Edition)

// #region Edition getters

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

// #endregion

// #region Edition setters

func setEditions() {
	directories, err := tools.GetDirectoriesFrom(editionsDirPath)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	editions := make(map[string]*Edition)

	for _, directory := range directories {
		wg.Add(1)
		go setEdition(directory, editionsDirPath, editions, &wg)
	}

	wg.Wait()
}

func setEdition(directory string, editionsDirPath string, editions map[string]*Edition, wg *sync.WaitGroup) {
	defer wg.Done()

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
			template := new(Character)
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

		storage := new(EditionStorage)
		err := json.Unmarshal(raw, storage)
		if err != nil {
			panic(err)
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
