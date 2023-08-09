package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"path/filepath"
)

var locales = structs.Locale{}

func GetLocales() *structs.Locale {
	return &locales
}

func setLocales() {
	directories, err := tools.GetDirectoriesFrom(localesPath)
	if err != nil {
		panic(err)
	}

	localeData := structs.LocaleData{}
	structure := make(map[string]structs.LocaleData)

	localeFiles := [2]string{"locale.json", "menu.json"}

	for _, dir := range directories {

		dirPath := filepath.Join(localesPath, dir)

		for _, file := range localeFiles {
			var fileContent []byte

			fileContent, err := tools.ReadFile(filepath.Join(dirPath, file))
			if err != nil {
				panic(err)
			}

			raw := make(map[string]json.RawMessage)
			err = json.Unmarshal(fileContent, &raw)
			if err != nil {
				panic(err)
			}

			format := make(map[string]string)
			for key, val := range raw {
				format[key] = string(val)
			}

			if file == "locale.json" {
				localeData.Locale = format
			} else {
				localeData.Menu = format
			}
		}

		structure[dir] = localeData
	}

	jsonData, err := json.Marshal(structure)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &locales)
	if err != nil {
		panic(err)
	}
}
