package es

import (
	"context"
	"fmt"
	"sports-common/consts"
	"strings"
	"sync"

	"github.com/olivere/elastic/v7"
)

// 所有的 es 连接
// var Servers map[string]*elastic.Client

// ServerUrls
var ServerUrls map[string][]string

// 连接池
var clientList = map[string]map[string]*elastic.Client{}

// 锁
var clientLock = sync.Mutex{}

// 建立ES连接
// 格式: user:password@host:port
func LoadConfig(platform, connString string) {

	// 先判断是否已经存在
	if _, exists := ServerUrls[platform]; exists {
		return
	}

	// 判断是否有没可能是空值
	if len(ServerUrls) == 0 {
		ServerUrls = map[string][]string{}
	}

	confArr := strings.Split(connString, ",")
	connArr := []string{}
	connArr = append(connArr, confArr...)
	ServerUrls[platform] = connArr

	//fmt.Printf("[全局] elasticSearch 默认平台配置信息 平台名字：%s esUrl：%s ... \n", platform, connString)
	//fmt.Printf("EsServersUrl %v\n", ServerUrls)

	mappings, ok := consts.EsInitMappings["default"] // Get("elastic.es_platform_name")]
	if !ok {
		panic("elastic的配置平台 " + platform + " 对应的索引mapping不存在 ... \n")
	}

	client, err := GetClientByPlatform(platform)
	if err != nil {
		panic(err)
	}
	defer client.Stop()

	ctx := context.Background()
	for indexName, mapping := range mappings {
		realIndexName := platform + "_" + indexName
		exists, err := client.IndexExists(realIndexName).Do(ctx) // 判断索引是否存在, 如果不存在, 则自动创建
		if err != nil {
			panic(err)
		}
		if !exists {
			fmt.Println("索引 ", realIndexName, " 不存在, 创建索引 ... ")
			if _, err := client.CreateIndex(realIndexName).BodyString(mapping).Do(ctx); err != nil {
				panic(err)
			}
		}
	}
}
