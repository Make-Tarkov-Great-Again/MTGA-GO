// Pog :)
package EscapeFromHell

import (
	"MT-GO/database"
	"MT-GO/tools"
	items "MT-GO/user/mods/EFHDev/mod"
	"fmt"
	"github.com/goccy/go-json"
	"path/filepath"
	"runtime"
)

var modInfo = SetModConfig()

func Mod() {
	fmt.Println("Loading Escape from Hell....")

	items.Modify(&modInfo)

	fmt.Println("Loaded mod Escape From hell lol.")
}

/* ---------------------- Boring mod bindings below lol --------------------- */
//TODO: Save directory path to reduce imports

func SetModConfig() database.ModInfo {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic(fmt.Errorf("failed to get the current file's path"))
	}
	readFile, err := tools.ReadFile(filepath.Join(filepath.Dir(filename), "mod-info.json"))
	if err != nil {
		panic(err)
	}

	config := new(database.ModInfo)
	err = json.Unmarshal(readFile, &config)
	if err != nil {
		panic(err)
	}
	return *config
}
