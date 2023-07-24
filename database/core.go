package database

import (
	"MT-GO/database/structs"
	"MT-GO/tools"
	"encoding/json"
)

const (
	CORE_FILE_PATH         string = DATABASE_LIB_PATH + "/core"
	BOT_TEMPLATE_FILE_PATH string = CORE_FILE_PATH + "/botTemplate.json"
	CLIENT_SETTINGS_PATH   string = CORE_FILE_PATH + "/client.settings.json"
	GLOBALS_FILE_PATH      string = CORE_FILE_PATH + "/globals.json"
	LOCATIONS_FILE_PATH    string = CORE_FILE_PATH + "/locations.json"
	MATCH_METRICS_PATH     string = CORE_FILE_PATH + "/matchMetrics.json"
	SERVER_CONFIG_PATH     string = CORE_FILE_PATH + "/server.json"
)

func LoadCore() *structs.CoreStruct {
	core := &structs.CoreStruct{}
	core.BotTemplate = processBotTemplate()
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
