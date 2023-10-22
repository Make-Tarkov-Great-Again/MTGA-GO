package mod

import (
	"MT-GO/server"
	"MT-GO/services"
	"encoding/json"
	"log"
	"net/http"
)

var coopRoutes = map[string]http.HandlerFunc{
	"/coop/connect":       coopConnect,
	"/coop/server-status": coopServerStatus,
	"/coop/get-invites":   coopGetInvites,
	"/coop/server/delete": coopDelete,
	//"/coop/server/update": handlers.CoopServerUpdate,
	//"/coop/server/read/players": handlers.CoopServerReadPlayers,
	//"/coop/server/join": handlers.CoopServerJoin,
	//"/coop/server/exist": handlers.CoopServerExists,
	//"/coop/server/create": handlers.CoopServerCreate,
	//"/coop/server/getAllForLocation": handlers.CoopServerGetAllForLocation,
	//"/coop/server/friendlyAI": handlers.CoopServerFriendlyAI,
	//"/coop/server/spawnPoint": handlers.CoopServerSpawnPoint,
}

func AddRoutes() {

}

func coopConnect(w http.ResponseWriter, r *http.Request) {
	services.ZlibJSONReply(w, "{}")
}

func coopDelete(w http.ResponseWriter, r *http.Request) {
	if r.RemoteAddr == "192.168.1.1" {
		log.Println("Deleting Coop Server")
		body := services.ApplyResponseBody(map[string]string{"response": "OK"})
		services.ZlibJSONReply(w, body)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func coopServerStatus(w http.ResponseWriter, r *http.Request) {
	count := server.CW.Count()
	body := services.ApplyResponseBody(map[string]int{"Players Connected: ": count})
	services.ZlibJSONReply(w, body)
}

func coopGetInvites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	obj := map[string]interface{}{
		"players": []interface{}{
			map[string]interface{}{},
			map[string]interface{}{},
		},
		"invite": []interface{}{},
		"group":  []interface{}{},
	}

	json.NewEncoder(w).Encode(obj)
}

func coopServerUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var info struct {
		ServerId string `json:"serverId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return
	}

	// ARE YOU FUCKING VALID YOU BITCH
	if info.ServerId == "" {
		log.Printf("/coop/server/update -- no info or serverId provided")
		json.NewEncoder(w).Encode(struct {
			Response string `json:"response"`
		}{
			Response: "ERROR",
		})
		return
	}

	// Check if coopMatch exists
	coopMatch := Main.getCoopMatch(info.ServerId) //Im changing the StayInTarkovInstance to main at some point lmfao
	if coopMatch == nil {
		log.Printf("/coop/server/update -- no coopMatch found to update")
		json.NewEncoder(w).Encode(struct{}{})
		return
	}

	// Process data
	coopMatch.ProcessData(r.Body)

	json.NewEncoder(w).Encode(struct{}{})
}
