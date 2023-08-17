package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"MT-GO/structs"
	"MT-GO/tools"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

const route_not_implemented = "Route is not implemented yet, using empty values instead"

// GetBundleList returns a list of custom bundles to the client
func GetBundleList(w http.ResponseWriter, _ *http.Request) {
	fmt.Println(route_not_implemented)
	services.ZlibJSONReply(w, []string{})
}

func GetWebSocketAddress(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	websocketURL := database.GetWebsocketURL()
	websocketURL = fmt.Sprintf(websocketURL, database.GetBackendAddress(), sessionID)
	services.ZlibReply(w, websocketURL)
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
	lang := strings.TrimPrefix(r.URL.Path, "/client/menu/locale/")
	menu := services.ApplyResponseBody(database.GetLocalesMenuByName(lang))
	services.ZlibJSONReply(w, menu)
}

func ClientVersionValidate(w http.ResponseWriter, _ *http.Request) {
	services.ZlibJSONReply(w, services.ApplyResponseBody(nil))
}

func ClientLanguages(w http.ResponseWriter, r *http.Request) {
	languages := services.ApplyResponseBody(database.GetLanguages())
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

	gameConfig := services.ApplyResponseBody(&GameConfig{
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
	})

	fmt.Println("Don't forget to create multiple servers later on you dumbshit!!!!!!!!!!!!")
	services.ZlibJSONReply(w, gameConfig)
}

const itemsRoute string = "/client/items"

func ClientItems(w http.ResponseWriter, r *http.Request) {
	ok := services.CheckIfResponseCanBeCached(itemsRoute)
	if ok {

		ok = services.CheckIfResponseIsCached(itemsRoute)
		if ok {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(itemsRoute))
			services.ZlibJSONReply(w, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetItems(), services.GetCachedCRC(itemsRoute))
			services.ZlibJSONReply(w, body)
		}
	}

	fmt.Println("You know you're going to have to go back and try creating structs in your database, you lazy twit!")
}

const customizationRoute string = "/client/customization"

func ClientCustomization(w http.ResponseWriter, r *http.Request) {
	ok := services.CheckIfResponseCanBeCached(customizationRoute)
	if ok {

		ok = services.CheckIfResponseIsCached(customizationRoute)
		if ok {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(customizationRoute))
			services.ZlibJSONReply(w, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetCustomization(), services.GetCachedCRC(customizationRoute))
			services.ZlibJSONReply(w, body)
		}
	}
}

const globalsRoute string = "/client/globals"

func ClientGlobals(w http.ResponseWriter, r *http.Request) {
	ok := services.CheckIfResponseCanBeCached(globalsRoute)
	if ok {

		ok = services.CheckIfResponseIsCached(globalsRoute)
		if ok {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(globalsRoute))
			services.ZlibJSONReply(w, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetGlobals(), services.GetCachedCRC(globalsRoute))
			services.ZlibJSONReply(w, body)
		}
	}
}

const traderSettingsRoute string = "/client/trading/api/traderSettings"

func ClientTraderSettings(w http.ResponseWriter, r *http.Request) {
	traders := database.GetTraders()
	data := make([]map[string]interface{}, 0, len(traders))

	for _, trader := range traders {
		base, ok := trader["Base"].(map[string]interface{})
		if ok {
			data = append(data, base)
		}
	}

	traders = nil

	body := services.ApplyResponseBody(&data)
	services.ZlibJSONReply(w, body)
}

const clientSettingsRoute string = "/client/settings"

func ClientSettings(w http.ResponseWriter, r *http.Request) {
	ok := services.CheckIfResponseCanBeCached(clientSettingsRoute)
	if ok {

		ok = services.CheckIfResponseIsCached(clientSettingsRoute)
		if ok {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(clientSettingsRoute))
			services.ZlibJSONReply(w, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetClientSettings(), services.GetCachedCRC(clientSettingsRoute))
			services.ZlibJSONReply(w, body)
		}
	}
}

func ClientProfileList(w http.ResponseWriter, r *http.Request) {

	sessionID := services.GetSessionID(r)
	character := database.GetCharacterByUID(sessionID)

	if character == nil {
		profiles := services.ApplyResponseBody([]interface{}{})
		services.ZlibJSONReply(w, profiles)
		fmt.Println("Character doesn't exist, begin creation")
	} else {

		playerScav := database.GetPlayerScav()
		playerScav.Info["RegistrationDate"] = tools.GetCurrentTimeInSeconds()

		aid, err := strconv.Atoi(sessionID)
		if err != nil {
			panic(err)
		}
		playerScav.AID = aid
		playerScav.ID = character.Savage

		body := services.ApplyResponseBody([]*structs.PlayerTemplate{playerScav, character})
		services.ZlibJSONReply(w, body)
	}
}

func ClientAccountCustomization(w http.ResponseWriter, r *http.Request) {
	customization := database.GetCustomization()
	output := []interface{}{}
	for _, c := range customization {
		custom, ok := c.(map[string]interface{})
		if ok {
			props, ok := custom["_props"].(map[string]interface{})
			if ok {
				side, ok := props["Side"].([]interface{})
				if ok {
					if side != nil && len(side) > 0 {
						output = append(output, custom)
					}
				}
			}
		}
	}

	custom := services.ApplyResponseBody(output)
	services.ZlibJSONReply(w, custom)
}

const clientLocaleRoute string = "/client/locale/"

func ClientLocale(w http.ResponseWriter, r *http.Request) {
	lang := strings.TrimPrefix(r.URL.Path, clientLocaleRoute)

	ok := services.CheckIfResponseCanBeCached(clientLocaleRoute)
	if ok {

		ok = services.CheckIfResponseIsCached(r.URL.Path)
		if ok {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(clientLocaleRoute))
			services.ZlibJSONReply(w, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetLocalesLocaleByName(lang), services.GetCachedCRC(clientLocaleRoute))
			services.ZlibJSONReply(w, body)
		}
	}
}

func KeepAlive(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Msg     string `json:"msg"`
		UtcTime string `json:"utc_time"`
	}{
		Msg:     "OK",
		UtcTime: tools.GetCurrentTimeInSeconds(),
	}

	body := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, body)
}

func NicknameReserved(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody("")
	services.ZlibJSONReply(w, body)
}

func NicknameValidate(w http.ResponseWriter, r *http.Request) {
	context := services.GetParsedBody(r).(map[string]interface{})

	nickname, ok := context["nickname"]
	if !ok {
		fmt.Println("For whatever reason, the nickname does not exist.")
	}

	available := services.IsNicknameAvailable(nickname.(string))
	if !available {
		body := services.ApplyResponseBody("The nickname is already in use")
		body.Err = 255

		services.ZlibJSONReply(w, body)
	} else {
		body := services.ApplyResponseBody(map[string]string{"status": "ok"})
		services.ZlibJSONReply(w, body)
	}
}
