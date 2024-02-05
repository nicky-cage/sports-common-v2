package request

import (
	"sports-common/consts"

	"github.com/gin-gonic/gin"
)

// GetPlatform 获取平台标识号
func GetPlatform(c *gin.Context) string {

	if platformInf, exists := c.Get("__platform"); exists { // 先检查之前是不是被设置过
		return platformInf.(string)
	}

	host := c.Request.Host
	if platform, exists := consts.PlatformUrls[host]; exists {
		c.Set("__platform", platform)
		return platform
	}

	// 如果没有平台识别号，有可能是内部调用, 需要传post/get的platform字段
	platform := c.DefaultQuery("platform", "")
	if platform != "" {
		c.Set("__platform", platform)
		return platform
	}

	postedData := GetPostedData(c)
	if v, exists := postedData["platform"]; exists {
		platform := v.(string)
		c.Set("__platform", platform)
		return platform
	}

	panic("无法获取平台识别号, host = " + host + ", 请检查 xx_v2.sites 表相关配置")
}
