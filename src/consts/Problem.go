package consts

type Question struct {
	Q string `json:"q"`
	A string `json:"a"`
}

var ProblemType = []map[string]interface{}{
	map[string]interface{}{
		"group": "常见问题",
		"itme": []Question{
			Question{
				Q: "1，我们代理后台如何登录",
				A: "安装代理后台APP后就是一个IOS原生新闻APP，专员会给您一个新闻APP安全码即可登录后台。",
			},
			Question{
				Q: "2，我们的安全应急机制",
				A: "为了您自身安全，必须使用IOS系统。当输入应急销毁码后进入后会显示空白页面同时闪退。此后无法登录后台，同时锁定您的账户，只能作为正常新闻APP使用。不过数据依然正常更新，但需要联系相关人员协助您解锁后才能正常登录。",
			},
		},
	},

	map[string]interface{}{
		"group": "佣金与盈利",
		"itme": []Question{
			Question{
				Q: "1，什么是净盈利？",
				A: "1.1净盈利＝所有产品的盈利金额－会员相关红利－会员存提手续费－平台费。1.2会员红利：客户相关优惠和返水。1.3如：平台内优惠活动以及代理在后台上分的活动。1.4会员存提手续费：根据不同的存款方式会有不同的费率产生。1.5平台费：平台费会以各个平台商的收费标准计算，月累计（根据各平台而定）",
			},
			Question{
				Q: "2，什么样的会员算活跃会员？",
				A: "活跃会员的标准：每月存款至少100元，并在网站相关游戏中至少一次的有效投注。",
			},
			Question{
				Q: "3，什么时候生成佣金报表？",
				A: "佣金报表生成时间为次月的5号之前，相关部门核实数据准确无误后即会生成。",
			},
			Question{
				Q: "4，什么时候派发佣金？",
				A: "风控部门审核且通过后，每月5号开始予以派发，10号之前派发完毕。",
			},
			Question{
				Q: "5，在哪里查收佣金，佣金派发到哪里？",
				A: "佣金会派发至您本人后台账号中（注：后台需绑定提款账号）, 在后台设置中提款即可。",
			},
			Question{
				Q: "6，佣金最少能提多少？",
				A: "佣金最低提取金额为 100 元. 佣金不满100元将累计至下月。",
			},

			Question{
				Q: "7，提款多长时间可以到账？",
				A: "提款会尽快安全到账。",
			},
		},
	},

	map[string]interface{}{
		"group": "注意事项",
		"itme": []Question{
			Question{
				Q: "1，违规行为怎么处理？",
				A: "为有效防止非诚信代理滥用我们所提供的优惠制度，审查部门将严格审核每位注册时提供的个人资料（包括姓名、邮件及电话等）经审核发现有任何不良营利企图或与其他代理商、会员进行合谋套利等不诚信行为，我们将永久关闭该合作，并停止发放佣金。",
			},
			Question{
				Q: "2.没有达到活跃标准产生的净盈利会累计到下个月吗？",
				A: "每月至少10 名有效活跃线下会员，若当月活跃会员少于 10 人，则当月无佣金，产生的正盈利不累计。",
			},
			Question{
				Q: "3.产生了负盈利会累计吗？？",
				A: "每月至少10 名有效活跃线下会员，若当月活跃会员少于 10 人，则当月无佣金，产生的正盈利不累计。",
			},
		},
	},

	map[string]interface{}{
		"group": "佣金计算方式",
		"itme": []Question{
			Question{
				Q: "1，我的佣金可以一直存在账户里吗？",
				A: "若您未能及时领取佣金，我们将安全保管佣金6个月，超过此时限将会归零。",
			},
			Question{
				Q: "2.为什么报表只显示1个月？",
				A: "为了您的个人及数据安全，只显示一个月，如有疑问请联系专员！。",
			},
			Question{
				Q: "3.佣金等级是什么意思？",
				A: "佣金分5个级别，根据相应的等级拿相应的佣金比例。3.1.净盈利为0-50000，当月活跃人数≥10，佣金比例为30%。3.2.净盈利为50001-200000，当月活跃人数≥15，佣金比例为35%。3.3.净盈利为200001-400000，当月活跃人数≥20，佣金比例为40%。3.4.净盈利为400001-800000，当月活跃人数≥25，佣金比例为45%。3.5.净盈利≥800001，当月活跃人数≥35，佣金比例为50%",
			},
			Question{
				Q: "4.会员上分制度",
				A: "1活跃会员达到35人以上方有资格申请。2请联系专员申请。3上分额度会在月底计算佣金时扣除。4最大上分月总额度不得超过10000元。5官方建议给出上分金额2倍流水起步。6专员拥有随时取消及停止此制度的权限。",
			},
		},
	},
}

type feedBackType struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	ImgUrl  string `json:"img_url"`
}

var FeedBackType = []feedBackType{
	feedBackType{1, "存款问题", "/frontend/feedback/icon/deposit.png"},
	feedBackType{2, "提款问题", "/frontend/feedback/icon/withdraw.png"},
	feedBackType{3, "游戏问题", "/frontend/feedback/icon/game.png"},
	feedBackType{4, "优惠问题", "/frontend/feedback/icon/discount.png"},
	feedBackType{5, "网站/App登录", "/frontend/feedback/icon/login.png"},
	feedBackType{6, "修改资料", "/frontend/feedback/icon/modify.png"},
	feedBackType{7, "流水问题", "/frontend/feedback/icon/flow.png"},
	feedBackType{8, "其他", "/frontend/feedback/icon/other.png"},
}

type VipIcon struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	ImgUrl    string `json:"img_url"`
	ActiveUrl string `json:"active_url"`
}

var VipIcons = []VipIcon{
	VipIcon{1, "豪礼赠送", "/frontend/vip/big_gift.png", "/frontend/vip/big_gift.png"},
	VipIcon{2, "场馆返水", "/frontend/vip/rebate.png", "/frontend/vip/rebate.png"},
	VipIcon{3, "每日提款", "/frontend/vip/withdraw.png", "/frontend/vip/withdraw.png"},
	VipIcon{4, "每月红包", "/frontend/vip/redpacket.png", "/frontend/vip/redpacket.png"},
	VipIcon{5, "生日礼金", "/frontend/vip/bithday_gift.png", "/frontend/vip/bithday_gift.png"},
	VipIcon{6, "晋级礼金", "/frontend/vip/promotion_gift.png", "/frontend/vip/promotion_gift.png"},
	VipIcon{7, "晋级优惠", "/frontend/vip/promotion_discount.png", "/frontend/vip/promotion_discount.png"},
}
