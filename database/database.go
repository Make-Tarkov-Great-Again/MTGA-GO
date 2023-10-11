// Package database contains all the database related code
package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"MT-GO/tools"
)

const (
	databaseLibPath       = "assets"
	coreFilePath          = databaseLibPath + "/core"
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
	botsPath              = databaseLibPath + "/bot/"
	botsDirectory         = botsPath + "bots/"
)

// InitializeDatabase initializes the database
func InitializeDatabase() {
	setRequiredFolders()

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
	}

	numWorkers := tools.CalculateWorkers() / 4

	workerCh := make(chan struct{}, numWorkers)

	startTime := time.Now()
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

	setProfiles()
	endTime := time.Now()
	fmt.Printf("\n\nDatabase initialized in %s with %d workers\n", endTime.Sub(startTime), numWorkers)
}

// #region Database setters

func setRequiredFolders() {
	var users string = "user"

	if !tools.FileExist(users) {
		err := os.Mkdir(users, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}

	profilesPath := filepath.Join(users, "profiles")
	if !tools.FileExist(profilesPath) {
		err := os.Mkdir(profilesPath, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// #endregion

// #region Database structs

type Database struct {
	Core *Core
	//Connections *ConnectionStruct

	Items     map[string]*DatabaseItem
	Locales   *Locale
	Languages map[string]interface{}
	Handbook  *Handbook
	Traders   map[string]*Trader
	Flea      *Flea
	Quests    map[string]interface{}
	Hideout   *Hideout

	Locations     *Locations
	Weather       *Weather
	Customization map[string]interface{}
	Editions      map[string]interface{}
	Bot           *Bots
	Profiles      map[string]*Profile
	//bundles  []map[string]interface{}
}

// #endregion
