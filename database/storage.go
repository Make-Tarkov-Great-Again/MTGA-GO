package database

import (
	"MT-GO/tools"
	"fmt"
	"path/filepath"
)

func (storage *Storage) SaveStorage(sessionID string) {
	storageFilePath := filepath.Join(profilesPath, sessionID, "storage.json")

	err := tools.WriteToFile(storageFilePath, storage)
	if err != nil {
		panic(err)
	}
	fmt.Println("Storage saved")
}
