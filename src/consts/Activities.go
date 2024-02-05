package consts

const (
	ActivityTypeFirst    = 1 // 首存优惠
	ActivityTypeDeposit  = 2 // 存送优惠
	ActivityTypeRegister = 3 // 注册优惠
	ActivityTypeSave     = 4 // 救援奖金
	ActivityTypeVip      = 5 // VIP礼包
	ActivityTypeSpecial  = 6 // 特别优惠
	ActivityTypeInvite   = 7 // 邀请朋友优惠
	ActivityTypeDispatch = 8 // 人工派发奖金
	ActivityTypeBack     = 9 // 回归优惠
)

//  ActPltType 优惠场馆 类型
var ActPltType = map[string]string{
	"0": "TianJi优惠",
	"1": "体育优惠",
	"2": "电竞优惠",
	"3": "真人优惠",
	"4": "电游优惠",
	"5": "捕鱼优惠",
	"6": "彩票优惠",
	"7": "棋牌优惠",
	"8": "VIP优惠",
}

// ActivityType 优惠活动类型
var ActivityType = map[string]string{
	"1": "首存优惠",
	"2": "存送优惠",
	"3": "注册优惠",
	"4": "救援奖金",
	"5": "VIP礼包",
	"6": "特别优惠",
	"7": "邀请朋友优惠",
	"8": "人工派发奖金",
	"9": "回归优惠",
}

