package MTGACoop

import (
	"MT-GO/database"
	"MT-GO/tools"
	"encoding/json"
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

// Fired when a new player joins the game...
func (cm *coopMatch) PlayerJoined(accountId string) {
	_, ok := cm.ConnectedPlayers[accountId]
	if !ok {
		cm.ConnectedPlayers[accountId] = struct{}{}
		fmt.Printf("this nigga %s: %s has joined", cm.ServerId, accountId)
	}
}

// Host has left
func (cm *coopMatch) KillMyself() {
	delete(coopMatches, cm.ServerId)
}

// Fired when a player leaves
func (cm *coopMatch) PlayerLeft(accountId string) {
	if accountId == cm.ServerId {
		fmt.Printf("Host nigga rage quit LOL")
		cm.KillMyself()
	}

	_, ok := cm.ConnectedPlayers[accountId]
	if ok {
		delete(cm.ConnectedPlayers, accountId)
		fmt.Printf("this nigga %s: %s has DIED", cm.ServerId, accountId)
	}

}

// Ping is used in /coop/server/update and wsOnConnection in `ProcessData` //mhm
func (cm *coopMatch) Ping(accountId string, timestamp int64) {
	pm := map[string]interface{}{ //ITS FINE ITS DOESNT MATTER MANE
		"pong": timestamp,
	}
	messageJson, err := json.Marshal(pm)
	if err != nil {
		fmt.Printf("Failed to ping %s %s", accountId, messageJson)
		return
	}
	//fucking websocoket thing here
	//sendtowebsockets shit
}

// # Status Codes:
//
// Loading: 	0
//
// InGame:  	1
//
// Complete:	2
func (cm *coopMatch) UpdateStatus(status status /*int8*/) {
	cm.Status = status
}

// End the session
//
// reason(string)
func (cm *coopMatch) EndSession(reason string) {
	fmt.Printf("[Coop] Session %s has ended with reason: %s\n", cm.ServerId, reason)
	//Websocket.SendMessageToWebsocket()
	cm.KillMyself()
}

func processWebsocketMessage(websocketMessage string) {

}

// async processMessage(msg) {
//         const msgStr = msg.toString();
//         this.processMessageString(msgStr);
//     }
//     async processMessageString(msgStr) {
//         // If is SIT serialized string -- This is NEVER stored.
//         if (msgStr.startsWith("MTC")) {
//             const messageWithoutSITPrefix = msgStr.substring(3, msgStr.length);
//             const serverId = messageWithoutSITPrefix.substring(0, 24); // get serverId (MongoIds are 24 characters)
//             const messageWithoutSITPrefixes = messageWithoutSITPrefix.substring(24, messageWithoutSITPrefix.length);
//             const match = CoopMatch_1.CoopMatch.CoopMatches[serverId];
//             if (match !== undefined) {
//                 match.ProcessData(messageWithoutSITPrefixes, this.logger);
//             }
//             return;
//         }
//         var jsonArray = this.TryParseJsonArray(msgStr);
//         if (jsonArray !== undefined) {
//             for (const key in jsonArray) {
//                 this.processObject(jsonArray[key]);
//             }
//         }
//         if (msgStr.charAt(0) !== '{')
//             return;
//         var jsonObject = JSON.parse(msgStr);
//         this.processObject(jsonObject);
//     }
//     async processObject(jsonObject) {
//         const match = CoopMatch_1.CoopMatch.CoopMatches[jsonObject["serverId"]];
//         if (match !== undefined) {
//             if (jsonObject["connect"] == true) {
//                 match.PlayerJoined(jsonObject["accountId"]);
//             }
//             else {
//                 match.ProcessData(jsonObject, this.logger);
//             }
//         }
//         this.sendToAllWebSockets(JSON.stringify(jsonObject));
//     }

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
	ConnectedPlayers   map[string]struct{}    // []string
	Characters         map[string]interface{} // []interface{}
	//State any
	//Ip any
	//Port any
	ExpectedPlayers     int16
	LastDataByAccountId interface{}
	//LastDataReceivedByAccountId interface{}
	//LastData                    interface{}
	//LastMoves                   interface{}
	//LastRotates                 interface{}
	//DamageArray                 []interface{}
	Status                 status // Loading = 0, InGame = 1, Complete = 2
	Settings               interface{}
	Loot                   interface{}
	Location               string
	Time                   int8
	TimeAndWeatherSettings timeAndWeatherSettings
	SpawnPoint             database.Vector3

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
