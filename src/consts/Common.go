package consts

import (
	"fmt"
	"time"
)

//  TimeLayoutYmdHis 时间格式
const TimeLayoutYmdHis = "2006-01-02 15:04:05"

// TimeLayoutYmd 年月日格式
const TimeLayoutYmd = "2006-01-02"

const TimeLayoutYmdNoLine = "20060102"

// TimeBillLayoutYmd 账户取数
const TimeBillLayoutYmd = "060102150405"

// DATETIME 时间+时区 格式
const DATETIME = "2006-01-02T15:04:05+08:00"

// HeaderKeyToken 常量定义
const HeaderKeyToken = "Authorization"

const UserPrefix = "USER:"

func t() {
	fmt.Print(UserPrefix)
}

// TokenKeyDuration token-Key持续时间
const TokenKeyDuration time.Duration = 4800 //1.5个小时

// AccessToken 访问Token
const AccessToken = "Authorization"

//设备ID
const DeviceID = "DeviceID"

// CurUserInfo 当前用户信息
const CurUserInfo = "user_info"

// TokenExpireTime token过期时间
const TokenExpireTime = 60 * 60 * 3

// AdminLoginErrMax 后台登录最多错误次数
const AdminLoginErrMax = 5

// AdminLoginLockSeconds 后台登录错误锁定时间(秒)
const AdminLoginLockSeconds = 3600

// DefaultExpiration 默认缓存时间
const DefaultExpiration = time.Duration(0)

// ForeverExpiration 永久缓存时间
const ForeverExpiration = time.Duration(-1)

const Limit = 50

const UploadFile = "upload:file:"

const IncrUserId = 9999999

/***************************************************/
// 0:主平台 1:体育 2:电竞 3:真人 4:电游 5:捕鱼 6:彩票 7:棋牌
const (
	SelfType = iota
	SportsType
	EsportsType
	LiveType
	EgameType
	FishType
	LotteryType
	ChessCardType
)

//游戏转账的流向
const (
	TranstypeInToGame    = 1 //转入游戏
	TranstypeOutFromGame = 2 //转出游戏
)

//代理相关
const (
	AgentBaseCode = 2000
)
