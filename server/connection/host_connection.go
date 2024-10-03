package connection

import (
	"errors"
	"log"
	"net/http"

	"github.com/husaria-dev/party/server"
)

func (h *Handler) handleHostWebsocket(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	roomId := req.FormValue("roomId")
	password := req.FormValue("password")
	connectionInfo := req.FormValue("connectionInfo")

	room, err := h.rs.GetRoom(ctx, roomId)
	if errors.Is(err, server.ErrNotFound) {
		room = server.Room{
			Id:             roomId,
			AdminPassword:  password,
			ConnectionInfo: connectionInfo,
		}
		err := h.rs.CreateRoom(ctx, room)
		if err != nil {
			http.Error(w, "cannot create room: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		http.Error(w, "cannot get room: "+err.Error(), http.StatusInternalServerError)
		return
	} else if room.AdminPassword != password {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	conn, err := h.upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("failed to upgrade to websocket: " + err.Error())
	}

	err = h.setRoomConnection(room.Id, conn)
	if err != nil {
		http.Error(w, "cannot connect to room: "+err.Error(), http.StatusConflict)
		return
	}

	conn.ReadMessage() // host should send anything, so this should block forever

	delete(h.connections, room.Id)
}
