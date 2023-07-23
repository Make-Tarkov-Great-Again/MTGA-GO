package structs

type Locations struct {
	Locations map[string]Location `json:"locations"`
	Paths     []Path              `json:"paths"`
}

type Location struct {
	Enabled    bool   `json:"Enabled,omitempty"`
	EnableCoop bool   `json:"EnableCoop,omitempty"`
	Locked     bool   `json:"Locked,omitempty"`
	Name       string `json:"Name,omitempty"`
	Scene      struct {
		Path string `json:"path,omitempty"`
		RCID string `json:"rcid,omitempty"`
	} `json:"Scene,omitempty"`
	Area                    float32                  `json:"Area,omitempty"`
	RequiredPlayerLevel     int                      `json:"RequiredPlayerLevel,omitempty"`
	PmcMaxPlayersInGroup    int                      `json:"PmcMaxPlayersInGroup,omitempty"`
	ScavMaxPlayersInGroup   int                      `json:"ScavMaxPlayersInGroup,omitempty"`
	MinPlayers              int                      `json:"MinPlayers,omitempty"`
	MaxPlayers              int                      `json:"MaxPlayers,omitempty"`
	MaxCoopGroup            int                      `json:"MaxCoopGroup,omitempty"`
	IconX                   int                      `json:"IconX,omitempty"`
	IconY                   int                      `json:"IconY,omitempty"`
	Waves                   []Wave                   `json:"waves,omitempty"`
	Limits                  [0]string                `json:"limits,omitempty"`
	AveragePlayTime         int                      `json:"AveragePlayTime,omitempty"`
	AveragePlayerLevel      int                      `json:"AveragePlayerLevel,omitempty"`
	EscapeTimeLimit         int                      `json:"EscapeTimeLimit,omitempty"`
	EscapeTimeLimitCoop     int                      `json:"EscapeTimeLimitCoop,omitempty"`
	Rules                   string                   `json:"Rules,omitempty"`
	IsSecret                bool                     `json:"IsSecret,omitempty"`
	Doors                   [0]string                `json:"doors,omitempty"`
	MaxDistToFreePoint      int                      `json:"MaxDistToFreePoint,omitempty"`
	MinDistToFreePoint      int                      `json:"MinDistToFreePoint,omitempty"`
	MaxBotPerZone           int                      `json:"MaxBotPerZone,omitempty"`
	OpenZones               string                   `json:"OpenZones,omitempty"`
	OcculsionCullingEnabled bool                     `json:"OcculsionCullingEnabled,omitempty"`
	OldSpawn                bool                     `json:"OldSpawn,omitempty"`
	NewSpawn                bool                     `json:"NewSpawn,omitempty"`
	BotMax                  int                      `json:"BotMax,omitempty"`
	BotStart                int                      `json:"BotStart,omitempty"`
	BotStop                 int                      `json:"BotStop,omitempty"`
	BotSpawnTimeOnMin       int                      `json:"BotSpawnTimeOnMin,omitempty"`
	BotSpawnTimeOnMax       int                      `json:"BotSpawnTimeOnMax,omitempty"`
	BotSpawnTimeOffMin      int                      `json:"BotSpawnTimeOffMin,omitempty"`
	BotSpawnTimeOffMax      int                      `json:"BotSpawnTimeOffMax,omitempty"`
	BotEasy                 int                      `json:"BotEasy,omitempty"`
	BotNormal               int                      `json:"BotNormal,omitempty"`
	BotHard                 int                      `json:"BotHard,omitempty"`
	BotImpossible           int                      `json:"BotImpossible,omitempty"`
	BotAssault              int                      `json:"BotAssault,omitempty"`
	BotMarksman             int                      `json:"BotMarksman,omitempty"`
	DisabledScavExits       string                   `json:"DisabledScavExits,omitempty"`
	AccessKeys              []string                 `json:"AccessKeys,omitempty"`
	UnixDateTime            int                      `json:"UnixDateTime,omitempty"`
	MinMaxBots              []MinMaxBot              `json:"MinMaxBots,omitempty"`
	BotLocationModifier     BotLocationModifier      `json:"BotLocationModifier,omitempty"`
	Exits                   []Exit                   `json:"exits,omitempty"`
	DisabledForScav         bool                     `json:"DisabledForScav,omitempty"`
	BossLocationSpawn       []BossLocationSpawn      `json:"BossLocationSpawn,omitempty"`
	SpawnPointParams        []SpawnPointParams       `json:"SpawnPointParams,omitempty"`
	MaxItemCountInLocation  []MaxItemCountInLocation `json:"maxItemCountInLocation,omitempty"`
	NameId                  string                   `json:"Id,omitempty"`
	LocationId              string                   `json:"_Id,omitempty"`
	Loot                    []Loot                   `json:"Loot,omitempty"`
	Banners                 []Banner                 `json:"Banners,omitempty"`
}

type MaxItemCountInLocation struct {
	TemplateID string `json:"TemplateId,omitempty"`
	Value      int    `json:"Value,omitempty"`
}

type MinMaxBot struct {
	Min           int    `json:"min,omitempty"`
	Max           int    `json:"max,omitempty"`
	WildSpawnType string `json:"WildSpawnType,omitempty"`
}

