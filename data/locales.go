package data

import (
	"MT-GO/tools"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
)

// #region Locale getters

func GetLocaleByName(input string) (*Locale, error) {
	name := strings.ToLower(input)
	if locale, ok := db.locale[name]; ok {
		return locale, nil
	}
	return nil, fmt.Errorf("locale %s doesn't exist", name)
}

func GetLocaleMenuByName(name string) (*LocaleMenu, error) {
	locale, err := GetLocaleByName(name)
	if err != nil {
		return nil, err
	}

	if locale.Menu != nil {
		return locale.Menu, nil
	}

	return nil, fmt.Errorf("locale %s menu doesn't exist", name)
}

func GetLocaleGlobalByName(name string) (map[string]any, error) {
	locale, err := GetLocaleByName(name)
	if err != nil {
		return nil, err
	}

	if locale.Global != nil {
		return locale.Global, nil
	}

	return nil, fmt.Errorf("locale %s globals doesn't exist", name)
}

// #endregion

// #region Locale setters

func setLocales() {
	directories, err := tools.GetDirectoriesFrom(localesPath)
	if err != nil {
		log.Fatalln(err)
	}

	db.locale = make(map[string]*Locale)

	for dir := range directories {
		localeData := &Locale{
			Global: make(map[string]any),
			Menu:   new(LocaleMenu),
		}
		dirPath := filepath.Join(localesPath, dir)
		localeFiles, err := tools.GetFilesFrom(dirPath)
		if err != nil {
			log.Fatalln(err)
		}

		for file := range localeFiles {
			fileContent := tools.GetJSONRawMessage(filepath.Join(dirPath, file))

			switch file {
			case "locale.json":
				if err := json.UnmarshalNoEscape(fileContent, &localeData.Global); err != nil {
					log.Fatalln(err)
				}
			case "menu.json":
				if err := json.UnmarshalNoEscape(fileContent, &localeData.Menu); err != nil {
					log.Fatalln(err)
				}
			default:
				log.Println("huh")
			}
		}

		db.locale[dir] = localeData
	}
}

// #endregion

// #region Locale struct

type Locale struct {
	Global map[string]any
	Menu   *LocaleMenu
}

type LocaleMenu struct {
	Menu map[string]string `json:"menu"`
}

// #endregion
