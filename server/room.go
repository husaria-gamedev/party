package server

import "context"

type Room struct {
	Id             string `json:"id"`
	ConnectionInfo string `json:"connectionInfo"`
	AdminPassword  string `json:"-"`
}

type RoomService interface {
	CreateRoom(context.Context, Room) error
	GetRoom(context.Context, string) (Room, error)
}
