package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
)

const profilesPath string = "user/profiles/"

var profiles = map[string]*structs.Profile{}

func GetProfiles() map[string]*structs.Profile {
	return profiles
}

func GetProfileByUID(uid string) *structs.Profile {
	if profile, ok := profiles[uid]; ok {
		return profile
	}

	fmt.Println("No profile with UID ", uid, ", you stupid motherfucker")
	return nil
}

func GetAccountByUID(uid string) *structs.Account {
	profile := GetProfileByUID(uid)
	if profile.Account != nil {
		return profile.Account
	}

	fmt.Println("Profile with UID ", uid, " does not have an account, how the fuck did you get here????!?!?!?!?!?")
	return nil
}

func GetCharacterByUID(uid string) *structs.PlayerTemplate {
	if profile, ok := profiles[uid]; ok {
		return profile.Character
	}

	fmt.Println("Profile with UID ", uid, " does not have a character")
	return nil
}

func GetStorageByUID(uid string) *structs.Storage {
	if profile, ok := profiles[uid]; ok {
		return profile.Storage
	}

	fmt.Println("Profile with UID ", uid, " does not have a storage")
	return nil
}

func GetDialogueByUID(uid string) *map[string]interface{} {
	if profile, ok := profiles[uid]; ok {
		return &profile.Dialogue
	}

	fmt.Println("Profile with UID ", uid, " does not have dialogue")
	return nil
}

func setProfiles() map[string]*structs.Profile {
	users, err := tools.GetDirectoriesFrom(profilesPath)
	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		return profiles
	}
	for _, user := range users {
		profile := &structs.Profile{}
		userPath := filepath.Join(profilesPath, user)
		files, err := tools.GetFilesFrom(userPath)
		if err != nil {
			panic(err)
		}

		dynamic := make(map[string]json.RawMessage)
		for _, file := range files {
			name := strings.TrimSuffix(file, ".json")
			data := tools.GetJSONRawMessage(filepath.Join(userPath, file))
			dynamic[name] = data
		}

		jsonData, err := json.Marshal(dynamic)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(jsonData, profile)
		if err != nil {
			panic(err)
		}
		profiles[user] = profile
	}
	return profiles
}
