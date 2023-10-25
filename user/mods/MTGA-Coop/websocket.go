package MTGACoop

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// type webSocketHandler map[string]*websocket.Conn
type webSocketHandler struct {
	connections map[string]*websocket.Conn
}

var ws webSocketHandler

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
	Method    string //"PlayerInventoryController_UnloadMagazine"
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

//SendDataToPool -> ToggleLauncher

var packetJumpTable = map[string]func(map[string]interface{}){
	"SetTriggerPressed": func(packet map[string]interface{}) {
		printMethod(packet["Method"].(string))
	},
}

func printMethod(method string) {
	fmt.Println(method)
}

func handlePacket(packet interface{}) {
	packetMap, ok := packet.(map[string]interface{})
	if !ok {
		fmt.Println("God is a mean kid with a magnifying glass")
		return
	}

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
}

func setWebSocketHandler() {
	ws = webSocketHandler{
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

}
