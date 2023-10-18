// Pog :)
package EscapeFromSmell

import (
	items "MT-GO/user/mods/EFHLev/mod"
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
)

var passed ModInfo

func Mod() {
	fmt.Println("Loading Escape from Smell....")
	items.Modify(passed)
	items.AmmoStacks()
	defer fmt.Println("Loaded mod Escape From Smell lol.")

}

/* ---------------------- Boring mod bindings below lol --------------------- */

func GetModConfig() (*ModInfo, error) {
	data, err := config.ReadFile("mod-info.json")
	if err != nil {
		return nil, err
	}

	var modInfo ModInfo
	if err := json.Unmarshal(data, &modInfo); err != nil {
		return nil, err
	}
	passed = modInfo
	return &modInfo, nil
}

type ModInfo struct {
	NameSpace       string
	ModNameNoSpaces string
	Advanced        struct {
		CustomRoutes bool
	}
	Config map[string]interface{}
}

var Router *http.ServeMux

func GetRouter(mux *http.ServeMux) {
	config, err := GetModConfig()
	if err != nil {
		fmt.Errorf("Error setting router for %s: %s", config.ModNameNoSpaces, err)
	}
	if config.Advanced.CustomRoutes {
		Router = mux
		return
	}
}

//go:embed mod-info.json
var config embed.FS

func (m ModInfo) GetConfig() map[string]interface{} {
	return m.Config
}
