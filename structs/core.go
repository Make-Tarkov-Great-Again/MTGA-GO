package structs

type Core struct {
	PlayerTemplate    *PlayerTemplate
	PlayerScav        *PlayerScavTemplate
	MainSettings      *MainSettings
	ServerConfig      *ServerConfig
	Globals           *Globals
	GlobalBotSettings *map[string]interface{}
	//gameplay        map[string]interface{}
	//blacklist       []interface{}
	MatchMetrics *MatchMetrics
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
