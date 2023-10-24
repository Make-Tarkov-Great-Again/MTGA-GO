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
	Type                    string      `json:"m"` //"IC_MOVE"
	Time                    string      `json:"t"`
	GridItem                interface{} `json:"grad"`
	SlotItem                interface{} `json:"sitad"`
	StackSlotItem           interface{} `json:"ssad"`
	Id                      string      `json:"id"`
	Tpl                     string      `json:"tpl"`
	ItemControllerId        string      `json:"icId"`
	ItemControllerCurrentId string      `json:"icCId"`
}

type playerJump struct {
	Method    string `json:"Method"`
	AccountId string `json:"AccountId"`
}

type playerMove struct {
	AccountId string      `json:"AccountId"`
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

//SendDataToPool -> UnloadMagazine

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
