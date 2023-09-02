package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"MT-GO/tools"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
)

const route_not_implemented = "Route is not implemented yet, using empty values instead"

// GetBundleList returns a list of custom bundles to the client
func GetBundleList(w http.ResponseWriter, _ *http.Request) {
	fmt.Println(route_not_implemented)
	services.ZlibJSONReply(w, []string{})
}

/* func GetWebSocketAddress(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	database.SetWebSocketAddress(sessionID)
	websocketURL := database.GetWebSocketAddress()
	services.ZlibReply(w, websocketURL)
} */

func ShowPersonKilledMessage(w http.ResponseWriter, _ *http.Request) {
	services.ZlibJSONReply(w, "true")
}

func MainGameStart(w http.ResponseWriter, _ *http.Request) {
	data := map[string]interface{}{
		"utc_time": tools.GetCurrentTimeInSeconds(),
	}

	start := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, start)
}

func MainMenuLocale(w http.ResponseWriter, r *http.Request) {
	lang := strings.TrimPrefix(r.URL.Path, "/client/menu/locale/")
	menu := services.ApplyResponseBody(database.GetLocalesMenuByName(lang))
	services.ZlibJSONReply(w, menu)
}

func MainVersionValidate(w http.ResponseWriter, _ *http.Request) {
	services.ZlibJSONReply(w, services.ApplyResponseBody(nil))
}

