package database

import (
	"fmt"
	"log"
	"path/filepath"

	"MT-GO/tools"
)

func (storage *Storage) SaveStorage(sessionID string) {
	storageFilePath := filepath.Join(profilesPath, sessionID, "storage.json")

	err := tools.WriteToFile(storageFilePath, storage)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Storage saved")
}
