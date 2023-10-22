package main

import (
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

type CoopMatch struct {
	CoopMatches                 *CoopMatches
	CreatedDateTime             time.Time
	LastUpdateDateTime          time.Time
	ExpectedNumberOfPlayers     int
	ConnectedPlayers            []int
	Characters                  []Character
	LastDataByAccountId         map[int]map[string]interface{}
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
	AccountId int
	IsDead    bool
}

type SpawnPoint struct {
	X float64
	Y float64
	Z float64
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
		ConnectedPlayers:            []int{},
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
		func() {
			if !WebSocketHandler.Instance.AreThereAnyWebSocketsOpen(cm.ConnectedPlayers) {
				cm.EndSession(HostTimeoutMessage)
			}
		},
	)

	return cm
}

func (cm *CoopMatch) ProcessData(info map[string]interface{}) {
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

	if info["m"] == "Ping" && info["t"] != nil && info["accountId"] != nil {
		cm.Ping(info["accountId"].(int), info["t"].(int))
		return
	}

	if info["m"] == "SpawnPointForCoop" {
		cm.SpawnPoint.X = info["x"].(float64)
		cm.SpawnPoint.Y = info["y"].(float64)
		cm.SpawnPoint.Z = info["z"].(float64)
		return
	}

	if info["accountId"] != nil && info["m"] == "PlayerLeft" {
		cm.PlayerLeft(info["accountId"].(int))
		if len(cm.ConnectedPlayers) == 0 {
			fmt.Println("connected players is 0")
		}
		cm.EndSession(HostShutdownMessage)
		return
	}

	if info["accountId"] != nil {
		cm.PlayerJoined(info["accountId"].(int))
	}

	if info["m"] == nil {
		cm.LastUpdateDateTime = time.Now()
		return
	}

	if info["m"] != "PlayerSpawn" {
		if cm.LastDataByAccountId[info["accountId"].(int)] == nil {
			cm.LastDataByAccountId[info["accountId"].(int)] = make(map[string]interface{})
		}
		cm.LastDataByAccountId[info["accountId"].(int)][info["m"].(string)] = info
	}

	if info["m"] == "PlayerSpawn" {
		foundExistingPlayer := false
		for _, c := range cm.Characters {
			if info["accountId"] == c.AccountId {
				foundExistingPlayer = true
				break
			}
		}
		if !foundExistingPlayer {
			cm.Characters = append(cm.Characters, Character{AccountId: info["accountId"].(int)})
		}
	}

	if info["m"] == "Kill" {
		for _, c := range cm.Characters {
			if info["accountId"] == c.AccountId {
				c.IsDead = true
				break
			}
		}
	}

	cm.LastUpdateDateTime = time.Now()
	WebSocketHandler.Instance.SendToWebSockets(cm.ConnectedPlayers, infoJSON)
}

func (cm *CoopMatch) UpdateStatus(inStatus CoopMatchStatus) {
	cm.Status = inStatus
}

func (cm *CoopMatch) PlayerJoined(accountId int) {
	if cm.ConnectedPlayersIndex(accountId) == -1 {
		cm.ConnectedPlayers = append(cm.ConnectedPlayers, accountId)
		fmt.Printf("%d: %d has joined\n", cm.ServerId, accountId)
	}
}

func (cm *CoopMatch) PlayerLeft(accountId int) {
	index := cm.ConnectedPlayersIndex(accountId)
	if index != -1 {
		cm.ConnectedPlayers = append(cm.ConnectedPlayers[:index], cm.ConnectedPlayers[index+1:]...)
		fmt.Printf("%d: %d has left\n", cm.ServerId, accountId)
		if cm.ServerId == accountId {
			fmt.Println("Player left HOST left")
			cm.EndSession(HostShutdownMessage)
		}
	}
}

func (cm *CoopMatch) Ping(accountId, timestamp int) {
	WebSocketHandler.Instance.SendToWebSockets([]int{accountId}, json.Marshal(map[string]interface{}{"pong": timestamp}))
}

func (cm *CoopMatch) EndSession(reason string) {
	fmt.Printf("COOP SESSION %d HAS BEEN ENDED reason: %s\n", cm.ServerId, reason)
	WebSocketHandler.Instance.SendToWebSockets(cm.ConnectedPlayers, json.Marshal(map[string]interface{}{"endSession": true, "reason": reason}))
	cm.Status = Complete
	cm.CheckStillRunningInterval.Stop()
	cm.SendLastDataInterval.Stop()
	delete(CoopMatch.CoopMatches, cm.ServerId)
	StayInTarkovMod.Instance.LocationData = make(map[string]interface{})
}

func (cm *CoopMatch) ConnectedPlayersIndex(accountId int) int {
	for i, id := range cm.ConnectedPlayers {
		if id == accountId {
			return i
		}
	}
	return -1
}
