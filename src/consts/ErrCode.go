package consts

// ErrCode 错误代码
type ErrCode int

// 全局错误码
const (
	Success                ErrCode = 0      //成功
	ErrorGameGet           ErrCode = 1001   //游戏列表获取出错
	ErrorGameRegFail       ErrCode = 1002   //游戏场馆失败
	ErrorEgameCodeErr      ErrCode = 1003   //游戏场馆失败
	ErrorRecoveryNoBalance ErrCode = 1005   //游戏回收余额的时候，如果钱是0.00，直接返回特殊的错误，让回收过滤掉前台游戏的余额刷新
	ErrorIpBan             ErrCode = 2001   //ip禁止登录（nginx做的）统一ip禁止页面
	ErrorLoginExipre       ErrCode = 2002   //token(过期，单点登录的原因) 重定向到登录页面
	ErrorMaintain          ErrCode = 10003  //系统维护 前台统一使用 统一系统维护页面
	ErrorWagers            ErrCode = 5000   //采集数据出错
	ErrorIOSApp            ErrCode = 8000   //IOS特殊错误.
	ErrorInvalidParam      ErrCode = 10000  //参数错误
	ErrorCommon            ErrCode = 10001  //一般错误
	ErrorInternal          ErrCode = 10002  //内部错误
	ErrorLoginExpection    ErrCode = 10004  //登录异常
	ErrorNotHasPriv        ErrCode = 20001  //没权限
	ErrorPermissionDenied  ErrCode = 100000 //验证失败
	ErrorIPDenied          ErrCode = 100001 //ip禁止
	ErrorNoLogin           ErrCode = 40005  //没有登录
	AgentBanCode           ErrCode = 70001  //代理账号禁用
)

//
var errMap = map[ErrCode]string{
	Success:               "成功",
	ErrorGameGet:          "游戏列表获取出错",
	ErrorInvalidParam:     "参数错误",
	ErrorCommon:           "一般错误",
	ErrorInternal:         "内部错误",
	ErrorPermissionDenied: "验证失败",
	ErrorIOSApp:           "IOS错误",
	ErrorGameRegFail:      "游戏场馆开通失败",
	ErrorNotHasPriv:       "您尚未登录，请登录后重试！",
	ErrorEgameCodeErr:     "电子游戏编码错误",
	ErrorWagers:           "游戏场馆采集数据出错",
	ErrorIpBan:            "ip被禁止",
	ErrorLoginExipre:      "您尚未登录，请登录后再试",
	ErrorMaintain:         "系统维护",
}

// ErrorInfo 错误信息
func ErrorInfo(code ErrCode) string {
	return errMap[code]
}
