package data

import (
	"log"
	"path/filepath"

	"MT-GO/tools"
)

func (storage *Storage) SaveStorage(sessionID string) {
	storageFilePath := filepath.Join(profilesPath, sessionID, "storage.json")

	err := tools.WriteToFile(storageFilePath, storage)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Storage saved")
}
