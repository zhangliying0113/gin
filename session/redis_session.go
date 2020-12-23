package session

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"sync"
)

type RedisSession struct {
	sessionId  string
	pool       *redis.Pool
	sessionMap map[string]interface{}
	rwlock     sync.RWMutex
	flag       int
}

// 用常量去定义状态
const (
	SessionFlagNone = iota
	SessionFlagModify
)

// 构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		sessionMap: make(map[string]interface{}),
		pool:       pool,
		flag:       SessionFlagNone,
	}
	return s
}

// session 存储到内存中的 map
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 设置值
	r.sessionMap[key] = value
	// 标记记录
	r.flag = SessionFlagModify
	return
}

// session 存储到 redis
func (r *RedisSession) Save() (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	if r.flag != SessionFlagModify {
		return
	}
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	// 获取 redis 连接
	coon := r.pool.Get()
	_, err = coon.Do("SET", r.sessionId, string(data))
	if err != nil {
		return
	}
	// 改状态
	r.flag = SessionFlagNone
	return
}

// 取 session
func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 先判断内存
	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key not exists")
	}
	return
}

// 从 redis 里再次加载
func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	// 转字符串
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	// 取到的东西反序列化到内存 map
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}
	return
}

func (r *RedisSession) Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}
