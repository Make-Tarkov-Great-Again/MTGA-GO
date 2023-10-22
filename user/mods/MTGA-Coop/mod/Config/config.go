package config

import (
	"encoding/json"
	"fmt"
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

	configFilePath := filepath.Join(filepath.Dir(os.Args[0]), "config")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		os.MkdirAll(configFilePath, os.ModePerm)
		fmt.Println("[COOP] Loading config")
		coopConfigFilePath := filepath.Join(configFilePath, "coopConfig.json")
		if _, err := os.Stat(coopConfigFilePath); os.IsNotExist(err) {
			fmt.Println("[COOP] No config found. Defaulting...")
			fmt.Println("[COOP] External IP finder is active. Your config won't affect it.")
			coopcfgString, _ := json.MarshalIndent(coopConfig, "", "    ")
			os.WriteFile(coopConfigFilePath, coopcfgString, os.ModePerm)
		} else {
			configData, _ := os.ReadFile(coopConfigFilePath)
			json.Unmarshal(configData, coopConfig)
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
