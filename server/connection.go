package server

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ConnectionPool struct {
	connMu *sync.RWMutex
	conns  map[uuid.UUID]*websocket.Conn
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		connMu: &sync.RWMutex{},
		conns:  map[uuid.UUID]*websocket.Conn{},
	}
}

func (c *ConnectionPool) AddConnection(id uuid.UUID, conn *websocket.Conn) error {
	c.connMu.Lock()
	defer c.connMu.Unlock()

	if _, exists := c.conns[id]; exists {
		return ErrConflict
	}
	c.conns[id] = conn

	return nil
}

func (c *ConnectionPool) RemoveConnection(id uuid.UUID) error {
	c.connMu.Lock()
	defer c.connMu.Unlock()

	delete(c.conns, id)
	return nil
}

func (c *ConnectionPool) getConnection(id uuid.UUID) *websocket.Conn {
	c.connMu.RLock()
	defer c.connMu.RUnlock()

	return c.conns[id]
}

func (c *ConnectionPool) SendMessage(msg Message) error {
	conn := c.getConnection(msg.RecipientId)
	if conn == nil {
		return ErrNotFound
	}

	return conn.WriteJSON(msg)
}
