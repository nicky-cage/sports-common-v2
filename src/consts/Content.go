package consts

// 关联场景
var VenueTypes = map[uint8]string{
	1: "存款",
	2: "取款",
	3: "转账",
	4: "奖励",
	5: "VIP等级",
	6: "合营计划",
}

// 平台 类型
var PlatformTypes = map[uint8]string{
	1: "PC",
	2: "H5",
	3: "安卓全站",
	4: "安卓体育",
	5: "IOS全站",
	6: "IOS体育",
}

// 底中部信息类型
var BottomTypes = map[uint8]string{
	1: "合作方图标",
	2: "牌照文本",
	3: "牌照图标",
	4: "帮助跳转",
	5: "赞助图标",
	6: "版权文本",
}

// 底部信息配置 - 内容类型
var BottomContentTypes = map[uint8]string{
	0: "图片",
	1: "图片 + 链接",
	2: "文本",
	3: "文本 + 链接",
}
