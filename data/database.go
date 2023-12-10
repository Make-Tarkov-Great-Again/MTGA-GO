// Package data contains all the data related code
package data

import (
	"MT-GO/tools"
	"github.com/alphadose/haxmap"
	"sync"
)

const (
	databaseLibPath   = "assets"
	coreFilePath      = databaseLibPath + "/core"
	airdropFilePath   = coreFilePath + "/airdrop.json"
	playerScavPath    = coreFilePath + "/playerScav.json"
	MainSettingsPath  = coreFilePath + "/client.settings.json"
	globalsFilePath   = coreFilePath + "/globals.json"
	locationsFilePath = coreFilePath + "/locations.json"
	locationsPath     = databaseLibPath + "/locations"
	matchMetricsPath  = coreFilePath + "/matchMetrics.json"
	serverConfigPath  = coreFilePath + "/server.json"
	editionsDirPath   = databaseLibPath + "/editions/"
	itemsPath         = databaseLibPath + "/items.json"
	localesPath       = databaseLibPath + "/locales"
	handbookPath      = databaseLibPath + "/handbook.json"
	traderPath        = databaseLibPath + "/traders/"
	questsPath        = databaseLibPath + "/quests.json"
	hideoutPath       = databaseLibPath + "/hideout/"
	customizationPath = databaseLibPath + "/customization.json"
	botMainDir        = databaseLibPath + "/bot/"
	botsMainDir       = botMainDir + "bots/"
)

var db *database

type database struct {
	cache         *Cache
	core          *Core
	customization *haxmap.Map[string, *Customization] //map[string]*Customization
	bot           *Bots
	edition       *haxmap.Map[string, *Edition] //map[string]*Edition
	template      *Template
	hideout       *Hideout
	item          *haxmap.Map[string, *DatabaseItem]
	location      *Location
	locale        *haxmap.Map[string, *Locale]  //map[string]*Locale
	profile       *haxmap.Map[string, *Profile] //map[string]*Profile
	trader        *Traders                      //map[string]*Trader
	quest         *Quest
	ragfair       *Ragfair
	weather       *Weather
}

var workers = tools.CalculateWorkers() / 3

func SetPrimaryDatabase() {
	db = &database{
		cache: &Cache{
			player: haxmap.New[string, *PlayerCache](),
			channel: &Channels{
				Template: &Channel{
					Status: "ok",
					Notifier: &Notifier{
						Server:         "",
						ChannelID:      "",
						URL:            "",
						NotifierServer: "",
						WS:             "",
					},
					NotifierServer: "",
				},
				channels: haxmap.New[string, Channel](),
			},
			websocket: haxmap.New[string, *Connect](),
			nicknames: haxmap.New[string, struct{}](),
		},
	}

	wg := &sync.WaitGroup{}
	tools.RunTasks(wg,
		[]func(){
			setBots,
			setEditions,
			setHideout,
			setLocales,
			setTraders,
			setCore,
			setHandbook,
			setQuests,
			setItems,
			setWeather,
			setLocations,
			setCustomization,
		},
		workers)
}

func SetCache() {
	wg := &sync.WaitGroup{}
	tools.RunTasks(wg,
		[]func(){
			setProfiles,
			setQuestLookup,
			setTraderOfferLookup,
			setServerConfig,
			setHideoutAreaLookup,
			setHideoutRecipeLookup,
			setScavcaseRecipeLookup,
			setCachedResponses,
			setHandbookIndex,
		},
		workers)

	setProfiles()
}

// #region Database setters

// #endregion

// #region Database structs

type Vector3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Prefab struct {
	Path string `json:"path"`
	Rcid string `json:"rcid"`
}

type Value struct {
	Value int `json:"value"`
}

type PriceModifier struct {
	PriceModifier float64 `json:"PriceModifier"`
}

type RepairStrategy struct {
	BuffTypes []string `json:"BuffTypes"`
	Filter    []string `json:"Filter"`
}

// #endregion
