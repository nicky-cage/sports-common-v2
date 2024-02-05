package tools

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TimeDebugEnable 是否启用时间调式
var TimeDebugEnable = false

// NowTime 标准获取当前时间 - 考虑时区因素
func NowTime() time.Time {
	_, _ = time.LoadLocation("Asia/Shanghai")
	return time.Now()
}

// CurrentTime 标准获取当前时间 - 考虑时区因素
func CurrentTime() time.Time {
	return NowTime()
}

// Now 标准获取当前时间
func Now() int64 {
	return NowTime().Unix()
}

// CurrentTimestamp 标准获取当前时间
func CurrentTimestamp() int64 {
	return NowTime().Unix()
}

// Unix 通过unix获取时间
func Unix(ts int64) time.Time {
	_, _ = time.LoadLocation("Asia/Shanghai")
	return time.Unix(ts, 0)
}

// GetDateTimeByTimeStamp 依据时间戳获取日期时间
func GetDateTimeByTimeStamp(ts int64) string {
	_, _ = time.LoadLocation("Asia/Shanghai")
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}

// GetTimestamp 得到时间戳
func GetTimestamp(dateStr string) int64 {
	return GetTimeStampByString(dateStr)
}

// GetTimeStampByDate 依据时间得到时间
func GetTimeStampByDate(dateStr string, args ...string) int64 {
	if len(args) >= 1 {
		return GetTimeStampByString(dateStr + " " + args[0])
	}
	return GetTimeStampByString(dateStr + " 00:00:00")
}

// GetTimeStampByString 通过字符串转化为时间戳: YYYY-mm-dd HH:MM::SS
func GetTimeStampByString(dateStr string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, loc)
	if err != nil {
		return 0
	}
	return theTime.Unix()
}

// GetTimeStampByString2 通过字符串转化为时间戳: mm-dd-YYYY HH:MM:SS
func GetTimeStampByString2(dateStr string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, err := time.ParseInLocation("01-02-2006 15:04:05", dateStr, loc)
	if err != nil {
		return 0
	}
	return theTime.Unix()
}

// GetTimeStampByString3 通过字符串转化为时间戳: YYYY/mm/dd HH:MM:SS
func GetTimeStampByString3(dateStr string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	strList := strings.Split(dateStr, " ")
	if len(strList) != 2 {
		return 0
	}
	if len(strList[1]) == 7 {
		dateStr = strList[0] + " 0" + strList[1]
	}
	theTime, err := time.ParseInLocation("2006/01/02 15:04:05", dateStr, loc)
	if err != nil { // log.Err(dateStr, err.Error())
		return 0
	}
	return theTime.Unix()
}

// GetTimeStampByString4 通过字符串转化为时间戳: mm-dd-YYYY HH:MM:SS(小数)
func GetTimeStampByString4(dateStr string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, err := time.ParseInLocation("01-02-2006 15:04:05.000000", dateStr, loc)
	if err != nil {
		return 0
	}
	return theTime.Unix()
}

// GetTimeStampByString5 通过字符串转化为时间戳: dd-mm-YYYY HH:MM:SS
func GetTimeStampByString5(dateStr string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, err := time.ParseInLocation("02-01-2006 15:04:05", dateStr, loc)
	if err != nil {
		return 0
	}
	return theTime.Unix()
}

// GetDateStringByTimeStamp 通过时间戳，获取日期
func GetDateStringByTimeStamp(timeUnix int64, need string) string {
	return time.Unix(timeUnix, 0).Format("2006" + need + "01" + need + "02")
}

// GetDateOnlyHisStringByTimeStamp 获取日期
func GetDateOnlyHisStringByTimeStamp(timeUnix int64, need string) string {
	return time.Unix(timeUnix, 0).Format("15" + need + "04" + need + "05")
}

// GetDateHisStringByTimeStamp 通过时间戳，获取日期时分秒
func GetDateHisStringByTimeStamp(timeUnix int64, need string) string {
	return time.Unix(timeUnix, 0).Format("2006" + need + "01" + need + "02 15:04:05")
}

// GetDateHisStringByTimeStampT 获取日期
func GetDateHisStringByTimeStampT(timeUnix int64, need string) string {
	return time.Unix(timeUnix, 0).Format("2006" + need + "01" + need + "02\\T15:04:05")
}

// GetDateHisStringByTimeStampTnot 获取日期
func GetDateHisStringByTimeStampTnot(timeUnix int64, need string) string {
	return time.Unix(timeUnix, 0).Format("2006" + need + "01" + need + "02T15:04:05")
}

