package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sports-common/tools"
	"strings"

	"github.com/gin-gonic/gin"
)

// IsAjax 是否是ajax请求
func IsAjax(c *gin.Context) bool {
	return strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest"
}

// Get 提交get请求
func Get(queryURL string, data ...map[string]string) (string, error) {
	baseURL, _ := url.Parse(queryURL)
	if len(data) >= 1 {
		params := url.Values{}
		for k, v := range data[0] {
			params.Add(k, v)
		}
		baseURL.RawQuery = params.Encode()
	}
	realURL := baseURL.String()
	res, err := http.Get(realURL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Post 发送post请求
func Post(url string, data ...map[string]string) (string, error) {
	formData := ""
	if len(data) >= 1 {
		arr := []string{}
		for k, v := range data[0] {
			arr = append(arr, fmt.Sprintf("%s=%s", k, v))
		}
		formData = strings.Join(arr, "&")
	}
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(formData))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetTimesByQuery 得到时间区间
func GetTimesByQuery(c *gin.Context, field string) (int64, int64) {
	value, exists := c.GetQuery(field)
	if !exists {
		todayBegin := tools.GetTodayBegin()
		return todayBegin, todayBegin + 86399
	}
	areas := strings.Split(value, " - ")
	return tools.GetTimeStampByString(areas[0]), tools.GetTimeStampByString(areas[1])
}

// GetTimesByQuery 得到时间区间
func GetMicroTimesByQuery(c *gin.Context, field string) (int64, int64) {
	value, exists := c.GetQuery(field)
	if !exists {
		todayBegin := tools.GetTodayBegin()
		return tools.SecondToMicro(todayBegin), tools.SecondToMicro(todayBegin + 86399)
	}
	areas := strings.Split(value, " - ")
	return tools.GetMicroTimeStampByString(areas[0]), tools.GetMicroTimeStampByString(areas[1])
}

// IsExportExcel 是否导出文件
func IsExportExcel(c *gin.Context) bool {
	if exportExcel := c.DefaultQuery("export_excel", ""); exportExcel != "" {
		return true
	}
	return false
}

// GetLang 得到语言信息
func GetLang(c *gin.Context) {}

// GetFingerPrint 得到浏览器指纹
func GetFingerPrint(c *gin.Context) string {
	headers := c.Request.Header
	iArr := []string{}
	kArr := []string{
		"Accept",
		"Accept-Encoding",
		"Accept-Language",
		//"Cache-Control",
		//"Connection",
		"Http_x_forwarded_for",
		"Upgrade-Insecure-Requests",
		"User-Agent",
		"X-Forwarded-For",
		"X-Real-Ip",
	}

	for _, k := range kArr {
		if v, exists := headers[k]; exists {
			val := strings.Join(v, "##")
			iArr = append(iArr, val)
		}
	}

	info := strings.Join(iArr, "||")
	return tools.MD5(info)
}
