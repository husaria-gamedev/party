package connection

type joinRoomRequest struct {
	RoomId         string `json:"roomId"`
	ConnectionInfo string `json:"connectionInfo"`
}

type joinRoomResponse struct {
	ConnectionInfo string `json:"connectionInfo"`
}

type WebsocketMessage struct {
	Type WebsocketMessageType `json:"type"`
	Data any                  `json:"data"`
}

type WebsocketMessageType string

const (
	WebsocketMessageTypeJoined WebsocketMessageType = "Joined"
)
