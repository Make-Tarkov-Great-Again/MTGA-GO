// Handles the transfer of data between packages breaking our circular imports
package common

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	webSockets  map[string]*websocket.Conn
	mu          sync.Mutex
	CoopMatches map[string]*CoopMatch // Map to store CoopMatches by server ID
}

type CoopMatch struct {
	CreatedDateTime             time.Time
	LastUpdateDateTime          time.Time
	ExpectedNumberOfPlayers     int
	ConnectedPlayers            []ConnectedPlayer
	Characters                  []Character
	LastDataByAccountId         map[int]map[string]interface{}
	LastDataReceivedByAccountId map[int]map[string]interface{}
	LastData                    map[string]interface{}
	LastMoves                   map[string]interface{}
	LastRotates                 map[string]interface{}
	DamageArray                 []int
	Status                      CoopMatchStatus
	Settings                    map[string]interface{}
	Loot                        map[string]interface{}
	SpawnPoint                  SpawnPoint
	Private                     bool
	Name                        string
	ServerId                    int
	Location                    string
	Time                        string
	WeatherSettings             map[string]interface{}
	CheckStillRunningInterval   *time.Timer
	SendLastDataInterval        *time.Timer
}

type ConnectedPlayer struct {
	AccountID string
	IsOnline  bool
}

func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{
		webSockets:  make(map[string]*websocket.Conn),
		CoopMatches: make(map[string]*CoopMatch), // Initialize the map
	}
}

func (h *WebSocketHandler) AddWebSocket(sessionID string, conn *websocket.Conn) {
	h.mu.Lock()
	h.webSockets[sessionID] = conn
	h.mu.Unlock()
}

func (h *WebSocketHandler) RemoveWebSocket(sessionID string) {
	h.mu.Lock()
	delete(h.webSockets, sessionID)
	h.mu.Unlock()
}

func (h *WebSocketHandler) SendToAllWebSockets(data map[string]interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for sessionID, conn := range h.webSockets {
		if err := conn.WriteJSON(data); err != nil {
			fmt.Printf("[COOP] Failed to send data to websocket with ID %s", sessionID)
		}
	}
}

func (h *WebSocketHandler) SendToWebSockets(sessions string, data map[string]interface{}) {
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

func (h *WebSocketHandler) AreThereAnyWebSocketsOpen(sessions []ConnectedPlayer) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, session := range sessions {
		conn, ok := h.webSockets[session.AccountID]
		if ok && conn.WriteMessage(websocket.PingMessage, nil) == nil {
			return true
		}
	}

	return false
}

type CoopMatches struct {
	name        string
	serverId    string
	location    string
	playerCount int
	password    bool
	state       any
}
type Character struct {
	AccountId any
	IsDead    bool
}

type SpawnPoint struct {
	X float64
	Y float64
	Z float64
}

type CoopMatchStatus int

const (
	Loading CoopMatchStatus = iota
	InGame
	Complete
)
