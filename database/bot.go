package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"path/filepath"
	"strings"
)

var bots = structs.Bots{}

func GetBots() *structs.Bots {
	return &bots
}

func setBots() {
	bots.BotTypes = processBotTypes()
	bots.BotAppearance = processBotAppearance()
	bots.BotNames = processBotNames()
}

func processBotTypes() map[string]*structs.BotType {
	botTypes := make(map[string]*structs.BotType)

	directory, err := tools.GetDirectoriesFrom(botsDirectory)
	if err != nil {
		panic(err)
	}

	for _, directory := range directory {
		botType := structs.BotType{}
		var dirPath = filepath.Join(botsDirectory, directory)

		var diffPath = filepath.Join(dirPath, "difficulties")
		files, err := tools.GetFilesFrom(diffPath)
		if err != nil {
			panic(err)
		}

		difficulties := make(map[string]json.RawMessage)
		botDifficulty := structs.BotDifficulties{}
		for _, difficulty := range files {
			difficultyPath := filepath.Join(diffPath, difficulty)
			raw := tools.GetJSONRawMessage(difficultyPath)
			name := strings.TrimSuffix(difficulty, ".json")
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

		healthPath := filepath.Join(dirPath, "health.json")
		if tools.FileExist(healthPath) {
			health := structs.BotHealth{}

			raw := tools.GetJSONRawMessage(healthPath)
			err = json.Unmarshal(raw, &health)
			if err != nil {
				panic(err)
			}
			botType.Health = &health
		}

		loadoutPath := filepath.Join(dirPath, "loadout.json")
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

	raw := tools.GetJSONRawMessage(filepath.Join(botsPath, "appearance.json"))
	err := json.Unmarshal(raw, &botAppearance)
	if err != nil {
		panic(err)
	}
	return botAppearance
}

func processBotNames() *structs.BotNames {
	names := structs.BotNames{}

	raw := tools.GetJSONRawMessage(filepath.Join(botsPath, "names.json"))
	err := json.Unmarshal(raw, &names)
	if err != nil {
		panic(err)
	}
	return &names
}
