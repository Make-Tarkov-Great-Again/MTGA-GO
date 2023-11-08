// Package database contains all the database related code
package database

import (
	"MT-GO/tools"
	"sync"
)

const (
	databaseLibPath       = "assets"
	coreFilePath          = databaseLibPath + "/core"
	airdropFilePath       = coreFilePath + "/airdrop.json"
	botTemplateFilePath   = coreFilePath + "/botTemplate.json"
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

// SetDatabase initializes the database
//var db *Database

func SetDatabase() {
	var wg sync.WaitGroup
	completionCh := make(chan struct{})

	tasks := []struct {
		name     string
		function func()
	}{
		{"Core", setCore},
		{"Items", setItems},
		{"Locales", setLocales},
		{"Languages", setLanguages},
		{"Handbook", setHandbook},
		{"Traders", setTraders},
		{"Locations", setLocationsMaster},
		{"Quests", setQuests},
		{"Hideout", setHideout},
		{"Weather", setWeather},
		{"Customization", setCustomization},
		{"Bots", setBots},
		{"Editions", setEditions},
		{"Flea", setFlea},
		{"Sacrificial Bot", setSacrificialBot},
	}

	numWorkers := tools.CalculateWorkers() / 4

	workerCh := make(chan struct{}, numWorkers)

	for _, task := range tasks {
		wg.Add(1)
		go func(taskName string, taskFunc func()) {
			defer wg.Done()
			workerCh <- struct{}{}
			taskFunc()
			<-workerCh
			completionCh <- struct{}{}
		}(task.name, task.function)
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
