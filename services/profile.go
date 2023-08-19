package services

import (
	"MT-GO/database"
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func IsNicknameAvailable(nickname string) bool {
	profiles := database.GetProfiles()

	for _, profile := range profiles {
		if profile.Character == nil || profile.Character.Info.Nickname == "" {
			continue
		}

		Nickname := profile.Character.Info.Nickname
		if Nickname == nickname {
			return false
		}

	}
	return true
}

const profilesPath string = "user/profiles/"

func SaveProfile(profile *structs.Profile) {
	sessionID := profile.Account.UID
	profileDirPath := filepath.Join(profilesPath, sessionID)
	if !tools.FileExist(profileDirPath) {
		os.Mkdir(profileDirPath, 0755)
	}

	SaveAccount(*profile.Account)
	SaveCharacter(sessionID, *profile.Character)
	SaveDialogue(sessionID, profile.Dialogue)
	SaveStorage(sessionID, *profile.Storage)
	fmt.Println()
	fmt.Println("Profile saved")
}

func SaveAccount(account structs.Account) {
	accountFilePath := filepath.Join(profilesPath, account.UID, "account.json")
	data, err := json.MarshalIndent(account, "", "    ")
	if err != nil {
		panic(err)
	}

	err = tools.WriteToFile(accountFilePath, string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("Account saved")
}

func SaveCharacter(sessionID string, character structs.PlayerTemplate) {
	characterFilePath := filepath.Join(profilesPath, sessionID, "character.json")
	data, err := json.MarshalIndent(character, "", "    ")
	if err != nil {
		panic(err)
	}
	err = tools.WriteToFile(characterFilePath, string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("Character saved")
}

func SaveStorage(sessionID string, storage structs.Storage) {
	storageFilePath := filepath.Join(profilesPath, sessionID, "storage.json")
	data, err := json.MarshalIndent(storage, "", "    ")
	if err != nil {
		panic(err)
	}

	err = tools.WriteToFile(storageFilePath, string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("Storage saved")
}

func SaveDialogue(sessionID string, dialogue map[string]interface{}) {
	dialogueFilePath := filepath.Join(profilesPath, sessionID, "dialogue.json")
	data, err := json.MarshalIndent(dialogue, "", "    ")
	if err != nil {
		panic(err)
	}

	err = tools.WriteToFile(dialogueFilePath, string(data))
	if err != nil {
		panic(err)
	}
	fmt.Println("Dialogue saved")
}
