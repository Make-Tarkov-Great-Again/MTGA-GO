package MTGACoop

import (
	"MT-GO/services"
	"MT-GO/tools"
	"fmt"
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

var coopRoutes = map[string]http.HandlerFunc{
	"/coop/connect":       coopConnect,
	"/coop/server-status": coopServerStatus,
	"/coop/get-invites":   coopGetInvites,
	"/coop/server/delete": coopServerDelete,
	//"/coop/server/update":       coopServerUpdate,
	"/coop/server/update/weatherSettings": coopServerUpdateWeather,
	"/coop/server/update/spawnPoint":      coopServerUpdateSpawnPoint,
	"/coop/server/read/players":           coopServerReadPlayers,
	//"/coop/server/join": handlers.CoopServerJoin,
	"/coop/server/exist":             coopServerExist,
	"/coop/server/state":             coopServerState,
	"/coop/server/create":            coopServerCreate,
	"/coop/server/getAllForLocation": coopServerGetAllForLocation,
	//"/coop/server/friendlyAI": handlers.CoopServerFriendlyAI,
	//"/coop/server/spawnPoint": handlers.CoopServerSpawnPoint,
}

func coopServerStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Coop Server Match Status")
	services.ZlibReply(w, r.RequestURI, "")
}

// var coopStatuses = map[string]coopInvites

type coopInvites struct {
	Players []map[string]interface{} `json:"players"`
	Invite  []interface{}            `json:"invite"`
	Group   []interface{}            `json:"group"`
}

var coopStatusOutput = coopInvites{
	Players: make([]map[string]interface{}, 0),
	Invite:  make([]interface{}, 0),
	Group:   make([]interface{}, 0),
}

func coopGetInvites(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting Coop Server Invites")
	services.ZlibJSONReply(w, r.RequestURI, coopStatusOutput)
}

func coopServerDelete(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting Coop Server")
	//body := services.ApplyResponseBody(map[string]string{"response": "OK"})
	services.ZlibJSONReply(w, r.RequestURI, map[string]string{"response": "OK"})
}

func coopConnect(w http.ResponseWriter, r *http.Request) {
	//body := services.ApplyResponseBody(map[string]interface{}{})
	services.ZlibJSONReply(w, r.RequestURI, map[string]interface{}{})
}

func coopServerCreate(w http.ResponseWriter, r *http.Request) {
	info := new(serverInfo)
	data, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &info)
	if err != nil {
		log.Fatalln(err)
	}

	if checkIfMatchExists(info.ServerId) {
		//TODO: Delete existing match I guess
	}

	createCoopMatch(info)

	output := map[string]string{"serverId": info.ServerId}
	//body := services.ApplyResponseBody(output)
	services.ZlibJSONReply(w, r.RequestURI, output)
}

func coopServerState(w http.ResponseWriter, r *http.Request) {
	info := new(raidSettings)
	data, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &info)
	if err != nil {
		log.Fatalln(err)
	}

	match := getCoopMatch(info.KeyId)
	if match == nil {
		fmt.Println("Match does not exist")
		services.ZlibJSONReply(w, r.RequestURI, nil)
		return
	}

	if match.Location != info.LocationId ||
		match.Time != info.TimeVariant ||
		match.Status == complete ||
		match.LastUpdateDateTime < (tools.GetCurrentTimeInSeconds()-5) {

		fmt.Println("Match is over")
		services.ZlibJSONReply(w, r.RequestURI, nil)
		return
	}

	//body := services.ApplyResponseBody("hell yeah brother")
	fmt.Println("Match is alive")
	services.ZlibJSONReply(w, r.RequestURI, "hell yeah brother")
}

func coopServerExist(w http.ResponseWriter, r *http.Request) {
	parsedBody := services.GetParsedBody(r).(map[string]interface{})
	sid, ok := parsedBody["serverId"].(string)
	if !ok {
		fmt.Println("Server does not exist")
	}

	if checkIfMatchExists(sid) {
		fmt.Println("Match exists!")
		services.ZlibJSONReply(w, r.RequestURI, "hell yeah brother")
		return
	}

	fmt.Println("Match does not exist!")
	services.ZlibJSONReply(w, r.RequestURI, nil)
}

