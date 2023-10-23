package config

import (
	"MT-GO/tools"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type CoopConfig struct {
	Protocol                          string `json:"protocol"`
	ExternalIP                        string `json:"externalIP"`
	WebSocketPort                     int    `json:"webSocketPort"`
	UseExternalIPFinder               bool   `json:"useExternalIPFinder"`
	WebSocketTimeoutSeconds           int    `json:"webSocketTimeoutSeconds"`
	WebSocketTimeoutCheckStartSeconds int    `json:"webSocketTimeoutCheckStartSeconds"`
}

var Instance *CoopConfig

func NewCoopConfig() *CoopConfig {
	coopConfig := &CoopConfig{
		Protocol:                          "http",
		ExternalIP:                        "127.0.0.1",
		WebSocketPort:                     6970,
		UseExternalIPFinder:               true,
		WebSocketTimeoutSeconds:           10,
		WebSocketTimeoutCheckStartSeconds: 120,
	}

	configFilePath := tools.GetAbsolutePathFrom(filepath.Join(filepath.Dir(os.Args[0]), "config"))
	if !tools.FileExist(configFilePath) {
		if err := tools.CreateDirectory(configFilePath); err != nil {
			log.Fatal(err)
		}

		fmt.Println("[COOP] Loading config")
		coopConfigFilePath := filepath.Join(configFilePath, "coopConfig.json")

		if !tools.FileExist(coopConfigFilePath) {
			fmt.Println("[COOP] No config found. Defaulting...")
			fmt.Println("[COOP] External IP finder is active. Your config won't affect it.")
			coopcfgString, _ := json.MarshalIndent(coopConfig, "", "    ")
			if err := tools.WriteToFile(coopConfigFilePath, coopcfgString); err != nil {
				log.Fatal(err)
			}
		} else {
			configData, err := tools.ReadFile(coopConfigFilePath)
			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(configData, coopConfig)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("[COOP] Config loaded")
			if coopConfig.UseExternalIPFinder {
				fmt.Println("[COOP] External IP finder is active. Your config won't affect it.")
			}
		}
	}

	Instance = coopConfig
	return coopConfig
}

// GetConfig returns the value of the specified configuration key.
func GetConfig(key string) interface{} {
	switch key {
	case "Protocol":
		return Instance.Protocol
	case "ExternalIP":
		return Instance.ExternalIP
	case "WebSocketPort":
		return Instance.WebSocketPort
	case "UseExternalIPFinder":
		return Instance.UseExternalIPFinder
	case "WebSocketTimeoutSeconds":
		return Instance.WebSocketTimeoutSeconds
	case "WebSocketTimeoutCheckStartSeconds":
		return Instance.WebSocketTimeoutCheckStartSeconds
	default:
		return nil
	}
}
