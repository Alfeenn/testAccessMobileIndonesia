package model

import "github.com/gorilla/websocket"

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	From    string
	Type    string
	Message string
}

type WebSocketConnection struct {
	*websocket.Conn
	Username string
}

type WebSocketInterface interface {
	ConnectionSocket() []*WebSocketConnection
}

func NewSocketConnections() WebSocketInterface {
	return &WebSocketConnection{}
}

func (m *WebSocketConnection) ConnectionSocket() []*WebSocketConnection {
	var connections = make([]*WebSocketConnection, 0)
	return connections
}
