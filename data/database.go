// Package data contains all the data related code
package data

import (
	"MT-GO/tools"
	"MT-GO/user/mods"
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

// SetDatabase initializes the data
// var db *Database
var primary = []func(){
	setBots, setEditions, setHideout,
	setLocalLoot, setLocales, setTraders,
	setCore, setLanguages, setHandbook,
	setQuests, setItems, setWeather,
	setLocations, setCustomization, //setFlea,
}

var secondary = []func(){
	SetProfiles, IndexQuests, IndexTradeOffers,
	SetServerConfig, IndexHideoutAreas, IndexHideoutRecipes,
	IndexScavcase,
}

func SetDatabase() {
	var wg sync.WaitGroup
	numWorkers := tools.CalculateWorkers() / 3

	runTasks(&wg, primary, numWorkers)

	mods.Init()
	LoadBundleManifests()
	LoadCustomItems()

	runTasks(&wg, secondary, numWorkers)
}

func runTasks(wg *sync.WaitGroup, tasks []func(), numWorkers int) {
	workerCh := make(chan bool, numWorkers)
	completionCh := make(chan bool)

	for _, task := range tasks {
		wg.Add(1)
		go func(taskFunc func()) {
			defer wg.Done()
			workerCh <- true
			taskFunc()
			<-workerCh
			completionCh <- true
		}(task)
	}

	go func() {
		wg.Wait()
		close(completionCh)
	}()

	for range tasks {
		<-completionCh
	}
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
