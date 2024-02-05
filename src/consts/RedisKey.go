package consts

import "time"

//注解： 必须有
//1.需要组装的是什么参数
//2.是是否有关后台
//3.数据是单条还是整体
//4.关联的表是什么
var (
	// GlobalCacheExpire 默认缓存超时时间, caches 模块使用
	GlobalCacheExpire = time.Hour * 24 * 30 //30天
	CachePrefix       = "ls:"
	// --
	// 所有的redis的key请带上ls:的前缀
	//

	// PwdError 密码错误,无关后台
	PwdError = "ls:pwd:error:"

	// PwdErrorExpire 过期时间，无关后台
	PwdErrorExpire = "ls:pwd:error:expire:"

	// Member 会员，无关后台
	Member = "ls:member:"

	// RedisKeyLoginUser 登录会员名称token，无关后台
	RedisKeyLoginUser = "ls:login:user:token:"
	//用户钱包信息 ，与user_id组装，有关后台只要 调用accout.ResetCacheData
	RedisKeyUserAccount = "ls:login:user:account:"

	//用户基本信息， 与user_id组装，有关后台修改后要删除对应的key
	CtxKeyLoginUser = "ls:login:user:"

	//获取游戏场馆和游戏列表的所有数据，有关后台只要删除该key  关于game_venues这个表 无关后台
	CkeyGameList = "ls:game:list:"

	//获取用户某个游戏场馆的本地记录，与game_code组装， 关于user_games这个表 无关后台
	CkeyUserGame = "ls:user:game:"

	//获取用户名下的所有开通了的游戏场馆列表，主要用于一键回收，拿到所有的game_code 关于user_games这个表 无关后台
	CkeyUserGameList = "ls:user:game:list:"

	//用户启动游戏的url,无关后台
	SportsLanunchUrl = "ls:sprot:lanun:"

	// CtxGameList 获取游戏平台
	//CtxGameList = "ls:game:list:"

	//邮箱验证码，无关后台
	BindEmail  = "ls:v_code:bd_email:"
	UpNewEmail = "ls:v_code:up_new_email:"
	//手机验证码，无关后台
	BindPhone  = "ls:v_code:bd_phone:"
	UpNewPhone = "ls:v_code:up_new_phone:"
	//修改邮箱，无关后台
	UpOldEmail = "ls:v_code:up_old_email:"
	//无关后台
	TransTypeCacheKey = "ls:trans:type:"
	//电子游戏列表
	EgameCacheKey = "ls:egame:list:gcode"

	//AgentLogin = "ls:agent_login:"

	//PinError = "ls:pin_error:"

	//无关后台
	UpOldPhone   = "ls:v_code:up_old_phone"
	FgPwdEmail   = "ls:v_code:fg_pwd_email:"
	FgPwdPhone   = "ls:v_code:fg_pwd_phone:"
	BindBankCard = "ls:v_code:bd_bank_card:"
	//验证设备或IP,无关后台
	VerifyDeviceIP = "ls:v_code:verify_device_ip:"
	//成为代理
	BecomeAgent = "ls:become_agent:"
	//ip注册，无关后台
	IPRegister = "ls:ip:register:"
	//验证码次数，无关后台
	PhoneVCodeCount   = "ls:phone:vcode:count:"
	EmailVCodeCount   = "ls:email:vcode:count"
	OfflineBanfChange = "ls:offline_bank_change:"

	//用户免转的时候，当前的gameCode值 无关后台
	MianZhuanKey = "ls:mianzhuan:"
	RiskBatchUp  = "ls:risk:dispatch:"

	// CacheKeyTokenPrefix 缓存token前缀 无关后台
	CacheKeyTokenPrefix = "token:"

	//username对应token，无关后台
	CacheKeyUserNameToken = "username:token:"

	//username对应user，无关后台
	CacheKeyUserNameUser = "username:user:"

	//登录密码验证是否通过，无关后台
	CacheKeyResetPwdVerify  = "reset_login_pwd:"
	CacheKeyModifyPwdVerify = "modify_login_pwd:"
	//用户名对应的用户id，无关后台
	UserNameFindId = "ls:username:to:uid"

	//用户余额缓存，无关后台
	GameBalanceKey = "ls:game:balance:"

	//轮播图有关后台,参数1,2全删, ad_carousels表
	BannerList = "banner:list:"
	//启动页广告, ad_lanuches表
	AdLaunchList = "ad:launch:list:"
	//用户余额，无关后台
	UserBalance = "user:balance:"
	//地区，无关后台
	PCDArea = "pcd:area"
	//有关后台，paramters表的key
	ParamCacheKey = "ls:parameters:key:"
	//二级密码
	WithdrawPhone = "username:withdraw:password"

	ReplenishmenKey = ":bet:replenishment:"
	//整天的时间戳字符串 可以用分号，带多个。
	AgDownLoadKey = ""
	LiveChannel   = "live_channel:"
	// 用于通知用户财务id-list
	KeyFinanceNotify = "fn:user_id_list"
)

var VerifyType = map[string]string{
	"4":  FgPwdPhone, // 获取忘记密码手机验证码
	"14": FgPwdEmail, // 获取忘记密码邮箱验证码
}

var VerifyTypePhone = map[string]string{
	"1": BindPhone,      // 获取绑定手机验证码
	"2": UpOldPhone,     // 获取旧手机验证码
	"3": UpNewPhone,     // 获取新手机验证码
	"4": FgPwdPhone,     // 获取忘记密码手机验证码
	"5": BindBankCard,   // 获取绑定银行卡手机验证码
	"6": VerifyDeviceIP, // 登录验证
	"7": BecomeAgent,    // 成为代理
	"8": WithdrawPhone,  // 资金密码
}

var VerifyTypeEmail = map[string]string{
	"11": BindEmail,  // 获取绑定邮箱验证码
	"12": UpOldEmail, // 获取旧邮箱验证码
	"13": UpNewEmail, // 获取新邮箱验证码
	"14": FgPwdEmail, // 获取忘记密码邮箱验证码
}

//联赛数据采集
const (
	FootballLeagueMatches   = "football_league_matches"
	BasketballLeagueMatches = "basketball_league_matches"
	FootballTextLive        = "football_text_live"
	BasketballTextLive      = "basketball_text_live"
	LiveList                = "live_list"
)
