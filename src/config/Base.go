package config

import (
	"sports-common/consts"

	"gopkg.in/ini.v1"
)

// TableFields 数据库: 平台名称 - 表名称 - 字段
var TableFields map[string]map[string][]string

// Ini 配置信息
var Ini *ini.File

// Get 获取配置项信息, 注意: 此方法必须在 LoadConfigs() 方法之后执行, 否则将无法获取配置信息
// key: 配置项名称
// 顶级配置项获取: config.Get("app_name")
// 二级配置项获取: config.Get("database.name")
func Get(key string, args ...interface{}) string {
	defaultValue := ""
	if len(args) >= 1 {
		defaultValue = args[0].(string)
	}
	if Ini == nil {
		return defaultValue
	}

	val := getIniKey(key).String()
	if val == "" {
		return defaultValue
	}
	return val
}

// GetInt 得到Int的值
// 请依据自己需要自动对int进行转换
func GetInt(key string) int {
	if Ini == nil {
		return 0
	}
	v, err := getIniKey(key).Int()
	if err != nil {
		return 0
	}
	return v
}

// GetBool 得到true/false的值
func GetBool(key string) bool {
	if Ini == nil {
		return false
	}
	v, err := getIniKey(key).Bool()
	if err != nil {
		return false
	}
	return v
}

// EnvIsProduct 是生产环境
func EnvIsProduct() bool {
	return Get("sys.run_mode") == "release"
}

// EnvIsTest 是测试环境
func EnvIsTest() bool {
	return Get("sys.run_mode") == "test"
}

// EnvIsDevelop 是开发环境
func EnvIsDevelop() bool {
	return Get("sys.run_mode") == "develop"
}

// GetPlatformByURL 依据url得到平台名称
func GetPlatformByURL(url string) string {
	if v, exists := consts.PlatformUrls[url]; exists {
		return v
	}
	return ""
}
