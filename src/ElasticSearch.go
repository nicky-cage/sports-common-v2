package common

import (
	"sports-common/es"

	"github.com/olivere/elastic/v7"
)

// ElasticSearch 依据平台得到客户端
func ElasticSearch(platform string) *elastic.Client {
	return es.GetClient(platform)
	//client, err := es.GetClientByPlatform(platform)
	//if err != nil {
	//	panic(err)
	//}
	//return client
}

// ElasticRestore
func ElasticRestore(platform string, client *elastic.Client) {
	es.ReturnClient(platform, client)
}
