package tools

import (
	"fmt"
	"hash/crc32"
	"sports-common/consts"
	"strconv"

	"github.com/go-redis/redis/v7"
)

// UserNameToIdHashNum 通过用户名按id取模分配的hash缓存，通过用户名来获取用户id
var UserNameToIdHashNum = 2000

// HashCodeByString 加码以字符串
func HashCodeByString(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

// SetUserNameToId 通过用户名的hashcode来写入到hash缓存
func SetUserNameToId(username string, userId uint64, rClient *redis.Conn) {
	hc := HashCodeByString(username)
	modId := hc % UserNameToIdHashNum                       //通过用户名按id取模分配的hash缓存，通过用户名来获取用户id //设置hash缓存
	cacheKey := consts.UserNameFindId + strconv.Itoa(modId) //设置hash缓存
	//写入的时候，可以不存在当前的cacheKey
	hasInt, err := rClient.HSet(cacheKey, username, strconv.Itoa(int(userId))).Result()
	if hasInt > 0 && err != nil {
		rClient.Expire(cacheKey, consts.ForeverExpiration)
	}
}

// GetIdByUserName 通过用户名获取hash的缓存的key来获取值
func GetIdByUserName(username string, rClient *redis.Conn) (uint64, error) {
	hc := HashCodeByString(username)
	modId := hc % UserNameToIdHashNum
	//通过用户名按id取模分配的hash缓存，通过用户名来获取用户id
	cacheKey := consts.UserNameFindId + strconv.Itoa(modId)
	isExists, err := rClient.HExists(cacheKey, username).Result()
	if err != nil {
		fmt.Println("无法获取用户ID: ", err)
		return 0, err
	}
	//也就是检查exists的时候，即使没值，err 是nil, isExists是0
	if !isExists {
		//fmt.Println("err 1")
		return 0, nil
	}

	//设置hash缓存
	v, err := rClient.HGet(cacheKey, username).Result()
	//这种直接取值，就要 注意err == redis.Nil
	if err != nil {
		return 0, err
	}
	if v == "" {
		return 0, nil
	}
	i, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return i, nil
}
