package response

import (
	"sports-common/config"
	"sports-common/consts"
	"strings"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

// RespInfo 响应结构体
type RespInfo struct {
	ErrCode consts.ErrCode `json:"errcode"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
}

type ViewData = pongo2.Context

// 得到临时平台识别号
func getPlatform(c *gin.Context) string {
	path := c.Request.URL.Path
	pArr := strings.Split(path, "/")
	if len(pArr) < 2 {
		return ""
	}

	host := c.Request.Host
	if platform, exists := consts.PlatformUrls[host]; exists {
		return platform
	}
	return ""
}

// Ok 输出成功消息
func Ok(c *gin.Context) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Message string         `json:"message"`
	}{
		ErrCode: consts.Success,
		Message: "",
	})
}

// Err 输出错误信息
func Err(c *gin.Context, message string) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Message string         `json:"message"`
	}{
		ErrCode: consts.ErrorCommon,
		Message: message,
	})
	c.Abort()
}

// ErrWithCode 通过指定的错误代码，输出错误信息
func ErrWithCode(c *gin.Context, code consts.ErrCode) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Message string         `json:"message"`
	}{
		ErrCode: consts.ErrorCommon,
		Message: consts.ErrorInfo(code),
	})
	c.Abort()
}

// ErrCodeAndMsg 带错误码信息
func ErrCodeAndMsg(c *gin.Context, code consts.ErrCode, message string) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Message string         `json:"message"`
	}{
		ErrCode: code,
		Message: message,
	})
	c.Abort()
}

func ErrCodeAndData(c *gin.Context, code consts.ErrCode, data interface{}) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Data    interface{}    `json:"data"`
	}{
		ErrCode: code,
		Data:    data,
	})
	c.Abort()
}

func ErrCodeMsgAndData(c *gin.Context, code consts.ErrCode, msg string, data interface{}) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Message string         `json:"message"`
		Data    interface{}    `json:"data"`
	}{
		ErrCode: code,
		Message: msg,
		Data:    data,
	})
	c.Abort()
}

//errErrCode
//message
//data

// Message 输出消息
func Message(c *gin.Context, message string) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Message string         `json:"message"`
	}{
		ErrCode: consts.Success,
		Message: message,
	})
}

// ResultWithMsg 输出结果带信息
func ResultWithMsg(c *gin.Context, msg string, result interface{}) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Message string         `json:"message"`
		Data    interface{}    `json:"data"`
	}{
		ErrCode: consts.Success,
		Message: msg,
		Data:    result,
	})
}

func ResultAll(c *gin.Context, result interface{}) {
	c.JSON(0, result)
}

// Result 输出结果
func Result(c *gin.Context, result interface{}) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Message string         `json:"message"`
		Data    interface{}    `json:"data"`
	}{
		ErrCode: consts.Success,
		Message: "",
		Data:    result,
	})
}

// Pager 输出带分页的结果
func Pager(c *gin.Context, list interface{}, total uint64) {
	c.JSON(0, struct {
		ErrCode consts.ErrCode `json:"errcode"`
		Message string         `json:"message"`
		Data    interface{}    `json:"data"`
	}{
		ErrCode: consts.Success,
		Message: "",
		Data: struct {
			List  interface{} `json:"list"`
			Total uint64      `json:"total"`
		}{
			List:  list,
			Total: total,
		},
	})
}

// ErrorHTML 输出错误页面
func ErrorHTML(c *gin.Context, message ...string) {
	errorMessage := "好像出错了呢"
	viewData := pongo2.Context{}
	platform := getPlatform(c) // 多个平台可以共用一个静态文件服务
	//viewData["STATIC_URL"] = config.Get("config.static_url", "/static") // 全局静态URL
	viewData["STATIC_URL"] = func() string { // config.Get("config.static_url", "/static") // 全局静态URL
		if val := consts.Integrated.GetStaticURL(platform); val != "" {
			return val
		}
		return config.Get("config.static_url", "/static") // 全局静态URL
	}()
	if len(message) > 0 {
		errorMessage = message[0]
	}
	viewData["message"] = errorMessage
	c.HTML(200, "error.html", viewData)
}

// RenderString 输入纯字符串
func RenderString(c *gin.Context, data string) {
	c.String(200, data)
}

// Render 通用默认输出渲染
func Render(c *gin.Context, viewFile string, args ...pongo2.Context) {
	viewData := pongo2.Context{}
	if len(args) >= 1 {
		viewData = args[0]
	}
	platform := getPlatform(c)               // 多个平台可以共用一个静态文件服务
	viewData["STATIC_URL"] = func() string { // config.Get("config.static_url", "/static") // 全局静态URL
		if c.Request.Host == "bms.tjbet100.com" { // 特定情况下 - 备用域名指向特定static
			return "https://static.pusta.click/backend/static"
		}
		if val := consts.Integrated.GetStaticURL(platform); val != "" {
			return val + "/backend/static"
		}
		return config.Get("config.static_url", "/static") // 全局静态URL
	}()
	viewData["UPLOAD_URL"] = consts.Integrated.GetUploadURL(platform) // config.Get("config.upload_url", "/static") // 全局上传地址
	viewData["PLATFORM"] = platform
	for k, f := range templateFuncs {
		viewData[k] = f
	}
	if admin, exists := c.Get("_admin"); exists {
		viewData["ADMIN"] = admin
	}
	c.HTML(200, viewFile, viewData)
}
