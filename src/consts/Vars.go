package consts

// AppName 应用程序名称
var AppName = ""

// RunMode 运行模式
// develop: 开发模式
// release: 生产环境
var RunMode = "develop"

// PlatformDefault 默认平台识别号
var PlatformDefault = ""

// PlatformIntegrated 综合平台(总台)
var PlatformIntegrated = ""

// PlatformUserNamePrefix 平台用户名称前缀
var PlatformUserNamePrefix = "tj"

// UploadPath 上传路径
var UploadPath = "./uploads"

// UploadURLPath 上传之后前缀
var UploadURLPath = "/uploads"

// LogPath
// elastic search 日志记录 最好带目录,项目自身配置目录，可以覆盖
// 默认当前目录
var LogPath = "./"

// DbHasRwSplit 是否开启读写分离,最终会通过配置来断定是否开启了读写分离
var DbHasRwSplit = true

// DbServerUrls mysql数据库配置
var DbServerUrls map[string][]string

// RedisServerUrls redis -> 平台识别号: redis连接信息
var RedisServerUrls map[string]string

// ElasticSearchServerUrls EsServerUrls
var ElasticSearchServerUrls map[string]string

// KafkaServerUrls kafka配置
var KafkaServerUrls map[string]string

//redis服务配置
//var RedisServerUrl string

// InternalIpList 内部IP
var InternalIpList []string

// CustomDebug 自定义debug开关
var CustomDebug = true

// FrontPageSize 每页分页的数量
var FrontPageSize = 20

// AdminLogHeaderEsIdName 后台日志得header的name
var AdminLogHeaderEsIdName = "admin_log_header_es_id"

// CenterCodeName 不能修改
var CenterCodeName = "CENTERWALLET"

// DisableSysMaintain 默认不禁用系统维护的标记
var DisableSysMaintain = false

// InternalGameServUrl 内部调用game 服务
var InternalGameServUrl = ""

// InternalMemberServUrl 内部调用member 服务
var InternalMemberServUrl = ""

// InternalOssServUrl 内部调用oss 文件上传 服务
var InternalOssServUrl = ""

// InternalAdminServUrl 内部调用后台地址
var InternalAdminServUrl = ""

//kafka-borker节点列表
// var KafkaBrokerList []string

// KafkaPartitionNum //kafka分区数量
var KafkaPartitionNum = 0

//kafka-topic节点列表
// var KafkaTopicList []string

// kafka 主题

// KafkaMaxLen 每次最大一次性数量是100
var KafkaMaxLen = 500

// KafkaVersion 版本
var KafkaVersion = ""

// WagersSettleArr 状态 0:未结算 1:赢 2:输 3:平局 4:取消(无效注单) 5:提前结算
var WagersSettleArr = []int{1, 2, 3}

// WagersNotSettleArr 注单未结算
var WagersNotSettleArr = []int{0, 5}

// WagersStatusArr 投注状态
var WagersStatusArr = []int{0, 1, 2, 3, 4, 5, 6, 7}

// WagersSports 原生体育注单列表的game_code,如BTI
var WagersSports = []string{
	"BTI-1",
}

// WagersSportsWallet 体育钱包
var WagersSportsWallet = []string{
	"BTI",
}

// GameTypes 游戏类型
var GameTypes = map[uint8]string{
	1: "彩票",
	2: "棋牌",
}

// IpDbPath ip数据库文件地址
var IpDbPath = ""

// NoCheckTranferGameCodeList 不能转账确认的单子
var NoCheckTranferGameCodeList = []string{
	"CQ9",
	"VR",
}

// LogTypes 日志类型: 0:普通操作 1:登录退出 2:权限相关 3:财务相关 4:内容相关 9:其他类型
var LogTypes = map[uint8]string{
	0: "普通操作",
	1: "退出登录",
	2: "权限相关",
	3: "财务相关",
	4: "内容相关",
	9: "其他类型",
}

// AdminLogTypes 管理员日志类型
var AdminLogTypes = struct {
	Normal     int
	InOut      int
	Permission int
	Finance    int
	Content    int
	Other      int
}{
	Normal:     0,
	InOut:      1,
	Permission: 2,
	Finance:    3,
	Content:    4,
	Other:      9,
}

// LogLevels 日志级别: 0:普通 1:重要 2:警告, 3:致命 4:错误 5:特殊 9:其他
var LogLevels = map[uint8]string{
	0: "普通",
	1: "重要",
	2: "警告",
	3: "致命",
	4: "错误",
	5: "特殊",
	9: "其他",
}

// AdminLogLevels 管理员日志级别
var AdminLogLevels = struct {
	Normal    int
	Important int
	Warn      int
	Fatal     int
	Error     int
	Special   int
	Other     int
}{
	Normal:    0,
	Important: 1,
	Warn:      2,
	Fatal:     3,
	Error:     4,
	Special:   5,
	Other:     9,
}

//系统代理
const (
	SysTestAgent     = "sys_test_agent"
	SysOfficialAgent = "sys_official_agent"
	SysAgentAgent    = "sys_agent_agent"
)

// ResourceIsOnce 资源写入map只是1次
var ResourceIsOnce = true

// ReplenishmentContent 补单结构
type ReplenishmentContent struct {
	GameCode  string
	StartTime string
	EndTime   string
}
