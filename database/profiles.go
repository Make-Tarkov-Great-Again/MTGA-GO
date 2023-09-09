package database

import (
	"MT-GO/tools"
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"
)

const profilesPath string = "user/profiles/"

var profiles = map[string]*Profile{}

// #region Profile getters

func GetProfiles() map[string]*Profile {
	return profiles
}

func GetProfileByUID(uid string) *Profile {
	if profile, ok := profiles[uid]; ok {
		return profile
	}

	fmt.Println("No profile with UID ", uid, ", you stupid motherfucker")
	return nil
}

func GetStorageByUID(uid string) *Storage {
	if profile, ok := profiles[uid]; ok {
		return profile.Storage
	}

	fmt.Println("Profile with UID ", uid, " does not have a storage")
	return nil
}

// #endregion

// #region Profile setters

func setProfiles() map[string]*Profile {
	users, err := tools.GetDirectoriesFrom(profilesPath)
	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		return profiles
	}
	for _, user := range users {
		profile := &Profile{}
		userPath := filepath.Join(profilesPath, user)

		path := filepath.Join(userPath, "account.json")
		if tools.FileExist(path) {
			profile.Account = setAccount(path)
		}

		path = filepath.Join(userPath, "character.json")
		if tools.FileExist(path) {
			profile.Character = setCharacter(path)
			if profile.Character.Info.Nickname != "" {
				Nicknames[profile.Character.Info.Nickname] = struct{}{}
			}
		}

		path = filepath.Join(userPath, "storage.json")
		if tools.FileExist(path) {
			profile.Storage = setStorage(path)
		}

		path = filepath.Join(userPath, "dialogue.json")
		if tools.FileExist(path) {
			profile.Dialogue = setDialogue(path)
		}

		profile.Cache = setCache()
		profiles[user] = profile
	}
	return profiles
}

func setAccount(path string) *Account {
	output := new(Account)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		panic(err)
	}

	return output
}

func setCharacter(path string) *Character {
	output := &Character{}

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		panic(err)
	}

	return output
}

func setCache() *Cache {
	cache := &Cache{
		Quests: QuestCache{
			Index:  map[string]int8{},
			Quests: map[string]CharacterQuest{},
		},
		Traders: TraderCache{
			Index:         map[string]*AssortIndex{},
			Assorts:       map[string]*Assort{},
			LoyaltyLevels: map[string]int8{},
		},
	}
	return cache
}

func setStorage(path string) *Storage {
	output := new(Storage)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		panic(err)
	}

	return output
}

func setDialogue(path string) *Dialogue {
	output := make(Dialogue)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, &output)
	if err != nil {
		panic(err)
	}

	return &output
}

// #endregion

// #region Profile save

func (profile *Profile) SaveProfile() {
	sessionID := profile.Account.UID
	profileDirPath := filepath.Join(profilesPath, sessionID)
	if !tools.FileExist(profileDirPath) {
		err := os.Mkdir(profileDirPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	profile.Account.SaveAccount()
	profile.Character.SaveCharacter(sessionID)
	profile.Dialogue.SaveDialogue(sessionID)
	saveStorage(sessionID, *profile.Storage)
	fmt.Println()
	fmt.Println("Profile saved")
}

func saveStorage(sessionID string, storage Storage) {
	storageFilePath := filepath.Join(profilesPath, sessionID, "storage.json")

	err := tools.WriteToFile(storageFilePath, storage)
	if err != nil {
		panic(err)
	}
	fmt.Println("Storage saved")
}

// #endregion

type Profile struct {
	Account   *Account
	Character *Character
	Storage   *Storage
	Dialogue  *Dialogue
	Cache     *Cache
}

type Dialogue map[string]*Dialog

var Nicknames = make(map[string]struct{})

type Storage struct {
	//ID        string                 `json:"_id"`
	Suites    []string      `json:"suites"`
	Builds    Builds        `json:"builds"`
	Insurance []interface{} `json:"insurance"`
	Mailbox   []interface{} `json:"mailbox"`
}

type Builds struct {
	EquipmentBuilds []EquipmentBuild `json:"equipmentBuilds"`
	WeaponBuilds    []interface{}    `json:"weaponBuilds"`
}

type EquipmentBuild struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Root      string        `json:"root"`
	Items     []interface{} `json:"items"`
	Type      string        `json:"type"`
	FastPanel []interface{} `json:"fastPanel"`
	BuildType string        `json:"buildType"`
}

// #endregion
