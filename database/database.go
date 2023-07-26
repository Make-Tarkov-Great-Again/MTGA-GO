package database

import (
	"MT-GO/database/structs"
	"MT-GO/tools"
	"encoding/json"
)

type DatabaseStruct struct {
	Core *structs.CoreStruct
	//Connections *ConnectionStruct
	Items     map[string]*structs.DatabaseItem
	Locales   *structs.Locale
	Languages map[string]string
	Handbook  *structs.Handbook
	Traders   map[string]*structs.Trader
	//Flea          *FleaStruct
	Quests  map[string]*structs.Quest
	Hideout *structs.Hideout

	Locations     *structs.Locations
	Weather       *structs.Weather
	Customization map[string]*structs.Customization
	//Editions      map[string]interface{}
	//Bot           *BotStruct
	//Profiles      map[string]ProfileStruct
	//bundles  []map[string]interface{}
}

var db = DatabaseStruct{}

func GetDatabase() *DatabaseStruct {
	return &db
}

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
}

const (
	DATABASE_LIB_PATH      string = "database/lib"
	CORE_FILE_PATH         string = DATABASE_LIB_PATH + "/core"
	BOT_TEMPLATE_FILE_PATH string = CORE_FILE_PATH + "/botTemplate.json"
	CLIENT_SETTINGS_PATH   string = CORE_FILE_PATH + "/client.settings.json"
	GLOBALS_FILE_PATH      string = CORE_FILE_PATH + "/globals.json"
	LOCATIONS_FILE_PATH    string = CORE_FILE_PATH + "/locations.json"
	MATCH_METRICS_PATH     string = CORE_FILE_PATH + "/matchMetrics.json"
	SERVER_CONFIG_PATH     string = CORE_FILE_PATH + "/server.json"
	ITEMS_PATH             string = DATABASE_LIB_PATH + "/items.json"
	LOCALES_PATH           string = DATABASE_LIB_PATH + "/locales"
	HANDBOOK_PATH          string = DATABASE_LIB_PATH + "/handbook.json"
	TRADER_PATH            string = DATABASE_LIB_PATH + "/traders/"
	QUESTS_PATH            string = DATABASE_LIB_PATH + "/quests.json"
	HIDEOUT_PATH           string = DATABASE_LIB_PATH + "/hideout/"
	WEATHER_PATH           string = DATABASE_LIB_PATH + "/weather.json"
	CUSTOMIZATION_PATH     string = DATABASE_LIB_PATH + "/customization.json"
)

func setCustomization() map[string]*structs.Customization {
	raw := tools.GetJSONRawMessage(CUSTOMIZATION_PATH)

	customization := make(map[string]*structs.Customization)
	err := json.Unmarshal(raw, &customization)
	if err != nil {
		panic(err)
	}
	return customization
}

func setWeather() *structs.Weather {
	raw := tools.GetJSONRawMessage(WEATHER_PATH)

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
		raw := tools.GetJSONRawMessage(HIDEOUT_PATH + file + ".json")
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

	raw := tools.GetJSONRawMessage(HANDBOOK_PATH)
	err := json.Unmarshal(raw, &handbook)
	if err != nil {
		panic(err)
	}
	return &handbook
}

func setLocales() *structs.Locale {
	directories, err := tools.GetDirectoriesFrom(LOCALES_PATH)
	if err != nil {
		panic(err)
	}

	locales := structs.Locale{}
	localeData := structs.LocaleData{}

	structure := make(map[string]structs.LocaleData)

	localeFiles := [2]string{"locale.json", "menu.json"}

	for _, dir := range directories {

		dirPath := LOCALES_PATH + "/" + dir

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

	raw := tools.GetJSONRawMessage(LOCALES_PATH + "/languages.json")
	err := json.Unmarshal(raw, &languages)
	if err != nil {
		panic(err)
	}
	return languages
}

func setItems() map[string]*structs.DatabaseItem {
	items := make(map[string]*structs.DatabaseItem)

	raw := tools.GetJSONRawMessage(ITEMS_PATH)
	err := json.Unmarshal(raw, &items)
	if err != nil {
		panic(err)
	}
	return items
}

func setCore() *structs.CoreStruct {
	core := structs.CoreStruct{}
	core.PlayerTemplate = setBotTemplate()
	core.ClientSettings = setClientSettings()
	core.ServerConfig = setServerConfig()
	core.Globals = setGlobals()
	core.MatchMetrics = setMatchMetrics()
	return &core
}

func setBotTemplate() structs.PlayerTemplate {
	raw := tools.GetJSONRawMessage(BOT_TEMPLATE_FILE_PATH)

	var botTemplate structs.PlayerTemplate
	err := json.Unmarshal(raw, &botTemplate)
	if err != nil {
		panic(err)
	}
	return botTemplate
}

func setClientSettings() structs.ClientSettings {
	raw := tools.GetJSONRawMessage(CLIENT_SETTINGS_PATH)

	var data structs.ClientSettings
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func setServerConfig() structs.ServerConfig {
	raw := tools.GetJSONRawMessage(SERVER_CONFIG_PATH)

	var data structs.ServerConfig
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func setMatchMetrics() structs.MatchMetrics {
	raw := tools.GetJSONRawMessage(MATCH_METRICS_PATH)

	var data structs.MatchMetrics
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func setGlobals() structs.Globals {
	raw := tools.GetJSONRawMessage(GLOBALS_FILE_PATH)

	var global = structs.Globals{}
	err := json.Unmarshal(raw, &global)
	if err != nil {
		panic(err)
	}

	return global
}

func setLocations() *structs.Locations {
	raw := tools.GetJSONRawMessage(LOCATIONS_FILE_PATH)

	var locations = structs.Locations{}
	err := json.Unmarshal(raw, &locations)
	if err != nil {
		panic(err)
	}

	return &locations
}
