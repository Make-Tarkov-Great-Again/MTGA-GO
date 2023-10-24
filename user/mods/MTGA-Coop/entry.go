package MTGACoop

import (
	"MT-GO/database"
	"MT-GO/tools"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
		serverConfig.IP = coopConfig.ExternalIP
	}

	switch coopConfig.ProtocolConfiguration {
	case "http":
		serverConfig.Secure = false
		break
	case "https":
		serverConfig.Secure = true
		break
	default:
		log.Fatalln(coopConfig.ProtocolConfiguration, "is not a proper protocol, adjust in your coopConfig.json")
		return
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
	//path := filepath.Join("/user/mods/MTGA-Coop", "mod-info.json")
	data, err := tools.ReadFile("mod-info.json")
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

func getExternalIP() string {
	resp, err := http.Get("https://api.ipify.org?format=text") //MY FUNCTION!!
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	return string(ip)
}
