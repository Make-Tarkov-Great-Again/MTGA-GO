// Package data contains all the data related code
package data

import (
	"MT-GO/tools"
	"sync"
)

const (
	databaseLibPath       = "assets"
	coreFilePath          = databaseLibPath + "/core"
	airdropFilePath       = coreFilePath + "/airdrop.json"
	playerScavPath        = coreFilePath + "/playerScav.json"
	MainSettingsPath      = coreFilePath + "/client.settings.json"
	globalBotSettingsPath = coreFilePath + "/__BotGlobalSettings.json"
	globalsFilePath       = coreFilePath + "/globals.json"
	locationsFilePath     = coreFilePath + "/locations.json"
	matchMetricsPath      = coreFilePath + "/matchMetrics.json"
	serverConfigPath      = coreFilePath + "/server.json"
	editionsDirPath       = databaseLibPath + "/editions/"
	itemsPath             = databaseLibPath + "/items.json"
	localesPath           = databaseLibPath + "/locales"
	handbookPath          = databaseLibPath + "/handbook.json"
	traderPath            = databaseLibPath + "/traders/"
	questsPath            = databaseLibPath + "/quests.json"
	hideoutPath           = databaseLibPath + "/hideout/"
	weatherPath           = databaseLibPath + "/weather.json"
	customizationPath     = databaseLibPath + "/customization.json"
	botMainDir            = databaseLibPath + "/bot/"
	botsMainDir           = botMainDir + "bots/"
)

// SetDatabase initializes the data
// var db *Database
var databaseTasks = []func(){
	setBots, setEditions, setHideout,
	setLocalLoot, setLocales, setTraders,
	setCore, setLanguages, setHandbook,
	setQuests, setItems, setWeather,
	setLocations, setCustomization, setFlea,
}

func SetDatabase() {
	var wg sync.WaitGroup
	completionCh := make(chan bool)
	numWorkers := tools.CalculateWorkers() / 3

	workerCh := make(chan bool, numWorkers)

	for _, task := range databaseTasks {
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

	for range databaseTasks {
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
