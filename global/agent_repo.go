package global

import (
	"github.com/gorilla/websocket"
	"github.com/issue9/jsonrpc"
)

var AgentRepo = setupAgentRepository()

type AgentRepository struct {
	transport jsonrpc.Transport
	conn      *websocket.Conn
}

func setupAgentRepository() *AgentRepository {
	return &AgentRepository{}
}

func (s *AgentRepository) SetConn(c *websocket.Conn) {
	s.conn = c
	s.transport = jsonrpc.NewWebsocketTransport(c)
}

func (s *AgentRepository) Conn() *websocket.Conn {
	return s.conn
}

func (s *AgentRepository) Transport() jsonrpc.Transport {
	return s.transport
}
