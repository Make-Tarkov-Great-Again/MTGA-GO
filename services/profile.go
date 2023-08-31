package services

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"fmt"
	"os"
	"path/filepath"
)

func IsNicknameAvailable(nickname string, profiles map[string]*structs.Profile) bool {
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

	err := tools.WriteToFile(accountFilePath, account)
	if err != nil {
		panic(err)
	}
	fmt.Println("Account saved")
}

func SaveCharacter(sessionID string, character structs.PlayerTemplate) {
	characterFilePath := filepath.Join(profilesPath, sessionID, "character.json")

	err := tools.WriteToFile(characterFilePath, character)
	if err != nil {
		panic(err)
	}
	fmt.Println("Character saved")
}

func SaveStorage(sessionID string, storage structs.Storage) {
	storageFilePath := filepath.Join(profilesPath, sessionID, "storage.json")

	err := tools.WriteToFile(storageFilePath, storage)
	if err != nil {
		panic(err)
	}
	fmt.Println("Storage saved")
}

func SaveDialogue(sessionID string, dialogue map[string]interface{}) {
	dialogueFilePath := filepath.Join(profilesPath, sessionID, "dialogue.json")

	err := tools.WriteToFile(dialogueFilePath, dialogue)
	if err != nil {
		panic(err)
	}
	fmt.Println("Dialogue saved")
}
