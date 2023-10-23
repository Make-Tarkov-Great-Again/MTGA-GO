package coopmatch

import (
	common "MT-GO/user/mods/MTGA-Coop/mod/Common"
	config "MT-GO/user/mods/MTGA-Coop/mod/Config"
	"encoding/json"
	"fmt"
	"time"
)

type CoopMatchEndSessionMessages struct{}

const (
	HostShutdownMessage = "host-shutdown"
	HostTimeoutMessage  = "host-timeout-ws"
)

var ci *common.WebSocketHandler

func Init(instance *common.WebSocketHandler) {
	ci = instance
}

type Player struct {
	AccountId string
	IsDead    bool `json:",omitempty"`
	IsOnline  bool `json:",omitempty"`
}

type CoopMatch struct {
	CreatedDateTime             time.Time
	LastUpdateDateTime          time.Time
	ExpectedNumberOfPlayers     int
	ConnectedPlayers            []ConnectedPlayer              //map[string]Player              //TODO: Show me, but also these can be combined AND we should consider not using a fucking array/slice
	Characters                  []Character                    //map[string]Player                    //Some of the functions right now, wont work, as they "Arent the same struct" despite being the exact same
	LastDataByAccountId         map[int]map[string]interface{} // Ight let me remember which one it was
	LastDataReceivedByAccountId map[int]map[string]interface{}
	LastData                    map[string]interface{}
	LastMoves                   map[string]interface{}
	LastRotates                 map[string]interface{}
	DamageArray                 []int
	Status                      CoopMatchStatus
	Settings                    map[string]interface{}
	Loot                        map[string]interface{}
	SpawnPoint                  SpawnPoint
	Private                     bool
	Name                        string
	ServerId                    int
	Location                    string
	Time                        string
	WeatherSettings             map[string]interface{}
	CheckStillRunningInterval   *time.Timer
	SendLastDataInterval        *time.Timer
}
type CoopMatches struct {
	name        string
	serverId    string
	location    string
	playerCount int
	password    bool
	state       any
}
type Character struct {
	accountId string
	IsDead    bool
}

type SpawnPoint struct {
	X float64
	Y float64
	Z float64
}

type ConnectedPlayer struct {
	AccountID string
	IsOnline  bool
}

type CoopMatchStatus int

const (
	Loading CoopMatchStatus = iota
	InGame
	Complete
)

func NewCoopMatch(inData map[string]interface{}) *CoopMatch {

	cm := &CoopMatch{
		CreatedDateTime:             time.Now(),
		LastUpdateDateTime:          time.Now(),
		ExpectedNumberOfPlayers:     1,
		ConnectedPlayers:            []ConnectedPlayer{},
		Characters:                  []Character{},
		LastDataByAccountId:         make(map[int]map[string]interface{}),
		LastDataReceivedByAccountId: make(map[int]map[string]interface{}),
		LastData:                    make(map[string]interface{}),
		LastMoves:                   make(map[string]interface{}),
		LastRotates:                 make(map[string]interface{}),
		DamageArray:                 []int{},
		Status:                      Loading,
		Settings:                    make(map[string]interface{}),
		Loot:                        make(map[string]interface{}),
		Private:                     inData["isPrivate"].(bool),
		Name:                        inData["name"].(string),
		ServerId:                    inData["serverId"].(int),
		Location:                    inData["settings"].(map[string]interface{})["location"].(string),
		Time:                        inData["settings"].(map[string]interface{})["timeVariant"].(string),
		WeatherSettings:             inData["settings"].(map[string]interface{})["timeAndWeatherSettings"].(map[string]interface{}),
	}

	cm.CheckStillRunningInterval = time.AfterFunc(
		time.Duration(config.Instance.WebSocketTimeoutCheckStartSeconds)*time.Second,
		//Common is there as a medium to prevent cyclidic errors
		//Websocket relies on CoopMatches, and vice versa, and both of these are nesscary functions for the mod, 
		//Best work around i could come up with // Tried converting, threw "Cannot convert to type common.ConnectedPlayers"
		func() { //so you've got all this websocket shit in common... why even do a websocket.go then if it's all in common
			if !ci.AreThereAnyWebSocketsOpen(cm.ConnectedPlayers) { //This makes me want to shoot myself
				cm.EndSession(HostTimeoutMessage) //For some reason ^^ unlike everything else this wants a specific.
			} //common.ConnectedPlayers
		},
	)

	return cm
}

