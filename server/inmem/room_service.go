package inmem

import (
	"context"

	"github.com/google/uuid"
	"github.com/husaria-dev/party/server"
	"github.com/husaria-dev/party/server/randx"
)

type RoomService struct {
	Rooms []server.Room
}

var _ server.RoomService = (*RoomService)(nil)

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (rs *RoomService) CreateRoom(ctx context.Context, userId uuid.UUID) (server.Room, error) {
	r := server.Room{
		Code:       randx.String(5),
		HostUserId: userId,
	}
	rs.Rooms = append(rs.Rooms, r)
	return r, nil
}

func (rs *RoomService) GetRoom(ctx context.Context, id string) (server.Room, error) {
	for _, r := range rs.Rooms {
		if r.Code == id {
			return r, nil
		}
	}
	return server.Room{}, server.ErrNotFound
}