type Loot struct {
	Id              string     `json:"Id,omitempty"`
	IsStatic        bool       `json:"IsStatic,omitempty"`
	UseGravity      bool       `json:"useGravity,omitempty"`
	RandomRotation  bool       `json:"randomRotation,omitempty"`
	Position        XYZ        `json:"Position,omitempty"`
	Rotation        XYZ        `json:"Rotation,omitempty"`
	IsGroupPosition bool       `json:"IsGroupPosition,omitempty"`
	GroupPositions  []string   `json:"GroupPositions,omitempty"`
	Root            string     `json:"Root,omitempty"`
	Items           []LootItem `json:"Items,omitempty"`
}

type LootItem struct {
	Id       string `json:"_id,omitempty"`
	Tpl      string `json:"_tpl,omitempty"`
	ParentId string `json:"parentId,omitempty,omitempty"`
	SlotId   string `json:"slotId,omitempty,omitempty"`
	Location XYZ    `json:"location,omitempty,omitempty"`
	Upd      struct {
		FireMode struct {
			FireMode string `json:"FireMode,omitempty,omitempty"`
		} `json:"FireMode,omitempty,omitempty"`
		StackObjectsCount int `json:"StackObjectsCount,omitempty,omitempty"`
	} `json:"upd,omitempty,omitempty"`
}

type Wave struct {
	Number        int    `json:"number,omitempty"`
	TimeMin       int    `json:"time_min,omitempty"`
	TimeMax       int    `json:"time_max,omitempty"`
	SlotsMin      int    `json:"slots_min,omitempty"`
	SlotsMax      int    `json:"slots_max,omitempty"`
	SpawnPoints   string `json:"SpawnPoints,omitempty"`
	BotSide       string `json:"BotSide,omitempty"`
	BotPreset     string `json:"BotPreset,omitempty"`
	IsPlayers     bool   `json:"isPlayers,omitempty"`
	WildSpawnType string `json:"WildSpawnType,omitempty"`
}

type Banner struct {
	ID  string `json:"id,omitempty"`
	Pic struct {
		Path string `json:"path,omitempty"`
		RCID string `json:"rcid,omitempty"`
	} `json:"pic,omitempty"`
}

type Path struct {
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
}

type BossLocationSpawn struct {
	BossName            string    `json:"BossName,omitempty"`
	BossChance          float32   `json:"BossChance,omitempty"`
	BossZone            string    `json:"BossZone,omitempty"`
	BossPlayer          bool      `json:"BossPlayer,omitempty"`
	BossDifficult       string    `json:"BossDifficult,omitempty"`
	BossEscortType      string    `json:"BossEscortType,omitempty"`
	BossEscortDifficult string    `json:"BossEscortDifficult,omitempty"`
	BossEscortAmount    string    `json:"BossEscortAmount,omitempty"`
	Time                int       `json:"Time,omitempty"`
	Supports            []Support `json:"Supports,omitempty"`
	RandomTimeSpawn     bool      `json:"RandomTimeSpawn,omitempty"`
}

type Support struct {
	BossEscortType      string   `json:"BossEscortType,omitempty"`
	BossEscortDifficult []string `json:"BossEscortDifficult,omitempty"`
	BossEscortAmount    string   `json:"BossEscortAmount,omitempty"`
}

type ColliderParams struct {
	Parent string `json:"_parent,omitempty"`
	Props  struct {
		Center XYZ     `json:"Center,omitempty"`
		Radius float32 `json:"Radius,omitempty"`
	} `json:"_props,omitempty"`
}

type SpawnPointParams struct {
	ID                 string         `json:"Id,omitempty"`
	Position           XYZ            `json:"Position,omitempty"`
	Rotation           float32        `json:"Rotation,omitempty"`
	Sides              []string       `json:"Sides,omitempty"`
	Categories         []string       `json:"Categories,omitempty"`
	Infiltration       string         `json:"Infiltration,omitempty"`
	DelayToCanSpawnSec float32        `json:"DelayToCanSpawnSec,omitempty"`
	ColliderParams     ColliderParams `json:"ColliderParams,omitempty"`
	BotZoneName        string         `json:"BotZoneName,omitempty"`
}

type Exit struct {
	Name               string `json:"Name,omitempty"`
	EntryPoints        string `json:"EntryPoints,omitempty"`
	Chance             int    `json:"Chance,omitempty"`
	MinTime            int    `json:"MinTime,omitempty"`
	MaxTime            int    `json:"MaxTime,omitempty"`
	PlayersCount       int    `json:"PlayersCount,omitempty"`
	ExfiltrationTime   int    `json:"ExfiltrationTime,omitempty"`
	PassageRequirement string `json:"PassageRequirement,omitempty"`
	ExfiltrationType   string `json:"ExfiltrationType,omitempty"`
	RequiredSlot       string `json:"RequiredSlot,omitempty"`
	Id                 string `json:"Id,omitempty"`
	RequirementTip     string `json:"RequirementTip,omitempty"`
	Count              int    `json:"Count,omitempty"`
}

type Exits struct {
	Exits []Exit `json:"exits,omitempty"`
}

type BotLocationModifier struct {
	AccuracySpeed          float32 `json:"AccuracySpeed,omitempty"`
	Scattering             float32 `json:"Scattering,omitempty"`
	GainSight              float32 `json:"GainSight,omitempty"`
	MarksmanAccuratyCoef   float32 `json:"MarksmanAccuratyCoef,omitempty"`
	VisibleDistance        float32 `json:"VisibleDistance,omitempty"`
	DistToPersueAxemanCoef float32 `json:"DistToPersueAxemanCoef,omitempty"`
	KhorovodChance         int     `json:"KhorovodChance,omitempty"`
}
