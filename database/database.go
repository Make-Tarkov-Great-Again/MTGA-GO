// Package database contains all the database related code
package database

import (
	"MT-GO/database/structs"
	"MT-GO/tools"
	"encoding/json"
	"path/filepath"
	"strings"
)

var db = structs.Database{}

// GetDatabase returns a pointer to the database
func GetDatabase() *structs.Database {
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
	db.Profiles = setProfiles()
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
	profilesPath          string = "user/profiles/"
)

func setProfiles() map[string]*structs.Profile {

	users, err := tools.GetDirectoriesFrom(profilesPath)
	if err != nil {
		panic(err)
	}
	profiles := make(map[string]*structs.Profile)
	if len(users) == 0 {
		return profiles
	}
	for _, user := range users {
		profile := &structs.Profile{}
		userPath := filepath.Join(profilesPath, user)
		files, err := tools.GetFilesFrom(userPath)
		if err != nil {
			panic(err)
		}

		dynamic := make(map[string]json.RawMessage)
		for _, file := range files {
			name := strings.TrimSuffix(file, ".json")
			data := tools.GetJSONRawMessage(filepath.Join(userPath, file))
			dynamic[name] = data
		}

		jsonData, err := json.Marshal(dynamic) //gos syntax is fucking pog.
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(jsonData, profile)
		if err != nil {
			panic(err)
		}
		profiles[user] = profile
	}
	return profiles
}

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

		editionPath := filepath.Join(editionsDirPath, directory)
		files, err := tools.GetFilesFrom(editionPath)
		if err != nil {
			panic(err)
		}

		dynamic := make(map[string]interface{})
		for _, file := range files {
			raw := tools.GetJSONRawMessage(filepath.Join(editionPath, file))
			removeJSON := strings.TrimSuffix(file, ".json")
			name := strings.TrimPrefix(removeJSON, "character_")

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
	files, err := tools.GetFilesFrom(hideoutPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		name := strings.TrimSuffix(file, ".json")
		raw := tools.GetJSONRawMessage(filepath.Join(hideoutPath, file))
		dynamic[name] = raw
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

	return &locales
}

func setLanguages() map[string]string {
	languages := make(map[string]string)

	raw := tools.GetJSONRawMessage(filepath.Join(localesPath, "/languages.json"))
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

func setCore() *structs.Core {
	core := structs.Core{}
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
