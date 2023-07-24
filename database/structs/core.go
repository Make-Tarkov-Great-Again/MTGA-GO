package structs

type CoreStruct struct {
	BotTemplate    PlayerTemplate
	ClientSettings ClientSettings
	ServerConfig   ServerConfig
	Globals        Globals
	Locations      Locations
	//gameplay        map[string]interface{}
	//blacklist       []interface{}
	MatchMetrics MatchMetrics
}

type ClientSettings struct {
	MemoryManagementSettings struct {
		HeapPreAllocationEnabled               bool `json:"HeapPreAllocationEnabled"`
		HeapPreAllocationMB                    int  `json:"HeapPreAllocationMB"`
		OverrideRamCleanerSettings             bool `json:"OverrideRamCleanerSettings"`
		RamCleanerEnabled                      bool `json:"RamCleanerEnabled"`
		GigabytesRequiredToDisableGCDuringRaid int  `json:"GigabytesRequiredToDisableGCDuringRaid"`
		AggressiveGC                           bool `json:"AggressiveGC"`
	} `json:"MemoryManagementSettings"`
	ReleaseProfiler struct {
		Enabled            bool `json:"Enabled"`
		RecordTriggerValue int  `json:"RecordTriggerValue"`
		MaxRecords         int  `json:"MaxRecords"`
	} `json:"ReleaseProfiler"`
	FramerateLimit struct {
		MinFramerateLimit      int `json:"MinFramerateLimit"`
		MaxFramerateLobbyLimit int `json:"MaxFramerateLobbyLimit"`
		MaxFramerateGameLimit  int `json:"MaxFramerateGameLimit"`
	} `json:"FramerateLimit"`
	ClientSendRateLimit            int       `json:"ClientSendRateLimit"`
	TurnOffLogging                 bool      `json:"TurnOffLogging"`
	NVidiaHighlights               bool      `json:"NVidiaHighlights"`
	WebDiagnosticsEnabled          bool      `json:"WebDiagnosticsEnabled"`
	KeepAliveInterval              int       `json:"KeepAliveInterval"`
	GroupStatusInterval            int       `json:"GroupStatusInterval"`
	GroupStatusButtonInterval      int       `json:"GroupStatusButtonInterval"`
	PingServersInterval            int       `json:"PingServersInterval"`
	PingServerResultSendInterval   int       `json:"PingServerResultSendInterval"`
	WeaponOverlapDistanceCulling   int       `json:"WeaponOverlapDistanceCulling"`
	FirstCycleDelaySeconds         int       `json:"FirstCycleDelaySeconds"`
	SecondCycleDelaySeconds        int       `json:"SecondCycleDelaySeconds"`
	NextCycleDelaySeconds          int       `json:"NextCycleDelaySeconds"`
	AdditionalRandomDelaySeconds   int       `json:"AdditionalRandomDelaySeconds"`
	Mark502and504AsNonImportant    bool      `json:"Mark502and504AsNonImportant"`
	DefaultRetriesCount            int       `json:"DefaultRetriesCount"`
	CriticalRetriesCount           int       `json:"CriticalRetriesCount"`
	AFKTimeoutSeconds              int       `json:"AFKTimeoutSeconds"`
	RequestsMadeThroughLobby       []string  `json:"RequestsMadeThroughLobby"`
	LobbyKeepAliveInterval         int       `json:"LobbyKeepAliveInterval"`
	RequestConfirmationTimeouts    []float32 `json:"RequestConfirmationTimeouts"`
	ShouldEstablishLobbyConnection bool      `json:"ShouldEstablishLobbyConnection"`
}

type ServerConfig struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Discord  string `json:"discord"`
	Website  string `json:"website"`
	Version  string `json:"version"`
}

type MatchMetrics struct {
	Keys                  []int `json:"Keys"`
	NetProcessingBins     []int `json:"NetProcessingBins"`
	RenderBins            []int `json:"RenderBins"`
	GameUpdateBins        []int `json:"GameUpdateBins"`
	MemoryMeasureInterval int   `json:"MemoryMeasureInterval"`
}
