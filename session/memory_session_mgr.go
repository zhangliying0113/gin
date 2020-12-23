package session

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// 定义对象
type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

// 创建 session
func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	// go get github.com/satori/go.uuid
	// 用uuid作为sessionId
	id, err := uuid.NewV4()
	if err != nil {
		return
	}
	// 转string
	sessionId := id.String()
	// 创建个session
	session = NewMemorySession(sessionId)
	// 加到大 map
	s.sessionMap[sessionId] = session

	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
