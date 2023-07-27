package database

import (
	"MT-GO/database/structs"
	"MT-GO/tools"
	"encoding/json"
	"strings"
)

func setBots() *structs.Bots {
	bots := structs.Bots{}

	bots.BotTypes = processBotTypes()
	bots.BotAppearance = processBotAppearance()
	bots.BotNames = processBotNames()
	return &bots
}

const botsDirectory string = botsPath + "bots/"

func processBotTypes() map[string]*structs.BotType {
	botTypes := make(map[string]*structs.BotType)

	directory, err := tools.GetDirectoriesFrom(botsDirectory)
	if err != nil {
		panic(err)
	}

	for _, directory := range directory {
		botType := structs.BotType{}
		var dirPath = botsDirectory + directory + "/"

		var diffPath = dirPath + "difficulties/"
		files, err := tools.GetFilesFrom(diffPath)
		if err != nil {
			panic(err)
		}

		difficulties := make(map[string]json.RawMessage)
		botDifficulty := structs.BotDifficulties{}
		for _, difficulty := range files {
			raw := tools.GetJSONRawMessage(diffPath + difficulty)
			name := strings.Replace(difficulty, ".json", "", -1)
			difficulties[name] = raw
		}

		jsonData, err := json.Marshal(difficulties)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(jsonData, &botDifficulty)
		if err != nil {
			panic(err)
		}
		botType.Difficulties = &botDifficulty

		healthPath := dirPath + "health.json"
		if tools.FileExist(healthPath) {
			health := structs.BotHealth{}

			raw := tools.GetJSONRawMessage(healthPath)
			err = json.Unmarshal(raw, &health)
			if err != nil {
				panic(err)
			}
			botType.Health = &health
		}

		loadoutPath := dirPath + "loadout.json"
		if tools.FileExist(loadoutPath) {
			loadout := structs.BotLoadout{}
			raw := tools.GetJSONRawMessage(loadoutPath)
			err = json.Unmarshal(raw, &loadout)
			if err != nil {
				panic(err)
			}
			botType.Loadout = &loadout
		}
		botTypes[directory] = &botType
	}

	return botTypes
}

func processBotAppearance() map[string]*structs.BotAppearance {
	botAppearance := make(map[string]*structs.BotAppearance)

	raw := tools.GetJSONRawMessage(botsPath + "appearance.json")
	err := json.Unmarshal(raw, &botAppearance)
	if err != nil {
		panic(err)
	}
	return botAppearance
}

func processBotNames() *structs.BotNames {
	names := structs.BotNames{}

	raw := tools.GetJSONRawMessage(botsPath + "names.json")
	err := json.Unmarshal(raw, &names)
	if err != nil {
		panic(err)
	}
	return &names
}
