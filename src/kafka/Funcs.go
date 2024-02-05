package kafka

import (
	"fmt"
	"sports-common/log"
	"time"

	"github.com/Shopify/sarama"
)

// BrokderList 得到平台相关配置
func BrokerList(platform string) []string {
	conf, exists := ServerUrls[platform]
	if !exists {
		fmt.Println("无法获取kafka配置信息")
		return nil
	}

	return conf.Brokers
}

//  AsyncProducer 生产者
func AsyncProducer(platform string) sarama.AsyncProducer {
	conf, exists := ServerUrls[platform]
	if !exists {
		fmt.Println("无法获取kafka配置信息")
		return nil
	}

	config := sarama.NewConfig()
	/*config.Net.TLS.Enable = true
	config.Net.TLS.Config = tlsConfig*/
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	//config.Producer.Return.Successes = true
	//config.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer(conf.Brokers, config)
	if err != nil {
		fmt.Println("获取kafka producer 失败: ", err.Error())
		log.Logger.Errorf("kakfa Failed to report wager start Sarama producer, consts.KafkaBrokerList: %s, err: %v", conf, err)
		return nil
	}
	return producer
}
