package redis

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// 获取随机的字符串
func randString(n int) string {
	var Letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	b := make([]rune, n)
	total := len(Letters)
	for i := range b {
		b[i] = Letters[rand.Intn(total)]
	}
	return string(b)
}

// Lock 分布式锁
func Lock(platform string, idStr string) (string, error) {
	conn := GetConn(platform)
	defer ReturnConn(platform, conn)
	key := fmt.Sprintf("LOCK:%d:%s", time.Now().Unix(), idStr)
	val, err := conn.Incr(key).Result()
	if err != nil { // 操作出现问题
		return "", err
	}
	conn.Expire(key, time.Minute*3) // 设置自动超时
	if val > 1 {                    // 表示已经被锁定
		return "", errors.New("locked data")
	}
	return key, nil
}

// Unlock 解锁
func Unlock(platform, key string) {
	redis := GetConn(platform)
	defer ReturnConn(platform, redis)
	_ = redis.Del(key)
}
