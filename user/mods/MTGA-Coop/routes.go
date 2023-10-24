package MTGACoop

import (
	"MT-GO/services"
	"MT-GO/tools"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"net/http"
)

var coopRoutes = map[string]http.HandlerFunc{
	"/coop/connect":       coopConnect,
	"/coop/server-status": coopServerStatus,
	"/coop/get-invites":   coopGetInvites,
	"/coop/server/delete": coopServerDelete,
	//"/coop/server/update": handlers.CoopServerUpdate,
	"/coop/server/read/players": coopServerReadPlayers,
	//"/coop/server/join": handlers.CoopServerJoin,
	"/coop/server/exist":  coopServerExist,
	"/coop/server/create": coopServerCreate,
	//"/coop/server/getAllForLocation": handlers.CoopServerGetAllForLocation,
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

func coopServerExist(w http.ResponseWriter, r *http.Request) {
	parsedBody := services.GetParsedBody(r).(map[string]interface{})
	sid, ok := parsedBody["serverId"].(string)
	if !ok { //then it has raid settings
		info := new(raidSettings)
		data, err := json.Marshal(parsedBody)
		if err != nil {
			log.Fatalln(err)
		}
		err = json.Unmarshal(data, &info)
		if err != nil {
			log.Fatalln(err)
		}

		for _, match := range coopMatches {
			if match.Location != info.LocationId ||
				match.Time != info.TimeVariant ||
				match.Status != complete ||
				match.LastUpdateDateTime < (tools.GetCurrentTimeInSeconds()-5) {
				continue
			}

			//body := services.ApplyResponseBody("hell yeah brother")
			fmt.Println("Match exists!")
			services.ZlibJSONReply(w, r.RequestURI, "hell yeah brother")
			return
		}

		fmt.Println("Match does not exist!")
		//body := services.ApplyResponseBody(nil)
		services.ZlibJSONReply(w, r.RequestURI, nil)
		return
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
