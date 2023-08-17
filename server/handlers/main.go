package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"MT-GO/structs"
	"MT-GO/tools"
	"fmt"
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
	database.SetWebSocketAddress(sessionID)
	websocketURL := database.GetWebSocketAddress()
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
	UseProtobuf       bool              `json:"useProtobuf"`
	UtcTime           float64           `json:"utc_time"`
	TotalInGame       int               `json:"totalInGame"`
	ReportAvailable   bool              `json:"reportAvailable"`
	TwitchEventMember bool              `json:"twitchEventMember"`
}

func ClientGameConfig(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	lang := database.GetAccountByUID(sessionID).Lang
	if lang == "" {
		lang = "en"
	}

	timeAsFloat, _ := strconv.ParseFloat(tools.GetCurrentTimeInSeconds(), 64)
	mainAddress := database.GetMainAddress()

	gameConfig := services.ApplyResponseBody(&GameConfig{
		Aid:             sessionID,
		Lang:            lang,
		Languages:       database.GetLanguages(),
		NdaFree:         false,
		Taxonomy:        6,
		ActiveProfileID: sessionID,
		Backend: Backend{
			Lobby:     mainAddress, //database.GetLobbyAddress(),
			Trading:   mainAddress, //database.GetTradingAddress(),
			Messaging: mainAddress, //database.GetMessageAddress(),
			Main:      mainAddress, //database.GetMainAddress()
			RagFair:   mainAddress, //database.GetRagFairAddress(),
		},
		UseProtobuf:       false,
		UtcTime:           timeAsFloat,
		TotalInGame:       0, //account.GetTotalInGame
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
	output := []string{}
	for id, c := range customization {
		custom, ok := c.(map[string]interface{})
		if !ok {
			panic("customization is not a map[string]interface{}")
		}
		props, ok := custom["_props"].(map[string]interface{})
		if !ok {
			panic("customization properties are not map[string]interface{}")
		}
		side, ok := props["Side"].([]interface{})
		if !ok {
			continue
		}

		if side != nil && len(side) > 0 {
			output = append(output, id)
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
	timeToFloat, _ := strconv.ParseFloat(tools.GetCurrentTimeInSeconds(), 32)

	data := struct {
		Msg     string  `json:"msg"`
		UtcTime float32 `json:"utc_time"`
	}{
		Msg:     "OK",
		UtcTime: float32(timeToFloat),
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

	if len(nickname.(string)) == 0 {
		body := services.ApplyResponseBody(nil)
		body.Err = 226
		body.Errmsg = "226 - "

		services.ZlibJSONReply(w, body)
	} else {
		available := services.IsNicknameAvailable(nickname.(string))
		if !available {
			body := services.ApplyResponseBody(nil)
			body.Err = 225
			body.Errmsg = "225 - "

			services.ZlibJSONReply(w, body)
		} else {
			status := struct {
				Status interface{} `json:"status"`
			}{
				Status: "ok",
			}
			body := services.ApplyResponseBody(status)
			services.ZlibJSONReply(w, body)
		}
	}
}

func ProfileCreate(w http.ResponseWriter, r *http.Request) {
	body := services.GetParsedBody(r).(map[string]string)
	/* 	data, err := json.Marshal(services.GetParsedBody(r))
	   	if err != nil {
	   		panic(err)
	   	}
	   	err = json.Unmarshal(data, &body)
	   	if err != nil {
	   		panic(err)
	   	} */

	fmt.Println(body)
}
