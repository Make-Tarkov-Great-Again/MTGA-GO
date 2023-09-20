package database

import (
	"MT-GO/tools"
	"fmt"
	"path/filepath"

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
	err := json.Unmarshal(raw, &languages)
	if err != nil {
		panic(err)
	}
}

// #endregion

// #region Locale getters

func GetLocales() *Locale {
	return locales
}

func GetLocaleByName(name string) *LocaleData {
	if locale, ok := localeMap[name]; ok {
		return locale
	}
	fmt.Println("Locale doesnt exist, returning EN")
	return &locales.EN
}

func GetLocalesMenuByName(name string) *LocaleMenu {
	if locale, ok := localeMap[name]; ok {
		return locale.Menu
	}
	fmt.Println("Locale menu doesnt exist , returning EN")
	return locales.EN.Menu
}

func GetLocalesLocaleByName(name string) map[string]interface{} {
	if locale, ok := localeMap[name]; ok {
		return locale.Locale
	}
	fmt.Println("Locale doesnt exist, returning EN")
	return locales.EN.Locale
}

// #endregion

// #region Locale setters

func setLocales() {
	directories, err := tools.GetDirectoriesFrom(localesPath)
	if err != nil {
		panic(err)
	}

	structure := make(map[string]*LocaleData)
	localeFiles := [2]string{"locale.json", "menu.json"}

	for _, dir := range directories {
		localeData := &LocaleData{}
		dirPath := filepath.Join(localesPath, dir)

		for _, file := range localeFiles {

			fileContent := tools.GetJSONRawMessage(filepath.Join(dirPath, file))
			if err != nil {
				panic(err)
			}

			raw := make(map[string]interface{})

			if file == "locale.json" {
				err = json.Unmarshal(fileContent, &raw)
				if err != nil {
					panic(err)
				}
				localeData.Locale = raw
			} else {
				localeMenu := &LocaleMenu{}
				err = json.Unmarshal(fileContent, &localeMenu)
				if err != nil {
					panic(err)
				}

				localeData.Menu = localeMenu
			}
		}

		structure[dir] = localeData
	}

	bytes, err := json.Marshal(structure)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &locales)
	if err != nil {
		panic(err)
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
	Locale map[string]interface{}
	Menu   *LocaleMenu
}

type LocaleMenu struct {
	Menu map[string]string `json:"menu"`
}

// #endregion