// GetDateHisStringByTimeStampTZ 以时区获取
func GetDateHisStringByTimeStampTZ(timeUnix int64, need string) string {
	return time.Unix(timeUnix, 0).Format("2006" + need + "01" + need + "02\\T15:04:05\\Z")
}

// GetDateHisStringByTimeStampTZnot 以时区获取
func GetDateHisStringByTimeStampTZnot(timeUnix int64, need string) string {
	return time.Unix(timeUnix, 0).Format("2006" + need + "01" + need + "02T15:04:05Z")
}

// Timestamp 时间戳
func Timestamp() int64 {
	return NowTime().Unix()
}

// Date 得到当前的年/月/日
func Date() (int, int, int) {
	_, _ = time.LoadLocation("Asia/Shanghai")
	now := NowTime()
	return DateOf(now)
}

// DateOf 得到当前的年/月/日
func DateOf(now time.Time) (int, int, int) {
	ymd := now.Format("2006-01-02")
	ymdArr := strings.Split(ymd, "-")
	year, _ := strconv.Atoi(ymdArr[0])
	month, _ := strconv.Atoi(ymdArr[1])
	day, _ := strconv.Atoi(ymdArr[2])
	return year, month, day
}

// GetWeekStart 本周第一天
func GetWeekStart() time.Time {
	now := NowTime()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart
}

// GetMonthStart 得到本月第一天
func GetMonthStart() time.Time {
	now := NowTime()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	return time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
}

// FirstDayOfMon 当月第一天
func FirstDayOfMon(t time.Time) string {
	currentYear, currentMonth, _ := t.Date()
	currentLocation := t.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	return firstOfMonth.Format("2006-01-02")
}

// LastDayOfMon 当月最后一天
func LastDayOfMon(t time.Time) string {
	currentYear, currentMonth, _ := t.Date()
	currentLocation := t.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	return lastOfMonth.Format("2006-01-02")
}

// FirstDayOfLastMon 上月第一天
func FirstDayOfLastMon(t time.Time) string {
	year, month, _ := t.Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	firstDay := thisMonth.AddDate(0, -1, 0).Format("2006-01-02")
	return firstDay
}

// LastDayOfLastMon 上月最后一天
func LastDayOfLastMon(t time.Time) string {
	year, month, _ := t.Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	endDay := thisMonth.AddDate(0, 0, -1).Format("2006-01-02")
	return endDay
}

// GetTodayBegin 获取今天的零点时间
func GetTodayBegin() int64 {
	timeStr := NowTime().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t.Unix()
}

// GetTodayEnd 获取今天的最后结束时间
func GetTodayEnd() int64 {
	timeStr := NowTime().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	return t.Unix()
}

// GetDayBegin 某天开始时间
func GetDayBegin(dateStr string) int64 {
	return GetTimeStampByString(dateStr + " 00:00:00")
}

// GetDayEnd 某天最后时间
func GetDayEnd(dateStr string) int64 {
	return GetTimeStampByString(dateStr + " 23:59:59")
}

// TimeDebugBegin 在开始打印信息
func TimeDebugBegin(args ...interface{}) *time.Time {
	if !TimeDebugEnable {
		return nil
	}
	currentTime := NowTime()
	debugInfo := append([]interface{}{"[Debug][", currentTime.Format("2006-01-02 15:04:05"), "]"}, args...)
	fmt.Println(debugInfo...)
	return &currentTime
}

// TimeDebugAt 在某个时间点打印信息
func TimeDebugAt(lastTime *time.Time, args ...interface{}) *time.Time {
	if !TimeDebugEnable {
		return nil
	}

	currentTime := NowTime()
	past := currentTime.Sub(*lastTime)
	us := past.Microseconds()
	message := fmt.Sprintf("%.6f ms", float64(us)/1000.0)

	debugInfo := append([]interface{}{"[Debug][", currentTime.Format("2006-01-02 15:04:05"), fmt.Sprintf("][PAST:%s]", message)}, args...)
	fmt.Println(debugInfo...)
	return &currentTime
}

func NowMicro() int64 {
	return NowTime().UnixMicro()
}

// GetMicroTimeStampByString 通过字符串转化为时间戳: YYYY-mm-dd HH:MM::SS
func GetMicroTimeStampByString(dateStr string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, loc)
	if err != nil {
		return 0
	}
	return theTime.UnixMicro()
}

func MicroToSecond(micro int64) int64 {
	if micro < 1000000000*1000000 {
		return micro
	}
	return micro / 1000000
}

func SecondToMicro(second int64) int64 {
	if second > 1000000000*1000000 {
		return second
	}
	return second * 1000000
}
