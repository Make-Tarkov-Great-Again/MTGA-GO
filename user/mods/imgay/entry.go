package imgay

import (
	"MT-GO/database"
	"MT-GO/tools"
	"MT-GO/user/mods/imgay/custom/items"
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

		database.SortAndQueueCustomItems(modConfig.PackageName, items.Test9)
		database.SortAndQueueCustomItems(modConfig.PackageName, items.FAMAS)

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
