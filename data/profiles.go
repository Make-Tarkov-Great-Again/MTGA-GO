package data

import (
	"fmt"
	"github.com/alphadose/haxmap"
	"log"
	"os"
	"path/filepath"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

const profilesPath string = "user/profiles/"

// #region Profile getters

func GetProfiles() *haxmap.Map[string, *Profile] {
	return db.profile
}

const profileNotExist string = "Profile for %s does not exist"

func GetProfileByUID(uid string) (*Profile, error) {
	if profile, ok := db.profile.Get(uid); ok {
		return profile, nil
	}
	return nil, fmt.Errorf(profileNotExist, uid)
}

// #endregion

// #region Profile setters

func setProfiles() {
	users, err := tools.GetDirectoriesFrom(profilesPath)
	if err != nil {
		log.Println(err)
		return
	}

	db.profile = haxmap.New[string, *Profile]()

	if len(users) == 0 {
		return
	}
	for user := range users {
		profile := new(Profile)
		userPath := filepath.Join(profilesPath, user)
		done := make(chan struct{})

		go func() {
			path := filepath.Join(userPath, "account.json")
			if tools.FileExist(path) {
				profile.Account = setAccount(path)
			}
			done <- struct{}{}
		}()

		go func() {
			path := filepath.Join(userPath, "character.json")
			if tools.FileExist(path) {
				profile.Character = setCharacter(path)
				if profile.Character.Info.Nickname != "" {
					db.cache.nicknames.Set(profile.Character.Info.Nickname, struct{}{})
				}

				if profile.Character.Inventory.CleanInventoryOfDeletedItemMods() {
					if err := profile.Character.SaveCharacter(); err != nil {
						log.Println(err)
						return
					}
				}
			}
			done <- struct{}{}
		}()

		go func() {
			path := filepath.Join(userPath, "storage.json")
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
			done <- struct{}{}
		}()

		go func() {
			path := filepath.Join(userPath, "dialogue.json")
			if tools.FileExist(path) {
				profile.Dialogue = setDialogue(path)
			} else {
				profile.Dialogue = &Dialogue{}
			}
			done <- struct{}{}
		}()

		go func() {
			path := filepath.Join(userPath, "friends.json")
			if tools.FileExist(path) {
				profile.Friends = setFriends(path)
			} else {
				profile.Friends = &Friends{}
			}
			done <- struct{}{}
		}()

		for i := 0; i < 5; i++ {
			<-done
		}

		db.profile.Set(user, profile)
		SetProfileCache(user) //MAYBE I SET THIS SECONDARY
	}
}

func setAccount(path string) *Account {
	output := new(Account)

	data := tools.GetJSONRawMessage(path)
	if err := json.UnmarshalNoEscape(data, output); err != nil {
		log.Println(err)
		return nil
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
			log.Println(err)
			return
		}
	}
	done := make(chan struct{})
	go func() {
		if err := profile.Account.SaveAccount(); err != nil {
			log.Println(err)
			return
		}
		done <- struct{}{}
	}()
	go func() {
		if err := profile.Character.SaveCharacter(); err != nil {
			log.Println(err)
			return
		}
		done <- struct{}{}
	}()
	go func() {
		if err := profile.Dialogue.SaveDialogue(sessionID); err != nil {
			log.Println(err)
			return
		}
		done <- struct{}{}
	}()
	go func() {
		if err := profile.Storage.SaveStorage(sessionID); err != nil {
			log.Println(err)
			return
		}
		done <- struct{}{}
	}()
	go func() {
		if err := profile.Friends.SaveFriends(sessionID); err != nil {
			log.Println(err)
			return
		}
		done <- struct{}{}
	}()

	for i := 0; i < 5; i++ {
		<-done
	}

	log.Println("Profile saved")
}

// #endregion

type Profile struct {
	Account   *Account
	Character *Character[map[string]PlayerTradersInfo]
	Friends   *Friends
	Storage   *Storage
	Dialogue  *Dialogue
	Cache     *PlayerCache
}

type Dialogue map[string]*Dialog

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
	MagazineBuilds  []*struct{}       `json:"magazineBuilds"`
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
