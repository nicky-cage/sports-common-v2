package tools

import (
	"math/rand"
	"time"
)

// Letters 默认的可用字符
var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012345678")

// RandInt64 随机int64
func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// RandString 生成随机字符串
func RandString(n int) string {
	b := make([]rune, n)
	total := len(Letters)
	for i := range b {
		b[i] = Letters[rand.Intn(total)]
	}
	return string(b)
}

// GenerateRangeNum 随机生成指定范围的数字比如6-9
func GenerateRangeNum(min, max int) int {
	if min > max {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}
