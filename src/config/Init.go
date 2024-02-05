package config

import (
	"fmt"
	"os"
	"sports-common/consts"
	"sports-common/db"
	"sports-common/es"
	"sports-common/kafka"
	"sports-common/payment"
	"sports-common/redis"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"gopkg.in/ffmt.v1"
	"gopkg.in/ini.v1"
)

// LoadConfigs 加载配置信息
func LoadConfigs(configFile string) {
	if db.Servers != nil {
		return
	}
	if configFile == "" {
		configFile = "setting.ini"
	}
	cfg, err := ini.Load(configFile) //加载配置文件
	if err != nil {
		fmt.Printf("读取配置文件出错: %v\n", err)
		os.Exit(1)

	}
	Ini = cfg //设置全局Ini信息

	// ---------------------------------------------------------------------
	// 注意: 需要加什么额外的ini配置信息, 可以追加到这里面
	// 此段区域为除数据库/reids/mango之外的ini配置信息
	// 开始 ->
	// ---------------------------------------------------------------------
	consts.UploadPath = Get("upload.save_path")   //上传保存文件位置
	consts.UploadURLPath = Get("upload.url_path") //上传文件url前缀
	consts.IpDbPath = Get("ip.db_path")           //IP数据库保存路径
	consts.LogPath = Get("log.path")

	//-----------加载内网虚拟域名的配置------------
	consts.InternalMemberServUrl = Get("internal.internal_member_service")
	consts.InternalGameServUrl = Get("internal.internal_game_service")
	consts.InternalOssServUrl = Get("internal.internal_oss_service")
	consts.InternalAdminServUrl = Get("internal.internal_admin_service")

	//-----------加载kafka配置信息------------
	// consts.KafkaBrokerList = strings.Split(Get("kafka.broker_list_str"), ",")
	// fmt.Printf("KafkaBrokerList: %s\n", consts.KafkaBrokerList)
	// consts.KafkaTopicList = strings.Split(Get("kafka.topic_list_str"), ",")
	// consts.KafkaVersion = Get("kafka.version")

	// <- 结束
	// ---------------------------------------------------------------------

	// 全局变量设定
	// 设置运行模式, develop: 开发环境, release: 正式环境
	runMode := Get("sys.run_mode")
	if runMode == "develop" || runMode == "release" {
		consts.RunMode = runMode
	}
	if consts.RunMode == "release" {
		//gin.SetMode("release")
		gin.SetMode(gin.ReleaseMode)
	}
	// 设置运行模式
	if consts.AppName == "" {
		consts.AppName = Get("app_name") //应用程序名称
	}
	consts.PlatformIntegrated = Get("platform_integrated") //综合平台名称 - 总的平台名称
	//consts.PlatformDefault = Get("platform_default")       //设定默认平台 - 默认平台名称
	if Get("platform.custom_debug") != "" { //说明主的没配置，就不管
		consts.CustomDebug = Get("platform.custom_debug") == "1" //自定义debug是否开启
	}

	db.Servers = map[string]*xorm.EngineGroup{}     // 初始化 -> 所有数据库
	consts.PlatformUrls = map[string]string{}       // 域名 -> 平台识别号初始化
	consts.PlatformCodes = map[string]string{}      // 编码 -> 平台识别号初始化
	consts.PlatformStaticURLs = map[string]string{} // 静态域名 -> 平台识别号
	consts.PlatformUploadURLs = map[string]string{}

	// 加载总的平台信息
	// ----------------------->> 开始 <<----------------------------------------------
	integratedPlatform := Get("platform.name") // 平台名称
	dbHost := Get("platform.host")             // 主机地址
	dbUser := Get("platform.user")             // 用户名称
	dbPassword := Get("platform.password")     // 默认密码
	dbName := Get("platform.database")         // 数居库名
	dbPort := Get("platform.port")             // 默认端口
	connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)
	dataSources := []string{ // 默认平台数据源
		connString,
	}
	cacheConnString := Get("platform.redis_host") // 默认平台的redis配置信息

	if consts.RunMode != "release" {
		integratedConfig := map[string]interface{}{
			"db_source": dataSources,
			"redis":     cacheConnString,
		}
		_, _ = ffmt.Puts(map[string]map[string]interface{}{
			integratedPlatform: integratedConfig,
		})
	}
	redis.LoadConfig(integratedPlatform, cacheConnString)
	db.InitDbServers(integratedPlatform, "mysql", dataSources) //初始化默认的系统平台数据库
	// ----------------------->> 结束 <<----------------------------------------------

	// 因为默认只有一个平台级的数据库, 所以在此只取1个, 需要注意:
	// 1. 以下数据库连接信息是从系统平台的数据库当中读取
	// 2. 在系统平台库的 site_configs 表
	dbDefault := db.Servers[integratedPlatform]
	confs := GetPlatformConfigs(dbDefault)

	if consts.RunMode != "release" {
		fmt.Println("---------- platforms - site - configs ----------")
		_, _ = ffmt.Puts(confs)
		fmt.Println("---------- platforms - codes ----------")
		_, _ = ffmt.Puts(consts.PlatformCodes)
		fmt.Println("---------- platforms - urls ----------")
		_, _ = ffmt.Puts(consts.PlatformUrls)
		fmt.Println("---------- platforms - upload urls ----------")
		_, _ = ffmt.Puts(consts.PlatformUploadURLs)
		fmt.Println("---------- platforms - static urls ----------")
		_, _ = ffmt.Puts(consts.PlatformStaticURLs)
	}

	if len(confs) == 0 {
		panic("未加载到任何平台信息")
	}

	// 读取所有平台配置信息并加载至缓存
	consts.DbServerUrls = map[string][]string{}          // mysql 数据库信息
	consts.RedisServerUrls = map[string]string{}         // redis 配置
	consts.ElasticSearchServerUrls = map[string]string{} // elasticsearch
	consts.KafkaServerUrls = map[string]string{}         // kafka

	for _, conf := range confs {
		if platform, exists := conf["platform"]; exists {
			// 加载mysql信息
			if connString, exists := conf["conn_strings"]; exists {
				//fmt.Printf("[平台: %v] 加载Mysql配置信息 ...\n", platform)
				var connStrs []string
				//site_config的多个mysql配置请以逗号，隔开，并且不允许其他的逗号出现
				if strings.Index(connString, ",") > 0 {
					connStrs = strings.Split(connString, ",")
					consts.DbHasRwSplit = true
				} else {
					connStrs = append(connStrs, connString)
					consts.DbHasRwSplit = false
				}
				consts.DbServerUrls[platform] = connStrs
				db.InitDbServers(platform, "mysql", connStrs) //默认使用mysql数据库
			}
			// 加载redis信息 => redis#2020Yb@156.227.26.69:6379
			if connString, exists := conf["redis_strings"]; exists {
				//fmt.Printf("[平台: %v] 加载Redis配置信息 ...\n", platform)
				consts.RedisServerUrls[platform] = connString
				redis.LoadConfig(platform, connString) //只加载配置文件, 懒连接模式, 需要连接时才真正连接
			}
			// 再加载es信息
			if connString, exists := conf["elastic_strings"]; exists { // 加载elastic
				//fmt.Println("[平台: ", platform, "] 加载elasticsearch信息 ...")
				consts.ElasticSearchServerUrls[platform] = connString
				es.LoadConfig(platform, connString)
			}
			// 现加载kafka信息
			if connString, exists := conf["kafka_strings"]; exists { // kafka
				//fmt.Println("[平台: ", platform, "] 加载Kafka信息: ...")
				consts.KafkaServerUrls[platform] = connString
				kafka.LoadConfig(platform, connString)
			}
			// 加载支付信息
			if connString, exists := conf["pay_strings"]; exists {
				//fmt.Println("[平台: ", platform, "] 加载支付相关信息: ...")
				payment.LoadConfig(platform, connString)
			}
		}
	}

	//之后可以存放到数据库中
	internalIpListStr := Get("platform.internal_ip_list")
	if internalIpListStr != "" {
		if strings.Index(internalIpListStr, ",") > 0 {
			consts.InternalIpList = strings.Split(",", internalIpListStr)
		} else {
			consts.InternalIpList = append(consts.InternalIpList, internalIpListStr)
		}
	}

	if TableFields == nil {
		//fmt.Printf("[全局] 加载数据库/表/字段信息 ...\n")
		LoadDbTableFields()
	}
}
