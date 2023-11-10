package data

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"path/filepath"

	"MT-GO/tools"
)

func setStorage(path string) *Storage {
	output := new(Storage)

	data := tools.GetJSONRawMessage(path)
	err := json.Unmarshal(data, output)
	if err != nil {
		log.Println(err)
	}

	return output
}

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

func (storage *Storage) SaveStorage(sessionID string) error {
	storageFilePath := filepath.Join(profilesPath, sessionID, "storage.json")

	err := tools.WriteToFile(storageFilePath, storage)
	if err != nil {
		return fmt.Errorf(storageNotSaved, sessionID, err)
	}
	log.Println("Storage saved")
	return nil
}

const (
	storageNotSaved string = "Account for %s was not saved: %s"
	storageNotExist string = "Storage for UID %s does not exist"
)
