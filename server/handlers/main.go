package handlers

import (
	"MT-GO/database"
	"MT-GO/services"
	"MT-GO/tools"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
)

const routeNotImplemented = "Route is not implemented yet, using empty values instead"

// GetBundleList returns a list of custom bundles to the client
func GetBundleList(w http.ResponseWriter, _ *http.Request) {
	fmt.Println(routeNotImplemented)
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

func MainPutMetrics(w http.ResponseWriter, _ *http.Request) {
	services.ZlibJSONReply(w, services.ApplyResponseBody(nil))
}

func MainMenuLocale(w http.ResponseWriter, r *http.Request) {
	lang := strings.TrimPrefix(r.URL.Path, "/client/menu/locale/")
	menu := services.ApplyResponseBody(database.GetLocalesMenuByName(lang))
	services.ZlibJSONReply(w, menu)
}

func MainVersionValidate(w http.ResponseWriter, _ *http.Request) {
	services.ZlibJSONReply(w, services.ApplyResponseBody(nil))
}

func MainLanguages(w http.ResponseWriter, _ *http.Request) {
	languages := services.ApplyResponseBody(database.GetLanguages())
	services.ZlibJSONReply(w, languages)
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
			Lobby:     database.GetLobbyAddress(),
			Trading:   database.GetTradingAddress(),
			Messaging: database.GetMessageAddress(),
			Main:      database.GetMainAddress(),
			RagFair:   database.GetRagFairAddress(),
		},
		UseProtobuf:       false,
		UtcTime:           tools.GetCurrentTimeInSeconds(),
		TotalInGame:       0, //account.GetTotalInGame
		ReportAvailable:   true,
		TwitchEventMember: false,
	})

	services.ZlibJSONReply(w, gameConfig)
}

const itemsRoute string = "/client/items"

