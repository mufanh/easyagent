package global

import (
	"github.com/gorilla/websocket"
	"github.com/issue9/jsonrpc"
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
	agentInfo *model.AgentInfo
	conn      *websocket.Conn
	jConn     *jsonrpc.Conn
}

func setupServerRepository() *ServerRepository {
	return &ServerRepository{
		lock:     &sync.RWMutex{},
		Sessions: make(map[string]*SessionInfo),
	}
}

func (s *ServerRepository) SessionJConn(token string) *jsonrpc.Conn {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if info, ok := s.Sessions[token]; !ok {
		return nil
	} else {
		return info.jConn
	}
}

func (s *ServerRepository) AddSession(agentInfo *model.AgentInfo, conn *websocket.Conn, jConn *jsonrpc.Conn) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.Sessions[agentInfo.Token]; ok {
		return errors.New("添加连接失败，连接标识重复")
	}

	s.Sessions[agentInfo.Token] = &SessionInfo{agentInfo: agentInfo, conn: conn, jConn: jConn}
	return nil
}

func (s *ServerRepository) DeleteSession(token string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.Sessions[token]; !ok {
		return errors.New("连接不存在，删除失败")
	}

	conn := s.Sessions[token].conn
	if err := conn.Close(); err != nil {
		return err
	}
	delete(s.Sessions, token)
	return nil
}

func (s *ServerRepository) AgentInfos() []*model.AgentInfo {
	s.lock.RLock()
	defer s.lock.RUnlock()

	var r = make([]*model.AgentInfo, 0)
	for _, v := range s.Sessions {
		r = append(r, v.agentInfo)
	}
	return r
}
