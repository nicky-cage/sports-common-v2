package consts

// PayChannelIds 支付渠道编号
var PayChannelIds = []string{
	"1", "2", "3", "4", "5", "6", "7", "8",
}

// PayChannels 支付渠道
var PayChannels = map[string]string{
	"0": "银行转账(离线)",
	"1": "网银转账",
	"2": "支付宝",
	"3": "微信",
	"4": "QQ钱包",
	"5": "快捷支付",
	"6": "京东支付",
	"7": "银联扫码",
	"8": "虚拟币",
	"9": "云闪付",
}

// PayChannelTypes 支付渠道类型
var PayChannelTypes = map[string]int{
	"transfer":     0,
	"ebank":        1,
	"alipay":       2,
	"weixin":       3,
	"qqpay":        4,
	"quickpay":     5,
	"jdpay":        6,
	"unionpay":     7,
	"virtual_coin": 8,
	"cloud":        9,
}

var OfflineDepositCheck = map[string]string{
	"1": "银行未收到该笔存款",
	"2": "重复提交",
	"3": "提交金额与实际到账金额不符",
	"4": "存款姓名与提交姓名不相符",
}

// 可用渠道列表
var PayChannelList = []map[string]interface{}{
	{
		"channel_name": "网银",
		"channel_code": "ebank",
		"vip_list":     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	},
	{
		"channel_name": "支付宝",
		"channel_code": "alipay",
		"vip_list":     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	},
	{
		"channel_name": "微信",
		"channel_code": "weixin",
		"vip_list":     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	},
	{
		"channel_name": "QQ钱包",
		"channel_code": "qqpay",
		"vip_list":     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	},
	{
		"channel_name": "快捷",
		"channel_code": "quickpay",
		"vip_list":     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	},
	{
		"channel_name": "京东",
		"channel_code": "jdpay",
		"vip_list":     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	},
	{
		"channel_name": "银联扫码",
		"channel_code": "unionpay",
		"vip_list":     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	},
}
