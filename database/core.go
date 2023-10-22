package database

import (
	"fmt"
	"log"
	"net"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var core = Core{}
var coreServerData = &serverData{}

// #region Core getters

func GetCore() *Core {
	return &core
}

func GetGlobals() *Globals {
	return core.Globals
}

func GetMainSettings() *MainSettings {
	return core.MainSettings
}

func GetAirdropParameters() *AirdropParameters {
	return core.AirdropParameters
}

func GetMatchMetrics() *MatchMetrics {
	return core.MatchMetrics
}

func GetServerConfig() *ServerConfig {
	return core.ServerConfig
}

func GetGlobalBotSettings() *map[string]interface{} {
	return core.GlobalBotSettings
}

func GetPlayerScav() *Scav {
	return core.Scav
}

func GetBotTemplate() *Character {
	return core.Character
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
func GetWebSocketAddress() string {
	if coreServerData.HTTPS != nil {
		return coreServerData.HTTPS.WSSAddress
	}
	return coreServerData.HTTP.WSAddress

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

// #endregion

// #region Core setters

func setCore() {
	core.Character = setBotTemplate()
	core.Scav = setPlayerScav()
	core.MainSettings = setMainSettings()
	core.ServerConfig = setServerConfig()
	core.Globals = setGlobals()
	core.GlobalBotSettings = setGlobalBotSettings()
	core.MatchMetrics = setMatchMetrics()
	core.AirdropParameters = setGetAirdropSettings()
}

func setGetAirdropSettings() *AirdropParameters {
	raw := tools.GetJSONRawMessage(airdropFilePath)

	airDropParameters := new(AirdropParameters)
	err := json.Unmarshal(raw, &airDropParameters)
	if err != nil {
		log.Fatalln(err)
	}
	return airDropParameters
}

func setGlobalBotSettings() *map[string]interface{} {
	raw := tools.GetJSONRawMessage(globalBotSettingsPath)

	globalBotSettings := map[string]interface{}{}
	err := json.Unmarshal(raw, &globalBotSettings)
	if err != nil {
		log.Fatalln(err)
	}
	return &globalBotSettings
}

func setPlayerScav() *Scav {
	raw := tools.GetJSONRawMessage(playerScavPath)

	var playerScav Scav
	err := json.Unmarshal(raw, &playerScav)
	if err != nil {
		log.Fatalln(err)
	}
	return &playerScav
}

func setBotTemplate() *Character {
	raw := tools.GetJSONRawMessage(botTemplateFilePath)

	var botTemplate Character
	err := json.Unmarshal(raw, &botTemplate)
	if err != nil {
		log.Fatalln(err)
	}
	return &botTemplate
}

func setMainSettings() *MainSettings {
	raw := tools.GetJSONRawMessage(MainSettingsPath)

	var data MainSettings
	err := json.Unmarshal(raw, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return &data
}

const (
	WSSTemplate   = "wss://%s"
	HTTPSTemplate = "https://%s"
	WSTemplate    = "ws://%s"
	HTTPTemplate  = "http://%s"
)

func setServerConfig() *ServerConfig {
	raw := tools.GetJSONRawMessage(serverConfigPath)

	data := new(ServerConfig)
	err := json.Unmarshal(raw, &data)
	if err != nil {
		log.Fatalln(err)
	}

	coreServerData.MainIPandPort = net.JoinHostPort(data.IP, data.Ports.Main)
	coreServerData.MessagingIPandPort = net.JoinHostPort(data.IP, data.Ports.Messaging)
	coreServerData.TradingIPandPort = net.JoinHostPort(data.IP, data.Ports.Trading)
	coreServerData.RagFairIPandPort = net.JoinHostPort(data.IP, data.Ports.Flea)
	coreServerData.LobbyIPandPort = net.JoinHostPort(data.IP, data.Ports.Lobby)

	if data.Secure {
		coreServerData.HTTPS = new(serverDataHTTPS)

		coreServerData.HTTPS.HTTPSAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.MainIPandPort)
		coreServerData.HTTPS.WSSAddress = fmt.Sprintf(WSSTemplate, coreServerData.LobbyIPandPort)

		coreServerData.MainAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.MainIPandPort)
		coreServerData.MessageAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.MessagingIPandPort)
		coreServerData.TradingAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.TradingIPandPort)
		coreServerData.RagFairAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.RagFairIPandPort)
		coreServerData.LobbyAddress = fmt.Sprintf("wss://%s/sws", coreServerData.LobbyIPandPort)
	} else {
		coreServerData.HTTP = new(serverDataHTTP)

		coreServerData.HTTP.HTTPAddress = fmt.Sprintf(HTTPTemplate, coreServerData.LobbyIPandPort)
		coreServerData.HTTP.WSAddress = fmt.Sprintf(WSTemplate, coreServerData.LobbyIPandPort)

		coreServerData.MainAddress = fmt.Sprintf(HTTPTemplate, coreServerData.MainIPandPort)
		coreServerData.MessageAddress = fmt.Sprintf(HTTPTemplate, coreServerData.MessagingIPandPort)
		coreServerData.TradingAddress = fmt.Sprintf(HTTPTemplate, coreServerData.TradingIPandPort)
		coreServerData.RagFairAddress = fmt.Sprintf(HTTPTemplate, coreServerData.RagFairIPandPort)
		coreServerData.LobbyAddress = fmt.Sprintf("ws://%s/sws", coreServerData.LobbyIPandPort)
	}

	return data
}

