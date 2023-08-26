package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"fmt"
	"path/filepath"
)

// Locales contains various locale information for all readable text in-game
var Locales = &structs.Locale{}

var localeMap = map[string]*structs.LocaleData{
	"en":    &Locales.EN,
	"fr":    &Locales.FR,
	"ru":    &Locales.RU,
	"es":    &Locales.ES,
	"es-mx": &Locales.ESMX,
	"ch":    &Locales.CH,
	"cz":    &Locales.CZ,
	"ge":    &Locales.GE,
	"hu":    &Locales.HU,
	"it":    &Locales.IT,
	"jp":    &Locales.JP,
	"kr":    &Locales.KR,
	"pl":    &Locales.PL,
	"po":    &Locales.PO,
	"sk":    &Locales.SK,
	"tu":    &Locales.TU,
}

func GetLocales() *structs.Locale {
	return Locales
}

func GetLocaleByName(name string) *structs.LocaleData {
	if locale, ok := localeMap[name]; ok {
		return locale
	}
	fmt.Println("No such locale, returning EN")
	return &Locales.EN
}

func GetLocalesMenuByName(name string) *structs.LocaleMenu {
	if locale, ok := localeMap[name]; ok {
		return locale.Menu
	}
	fmt.Println("No such locale menu, returning EN")
	return Locales.EN.Menu
}

func GetLocalesLocaleByName(name string) map[string]interface{} {
	if locale, ok := localeMap[name]; ok {
		return locale.Locale
	}
	fmt.Println("No such locale ...locale, returning EN")
	return Locales.EN.Locale
}

func setLocales() {
	directories, err := tools.GetDirectoriesFrom(localesPath)
	if err != nil {
		panic(err)
	}

	structure := make(map[string]*structs.LocaleData)
	localeFiles := [2]string{"locale.json", "menu.json"}

	for _, dir := range directories {
		localeData := &structs.LocaleData{}
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
				localeMenu := &structs.LocaleMenu{}
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

	err = json.Unmarshal(bytes, &Locales)
	if err != nil {
		panic(err)
	}

}
