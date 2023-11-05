package handlers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"MT-GO/database"
	"MT-GO/services"
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

const routeNotImplemented = "Route is not implemented yet, using empty values instead"

// GetBundleList returns a list of custom bundles to the client
func GetBundleList(w http.ResponseWriter, r *http.Request) {
	manifests := database.GetBundleManifests()
	services.ZlibJSONReply(w, r.RequestURI, manifests)
}

func GetBrandName(w http.ResponseWriter, r *http.Request) {
	brand := map[string]string{"name": database.GetServerConfig().BrandName}
	services.ZlibJSONReply(w, r.URL.Path, brand)
}

func ShowPersonKilledMessage(w http.ResponseWriter, r *http.Request) {
	services.ZlibJSONReply(w, r.RequestURI, "true")
}

func MainGameStart(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"utc_time": tools.GetCurrentTimeInSeconds(),
	}

	start := services.ApplyResponseBody(data)
	services.ZlibJSONReply(w, r.RequestURI, start)
}

func MainPutMetrics(w http.ResponseWriter, r *http.Request) {
	services.ZlibJSONReply(w, r.RequestURI, services.ApplyResponseBody(nil))
}

func MainMenuLocale(w http.ResponseWriter, r *http.Request) {
	lang := strings.TrimPrefix(r.URL.Path, "/client/menu/locale/")
	menu, err := database.GetLocalesMenuByName(lang)
	if err != nil {
		log.Fatalln(err)
		return
	}

	body := services.ApplyResponseBody(menu)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainVersionValidate(w http.ResponseWriter, r *http.Request) {
	services.ZlibJSONReply(w, r.RequestURI, services.ApplyResponseBody(nil))
}

func MainLanguages(w http.ResponseWriter, r *http.Request) {
	languages := services.ApplyResponseBody(database.GetLanguages())
	services.ZlibJSONReply(w, r.RequestURI, languages)
}

func MainGameConfig(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	lang := "en"
	if account, err := database.GetAccountByUID(sessionID); err != nil {
		log.Fatalln(err)
	} else if account.Lang != "" {
		lang = account.Lang
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

	services.ZlibJSONReply(w, r.RequestURI, gameConfig)
}

const itemsRoute string = "/client/items"

func MainItems(w http.ResponseWriter, r *http.Request) {
	if services.CheckIfResponseCanBeCached(itemsRoute) {
		if services.CheckIfResponseIsCached(itemsRoute) {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(itemsRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetItems(), services.GetCachedCRC(itemsRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		}
	}

	fmt.Println("You know you're going to have to go back and try creating structs in your database, you lazy twit!")
}

const customizationRoute string = "/client/customization"

func MainCustomization(w http.ResponseWriter, r *http.Request) {
	if services.CheckIfResponseCanBeCached(customizationRoute) {
		if services.CheckIfResponseIsCached(customizationRoute) {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(customizationRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetCustomizations(), services.GetCachedCRC(customizationRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		}
	}
}

const globalsRoute string = "/client/globals"

func MainGlobals(w http.ResponseWriter, r *http.Request) {
	if services.CheckIfResponseCanBeCached(globalsRoute) {
		if services.CheckIfResponseIsCached(globalsRoute) {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(globalsRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetGlobals(), services.GetCachedCRC(globalsRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		}
	}
}

const MainSettingsRoute string = "/client/settings"

func MainSettings(w http.ResponseWriter, r *http.Request) {
	if services.CheckIfResponseCanBeCached(MainSettingsRoute) {
		if services.CheckIfResponseIsCached(MainSettingsRoute) {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(MainSettingsRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		} else {
			body := services.ApplyCRCResponseBody(database.GetMainSettings(), services.GetCachedCRC(MainSettingsRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		}
	}
}

func MainProfileList(w http.ResponseWriter, r *http.Request) {

	sessionID := services.GetSessionID(r)
	character := database.GetCharacterByUID(sessionID)

	if character == nil || character.ID == "" {
		profiles := services.ApplyResponseBody([]any{})
		services.ZlibJSONReply(w, r.RequestURI, profiles)
		fmt.Println("Character doesn't exist, begin creation")
	} else {

		playerScav := database.GetPlayerScav()
		playerScav.Info.RegistrationDate = int32(tools.GetCurrentTimeInSeconds())
		playerScav.AID = character.AID
		playerScav.ID = *character.Savage

		slice := []any{*playerScav, *character}
		body := services.ApplyResponseBody(slice)
		services.ZlibJSONReply(w, r.RequestURI, body)
	}
}

func MainAccountCustomization(w http.ResponseWriter, r *http.Request) {
	customization := database.GetCustomizations()
	var output []string
	for id, c := range customization {
		if c.Props.Side != nil && len(c.Props.Side) > 0 {
			output = append(output, id)
		}
	}

	custom := services.ApplyResponseBody(output)
	services.ZlibJSONReply(w, r.RequestURI, custom)
}

const MainLocaleRoute string = "/client/locale/"

func MainLocale(w http.ResponseWriter, r *http.Request) {
	lang := strings.TrimPrefix(r.URL.Path, MainLocaleRoute)

	if services.CheckIfResponseCanBeCached(MainLocaleRoute) {
		if services.CheckIfResponseIsCached(r.URL.Path) {
			body := services.ApplyCRCResponseBody(nil, services.GetCachedCRC(MainLocaleRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		} else {
			locale, err := database.GetLocalesLocaleByName(lang)
			if err != nil {
				log.Fatalln(err)
			}

			body := services.ApplyCRCResponseBody(locale, services.GetCachedCRC(MainLocaleRoute))
			services.ZlibJSONReply(w, r.RequestURI, body)
		}
	}
}

var keepAlive = &KeepAlive{
	Msg:     "OK",
	UtcTime: 0,
}

func MainKeepAlive(w http.ResponseWriter, r *http.Request) {
	keepAlive.UtcTime = tools.GetCurrentTimeInSeconds()

	body := services.ApplyResponseBody(keepAlive)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainNicknameReserved(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody("")
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainNicknameValidate(w http.ResponseWriter, r *http.Request) {
	parsedData := services.GetParsedBody(r).(map[string]any)

	nickname, ok := parsedData["nickname"]
	if !ok {
		fmt.Println("For whatever reason, the nickname does not exist.")
	}

	if len(nickname.(string)) == 0 {
		body := services.ApplyResponseBody(nil)
		body.Err = 226
		body.Errmsg = "226 - "

		services.ZlibJSONReply(w, r.RequestURI, body)
		return
	}

	_, ok = database.Nicknames[nickname.(string)]
	if ok {
		body := services.ApplyResponseBody(nil)
		body.Err = 225
		body.Errmsg = "225 - "

		services.ZlibJSONReply(w, r.RequestURI, body)
		return
	}

	status := &NicknameValidate{
		Status: "ok",
	}
	body := services.ApplyResponseBody(status)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

type profileCreate struct {
	UID string `json:"uid"`
}

func MainProfileCreate(w http.ResponseWriter, r *http.Request) {
	request := new(ProfileCreateRequest)
	body, _ := json.Marshal(services.GetParsedBody(r))
	if err := json.Unmarshal(body, request); err != nil {
		log.Fatalln(err)
	}

	sessionID := services.GetSessionID(r)

	profile, err := database.GetProfileByUID(sessionID)
	if err != nil {
		log.Fatalln(err)
	}

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

	if customization, err := database.GetCustomization(request.VoiceID); err != nil {
		log.Fatalln(err)
	} else {
		pmc.Info.Voice = customization.Name
	}

	time := int32(tools.GetCurrentTimeInSeconds())
	pmc.Info.RegistrationDate = time
	pmc.Health.UpdateTime = time

	pmc.Customization.Head = request.HeadID

	stats := &pmc.Stats.Eft
	stats.SessionCounters = nil
	stats.OverallCounters = map[string]any{"Items": []any{}}
	stats.Aggressor = nil
	stats.DroppedItems = make([]any, 0)
	stats.FoundInRaidItems = make([]any, 0)
	stats.Victims = make([]any, 0)
	stats.CarriedQuestItems = make([]any, 0)
	stats.DamageHistory = map[string]any{
		"BodyParts":        []any{},
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

	hideout.Improvement = make(map[string]any)

	profile.Character = &pmc
	profile.Cache = profile.SetCache()
	profile.SaveProfile()

	data := services.ApplyResponseBody(&profileCreate{UID: sessionID})
	services.ZlibJSONReply(w, r.RequestURI, data)
}

var channel = &Channel{}

func MainChannelCreate(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(channel.Notifier)
	services.ZlibJSONReply(w, r.RequestURI, body)
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
	services.ZlibJSONReply(w, r.RequestURI, body)
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
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainWeather(w http.ResponseWriter, r *http.Request) {
	weather := database.GetWeather()
	body := services.ApplyResponseBody(weather)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainLocations(w http.ResponseWriter, r *http.Request) {
	locations := database.GetLocations()
	body := services.ApplyResponseBody(locations)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainTemplates(w http.ResponseWriter, r *http.Request) {
	templates := database.GetHandbook()
	body := services.ApplyResponseBody(templates)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutAreas(w http.ResponseWriter, r *http.Request) {
	areas := database.GetHideout().Areas
	body := services.ApplyResponseBody(areas)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutQTE(w http.ResponseWriter, r *http.Request) {
	qte := database.GetHideout().QTE
	body := services.ApplyResponseBody(qte)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutSettings(w http.ResponseWriter, r *http.Request) {
	settings := database.GetHideout().Settings
	body := services.ApplyResponseBody(settings)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutRecipes(w http.ResponseWriter, r *http.Request) {
	recipes := database.GetHideout().Recipes
	body := services.ApplyResponseBody(recipes)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutScavRecipes(w http.ResponseWriter, r *http.Request) {
	scavCaseRecipes := database.GetHideout().ScavCase
	body := services.ApplyResponseBody(scavCaseRecipes)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainBuildsList(w http.ResponseWriter, r *http.Request) {
	storage, err := database.GetStorageByUID(services.GetSessionID(r))
	if err != nil {
		log.Fatalln(err)
	}

	body := services.ApplyResponseBody(storage.Builds)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainQuestList(w http.ResponseWriter, r *http.Request) {
	sessionID := services.GetSessionID(r)
	quests := database.GetCharacterByUID(sessionID).GetQuestsAvailableToPlayer()
	body := services.ApplyResponseBody(quests)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainCurrentGroup(w http.ResponseWriter, r *http.Request) {
	group := &CurrentGroup{
		Squad: []any{},
	}
	body := services.ApplyResponseBody(group)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainRepeatableQuests(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody([]any{})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainServerList(w http.ResponseWriter, r *http.Request) {
	var serverListings []ServerListing
	port, _ := strconv.Atoi(database.GetServerConfig().Ports.Main)

	serverListings = append(serverListings, ServerListing{
		IP:   database.GetServerConfig().IP,
		Port: port,
	})

	body := services.ApplyResponseBody(serverListings)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainCheckVersion(w http.ResponseWriter, r *http.Request) {
	check := strings.TrimPrefix(r.Header.Get("App-Version"), "EFT Client ")
	version := &Version{
		IsValid:       true,
		LatestVersion: check,
	}
	body := services.ApplyResponseBody(version)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainLogout(w http.ResponseWriter, r *http.Request) {
	if profile, err := database.GetProfileByUID(services.GetSessionID(r)); err != nil {
		log.Fatalln(err)
	} else {
		profile.SaveProfile()
	}

	body := services.ApplyResponseBody(map[string]any{"status": "ok"})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MainPrices(w http.ResponseWriter, r *http.Request) {
	prices := database.GetPrices()
	nextResupply := database.SetResupplyTimer()

	supplyData := &SupplyData{
		SupplyNextTime: nextResupply,
		Prices:         prices,
		CurrencyCourses: CurrencyCourses{
			RUB: *prices["5449016a4bdc2d6f028b456f"],
			EUR: *prices["569668774bdc2da2298b4568"],
			DOL: *prices["5696686a4bdc2da3298b456a"],
		},
	}

	body := services.ApplyResponseBody(supplyData)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

var actionHandlers = map[string]func(map[string]any, *database.Character, *database.ProfileChangesEvent){
	"QuestAccept": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.QuestAccept(moveAction["qid"].(string), profileChangeEvent)
	},
	"Examine": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.ExamineItem(moveAction)
	},
	"Move": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.MoveItemInStash(moveAction, profileChangeEvent)
	},
	"Swap": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.SwapItemInStash(moveAction, profileChangeEvent)
	},
	"Fold": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.FoldItem(moveAction, profileChangeEvent)
	},
	"Merge": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.MergeItem(moveAction, profileChangeEvent)
	},
	"Transfer": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.TransferItem(moveAction)
	},
	"Split": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.SplitItem(moveAction, profileChangeEvent)
	},
	"ApplyInventoryChanges": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.ApplyInventoryChanges(moveAction)
	},
	"ReadEncyclopedia": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.ReadEncyclopedia(moveAction)
	},
	"TradingConfirm": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.TradingConfirm(moveAction, profileChangeEvent)
	},
	"Remove": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.RemoveItem(moveAction, profileChangeEvent)
	},
	"CustomizationBuy": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.CustomizationBuy(moveAction)
	},
	"CustomizationWear": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.CustomizationWear(moveAction)
	},
	"Bind": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.BindItem(moveAction)
	},
	"Tag": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.TagItem(moveAction)
	},
	"Toggle": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.ToggleItem(moveAction)
	},
	"HideoutUpgrade": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.HideoutUpgrade(moveAction, profileChangeEvent)
	},
	//HideoutUpgradeComplete
	"HideoutUpgradeComplete": func(moveAction map[string]any, character *database.Character, profileChangeEvent *database.ProfileChangesEvent) {
		character.HideoutUpgradeComplete(moveAction, profileChangeEvent)
	},
}

func MainItemsMoving(w http.ResponseWriter, r *http.Request) {
	parsed := services.GetParsedBody(r)
	data := parsed.(map[string]any)["data"].([]any)
	length := int8(len(data)) - 1

	character := database.GetCharacterByUID(services.GetSessionID(r))
	profileChangeEvent := database.CreateProfileChangesEvent(character)

	for i, move := range data {
		moveAction := move.(map[string]any)
		action := moveAction["Action"].(string)
		log.Println("[", i, "/", length, "] Action: ", action)

		if handler, ok := actionHandlers[action]; ok {
			handler(moveAction, character, profileChangeEvent)
		} else {
			fmt.Println(action, "is not supported, sending empty response")
		}
	}

	character.SaveCharacter(character.ID)
	services.ZlibJSONReply(w, r.RequestURI, services.ApplyResponseBody(profileChangeEvent))
}

func ExitFromMenu(w http.ResponseWriter, r *http.Request) {
	//TODO: IDK WHAT SIT NEEDS HERE
	body := services.ApplyResponseBody(nil)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

type localLoot struct {
	LocationID string `json:"locationId"`
	VariantID  int8   `json:"variantId"`
}

func GetLocalLoot(w http.ResponseWriter, r *http.Request) {
	localloot := new(localLoot)
	data, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, localloot)
	if err != nil {
		fmt.Println(err)
	}

	loot := database.GetLocalLootByNameAndIndex(localloot.LocationID, localloot.VariantID)
	body := services.ApplyResponseBody(loot)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func RaidConfiguration(w http.ResponseWriter, r *http.Request) {
	/*
		TODO: Pre-raid nonsense that we might need to do
		AKI does some shit with setting difficulties to bots or something? IDK
		IDC
		IM THE GREATEST
	*/

	body := services.ApplyResponseBody(nil)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

type insuranceList struct {
	Traders []string `json:"traders"`
	Items   []string `json:"items"`
}

type traderInsuranceInfo struct {
	LoyaltyLevel int8
	PriceCoef    int16
}

func InsuranceListCost(w http.ResponseWriter, r *http.Request) {
	insuranceListCost := new(insuranceList)
	data, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, insuranceListCost)
	if err != nil {
		fmt.Println(err)
	}

	sessionID := services.GetSessionID(r)
	output := make(map[string]map[string]int32)
	character := database.GetCharacterByUID(sessionID)
	invCache := database.GetCacheByUID(sessionID).Inventory

	Traders := make(map[string]traderInsuranceInfo)
	for _, TID := range insuranceListCost.Traders {
		trader, err := database.GetTraderByUID(TID)
		if err != nil {
			log.Fatalln("InsuranceListCost:", err)
		}

		trader.GetTraderLoyaltyLevel(character)

		Traders[TID] = traderInsuranceInfo{
			LoyaltyLevel: character.TradersInfo[TID].LoyaltyLevel,
			PriceCoef:    trader.Base.LoyaltyLevels[character.TradersInfo[TID].LoyaltyLevel].InsurancePriceCoef,
		}

		output[TID] = make(map[string]int32)
	}

	for _, itemID := range insuranceListCost.Items {
		itemInInventory := character.Inventory.Items[*invCache.GetIndexOfItemByUID(itemID)]
		itemPrice, err := database.GetPriceByID(itemInInventory.TPL)
		if err != nil {
			log.Fatalln(err)
		}

		for key, insuranceInfo := range Traders {
			insuranceCost := int32(math.Round(float64(*itemPrice) * 0.3))
			if insuranceInfo.PriceCoef > 0 {
				insuranceCost *= int32(1 - insuranceInfo.PriceCoef/100)
			}

			output[key][itemInInventory.TPL] = insuranceCost
		}
	}

	body := services.ApplyResponseBody(output)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func InviteCancelAll(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(nil)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func MatchAvailable(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(false)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func RaidNotReady(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(map[string]any{})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func RaidReady(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(map[string]any{})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

type groupStatus struct {
	Players []any `json:"players"`
	Invite  []any `json:"invite"`
	Group   []any `json:"group"`
}

var groupStatusOutput = groupStatus{
	Players: make([]any, 0),
	Invite:  make([]any, 0),
	Group:   make([]any, 0),
}

func GroupStatus(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(groupStatusOutput)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func LookingForGroupStart(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(nil)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func LookingForGroupStop(w http.ResponseWriter, r *http.Request) {
	body := services.ApplyResponseBody(nil)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

type botDifficulties struct {
	Easy       any `json:"easy"`
	Normal     any `json:"normal"`
	Hard       any `json:"hard"`
	Impossible any `json:"impossible"`
}

func GetBotDifficulty(w http.ResponseWriter, r *http.Request) {
	//TODO: For change
	/*
		bots := services.GetParsedBody(r).([]string)
			data := make(map[string]*botDifficulties)
			for _, key := range bots {
				difficulties := new(botDifficulties)
				if bot, _ := database.GetBotTypeByName(strings.ToLower(services.GetParsedBody(r).(map[string]any)["name"].(string))); bot != nil {
					difficulties.Easy = bot.Difficulties["easy"]
					difficulties.Normal = bot.Difficulties["normal"]
					difficulties.Hard = bot.Difficulties["hard"]
					difficulties.Impossible = bot.Difficulties["impossible"]
				}
				data[key] = difficulties
			}
			services.ZlibJSONReply(w, r.RequestURI, data)
	*/

	difficulties := new(botDifficulties)
	if bot, _ := database.GetBotTypeByName(strings.ToLower(services.GetParsedBody(r).(map[string]any)["name"].(string))); bot != nil {
		difficulties.Easy = bot.Difficulties["easy"]
		difficulties.Normal = bot.Difficulties["normal"]
		difficulties.Hard = bot.Difficulties["hard"]
		difficulties.Impossible = bot.Difficulties["impossible"]
	}

	services.ZlibJSONReply(w, r.RequestURI, difficulties)
}

type botConditions struct {
	Conditions []botCondition `json:"conditions"`
}
type botCondition struct {
	Role       string
	Limit      int8
	Difficulty string
}

func BotGenerate(w http.ResponseWriter, r *http.Request) {
	parsedBody := services.GetParsedBody(r)

	conditions := new(botConditions)
	data, err := json.Marshal(parsedBody)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &conditions)
	if err != nil {
		fmt.Println(err)
	}
	//TODO: Send bots lol
	body := services.ApplyResponseBody([]any{})
	services.ZlibJSONReply(w, r.RequestURI, body)
}

type offlineMatchEnd struct {
	ExitName    string  `json:"exitName"`
	ExitStatus  string  `json:"exitStatus"`
	RaidSeconds float64 `json:"raidSeconds"`
}

func OfflineMatchEnd(w http.ResponseWriter, r *http.Request) {
	matchEnd := new(offlineMatchEnd)
	data, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &matchEnd)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n:::::::::::: Offline Match Status ::::::::::::\nExitName:", matchEnd.ExitName, "\nExitStatus:", matchEnd.ExitStatus, "\nRaidSeconds:", matchEnd.RaidSeconds)
	fmt.Println()
	body := services.ApplyResponseBody(nil)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

type raidProfileSave struct {
	Exit                  string         `json:"exit"`
	Profile               map[string]any `json:"profile"`
	IsPlayerScav          bool           `json:"isPlayerScav"`
	Health                saveHealth     `json:"health"`
	DisableProgressionNow bool           `json:"disableProgressionNow"`
}

type saveHealth struct {
	IsAlive     bool
	Health      map[string]healthPart
	Hydration   float64
	Energy      float64
	Temperature float64
}

type healthPart struct {
	Maximum float64
	Current float64
	Effects map[string]any
}

func RaidProfileSave(w http.ResponseWriter, r *http.Request) {
	save := new(raidProfileSave)
	data, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &save)
	if err != nil {
		fmt.Println(err)
	}

	//TODO: Raid Profile Save
	err = tools.WriteToFile("/faggot.json", save)
	if err != nil {
		return
	}

	fmt.Println("Raid Profile Save not implemented yet!")
	body := services.ApplyResponseBody(nil)
	services.ZlibJSONReply(w, r.RequestURI, body)
}

func AirdropConfig(w http.ResponseWriter, r *http.Request) {
	airdropParams := database.GetAirdropParameters()
	services.ZlibJSONReply(w, r.RequestURI, airdropParams)
}
