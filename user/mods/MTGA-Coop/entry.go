package MTGACoop

import (
	"MT-GO/database"
	"MT-GO/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var coopConfig = getCoopConfig()
var serverConfig = database.GetServerConfig()
var modConfig = getModConfig()
var coop = MTGACoop{}

//var webSocketHandler struct{}

func init() {
	//mod.AddRoutes()

	if coopConfig.UseExternalIPFinder {
		coopConfig.ExternalIP = getExternalIP()
	}
}

func Mod() {
	Load()
}

type MTGACoop struct {
	LocationData  map[string]interface{}
	LocationData2 map[string]interface{}
	//Traders       []interface{}
}

func getModConfig() *database.ModInfo {
	path := filepath.Join("/user/mods/MTGA-Coop", "mod-info.json")
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

func Load() {}

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
	err = json.Unmarshal(body, &ip)
	if err != nil {
		return ""
	}

	return ip.Query
}
