package caches

import (
	"sync"
	"time"

	"github.com/go-redis/redis/v7"
)

// GetWagerRedisKey getWagerRedisKey
var GetWagerRedisKey = func(billNo string) string { // 生成redis的key
	return "wager:" + billNo
}

// StoredWagers
// step 1 => 以下数据保存在系统内存当中, 优先从这当中获取
// step 2 => 读取redis内存中的数据
// step 3 => 再查数据库
// 用于保存已成功处理(已结算)的订单号
var StoredWagers = struct {
	Data map[string]int64 // 最多保存10天
	Lock sync.Mutex
}{
	Data: map[string]int64{},
	Lock: sync.Mutex{},
}

// WagerCache 用于保存订单缓存
var WagerCache = struct {
	Save   func(*redis.Conn, ...string)
	Exists func(*redis.Conn, string) bool
}{
	Save: func(redisClient *redis.Conn, wagerIds ...string) {
		StoredWagers.Lock.Lock()
		defer StoredWagers.Lock.Unlock()

		currentTime := time.Now()
		for _, wagerId := range wagerIds {
			redisKey := GetWagerRedisKey(wagerId)                                // 将订单号码写入redis
			val := currentTime.Format("2006-01-02 15:04:05")                     // 订单写入时间
			_, _ = redisClient.Set(redisKey, val, time.Second*86400*10).Result() // 写入redis - 10 天超时
			StoredWagers.Data[wagerId] = time.Now().Unix()                       // 写入当前内部缓存
		}

		outTimestamp := currentTime.Unix() - 10*86400 // 超过10天
		for k, v := range StoredWagers.Data {         // 如果超过一定时间,  则从缓存当中删除
			if v < outTimestamp {
				delete(StoredWagers.Data, k)
			}
		}
	},
	Exists: func(redisClient *redis.Conn, wagerId string) bool {
		if _, exists := StoredWagers.Data[wagerId]; exists {
			return true
		}
		redisKey := GetWagerRedisKey(wagerId)
		res, err := redisClient.Get(redisKey).Result()
		if err != nil || err == redis.Nil { // 表示没有存在
			return false
		}
		if res == "" { // 表示没有找到
			return false
		}
		return true
	},
}
