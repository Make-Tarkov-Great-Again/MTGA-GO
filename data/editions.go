package data

import (
	"MT-GO/tools"
	"log"
	"path/filepath"

	"github.com/goccy/go-json"
)

// #region Edition getters

func GetEditionByName(version string) *Edition {
	edition, _ := db.edition[version]
	return edition
}

// #endregion

// #region Edition setters

func setEditions() {
	editions := make(map[string]*Edition)
	directories, err := tools.GetDirectoriesFrom(editionsDirPath)
	if err != nil {
		log.Fatalln(err)
	}

	for directory := range directories {
		editions[directory] = setEdition(filepath.Join(editionsDirPath, directory))
	}
	db.edition = editions
}

func setEdition(editionPath string) *Edition {
	edition := &Edition{
		Bear:    new(Character),
		Usec:    new(Character),
		Storage: new(EditionStorage),
	}

	done := make(chan struct{})
	go func() {
		raw := tools.GetJSONRawMessage(filepath.Join(editionPath, "storage.json"))
		if err := json.UnmarshalNoEscape(raw, edition.Storage); err != nil {
			log.Fatalln(err)
		}
		done <- struct{}{}
	}()
	go func() {
		raw := tools.GetJSONRawMessage(filepath.Join(editionPath, "character_usec.json"))
		if err := json.UnmarshalNoEscape(raw, edition.Usec); err != nil {
			log.Fatalln(err)
		}
		done <- struct{}{}
	}()
	go func() {
		raw := tools.GetJSONRawMessage(filepath.Join(editionPath, "character_bear.json"))
		if err := json.UnmarshalNoEscape(raw, edition.Bear); err != nil {
			log.Fatalln(err)
		}
		done <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		<-done
	}
	return edition
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
