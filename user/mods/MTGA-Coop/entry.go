package MTGACoop

import (
	"MT-GO/database"
	"MT-GO/tools"
	"MT-GO/user/mods/MTGA-Coop/mod"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var config *coopConfig

func init() {
	initializeVariables()
	//Cert shit but fuck you we ignore it for now and use my cert
	mod.AddRoutes()

	// if config.UseExternalIPFinder {
	//     config.ExternalIP = getExternalIP()
	// }
}

func Mod() {
	Load()
}

var coop = MTGACoop{}

type MTGACoop struct {
	LocationData  map[string]interface{}
	LocationData2 map[string]interface{}
	//Traders       []interface{}
}

// takes sessionID which is a string
func GetCoopMatch(sessionID string) string {
	if sessionID == "" {
		fmt.Println("YO SHIT IS BROKE.")
		return ""
	}
	// TODO: Return coop match via key of sessionID
	if true {
		return "GOD"
	} else {
		return ""
	}
} // Just so it doesn't bug us, ignore it

var serverConfig = database.GetServerConfig()

var modConfig = getModConfig()

func getModConfig() *database.ModInfo {
	filepath := filepath.Join("/user/mods/MTGA-Coop", "mod-info.json")

	data, err := tools.ReadFile(filepath)
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

var coopConfig = getCoopConfig()
var webSocketHandler struct{}

type coopConfigs struct {
	ProtocolConfiguration             string `json:"protocol"`
	ExternalIP                        string `json:"externalIP"`
	WebSocketPort                     int    `json:"webSocketPort"`
	UseExternalIPFinder               bool   `json:"useExternalIPFinder"`
	WebSocketTimeoutSeconds           int    `json:"webSocketTimeoutSeconds"`
	WebSocketTimeoutCheckStartSeconds int    `json:"webSocketTimeoutCheckStartSeconds"`
}

func getCoopConfig() *coopConfigs {
	output := new(coopConfigs)

	data, err := os.ReadFile("coopConfig.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = json.Unmarshal(data, &output) // pog
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return output
}

func initializeVariables() {
}

func Load() {}

/* ---------------------- Boring mod bindings below lol --------------------- */

type IP struct {
	Query string
}

func getExternalIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error() //lmfao im trying to figure out what os has that reads all now
	}
	defer req.Body.Close() //bam VV

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}
