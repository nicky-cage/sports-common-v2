package consts

import "time"

// DurationOffsets 各个游戏最大拉单区间限制
var DurationOffsets = map[string]time.Duration{
	"IM":   time.Hour * 1,    // 1小时 - 体育 - ok
	"BTI":  time.Hour * 1,    // 1小时 - 体育 - ok
	"SABA": time.Hour * 1,    // 1小时 - 体育 - ok
	"KY":   time.Minute * 30, // 30分 - 棋牌 - ok
	"LEG":  time.Minute * 30, // 30分 - 棋牌 - ok
	"DT":   time.Hour * 1,    // 1小时 - 棋牌
	"VG":   time.Hour * 1,    // 1小时 - 棋牌 - ok
	"VR":   time.Hour * 1,    // 1小时 - 彩票 - ok
	"EBET": time.Hour * 1,    // 2小时 - 真人
	"BBIN": time.Hour * 1,    // 1小时 - 电子 + 真人
	"MG":   time.Hour * 1,    // 1小时 - 电子
	"JDB":  time.Minute * 5,  // 1小时 - 电子 - ok
	"PT":   time.Minute * 30, // 1小时 - 电子
	//"AG":     time.Hour * 1,    // 1天  - 电子 - ok
	"AG":     time.Minute * 10, // 10分钟 - AG在改为HTTP拉取后最大时间间隔为10分钟
	"LEIHUO": time.Hour * 1,    // 1小时 - 电竞
	"IME":    time.Hour * 2,    // 1小时 - 电竞 - ok
	"HL":     time.Minute * 60, // 欢乐 - 棋牌
	"ALB":    time.Minute * 15, // 15分钟 - 体育
	"BG":     time.Hour * 1,    // 1小时 - 真人
	"CQ9":    time.Hour * 1,    // 1小时 - 电子
	"SG":     time.Hour * 1,    // 1小时 - 彩票
	"AVIA":   time.Hour * 1,    // 1小时 - 电竞
	"XJ188":  time.Hour * 1,    // 1小时 - 小金体育
	"BAOLI":  time.Minute * 20, // 20分钟 - 保利体育(最大间隔)
	"WE":     time.Minute * 30, // 30分钟 - WE真人
	"UG":     time.Minute * 30, // 30分钟 - UG体育
	"OG":     time.Minute * 10, // 10分钟 - OG真人(最大时间间隔为10分钟)
	"WM":     time.Hour * 1,    // 1小时 - WM真人
	"OB":     time.Hour * 1,    // 1小时 - OB体育
}

// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%KY%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%LEG%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%VG%' LIMIT 1 \G;

// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%DT%' LIMIT 1 \G;

// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%SABA%' LIMIT 1 \G;

// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%BTI%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%IM%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%VR%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%EBET%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%BBIN%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%MG%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%JDB%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%AG%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%LEIHUO%' LIMIT 1 \G;
// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%IME%' LIMIT 1 \G;

// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%PT%' LIMIT 1 \G;

// SELECT * FROM wagers_statistics_status WHERE wager_id LIKE '%HL%' LIMIT 1 \G;
