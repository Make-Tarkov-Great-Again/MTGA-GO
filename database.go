package main

import (
	"MT-GO/tools"
	"fmt"
)

type DatabaseStruct struct {
	core *CoreStruct
	//items         map[string]interface{}
	//locales       map[string]interface{}
	//languages     map[string]interface{}
	//templates     map[string]interface{}
	//traders       map[string]interface{}
	//flea          map[string]interface{}
	//quests        map[string]interface{}
	//hideout       map[string]interface{}
	//locations     map[string]interface{}
	//weather       map[string]interface{}
	//customization map[string]interface{}
	//editions      map[string]interface{}
	//presets       map[string]interface{}
	//bot           map[string]interface{}
	//profiles      map[string]interface{}
	//bundles       []interface{}
}

type CoreStruct struct {
	botTemplate    map[string]interface{}
	clientSettings map[string]interface{}
	serverConfig   map[string]interface{}
	globals        map[string]interface{}
	locations      map[string]interface{}
	//gameplay        map[string]interface{}
	//hideoutSettings map[string]interface{}
	//blacklist       []interface{}
	matchMetrics map[string]interface{}
	connections  ConnectionStruct
}

type ConnectionStruct struct {
	webSocket      map[string]interface{}
	webSocketPings map[string]interface{}
}

var Database = &DatabaseStruct{}

func setDatabase() error {
	if err := setDatabaseCore(); err != nil {
		return err
	}
	return nil
}

func setDatabaseCore() error {
	core := &CoreStruct{}

	if err := setServerConfigCore(core); err != nil {
		return fmt.Errorf("error setting server config: %w", err)
	}
	if err := setMatchMetricsCore(core); err != nil {
		return fmt.Errorf("error setting match metrics: %w", err)
	}
	if err := setGlobalsCore(core); err != nil {
		return fmt.Errorf("error setting globals: %w", err)
	}
	if err := setClientSettingsCore(core); err != nil {
		return fmt.Errorf("error setting client settings: %w", err)
	}
	if err := setLocationsCore(core); err != nil {
		return fmt.Errorf("error setting locations: %w", err)
	}
	if err := setBotTemplateCore(core); err != nil {
		return fmt.Errorf("error setting bot template: %w", err)
	}

	Database.core = core

	return nil
}

func checkAndReturnIfDataPropertyExists(data interface{}) map[string]interface{} {
	ifData, ok := data.(map[string]interface{})["data"].(map[string]interface{})
	if !ok {
		return data.(map[string]interface{})
	}
	return ifData
}

func setServerConfigCore(core *CoreStruct) error {
	serverConfig, err := tools.ReadParsed("assets/server.json")

	if err != nil {
		return fmt.Errorf("error reading server.json: %w", err)
	}

	core.serverConfig = checkAndReturnIfDataPropertyExists(serverConfig)
	return nil
}

func setMatchMetricsCore(core *CoreStruct) error {
	matchMetrics, err := tools.ReadParsed("assets/matchMetrics.json")

	if err != nil {
		return fmt.Errorf("error reading matchMetrics.json: %w", err)
	}

	core.matchMetrics = checkAndReturnIfDataPropertyExists(matchMetrics)
	return nil
}

func setGlobalsCore(core *CoreStruct) error {
	globals, err := tools.ReadParsed("assets/globals.json")

	if err != nil {
		return fmt.Errorf("error reading globals.json: %w", err)
	}

	core.globals = checkAndReturnIfDataPropertyExists(globals)
	return nil
}

func setClientSettingsCore(core *CoreStruct) error {
	clientSettings, err := tools.ReadParsed("assets/client.settings.json")

	if err != nil {
		return fmt.Errorf("error reading client.settings.json: %w", err)
	}
	core.clientSettings = checkAndReturnIfDataPropertyExists(clientSettings)
	return nil
}

func setLocationsCore(core *CoreStruct) error {
	locations, err := tools.ReadParsed("assets/locations.json")

	if err != nil {
		return fmt.Errorf("error reading locations.json: %w", err)
	}

	core.locations = checkAndReturnIfDataPropertyExists(locations)
	return nil
}

func setBotTemplateCore(core *CoreStruct) error {
	botTemplate, err := tools.ReadParsed("assets/botTemplate.json")

	if err != nil {
		return fmt.Errorf("error reading botTemplate.json: %w", err)
	}

	core.botTemplate = checkAndReturnIfDataPropertyExists(botTemplate)
	return nil
}
