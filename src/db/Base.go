package db

import (
	"fmt"
	"os"
	"sports-common/consts"
	"time"

	"github.com/go-xorm/xorm"
)

// Servers 所有数据库服务器
var Servers map[string]*xorm.EngineGroup

// Configs 数据库配置信息
var Configs map[string][]string

// IsDir 判断所给路径是否为文件夹
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 是否文件
func isFile(path string) bool {
	return !isDir(path)
}

// InitDbServers 初始化数据库服务器
// platform:  平台名称
// dbType: 数据加类型, 所有xorm支持的数据库类型都可以
// dataSources: 数据源信息
func InitDbServers(platform string, dbType string, dataSources []string) {
	if _, exists := Servers[platform]; exists { //如果已经存在则不处理
		return
	}
	engineGroup, err := xorm.NewEngineGroup(dbType, dataSources)
	if err != nil {
		fmt.Printf("[错误] 连接数据库失败: %v\n", err)
		return
	}

	setDbLogger := func() {
		fileDate := time.Now().Format("200601")
		fileDir := consts.LogPath + "/sql/" + fileDate
		if consts.LogPath == "./" {
			fileDir = consts.LogPath + "sql/" + fileDir
		}
		if !isDir(fileDir) {
			err := os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		fileName := "sql-" + fileDate + ".log"
		fileName = fileDir + "/" + fileName
		var logFile *os.File
		if !isFile(fileName) {
			logFile, err = os.Create(fileName)
			if err != nil {
				panic("无法创建日志文件: " + fileName)
			}
			err = os.Chmod(fileName, 0766)
			if err != nil {
				panic("无法修改文件权限: " + fileName)
			}
		}
		logFile, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
		if err != nil {
			return
		}
		logger := xorm.NewSimpleLogger(logFile)
		engineGroup.SetLogger(logger)
	}
	logDebug := false
	if logDebug {
		setDbLogger()
	}
	engineGroup.ShowExecTime(false)                 // 不显示执行时间
	engineGroup.ShowSQL(false)                      // 不显示sql语句
	engineGroup.SetMaxOpenConns(200)                // 最大空闲
	engineGroup.SetMaxIdleConns(20)                 // 最大
	engineGroup.SetConnMaxLifetime(time.Minute * 5) // 5分钟
	Servers[platform] = engineGroup
}

// LoadConfig 加载配置信息
func LoadConfig(platform string, conf []string) {
	Configs[platform] = conf
}

// Default 默认的db
// func Default() *xorm.EngineGroup {
//		if db, exists := Servers[consts.PlatformDefault]; exists {
//			return db
//		}
//		return nil
// }

// GetDbByPlatform 默认的db
func GetDbByPlatform(platform string, masterOnly ...bool) *xorm.Session {
	isMasterOnly := false
	if len(masterOnly) >= 1 {
		isMasterOnly = masterOnly[0]
	}

	reConnect := func() *xorm.Session {
		if conf, exists := Configs[platform]; exists {
			InitDbServers(platform, "mysql", conf)
			db := Servers[platform]
			if isMasterOnly {
				return db.Master().NewSession()
			}
			return db.NewSession()
		}
		return nil
	}

	if db, exists := Servers[platform]; exists { //拿到默认的值
		//if err := db.Ping(); err != nil {
		//	fmt.Println("-------------------- 数据库连接错误: engineGroup ---------------------")
		//	return reConnect()
		//}
		if isMasterOnly {
			return db.Master().NewSession()
		}
		return db.NewSession()
	}

	return reConnect()
}
