package log

/// 日志类型
const (
	TypeNormal     = 0 //普通
	TypeLogin      = 1 //登录
	TypePermission = 2 //权限
	TypeFinance    = 3 //财务
	TypeContent    = 4 //内容
	TypeOther      = 9 //其他
)

/// 日志级别
const (
	LevelNormal    = 0 //普通
	LevelImportant = 1 //重要
	LevelWarn      = 2 //警告
	LevelFault     = 3 //致命
	LevelError     = 4 //错误
	LevelSpecial   = 5 //特别
	LevelOther     = 9 //其他
)
