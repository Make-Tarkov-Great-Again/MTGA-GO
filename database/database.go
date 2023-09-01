// Package database contains all the database related code
package database

import (
	"MT-GO/tools"
	"os"
	"path/filepath"
)

const (
	databaseLibPath       string = "assets/database"
	coreFilePath          string = databaseLibPath + "/core"
	botTemplateFilePath   string = coreFilePath + "/botTemplate.json"
	playerScavPath        string = coreFilePath + "/playerScav.json"
	MainSettingsPath      string = coreFilePath + "/client.settings.json"
	globalBotSettingsPath string = coreFilePath + "/__BotGlobalSettings.json"
	globalsFilePath       string = coreFilePath + "/globals.json"
	locationsFilePath     string = coreFilePath + "/locations.json"
	matchMetricsPath      string = coreFilePath + "/matchMetrics.json"
	serverConfigPath      string = coreFilePath + "/server.json"
	editionsDirPath       string = databaseLibPath + "/editions/"
	itemsPath             string = databaseLibPath + "/items.json"
	localesPath           string = databaseLibPath + "/locales"
	handbookPath          string = databaseLibPath + "/handbook.json"
	traderPath            string = databaseLibPath + "/traders/"
	questsPath            string = databaseLibPath + "/quests.json"
	hideoutPath           string = databaseLibPath + "/hideout/"
	weatherPath           string = databaseLibPath + "/weather.json"
	customizationPath     string = databaseLibPath + "/customization.json"
	botsPath              string = databaseLibPath + "/bot/"
	botsDirectory         string = botsPath + "bots/"
)

// InitializeDatabase initializes the database
func InitializeDatabase() {
	setRequiredFolders()
	setCore()
	setItems()
	setLocales()
	setLanguages()
	setHandbook()
	setTraders()
	setLocations()
	setQuests()
	setHideout()
	setWeather()
	setCustomization()
	setBots()
	setEditions()
	setFlea()
	setProfiles()
}

// #region Database setters

func setRequiredFolders() {
	var users string = "user"

	if !tools.FileExist(users) {
		err := os.Mkdir(users, 0755)
		if err != nil {
			panic(err)
		}
	}

	profilesPath := filepath.Join(users, "profiles")
	if !tools.FileExist(profilesPath) {
		err := os.Mkdir(profilesPath, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// #endregion

// #region Database structs

type Database struct {
	Core *Core
	//Connections *ConnectionStruct
	Items     map[string]interface{}
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
