package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sports-common/config"
	"sports-common/consts"
	"strings"
	"sync"

	"github.com/ipipdotnet/ipdb-go"
)

// IP_V6_URL 用于临时处理ipv6的数据库
const IP_V6_URL = "https://ip.zxinc.org/api.php?type=json&ip="

// 用于ipv6的临时处理数组
var ipV6Arr = map[string]string{}

// ZxincV6 从zxinc网址拿来下的数据结构
type ZxincV6 struct {
	Code int `json:"code"`
	Data struct {
		Country string `json:"country"`
		Ip      struct {
			End   string `json:"end"`
			Start string `json:"start"`
		} `json:"ip"`
		Local    string `json:"local"`
		Location string `json:"location"`
		Myip     string `json:"myip"`
	} `json:"data"`
}

// ip数据信息
type ipData = struct {
	locker sync.Mutex        //锁
	data   map[string]string //map[索引]map[ip]map[地区]频率
	max    int               //最大保存次数
}

// ip信息获取
var ipService = &ipData{
	locker: sync.Mutex{},
	data:   map[string]string{},
	max:    10240,
}

// IsIPv4 是否ipv4
func IsIPv4(ipAddress string) bool {
	ip := net.ParseIP(ipAddress)
	return ip != nil && strings.Contains(ipAddress, ".")
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(ipAddress string) bool {
	ip := net.ParseIP(ipAddress)
	return ip != nil && strings.Contains(ipAddress, ":")
}

// LoadIpV6Data 加载ipv6数据相关信息
func LoadIpV6Data() {
	if len(ipV6Arr) > 0 { // 如果已有读取过内存, 则跳过
		return
	}
	path := config.Get("ip.db_path_v6")
	fmt.Println("ipv6 path: ", path)

	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("读取文件内容有误: ", path)
		return
	}

	ipV6Arr = map[string]string{}
	lines := strings.Split(string(f), "\n")
	for _, v := range lines {
		tempArr := strings.Split(v, "|")
		if len(tempArr) >= 2 {
			ipV6Arr[tempArr[0]] = tempArr[1]
		}
	}
}

// GetIpV6AreaFromUrl 从url当中得到ipv6结果
func GetIpV6AreaFromUrl(ipAddress string) string {
	url := IP_V6_URL + ipAddress
	body, err := HttpGet(url)
	if err != nil {
		fmt.Println("获取ipv6结果出错: ", err)
		return ""
	}

	fmt.Println("ipv6 body: ", body)
	data := ZxincV6{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		fmt.Println("获取ipv6转化结果出错: ", err)
	}

	return data.Data.Location
}

// WriteIpV6Data 将ipv6数据写入到文件当中
func WriteIpV6Data(ipAddress, area string) {
	path := config.Get("ip.db_path_v6")
	fmt.Println("ipv6 path: ", path)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开以准备写入ipv6信息出错", err)
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.WriteString(ipAddress + "|" + area + "\n")
	writer.Flush()

	ipV6Arr[ipAddress] = area
}

// GetAreaByIp 依据IP得到地区
func GetAreaByIp(ipAddress string) string {
	ipService.locker.Lock()
	defer ipService.locker.Unlock()

	if v, exists := ipService.data[ipAddress]; exists { //先检查保存的是否有
		return v
	}

	if len(ipService.data) > ipService.max { //如果超出最大长度
		i := 0
		for k := range ipService.data {
			delete(ipService.data, k)
			i++
			if i >= 1024 { //只删除前面1024个
				break
			}
		}
	}

	if IsIPv4(ipAddress) {
		//设在ini文件设置绝对路径, 如 "D:/shares/go-v/src/sports-common/tools/ipip_v4.ipdb"
		if consts.IpDbPath == "" {
			panic("请设置IP数据库的文件地址")
		}
		ipDB, err := ipdb.NewCity(consts.IpDbPath)
		if err != nil {
			fmt.Printf("%v", err)
			return ""
		}
		areaString, err := ipDB.Find(ipAddress, "CN")
		if err != nil {
			fmt.Printf("%v", err)
			return ""
		}
		areaInfo := "[" + strings.Join(areaString, ",") + "]"
		ipService.data[ipAddress] = areaInfo
		return areaInfo
	} else if IsIPv6(ipAddress) { // 如果是ipv6
		LoadIpV6Data()
		if val, exists := ipV6Arr[ipAddress]; exists { // 如果有存在
			return val
		} else { // 如果不存在
			area := GetIpV6AreaFromUrl(ipAddress)
			if area != "" {
				WriteIpV6Data(ipAddress, area)
				return area
			}
		}
		return ""
	}
	return ""
}

// GetMillisecond 得到毫秒
func GetMillisecond() int64 {
	return NowTime().UnixNano() / 1e6
}
