package handler

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/husaria-dev/party/server"
)

type Handler struct {
	rs       server.RoomService
	upgrader websocket.Upgrader

	connectionPool *server.ConnectionPool
}

func New(rs server.RoomService, connectionPool *server.ConnectionPool) http.Handler {
	h := &Handler{
		rs:             rs,
		connectionPool: connectionPool,
		upgrader:       websocket.Upgrader{},
	}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ws", h.handleWS)
	mux.HandleFunc("GET /room/{roomId}", h.handleGetRoom)
	mux.HandleFunc("POST /room", h.handleCreateRoom)
	return mux
}
