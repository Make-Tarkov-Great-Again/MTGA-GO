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

var coreServerData = &serverData{
	MainIPandPort:      "",
	MainAddress:        "",
	MessagingIPandPort: "",
	MessageAddress:     "",
	TradingIPandPort:   "",
	TradingAddress:     "",
	RagFairIPandPort:   "",
	RagFairAddress:     "",
	LobbyIPandPort:     "",
	LobbyAddress:       "",
}

func SetServerConfig() {

	coreServerData.MainIPandPort = net.JoinHostPort(core.ServerConfig.IP, core.ServerConfig.Ports.Main)
	coreServerData.MessagingIPandPort = net.JoinHostPort(core.ServerConfig.IP, core.ServerConfig.Ports.Messaging)
	coreServerData.TradingIPandPort = net.JoinHostPort(core.ServerConfig.IP, core.ServerConfig.Ports.Trading)
	coreServerData.RagFairIPandPort = net.JoinHostPort(core.ServerConfig.IP, core.ServerConfig.Ports.Flea)
	coreServerData.LobbyIPandPort = net.JoinHostPort(core.ServerConfig.IP, core.ServerConfig.Ports.Lobby)

	if core.ServerConfig.Secure {
		coreServerData.HTTPS = new(serverDataHTTPS)

		coreServerData.HTTPS.HTTPSAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.MainIPandPort)
		coreServerData.HTTPS.WSSAddress = fmt.Sprintf(WSSTemplate, coreServerData.LobbyIPandPort)

		coreServerData.MainAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.MainIPandPort)
		coreServerData.MessageAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.MessagingIPandPort)
		coreServerData.TradingAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.TradingIPandPort)
		coreServerData.RagFairAddress = fmt.Sprintf(HTTPSTemplate, coreServerData.RagFairIPandPort)
		coreServerData.LobbyAddress = fmt.Sprintf("wss://%s/sws", coreServerData.LobbyIPandPort)
	} else {
		coreServerData.HTTP = new(serverDataHTTP)

		coreServerData.HTTP.HTTPAddress = fmt.Sprintf(HTTPTemplate, coreServerData.LobbyIPandPort)
		coreServerData.HTTP.WSAddress = fmt.Sprintf(WSTemplate, coreServerData.LobbyIPandPort)

		coreServerData.MainAddress = fmt.Sprintf(HTTPTemplate, coreServerData.MainIPandPort)
		coreServerData.MessageAddress = fmt.Sprintf(HTTPTemplate, coreServerData.MessagingIPandPort)
		coreServerData.TradingAddress = fmt.Sprintf(HTTPTemplate, coreServerData.TradingIPandPort)
		coreServerData.RagFairAddress = fmt.Sprintf(HTTPTemplate, coreServerData.RagFairIPandPort)
		coreServerData.LobbyAddress = fmt.Sprintf("ws://%s/sws", coreServerData.LobbyIPandPort)
	}
}

func GetMainAddress() string {
	return coreServerData.MainAddress
}

func GetTradingAddress() string {
	return coreServerData.TradingAddress
}

func GetMessageAddress() string {
	return coreServerData.MessageAddress
}

func GetRagFairAddress() string {
	return coreServerData.RagFairAddress
}

func GetLobbyAddress() string {
	return coreServerData.LobbyAddress
}
func GetWebSocketAddress() string {
	if coreServerData.HTTPS != nil {
		return coreServerData.HTTPS.WSSAddress
	}
	return coreServerData.HTTP.WSAddress
}

func GetMainIPandPort() string {
	return coreServerData.MainIPandPort
}

func GetTradingIPandPort() string {
	return coreServerData.TradingIPandPort
}

func AddToItemPresets(key string, value globalItemPreset) {
	core.Globals.ItemPresets[key] = value
}

func GetMessagingIPandPort() string {
	return coreServerData.MessagingIPandPort
}

func GetLobbyIPandPort() string {
	return coreServerData.LobbyIPandPort
}

func GetRagFairIPandPort() string {
	return coreServerData.RagFairIPandPort
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
