package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) handleGetRoom(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	roomId := req.PathValue("roomId")

	room, err := h.rs.GetRoom(ctx, roomId)
	if err != nil {
		http.Error(w, "cannot get room", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, "cannot send response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) handleCreateRoom(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	userId, err := uuid.Parse(req.URL.Query().Get("userId"))
	if err != nil {
		http.Error(w, "invalid userId: "+err.Error(), http.StatusBadRequest)
		return
	}
	room, err := h.rs.CreateRoom(ctx, userId)
	if err != nil {
		http.Error(w, "cannot create room: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		http.Error(w, "cannot send response: "+err.Error(), http.StatusInternalServerError)
	}
}
