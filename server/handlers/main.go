package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"MT-GO/structs"
	"MT-GO/tools"
	"fmt"
	"net"
	"net/http"
	"strings"
)

const ROUTE_NOT_IMPLEMENTED = "Route is not implemented yet, using empty values instead"

// GetBundleList returns a list of custom bundles to the client
func GetBundleList(w http.ResponseWriter, _ *http.Request) {
	fmt.Println(ROUTE_NOT_IMPLEMENTED)
	services.ZlibJSONReply(w, []string{})
}

func ShowPersonKilledMessage(w http.ResponseWriter, _ *http.Request) {
	services.ZlibJSONReply(w, "true")
}

func ClientGameStart(w http.ResponseWriter, _ *http.Request) {
	data := map[string]interface{}{
		"utc_time": tools.GetCurrentTimeInSeconds(),
	}

	start := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, start)
}

func ClientMenuLocale(w http.ResponseWriter, r *http.Request) {
	locale := strings.TrimPrefix(r.URL.Path, "/client/menu/locale/")

	menu := struct {
		Data   *structs.LocaleMenu
		Err    int
		Errmsg interface{}
	}{
		Data: database.GetLocalesMenuByName(locale),
	}

	services.ZlibJSONReply(w, menu)
}

func ClientGameVersionValidate(w http.ResponseWriter, _ *http.Request) {
	verValidate := struct {
		Data   interface{}
		Err    int
		Errmsg interface{}
	}{
		Err: 0,
	}

	services.ZlibJSONReply(w, verValidate)
}

func ClientLanguages(w http.ResponseWriter, r *http.Request) {
	languages := struct {
		Data   map[string]string
		Err    int
		Errmsg interface{}
	}{
		Data: database.GetLanguages(),
	}

	services.ZlibJSONReply(w, languages)
}

func ClientGameConfig(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	account := database.GetAccountByUID(sessionID)

	serverConfig := database.GetCore().ServerConfig
	//Main := database.GetCore().ServerConfig.Ports.Main
	IP := serverConfig.IP

	mainAddress := net.JoinHostPort(IP, serverConfig.Ports.Main)
	//messageAddress := net.JoinHostPort(IP, serverConfig.Ports.Messaging)
	//tradingAddress := net.JoinHostPort(IP, serverConfig.Ports.Trading)
	//ragFairAddress := net.JoinHostPort(IP, serverConfig.Ports.Flea)
	//lobbyAddress := net.JoinHostPort(IP, serverConfig.Ports.Lobby)

	type Backend struct {
		Lobby     string `json:"Lobby"`
		Trading   string `json:"Trading"`
		Messaging string `json:"Messaging"`
		Main      string `json:"Main"`
		RagFair   string `json:"RagFair"`
	}

	type GameConfig struct {
		Aid               string            `json:"aid"`
		Lang              string            `json:"lang"`
		Languages         map[string]string `json:"languages"`
		NdaFree           bool              `json:"ndaFree"`
		Taxonomy          int               `json:"taxonomy"`
		ActiveProfileID   string            `json:"activeProfileId"`
		Backend           Backend           `json:"backend"`
		UtcTime           string            `json:"utc_time"`
		TotalInGame       int               `json:"totalInGame"`
		ReportAvailable   bool              `json:"reportAvailable"`
		TwitchEventMember bool              `json:"twitchEventMember"`
	}

	config := &GameConfig{
		Aid:             account.UID,
		Lang:            account.Lang,
		Languages:       database.GetLanguages(),
		NdaFree:         false,
		Taxonomy:        6,
		ActiveProfileID: account.UID,
		Backend: Backend{
			Lobby:     mainAddress, //lobbyAddress,
			Trading:   mainAddress, //tradingAddress,
			Messaging: mainAddress, //messageAddress,
			Main:      mainAddress,
			RagFair:   mainAddress, //ragFairAddress,
		},
		UtcTime:           tools.GetCurrentTimeInSeconds(),
		TotalInGame:       1,
		ReportAvailable:   true,
		TwitchEventMember: false,
	}

	gameConfig := struct {
		Err    int
		Errmsg interface{}
		Data   *GameConfig
	}{
		Data: config,
	}

	fmt.Println("Don't forget to create multiple servers later on you dumbshit!!!!!!!!!!!!")
	services.ZlibJSONReply(w, gameConfig)
}

func ClientItems(w http.ResponseWriter, r *http.Request) {
	type clientItems struct {
		Err    int
		Errmsg interface{}
		Data   *map[string]*structs.DatabaseItem
	}

	items := &clientItems{
		Data: database.GetItems(),
	}

	services.ZlibJSONReply(w, items)
}
