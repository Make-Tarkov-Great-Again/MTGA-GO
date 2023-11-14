package hndlr

import (
	"MT-GO/data"
	"log"
	"net/http"
	"strconv"
	"strings"

	"MT-GO/pkg"
	"MT-GO/tools"

	"github.com/goccy/go-json"
)

const routeNotImplemented = "Route is not implemented yet, using empty values instead"

// GetBundleList returns a list of custom bundles to the client
func GetBundleList(w http.ResponseWriter, r *http.Request) {
	manifests := data.GetBundleManifests()
	pkg.ZlibJSONReply(w, r.RequestURI, manifests)
}

func GetBrandName(w http.ResponseWriter, r *http.Request) {
	brand := pkg.GetBrandName()
	pkg.ZlibJSONReply(w, r.URL.Path, brand)
}

func ShowPersonKilledMessage(w http.ResponseWriter, r *http.Request) {
	pkg.ZlibJSONReply(w, r.RequestURI, "true")
}

func MainGameStart(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody(map[string]any{
		"utc_time": tools.GetCurrentTimeInSeconds(),
	})
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainPutMetrics(w http.ResponseWriter, r *http.Request) {
	pkg.ZlibJSONReply(w, r.RequestURI, pkg.ApplyResponseBody(nil))
}

func MainMenuLocale(w http.ResponseWriter, r *http.Request) {
	menu, err := data.GetLocalesMenuByName(r.URL.Path[20:])
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(menu)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainVersionValidate(w http.ResponseWriter, r *http.Request) {
	pkg.ZlibJSONReply(w, r.RequestURI, pkg.ApplyResponseBody(nil))
}

func MainLanguages(w http.ResponseWriter, r *http.Request) {
	languages := data.GetLanguages()
	body := pkg.ApplyResponseBody(languages)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainGameConfig(w http.ResponseWriter, r *http.Request) {
	gameConfig, err := pkg.GetGameConfig(pkg.GetSessionID(r))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(gameConfig)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainItems(w http.ResponseWriter, r *http.Request) {
	body := pkg.GetMainItems()
	pkg.ZlibJSONReply(w, r.RequestURI, body)

	log.Println("You know you're going to have to go back and try creating structs in your data, you lazy twit!")
}

func MainCustomization(w http.ResponseWriter, r *http.Request) {
	body := pkg.GetMainCustomization()
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainGlobals(w http.ResponseWriter, r *http.Request) {
	body := pkg.GetMainGlobals()
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainSettings(w http.ResponseWriter, r *http.Request) {
	body := pkg.GetMainSettings()
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainProfileList(w http.ResponseWriter, r *http.Request) {
	profileList := pkg.GetMainProfileList(pkg.GetSessionID(r))
	body := pkg.ApplyResponseBody(profileList)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainAccountCustomization(w http.ResponseWriter, r *http.Request) {
	accountCustomization := pkg.GetMainAccountCustomization()
	body := pkg.ApplyResponseBody(accountCustomization)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainLocale(w http.ResponseWriter, r *http.Request) {
	body, err := pkg.GetMainLocale(r.URL.Path[15:])
	if err != nil {
		log.Println(err)
	}

	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

var keepAlive = &KeepAlive{
	Msg:     "OK",
	UtcTime: 0,
}

func MainKeepAlive(w http.ResponseWriter, r *http.Request) {
	keepAlive.UtcTime = tools.GetCurrentTimeInSeconds()

	body := pkg.ApplyResponseBody(keepAlive)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainNicknameReserved(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody("")
	pkg.ZlibJSONReply(w, r.RequestURI, body)
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
	pkg.ZlibJSONReply(w, r.RequestURI, body)
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
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainChannelCreate(w http.ResponseWriter, r *http.Request) {
	notifier, err := pkg.GetChannelNotifier(pkg.GetSessionID(r))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(notifier)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainProfileSelect(w http.ResponseWriter, r *http.Request) {
	channel := pkg.GetChannel(pkg.GetSessionID(r))

	body := pkg.ApplyResponseBody(channel)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainProfileStatus(w http.ResponseWriter, r *http.Request) {
	statuses := pkg.GetProfileStatuses(pkg.GetSessionID(r))

	body := pkg.ApplyResponseBody(statuses)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainWeather(w http.ResponseWriter, r *http.Request) {
	weather := data.GetWeather()
	body := pkg.ApplyResponseBody(weather)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainLocations(w http.ResponseWriter, r *http.Request) {
	locations := data.GetLocations()
	body := pkg.ApplyResponseBody(locations)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainTemplates(w http.ResponseWriter, r *http.Request) {
	templates := data.GetHandbook()
	body := pkg.ApplyResponseBody(templates)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutAreas(w http.ResponseWriter, r *http.Request) {
	areas, err := data.GetHideoutAreas()
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(areas)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutQTE(w http.ResponseWriter, r *http.Request) {
	qte, err := data.GetHideoutQTE()
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(qte)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := data.GetHideoutSettings()
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(settings)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := data.GetHideoutRecipes()
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(recipes)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainHideoutScavRecipes(w http.ResponseWriter, r *http.Request) {
	scavCaseRecipes, err := data.GetHideoutScavcase()
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(scavCaseRecipes)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainBuildsList(w http.ResponseWriter, r *http.Request) {
	builds, err := pkg.GetBuildsList(pkg.GetSessionID(r))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(builds)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainQuestList(w http.ResponseWriter, r *http.Request) {
	quests, err := pkg.GetQuestList(pkg.GetSessionID(r))
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(quests)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainCurrentGroup(w http.ResponseWriter, r *http.Request) {
	group := &CurrentGroup{
		Squad: []any{},
	}
	body := pkg.ApplyResponseBody(group)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainRepeatableQuests(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody([]any{})
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

var serverListings []ServerListing

func GetServerList(w http.ResponseWriter, r *http.Request) {
	srvcfg := data.GetServerConfig()
	port, _ := strconv.Atoi(srvcfg.Ports.Main)
	serverListings = append(serverListings, ServerListing{
		IP:   srvcfg.IP,
		Port: port,
	})

	body := pkg.ApplyResponseBody(serverListings)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainCheckVersion(w http.ResponseWriter, r *http.Request) {
	check := strings.TrimPrefix(r.Header.Get("App-Version"), "EFT Client ")
	version := &Version{
		IsValid:       true,
		LatestVersion: check,
	}
	body := pkg.ApplyResponseBody(version)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainLogout(w http.ResponseWriter, r *http.Request) {
	if profile, err := data.GetProfileByUID(pkg.GetSessionID(r)); err != nil {
		log.Fatalln(err)
	} else {
		profile.SaveProfile()
	}

	body := pkg.ApplyResponseBody(map[string]any{"status": "ok"})
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MainPrices(w http.ResponseWriter, r *http.Request) {
	supplyData := pkg.GetMainPrices()

	body := pkg.ApplyResponseBody(supplyData)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func ExitFromMenu(w http.ResponseWriter, r *http.Request) {
	//TODO: IDK WHAT SIT NEEDS HERE
	body := pkg.ApplyResponseBody(nil)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

type localLoot struct {
	LocationID string `json:"locationId"`
	VariantID  int8   `json:"variantId"`
}

func GetLocalLoot(w http.ResponseWriter, r *http.Request) {
	loot := new(localLoot)
	input, err := json.Marshal(pkg.GetParsedBody(r))
	if err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal(input, loot); err != nil {
		log.Println(err)
	}

	output := data.GetLocalLootByNameAndIndex(loot.LocationID, loot.VariantID)
	body := pkg.ApplyResponseBody(output)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func RaidConfiguration(w http.ResponseWriter, r *http.Request) {
	/*
		TODO: Pre-raid nonsense that we might need to do
		AKI does some shit with setting difficulties to bots or something? IDK
		IDC
		IM THE GREATEST
	*/

	body := pkg.ApplyResponseBody(nil)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

type insuranceList struct {
	Traders []string `json:"traders"`
	Items   []string `json:"items"`
}

func InsuranceListCost(w http.ResponseWriter, r *http.Request) {
	insurances := new(insuranceList)
	input, err := json.Marshal(pkg.GetParsedBody(r))
	if err != nil {
		log.Println(err)
	}
	if err := json.Unmarshal(input, insurances); err != nil {
		log.Println(err)
	}

	costs, err := pkg.GetInsuranceCosts(pkg.GetSessionID(r), insurances.Traders, insurances.Items)
	if err != nil {
		log.Println(err)
	}

	body := pkg.ApplyResponseBody(costs)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func InviteCancelAll(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody(nil)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func MatchAvailable(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody(false)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func RaidNotReady(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody(map[string]any{})
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func RaidReady(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody(map[string]any{})
	pkg.ZlibJSONReply(w, r.RequestURI, body)
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
	body := pkg.ApplyResponseBody(groupStatusOutput)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func LookingForGroupStart(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody(nil)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func LookingForGroupStop(w http.ResponseWriter, r *http.Request) {
	body := pkg.ApplyResponseBody(nil)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
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
		bots := pkg.GetParsedBody(r).([]string)
			data := make(map[string]*botDifficulties)
			for _, key := range bots {
				difficulties := new(botDifficulties)
				if bot, _ := data.GetBotByName(strings.ToLower(pkg.GetParsedBody(r).(map[string]any)["name"].(string))); bot != nil {
					difficulties.Easy = bot.Difficulties["easy"]
					difficulties.Normal = bot.Difficulties["normal"]
					difficulties.Hard = bot.Difficulties["hard"]
					difficulties.Impossible = bot.Difficulties["impossible"]
				}
				data[key] = difficulties
			}
			pkg.ZlibJSONReply(w, r.RequestURI, data)
	*/

	difficulties := new(botDifficulties)
	if bot, _ := data.GetBotByName(strings.ToLower(pkg.GetParsedBody(r).(map[string]any)["name"].(string))); bot != nil {
		difficulties.Easy = bot.Difficulties["easy"]
		difficulties.Normal = bot.Difficulties["normal"]
		difficulties.Hard = bot.Difficulties["hard"]
		difficulties.Impossible = bot.Difficulties["impossible"]
	}

	pkg.ZlibJSONReply(w, r.RequestURI, difficulties)
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
	pkg.ZlibJSONReply(w, r.RequestURI, body)
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
	pkg.ZlibJSONReply(w, r.RequestURI, body)
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

	log.Println("Raid Profile Save not implemented yet!")
	body := pkg.ApplyResponseBody(nil)
	pkg.ZlibJSONReply(w, r.RequestURI, body)
}

func AirdropConfig(w http.ResponseWriter, r *http.Request) {
	airdropParams := data.GetAirdropParameters()
	pkg.ZlibJSONReply(w, r.RequestURI, airdropParams)
}
