package database

const (
	Loading  int8 = 0
	InGame   int8 = 1
	Complete int8 = 2

	HostShutDownMethod      string = "host-shutdown"
	WebsocketTimeoutMessage string = "websocket-timeout"
	NoPlayersMessage        string = "no-players"
)

var coopMatches = make(map[string]CoopMatch)

type CoopMatch struct {
}