var HumanActivityType = map[int32]string{
	1: "限时活动",
	2: "新人首存",
	3: "日常活动",
	4: "体育优惠",
	5: "VIP礼包",
	6: "高额返水",
	7: "VIP特权",
}
var SpecialActivityListNew = `[{
	"id": 9,
	"type": 0,
	"activity_type": 1,
	"name": "欢乐时光",
	"title": "开赛巨献",
	"contents": "豪礼加持，首存就送",
	"details": "会员注册当日首存，达到以下条件即可获赠丰厚礼品，前往客服申请！",
	"rule": "首存达到申请要求请联系客服领取",
	"img_url": "/upload/images/activity/save_all03_new.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/YobetLea-App.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"items": [{
		"title_list": ["首存金额", "赠送礼品", "流水倍数"],
		"data_list": [
			["10万", "C罗或梅西亲笔签名球衣或球鞋", "1倍水"],
			["15万", "苹果手机 iPhone Xs Max 512GB", "1倍水"],
			["20万", "C罗或梅西亲笔签名球衣+球鞋+苹果手机 iPhone Xs Max 512GB", "1倍水"]
		]
	}]
}, {
	"id": 122,
	"type": 3,
	"activity_type": 6,
	"name": "",
	"title": "真人视讯 返水码上加码",
	"contents": "真人返水，史无前例，在业界最高真人洗码1.05%，返水无上限的基础上，额外在赠送0.5%的返水，总体返水额度高达1.55%；rn每天真人视讯投注20万以上，真人视讯有效投注额/(当日零时账户余额+当日存款)大于等于8，即可获取最高0.5%的额度返水奖励！rn额外返水奖励，彩金上限为18888元，1倍流水即可提现，且必须在次日12:00-24:00之间手动申请，否则视为放弃rn",
	"details": " 真人返水，史无前例，在业界最高真人洗码1.05%，返水无上限的基础上，额外在赠送0.5%的返水，总体返水额度高达1.55%；rn每天真人视讯投注20万以上，真人视讯有效投注额/(当日零时账户余额+当日存款)大于等于8，即可获取最高0.5%的额度返水奖励！rn额外返水奖励，彩金上限为18888元，1倍流水即可提现，且必须在次日12:00-24:00之间手动申请，否则视为放弃rn",
	"rule": " ",
	"img_url": "/upload/images/activity/ouzhoubeizhuli-700-500.jpg",
	"join_url": "liveCasinoRebate",
	"h5_img_url": "/upload/images/activity/ouzhoubei-app-promozhuli.png",
	"h5_join_url": "https:testh5.dixao.com/phonehtml/liveCasinoRebate/index.html",
	"multiple": 1,
	"active_url": "https:testh5.dixao.com/phonehtml/liveCasinoRebate/index.html",
	"items": [{
		"title_list": [],
		"data_list": []
	}]
}, {
	"id": 121,
	"type": 1,
	"activity_type": 6,
	"name": "",
	"title": "助力欧洲杯第二波,188888彩金存就送",
	"contents": "一重豪礼:累计存款送彩金，最高188888元等着您。二重豪礼:体育投注赢积分，400万现金等您提",
	"details": " 一重豪礼:累计存款送彩金，最高188888元等着您。二重豪礼:体育投注赢积分，400万现金等您提",
	"rule": " ",
	"img_url": "/upload/images/activity/ouzhoubeizhuli-700-500.jpg",
	"join_url": "freeCar",
	"h5_img_url": "/upload/images/activity/ouzhoubei-app-promozhuli.png",
	"h5_join_url": "https:www.yobet28.com/phonehtml/freeCar/index.html",
	"multiple": 1,
	"active_url": "https:www.yobet28.com/phonehtml/freeCar/index.html",
	"items": [{
		"title_list": [],
		"data_list": []
	}]
}, {
	"id": 120,
	"type": 1,
	"activity_type": 6,
	"name": "",
	"title": "挥别疫情-拥抱欧洲杯",
	"contents": "全体YOBET会员，在北京时间2月19号-3月3号内存款超过7天以上的即可享受浪漫欧洲3天2夜观赛游+英冠斯旺西城球星签名球衣，签名足球!",
	"details": "全体YOBET会员，在北京时间2月19号-3月3号内存款超过7天以上的即可享受浪漫欧洲3天2夜观赛游+英冠斯旺西城球星签名球衣，签名足球!",
	"rule": " ",
	"img_url": "/upload/images/activity/ouzhoubei-700-500.jpg",
	"join_url": "championsLeague",
	"h5_img_url": "/upload/images/activity/ouzhoubei-app-promo.png",
	"h5_join_url": "https:www.yobet28.com/phonehtml/championsLeague/index.html",
	"multiple": 1,
	"active_url": "https:www.yobet28.com/phonehtml/championsLeague/index.html",
	"items": [{
		"title_list": [],
		"data_list": []
	}]
}, {
	"id": 119,
	"type": 1,
	"activity_type": 6,
	"name": "",
	"title": "体育洗码王-输赢全返利",
	"contents": "",
	"details": "",
	"rule": " ",
	"img_url": "/upload/images/activity/tiyuximawang-700-500.jpg",
	"join_url": "sportWashingKing",
	"h5_img_url": "/upload/images/activity/tiyuximawang-app-promo.png",
	"h5_join_url": "https:www.yobet28.com/phonehtml/sportWashingKing/index.html",
	"multiple": 1,
	"active_url": "https:www.yobet28.com/phonehtml/sportWashingKing/index.html",
	"items": [{
		"title_list": [],
		"data_list": []
	}]
}, {
	"id": 118,
	"type": 0,
	"activity_type": 6,
	"name": "元宵喜乐会",
	"title": "元宵喜乐会",
	"contents": "元宵喜乐会，猜灯谜大赛，情人节浪漫礼＆驰援疫情送全体YOBET新老会员在2020年2月8号-2月14号【每日存款/每日有效投注额】= 3倍以上，即可享受288888元回馈！或者爱马仕手提包，赋予心爱的她",
	"details": "元宵喜乐会，猜灯谜大赛，情人节浪漫礼＆驰援疫情送全体YOBET新老会员在2020年2月8号-2月14号【每日存款/每日有效投注额】= 3倍以上，即可享受288888元回馈！或者爱马仕手提包，赋予心爱的她",
	"rule": " ",
	"img_url": "/upload/images/activity/yuanxiaohuanlehui-700-500.jpg",
	"join_url": "lanternFestival",
	"h5_img_url": "/upload/images/activity/yuanxiaohuanlehui-app-promo.png",
	"h5_join_url": "https:www.yobet28.com/phonehtml/lanternFestival/index.html",
	"multiple": 1,
	"active_url": "https:www.yobet28.com/phonehtml/lanternFestival/index.html",
	"items": [{
		"title_list": ["元宵喜乐会", "元宵喜乐会", "元宵喜乐会"],
		"data_list": [{
			"give": {
				"7": {
					"integral": 7,
					"gift": "玫瑰花一束10个",
					"bonus": 118
				},
				"15": {
					"integral": 15,
					"gift": "1壶食用油【金龙鱼】+一袋大米",
					"bonus": 288
				},
				"25": {
					"integral": 25,
					"gift": "Dior迪奥香水",
					"bonus": 388
				},
				"38": {
					"integral": 38,
					"gift": "N95口罩+30个【海外购】",
					"bonus": 888
				},
				"58": {
					"integral": 58,
					"gift": "潘多拉手链",
					"bonus": 1888
				},
				"68": {
					"integral": 68,
					"gift": "GUICC手提包",
					"bonus": 5888
				},
				"88": {
					"integral": 88,
					"gift": "苹果11PRo",
					"bonus": 9888
				},
				"98": {
					"integral": 98,
					"gift": "宝格丽手镯+纪梵希上衣",
					"bonus": 58888
				},
				"158": {
					"integral": 158,
					"gift": "劳力士手表",
					"bonus": 188888
				},
				"188": {
					"integral": 188,
					"gift": "Hermes 手提袋",
					"bonus": 288888
				}
			}
		}]
	}]
}, {
	"id": 117,
	"type": 0,
	"activity_type": 6,
	"name": "注册新人礼",
	"title": "首存三重送",
	"contents": "迎新礼-2020年2月3号之后注册的会员，即可申请首存，二存，三存，此项优惠活动！",
	"details": "迎新礼-2020年2月3号之后注册的会员，即可申请首存，二存，三存，此项优惠活动！",
	"rule": " ",
	"img_url": "/upload/images/activity/shoucunsanrensong-700-500.jpg",
	"join_url": "registrationCeremony",
	"h5_img_url": "/upload/images/activity/shoucunsanrensong-app.png",
	"h5_join_url": "https:www.yobet28.com/phonehtml/registrationCeremony/index.html",
	"multiple": 1,
	"active_url": "https:www.yobet28.com/phonehtml/registrationCeremony/index.html",
	"items": [{
		"title_list": ["1", "2", "3"],
		"data_list": [
			["1", "2", "3"]
		]
	}]
}, {
	"id": 116,
	"type": 0,
	"activity_type": 9,
	"name": "王者归来",
	"title": "老玩家-凯旋归来",
	"contents": "老玩家-凯旋回归",
	"details": "",
	"rule": " ",
	"img_url": "/upload/images/activity/wangzheguilai-pc-700-500.jpg",
	"join_url": "regress",
	"h5_img_url": "/upload/images/activity/wangzheguilai-app--app.png",
	"h5_join_url": "https:www.yobet28.com/phonehtml/regress/index.html",
	"multiple": 1,
	"active_url": "https:www.yobet28.com/phonehtml/regress/index.html",
	"items": [{
		"title_list": ["1", "2", "3"],
		"data_list": [
			["1", "2", "3"]
		]
	}]
}, {
	"id": 20,
	"type": 0,
	"activity_type": 6,
	"name": "独家直播",
	"title": "篮球季前赛",
	"contents": "您亏损，我买单",
	"details": "看球赛赢取保险奖上奖，您亏损，我买单最高3888元！",
	"rule": "1.参加活动会员需在东西部赛事篮球当天24小时内【北京时间】存款达600元或以上即可获得参加优惠资格。按照当日篮球亏损，第二日派发保险金额。rn2.VIP会员等级以赛事当日北京时间下午14:00后系统等级为准。rn3.以上负盈利，仅对已结算并产生输赢结果的投注额进行计算，任何平局、串关、特殊投注、取消的赛事将不计算在有效投注。负盈利计算仅限独赢，让球，大小，单双四个盘口的全场与半场。任何低于欧洲盘1.5或亚洲盘0.75水位的投注以及在同一赛事中同时投注对等盘口，将不计算在投注额内。rn4.本活动不与其他存送优惠共享（返水及存款优惠除外 ）。rn5.每位真实有效玩家/每一手机号码/电子邮箱/户籍地址/现居地址/同一银行卡/每一IP地址/每一台电脑或上网设备，每场赛事仅能参加并享受一次优惠活动，若有违规者，将不享受此红利。rn6.彩金仅需一倍流水即可出款。rn7.会员无需申请，满足申请条件的会员系统会在次日18：00之前进行派奖，请注意查收。rn8.若发现有套利客户，对赌或不诚实获取盈利之行为，将取消其优惠资格。rn9.此活动最终解释权归YOBET体育所有。",
	"img_url": "/upload/images/activity/NBA-PC.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/NBA-APP.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"sort": 0,
	"items": [{
		"title_list": ["等级会员", "最高返还金额", "彩金要求"],
		"data_list": [
			["VIP2及以下", "58", "1倍水"],
			["VIP3", "68", "1倍水"],
			["VIP4", "88", "1倍水"],
			["VIP5", "108", "1倍水"],
			["VIP6", "168", "1倍水"],
			["VIP7", "288", "1倍水"],
			["VIP8", "588", "1倍水"],
			["VIP9", "1888", "1倍水"],
			["VIP10", "3888", "1倍水"]
		]
	}]
}, {
	"id": 9,
	"type": 0,
	"activity_type": 6,
	"name": "欢乐时光",
	"title": "开赛巨献",
	"contents": "豪礼加持，首存就送",
	"details": "会员注册当日首存，达到以下条件即可获赠丰厚礼品，前往客服申请！",
	"rule": "首存达到申请要求请联系客服领取",
	"img_url": "/upload/images/activity/save_all03_new.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/YobetLea-App.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"items": [{
		"title_list": ["首存金额", "赠送礼品", "流水倍数"],
		"data_list": [
			["10万", "C罗或梅西亲笔签名球衣或球鞋", "1倍水"],
			["15万", "苹果手机 iPhone Xs Max 512GB", "1倍水"],
			["20万", "C罗或梅西亲笔签名球衣+球鞋+苹果手机 iPhone Xs Max 512GB", "1倍水"]
		]
	}]
}, {
	"id": 39,
	"type": 0,
	"activity_type": 6,
	"name": "欢乐时光",
	"title": "全平台首存优惠",
	"contents": "最高送1888元1倍流水",
	"details": "会员首存达到对应金额，即可领取相应奖金，奖金高达1888，仅需1倍流水！",
	"rule": "（与Yobet迎战五大联赛只能参与其一，必须通过离线存款方式参与，如有疑问，请联系客服）",
	"img_url": "/upload/images/activity/special_yobet.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/h5_special_yobet.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"items": [{
		"title_list": ["存款额度", "赠送红利", "流水倍数"],
		"data_list": [
			["200", "18", "1倍水"],
			["500", "38", "1倍水"],
			["1000", "58", "1倍水"],
			["5000", "188", "1倍水"],
			["10000", "288", "1倍水"],
			["50000", "1888", "1倍水"]
		]
	}]
}, {
	"id": 1,
	"type": 1,
	"activity_type": 6,
	"name": "欢乐时光",
	"title": "体育投注周",
	"contents": "周周有奖投越多送越多",
	"details": "会员于YOBET体育游戏中，按照北京时间计算当周有效投注达到以下等级者，即可申请，奖金高达8888元，仅需1倍水！",
	"rule": "每位会员每周内仅限申请一次，计算周期为北京时间。",
	"img_url": "/upload/images/activity/special_sport_zzyj.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/h5_special_sport_zzyj.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"items": [{
		"title_list": ["周有效投注", "可获彩金", "流水要求"],
		"data_list": [
			["20000+", "88", "1倍水"],
			["50000+", "288", "1倍水"],
			["100000+", "388", "1倍水"],
			["200000+", "588", "1倍水"],
			["500000+", "888", "1倍水"],
			["800000+", "1888", "1倍水"],
			["2000000+", "8888", "1倍水"]
		]
	}]
}, {
	"id": 2,
	"type": 2,
	"activity_type": 6,
	"name": "欢乐时光",
	"title": "电竞拯救日",
	"contents": "您亏损我买单",
	"details": "会员于YOBET于(IM电子竞技)游戏中，按照北京时间计算当日亏损达到500元或以上，即可申请，奖金高达500元，仅需3倍水！",
	"rule": "每位会员每天内仅限申请一次，计算周期为北京时间。",
	"img_url": "/upload/images/activity/special_esport_zjr.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/h5_special_esport_zjr.png",
	"h5_join_url": "",
	"multiple": 3,
	"active_url": "123",
	"items": [{
		"title_list": ["日亏损金额", "可获救援金比例", "流水要求"],
		"data_list": [
			["500+", "亏损金额的20%", "3倍水"],
			["1000+", "亏损金额的50%", "3倍水"]
		]
	}]
}, {
	"id": 3,
	"type": 3,
	"activity_type": 6,
	"name": "欢乐时光",
	"title": "真人流水送",
	"contents": "日奖金高达1888",
	"details": "会员于YOBET真人游戏中，按照北京时间计算当日有效投注达到以下等级者，即可申请，奖金高达1888元，仅需1倍水！",
	"rule": "每位会员每天内仅限申请一次，计算周期为北京时间。",
	"img_url": "/upload/images/activity/special_live_casino.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/h5_special_live_casino.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"items": [{
		"title_list": ["日有效投注", "可获彩金", "流水要求"],
		"data_list": [
			["60000+", "88", "1倍水"],
			["200000+", "188", "1倍水"],
			["500000+", "288", "1倍水"],
			["1588888+", "588", "1倍水"],
			["5888888+", "1888", "1倍水"]
		]
	}]
}, {
	"id": 4,
	"type": 4,
	"activity_type": 6,
	"name": "欢乐时光",
	"title": "电游拯救日",
	"contents": "亏损我买单",
	"details": "会员于YOBET于电子游戏中（所有捕鱼游戏除外），按照北京时间计算当日亏损达到500元或以上，即可申请，奖金高达88888元，仅需1倍水！",
	"rule": "每位会员每天内仅限申请一次，计算周期为北京时间。",
	"img_url": "/upload/images/activity/special_slot_game.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/h5_special_slot_game.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"items": [{
		"title_list": ["日亏损金额", "可获救援金比例", "流水要求"],
		"data_list": [
			["500+", "亏损金额的3%", "1倍水"],
			["3000+", "亏损金额的4%", "1倍水"],
			["8000+", "亏损金额的5%", "1倍水"],
			["50000+", "亏损金额的6%", "1倍水"]
		]
	}]
}, {
	"id": 5,
	"type": 7,
	"activity_type": 6,
	"name": "欢乐时光",
	"title": "棋牌对战周",
	"contents": "火利全开1888等你拿",
	"details": "会员于YOBET棋牌游戏中，按照北京时间计算当周有效投注达到以下等级者，即可申请，奖金高达1888元，仅需1倍水！",
	"rule": "每位会员每周内仅限申请一次，计算周期为北京时间。",
	"img_url": "/upload/images/activity/special_card_hlqk.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/h5_special_card_hlqk.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"items": [{
		"title_list": ["周有效投注", "可获彩金", "流水要求"],
		"data_list": [
			["58888+", "38", "1倍水"],
			["188888+", "88", "1倍水"],
			["588888+", "288", "1倍水"],
			["1888888+", "588", "1倍水"],
			["8888888+", "1888", "1倍水"]
		]
	}]
}, {
	"id": 6,
	"type": 5,
	"activity_type": 6,
	"name": "欢乐时光",
	"title": "捕鱼大作战",
	"contents": "奖金天天送",
	"details": "会员于YOBET于捕鱼游戏中，按照北京时间计算当日有效投注2000元或以上，即可申请，奖金高达28888，仅需1倍水！",
	"rule": "每位会员每天内仅限申请一次，计算周期为北京时间。",
	"img_url": "/upload/images/activity/special_fish_game.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/h5_special_fish_game.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"items": [{
		"title_list": ["日有效投注", "可获彩金", "流水要求"],
		"data_list": [
			["2000+", "28", "1倍水"],
			["8000+", "88", "1倍水"],
			["20000+", "188", "1倍水"],
			["50000+", "388", "1倍水"],
			["100000+", "888", "1倍水"],
			["8000000+", "5888", "1倍水"],
			["20000000+", "28888", "1倍水"]
		]
	}]
}, {
	"id": 7,
	"type": 7,
	"activity_type": 6,
	"name": "欢乐时光",
	"title": "棋牌对战日",
	"contents": "拯救不开心5888大放送",
	"details": "会员于YOBET棋牌游戏中，按照北京时间计算当日亏损达到2000或以上即可申请，奖金高达5888元，仅需1倍水！",
	"rule": "每位会员每天内仅限申请一次，计算周期为北京时间。",
	"img_url": "/upload/images/activity/special_card_zjbkx.png",
	"join_url": "",
	"h5_img_url": "/upload/images/activity/h5_special_card_zjbkx.png",
	"h5_join_url": "",
	"multiple": 1,
	"active_url": "123",
	"items": [{
		"title_list": ["日亏损金额", "可获彩金", "流水要求"],
		"data_list": [
			["2000+", "28", "1倍水"],
			["5000+", "58", "1倍水"],
			["10000+", "88", "1倍水"],
			["20000+", "188", "1倍水"],
			["50000+", "588", "1倍水"],
			["200000+", "1588", "1倍水"],
			["500000+", "2888", "1倍水"],
			["1000000+", "5888", "1倍水"]
		]
	}]
}]`
