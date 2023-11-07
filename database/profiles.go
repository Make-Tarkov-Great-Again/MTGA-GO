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
		Quests:          make([]any, 0),
		RagfairOffers:   make([]any, 0),
		WeaponBuilds:    make([]any, 0),
		EquipmentBuilds: make([]any, 0),
		Items:           ItemChanges{},
		Improvements:    make(map[string]any),
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

const profileNotExist string = "Profile for %s does not exist"

func GetProfileByUID(uid string) (*Profile, error) {
	if profile, ok := profiles[uid]; ok {
		return profile, nil
	}
	return nil, fmt.Errorf(profileNotExist, uid)
}

const storageNotExist string = "Storage for UID %s does not exist"

func GetStorageByUID(uid string) (*Storage, error) {
	profile, err := GetProfileByUID(uid)
	if err != nil {
		return nil, err
	}

	if profile.Storage != nil {
		return profile.Storage, nil
	}

	return nil, fmt.Errorf(storageNotExist, uid)
}

// #endregion

// #region Profile setters

func SetProfiles() {
	users, err := tools.GetDirectoriesFrom(profilesPath)
	if err != nil {
		log.Fatalln(err)
	}

	if len(users) == 0 {
		return
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
				Nicknames[profile.Character.Info.Nickname] = nil
			}

			if profile.Character.Inventory.CleanInventoryOfDeletedItemMods() {
				profile.Character.SaveCharacter()
			}

		} else {
			profile.Character = &Character{
				ID: profile.Account.UID,
			}
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
				Insurance: make([]any, 0),
				Mailbox:   make([]*Notification, 0),
			}
		}

		path = filepath.Join(userPath, "dialogue.json")
		if tools.FileExist(path) {
			profile.Dialogue = setDialogue(path)
		} else {
			profile.Dialogue = &Dialogue{}
		}

		path = filepath.Join(userPath, "friends.json")
		if tools.FileExist(path) {
			profile.Friends = setFriends(path)
		} else {
			profile.Friends = &Friends{}
		}

		profile.Cache = profile.SetCache()

		profiles[user] = profile
	}
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

func setFriends(path string) *Friends {
	output := new(Friends)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, &output)
	if err != nil {
		log.Fatalln(err)
	}

	return output
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
	profile.Character.SaveCharacter()
	profile.Dialogue.SaveDialogue(sessionID)
	profile.Storage.SaveStorage(sessionID)
	profile.Friends.SaveFriends(sessionID)

	fmt.Println()
	fmt.Println("Profile saved")
}

// #endregion

type Profile struct {
	Account   *Account
	Character *Character
	Friends   *Friends
	Storage   *Storage
	Dialogue  *Dialogue
	Cache     *Cache
}

type Dialogue map[string]*Dialog

var Nicknames = make(map[string]*struct{})

type Friends struct {
	Friends             []FriendRequest `json:"Friends"`
	Ignore              []string        `json:"Ignore"`
	InIgnoreList        []string        `json:"InIgnoreList"`
	Matching            Matching        `json:"Matching"`
	FriendRequestInbox  []any           `json:"friendRequestInbox"`
	FriendRequestOutbox []any           `json:"friendRequestOutbox"`
}

type Matching struct {
	LookingForGroup bool `json:"LookingForGroup"`
}

type FriendRequest struct {
	ID      string               `json:"_id"`
	From    string               `json:"from"`
	To      string               `json:"to"`
	Date    int32                `json:"date"`
	Profile FriendRequestProfile `json:"profile"`
}

type FriendRequestProfile struct {
	ID   int32
	Info struct {
		Nickname       string         `json:"Nickname"`
		Side           string         `json:"Side"`
		Level          int8           `json:"Level"`
		MemberCategory MemberCategory `json:"MemberCategory"`
	}
}

type Storage struct {
	//ID        string                 `json:"_id"`
	Suites    []string        `json:"suites"`
	Builds    Builds          `json:"builds"`
	Insurance []any           `json:"insurance"`
	Mailbox   []*Notification `json:"mailbox"`
}

type Builds struct {
	EquipmentBuilds []*EquipmentBuild `json:"equipmentBuilds"`
	WeaponBuilds    []*WeaponBuild    `json:"weaponBuilds"`
}

type WeaponBuild struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Root  string `json:"root"`
	Items []any  `json:"items"`
}

type EquipmentBuild struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Root      string `json:"root"`
	Items     []any  `json:"items"`
	Type      string `json:"type"`
	FastPanel []any  `json:"fastPanel"`
	BuildType int8   `json:"buildType"`
}

// #endregion

// #region Profile Change struct

type ProfileChangesEvent struct {
	Warnings       []*Warning                 `json:"warnings"`
	ProfileChanges map[string]*ProfileChanges `json:"profileChanges"`
}

type Warning struct {
	Index  int    `json:"index"`
	Errmsg string `json:"errmsg"`
	Code   string `json:"code,omitempty"`
	Data   any    `json:"data,omitempty"`
}

type ItemChanges struct {
	New    []InventoryItem `json:"new,omitempty"`
	Change []InventoryItem `json:"change,omitempty"`
	Del    []InventoryItem `json:"del,omitempty"`
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
	Quests                []any                        `json:"quests"`
	QuestsStatus          []CharacterQuest             `json:"questsStatus"`
	RagfairOffers         []any                        `json:"ragFairOffers"`
	WeaponBuilds          []any                        `json:"weaponBuilds"`
	EquipmentBuilds       []any                        `json:"equipmentBuilds"`
	Items                 ItemChanges                  `json:"items"`
	Production            *map[string]any              `json:"production"`
	Improvements          map[string]any               `json:"improvements"`
	Skills                PlayerSkills                 `json:"skills"`
	Health                HealthInfo                   `json:"health"`
	TraderRelations       map[string]PlayerTradersInfo `json:"traderRelations"`
	RepeatableQuests      *[]any                       `json:"repeatableQuests,omitempty"`
	RecipeUnlocked        *map[string]bool             `json:"recipeUnlocked,omitempty"`
	ChangedHideoutStashes *map[string]any              `json:"changedHideoutStashes,omitempty"`
}

// #endregion
