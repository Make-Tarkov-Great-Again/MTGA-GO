package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
)

var core = structs.Core{}

func GetCore() *structs.Core {
	return &core
}

func setCore() {
	core.PlayerTemplate = setBotTemplate()
	core.PlayerScav = setPlayerScav()
	core.ClientSettings = setClientSettings()
	core.ServerConfig = setServerConfig()
	core.Globals = setGlobals()
	core.GlobalBotSettings = setGlobalBotSettings()
	core.MatchMetrics = setMatchMetrics()
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
