package data

import (
	"MT-GO/tools"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
)

// #region Edition getters

func GetEditionByName(version string) (*Edition, error) {
	edition, ok := db.edition[version]
	if !ok {
		return edition, fmt.Errorf("Edition %s does not exist", version)
	}
	return edition, nil
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
		editions[strings.ToLower(directory)] = setEdition(filepath.Join(editionsDirPath, directory))
	}
	db.edition = editions
}

func setEdition(editionPath string) *Edition {
	edition := &Edition{
		Bear:    new(Character[map[string]PlayerTradersInfo]),
		Usec:    new(Character[map[string]PlayerTradersInfo]),
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
	Bear    *Character[map[string]PlayerTradersInfo] `json:"bear"`
	Usec    *Character[map[string]PlayerTradersInfo] `json:"usec"`
	Storage *EditionStorage                          `json:"storage"`
}

type EditionStorage struct {
	Bear []string `json:"bear"`
	Usec []string `json:"usec"`
}

// #endregion
