package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"fmt"
	"net"

	"github.com/goccy/go-json"
)

var core = structs.Core{}

func GetCore() *structs.Core {
	return &core
}

func GetGlobals() *structs.Globals {
	return core.Globals
}

func GetMainSettings() *structs.MainSettings {
	return core.MainSettings
}

func GetMatchMetrics() *structs.MatchMetrics {
	return core.MatchMetrics
}

func GetServerConfig() *structs.ServerConfig {
	return core.ServerConfig
}

func GetGlobalBotSettings() *map[string]interface{} {
	return core.GlobalBotSettings
}

func GetPlayerScav() *structs.PlayerScavTemplate {
	return core.PlayerScav
}

func GetBotTemplate() *structs.PlayerTemplate {
	return core.PlayerTemplate
}

func setCore() {
	core.PlayerTemplate = setBotTemplate()
	core.PlayerScav = setPlayerScav()
	core.MainSettings = setMainSettings()
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

func setPlayerScav() *structs.PlayerScavTemplate {
	raw := tools.GetJSONRawMessage(playerScavPath)

	var playerScav structs.PlayerScavTemplate
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

func setMainSettings() *structs.MainSettings {
	raw := tools.GetJSONRawMessage(MainSettingsPath)

	var data structs.MainSettings
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

type serverData struct {
	HTTPSTemplate string
	WSSTemplate   string
	WSSAddress    string

	MainIPandPort string
	MainAddress   string

	MessagingIPandPort string
	MessageAddress     string

	TradingIPandPort string
	TradingAddress   string

	RagFairIPandPort string
	RagFairAddress   string

	LobbyIPandPort string
	LobbyAddress   string
}

var coreServerData = &serverData{}

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
func GetWebSocketAddress() string {
	return coreServerData.WSSAddress
}

func GetMainIPandPort() string {
	return coreServerData.MainIPandPort
}

func GetTradingIPandPort() string {
	return coreServerData.TradingIPandPort
}

func GetMessagingIPandPort() string {
	return coreServerData.MessagingIPandPort
}

func GetLobbyIPandPort() string {
	return coreServerData.LobbyIPandPort
}

func GetRagFairIPandPort() string {
	return coreServerData.RagFairIPandPort
}

func setServerConfig() *structs.ServerConfig {
	raw := tools.GetJSONRawMessage(serverConfigPath)

	var data structs.ServerConfig
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}

	coreServerData.HTTPSTemplate = "https://%s"
	coreServerData.WSSTemplate = "wss://%s"

	coreServerData.MainIPandPort = net.JoinHostPort(data.IP, data.Ports.Main)
	coreServerData.MessagingIPandPort = net.JoinHostPort(data.IP, data.Ports.Messaging)
	coreServerData.TradingIPandPort = net.JoinHostPort(data.IP, data.Ports.Trading)
	coreServerData.RagFairIPandPort = net.JoinHostPort(data.IP, data.Ports.Flea)
	coreServerData.LobbyIPandPort = net.JoinHostPort(data.IP, data.Ports.Lobby)

	coreServerData.WSSAddress = fmt.Sprintf(coreServerData.WSSTemplate, coreServerData.MainIPandPort)

	coreServerData.MainAddress = fmt.Sprintf(coreServerData.HTTPSTemplate, coreServerData.MainIPandPort)

	coreServerData.MessageAddress = fmt.Sprintf(coreServerData.HTTPSTemplate, coreServerData.MessagingIPandPort)

	coreServerData.TradingAddress = fmt.Sprintf(coreServerData.HTTPSTemplate, coreServerData.TradingIPandPort)

	coreServerData.RagFairAddress = fmt.Sprintf(coreServerData.HTTPSTemplate, coreServerData.RagFairIPandPort)

	coreServerData.LobbyAddress = fmt.Sprintf("wss://%s/sws", coreServerData.LobbyIPandPort)

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