func (cm *CoopMatch) ProcessData(info map[string]interface{}) { //This is fucked i feel like lmfao 
	if info == nil {
		return
	}

	infoJSON, err := json.Marshal(info)
	if err != nil {
		return
	}

	if infoJSON[0] == byte('[') {
		for _, _info := range info.([]map[string]interface{}) {
			cm.ProcessData(_info)
		}
		return
	}

	if m, ok := info["m"]; ok {
		switch m {
		case "Ping":
			if t, tOk := info["t"]; tOk {
				if accountId, aOk := info["accountId"]; aOk {
					cm.Ping(accountId.(int), t.(int))
				}
			}
		case "SpawnPointForCoop":
			if x, xOk := info["x"].(float64); xOk {
				if y, yOk := info["y"].(float64); yOk {
					if z, zOk := info["z"].(float64); zOk {
						cm.SpawnPoint.X = x
						cm.SpawnPoint.Y = y
						cm.SpawnPoint.Z = z
					}
				}
			}
		case "PlayerLeft":
			if accountId, aOk := info["accountId"]; aOk {
				cm.PlayerLeft(accountId.(string))
				if len(cm.ConnectedPlayers) == 0 {
					fmt.Println("connected players is 0")
				}
				cm.EndSession(HostShutdownMessage)
			}
		case "PlayerJoined":
			if accountId, aOk := info["accountId"]; aOk {
				cm.PlayerJoined(accountId.(string))
			}
		case "PlayerSpawn":
			if accountId, aOk := info["accountId"].(int); aOk {
				if cm.LastDataByAccountId[accountId] == nil {
					cm.LastDataByAccountId[accountId] = make(map[string]interface{})
				}
				cm.LastDataByAccountId[accountId]["PlayerSpawn"] = info
			}
		case "Kill":
			if accountId, aOk := info["accountId"].(int); aOk {
				if cm.LastDataByAccountId[accountId] == nil {
					cm.LastDataByAccountId[accountId] = make(map[string]interface{})
				}
				cm.LastDataByAccountId[accountId]["Kill"] = info
				for _, c := range cm.Characters {
					if c.accountId == accountId {
						c.IsDead = true
						break
					}
				}
			}
		}
	}

	cm.LastUpdateDateTime = time.Now()
	ci.SendToWebSockets(cm.ConnectedPlayers, infoJSON)
}

func (cm *CoopMatch) UpdateStatus(inStatus CoopMatchStatus) {
	cm.Status = inStatus
}

func (cm *CoopMatch) PlayerJoined(accountId string) {
	if cm.ConnectedPlayersIndex(accountId) == -1 {
		// Create a new ConnectedPlayer instance and append it
		newConnectedPlayer := ConnectedPlayer{
			AccountID: accountId,
			IsOnline:  true,
		}
		cm.ConnectedPlayers = append(cm.ConnectedPlayers, newConnectedPlayer)
		fmt.Printf("%d: %s has joined\n", cm.ServerId, accountId)
	}
}

func (cm *CoopMatch) PlayerLeft(accountId any) {
	index := cm.ConnectedPlayersIndex(accountId)
	if index != -1 {
		cm.ConnectedPlayers = append(cm.ConnectedPlayers[:index], cm.ConnectedPlayers[index+1:]...)
		fmt.Printf("%d: %s has left\n", cm.ServerId, accountId)
		if cm.ServerId == accountId {
			fmt.Println("Player left HOST left")
			cm.EndSession(HostShutdownMessage)
		}
	}
}

func (cm *CoopMatch) Ping(accountId, timestamp int) {
	ci.SendToWebSockets([]int{accountId}, json.Marshal(map[string]interface{}{"pong": timestamp}))
}

func (cm *CoopMatch) EndSession(reason string) {
	fmt.Printf("COOP SESSION %d HAS BEEN ENDED reason: %s\n", cm.ServerId, reason)
	ci.SendToWebSockets(cm.ConnectedPlayers, json.Marshal(map[string]interface{}{"endSession": true, "reason": reason}))
	cm.Status = Complete //So this needs websocket right? We go to websocket, 
	cm.CheckStillRunningInterval.Stop()
	cm.SendLastDataInterval.Stop()
	delete(CoopMatch.CoopMatches, cm.ServerId)
	StayInTarkovMod.Instance.LocationData = make(map[string]interface{})
}

func (cm *CoopMatch) ConnectedPlayersIndex(accountId string) int {
	for i, connectedPlayer := range cm.ConnectedPlayers {
		if connectedPlayer.AccountID == accountId {
			return i
		}
	}
	return -1
}
