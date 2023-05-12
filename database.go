package main

import (
	"MT-GO/tools"
	"fmt"
)

type DatabaseStruct struct {
	core *CoreStruct
	/* 	items         map[string]interface{}
	   	locales       map[string]interface{}
	   	languages     map[string]interface{}
	   	templates     map[string]interface{}
	   	traders       map[string]interface{}
	   	flea          map[string]interface{}
	   	quests        map[string]interface{}
	   	hideout       map[string]interface{}
	   	locations     map[string]interface{}
	   	weather       map[string]interface{}
	   	customization map[string]interface{}
	   	editions      map[string]interface{}
	   	presets       map[string]interface{}
	   	bot           map[string]interface{}
	   	profiles      map[string]interface{}
	   	bundles       []interface{} */
}

type CoreStruct struct {
	/* 	botTemplate     map[string]interface{}
	   	clientSettings  map[string]interface{} */
	serverConfig map[string]interface{}
	/* 	globals         map[string]interface{}
	   	gameplay        map[string]interface{}
	   	hideoutSettings map[string]interface{}
	   	blacklist       []interface{}
	   	matchMetrics    map[string]interface{}
	   	connections     ConnectionStruct */
}

/* type ConnectionStruct struct {
	webSocket      map[string]interface{}
	webSocketPings map[string]interface{}
} */

var Database = &DatabaseStruct{}

func InitializeDatabase() error {
	if err := setCoreDatabase(); err != nil {
		return err
	}
	return nil
}

func setCoreDatabase() error {
	core := &CoreStruct{}

	serverConfig, err := tools.ReadParsed("assets/server.json")
	if err != nil {
		return fmt.Errorf("error reading server.json: %w", err)
	}
	core.serverConfig = serverConfig.(map[string]interface{})
	Database.core = core

	return nil
}
