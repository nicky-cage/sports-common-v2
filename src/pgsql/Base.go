package pgsql

import (
	"context"
	"fmt"
	"math/rand"
	"sports-common/config"
	"sports-common/consts"
	"strings"

	"github.com/go-pg/pg/v10"
)

// ServerURLs 服务器连接信息
var ServerURLs map[string][]string

// Servers 服务器连接 -> 平台识别号/数据库
var Servers map[string][]*pg.DB

// LoadConfigs 加载配置文件
func LoadConfigs() {
	if len(ServerURLs) > 0 {
		return
	}

	// 格式: "postgres://username:password@host/database?sslmode=disable"
	connStrings := config.Get("pgsql.conn_strings")
	cArr := strings.Split(connStrings, ",")
	if len(cArr) == 0 {
		return
	}

	ServerURLs = map[string][]string{}
	Servers = map[string][]*pg.DB{}

	// 目前只有 sports_xx 库
	for _, platform := range consts.PlatformCodes {
		ServerURLs[platform] = []string{}
		Servers[platform] = []*pg.DB{} // 每个平台有 1-3台pg数据库
	}

	for _, platform := range consts.PlatformCodes {
		for _, v := range cArr {
			originString := strings.TrimSpace(v)
			if originString == "" {
				continue
			}

			database := "sports_" + platform
			connString := strings.ReplaceAll(originString, "_database", database) // 替换为真正的数据库连接
			opt, err := pg.ParseURL(connString)
			if err != nil {
				fmt.Println("加载PG出错:", err)
				panic(err)
			}
			db := pg.Connect(opt)
			if db == nil {
				panic("建拉PG连接有误")
			}
			cxt := context.Background()
			if err := db.Ping(cxt); err != nil {
				panic("建立PG连接出错:" + err.Error())
			}
			Servers[platform] = append(Servers[platform], db)
			ServerURLs[platform] = append(ServerURLs[platform], connString)
		}
	}
}

// Reconnect 重连
func Reconnect(platform string) {
	Servers[platform] = []*pg.DB{}
	cArr := ServerURLs[platform]
	for _, v := range cArr {
		connString := strings.TrimSpace(v)
		if connString == "" {
			continue
		}
		opt, err := pg.ParseURL(connString)
		if err != nil {
			fmt.Println("加载PG出错:", err)
			panic(err)
		}
		db := pg.Connect(opt)
		if db == nil {
			panic("建拉PG连接有误")
		}
		cxt := context.Background()
		if err := db.Ping(cxt); err != nil {
			panic("建立PG连接出错:" + err.Error())
		}
		Servers[platform] = append(Servers[platform], db)
	}
}

// GetConnForReading 得到读服务器 - 返回随机一台服务器
func GetConnForReading(platform string) *pg.Conn {
	LoadConfigs() // 加载信息

	if pArr, exists := Servers[platform]; exists {
		length := len(pArr)
		if length > 0 {
			randIndex := rand.Intn(length)
			conn := pArr[randIndex].Conn()
			if conn == nil {
				fmt.Println("未连接成功或连接已经断开, 尝试重新连接 ...")
				Reconnect(platform)
				conn = Servers[platform][rand.Intn(length)].Conn()
			}
			ctx := context.Background()
			if err := conn.Ping(ctx); err != nil {
				fmt.Println("无法ping通, 可能连接已经断开, 尝试重新连接 ...")
				Reconnect(platform)
				conn = Servers[platform][rand.Intn(length)].Conn()
			}
			return conn
		}
	}

	return nil
}

// GetConn 得到读取的连接
func GetConn(platform string) *pg.Conn {
	return GetConnForReading(platform)
}

// CloseConn 关闭连接
func CloseConn(pgConn *pg.Conn) {
	_ = pgConn.Close()
}

// GetConnGroup 得到写服务器组
func GetConnGroup(platform string) *ConnGroup {
	LoadConfigs()

	var cArr []*pg.Conn
	if pArr, exists := Servers[platform]; exists {
		for _, p := range pArr {
			cArr = append(cArr, p.Conn())
		}
	}

	if len(cArr) == 0 { // 没有任何连接
		fmt.Println("--- 没有任何 pgsql 连接 ---")
	}

	return &ConnGroup{
		Platform: platform,
		Conns:    cArr,
	}
}
