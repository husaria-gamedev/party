package connection

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/husaria-dev/party/server"
)

type Handler struct {
	rs        server.RoomService
	roomIdLen int
	upgrader  websocket.Upgrader

	connectionsMu *sync.RWMutex
	connections   map[string]*websocket.Conn
}

var _ http.Handler = (*Handler)(nil)

func NewHandler(rs server.RoomService, roomIdLen int) *Handler {
	return &Handler{
		rs:            rs,
		roomIdLen:     roomIdLen,
		upgrader:      websocket.Upgrader{},
		connectionsMu: &sync.RWMutex{},
		connections:   make(map[string]*websocket.Conn),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /websocket", h.handleHostWebsocket)
	mux.HandleFunc("GET /room/:id", h.handleJoinRoom)
}

func (h *Handler) getRoomConnection(roomId string) *websocket.Conn {
	h.connectionsMu.RLock()
	defer h.connectionsMu.RUnlock()
	return h.connections[roomId]
}

func (h *Handler) setRoomConnection(roomId string, conn *websocket.Conn) error {
	h.connectionsMu.Lock()
	defer h.connectionsMu.Unlock()
	if h.connections[roomId] != nil {
		return server.ErrConflict
	}
	h.connections[roomId] = conn
	return nil
}
