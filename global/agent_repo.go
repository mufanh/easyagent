package global

import (
	"github.com/gorilla/websocket"
	"github.com/issue9/jsonrpc"
	"sync"
)

var AgentRepo = setupAgentRepository()

type AgentRepository struct {
	transport jsonrpc.Transport
	conn      *websocket.Conn
	connected bool
	wg        *sync.WaitGroup
}

func setupAgentRepository() *AgentRepository {
	var wg sync.WaitGroup
	wg.Add(1)

	return &AgentRepository{
		connected: false,
		wg:        &wg,
	}
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

func (s *AgentRepository) SetConnected(connected bool) {
	s.connected = connected
}

func (s *AgentRepository) IsConnected() bool {
	return s.connected
}

func (s *AgentRepository) SetDone() {
	s.wg.Done()
}

func (s *AgentRepository) Wait() {
	s.wg.Wait()
}
