package redis

import (
	"fmt"
	"sports-common/consts"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v7"
)

// Config redis连接配置信息
type Config struct {
	Host     string
	Password string
	Port     int
	Db       int
}

// Servers 所有redis服务器 => map[平台识别号]缓存客户端
var Servers = map[string]*redis.Client{}

// Configs 所有redis配置信息 => map[平台识别号]配置信息
var Configs = map[string]*Config{}

// 实现连接管理 平台名称 => 唯一标识 => 连接
var connList = map[string]map[string]*redis.Conn{}

// 连接管理的锁
var connLock = sync.Mutex{}

// LoadConfig 加载配置 - password@host:port
func LoadConfig(platform string, config string) {
	if Configs[platform] != nil {
		return
	}
	host := "127.0.0.1"
	port := 6379
	password := ""
	db := 0
	info := config
	// 设定密码
	if strings.Index(info, "@") > 0 {
		uArr := strings.Split(info, "@")
		password = uArr[0]
		info = uArr[1]
	}

	if strings.Index(info, "/") > 0 {
		sArr := strings.Split(info, "/")
		if dbNumber, err := strconv.Atoi(sArr[1]); err == nil {
			db = dbNumber
		}
		info = sArr[0]
	}

	if strings.Index(info, ":") > 0 {
		sArr := strings.Split(info, ":")
		host = sArr[0]
		if portNumber, err := strconv.Atoi(sArr[1]); err == nil {
			port = portNumber
		}
	}

	conf := &Config{
		Host:     host,
		Password: password,
		Port:     port,
		Db:       db,
	}

	Configs[platform] = conf
}

// Connect 连接redis
func Connect(conf *Config) *redis.Client {
	addr := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	var client *redis.Client
	options := redis.Options{
		DB:          conf.Db,
		Addr:        addr,
		PoolSize:    2048,
		PoolTimeout: time.Minute * 3,
		IdleTimeout: time.Minute * 1,  // 空闲超过60s即自动断开
		MaxConnAge:  time.Second * 55, // 单个连接最大活路时间
	}
	if conf.Password != "" {
		options.Password = conf.Password
	}
	client = redis.NewClient(&options)
	if client == nil {
		fmt.Println("------------------- 加载redis服务出错 ----------------------------------")
	}
	return client
}

// GetConn 得到默认连接
func GetConn(platform string) *redis.Conn {
	connLock.Lock()
	defer connLock.Unlock()
	if list, exists := connList[platform]; exists {
		for k, v := range list {
			_, err := v.Ping().Result()
			if err != nil {
				fmt.Println("当前链接已失效: ", err)
				delete(list, k)
				continue
			}
			conn := v
			delete(list, k)
			return conn
		}
	}

	// 所有可用连接已经用完, 才会执行以下
	client := GetClientByPlatform(platform)
	_, err := client.Ping().Result()
	if err != nil {
		_ = client.Close()
		fmt.Println("------------------- *** 加载redis服务出错 *** ----------------------------------")
		fmt.Println("redis 连接失败 ==> ERROR: ", err, ", 正在尝试重新连接 ...")
		fmt.Println("------------------- *** 加载redis服务出错 *** ----------------------------------")
		client = GetClientByPlatform(platform, true)
		time.Sleep(time.Second * 2) // 连接时间
		return client.Conn()
	}

	return client.Conn()
}

// ReturnConn 归还连接
func ReturnConn(platform string, conn *redis.Conn) {
	connLock.Lock()
	defer connLock.Unlock()
	k := fmt.Sprintf("%p", conn)
	if len(connList[platform]) > 0 { // 如果不为0
		connList[platform][k] = conn
		return
	}
	connList[platform] = map[string]*redis.Conn{
		k: conn,
	}
}

// Default 获取默认的数据库连接
// func Default() *redis.Client {
// 	return GetClientByPlatform(consts.PlatformDefault)
// }

// GetClientByPlatform 依据平台识别号获取redis
func GetClientByPlatform(platform string, args ...bool) *redis.Client {
	var reConnect bool = false // 不需要重新连接
	if len(args) >= 1 {
		reConnect = args[0]
	}
	if !reConnect { // 如果不需要重新连接
		if client, exists := Servers[platform]; exists {
			if _, err := client.Ping().Result(); err != nil { // 如果ping失败则进行重连
				return GetClientByPlatform(platform, true)
			}
			return client
		}
	}

	// 如果没有连接，再检测配置是否存在
	if conf, exists := Configs[platform]; exists {
		if client := Connect(conf); client != nil {
			if consts.ResourceIsOnce { //用于区分定时任务
				Servers[platform] = client
			}
			return client
		}
	}

	// 打印错误信息
	panic(" *** redis没找到对应平台的信息，缺少平台识别号: [" + platform + "] *** ")
}
