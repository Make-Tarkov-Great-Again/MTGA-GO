package database

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"MT-GO/tools"

	"github.com/goccy/go-json"
)

var locations = Locations{}

// #region Location getters

func GetLocations() *Locations {
	return &locations
}

// #endregion

// #region Location setters

func setLocationsMaster() {
	setLocations()
	setLocalLoot()
}

func setLocations() {
	raw := tools.GetJSONRawMessage(locationsFilePath)
	err := json.Unmarshal(raw, &locations)
	if err != nil {
		log.Fatalln(err)
	}
}

var localLoot = make(map[string][]interface{})

func GetLocalLootByNameAndIndex(name string, index int8) interface{} {
	location, ok := localLoot[name]
	if !ok {
		fmt.Println("Location", name, "doesn't exist in localLoot map")
		return nil
	}

	loot := location[index]
	if loot == nil {
		fmt.Println("Loot at index", index, "does not exist")
		return nil
	}

	return loot
}

func setLocalLoot() {
	files, err := tools.GetFilesFrom("/locationTest")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fileNameSplit := strings.Split(file, ".")
		name := fileNameSplit[0][:len(fileNameSplit[0])-1]

		if _, ok := localLoot[name]; !ok {
			localLoot[name] = make([]interface{}, 0, 6)
		}
		filePath := filepath.Join("/locationTest", file)

		formatt := new(interface{})
		readFile, err := tools.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(readFile, formatt)
		if err != nil {
			fmt.Println(err)
		}

		localLoot[name] = append(localLoot[name], formatt)
	}
}

// #endregion

// #region Location structs

type Locations struct {
	Locations map[string]interface{} `json:"locations"`
	Paths     []interface{}          `json:"paths"`
}

// #endregion
