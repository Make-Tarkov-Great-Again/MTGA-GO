package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

const profilesPath string = "user/profiles/"

var profiles = make(map[string]*Profile)

// #region Profile getters

func CreateProfileChangesEvent(character *Character) *ProfileChangesEvent {
	output := &ProfileChangesEvent{
		Warnings:       []*Warning{},
		ProfileChanges: make(map[string]*ProfileChanges),
	}
	output.ProfileChanges[character.ID] = &ProfileChanges{
		ID:              character.ID,
		Experience:      character.Info.Experience,
		Quests:          make([]interface{}, 0),
		RagfairOffers:   make([]interface{}, 0),
		WeaponBuilds:    make([]interface{}, 0),
		EquipmentBuilds: make([]interface{}, 0),
		Items:           ItemChanges{},
		Improvements:    make(map[string]interface{}),
		Skills:          character.Skills,
		Health:          character.Health,
		TraderRelations: make(map[string]PlayerTradersInfo),
		QuestsStatus:    make([]CharacterQuest, 0),
	}

	return output
}

func GetProfiles() map[string]*Profile {
	return profiles
}

func GetProfileByUID(uid string) *Profile {
	if profile, ok := profiles[uid]; ok {
		return profile
	}

	fmt.Println("No profile with UID ", uid, ".")
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

func SetProfiles() map[string]*Profile {
	users, err := tools.GetDirectoriesFrom(profilesPath)
	if err != nil {
		log.Fatalln(err)
	}

	if len(users) == 0 {
		return profiles
	}
	for user := range users {
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
		} else {
			profile.Character = &Character{}
		}

		path = filepath.Join(userPath, "storage.json")
		if tools.FileExist(path) {
			profile.Storage = setStorage(path)
		} else {
			profile.Storage = &Storage{
				Suites: make([]string, 0),
				Builds: Builds{
					EquipmentBuilds: make([]*EquipmentBuild, 0),
					WeaponBuilds:    make([]*WeaponBuild, 0),
				},
				Insurance: make([]interface{}, 0),
				Mailbox:   make([]*Notification, 0),
			}
		}

		path = filepath.Join(userPath, "dialogue.json")
		if tools.FileExist(path) {
			profile.Dialogue = setDialogue(path)
		} else {
			profile.Dialogue = &Dialogue{}
		}

		profile.Cache = profile.SetCache()

		profiles[user] = profile
	}
	return profiles
}

func setAccount(path string) *Account {
	output := new(Account)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		log.Fatalln(err)
	}

	return output
}

func setCharacter(path string) *Character {
	output := &Character{}

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		log.Fatalln(err)
	}

	return output
}

func setStorage(path string) *Storage {
	output := new(Storage)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		log.Fatalln(err)
	}

	return output
}

func setDialogue(path string) *Dialogue {
	output := make(Dialogue)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, &output)
	if err != nil {
		log.Fatalln(err)
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
			log.Fatalln(err)
		}
	}

	profile.Account.SaveAccount()
	profile.Character.SaveCharacter(sessionID)
	profile.Dialogue.SaveDialogue(sessionID)
	profile.Storage.SaveStorage(sessionID)

	fmt.Println()
	fmt.Println("Profile saved")
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
	Suites    []string        `json:"suites"`
	Builds    Builds          `json:"builds"`
	Insurance []interface{}   `json:"insurance"`
	Mailbox   []*Notification `json:"mailbox"`
}

type Builds struct {
	EquipmentBuilds []*EquipmentBuild `json:"equipmentBuilds"`
	WeaponBuilds    []*WeaponBuild    `json:"weaponBuilds"`
}

type WeaponBuild struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	Root  string        `json:"root"`
	Items []interface{} `json:"items"`
}

type EquipmentBuild struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Root      string        `json:"root"`
	Items     []interface{} `json:"items"`
	Type      string        `json:"type"`
	FastPanel []interface{} `json:"fastPanel"`
	BuildType int8          `json:"buildType"`
}

// #endregion

// #region Profile Change struct

type ProfileChangesEvent struct {
	Warnings       []*Warning                 `json:"warnings"`
	ProfileChanges map[string]*ProfileChanges `json:"profileChanges"`
}

type Warning struct {
	Index  int         `json:"index"`
	Errmsg string      `json:"errmsg"`
	Code   string      `json:"code,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type ItemChanges struct {
	New    []*InventoryItem `json:"new,omitempty"`
	Change []*InventoryItem `json:"change,omitempty"`
	Del    []*InventoryItem `json:"del,omitempty"`
}

type ItemLocation struct {
	X          float32 `json:"x"`
	Y          float32 `json:"y"`
	Z          float32 `json:"z"`
	IsSearched bool    `json:"isSearched,omitempty"`
}

type ProfileChanges struct {
	ID                    string                       `json:"_id"`
	Experience            int32                        `json:"experience"`
	Quests                []interface{}                `json:"quests"`
	QuestsStatus          []CharacterQuest             `json:"questsStatus"`
	RagfairOffers         []interface{}                `json:"ragFairOffers"`
	WeaponBuilds          []interface{}                `json:"weaponBuilds"`
	EquipmentBuilds       []interface{}                `json:"equipmentBuilds"`
	Items                 ItemChanges                  `json:"items"`
	Production            *map[string]interface{}      `json:"production"`
	Improvements          map[string]interface{}       `json:"improvements"`
	Skills                PlayerSkills                 `json:"skills"`
	Health                HealthInfo                   `json:"health"`
	TraderRelations       map[string]PlayerTradersInfo `json:"traderRelations"`
	RepeatableQuests      *[]interface{}               `json:"repeatableQuests,omitempty"`
	RecipeUnlocked        *map[string]bool             `json:"recipeUnlocked,omitempty"`
	ChangedHideoutStashes *map[string]interface{}      `json:"changedHideoutStashes,omitempty"`
}

// #endregion
