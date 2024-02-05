package common

import (
	"sports-common/consts"
	redisCache "sports-common/redis"
	"strings"

	"github.com/go-redis/redis/v7"
)

// Redis 获得Redis
func Redis(platform string) *redis.Conn {
	return redisCache.GetConn(platform)
}

// RedisRestore 放回redis
func RedisRestore(platform string, conn *redis.Conn) {
	redisCache.ReturnConn(platform, conn)
}

// PlatformsRedis 所有redis
func PlatformsRedis() map[string]*redis.Conn {
	arr := map[string]*redis.Conn{}
	for _, platform := range consts.PlatformCodes {
		arr[platform] = Redis(platform)
	}
	return arr
}

// PlatformsRedisRestore 关闭所有redis
func PlatformsRedisRestore(all map[string]*redis.Conn) {
	for code, r := range all {
		platform := consts.Integrated.GetPlatformByCode(strings.ToLower(code))
		RedisRestore(platform, r)
	}
}
