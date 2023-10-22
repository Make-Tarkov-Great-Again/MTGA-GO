package MTGACoop

import (
	"MT-GO/database"
	"MT-GO/tools"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/goccy/go-json"
)

var modInfo = SetModConfig()

func Mod() {
	Load()
}

/* ---------------------- Boring mod bindings below lol --------------------- */
//TODO: Save directory path to reduce imports

func SetModConfig() database.ModInfo {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(fmt.Errorf("failed to get the current file's path"))
	}

	directory := filepath.Dir(filename)
	readFile, err := tools.ReadFile(filepath.Join(directory, "mod-info.json"))
	if err != nil {
		panic(err)
	}

	config := new(database.ModInfo)
	err = json.Unmarshal(readFile, &config)
	if err != nil {
		panic(err)
	}

	config.Dir = directory
	return *config
}

func Load() {

}
