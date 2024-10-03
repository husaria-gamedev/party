package server

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Message struct {
	SenderId    uuid.UUID       `json:"senderId"`
	RecipientId uuid.UUID       `json:"recipientId"`
	Type        string          `json:"type"`
	Data        json.RawMessage `json:"data"`
}
