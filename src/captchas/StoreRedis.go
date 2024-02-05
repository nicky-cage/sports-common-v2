package captchas

import (
	redisCache "sports-common/redis"
	"time"
)

// StoreRedis 生成图片保存redis
type StoreRedis struct {
	Platform string // 平台名称
}

// *** 警告 ***
// Set 设置redis的值
func (p *StoreRedis) Set(id string, value string) error {
	platform := p.Platform
    _, err := redisCache.GetConn(platform).Set(id, value, 300*time.Second).Result()
    return err
}

// *** 警告 ***
// Get 获取redis的值
func (p *StoreRedis) Get(id string, clear bool) string {
	platform := p.Platform
	conn := redisCache.GetConn(platform)
	v, _ := conn.Get(id).Result()
	if clear {
		conn.Del(id)
	}
	return v
}

// Verify 校验redis的值
func (p *StoreRedis) Verify(id, answer string, clear bool) bool {
	v := p.Get(id, clear)
	return v == answer
}
