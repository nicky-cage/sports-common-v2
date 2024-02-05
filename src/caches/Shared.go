package caches

import (
	"encoding/json"
	"sports-common/consts"
	"sync"

	"github.com/go-redis/redis/v7"
)

// GlobalCache 全局共享内存,以redis操作
type GlobalCache struct {
	Items  map[string]interface{}
	Locker sync.Mutex
}

// Global 所有的缓存保存
var Global = &GlobalCache{
	Items:  map[string]interface{}{},
	Locker: sync.Mutex{},
}

// Get 获取共局内存
// 先检查程序内存当中有没有, 如果有则返回
// 再检查redis当中有没有, 如果有则返回
// 如果以上都没有, 则执行函数, 获取结果, 并保存到redis及内存当中, 然后返回结果
func (ths *GlobalCache) Get(platform string, r *redis.Conn, key string, callback func() interface{}) interface{} {
	cacheKey := platform + "-" + key
	// 先在内存当中检查
	if val, exists := Global.Items[cacheKey]; exists {
		return val
	}
	// 再从redis当中检查
	var result interface{}
	strVal, err := r.Get(cacheKey).Result()
	if err == nil && strVal != "" {
		data := []byte(strVal)
		err := json.Unmarshal(data, &result)
		if err == nil { //如果从redis当中找到, 则直接返回
			return result
		}
		return nil
	}

	// 往内存和redis中增加
	Global.Locker.Lock()
	defer Global.Locker.Unlock()

	result = callback()                                          // 获取结果
	_, _ = r.Set(key, result, consts.GlobalCacheExpire).Result() // 写入redis

	Global.Items[cacheKey] = result //放入内存
	return result
}
