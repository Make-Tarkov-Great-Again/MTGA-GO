package database

import (
	"MT-GO/tools"
	"log"
	"path/filepath"
	"strings"

	"github.com/goccy/go-json"
)

// #region Bot getters

var bots = Bots{}

func GetBots() *Bots {
	return &bots
}

func GetBotTypeByName(name string) *BotType {
	botType, ok := bots.BotTypes[name]
	if !ok {
		log.Println("Bot of Type Name", name, "does not exist, returning nil!")
		return nil
	}
	return botType
}

func GetBotTypeDifficultyByName(name string, diff string) any {
	botType := GetBotTypeByName(name)
	if botType == nil {
		return nil
	}

	difficulty, ok := botType.Difficulties[diff]
	if !ok {
		log.Println("Difficulty", diff, "does not exist on Bot Type", name, ",returning nil!")
		return nil
	}

	return difficulty
}

// #endregion

// #region Bot setters

func setBots() {
	bots.BotTypes = setBotTypes()
	bots.BotAppearance = setBotAppearance()
	bots.BotNames = setBotNames()
}

func setBotTypes() map[string]*BotType {
	directory, err := tools.GetDirectoriesFrom(botsMainDir)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a channel to collect the results
	resultChan := make(chan map[string]*BotType, len(directory))

	for directory := range directory {
		go func(dir string) {
			botType := setBotType(filepath.Join(botsMainDir, dir))
			resultChan <- map[string]*BotType{dir: botType}
		}(directory)
	}

	// Create a map to store the results
	botTypes := make(map[string]*BotType)

	// Collect results from the channel
	for i := 0; i < len(directory); i++ {
		result := <-resultChan
		for key, value := range result {
			botTypes[key] = value
		}
	}

	return botTypes
}

func setBotType(dirPath string) *BotType {
	botType := new(BotType)

	var diffPath = filepath.Join(dirPath, "difficulties")
	files, err := tools.GetFilesFrom(diffPath)
	if err != nil {
		log.Fatalln(err)
	}

	difficulties := make(map[string]json.RawMessage)
	botDifficulty := map[string]any{}
	for difficulty := range files {
		difficultyPath := filepath.Join(diffPath, difficulty)
		raw := tools.GetJSONRawMessage(difficultyPath)
		name := strings.TrimSuffix(difficulty, ".json")
		difficulties[name] = raw
	}

	jsonData, err := json.Marshal(difficulties)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(jsonData, &botDifficulty)
	if err != nil {
		log.Fatalln(err)
	}
	botType.Difficulties = botDifficulty

	healthPath := filepath.Join(dirPath, "health.json")
	if tools.FileExist(healthPath) {
		health := make(map[string]any)

		raw := tools.GetJSONRawMessage(healthPath)
		err = json.Unmarshal(raw, &health)
		if err != nil {
			log.Fatalln(err)
		}
		botType.Health = health
	}

	loadoutPath := filepath.Join(dirPath, "loadout.json")
	if tools.FileExist(loadoutPath) {
		loadout := new(BotLoadout)
		raw := tools.GetJSONRawMessage(loadoutPath)
		err = json.Unmarshal(raw, &loadout)
		if err != nil {
			log.Fatalln(err)
		}
		botType.Loadout = loadout
	}

	return botType
}

func setBotAppearance() map[string]*BotAppearance {
	botAppearance := make(map[string]*BotAppearance)

	raw := tools.GetJSONRawMessage(filepath.Join(botMainDir, "appearance.json"))
	err := json.Unmarshal(raw, &botAppearance)
	if err != nil {
		log.Fatalln(err)
	}
	return botAppearance
}

func setBotNames() *BotNames {
	names := new(BotNames)

	raw := tools.GetJSONRawMessage(filepath.Join(botMainDir, "names.json"))
	err := json.Unmarshal(raw, names)
	if err != nil {
		log.Fatalln(err)
	}
	return names
}

// #endregion

// #region Bot structs

type Bots struct {
	BotTypes      map[string]*BotType
	BotAppearance map[string]*BotAppearance
	BotNames      *BotNames
}

type BotNames struct {
	BossGluhar       []string `json:"bossGluhar,omitempty"`
	BossZryachiy     []string `json:"bossZryachiy,omitempty"`
	FollowerZryachiy []string `json:"followerZryachiy,omitempty"`
	GeneralFollower  []string `json:"generalFollower,omitempty"`
	BossKilla        []string `json:"bossKilla,omitempty"`
	BossBully        []string `json:"bossBully,omitempty"`
	FollowerBully    []string `json:"followerBully,omitempty"`
	BossKojaniy      []string `json:"bossKojaniy,omitempty"`
	FollowerKojaniy  []string `json:"followerKojaniy,omitempty"`
	BossSanitar      []string `json:"bossSanitar,omitempty"`
	FollowerSanitar  []string `json:"followerSanitar,omitempty"`
	BossTagilla      []string `json:"bossTagilla,omitempty"`
	FollowerTagilla  []string `json:"followerTagilla,omitempty"`
	FollowerBigPipe  []string `json:"followerBigPipe,omitempty"`
	FollowerBirdEye  []string `json:"followerBirdEye,omitempty"`
	BossKnight       []string `json:"bossKnight,omitempty"`
	Gifter           []string `json:"gifter,omitempty"`
	Sectantpriest    []string `json:"sectantpriest,omitempty"`
	Sectantwarrior   []string `json:"sectantwarrior,omitempty"`
	Normal           []string `json:"normal,omitempty"`
	Scav             []string `json:"scav,omitempty"`
}

type BotAppearance struct {
	Voice []string
	Body  []string
	Head  []string
	Hands []string
	Feet  []string
}

type BotType struct {
	Difficulties map[string]any `json:"difficulties,omitempty"`
	Health       map[string]any `json:"health,omitempty"`
	Loadout      *BotLoadout    `json:"loadout,omitempty"`
}

type BotLoadout struct {
	Earpiece        []string `json:"earpiece,omitempty"`
	Headerwear      []string `json:"headerwear,omitempty"`
	Facecover       []string `json:"facecover,omitempty"`
	BodyArmor       []string `json:"bodyArmor,omitempty"`
	Vest            []string `json:"vest,omitempty"`
	Backpack        []string `json:"backpack,omitempty"`
	PrimaryWeapon   []string `json:"primaryWeapon,omitempty"`
	SecondaryWeapon []string `json:"secondaryWeapon,omitempty"`
	Holster         []string `json:"holster,omitempty"`
	Melee           []string `json:"melee,omitempty"`
	Pocket          []string `json:"pocket,omitempty"`
}

// #endregion
