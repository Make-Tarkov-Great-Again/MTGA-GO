package data

import (
	"MT-GO/tools"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
)

var languages = make(map[string]string)

// Locales contains various locale information for all readable text in-game
var locales = &Locale{}

var localeMap = map[string]*LocaleData{
	"en":    &locales.EN,
	"fr":    &locales.FR,
	"ru":    &locales.RU,
	"es":    &locales.ES,
	"es-mx": &locales.ESMX,
	"ch":    &locales.CH,
	"cz":    &locales.CZ,
	"ge":    &locales.GE,
	"hu":    &locales.HU,
	"it":    &locales.IT,
	"jp":    &locales.JP,
	"kr":    &locales.KR,
	"pl":    &locales.PL,
	"po":    &locales.PO,
	"sk":    &locales.SK,
	"tu":    &locales.TU,
}

// #region Language getters

func GetLanguages() map[string]string {
	return languages
}

// #endregion

// #region Language setters
func setLanguages() {
	raw := tools.GetJSONRawMessage(filepath.Join(localesPath, "/languages.json"))
	if err := json.UnmarshalNoEscape(raw, &languages); err != nil {
		log.Fatalln(err)
	}
}

// #endregion

// #region Locale getters

func GetLocales() map[string]*LocaleData {
	return localeMap
}

func GetLocaleByName(input string) (*LocaleData, error) {
	name := strings.ToLower(input)
	if locale, ok := localeMap[name]; ok {
		return locale, nil
	}
	return nil, fmt.Errorf("Locale %s doesn't exist", name)
}

func GetLocalesMenuByName(name string) (*LocaleMenu, error) {
	locale, err := GetLocaleByName(name)
	if err != nil {
		return nil, err
	}

	if locale.Menu != nil {
		return locale.Menu, nil
	}

	return nil, fmt.Errorf("Locale %s menu doesn't exist", name)
}

func GetLocalesGlobalByName(name string) (map[string]any, error) {
	locale, err := GetLocaleByName(name)
	if err != nil {
		return nil, err
	}

	if locale.Locale != nil {
		return locale.Locale, nil
	}

	return nil, fmt.Errorf("Locale %s globals doesn't exist", name)
}

// #endregion

// #region Locale setters

func setLocales() {
	directories, err := tools.GetDirectoriesFrom(localesPath)
	if err != nil {
		log.Fatalln(err)
	}

	structure := make(map[string]*LocaleData)

	for dir := range directories {
		localeData := new(LocaleData)
		dirPath := filepath.Join(localesPath, dir)
		localeFiles, err := tools.GetFilesFrom(dirPath)
		if err != nil {
			log.Fatalln(err)
		}

		for file := range localeFiles {
			fileContent := tools.GetJSONRawMessage(filepath.Join(dirPath, file))

			switch file {
			case "locale.json":
				raw := make(map[string]any)
				if err := json.UnmarshalNoEscape(fileContent, &raw); err != nil {
					log.Fatalln(err)
				}
				localeData.Locale = raw
			case "menu.json":
				localeMenu := &LocaleMenu{}
				if err := json.UnmarshalNoEscape(fileContent, &localeMenu); err != nil {
					log.Fatalln(err)
				}

				localeData.Menu = localeMenu
			default:
				log.Println("huh")
			}
		}

		structure[dir] = localeData
	}

	bytes, err := json.MarshalNoEscape(structure)
	if err != nil {
		log.Fatalln(err)
	}

	if err = json.UnmarshalNoEscape(bytes, &locales); err != nil {
		log.Fatalln(err)
	}
}

// #endregion

// #region Locale struct

type Locales struct {
	Locales   Locale
	Languages map[string]string
}

type Locale struct {
	CH   LocaleData `json:"ch"`
	CZ   LocaleData `json:"cz"`
	EN   LocaleData `json:"en"`
	FR   LocaleData `json:"fr"`
	GE   LocaleData `json:"ge"`
	HU   LocaleData `json:"hu"`
	IT   LocaleData `json:"it"`
	JP   LocaleData `json:"jp"`
	KR   LocaleData `json:"kr"`
	PL   LocaleData `json:"pl"`
	PO   LocaleData `json:"po"`
	SK   LocaleData `json:"sk"`
	ES   LocaleData `json:"es"`
	ESMX LocaleData `json:"es-mx"`
	TU   LocaleData `json:"tu"`
	RU   LocaleData `json:"ru"`
}

type LocaleData struct {
	Locale map[string]any
	Menu   *LocaleMenu
}

type LocaleMenu struct {
	Menu map[string]string `json:"menu"`
}

// #endregion
