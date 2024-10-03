package connection

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) handleJoinRoom(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	var joinRequest joinRoomRequest
	err := json.NewDecoder(req.Body).Decode(&joinRequest)
	if err != nil {
		http.Error(w, "cannot parse request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	room, err := h.rs.GetRoom(ctx, joinRequest.RoomId)
	if err != nil {
		http.Error(w, "cannot get room", http.StatusNotFound)
		return
	}

	conn := h.getRoomConnection(joinRequest.RoomId)
	if conn == nil {
		http.Error(w, "cannot get room connection", http.StatusInternalServerError)
		return
	}

	err = conn.WriteJSON(WebsocketMessage{WebsocketMessageTypeJoined, joinRequest.ConnectionInfo})
	if err != nil {
		http.Error(w, "cannot send message to host: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := joinRoomResponse{
		ConnectionInfo: room.ConnectionInfo,
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "cannot send response: "+err.Error(), http.StatusInternalServerError)
	}
}
