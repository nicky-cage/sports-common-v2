package tools

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
)

// 关于api相关用户信息
// https://openexchangerates.org/account/app-ids
// leon.leroy18@gmail.com
// qwe123QWE!@#
// 一天最多能调用 30 次
// https://openexchangerates.org/api/latest.json?app_id=d304f81fbbf043f8ab4f423b13fffb2d

// ExchangeResult {
//     "success": "1",
//     "result": {
//         "status": "ALREADY",
//         "scur": "USD", /*原币种*/
//         "tcur": "CNY", /*目标币种*/
//         "ratenm": "美元/人民币",
//         "rate": "6.5793", /*汇率结果(保留6位小数四舍五入) */
//         "update": "2016-06-24 08:30:37" /*数据更新时间*/
//     }
// }
//
type ExchangeResult struct {
	Desc      string             `json:"disclaimer"`
	License   string             `json:"license"`
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

// ExchangeRequestUrl 请求地址
var ExchangeRequestUrl = "https://openexchangerates.org/api/latest.json?app_id=d304f81fbbf043f8ab4f423b13fffb2d"

// ExchangeRateInfo 美元兑换人民币
var ExchangeRateInfo = struct {
	Rate     float64 // 美元兑换人民币汇率
	LastTime int64   // 上次获取时间
	Locker   sync.Mutex
}{
	Rate:     0,
	LastTime: 0,
	Locker:   sync.Mutex{},
}

// GetExchangeRate 获取美元兑换人民币汇率
func GetExchangeRate() float64 {

	currentTime := Now()
	if currentTime-ExchangeRateInfo.LastTime < 3600*20 { // 5分钟更新一次
		return ExchangeRateInfo.Rate
	}

	ExchangeRateInfo.Locker.Lock()
	defer ExchangeRateInfo.Locker.Unlock()

	content, err := HttpGet(ExchangeRequestUrl)
	if err != nil {
		fmt.Println("获取汇率相关数据出错: ", err)
		return 0
	}

	er := ExchangeResult{}
	if err := json.Unmarshal([]byte(content), &er); err != nil {
		fmt.Println("转换汇率相关数据出错: ", err)
		return 0
	}

	val, exists := er.Rates["CNY"]
	if !exists {
		fmt.Println("转换汇率相关数据出错: ", err)
		return 0
	}

	rate, err := strconv.ParseFloat(fmt.Sprintf("%.2f", val-0.02), 64) // 最小数后4位
	ExchangeRateInfo.Rate = rate                                       // 汇率 + 0.02
	ExchangeRateInfo.LastTime = Now()                                  // 最后访问时间

	return ExchangeRateInfo.Rate
}
