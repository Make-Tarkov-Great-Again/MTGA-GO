package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"fmt"
	"net"
)

var core = structs.Core{}

func GetCore() *structs.Core {
	return &core
}

func GetGlobals() *structs.Globals {
	return core.Globals
}

func GetClientSettings() *structs.ClientSettings {
	return core.ClientSettings
}

func GetMatchMetrics() *structs.MatchMetrics {
	return core.MatchMetrics
}

func GetServerConfig() *structs.ServerConfig {
	return core.ServerConfig
}

func GetMainAddress() string {
	return coreServerData.MainAddress
}

func GetTradingAddress() string {
	return coreServerData.TradingAddress
}

func GetMessageAddress() string {
	return coreServerData.MessageAddress
}

func GetRagFairAddress() string {
	return coreServerData.RagFairAddress
}

func GetLobbyAddress() string {
	return coreServerData.LobbyAddress
}
func GetWebsocketURL() string {
	return coreServerData.WSSAddress
}

func GetGlobalBotSettings() *map[string]interface{} {
	return core.GlobalBotSettings
}

func GetPlayerScav() *structs.PlayerTemplate {
	return core.PlayerScav
}

func GetBotTemplate() *structs.PlayerTemplate {
	return core.PlayerTemplate
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

func setGlobalBotSettings() *map[string]interface{} {
	raw := tools.GetJSONRawMessage(globalBotSettingsPath)

	globalBotSettings := map[string]interface{}{}
	err := json.Unmarshal(raw, &globalBotSettings)
	if err != nil {
		panic(err)
	}
	return &globalBotSettings
}

func setPlayerScav() *structs.PlayerTemplate {
	raw := tools.GetJSONRawMessage(playerScavPath)

	playerScav := structs.PlayerTemplate{}
	err := json.Unmarshal(raw, &playerScav)
	if err != nil {
		panic(err)
	}
	return &playerScav
}

func setBotTemplate() *structs.PlayerTemplate {
	raw := tools.GetJSONRawMessage(botTemplateFilePath)

	var botTemplate structs.PlayerTemplate
	err := json.Unmarshal(raw, &botTemplate)
	if err != nil {
		panic(err)
	}
	return &botTemplate
}

func setClientSettings() *structs.ClientSettings {
	raw := tools.GetJSONRawMessage(clientSettingsPath)

	var data structs.ClientSettings
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

type serverData struct {
	HTTPSTemplate  string
	WSSTemplate    string
	WSSAddress     string
	IPandPort      string
	MainAddress    string
	MessageAddress string
	TradingAddress string
	RagFairAddress string
	LobbyAddress   string
}

var coreServerData = &serverData{}

func SetWebSocketAddress(sessionID string) {
	coreServerData.WSSAddress = fmt.Sprintf(coreServerData.WSSTemplate, coreServerData.MainAddress, sessionID)
}

func GetWebSocketAddress() string {
	return coreServerData.WSSAddress
}

func GetIPandPort() string {
	return coreServerData.IPandPort
}

func setServerConfig() *structs.ServerConfig {
	raw := tools.GetJSONRawMessage(serverConfigPath)

	var data structs.ServerConfig
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}

	coreServerData.HTTPSTemplate = "https://%s"
	coreServerData.IPandPort = net.JoinHostPort(data.IP, data.Ports.Main)
	coreServerData.WSSTemplate = "wss://%s/socket/%s"

	coreServerData.MainAddress = fmt.Sprintf(coreServerData.HTTPSTemplate, coreServerData.IPandPort)
	coreServerData.MessageAddress = fmt.Sprintf(coreServerData.HTTPSTemplate, net.JoinHostPort(data.IP, data.Ports.Messaging))
	coreServerData.TradingAddress = fmt.Sprintf(coreServerData.HTTPSTemplate, net.JoinHostPort(data.IP, data.Ports.Trading))
	coreServerData.RagFairAddress = fmt.Sprintf(coreServerData.HTTPSTemplate, net.JoinHostPort(data.IP, data.Ports.Flea))
	coreServerData.LobbyAddress = fmt.Sprintf(coreServerData.HTTPSTemplate, net.JoinHostPort(data.IP, data.Ports.Lobby))

	return &data
}

func setMatchMetrics() *structs.MatchMetrics {
	raw := tools.GetJSONRawMessage(matchMetricsPath)

	var data structs.MatchMetrics
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

func setGlobals() *structs.Globals {
	raw := tools.GetJSONRawMessage(globalsFilePath)

	var global = structs.Globals{}
	err := json.Unmarshal(raw, &global)
	if err != nil {
		panic(err)
	}

	return &global
}
