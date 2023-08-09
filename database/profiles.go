package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"path/filepath"
	"strings"
)

const profilesPath string = "user/profiles/"

var profiles = make(map[string]*structs.Profile)

func GetProfiles() map[string]*structs.Profile {
	return profiles
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

		jsonData, err := json.Marshal(dynamic) //gos syntax is fucking pog.
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