func MainLanguages(w http.ResponseWriter, r *http.Request) {
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

func MainGameConfig(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	lang := database.GetAccountByUID(sessionID).Lang
	if lang == "" {
		lang = "en"
	}

	gameConfig := services.ApplyResponseBody(&GameConfig{
		Aid:             sessionID,
		Lang:            lang,
		Languages:       database.GetLanguages(),
		NdaFree:         false,
		Taxonomy:        6,
		ActiveProfileID: sessionID,
		Backend: Backend{
			Lobby:     database.GetMainAddress(),
			Trading:   database.GetTradingAddress(),
			Messaging: database.GetMessageAddress(),
			Main:      database.GetMainAddress(),
			RagFair:   database.GetRagFairAddress(),
		},
		UseProtobuf:       false,
		UtcTime:           float64(tools.GetCurrentTimeInSeconds()),
		TotalInGame:       0, //account.GetTotalInGame
		ReportAvailable:   true,
		TwitchEventMember: false,
	})

	services.ZlibJSONReply(w, gameConfig)
}

const itemsRoute string = "/client/items"

func MainItems(w http.ResponseWriter, r *http.Request) {
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

func MainCustomization(w http.ResponseWriter, r *http.Request) {
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

func MainGlobals(w http.ResponseWriter, r *http.Request) {
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

const MainSettingsRoute string = "/client/settings"

func MainSettings(w http.ResponseWriter, r *http.Request) {
	ok := services.CheckIfResponseCanBeCached(MainSettingsRoute)
	if ok {

		ok = services.CheckIfResponseIsCached(MainSettingsRoute)
		if ok {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(MainSettingsRoute))
			services.ZlibJSONReply(w, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetMainSettings(), services.GetCachedCRC(MainSettingsRoute))
			services.ZlibJSONReply(w, body)
		}
	}
}

func MainProfileList(w http.ResponseWriter, r *http.Request) {

	sessionID := services.GetSessionID(r)
	character := database.GetCharacterByUID(sessionID)

	if character.ID == "" {
		profiles := services.ApplyResponseBody([]interface{}{})
		services.ZlibJSONReply(w, profiles)
		fmt.Println("Character doesn't exist, begin creation")
	} else {

		playerScav := database.GetPlayerScav()
		playerScav.Info.RegistrationDate = int(tools.GetCurrentTimeInSeconds())
		playerScav.AID = character.AID
		playerScav.ID = *character.Savage

		slice := []interface{}{*playerScav, *character}
		body := services.ApplyResponseBody(slice)
		services.ZlibJSONReply(w, body)
	}
}

func MainAccountCustomization(w http.ResponseWriter, r *http.Request) {
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

		if len(side) > 0 {
			output = append(output, id)
		}
	}

	custom := services.ApplyResponseBody(output)
	services.ZlibJSONReply(w, custom)
}

const MainLocaleRoute string = "/client/locale/"

func MainLocale(w http.ResponseWriter, r *http.Request) {
	lang := strings.TrimPrefix(r.URL.Path, MainLocaleRoute)

	ok := services.CheckIfResponseCanBeCached(MainLocaleRoute)
	if ok {

		ok = services.CheckIfResponseIsCached(r.URL.Path)
		if ok {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(MainLocaleRoute))
			services.ZlibJSONReply(w, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetLocalesLocaleByName(lang), services.GetCachedCRC(MainLocaleRoute))
			services.ZlibJSONReply(w, body)
		}
	}
}

func MainKeepAlive(w http.ResponseWriter, r *http.Request) {

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

func MainNicknameReserved(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody("")
	services.ZlibJSONReply(w, body)
}

func MainNicknameValidate(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	if !services.IsNicknameAvailable(nickname.(string), database.GetProfiles()) {
		body := services.ApplyResponseBody(nil)
		body.Err = 225
		body.Errmsg = "225 - "

		services.ZlibJSONReply(w, body)
		return
	}

	status := struct {
		Status interface{} `json:"status"`
	}{
		Status: "ok",
	}
	body := services.ApplyResponseBody(status)
	services.ZlibJSONReply(w, body)
}

type ProfileCreateRequest struct {
	Side     string `json:"side"`
	Nickname string `json:"nickname"`
	HeadID   string `json:"headId"`
	VoiceID  string `json:"voiceId"`
}

func MainProfileCreate(w http.ResponseWriter, r *http.Request) {
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

	editions := database.GetEdition("Edge Of Darkness")
	var pmc database.Character

	if request.Side == "Bear" {
		pmc = *editions.Bear
		profile.Storage.Suites = editions.Storage.Bear
	} else {
		pmc = *editions.Usec
		profile.Storage.Suites = editions.Storage.Usec
	}

	pmc.ID = sessionID
	pmc.AID = profile.Account.AID

	sid := tools.GenerateMongoID()
	pmc.Savage = &sid

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
	stats.DroppedItems = make([]interface{}, 0)
	stats.FoundInRaidItems = make([]interface{}, 0)
	stats.Victims = make([]interface{}, 0)
	stats.CarriedQuestItems = make([]interface{}, 0)
	stats.DamageHistory = map[string]interface{}{
		"BodyParts":        []interface{}{},
		"LethalDamage":     nil,
		"LethalDamagePart": "Head",
	}
	stats.SurvivorClass = "Unknown"

	commonSkills := make([]database.SkillsCommon, 0, len(pmc.Skills.Common))
	commonSkills = append(commonSkills, pmc.Skills.Common...)
	pmc.Skills.Common = commonSkills

	hideout := &pmc.Hideout

	resizedAreas := make([]database.PlayerHideoutArea, 0, len(hideout.Areas))
	resizedAreas = append(resizedAreas, hideout.Areas...)
	hideout.Areas = resizedAreas

	hideout.Improvement = make(map[string]interface{})

	profile.Character = &pmc
	profile.SaveProfile()

	data := services.ApplyResponseBody(map[string]interface{}{"uid": sessionID})
	services.ZlibJSONReply(w, data)
}

type Notifier struct {
	Server         string `json:"server"`
	ChannelID      string `json:"channel_id"`
	URL            string `json:"url"`
	NotifierServer string `json:"notifierServer"`
	WS             string `json:"ws"`
}

type Channel struct {
	Status         string   `json:"status"`
	Notifier       Notifier `json:"notifier"`
	NotifierServer string   `json:"notifierServer"`
}

var channel = &Channel{}

func MainChannelCreate(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(channel.Notifier)
	services.ZlibJSONReply(w, body)
}
func MainProfileSelect(w http.ResponseWriter, r *http.Request) {

	sessionID := services.GetSessionID(r)

	notiServer := fmt.Sprintf("%s/push/notifier/get/%s", database.GetMainAddress(), sessionID)
	wssServer := fmt.Sprintf("%s/push/notifier/getwebsocket/%s", database.GetWebSocketAddress(), sessionID)

	channel.Status = "ok"
	Notifier := &channel.Notifier

	Notifier.Server = database.GetMainIPandPort() //probably will be lobby server
	fmt.Println("Probably need to set this to Lobby Server in the future")
	Notifier.ChannelID = sessionID
	Notifier.NotifierServer = notiServer
	Notifier.WS = wssServer

	body := services.ApplyResponseBody(channel)
	services.ZlibJSONReply(w, body)
}

type ProfileStatuses struct {
	MaxPVECountExceeded bool            `json:"maxPveCountExceeded"`
	Profiles            []ProfileStatus `json:"profiles"`
}

type ProfileStatus struct {
	ProfileID    string      `json:"profileid"`
	ProfileToken interface{} `json:"profileToken"`
	Status       string      `json:"status"`
	SID          string      `json:"sid"`
	IP           string      `json:"ip"`
	Port         int         `json:"port"`
}

func MainProfileStatus(w http.ResponseWriter, r *http.Request) {

	character := database.GetCharacterByUID(services.GetSessionID(r))

	scavProfile := &ProfileStatus{
		ProfileID: *character.Savage,
		Status:    "Free",
	}

	pmcProfile := &ProfileStatus{
		ProfileID: character.ID,
		Status:    "Free",
	}

	statuses := &ProfileStatuses{
		Profiles: []ProfileStatus{*scavProfile, *pmcProfile},
	}

	body := services.ApplyResponseBody(statuses)
	services.ZlibJSONReply(w, body)
}

func MainWeather(w http.ResponseWriter, r *http.Request) {
	weather := database.GetWeather()
	body := services.ApplyResponseBody(weather)
	services.ZlibJSONReply(w, body)
}

func MainLocations(w http.ResponseWriter, r *http.Request) {
	locations := database.GetLocations()
	body := services.ApplyResponseBody(locations)
	services.ZlibJSONReply(w, body)
}

func MainTemplates(w http.ResponseWriter, r *http.Request) {
	templates := database.GetHandbook()
	body := services.ApplyResponseBody(templates)
	services.ZlibJSONReply(w, body)
}

func MainHideoutAreas(w http.ResponseWriter, r *http.Request) {
	areas := database.GetHideout().Areas
	body := services.ApplyResponseBody(areas)
	services.ZlibJSONReply(w, body)
}

func MainHideoutQTE(w http.ResponseWriter, r *http.Request) {
	qte := database.GetHideout().QTE
	body := services.ApplyResponseBody(qte)
	services.ZlibJSONReply(w, body)
}

func MainHideoutSettings(w http.ResponseWriter, r *http.Request) {
	settings := database.GetHideout().Settings
	body := services.ApplyResponseBody(settings)
	services.ZlibJSONReply(w, body)
}

func MainHideoutRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := database.GetHideout().Recipes
	body := services.ApplyResponseBody(recipes)
	services.ZlibJSONReply(w, body)
}

func MainHideoutScavRecipes(w http.ResponseWriter, r *http.Request) {
	scavCaseRecipies := database.GetHideout().ScavCase
	body := services.ApplyResponseBody(scavCaseRecipies)
	services.ZlibJSONReply(w, body)
}

func MainBuildsList(w http.ResponseWriter, r *http.Request) {
	builds := database.GetProfileByUID(services.GetSessionID(r)).Storage.Builds
	body := services.ApplyResponseBody(builds)
	services.ZlibJSONReply(w, body)
}

func MainQuestList(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	quests := services.GetQuestsAvailableToPlayer(database.GetCharacterByUID(sessionID))
	body := services.ApplyResponseBody(quests)
	services.ZlibJSONReply(w, body)
}

type CurrentGroup struct {
	Squad []interface{} `json:"squad"`
}

func MainCurrentGroup(w http.ResponseWriter, r *http.Request) {
	group := &CurrentGroup{
		Squad: []interface{}{},
	}
	body := services.ApplyResponseBody(group)
	services.ZlibJSONReply(w, body)
}
func MainRepeatableQuests(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody([]interface{}{})
	services.ZlibJSONReply(w, body)
}

type ServerListing struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func MainServerList(w http.ResponseWriter, r *http.Request) {
	serverListings := []ServerListing{}
	port, _ := strconv.Atoi(database.GetServerConfig().Ports.Main)

	serverListings = append(serverListings, ServerListing{
		IP:   database.GetServerConfig().IP,
		Port: port,
	})

	body := services.ApplyResponseBody(serverListings)
	services.ZlibJSONReply(w, body)
}

func MainCheckVersion(w http.ResponseWriter, r *http.Request) {
	check := strings.TrimPrefix(r.Header["App-Version"][0], "EFT Client ")
	version := struct {
		IsValid       bool   `json:"isValid"`
		LatestVersion string `json:"latestVersion"`
	}{
		IsValid:       true,
		LatestVersion: check,
	}
	body := services.ApplyResponseBody(version)
	services.ZlibJSONReply(w, body)
}

func MainLogoout(w http.ResponseWriter, r *http.Request) {
	database.GetProfileByUID(services.GetSessionID(r)).SaveProfile()

	body := services.ApplyResponseBody(map[string]interface{}{"status": "ok"})
	services.ZlibJSONReply(w, body)
}

type SupplyData struct {
	SupplyNextTime  int            `json:"supplyNextTime"`
	Prices          map[string]int `json:"prices"`
	CurrencyCourses struct {
		RUB int `json:"5449016a4bdc2d6f028b456f"`
		EUR int `json:"569668774bdc2da2298b4568"`
		DOL int `json:"5696686a4bdc2da3298b456a"`
	} `json:"currencyCourses"`
}

func MainPrices(w http.ResponseWriter, r *http.Request) {
	prices := *database.GetPrices()
	nextResupply := database.SetResupplyTimer()

	supplyData := &SupplyData{
		SupplyNextTime: nextResupply,
		Prices:         prices,
		CurrencyCourses: struct {
			RUB int `json:"5449016a4bdc2d6f028b456f"`
			EUR int `json:"569668774bdc2da2298b4568"`
			DOL int `json:"5696686a4bdc2da3298b456a"`
		}{
			RUB: prices["5449016a4bdc2d6f028b456f"],
			EUR: prices["569668774bdc2da2298b4568"],
			DOL: prices["5696686a4bdc2da3298b456a"],
		},
	}

	body := services.ApplyResponseBody(supplyData)
	services.ZlibJSONReply(w, body)
}
