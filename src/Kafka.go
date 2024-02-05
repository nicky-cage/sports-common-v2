package common

import (
	"sports-common/consts"
	"sports-common/kafka"

	"github.com/Shopify/sarama"
)

// 所有平台数据库
func PlatformsProducer() map[string]sarama.AsyncProducer {
	arr := map[string]sarama.AsyncProducer{}
	for _, platform := range consts.PlatformCodes {
		arr[platform] = kafka.AsyncProducer(platform)
	}
	return arr
}

// 关闭所有session
func PlatformsProducerRestore(all map[string]sarama.AsyncProducer) {
	for _, producer := range all {
		producer.Close()
	}
}
