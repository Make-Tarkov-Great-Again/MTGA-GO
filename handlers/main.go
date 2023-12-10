package handlers

import (
	"MT-GO/data"
	"fmt"
	"github.com/go-chi/chi/v5"
	probing "github.com/prometheus-community/pro-bing"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"MT-GO/pkg"
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

const routeNotImplemented = "Route is not implemented yet, using empty values instead"

// GetBundleList returns a list of custom bundles to the client
func GetBundleList(w http.ResponseWriter, _ *http.Request) {
	manifests := data.GetBundleManifests()
	pkg.SendZlibJSONReply(w, manifests)
}

func GetBrandName(w http.ResponseWriter, _ *http.Request) {
	brand := pkg.GetBrandName()
	pkg.SendZlibJSONReply(w, brand)
}

func MainGameStart(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(map[string]any{
		"utc_time": tools.GetCurrentTimeInSeconds(),
	})
	pkg.SendZlibJSONReply(w, body)
}

func MainMenuLocale(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input, _ := data.GetLocaleMenuByName(chi.URLParam(r, "id"))
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainVersionValidate(w http.ResponseWriter, _ *http.Request) {
	pkg.SendZlibJSONReply(w, pkg.ApplyResponseBody(nil))
}

func MainLanguages(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input := data.GetLanguages()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainGameConfig(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	sessionID := pkg.GetSessionID(r)
	if !data.CheckRequestedResponseCache(route) {
		input, _ := pkg.GetGameConfig(sessionID)
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainItems(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		cache := pkg.CreateCachedResponse(data.GetItems())
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Response Time: %v\n", elapsedTime)
}

func MainCustomization(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input := data.GetCustomizations()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainGlobals(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input := data.GetGlobals()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainSettings(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input := data.GetMainSettings()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainProfileList(w http.ResponseWriter, r *http.Request) {
	sessionID := pkg.GetSessionID(r)
	profileList := pkg.GetMainProfileList(sessionID)
	body := pkg.ApplyResponseBody(profileList)
	pkg.SendZlibJSONReply(w, body)
}

func MainAccountCustomization(w http.ResponseWriter, _ *http.Request) {
	accountCustomization := pkg.GetMainAccountCustomization()
	body := pkg.ApplyResponseBody(accountCustomization)
	pkg.SendZlibJSONReply(w, body)
}

func MainLocale(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		lang := chi.URLParam(r, "id")
		input, _ := data.GetLocaleGlobalByName(lang)
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

var keepAlive = &KeepAlive{
	Msg: "OK",
}

func MainKeepAlive(w http.ResponseWriter, _ *http.Request) {
	keepAlive.UtcTime = tools.GetCurrentTimeInSeconds()
	//data.GetCachedResponses().SaveIfRequired()

	body := pkg.ApplyResponseBody(keepAlive)
	pkg.SendZlibJSONReply(w, body)
}

func MainNicknameReserved(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody("")
	pkg.SendZlibJSONReply(w, body)
}

type nicknameValidate struct {
	Nickname string `json:"nickname"`
}

func MainNicknameValidate(w http.ResponseWriter, r *http.Request) {
	validate := new(nicknameValidate)
	input, err := json.Marshal(pkg.GetParsedBody(r))
	if err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal(input, validate); err != nil {
		log.Println(err)
	}

	body := pkg.ValidateNickname(validate.Nickname)
	pkg.SendZlibJSONReply(w, body)
}

type profileCreate struct {
	UID string `json:"uid"`
}

func MainProfileCreate(w http.ResponseWriter, r *http.Request) {
	request := new(ProfileCreateRequest)
	input, err := json.Marshal(pkg.GetParsedBody(r))
	if err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal(input, request); err != nil {
		log.Println(err)
	}

	sessionID := pkg.GetSessionID(r)
	pkg.CreateProfile(sessionID, request.Side, request.Nickname, request.VoiceID, request.HeadID)
	body := pkg.ApplyResponseBody(&profileCreate{UID: sessionID})
	pkg.SendZlibJSONReply(w, body)
}

func MainChannelCreate(w http.ResponseWriter, r *http.Request) {
	notifier, err := pkg.GetChannelNotifier(pkg.GetSessionID(r))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(notifier)
	pkg.SendZlibJSONReply(w, body)
}

func MainProfileSelect(w http.ResponseWriter, r *http.Request) {
	channel := pkg.GetChannel(pkg.GetSessionID(r))

	body := pkg.ApplyResponseBody(channel)
	pkg.SendZlibJSONReply(w, body)
}

func ChangeVoice(w http.ResponseWriter, r *http.Request) {
	parsed := pkg.GetParsedBody(r)
	fmt.Println(parsed)
	//character := data.GetCharacterByID(pkg.GetSessionID(r))

	body := pkg.ApplyResponseBody(nil)
	pkg.SendZlibJSONReply(w, body)
}

func MainProfileStatus(w http.ResponseWriter, r *http.Request) {
	statuses := pkg.GetProfileStatuses(pkg.GetSessionID(r))

	body := pkg.ApplyResponseBody(statuses)
	pkg.SendZlibJSONReply(w, body)
}

func MainProfileSettings(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody("")
	pkg.SendZlibJSONReply(w, body)
}

func MainWeather(w http.ResponseWriter, _ *http.Request) {
	weather := data.GetWeather()
	body := pkg.ApplyResponseBody(weather)
	pkg.SendZlibJSONReply(w, body)
}

var locationsSet bool

func MainLocations(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !locationsSet {
		input := data.GetLocations()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
		locationsSet = true
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainTemplates(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input := data.GetHandbook()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainHideoutAreas(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input, _ := data.GetHideoutAreas()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainHideoutQTE(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input, _ := data.GetHideoutQTE()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainHideoutSettings(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input, _ := data.GetHideoutSettings()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainHideoutRecipes(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input, _ := data.GetHideoutRecipes()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainHideoutScavRecipes(w http.ResponseWriter, r *http.Request) {
	route := r.RequestURI
	if !data.CheckRequestedResponseCache(route) {
		input, _ := data.GetHideoutScavcase()
		cache := pkg.CreateCachedResponse(input)
		data.SetResponseCacheForRoute(route, cache)
	}

	input := data.GetRequestedResponseCache(route)
	pkg.SendJSONReply(w, input)
}

func MainBuildsList(w http.ResponseWriter, r *http.Request) {
	builds, err := pkg.GetBuildsList(pkg.GetSessionID(r))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(builds)
	pkg.SendZlibJSONReply(w, body)
}

func MainQuestList(w http.ResponseWriter, r *http.Request) {
	quests, err := pkg.GetQuestList(pkg.GetSessionID(r))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(quests)
	pkg.SendZlibJSONReply(w, body)
}

func MainCurrentGroup(w http.ResponseWriter, _ *http.Request) {
	group := &CurrentGroup{
		Squad: []any{},
	}
	body := pkg.ApplyResponseBody(group)
	pkg.SendZlibJSONReply(w, body)
}

func MainRepeatableQuests(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody([]any{})
	pkg.SendZlibJSONReply(w, body)
}

func GetServerList(w http.ResponseWriter, _ *http.Request) {
	serverConfig := data.GetServerConfig()
	port, _ := strconv.Atoi(serverConfig.Ports.Main)

	serverListings := data.HasGetServerListings()
	if len(serverListings) == 0 {
		serverListings = append(serverListings, data.ServerListing{
			IP:   serverConfig.IP,
			Port: port,
		})
	}

	body := pkg.ApplyResponseBody(serverListings)
	pkg.SendZlibJSONReply(w, body)
}

func MatchUpdatePing(w http.ResponseWriter, _ *http.Request) {
	serverListings := data.HasGetServerListings()
	for _, listing := range serverListings {
		addr := listing.IP + ":" + strconv.Itoa(listing.Port)
		probe, err := probing.NewPinger(addr)
		if err != nil {
			log.Fatalln(err)
		}
		probe.Count = 3
		err = probe.Run() // Blocks until finished.
		if err != nil {
			log.Fatalln(err)
		}
		avgRtt := probe.Statistics().AvgRtt
		listing.Ping = int(avgRtt)
	}

	body := pkg.ApplyResponseBody(serverListings)
	pkg.SendZlibJSONReply(w, body)
}

var version = &Version{
	IsValid:       true,
	LatestVersion: "",
}

func MainCheckVersion(w http.ResponseWriter, r *http.Request) {
	responseCache := data.GetCachedResponses()
	check := strings.TrimPrefix(r.Header.Get("App-Version"), "EFT Client ")
	if responseCache.Version != check {
		responseCache.Version = check
		responseCache.Save = true
	}

	version.LatestVersion = check
	body := pkg.ApplyResponseBody(version)
	pkg.SendZlibJSONReply(w, body)
}

func MainLogout(w http.ResponseWriter, r *http.Request) {
	profile, err := data.GetProfileByUID(pkg.GetSessionID(r))
	if err != nil {
		log.Fatalln(err)
	}

	profile.SaveProfile()
	//data.GetCachedResponses().SaveIfRequired()

	body := pkg.ApplyResponseBody(map[string]any{"status": "ok"})
	pkg.SendZlibJSONReply(w, body)
}

func MainPrices(w http.ResponseWriter, _ *http.Request) {
	supplyData := pkg.GetMainPrices()

	body := pkg.ApplyResponseBody(supplyData)
	pkg.SendZlibJSONReply(w, body)
}

func ExitFromMenu(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(map[string]any{})
	pkg.SendZlibJSONReply(w, body)
}

type localLoot struct {
	LocationID string `json:"locationId"`
	VariantID  int8   `json:"variantId"`
}

func GetLocalLoot(w http.ResponseWriter, r *http.Request) {
	loot := new(localLoot)
	input, err := json.Marshal(pkg.GetParsedBody(r))
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(input, loot); err != nil {
		log.Fatalln(err)
	}

	output, err := data.GetLocalLootByNameAndIndex(loot.LocationID, loot.VariantID)
	if err != nil {
		log.Fatal(err)
	}

	body := pkg.ApplyResponseBody(output)
	pkg.SendZlibJSONReply(w, body)
}

func RaidConfiguration(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(map[string]any{})
	pkg.SendZlibJSONReply(w, body)
}

type insuranceList struct {
	Traders []string `json:"traders"`
	Items   []string `json:"items"`
}

func InsuranceListCost(w http.ResponseWriter, r *http.Request) {
	insurances := new(insuranceList)
	if input, err := json.Marshal(pkg.GetParsedBody(r)); err != nil {
		log.Println(err)
	} else if err := json.Unmarshal(input, insurances); err != nil {
		log.Println(err)
	}

	costs, err := pkg.GetInsuranceCosts(pkg.GetSessionID(r), insurances.Traders, insurances.Items)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(costs)
	pkg.SendZlibJSONReply(w, body)
}

func InsuranceItemsCost(w http.ResponseWriter, r *http.Request) {
	parsedBody := pkg.GetParsedBody(r)
	fmt.Println(parsedBody)
	fmt.Println(routeNotImplemented)
	body := pkg.ApplyResponseBody(nil)
	pkg.SendZlibJSONReply(w, body)
}

func InviteCancelAll(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(map[string]any{})
	pkg.SendZlibJSONReply(w, body)
}

func MatchAvailable(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(false)
	pkg.SendZlibJSONReply(w, body)
}

func RaidNotReady(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(map[string]any{})
	pkg.SendZlibJSONReply(w, body)
}

func RaidReady(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(map[string]any{})
	pkg.SendZlibJSONReply(w, body)
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

func GroupStatus(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(groupStatusOutput)
	pkg.SendZlibJSONReply(w, body)
}

func LookingForGroupStart(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(nil)
	pkg.SendZlibJSONReply(w, body)
}

func LookingForGroupStop(w http.ResponseWriter, _ *http.Request) {
	body := pkg.ApplyResponseBody(nil)
	pkg.SendZlibJSONReply(w, body)
}

type botDifficulties struct {
	Easy       map[string]any `json:"easy"`
	Normal     map[string]any `json:"normal"`
	Hard       map[string]any `json:"hard"`
	Impossible map[string]any `json:"impossible"`
}

func GetBotDifficulty(w http.ResponseWriter, r *http.Request) {
	parsedBody := pkg.GetParsedBody(r)
	botName := strings.ToLower(parsedBody.(map[string]any)["name"].(string))

	difficulties := new(botDifficulties)
	if bot, _ := data.GetBotByName(botName); bot != nil {
		difficulties.Easy = bot.Difficulties["easy"]
		difficulties.Normal = bot.Difficulties["normal"]
		difficulties.Hard = bot.Difficulties["hard"]
		difficulties.Impossible = bot.Difficulties["impossible"]
	}

	pkg.SendZlibJSONReply(w, difficulties)
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
	conditions := new(botConditions)
	input, err := json.Marshal(pkg.GetParsedBody(r))
	if err != nil {
		log.Println(err)
	}
	if err = json.Unmarshal(input, conditions); err != nil {
		log.Println(err)
	}

	//TODO: Send bots lol
	bot := data.GetSacrificialBot()

	bots := make([]map[string]any, 0, 50)
	for _, condition := range conditions.Conditions {
		for i := int8(0); i < condition.Limit; i++ {
			clone := bot.Clone()
			clone["_id"] = tools.GenerateMongoID()
			clone["aid"] = i
			clone["Info"].(map[string]any)["Settings"].(map[string]any)["Role"] = condition.Role
			clone["Info"].(map[string]any)["Settings"].(map[string]any)["BotDifficulty"] = condition.Difficulty
			bots = append(bots, clone)
		}
	}
	body := pkg.ApplyResponseBody(bots)
	pkg.SendZlibJSONReply(w, body)
}

type offlineMatchEnd struct {
	ExitName    string  `json:"exitName"`
	ExitStatus  string  `json:"exitStatus"`
	RaidSeconds float64 `json:"raidSeconds"`
}

func OfflineMatchEnd(w http.ResponseWriter, r *http.Request) {
	matchEnd := new(offlineMatchEnd)
	input, err := json.Marshal(pkg.GetParsedBody(r))
	if err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal(input, matchEnd); err != nil {
		log.Println(err)
	}

	log.Println("\n:::::::::::: Offline Match Status ::::::::::::\nExitName:", matchEnd.ExitName, "\nExitStatus:", matchEnd.ExitStatus, "\nRaidSeconds:", matchEnd.RaidSeconds)
	log.Println()
	body := pkg.ApplyResponseBody(nil)
	pkg.SendZlibJSONReply(w, body)
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
	input, err := json.Marshal(pkg.GetParsedBody(r))
	if err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal(input, &save); err != nil {
		log.Println(err)
	}

	//TODO: Raid Profile Save
	err = tools.WriteToFile("/raidProfileSave.json", save)
	if err != nil {
		return
	}

	log.Println("Raid Profile Save not implemented yet!")
	body := pkg.ApplyResponseBody(nil)
	pkg.SendZlibJSONReply(w, body)
}

func AirdropConfig(w http.ResponseWriter, _ *http.Request) {
	airdropParams := data.GetAirdropParameters()
	pkg.SendZlibJSONReply(w, airdropParams)
}