func setMatchMetrics() *MatchMetrics {
	raw := tools.GetJSONRawMessage(matchMetricsPath)

	var data MatchMetrics
	err := json.Unmarshal(raw, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return &data
}

func setGlobals() *Globals {
	raw := tools.GetJSONRawMessage(globalsFilePath)

	var global = Globals{}
	err := json.Unmarshal(raw, &global)
	if err != nil {
		log.Fatalln(err)
	}

	return &global
}

// #endregion

// #region Core structs

type serverData struct {
	HTTPS *serverDataHTTPS
	HTTP  *serverDataHTTP

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

type serverDataHTTPS struct {
	HTTPSAddress string
	WSSAddress   string
}

type serverDataHTTP struct {
	HTTPAddress string
	WSAddress   string
}

type Core struct {
	Character         *Character
	Scav              *Scav
	MainSettings      *MainSettings
	ServerConfig      *ServerConfig
	Globals           *Globals
	GlobalBotSettings *map[string]interface{}
	//gameplay        map[string]interface{}
	//blacklist       []interface{}
	MatchMetrics      *MatchMetrics
	AirdropParameters *AirdropParameters
}

type AirdropParameters struct {
	AirdropChancePercent       airDropChance `json:"airdropChancePercent"`
	AirdropMinStartTimeSeconds int16         `json:"airdropMinStartTimeSeconds"`
	AirdropMaxStartTimeSeconds int16         `json:"airdropMaxStartTimeSeconds"`
	PlaneMinFlyHeight          int16         `json:"planeMinFlyHeight"`
	PlaneMaxFlyHeight          int16         `json:"planeMaxFlyHeight"`
	PlaneVolume                float32       `json:"planeVolume"`
	PlaneSpeed                 int16         `json:"planeSpeed"`
	CrateFallSpeed             int16         `json:"crateFallSpeed"`
}

type airDropChance struct {
	Bigmap        int8 `json:"bigmap"`
	Woods         int8 `json:"woods"`
	Lighthouse    int8 `json:"lighthouse"`
	Shoreline     int8 `json:"shoreline"`
	Interchange   int8 `json:"interchange"`
	Reserve       int8 `json:"reserve"`
	TarkovStreets int8 `json:"tarkovStreets"`
}

type Scav struct {
	ID                string                 `json:"_id"`
	AID               int                    `json:"aid"`
	Savage            *string                `json:"savage"`
	Info              PlayerInfo             `json:"Info"`
	Customization     PlayerCustomization    `json:"Customization"`
	Health            HealthInfo             `json:"Health"`
	Inventory         Inventory              `json:"Inventory"`
	Skills            PlayerSkills           `json:"Skills"`
	Stats             PlayerStats            `json:"Stats"`
	Encyclopedia      map[string]bool        `json:"Encyclopedia"`
	ConditionCounters ConditionCounters      `json:"ConditionCounters"`
	BackendCounters   map[string]interface{} `json:"BackendCounters"`
	InsuredItems      []interface{}          `json:"InsuredItems"`
	Hideout           interface{}            `json:"Hideout"`
	Bonuses           []interface{}          `json:"Bonuses"`
	Notes             struct {
		Notes [][]interface{} `json:"Notes"`
	} `json:"Notes"`
	Quests       []interface{}     `json:"Quests"`
	RagfairInfo  PlayerRagfairInfo `json:"RagfairInfo"`
	WishList     []interface{}     `json:"WishList"`
	TradersInfo  []interface{}     `json:"TradersInfo"`
	UnlockedInfo struct {
		UnlockedProductionRecipe []interface{} `json:"unlockedProductionRecipe"`
	} `json:"UnlockedInfo"`
}

type Globals struct {
	Config               map[string]interface{} `json:"config"`
	BotPresets           [18]interface{}        `json:"bot_presets"`
	BotWeaponScatterings [4]interface{}         `json:"BotWeaponScatterings"`
	ItemPresets          map[string]interface{} `json:"ItemPresets"`
}

type MainSettings struct {
	Config struct {
		MemoryManagementSettings       MemoryManagementSettings `json:"MemoryManagementSettings"`
		ReleaseProfiler                ReleaseProfiler          `json:"ReleaseProfiler"`
		FramerateLimit                 FramerateLimit           `json:"FramerateLimit"`
		ClientSendRateLimit            int                      `json:"ClientSendRateLimit"`
		TurnOffLogging                 bool                     `json:"TurnOffLogging"`
		NVidiaHighlights               bool                     `json:"NVidiaHighlights"`
		WebDiagnosticsEnabled          bool                     `json:"WebDiagnosticsEnabled"`
		KeepAliveInterval              int                      `json:"KeepAliveInterval"`
		GroupStatusInterval            int                      `json:"GroupStatusInterval"`
		GroupStatusButtonInterval      int                      `json:"GroupStatusButtonInterval"`
		PingServersInterval            int                      `json:"PingServersInterval"`
		PingServerResultSendInterval   int                      `json:"PingServerResultSendInterval"`
		WeaponOverlapDistanceCulling   int                      `json:"WeaponOverlapDistanceCulling"`
		FirstCycleDelaySeconds         int                      `json:"FirstCycleDelaySeconds"`
		SecondCycleDelaySeconds        int                      `json:"SecondCycleDelaySeconds"`
		NextCycleDelaySeconds          int                      `json:"NextCycleDelaySeconds"`
		AdditionalRandomDelaySeconds   int                      `json:"AdditionalRandomDelaySeconds"`
		Mark502And504AsNonImportant    bool                     `json:"Mark502and504AsNonImportant"`
		DefaultRetriesCount            int                      `json:"DefaultRetriesCount"`
		CriticalRetriesCount           int                      `json:"CriticalRetriesCount"`
		AFKTimeoutSeconds              int                      `json:"AFKTimeoutSeconds"`
		RequestsMadeThroughLobby       []string                 `json:"RequestsMadeThroughLobby"`
		LobbyKeepAliveInterval         int                      `json:"LobbyKeepAliveInterval"`
		RequestConfirmationTimeouts    []float64                `json:"RequestConfirmationTimeouts"`
		ShouldEstablishLobbyConnection bool                     `json:"ShouldEstablishLobbyConnection"`
	} `json:"config"`
	NetworkStateView struct {
		LossThreshold int `json:"LossThreshold"`
		RttThreshold  int `json:"RttThreshold"`
	} `json:"NetworkStateView"`
}

type MemoryManagementSettings struct {
	HeapPreAllocationEnabled               bool `json:"HeapPreAllocationEnabled"`
	HeapPreAllocationMB                    int  `json:"HeapPreAllocationMB"`
	OverrideRAMCleanerSettings             bool `json:"OverrideRamCleanerSettings"`
	RAMCleanerEnabled                      bool `json:"RamCleanerEnabled"`
	GigabytesRequiredToDisableGCDuringRaid int  `json:"GigabytesRequiredToDisableGCDuringRaid"`
	AggressiveGC                           bool `json:"AggressiveGC"`
}
type ReleaseProfiler struct {
	Enabled            bool `json:"Enabled"`
	RecordTriggerValue int  `json:"RecordTriggerValue"`
	MaxRecords         int  `json:"MaxRecords"`
}
type FramerateLimit struct {
	MinFramerateLimit      int `json:"MinFramerateLimit"`
	MaxFramerateLobbyLimit int `json:"MaxFramerateLobbyLimit"`
	MaxFramerateGameLimit  int `json:"MaxFramerateGameLimit"`
}

type ServerConfig struct {
	IP       string      `json:"ip"`
	Hostname string      `json:"hostname"`
	Name     string      `json:"name"`
	Version  string      `json:"version"`
	Secure   bool        `json:"secure"`
	Ports    ServerPorts `json:"ports"`
}

type ServerPorts struct {
	Main      string `json:"Main"`
	Messaging string `json:"Messaging"`
	Trading   string `json:"Trading"`
	Flea      string `json:"Flea"`
	Lobby     string `json:"Lobby"`
}

type MatchMetrics struct {
	Keys                  []int `json:"Keys"`
	NetProcessingBins     []int `json:"NetProcessingBins"`
	RenderBins            []int `json:"RenderBins"`
	GameUpdateBins        []int `json:"GameUpdateBins"`
	MemoryMeasureInterval int   `json:"MemoryMeasureInterval"`
	PauseReasons          []int `json:"PauseReasons"`
}

// #endregion
