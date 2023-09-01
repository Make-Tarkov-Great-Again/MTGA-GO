package database

import (
	"MT-GO/tools"
	"fmt"
	"net"

	"github.com/goccy/go-json"
)

var core = Core{}

func GetCore() *Core {
	return &core
}

func GetGlobals() *Globals {
	return core.Globals
}

func GetMainSettings() *MainSettings {
	return core.MainSettings
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

func setCore() {
	core.Character = setBotTemplate()
	core.Scav = setPlayerScav()
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

func setPlayerScav() *Scav {
	raw := tools.GetJSONRawMessage(playerScavPath)

	var playerScav Scav
	err := json.Unmarshal(raw, &playerScav)
	if err != nil {
		panic(err)
	}
	return &playerScav
}

func setBotTemplate() *Character {
	raw := tools.GetJSONRawMessage(botTemplateFilePath)

	var botTemplate Character
	err := json.Unmarshal(raw, &botTemplate)
	if err != nil {
		panic(err)
	}
	return &botTemplate
}

func setMainSettings() *MainSettings {
	raw := tools.GetJSONRawMessage(MainSettingsPath)

	var data MainSettings
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

func setServerConfig() *ServerConfig {
	raw := tools.GetJSONRawMessage(serverConfigPath)

	var data ServerConfig
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

func setMatchMetrics() *MatchMetrics {
	raw := tools.GetJSONRawMessage(matchMetricsPath)

	var data MatchMetrics
	err := json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

func setGlobals() *Globals {
	raw := tools.GetJSONRawMessage(globalsFilePath)

	var global = Globals{}
	err := json.Unmarshal(raw, &global)
	if err != nil {
		panic(err)
	}

	return &global
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
	MatchMetrics *MatchMetrics
}

type Scav struct {
	ID                string                 `json:"_id"`
	AID               int                    `json:"aid"`
	Savage            *string                `json:"savage"`
	Info              PlayerInfo             `json:"Info"`
	Customization     PlayerCustomization    `json:"Customization"`
	Health            HealthInfo             `json:"Health"`
	Inventory         InventoryInfo          `json:"Inventory"`
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
	Discord  string      `json:"discord"`
	Website  string      `json:"website"`
	Version  string      `json:"version"`
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
