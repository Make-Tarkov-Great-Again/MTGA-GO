package srv

import (
	"net"
	"net/http"
	"sync/atomic"
)

var CW = &ConnectionWatcher{}

type ConnectionWatcher struct {
	n int64
}

func (cw *ConnectionWatcher) OnStateChange(_ net.Conn, state http.ConnState) {
	switch state {
	case http.StateNew: //Connection open
		cw.Add(1)
	case http.StateHijacked, http.StateClosed: //Connection Closed
		cw.Add(-1)
	case http.StateActive, http.StateIdle:
		return
	default:
		panic("unhandled default case")
	}
}

func (cw *ConnectionWatcher) Count() int {
	return int(atomic.LoadInt64(&cw.n))
}

func (cw *ConnectionWatcher) Add(c int64) {
	atomic.AddInt64(&cw.n, c)
}
