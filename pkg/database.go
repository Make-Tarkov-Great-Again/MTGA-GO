package pkg

import (
	"MT-GO/data"
	"fmt"
)

func GetWeather() *data.Weather {
	return data.GetWeather()
}

func GetLocations() *data.Locations {
	return data.GetLocations()
}

func GetHandbook() *data.Handbook {
	return data.GetHandbook()
}

func GetLocalGlobalsByName(lang string) (map[string]any, error) {
	return data.GetLocalesLocaleByName(lang)
}

func GetHideoutAreas() ([]map[string]any, error) {
	hideout := data.GetHideout()

	if hideout.Areas != nil {
		return hideout.Areas, nil
	}

	return nil, fmt.Errorf("Hideout areas does not exist")
}

func GetHideoutQTE() ([]map[string]any, error) {
	hideout := data.GetHideout()

	if hideout.QTE != nil {
		return hideout.QTE, nil
	}

	return nil, fmt.Errorf("Hideout QTE does not exist")
}

func GetHideoutSettings() (*data.HideoutSettings, error) {
	hideout := data.GetHideout()

	if hideout.Settings != nil {
		return hideout.Settings, nil
	}

	return nil, fmt.Errorf("Hideout Settings does not exist")
}

func GetHideoutRecipes() ([]map[string]any, error) {
	hideout := data.GetHideout()

	if hideout.Recipes != nil {
		return hideout.Recipes, nil
	}

	return nil, fmt.Errorf("Hideout Recipes does not exist")
}

func GetHideoutScavcase() ([]map[string]any, error) {
	hideout := data.GetHideout()

	if hideout.ScavCase != nil {
		return hideout.ScavCase, nil
	}

	return nil, fmt.Errorf("Hideout ScavCase does not exist")
}

func GetPlayerStorage(sessionID string) (*data.Storage, error) {
	return data.GetStorageByID(sessionID)
}

func GetPlayerProfile(sessionID string) (*data.Profile, error) {
	return data.GetProfileByUID(sessionID)
}

func GetPlayerCharacter(sessionID string) *data.Character {
	return data.GetCharacterByID(sessionID)
}

func GetServerConfigs() *data.ServerConfig {
	return data.GetServerConfig()
}

func GetItems() map[string]*data.DatabaseItem {
	return data.GetItems()
}

func GetLocalLootByNameAndIndex(locationID string, variantID int8) any {
	return data.GetLocalLootByNameAndIndex(locationID, variantID)
}

func GetBotByName(name string) (*data.BotType, error) {
	return data.GetBotByName(name)
}

func GetFillerBot() *map[string]any {
	return data.GetSacrificialBot()
}

func GetAirdropParameters() *data.AirdropParameters {
	return data.GetAirdropParameters()
}

func SetTradersResupply() int {
	return data.SetResupplyTimer()
}
