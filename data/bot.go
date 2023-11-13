package data

import (
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"path/filepath"
)

var bots = Bots{
	BotTypes:      make(map[string]*BotType),
	BotAppearance: make(map[string]*BotAppearance),
	BotNames:      new(BotNames),
}
var sacrificialBot DummyBot

const (
	botNotExist        string = "Bot %s does not exist"
	difficultyNotExist string = "Difficulty %s does not exist on Bot %s"
)

func setBots() {
	done := make(chan bool)

	go func() {
		directory, err := tools.GetDirectoriesFrom(botsMainDir)
		if err != nil {
			log.Fatalln(err)
		}

		for dir := range directory {
			bots.BotTypes[dir] = setBotType(filepath.Join(botsMainDir, dir))
		}
		done <- true
	}()
	go func() {
		raw := tools.GetJSONRawMessage(filepath.Join(botMainDir, "appearance.json"))
		if err := json.Unmarshal(raw, &bots.BotAppearance); err != nil {
			log.Fatalln(err)
		}
		done <- true
	}()
	go func() {
		raw := tools.GetJSONRawMessage(filepath.Join(botMainDir, "names.json"))
		if err := json.Unmarshal(raw, bots.BotNames); err != nil {
			log.Fatalln(err)
		}
		done <- true
	}()
	go func() {
		raw := tools.GetJSONRawMessage(filepath.Join(databaseLibPath, "sacrificialBot.json"))
		if err := json.Unmarshal(raw, &sacrificialBot); err != nil {
			log.Println(err)
		}
		done <- true
	}()

	for i := 0; i < 4; i++ {
		<-done
	}
}

func setBotType(dirPath string) *BotType {
	botType := &BotType{
		Difficulties: make(map[string]map[string]any),
		Health:       make(map[string]any),
		Loadout:      &BotLoadout{},
	}

	diffPath := filepath.Join(dirPath, "difficulties")
	files, err := tools.GetFilesFrom(diffPath)
	if err != nil {
		log.Println(err)
		return nil
	}

	for difficulty := range files {
		botDifficulty := make(map[string]any)
		difficultyPath := filepath.Join(diffPath, difficulty)
		raw := tools.GetJSONRawMessage(difficultyPath)
		if err = json.Unmarshal(raw, &botDifficulty); err != nil {
			log.Println(err)
			return nil
		}
		botType.Difficulties[difficulty[:len(difficulty)-5]] = botDifficulty
	}

	healthPath := filepath.Join(dirPath, "health.json")
	if tools.FileExist(healthPath) {
		raw := tools.GetJSONRawMessage(healthPath)
		if err = json.Unmarshal(raw, &botType.Health); err != nil {
			log.Println(err)
			return nil
		}
	}

	loadoutPath := filepath.Join(dirPath, "loadout.json")
	if tools.FileExist(loadoutPath) {
		raw := tools.GetJSONRawMessage(loadoutPath)
		if err = json.Unmarshal(raw, &botType.Loadout); err != nil {
			log.Println(err)
			return nil
		}
	}

	return botType
}

type DummyBot map[string]any

func GetSacrificialBot() DummyBot {
	return sacrificialBot
}

func (i *DummyBot) Clone() DummyBot {
	clone := new(DummyBot)

	data, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &clone); err != nil {
		log.Fatal(err)
	}

	return *clone
}

func GetBots() *Bots {
	return &bots
}

func GetBotByName(name string) (*BotType, error) {
	botType, ok := bots.BotTypes[name]
	if !ok {
		return nil, fmt.Errorf(botNotExist, name)
	}
	return botType, nil
}

func GetBotTypeDifficultyByName(name string, diff string) (any, error) {
	botType, err := GetBotByName(name)
	if err != nil {
		return nil, err
	}

	difficulty, ok := botType.Difficulties[diff]
	if !ok {
		return nil, fmt.Errorf(difficultyNotExist, diff, name)
	}

	return difficulty, nil
}

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
	Difficulties map[string]map[string]any `json:"difficulties,omitempty"`
	Health       map[string]any            `json:"health,omitempty"`
	Loadout      *BotLoadout               `json:"loadout,omitempty"`
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
