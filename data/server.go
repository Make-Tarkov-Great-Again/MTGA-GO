package data

import (
	"fmt"
	"net"
)

const (
	WSSTemplate   = "wss://%s"
	HTTPSTemplate = "https://%s"
	WSTemplate    = "ws://%s"
	HTTPTemplate  = "http://%s"
)

func setServerConfig() {
	coreServerData := &serverData{
		MainIPandPort:      net.JoinHostPort(db.core.ServerConfig.IP, db.core.ServerConfig.Ports.Main),
		MessagingIPandPort: net.JoinHostPort(db.core.ServerConfig.IP, db.core.ServerConfig.Ports.Messaging),
		TradingIPandPort:   net.JoinHostPort(db.core.ServerConfig.IP, db.core.ServerConfig.Ports.Trading),
		RagFairIPandPort:   net.JoinHostPort(db.core.ServerConfig.IP, db.core.ServerConfig.Ports.Flea),
		LobbyIPandPort:     net.JoinHostPort(db.core.ServerConfig.IP, db.core.ServerConfig.Ports.Lobby),
	}

	if db.core.ServerConfig.Secure {
		coreServerData.HTTPS = &serverDataHTTPS{
			HTTPSAddress: fmt.Sprintf(HTTPSTemplate, coreServerData.MainIPandPort),
			WSSAddress:   fmt.Sprintf(WSSTemplate, coreServerData.LobbyIPandPort),
		}

		coreServerData.MainAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.MainIPandPort)
		coreServerData.MessageAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.MessagingIPandPort)
		coreServerData.TradingAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.TradingIPandPort)
		coreServerData.RagFairAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.RagFairIPandPort)
		coreServerData.LobbyAddress = fmt.Sprintf("wss://%s/sws", coreServerData.LobbyIPandPort)
	} else {
		coreServerData.HTTP = &serverDataHTTP{
			HTTPAddress: fmt.Sprintf(HTTPTemplate, coreServerData.LobbyIPandPort),
			WSAddress:   fmt.Sprintf(WSTemplate, coreServerData.LobbyIPandPort),
		}

		coreServerData.MainAddress = fmt.Sprintf(HTTPTemplate, coreServerData.MainIPandPort)
		coreServerData.MessageAddress = fmt.Sprintf(HTTPTemplate, coreServerData.MessagingIPandPort)
		coreServerData.TradingAddress = fmt.Sprintf(HTTPTemplate, coreServerData.TradingIPandPort)
		coreServerData.RagFairAddress = fmt.Sprintf(HTTPTemplate, coreServerData.RagFairIPandPort)
		coreServerData.LobbyAddress = fmt.Sprintf("ws://%s/sws", coreServerData.LobbyIPandPort)
	}
	db.cache.server.core = coreServerData
}

func GetMainAddress() string {
	return db.cache.server.core.MainAddress
}

func GetTradingAddress() string {
	return db.cache.server.core.TradingAddress
}

func GetMessageAddress() string {
	return db.cache.server.core.MessageAddress
}

func GetRagFairAddress() string {
	return db.cache.server.core.RagFairAddress
}

func GetLobbyAddress() string {
	return db.cache.server.core.LobbyAddress
}
func GetWebSocketAddress() string {
	if db.cache.server.core.HTTPS != nil {
		return db.cache.server.core.HTTPS.WSSAddress
	}
	return db.cache.server.core.HTTP.WSAddress
}

func GetMainIPandPort() string {
	return db.cache.server.core.MainIPandPort
}

func GetTradingIPandPort() string {
	return db.cache.server.core.TradingIPandPort
}

func AddToItemPresets(key string, value globalItemPreset) {
	db.core.Globals.ItemPresets[key] = value
}

func GetMessagingIPandPort() string {
	return db.cache.server.core.MessagingIPandPort
}

func GetLobbyIPandPort() string {
	return db.cache.server.core.LobbyIPandPort
}

func GetRagFairIPandPort() string {
	return db.cache.server.core.RagFairIPandPort
}

type serverData struct {
	HTTPS *serverDataHTTPS
	HTTP  *serverDataHTTP

	MainIPandPort string
	MainAddress   string

	MessagingIPandPort string
	MessageAddress     string

	TradingIPandPort string
	TradingAddress   string

	RagFairIPandPort string
	RagFairAddress   string

	LobbyIPandPort string
	LobbyAddress   string
}

type serverDataHTTPS struct {
	HTTPSAddress string
	WSSAddress   string
}

type serverDataHTTP struct {
	HTTPAddress string
	WSAddress   string
}

type ServerConfig struct {
	IP                 string      `json:"ip"`
	Hostname           string      `json:"hostname"`
	Name               string      `json:"name"`
	BrandName          string      `json:"brandName"`
	Version            string      `json:"version"`
	Secure             bool        `json:"secure"`
	DownloadImageFiles bool        `json:"downloadImageFiles"`
	Ports              ServerPorts `json:"ports"`
}

type ServerPorts struct {
	Main      string `json:"Main"`
	Messaging string `json:"Messaging"`
	Trading   string `json:"Trading"`
	Flea      string `json:"Flea"`
	Lobby     string `json:"Lobby"`
}
