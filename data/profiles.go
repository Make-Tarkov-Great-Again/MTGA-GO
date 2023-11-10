package data

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

func GetStorageByID(uid string) (*Storage, error) {
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
				Builds: &Builds{
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

		CreateCacheByID(user)
		profiles[user] = profile
		if cache, err := GetCacheByID(user); err == nil {
			cache.SetCache(profiles[user].Character)
		}
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

	log.Println()
	log.Println("Profile saved")
}

// #endregion

type Profile struct {
	Account   *Account
	Character *Character
	Friends   *Friends
	Storage   *Storage
	Dialogue  *Dialogue
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
	Builds    *Builds         `json:"builds"`
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

// #endregion
