package MTGACoop

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// type webSocketHandler map[string]*websocket.Conn
type webSocketHandler struct {
	connections map[string]*websocket.Conn
}

var wsHandler webSocketHandler

//BackendCommunication.Instance.SendDataToPool("{ \"HostPing\": " + DateTime.Now.Ticks + " }");
//BackendCommunication.Instance.SendDataToPool("{ \"RaidTimer\": " + GameTimer.SessionTime.Value.Ticks + " }");

type itemController struct {
	Method                  string //"IC_MOVE"
	Time                    string
	GridItem                interface{} `json:"grad"`
	SlotItem                interface{} `json:"sitad"`
	StackSlotItem           interface{} `json:"ssad"`
	Id                      string
	Tpl                     string
	ItemControllerId        string `json:"icId"`
	ItemControllerCurrentId string `json:"icCId"`
}

type playerJump struct {
	Method    string `json:"Method"` //"Jump"
	AccountId string `json:"AccountId"`
}

type playerMove struct {
	AccountId string      `json:"AccountId"`
	Method    string      //"Move"
	PosX      float64     `json:"pX"`
	PosY      float64     `json:"pY"`
	PosZ      float64     `json:"pZ"`
	DirX      float64     `json:"dX"`
	DirY      float64     `json:"dY"`
	Speed     interface{} `json:"spd"` //find proper type
}

type loadMagazine struct {
	AccountId          string
	Method             string // "PlayerInventoryController_LoadMagazine"
	SourceAmmoId       string
	SourceTemplateId   string
	MagazineId         string
	MagazineTemplateId string
	LoadCount          int
	IgnoreRestrictions bool
}

type unloadMagainze struct {
	AccountId          string
	Method             string //"PlayerInventoryController_UnloadMagazine"
	MagazineId         string
	MagazineTemplateId string
}

type throwItem struct {
	AccountId  string
	Method     string //"PlayerInventoryController_ThrowItem"
	TemplateId string
	ItemId     string
}

type toggleItem struct {
	AccountId  string
	Method     string //"PlayerInventoryController_ToggleItem"
	ItemId     string
	TemplateId string
	ParentId   string
}

type pickUp struct {
	AccountId string
	Method    string //FCPickup
	Pickup    bool
}

type lightState struct {
	AccountId string
	Method    string //"SetLightsState"
	Id        string
	IsActive  bool
	LightMode int
}

type triggerPressed struct {
	AccountId string
	Method    string  //"SetTriggerPressed"
	Pressed   bool    `json:"pr"`
	RX        float64 `json:"rX"`
	RY        float64 `json:"rY"`
}

type toggleLauncher struct {
	Method string //ToggleLauncher
	Time   string
}

type EDamageType int

const (
	Undefined EDamageType = 1 << iota
	Fall
	Explosion
	Barbed
	Flame
	GrenadeFragment
	Impact
	Existence
	Medicine
	Bullet
	Melee
	Landmine
	Sniper
	Blunt
	LightBleeding
	HeavyBleeding
	Dehydration
	Exhaustion
	RadExposure
	Stimulator
	Poison
	LethalToxin
)

type kill struct {
	AccountId  string
	Method     string //Kill
	DamageType EDamageType
}

type EBodyPart int8

const (
	Head EBodyPart = iota
	Chest
	Stomach
	LeftArm
	RightArm
	LeftLeg
	RightLeg
	Common
)

type removeNegativeEffects struct {
	AccountId string
	Method    string //RemoveNegativeEffects
	BodyPart  EBodyPart
}

type restoreBodyPart struct {
	AccountId     string
	Method        string //"RestoreBodyPart"
	BodyPart      EBodyPart
	HealthPenalty float64
}

type localPlayer struct {
	AccountId  string `json:"accountId"`
	ServerId   string `json:"serverId"`
	Time       string `json:"t"`
	ProfileId  string `json:"profileId"`
	ProfileId2 string `json:"pId"`
}

type worldInteractiveObject struct {
	Method string      `json:"m"` //WIO_Interact
	Time   interface{} `json:"t"`
	DoorId string      `json:"doorId"`
	Type   string      `json:"type"`
}

var packetJumpTable = map[string]func(map[string]interface{}){
	"SetTriggerPressed": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"WIO_Interact": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"RestoreBodyPart": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"RemoveNegativeEffects": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"Kill": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"ToggleLauncher": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"SetLightsState": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"FCPickup": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"PlayerInventoryController_ToggleItem": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"PlayerInventoryController_ThrowItem": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"PlayerInventoryController_UnloadMagazine": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"PlayerInventoryController_LoadMagazine": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"Move": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"Jump": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
	"IC_MOVE": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
}

func printMethod(method string) {
	fmt.Println(method)
}

//TODO: Debug handlePacket until my eyes and fingers bleed, hell yeah brother

func handlePacket(packet interface{}) {
	packetMap, ok := packet.(map[string]interface{})
	if !ok {
		fmt.Println("God is a mean kid with a magnifying glass")
		return
	}
	//we will handle each specific packet individually

	var method string
	if m, ok := packetMap["m"].(string); ok {
		method = m
	} else if meth, ok := packetMap["Method"].(string); ok {
		method = meth
	} else {
		fmt.Println("God is a mean kid with a magnifying glass")
		return
	}

	if jumpFunc, found := packetJumpTable[method]; found {
		jumpFunc(packetMap)
	} else {
		fmt.Println("Unknown condition")
	}
} //Yeah

func setWebSocketHandler() {
	wsHandler = webSocketHandler{
		connections: make(map[string]*websocket.Conn), //mmmmmmmmmmmm this is annoying
	}
}

func (wsh *webSocketHandler) areThereAnyWebSocketsOpen() bool {
	for _, ws := range wsh.connections {
		if ws != nil && ws.WriteMessage(websocket.TextMessage, nil) == nil {
			return true
		}

	}
	return false
}

func (wsh *webSocketHandler) wsOnConnection(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	sessionId := strings.TrimSuffix(r.URL.Path, "/")
	sessionId = strings.TrimPrefix(sessionId, "/pmc")

	fmt.Printf("%s connected to die again to a cheater lol", sessionId)

	wsh.connections[sessionId] = ws

	packet := new(interface{})
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if err := json.Unmarshal(msg, &packet); err != nil {
			log.Fatalln(err)
		}

		handlePacket(packet)
	}

}

func (wsh *webSocketHandler) sendToAllAvailableWebSockets(data []byte) {
	for sessionID, ws := range wsh.connections {
		if ws == nil {
			fmt.Println("Websocket Connection for", sessionID, "is invalid, deleting")
			delete(wsh.connections, sessionID)
			continue
		}

		if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
			fmt.Println("Could not write message to websocket:", err)
			delete(wsh.connections, sessionID)
			continue
		}
	}
}

func (wsh *webSocketHandler) sendToWebSocket(socketId string, data []byte) {
	ws, ok := wsh.connections[socketId]
	if !ok || ws == nil {
		log.Fatalln("Websocket doesn't exist or is invalid, idc")
	}

	if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
		log.Fatalln(err)
	}
}
