package websocket

import (
	common "MT-GO/user/mods/MTGA-Coop/mod/Common"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var ci *common.WebSocketHandler

func Init(instance *common.WebSocketHandler) {
	ci = instance
}

type WebSocketHandler struct {
	webSockets map[string]*websocket.Conn
	mu         sync.Mutex
}

var Instance *WebSocketHandler

func NewWebSocketHandler(webSocketPort int) *WebSocketHandler {
	h := &WebSocketHandler{
		webSockets: make(map[string]*websocket.Conn),
	}
	Instance = h
	go h.startServer(webSocketPort)
	go h.sendServerList()

	return h
}

func (h *WebSocketHandler) startServer(webSocketPort int) {
	server := &http.Server{Addr: fmt.Sprintf(":%d", webSocketPort)}
	http.HandleFunc("/ws", h.handleWebSocketConnection)

	fmt.Printf("[COOP] Web Socket Server is listening on port %d\n", webSocketPort)
	fmt.Println("[COOP] A temporary Web Socket Server until SPT-Aki opens theirs up for modding!")

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("[COOP] Error starting Web Socket Server:", err)
	}
}

func (h *WebSocketHandler) handleWebSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("[COOP] Error upgrading WebSocket connection:", err)
		return
	}
	defer conn.Close()

	sessionID := getSessionIDFromRequestURL(r.URL.String())

	h.mu.Lock()
	h.webSockets[sessionID] = conn
	h.mu.Unlock()

	fmt.Printf("%s has connected to Coop Web Socket\n", sessionID)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		h.processMessage(string(msg))
	}

	h.mu.Lock()
	delete(h.webSockets, sessionID)
	h.mu.Unlock()
}

func (h *WebSocketHandler) sendServerList() {
	for {
		time.Sleep(5 * time.Second)
		h.mu.Lock()

		var dataToSend []map[string]interface{}
		for _, match := range CoopMatches {
			if match.Private {
				continue
			}

			serverData := map[string]interface{}{
				"name":        match.Name,
				"serverId":    match.ServerID,
				"location":    match.Location,
				"playerCount": len(match.ConnectedPlayers),
				"password": func() string {
					if match.Password != "" {
						return "Yes"
					}
					return "No"
				}(),
				"state": match.Status,
			}
			dataToSend = append(dataToSend, serverData)
		}

		h.mu.Unlock()

		h.sendToAllWebSockets(map[string]interface{}{"serverList": dataToSend})
	}
}

func (h *WebSocketHandler) processMessage(msg string) {
	if strings.HasPrefix(msg, "MTC") {
		messageWithoutSITPrefix := msg[3:]
		serverID := messageWithoutSITPrefix[:24]
		messageWithoutSITPrefixes := messageWithoutSITPrefix[24:]

		h.mu.Lock()
		match, ok := CoopMatches[serverID]
		h.mu.Unlock()

		if ok {
			match.ProcessData(messageWithoutSITPrefixes)
		}
		return
	}

	if jsonArray := h.TryParseJSON(msg); jsonArray != nil {
		for _, obj := range jsonArray {
			h.processObject(obj)
		}
	} else if msg != "" && msg[0] == '{' {
		var jsonObject map[string]interface{}
		if err := json.Unmarshal([]byte(msg), &jsonObject); err != nil {
			fmt.Println("[COOP] Error parsing JSON:", err)
			return
		}
		h.processObject(jsonObject)
	}
}

func (h *WebSocketHandler) TryParseJSON(msg string) []map[string]interface{} {
	if msg != "" && msg[0] == '[' {
		var jsonArray []map[string]interface{}
		if err := json.Unmarshal([]byte(msg), &jsonArray); err != nil {
			fmt.Println("[COOP] Error parsing JSON array:", err)
			return nil
		}
		return jsonArray
	}
	return nil
}

func (h *WebSocketHandler) processObject(jsonObject map[string]interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()

	match, ok := CoopMatches[jsonObject["serverId"].(string)]
	if ok {
		if jsonObject["connect"] == true {
			match.PlayerJoined(jsonObject["accountId"].(string))
		} else {
			match.ProcessData(jsonObject)
		}
	}

	h.sendToAllWebSockets(jsonObject)
}

func (h *WebSocketHandler) sendToAllWebSockets(data map[string]interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for sessionID, conn := range h.webSockets {
		if err := conn.WriteJSON(data); err != nil {
			fmt.Printf("[COOP] Error sending data to %s: %s\n", sessionID, err)
		}
	}
}

func (h *WebSocketHandler) sendToWebSockets(sessions []string, data map[string]interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, session := range sessions {
		conn, ok := h.webSockets[session]
		if ok {
			if err := conn.WriteJSON(data); err != nil {
				fmt.Printf("[COOP] Error sending data to %s: %s\n", session, err)
			}
		}
	}
}

func getSessionIDFromRequestURL(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) > 0 {
		return strings.TrimSuffix(parts[len(parts)-1], "pmc")
	}
	return ""
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
