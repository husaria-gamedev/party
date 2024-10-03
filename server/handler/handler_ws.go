package handler

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/husaria-dev/party/server"
)

func (h *Handler) handleWS(w http.ResponseWriter, req *http.Request) {
	userId, err := uuid.Parse(req.FormValue("userId"))
	if err != nil {
		http.Error(w, "invalid userId: "+err.Error(), http.StatusBadRequest)
		return
	}
	conn, err := h.upgrader.Upgrade(w, req, nil)
	if err != nil {
		http.Error(w, "cannot upgrade to WS: "+err.Error(), http.StatusInternalServerError)
		return
	}
	h.connectionPool.AddConnection(userId, conn)
	defer h.connectionPool.RemoveConnection(userId)

	h.handleHostWsConn(conn, userId)
}

func (h *Handler) handleHostWsConn(conn *websocket.Conn, userId uuid.UUID) {
	for {
		msg := server.Message{}
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Invalid message: " + err.Error())
			break
		}
		msg.SenderId = userId
		h.connectionPool.SendMessage(msg)
	}
}
