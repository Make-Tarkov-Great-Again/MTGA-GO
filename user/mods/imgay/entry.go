package imgay

import (
	"MT-GO/database"
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"path/filepath"
)

var mainDir string
var modConfig *database.ModInfo

func Mod(dir string) {
	mainDir = dir
	modConfig = getModConfig()

	customDirectory := filepath.Join(mainDir, "custom")
	customDirectories, err := tools.GetDirectoriesFrom(customDirectory)
	if err != nil {
		log.Println(err)
		return
	}

	if modConfig.Parameters.CustomItems {
		if _, ok := customDirectories["items"]; !ok {
			log.Println("Where the fuck is your items directory faggot??!?!?!?!?")
			return
		}

		itemFilesDir := filepath.Join(customDirectory, "items")
		itemFiles, err := tools.GetFilesFrom(itemFilesDir)
		if err != nil {
			log.Println(err)
			return
		}

		for file := range itemFiles {
			filePath := filepath.Join(itemFilesDir, file)
			customItems := make(map[string]*database.CustomItemAPI)
			if err := json.Unmarshal(tools.GetJSONRawMessage(filePath), &customItems); err != nil {
				log.Println(err)
				return
			}

			database.SortAndQueueCustomItems(modConfig.PackageName, customItems)
		}

	}

	fmt.Println("im gay")
}

func getModConfig() *database.ModInfo {
	path := filepath.Join(mainDir, "mod-info.json")
	data, err := tools.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := new(database.ModInfo)
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}
