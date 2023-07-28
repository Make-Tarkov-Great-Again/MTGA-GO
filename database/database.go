// Package database contains all the database related code
package database

import (
	"MT-GO/database/structs"
	"MT-GO/tools"
	"encoding/json"
	"strings"
)

var db = structs.DatabaseStruct{}

// GetDatabase returns a pointer to the database
func GetDatabase() *structs.DatabaseStruct {
	return &db
}

// InitializeDatabase initializes the database
func InitializeDatabase() {
	db.Core = setCore()
	db.Items = setItems()
	db.Locales = setLocales()
	db.Languages = setLanguages()
	db.Handbook = setHandbook()
	db.Traders = setTraders()
	db.Locations = setLocations()
	db.Quests = setQuests()
	db.Hideout = setHideout()
	db.Weather = setWeather()
	db.Customization = setCustomization()
	db.Bot = setBots()
	db.Editions = setEditions()
	db.Flea = setFlea()
}

const (
	databaseLibPath       string = "assets/database"
	coreFilePath          string = databaseLibPath + "/core"
	botTemplateFilePath   string = coreFilePath + "/botTemplate.json"
	playerScavPath        string = coreFilePath + "/playerScav.json"
	clientSettingsPath    string = coreFilePath + "/client.settings.json"
	globalBotSettingsPath string = coreFilePath + "/__BotGlobalSettings.json"
	globalsFilePath       string = coreFilePath + "/globals.json"
	locationsFilePath     string = coreFilePath + "/locations.json"
	matchMetricsPath      string = coreFilePath + "/matchMetrics.json"
	serverConfigPath      string = coreFilePath + "/server.json"
	itemsPath             string = databaseLibPath + "/items.json"
	localesPath           string = databaseLibPath + "/locales"
	handbookPath          string = databaseLibPath + "/handbook.json"
	traderPath            string = databaseLibPath + "/traders/"
	questsPath            string = databaseLibPath + "/quests.json"
	hideoutPath           string = databaseLibPath + "/hideout/"
	weatherPath           string = databaseLibPath + "/weather.json"
	customizationPath     string = databaseLibPath + "/customization.json"
	botsPath              string = databaseLibPath + "/bot/"
	editionsDirPath       string = databaseLibPath + "/editions/"
)

func setFlea() *structs.Flea {
	return &structs.Flea{}
}

func setEditions() map[string]*structs.Edition {
	editions := make(map[string]*structs.Edition)

	directories, err := tools.GetDirectoriesFrom(editionsDirPath)
	if err != nil {
		panic(err)
	}

	for _, directory := range directories {
		edition := structs.Edition{}

		editionPath := editionsDirPath + directory + "/"
		files, err := tools.GetFilesFrom(editionPath)
		if err != nil {
			panic(err)
		}

		dynamic := make(map[string]interface{})
		for _, file := range files {
			raw := tools.GetJSONRawMessage(editionPath + file)
			removeJSON := strings.Replace(file, ".json", "", -1)
			name := strings.Replace(removeJSON, "character_", "", -1)

			dynamic[name] = raw
		}

		jsonData, err := json.Marshal(dynamic)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(jsonData, &edition)
		if err != nil {
			panic(err)
		}

		editions[directory] = &edition
	}
	return editions
}

func setCustomization() map[string]*structs.Customization {
	raw := tools.GetJSONRawMessage(customizationPath)

	customization := make(map[string]*structs.Customization)
	err := json.Unmarshal(raw, &customization)
	if err != nil {
		panic(err)
	}
	return customization
}

func setWeather() *structs.Weather {
	raw := tools.GetJSONRawMessage(weatherPath)

	weather := structs.Weather{}
	err := json.Unmarshal(raw, &weather)
	if err != nil {
		panic(err)
	}
	return &weather
}

