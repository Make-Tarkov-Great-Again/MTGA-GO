package MTGACoop

import (
	"MT-GO/database"
	"MT-GO/tools"
	"fmt"
)

//TODO: [COOP MATCHES] key is accountId/sessionId of player that created the match

var coopMatches = make(map[string]*coopMatch)

func checkIfMatchExists(sessionId string) bool {
	_, ok := coopMatches[sessionId]
	return ok
}

func createCoopMatch(info *serverInfo) {
	time := tools.GetCurrentTimeInSeconds()
	match := &coopMatch{
		Private:            info.IsPrivate,
		Name:               info.Name,
		ServerId:           info.ServerId,
		Status:             loading, //the mod has this as a string so we'll need to check this later
		CreatedDateTime:    time,
		LastUpdateDateTime: time,
	}

	if info.RaidSettings == nil {
		coopMatches[info.ServerId] = match
		return
	}

	match.Location = info.RaidSettings.LocationId
	match.Time = info.RaidSettings.TimeVariant //the mod has this as a string so we'll need to check this later
	match.TimeAndWeatherSettings = info.RaidSettings.TimeAndWeatherSettings

	coopMatches[info.ServerId] = match

	//TODO: Figure out how to implement this nonsense later
	// Reference trader.go/SetResupplyTimer() for how it's potentially being done
	/*
		setTimeout(() => {
			this.CheckStillRunningInterval = setInterval(() => {
			if(!WebSocketHandler.Instance.areThereAnyWebSocketsOpen(this.ConnectedPlayers)) {
			this.endSession(CoopMatchEndSessionMessages.HOST_TIMEOUT_MESSAGE);
			}
			}, CoopConfig.Instance.webSocketTimeoutSeconds * 1000);
		}, CoopConfig.Instance.webSocketTimeoutCheckStartSeconds * 1000);
	*/

}

func getCoopMatch(sessionID string) *coopMatch {
	match, ok := coopMatches[sessionID]
	if !ok {
		fmt.Println("Match not found using ServerID:", sessionID)
		return match
	}
	return match
}

// enums are kind of gay
type status int8

const (
	loading status = iota
	inGame
	complete
)

type coopMatch struct {
	ServerId           string
	Name               string
	Private            bool
	Password           string
	JoinCode           string
	CreatedDateTime    int64
	LastUpdateDateTime int64
	ConnectedPlayers   map[string]struct{}    //[]string
	Characters         map[string]interface{} // []interface{}
	//State any
	//Ip any
	//Port any
	ExpectedPlayers             int16
	LastDataByAccountId         interface{}
	LastDataReceivedByAccountId interface{}
	LastData                    interface{}
	LastMoves                   interface{}
	LastRotates                 interface{}
	DamageArray                 []interface{}
	Status                      status // Loading = 0, InGame = 1, Complete = 2
	Settings                    interface{}
	Loot                        interface{}
	Location                    string
	Time                        int8
	TimeAndWeatherSettings      timeAndWeatherSettings
	SpawnPoint                  database.Vector3

	//sendLastDataInterval      interface{} //these are apparently private
	//checkStillRunningInterval interface{}
}

type serverInfo struct {
	ServerId     string        `json:"serverId"`
	RaidSettings *raidSettings `json:"settings"`
	Name         string        `json:"name"`
	IsPrivate    bool          `json:"isPrivate"`
	JoinCode     string        `json:"JoinCode"`
}

type timeAndWeatherSettings struct {
	IsRandomTime    bool `json:"isRandomTime"`
	IsRandomWeather bool `json:"isRandomWeather"`
	CloudinessType  int8 `json:"cloudinessType"`
	//Clear = 0, PartlyCloud = 1, Cloudy = 2, CloudyWithGaps = 3, HeavyCloudCover = 4, Thundercloud = 5
	RainType int8 `json:"rainType"`
	//NoRain = 0, Drizzling = 1, Rain = 2, Heavy = 3, Shower = 4
	WindType int8 `json:"windType"`
	//Light = 0, Moderate = 1, Strong = 2, VeryStrong = 3, Hurricane = 4
	FogType int8 `json:"fogType"`
	//NoFog = 0, Faint = 1, Fog = 2, Heavy = 3, Continuous = 4
	TimeFlowType int8 `json:"timeFlowType"`
	//x0 = 0, x0_14 = 1, x0_25 = 2, x0_5 = 3, x1 = 4, x2 = 5, x4 = 6, x8 = 7
	HourOfDay int8 `json:"hourOfDay"`
}

type botSettings struct {
	IsScavWars bool `json:"isScavWars"`
	BotsAmount int8 `json:"botsAmount"`
	//AsOnline = 0, NotBots = 1, Low = 2, Medium = 3, High = 4, Horde = 5
	BossType int8 `json:"bossType"`
	//AsOnline = 0
}

type wavesSettings struct {
	BotsAmount int8 `json:"botsAmount"`
	//AsOnline = 0, NotBots = 1, Low = 2, Medium = 3, High = 4, Horde = 5
	BotDifficulty int8 `json:"botDifficulty"`
	//AsOnline = 0, Easy = 1, Medium = 2, Hard = 3, Impossible = 4, Random = 5
	IsBosses          bool `json:"isBosses"`
	IsTaggedAndCursed bool `json:"isTaggedAndCursed"`
}

type raidSettings struct {
	KeyId string `json:"keyId"`
	Side  int8   `json:"side"`
	//Pmc = 0, Savage = 1, Random = 2
	LocationId  string `json:"location"`
	TimeVariant int8   `json:"timeVariant"`
	//CURR = 0, PAST = 1
	RaidMode int8 `json:"raidMode"`
	//Online = 0, Local = 1, Coop = 2
	MetabolismDisabled bool `json:"metabolismDisabled"`
	PlayersSpawnPlace  int8 `json:"playersSpawnPlace"`
	//SamePlace = 0, DifferentPlaces = 1, AtEndsOfTheMap = 2
	TimeAndWeatherSettings timeAndWeatherSettings `json:"timeAndWeatherSettings"`
	BotSettings            botSettings            `json:"botSettings"`
	WavesSettings          wavesSettings          `json:"wavesSettings"`
}
