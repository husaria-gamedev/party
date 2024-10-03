package server

import (
	"context"

	"github.com/google/uuid"
)

type Room struct {
	Code       string    `json:"code"`
	HostUserId uuid.UUID `json:"hostUserId"`
}

type RoomService interface {
	CreateRoom(ctx context.Context, hostUserId uuid.UUID) (Room, error)
	GetRoom(ctx context.Context, code string) (Room, error)
}