func setHideout() *structs.Hideout {
	hideout := structs.Hideout{}

	dynamic := make(map[string]json.RawMessage)
	files := [5]string{"areas", "productions", "qte", "scavcase", "settings"}
	for _, file := range files {
		raw := tools.GetJSONRawMessage(hideoutPath + file + ".json")
		dynamic[file] = raw
	}

	jsonData, err := json.Marshal(dynamic)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonData, &hideout)
	if err != nil {
		panic(err)
	}
	return &hideout
}

func setHandbook() *structs.Handbook {
	handbook := structs.Handbook{}

	raw := tools.GetJSONRawMessage(handbookPath)
	err := json.Unmarshal(raw, &handbook)
	if err != nil {
		panic(err)
	}
	return &handbook
}

func setLocales() *structs.Locale {
	directories, err := tools.GetDirectoriesFrom(localesPath)
	if err != nil {
		panic(err)
	}

	locales := structs.Locale{}
	localeData := structs.LocaleData{}

	structure := make(map[string]structs.LocaleData)

	localeFiles := [2]string{"locale.json", "menu.json"}

	for _, dir := range directories {

		dirPath := localesPath + "/" + dir

		for _, file := range localeFiles {
			var fileContent []byte

			fileContent, err := tools.ReadFile(dirPath + "/" + file)
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

	return &locales
}

func setLanguages() map[string]string {
	languages := make(map[string]string)

	raw := tools.GetJSONRawMessage(localesPath + "/languages.json")
	err := json.Unmarshal(raw, &languages)
	if err != nil {
		panic(err)
	}
	return languages
}

func setItems() map[string]*structs.DatabaseItem {
	items := make(map[string]*structs.DatabaseItem)

	raw := tools.GetJSONRawMessage(itemsPath)
	err := json.Unmarshal(raw, &items)
	if err != nil {
		panic(err)
	}
	return items
}

func setCore() *structs.CoreStruct {
	core := structs.CoreStruct{}
	core.PlayerTemplate = setBotTemplate()
	core.PlayerScav = setPlayerScav()
	core.ClientSettings = setClientSettings()
	core.ServerConfig = setServerConfig()
	core.Globals = setGlobals()
	core.GlobalBotSettings = setGlobalBotSettings()
	core.MatchMetrics = setMatchMetrics()
	return &core
}

func setGlobalBotSettings() structs.GlobalBotSettings {
	raw := tools.GetJSONRawMessage(globalBotSettingsPath)

	globalBotSettings := structs.GlobalBotSettings{}
	err := json.Unmarshal(raw, &globalBotSettings)
	if err != nil {
		panic(err)
	}
	return globalBotSettings
}

func setPlayerScav() structs.PlayerTemplate {
	raw := tools.GetJSONRawMessage(playerScavPath)

	playerScav := structs.PlayerTemplate{}
	err := json.Unmarshal(raw, &playerScav)
	if err != nil {
		panic(err)
	}
	return playerScav
}

func setBotTemplate() structs.PlayerTemplate {
	raw := tools.GetJSONRawMessage(botTemplateFilePath)

	var botTemplate structs.PlayerTemplate
	err := json.Unmarshal(raw, &botTemplate)
	if err != nil {
		panic(err)
	}
	return botTemplate
}

func setClientSettings() structs.ClientSettings {
	raw := tools.GetJSONRawMessage(clientSettingsPath)

	var data structs.ClientSettings
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func setServerConfig() structs.ServerConfig {
	raw := tools.GetJSONRawMessage(serverConfigPath)

	var data structs.ServerConfig
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func setMatchMetrics() structs.MatchMetrics {
	raw := tools.GetJSONRawMessage(matchMetricsPath)

	var data structs.MatchMetrics
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func setGlobals() structs.Globals {
	raw := tools.GetJSONRawMessage(globalsFilePath)

	var global = structs.Globals{}
	err := json.Unmarshal(raw, &global)
	if err != nil {
		panic(err)
	}

	return global
}

func setLocations() *structs.Locations {
	raw := tools.GetJSONRawMessage(locationsFilePath)

	var locations = structs.Locations{}
	err := json.Unmarshal(raw, &locations)
	if err != nil {
		panic(err)
	}

	return &locations
}
