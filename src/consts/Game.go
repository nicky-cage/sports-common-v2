package consts

// 场馆类型 0:主平台;1:体育;2:电竞;3:真人;4:电游;5:捕鱼;6:彩票;7:棋牌;
var GameVenueTypes = map[uint8]string{
	0: "主平台",
	1: "体育",
	2: "电竞",
	3: "真人",
	4: "电游",
	5: "捕鱼",
	6: "彩票",
	7: "棋牌",
}

// 游戏平台类型
var GamePlatformTypes = map[uint8]string{
	0: "PC",
	1: "H5",
	2: "APP",
}

// 展示类型
var GameDisplayTypes = map[uint8]string{
	0: "热门",
	1: "最新",
}

//玩家反馈问题类型

var FeedbackTypes = map[uint8]string{
	1: "存款问题",
	2: "提款问题",
	3: "游戏问题",
	4: "优惠问题",
	5: "网站/APP登录",
	6: "修改资料",
	7: "流水问题",
	8: "其他",
}

//广场栏目类型
var SportTypes = map[uint8]string{
	1: "足球",
	2: "篮球",
	3: "电竞",
	4: "其他",
}

//投注头庄
var BetStatus = map[int]string{
	0:  "未结算",
	1:  "赢",
	2:  "输",
	3:  "平局",
	4:  "取消(无效注单)",
	5:  "提前结算-未结算",
	6:  "提前结算-赢",
	7:  "提前结算-输",
	8:  "提前结算-平",
	9:  "赢半",
	10: "输半",
}

var ActivityTypes = map[int32]string{
	0: "-",
	1: "限时活动",
	2: "新人首存",
	3: "日常活动",
	4: "体育优惠",
	5: "高额返水",
	6: "VIP特权",
}

//游戏名后缀。  防止注册不了游戏账号 测试时为空
var GameExtraName = "1"
