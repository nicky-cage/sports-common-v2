package consts

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// PlatformUrls 所有可被允许的访问的域名
// {
//  "admin.sports": "integrated-sports" // 域名: 平台识别号
// }
var PlatformUrls map[string]string

// PlatformCodes 所有可被允许的访问的域名
// {
//  "tj": "integrated-sports" // code: 平台识别号
// }
var PlatformCodes map[string]string

// PlatformStaticURLs 静态url -> 平台识别号
// {
//  "admin.sports": "integrated-sports" // 域名: 平台识别号
// }
var PlatformStaticURLs map[string]string

// PlatformUploadURLs 上传url -> 平台识别号
// {
//  "admin.sports": "integrated-sports" // 域名: 平台识别号
// }
var PlatformUploadURLs map[string]string

// Integrated 包网相关数据处理
var Integrated = struct {
	Allow             func(*gin.Context) bool    // 是否允许访问 - 是否有些盘口/网站
	AllowByURL        func(string) bool          // 依据url判断是否允许访问 - 是否有此平台/网站
	GetPlatformByCode func(string) string        // 依据编码得到平台识别号
	GetCodeByPlatform func(string) string        // 依据平台识别号得到编码
	HasCode           func(string) bool          // 判断是否有此识别号
	GetStaticURL      func(string) string        // 参数: 平台识别号
	GetUploadURL      func(string) string        // 台数: 平台识别号
	HasPlatform       func(platform string) bool // 是否有平台识别号
}{
	Allow: func(c *gin.Context) bool {
		url := c.Request.URL.Host
		_, exists := PlatformUrls[url]
		return exists
	},
	AllowByURL: func(url string) bool {
		_, exists := PlatformUrls[url]
		return exists
	},
	GetPlatformByCode: func(code string) string {
		if v, exists := PlatformCodes[strings.ToUpper(code)]; exists {
			return v
		}
		return ""
	},
	GetCodeByPlatform: func(platform string) string {
		for k, v := range PlatformCodes {
			if v == platform {
				return strings.ToUpper(k)
			}
		}
		return ""
	},
	HasCode: func(code string) bool {
		_, exists := PlatformCodes[strings.ToUpper(code)]
		return exists
	},
	HasPlatform: func(platform string) bool {
		_, exists := PlatformCodes[platform]
		return exists
	},
	GetStaticURL: func(platform string) string {
		for p, url := range PlatformStaticURLs {
			if p == platform {
				return url
			}
		}
		return ""
	},
	GetUploadURL: func(platform string) string {
		for p, url := range PlatformUploadURLs {
			if p == platform {
				return url
			}
		}
		return ""
	},
}
