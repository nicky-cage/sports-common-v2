package config

import (
	"fmt"
	"sports-common/consts"
	"sports-common/db"
	"strings"

	"github.com/go-xorm/xorm"
)

// GetPlatformConfigs 得到平台-站点-配置相关信
/*
	生成如下数据结构:  []map[string]string, 有多少个站点 ,就有多少个map做为[]的item
	[
		{
			"platform":      	"integrated-sports" // 平台识别号
			"code": 			"TJ", // 平台代码
			"conn_strings":  	"user:pass@tcp(127.0.0.1:3306)/sports_db?charset=utf8" // mysql
			"db_type":       	"mysql" // 默认为mysql
			"mongo_strings": 	"mongodb://156.227.26.65:36017" // 已经不用
			"redis_strings": 	"127.0.0.1:6379" // redis
			"site_name":     	"天际体育", 		// 站点名称
			"pay_strings": 		"appid:key", 	// 支付信息
			"static_url": 		"", 			// 静态文件
			"upload_url": 		"", 			// 上传目录
		}
	]
*/
func GetPlatformConfigs(db *xorm.EngineGroup) []map[string]string {
	// fmt.Printf("加载平台/盘口/配置信息 ...\n")
	var data []map[string]string
	session := db.NewSession()
	platforms, err := session.QueryString("SELECT id FROM platforms WHERE status = 2")
	if err != nil || len(platforms) == 0 {
		panic("无法获取平台列表信息: " + err.Error())
	}
	for _, p := range platforms {
		sites, err := session.QueryString("SELECT id, name, urls, admin_url, code, api FROM sites WHERE platform_id = ? AND status = 2", p["id"])
		if err != nil || len(sites) == 0 {
			fmt.Println("无法获取平台(id: ", p["id"], ")下属盘口信息: ", err)
			continue
		}
		for _, s := range sites {
			adminUrl := s["admin_url"]
			platformCode := s["code"]
			urls := strings.Split(s["urls"], ",") // 折份前台域名
			urls = append(urls, strings.Split(s["api"], ",")...)

			sql := "SELECT name, value FROM site_configs WHERE platform_id = ? AND site_id = ? AND status = 2"
			conf := map[string]string{}
			confs, err := session.QueryString(sql, p["id"], s["id"])
			if err != nil || len(confs) == 0 {
				panic(fmt.Sprintf("无法获取平台(id: %v)/盘口(id: %v)配置信息: %v\n", p["id"], s["id"], err))
			}
			for _, c := range confs {
				name := c["name"]
				value := c["value"]
				conf[name] = value
				if name == "platform" {
					uArr := strings.Split(adminUrl, ",")
					for _, u := range uArr {
						domainName := strings.TrimSpace(u)
						if domainName != "" {
							consts.PlatformUrls[domainName] = value // 一个平台对应n个url -- 此时 value = 平台识别号
						}
					}
					consts.PlatformCodes[platformCode] = value // 一个平台对应一个code
					for _, url := range urls {                 // 前台域名, 可以有多个
						domainName := strings.TrimSpace(url)
						if domainName != "" {
							consts.PlatformUrls[domainName] = value // 多个url可以对应一个平台 -- 此时 value = 平台识别号
						}
					}
				}
			}
			if staticURL, exists := conf["static_url"]; exists { // 前台静态文件地址
				consts.PlatformStaticURLs[conf["platform"]] = staticURL
			}
			if uploadURL, exists := conf["upload_url"]; exists { // 上传路径
				consts.PlatformUploadURLs[conf["platform"]] = uploadURL
			}
			data = append(data, conf)
		}
	}

	return data
}

// LoadDbTableFields 加载所有数据库/表/字段信息
func LoadDbTableFields() {
	TableFields = map[string]map[string][]string{}
	for platform, db := range db.Servers {
		if result := GetTableFieldsFromDb(platform, db); result != nil {
			TableFields[platform] = result
		}
	}
}

// GetTableFieldsFromDb 得到此数据库的所有表/字段
func GetTableFieldsFromDb(platform string, db *xorm.EngineGroup) map[string][]string {
	db.ShowSQL(false)  //关闭sql显示
	databaseName := "" // 得到当前数据库名称
	sql := "SELECT DATABASE() AS database_name"
	rows, err := db.QueryString(sql)
	if err != nil {
		fmt.Printf("获取数据库名称出错: %v\n", err)
		return nil
	}
	databaseName = rows[0]["database_name"]

	// 先得到所有表
	sql = "SELECT DISTINCT TABLE_NAME AS table_name FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = '" + databaseName + "'"
	tableNames := []string{}
	rows, err = db.QueryString(sql)
	if err != nil {
		fmt.Printf("获取所有表信息出错: %v\n", err)
		return nil
	}
	for _, r := range rows {
		tableNames = append(tableNames, r["table_name"])
	}

	// 初始化数库/表/字段信息
	tableFields := map[string][]string{}
	sql = "SELECT TABLE_NAME AS table_name, COLUMN_NAME AS column_name FROM information_schema.COLUMNS " +
		"WHERE TABLE_NAME IN ('" + strings.Join(tableNames, "','") + "') AND TABLE_SCHEMA = '" + databaseName + "'"
	rows, _ = db.QueryString(sql)
	for _, r := range rows {
		tableName := r["table_name"]
		tableFields[tableName] = append(tableFields[tableName], r["column_name"])
	}
	return tableFields
}

// GetDbTableFields 得到此平台的所有数据表/字段信息
func GetDbTableFields(platform string) map[string][]string {
	if tableFields, exists := TableFields[platform]; exists {
		return tableFields
	}
	return nil
}

// GetTableFields 获取表/字段信息
func GetTableFields(platform string, table string) []string {
	if tableFields := GetDbTableFields(platform); tableFields != nil {
		if fields, exists := tableFields[table]; exists {
			return fields
		}
	}
	return nil
}
