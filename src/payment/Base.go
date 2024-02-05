package payment

import "strings"

// PaymentConfig
type PaymentConfig struct {
	AppID     string
	AppSecret string
}

// Configs 支付配置信息
var Configs map[string]PaymentConfig

// LoadConfig 加载配置信息
func LoadConfig(platform string, conf string) {

	arr := strings.Split(conf, ":")
	if len(arr) != 2 {
		return
	}

	appID := arr[0]
	appSecret := arr[1]
	if appID == "" || appSecret == "" {
		return
	}

	if Configs == nil {
		Configs = map[string]PaymentConfig{
			platform: {
				AppID:     appID,
				AppSecret: appSecret,
			},
		}
		return
	}

	Configs[platform] = PaymentConfig{
		AppID:     appID,
		AppSecret: appSecret,
	}
}

// GetConfigByPlatform 依据平台识别号得到支付配置
func GetConfigByPlatform(platform string) *PaymentConfig {
	if val, exists := Configs[platform]; exists {
		return &val
	}
	// 默认使用天际体育的相关用户密码
	return &PaymentConfig{
		AppID:     "168000",
		AppSecret: "bAYZPYYpo2TN6vlNwqJfBGWGlNCDHUqP",
	}
}
