package kafka

type Conf struct {
	Brokers []string `json:"brokers"`
	Version string   `json:"version"`
}

// ServerUrls 默认配置信息
var ServerUrls = map[string]Conf{}

// LoadConfig 加载默认配置信息
func LoadConfig(platform string, serverUrl string) {
	if _, exists := ServerUrls[platform]; !exists {
		ServerUrls[platform] = Conf{
			Brokers: []string{serverUrl},
			Version: "V2_0_0_0",
		}
	}
}
