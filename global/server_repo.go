package global

import (
	"github.com/gorilla/websocket"
	"github.com/mufanh/easyagent/internal/model"
	"github.com/pkg/errors"
	"sync"
)

var ServerRepo = setupServerRepository()

type ServerRepository struct {
	lock     *sync.RWMutex
	Sessions map[string]*SessionInfo
}

type SessionInfo struct {
	AgentInfo *model.AgentInfo
	Conn      *websocket.Conn
}

func setupServerRepository() *ServerRepository {
	return &ServerRepository{
		lock:     &sync.RWMutex{},
		Sessions: make(map[string]*SessionInfo),
	}
}

func (s *ServerRepository) AddSession(agentInfo *model.AgentInfo, conn *websocket.Conn) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.Sessions[agentInfo.Token]; ok {
		return errors.New("添加连接失败，连接标识重复")
	}

	s.Sessions[agentInfo.Token] = &SessionInfo{AgentInfo: agentInfo, Conn: conn}
	return nil
}

func (s *ServerRepository) DeleteSession(token string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.Sessions, token)
}

func (s *ServerRepository) ListAgentInfos() []*model.AgentInfo {
	s.lock.RLock()
	defer s.lock.Unlock()

	r := make([]*model.AgentInfo, len(s.Sessions))
	index := 0
	for _, v := range s.Sessions {
		r[index] = v.AgentInfo
		index += 1
	}
	return r
}