type serverReadPlayers struct {
	ServerId   string   `json:"serverId"`
	PlayerList []string `json:"pL"`
}

func coopServerReadPlayers(w http.ResponseWriter, r *http.Request) {
	readPlayers := new(serverReadPlayers)
	data, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &readPlayers)
	if err != nil {
		log.Fatalln(err)
	}

	//TODO: Adjust for future struct
	output := make([]interface{}, 0, len(readPlayers.PlayerList))
	match := getCoopMatch(readPlayers.ServerId)
	if match != nil {
		for _, pid := range readPlayers.PlayerList {
			player, ok := match.Characters[pid]
			if !ok {
				continue
			}
			output = append(output, player)
		}
	}

	//body := services.ApplyResponseBody(output)
	services.ZlibJSONReply(w, r.RequestURI, output)
}

type matchResponse struct {
	HostProfileId string
	HostName      string
	Settings      interface{}
	RaidTime      string
	Location      string
	PlayerCount   int16
}

func coopServerGetAllForLocation(w http.ResponseWriter, r *http.Request) {
	matchResponses := make([]matchResponse, 0, len(coopMatches))
	for mid, match := range coopMatches {
		conPlayers := int16(len(match.ConnectedPlayers))
		if conPlayers == 0 {
			continue
		}

		response := matchResponse{
			HostProfileId: mid,
			HostName:      match.Name,
			Settings:      match.Settings,
			PlayerCount:   conPlayers,
			Location:      match.Location,
		}
		matchResponses = append(matchResponses, response)
	}

	services.ZlibJSONReply(w, r.RequestURI, matchResponses)
}

type updateWeatherSettings struct {
	Type           string `json:"m"`
	ServerId       string `json:"serverId"`
	CloudinessType int8   `json:"ct"`
	//Clear = 0, PartlyCloud = 1, Cloudy = 2, CloudyWithGaps = 3, HeavyCloudCover = 4, Thundercloud = 5
	RainType int8 `json:"rt"`
	//NoRain = 0, Drizzling = 1, Rain = 2, Heavy = 3, Shower = 4
	WindType int8 `json:"wt"`
	//Light = 0, Moderate = 1, Strong = 2, VeryStrong = 3, Hurricane = 4
	FogType int8 `json:"ft"`
	//NoFog = 0, Faint = 1, Fog = 2, Heavy = 3, Continuous = 4
	TimeFlowType int8 `json:"tft"`
	//x0 = 0, x0_14 = 1, x0_25 = 2, x0_5 = 3, x1 = 4, x2 = 5, x4 = 6, x8 = 7
	HourOfDay int8 `json:"hod"`
}

func coopServerUpdateWeather(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		log.Fatal(err)
	}

	timeAndWeather := new(updateWeatherSettings)
	err = json.Unmarshal(data, &timeAndWeather)
	if err != nil {
		log.Fatal(err)
	}

	coopMatch := getCoopMatch(timeAndWeather.ServerId)
	if coopMatch == nil {
		log.Fatal("no coop match found")
		return
	}
	ws := &coopMatch.TimeAndWeatherSettings

	ws.CloudinessType = timeAndWeather.CloudinessType
	ws.FogType = timeAndWeather.FogType
	ws.HourOfDay = timeAndWeather.HourOfDay
	ws.RainType = timeAndWeather.RainType
	ws.TimeFlowType = timeAndWeather.TimeFlowType
	ws.WindType = timeAndWeather.WindType
}

type spawnPointForCoop struct {
	Type     string  `json:"m"`
	ServerId string  `json:"serverId"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        float64 `json:"Z"`
}

func coopServerUpdateSpawnPoint(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(services.GetParsedBody(r))
	if err != nil {
		log.Fatal(err)
	}

	spawnPoint := new(spawnPointForCoop)
	err = json.Unmarshal(data, &spawnPoint)
	if err != nil {
		log.Fatal(err)
	}

	coopMatch := getCoopMatch(spawnPoint.ServerId)
	if coopMatch == nil {
		log.Fatal("no coop match found")
		return
	}
	coopMatch.SpawnPoint.X = spawnPoint.X
	coopMatch.SpawnPoint.Y = spawnPoint.Y
	coopMatch.SpawnPoint.Z = spawnPoint.Z

	fmt.Println("SpawnPointForCoop updated")
}
