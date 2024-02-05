package redis

// DeleteMatchedKeys 删除匹配的键
func DeleteMatchedKeys(platform string, matchKey string) {
	conn := GetConn(platform)
	defer ReturnConn(platform, conn)

	keys, err := conn.Keys(matchKey).Result()
	if err != nil {
		return
	}

	conn.Del(keys...)
}