func MainItems(w http.ResponseWriter, _ *http.Request) {
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

func MainCustomization(w http.ResponseWriter, _ *http.Request) {
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

func MainGlobals(w http.ResponseWriter, _ *http.Request) {
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

func MainSettings(w http.ResponseWriter, _ *http.Request) {
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

	if character == nil || character.ID == "" {
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

func MainAccountCustomization(w http.ResponseWriter, _ *http.Request) {
	customization := database.GetCustomizations()
	var output []string
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

var keepAlive = &KeepAlive{
	Msg:     "OK",
	UtcTime: 0,
}

func MainKeepAlive(w http.ResponseWriter, _ *http.Request) {
	keepAlive.UtcTime = tools.GetCurrentTimeInSeconds()

	body := services.ApplyResponseBody(keepAlive)
	services.ZlibJSONReply(w, body)
}

func MainNicknameReserved(w http.ResponseWriter, _ *http.Request) {
	body := services.ApplyResponseBody("")
	services.ZlibJSONReply(w, body)
}

func MainNicknameValidate(w http.ResponseWriter, r *http.Request) {
	parsedData := services.GetParsedBody(r).(map[string]interface{})

	nickname, ok := parsedData["nickname"]
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

	_, ok = database.Nicknames[nickname.(string)]
	if ok {
		body := services.ApplyResponseBody(nil)
		body.Err = 225
		body.Errmsg = "225 - "

		services.ZlibJSONReply(w, body)
		return
	}

	status := &NicknameValidate{
		Status: "ok",
	}
	body := services.ApplyResponseBody(status)
	services.ZlibJSONReply(w, body)
}

func MainProfileCreate(w http.ResponseWriter, r *http.Request) {
	request := new(ProfileCreateRequest)
	body, _ := json.Marshal(services.GetParsedBody(r))
	if err := json.Unmarshal(body, request); err != nil {
		panic(err)
	}

	sessionID := services.GetSessionID(r)

	profile := database.GetProfileByUID(sessionID)

	edition := database.GetEdition("Edge Of Darkness")
	if edition == nil {
		log.Fatalln("[MainProfileCreate] Edition is nil, this ain't good fella!")
	}
	var pmc database.Character

	if request.Side == "Bear" {
		pmc = *edition.Bear
		profile.Storage.Suites = edition.Storage.Bear
	} else {
		pmc = *edition.Usec
		profile.Storage.Suites = edition.Storage.Usec
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
	profile.Cache = profile.SetCache()
	profile.SaveProfile()

	data := services.ApplyResponseBody(map[string]interface{}{"uid": sessionID})
	services.ZlibJSONReply(w, data)
}

var channel = &Channel{}

func MainChannelCreate(w http.ResponseWriter, _ *http.Request) {
	body := services.ApplyResponseBody(channel.Notifier)
	services.ZlibJSONReply(w, body)
}

func MainProfileSelect(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)

	notifierServer := fmt.Sprintf("%s/push/notifier/get/%s", database.GetLobbyIPandPort(), sessionID)
	wssServer := fmt.Sprintf("%s/push/notifier/getwebsocket/%s", database.GetWebSocketAddress(), sessionID)

	channel.Status = "ok"
	Notifier := &channel.Notifier

	Notifier.Server = database.GetLobbyIPandPort()
	Notifier.ChannelID = sessionID
	Notifier.NotifierServer = notifierServer
	Notifier.WS = wssServer

	body := services.ApplyResponseBody(channel)
	services.ZlibJSONReply(w, body)
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

func MainWeather(w http.ResponseWriter, _ *http.Request) {
	weather := database.GetWeather()
	body := services.ApplyResponseBody(weather)
	services.ZlibJSONReply(w, body)
}

func MainLocations(w http.ResponseWriter, _ *http.Request) {
	locations := database.GetLocations()
	body := services.ApplyResponseBody(locations)
	services.ZlibJSONReply(w, body)
}

func MainTemplates(w http.ResponseWriter, _ *http.Request) {
	templates := database.GetHandbook()
	body := services.ApplyResponseBody(templates)
	services.ZlibJSONReply(w, body)
}

func MainHideoutAreas(w http.ResponseWriter, _ *http.Request) {
	areas := database.GetHideout().Areas
	body := services.ApplyResponseBody(areas)
	services.ZlibJSONReply(w, body)
}

func MainHideoutQTE(w http.ResponseWriter, _ *http.Request) {
	qte := database.GetHideout().QTE
	body := services.ApplyResponseBody(qte)
	services.ZlibJSONReply(w, body)
}

func MainHideoutSettings(w http.ResponseWriter, _ *http.Request) {
	settings := database.GetHideout().Settings
	body := services.ApplyResponseBody(settings)
	services.ZlibJSONReply(w, body)
}

func MainHideoutRecipes(w http.ResponseWriter, _ *http.Request) {
	recipes := database.GetHideout().Recipes
	body := services.ApplyResponseBody(recipes)
	services.ZlibJSONReply(w, body)
}

func MainHideoutScavRecipes(w http.ResponseWriter, _ *http.Request) {
	scavCaseRecipes := database.GetHideout().ScavCase
	body := services.ApplyResponseBody(scavCaseRecipes)
	services.ZlibJSONReply(w, body)
}

func MainBuildsList(w http.ResponseWriter, r *http.Request) {
	builds := database.GetProfileByUID(services.GetSessionID(r)).Storage.Builds
	body := services.ApplyResponseBody(builds)
	services.ZlibJSONReply(w, body)
}

func MainQuestList(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	quests := database.GetCharacterByUID(sessionID).GetQuestsAvailableToPlayer()
	body := services.ApplyResponseBody(quests)
	services.ZlibJSONReply(w, body)
}

func MainCurrentGroup(w http.ResponseWriter, _ *http.Request) {
	group := &CurrentGroup{
		Squad: []interface{}{},
	}
	body := services.ApplyResponseBody(group)
	services.ZlibJSONReply(w, body)
}

func MainRepeatableQuests(w http.ResponseWriter, _ *http.Request) {
	body := services.ApplyResponseBody([]interface{}{})
	services.ZlibJSONReply(w, body)
}

func MainServerList(w http.ResponseWriter, _ *http.Request) {
	var serverListings []ServerListing
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
	version := &Version{
		IsValid:       true,
		LatestVersion: check,
	}
	body := services.ApplyResponseBody(version)
	services.ZlibJSONReply(w, body)
}

func MainLogout(w http.ResponseWriter, r *http.Request) {
	database.GetProfileByUID(services.GetSessionID(r)).SaveProfile()

	body := services.ApplyResponseBody(map[string]interface{}{"status": "ok"})
	services.ZlibJSONReply(w, body)
}

func MainPrices(w http.ResponseWriter, _ *http.Request) {
	prices := *database.GetPrices()
	nextResupply := database.SetResupplyTimer()

	supplyData := &SupplyData{
		SupplyNextTime: nextResupply,
		Prices:         prices,
		CurrencyCourses: CurrencyCourses{
			RUB: prices["5449016a4bdc2d6f028b456f"],
			EUR: prices["569668774bdc2da2298b4568"],
			DOL: prices["5696686a4bdc2da3298b456a"],
		},
	}

	body := services.ApplyResponseBody(supplyData)
	services.ZlibJSONReply(w, body)
}

var actionHandlers = map[string]func(map[string]interface{}, *database.Character, *database.ProfileChangesEvent){
	"QuestAccept": func(moveAction map[string]interface{}, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.QuestAccept(moveAction["qid"].(string))
	},
	"Examine": func(moveAction map[string]interface{}, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.ExamineItem(moveAction)
	},
	"Move": func(moveAction map[string]interface{}, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.MoveItemInStash(moveAction, profileChangeEvent)
	},
	"Swap": func(moveAction map[string]interface{}, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.SwapItemInStash(moveAction, profileChangeEvent)
	},
	"Fold": func(moveAction map[string]interface{}, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.FoldItem(moveAction, profileChangeEvent)
	},
	/*
		"ReadEncyclopedia": func(moveAction map[string]interface{}, character *database.Character) interface{} {
			// Handle the "ReadEncyclopedia" action here
			return nil
		}, */
}

func MainItemsMoving(w http.ResponseWriter, r *http.Request) {
	data := services.GetParsedBody(r).(map[string]interface{})["data"].([]interface{})
	length := int8(len(data))

	character := database.GetCharacterByUID(services.GetSessionID(r))
	profileChangeEvent := database.GetProfileChangeByUID(character.ID)

	for i := int8(0); i < length; i++ {
		moveAction := data[i].(map[string]interface{})
		action := moveAction["Action"].(string)

		if handler, ok := actionHandlers[action]; ok {
			handler(moveAction, character, profileChangeEvent)
		} else {
			fmt.Println(action, "is not supported, sending empty response")
		}
	}

	//_ = tools.WriteToFile("/profileChangeEvent.json", profileChangeEvent)

	character.SaveCharacter(character.ID)
	services.ZlibJSONReply(w, services.ApplyResponseBody(profileChangeEvent))
}
