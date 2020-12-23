package session

import "fmt"

// 中间件让用户选择使用哪个版本
var sessionMgr SessionMgr

func Init(provider string, addr string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Errorf("不支持的 session 类型")
		return
	}
	err = sessionMgr.Init(addr, options...)
}
