package inmem

import (
	"context"

	"github.com/husaria-dev/party/server"
)

type RoomService struct {
	Rooms []server.Room
}

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (rs *RoomService) CreateRoom(ctx context.Context, r server.Room) error {
	rs.Rooms = append(rs.Rooms, r)
	return nil
}

func (rs *RoomService) GetRoom(ctx context.Context, id string) (server.Room, error) {
	for _, r := range rs.Rooms {
		if r.Id == id {
			return r, nil
		}
	}
	return server.Room{}, server.ErrNotFound
}
