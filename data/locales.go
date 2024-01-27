package data

import (
	"MT-GO/tools"
	"fmt"
	"github.com/alphadose/haxmap"
	"log"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
)

// #region Locale getters

func GetLocaleByName(input string) (*Locale, error) {
	name := strings.ToLower(input)
	if locale, ok := db.locale.Get(name); ok {
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

func GetLocaleGlobalByName(name string) (*haxmap.Map[string, any], error) {
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

	db.locale = haxmap.New[string, *Locale](uintptr(len(directories))) //make(map[string]*Locale)

	for dir := range directories {
		localeData := &Locale{
			Global: haxmap.New[string, any](), //make(map[string]any),
			Menu:   &LocaleMenu{Menu: haxmap.New[string, string]()},
		}

		dirPath := filepath.Join(localesPath, dir)
		fileContent := tools.GetJSONRawMessage(filepath.Join(dirPath, "locale.json"))
		if err := json.UnmarshalNoEscape(fileContent, &localeData.Global); err != nil {
			msg := tools.CheckParsingError(fileContent, err)
			log.Fatalln(msg)
		}

		fileContent = tools.GetJSONRawMessage(filepath.Join(dirPath, "menu.json"))
		if err := json.UnmarshalNoEscape(fileContent, &localeData.Menu); err != nil {
			msg := tools.CheckParsingError(fileContent, err)
			log.Fatalln(msg)
		}

		db.locale.Set(dir, localeData)
	}
}

// #endregion

// #region Locale struct

type Locale struct {
	Global *haxmap.Map[string, any] //map[string]any
	Menu   *LocaleMenu
}

type LocaleMenu struct {
	Menu *haxmap.Map[string, string] `json:"menu"` //map[string]string `json:"menu"`
}

// #endregion
