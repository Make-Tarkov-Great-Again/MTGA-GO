package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
	"fmt"
	"net/http"
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
		UtcTime:           float64(tools.GetCurrentTimeInSeconds()),
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
			body := services.ApplyCRCResponseBody(database.GetCustomizations(), services.GetCachedCRC(customizationRoute))
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
		if trader.Base == nil {
			fmt.Println()
		}

		data = append(data, trader.Base)
	}

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

	if character.ID == "" {
		profiles := services.ApplyResponseBody([]interface{}{})
		services.ZlibJSONReply(w, profiles)
		fmt.Println("Character doesn't exist, begin creation")
	} else {

		/* 		var data []interface{}
		   		file := tools.GetJSONRawMessage("dummyData.json")
		   		err := json.Unmarshal(file, &data)
		   		if err != nil {
		   			panic(err)
		   		} */

		playerScav := database.GetPlayerScav()
		playerScav.Info.RegistrationDate = int(tools.GetCurrentTimeInSeconds())
		playerScav.AID = character.AID
		playerScav.ID = character.Savage.(string)

		slice := []structs.PlayerTemplate{*playerScav, *character}
		body := struct {
			Err    int                      `json:"err"`
			Errmsg interface{}              `json:"errmsg"`
			Data   []structs.PlayerTemplate `json:"data"`
		}{
			Data: slice,
		}
		//body := services.ApplyResponseBody(data)
		services.ZlibJSONReply(w, body)
	}
}

func ClientAccountCustomization(w http.ResponseWriter, r *http.Request) {
	customization := database.GetCustomizations()
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

	data := struct {
		Msg     string `json:"msg"`
		UtcTime int    `json:"utc_time"`
	}{
		Msg:     "OK",
		UtcTime: int(tools.GetCurrentTimeInSeconds()),
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

type ProfileCreateRequest struct {
	Side     string `json:"side"`
	Nickname string `json:"nickname"`
	HeadID   string `json:"headId"`
	VoiceID  string `json:"voiceId"`
}

func ProfileCreate(w http.ResponseWriter, r *http.Request) {
	request := &ProfileCreateRequest{}
	body, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, request)
	if err != nil {
		panic(err)
	}

	sessionID := services.GetSessionID(r)

	profile := database.GetProfileByUID(sessionID)
	if profile.Storage == nil {
		profile.Storage = &structs.Storage{}
	}

	editions := database.GetEdition("Edge Of Darkness")
	var pmc structs.PlayerTemplate

	if request.Side == "Bear" {
		pmc = *editions.Bear
		profile.Storage.Suites = editions.Storage.Bear
	} else {
		pmc = *editions.Usec
		profile.Storage.Suites = editions.Storage.Usec
	}

	pmc.ID = sessionID
	pmc.AID = profile.Account.AID
	sid, _ := tools.GenerateMongoID()
	pmc.Savage = sid

	pmc.Info.Side = request.Side
	pmc.Info.Nickname = request.Nickname

	pmc.Info.LowerNickname = strings.ToLower(request.Nickname)
	pmc.Info.Voice = database.GetCustomization(request.VoiceID)["_name"].(string)

	time := int(tools.GetCurrentTimeInSeconds())
	pmc.Info.RegistrationDate = time

	pmc.Health.UpdateTime = time

	pmc.Customization.Head = request.HeadID

	stats := &pmc.Stats.Eft
	stats.SessionCounters = nil
	stats.OverallCounters = map[string]interface{}{"Items": []interface{}{}}
	stats.Aggressor = nil
	stats.DroppedItems = make([]interface{}, 0, 0)
	stats.FoundInRaidItems = make([]interface{}, 0, 0)
	stats.Victims = make([]interface{}, 0, 0)
	stats.CarriedQuestItems = make([]interface{}, 0, 0)
	stats.DamageHistory = map[string]interface{}{
		"BodyParts":        []interface{}{},
		"LethalDamage":     nil,
		"LethalDamagePart": "Head",
	}
	stats.SurvivorClass = "Unknown"

	commonSkills := make([]structs.SkillsCommon, 0, len(pmc.Skills.Common))
	for _, skill := range pmc.Skills.Common {
		commonSkills = append(commonSkills, skill)
	}
	pmc.Skills.Common = commonSkills

	hideout := &pmc.Hideout
	resizedAreas := make([]structs.PlayerHideoutArea, 0, len(hideout.Areas))
	for _, area := range hideout.Areas {
		resizedAreas = append(resizedAreas, area)
	}

	hideout.Areas = resizedAreas
	hideout.Improvement = make(map[string]interface{})

	services.SaveCharacter(sessionID, pmc)
	profile.Character = &pmc

	fmt.Println()
}
