package data

import (
	"MT-GO/tools"
	"log"

	"github.com/goccy/go-json"
)

var core = Core{
	Scav:              new(Scav),
	MainSettings:      new(MainSettings),
	ServerConfig:      new(ServerConfig),
	Globals:           new(Globals),
	MatchMetrics:      new(MatchMetrics),
	AirdropParameters: new(AirdropParameters),
}

// #region Core getters

func GetGlobals() *Globals {
	return core.Globals
}

func GetMainSettings() *MainSettings {
	return core.MainSettings
}

func GetAirdropParameters() *AirdropParameters {
	return core.AirdropParameters
}

func GetServerConfig() *ServerConfig {
	return core.ServerConfig
}

func GetPlayerScav() *Scav {
	return core.Scav
}

// #endregion

// #region Core setters

func setCore() {
	done := make(chan bool)

	go func() {
		raw := tools.GetJSONRawMessage(playerScavPath)
		if err := json.UnmarshalNoEscape(raw, &core.Scav); err != nil {
			log.Fatalln(err)
		}
		done <- true
	}()
	go func() {
		raw := tools.GetJSONRawMessage(MainSettingsPath)
		if err := json.UnmarshalNoEscape(raw, &core.MainSettings); err != nil {
			log.Fatalln(err)
		}
		done <- true
	}()
	go func() {
		raw := tools.GetJSONRawMessage(serverConfigPath)
		if err := json.UnmarshalNoEscape(raw, &core.ServerConfig); err != nil {
			log.Fatalln(err)
		}
		done <- true
	}()
	go func() {
		setGlobals()
		done <- true
	}()
	go func() {
		raw := tools.GetJSONRawMessage(matchMetricsPath)
		if err := json.UnmarshalNoEscape(raw, &core.MatchMetrics); err != nil {
			log.Println(err)
		}
		done <- true
	}()
	go func() {
		raw := tools.GetJSONRawMessage(airdropFilePath)
		if err := json.UnmarshalNoEscape(raw, &core.AirdropParameters); err != nil {
			log.Fatalln(err)
		}
		done <- true
	}()

	for i := 0; i < 6; i++ {
		<-done
	}
}

// #endregion

// #region Core structs

type Core struct {
	Scav              *Scav
	MainSettings      *MainSettings
	ServerConfig      *ServerConfig
	Globals           *Globals
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
	ID                string              `json:"_id"`
	AID               int                 `json:"aid"`
	Savage            *string             `json:"savage"`
	Info              PlayerInfo          `json:"Info"`
	Customization     PlayerCustomization `json:"Customization"`
	Health            HealthInfo          `json:"Health"`
	Inventory         Inventory           `json:"Inventory"`
	Skills            PlayerSkills        `json:"Skills"`
	Stats             PlayerStats         `json:"Stats"`
	Encyclopedia      map[string]bool     `json:"Encyclopedia"`
	ConditionCounters ConditionCounters   `json:"ConditionCounters"`
	BackendCounters   map[string]any      `json:"BackendCounters"`
	InsuredItems      []any               `json:"InsuredItems"`
	Hideout           any                 `json:"Hideout"`
	Bonuses           []any               `json:"Bonuses"`
	Notes             struct {
		Notes [][]any `json:"Notes"`
	} `json:"Notes"`
	Quests       []any             `json:"Quests"`
	RagfairInfo  PlayerRagfairInfo `json:"RagfairInfo"`
	WishList     []any             `json:"WishList"`
	TradersInfo  []any             `json:"TradersInfo"`
	UnlockedInfo struct {
		UnlockedProductionRecipe []any `json:"unlockedProductionRecipe"`
	} `json:"UnlockedInfo"`
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

type MatchMetrics struct {
	Keys                  []int `json:"Keys"`
	NetProcessingBins     []int `json:"NetProcessingBins"`
	RenderBins            []int `json:"RenderBins"`
	GameUpdateBins        []int `json:"GameUpdateBins"`
	MemoryMeasureInterval int   `json:"MemoryMeasureInterval"`
	PauseReasons          []int `json:"PauseReasons"`
}

// #endregion
