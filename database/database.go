package database

import (
	"MT-GO/database/structs"
	"MT-GO/tools"
	"encoding/json"
)

type DatabaseStruct struct {
	Core *structs.CoreStruct
	//Connections *ConnectionStruct
	Items map[string]*structs.DatabaseItem
	//Locales       *LocaleStruct
	//Templates     *TemplatesStruct
	//Traders       map[string]TraderStruct
	//Flea          *FleaStruct
	//Quests        map[string]map[string]interface{}
	//Hideout       *HideoutStruct
	//Locations     *LocationsStruct
	//Weather       map[string]interface{}
	//Customization map[string]map[string]interface{}
	//Editions      map[string]interface{}
	//Bot           *BotStruct
	//Profiles      map[string]ProfileStruct
	//bundles  []map[string]interface{}
}

var database = &DatabaseStruct{}

func GetDatabase() *DatabaseStruct {
	return database
}

func InitializeDatabase() error {
	database.Core = setCore()
	database.Items = setItems()
	return nil
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
)

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
	core := &structs.CoreStruct{}
	core.PlayerTemplate = processBotTemplate()
	core.ClientSettings = processClientSettings()
	core.ServerConfig = processServerConfig()
	core.Globals = processGlobals()
	core.Locations = processLocations()
	core.MatchMetrics = processMatchMetrics()
	return core
}

func processBotTemplate() structs.PlayerTemplate {
	raw := tools.GetJSONRawMessage(BOT_TEMPLATE_FILE_PATH)

	var botTemplate structs.PlayerTemplate
	err := json.Unmarshal(raw, &botTemplate)
	if err != nil {
		panic(err)
	}
	return botTemplate
}

func processClientSettings() structs.ClientSettings {
	raw := tools.GetJSONRawMessage(CLIENT_SETTINGS_PATH)

	var data structs.ClientSettings
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func processServerConfig() structs.ServerConfig {
	raw := tools.GetJSONRawMessage(SERVER_CONFIG_PATH)

	var data structs.ServerConfig
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func processMatchMetrics() structs.MatchMetrics {
	raw := tools.GetJSONRawMessage(MATCH_METRICS_PATH)

	var data structs.MatchMetrics
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func processGlobals() structs.Globals {
	raw := tools.GetJSONRawMessage(GLOBALS_FILE_PATH)

	var global = structs.Globals{}
	err := json.Unmarshal(raw, &global)
	if err != nil {
		panic(err)
	}

	return global
}

func processLocations() structs.Locations {
	raw := tools.GetJSONRawMessage(LOCATIONS_FILE_PATH)

	var locations = structs.Locations{}
	err := json.Unmarshal(raw, &locations)
	if err != nil {
		panic(err)
	}

	return locations
}
