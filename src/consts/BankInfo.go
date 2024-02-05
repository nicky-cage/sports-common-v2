package consts

// CdnImgUrl CDN图片URL
const CdnImgUrl = ""

// BankInfo 银行信息
type BankInfo struct {
	BankID   uint32 `json:"bank_id"`
	BankCode string `json:"bank_code"`
	BankName string `json:"bank_name"`
	Icon     string `json:"icon"`
	IconP    string `json:"-"`
	Color    string `json:"-"`
}

// BankList 银行列表
// TODO: 将来这些信息要同步到数据库当中, 便于统一管理
var BankList = map[string]BankInfo{
	"BOC": BankInfo{
		BankCode: "BOC",
		BankName: "中国银行",
		Icon:     "/upload/images/bank/BOC.png",
		IconP:    "/upload/images/bank/BOC_P.png",
		Color:    "#ffd8d8",
	},
	"ICBC": BankInfo{
		BankCode: "ICBC",
		BankName: "工商银行",
		Icon:     "/upload/images/bank/ICBC.png",
		IconP:    "/upload/images/bank/ICBC_P.png",
		Color:    "#ffd8d8",
	},
	"ABC": BankInfo{
		BankCode: "ABC",
		BankName: "农业银行",
		Icon:     "/upload/images/bank/ABC.png",
		IconP:    "/upload/images/bank/ABC_P.png",
		Color:    "#ccecdb",
	},
	"CCB": BankInfo{
		BankCode: "CCB",
		BankName: "建设银行",
		Icon:     "/upload/images/bank/CCB.png",
		IconP:    "/upload/images/bank/CCB_P.png",
		Color:    "#c3d9ff",
	},
	"CMB": BankInfo{
		BankCode: "CMB",
		BankName: "招商银行",
		Icon:     "/upload/images/bank/CMB.png",
		IconP:    "/upload/images/bank/CMB_P.png",
		Color:    "#ffd8d8",
	},
	"CMBC": BankInfo{
		BankCode: "CMBC",
		BankName: "民生银行",
		Icon:     "/upload/images/bank/CMBC.png",
		IconP:    "/upload/images/bank/CMBC_P.png",
		Color:    "#ccecdb",
	},
	"BCM": BankInfo{
		BankCode: "BCM",
		BankName: "交通银行",
		Icon:     "/upload/images/bank/BCM.png",
		IconP:    "/upload/images/bank/BCM_P.png",
		Color:    "#c3d9ff",
	},
	"BOG": BankInfo{
		BankCode: "BOG",
		BankName: "广州银行",
		Icon:     "/upload/images/bank/BOG.png",
		IconP:    "/upload/images/bank/BOG_P.png",
		Color:    "#ffd8d8",
	},
	"HXBANK": BankInfo{
		BankCode: "HXBANK",
		BankName: "华夏银行",
		Icon:     "/upload/images/bank/HXBANK.png",
		IconP:    "/upload/images/bank/HXBANK_P.png",
		Color:    "#ffd8d8",
	},
	"SPABANK": BankInfo{
		BankCode: "SPABANK",
		BankName: "平安银行",
		Icon:     "/upload/images/bank/SPABANK.png",
		IconP:    "/upload/images/bank/SPABANK_P.png",
		Color:    "#f2f6c7",
	},
	"QDCCB": BankInfo{
		BankCode: "QDCCB",
		BankName: "青岛银行",
		Icon:     "/upload/images/bank/QDCCB.png",
		IconP:    "/upload/images/bank/QDCCB_P.png",
		Color:    "#ffd8d8",
	},
	"PSBC": BankInfo{
		BankCode: "PSBC",
		BankName: "邮政银行",
		Icon:     "/upload/images/bank/PSBC.png",
		IconP:    "/upload/images/bank/PSBC_P.png",
		Color:    "#ccecdb",
	},
	"CNCB": BankInfo{
		BankCode: "CNCB",
		BankName: "中信银行",
		Icon:     "/upload/images/bank/CNCB.png",
		IconP:    "/upload/images/bank/CNCB_P.png",
		Color:    "#ffd8d8",
	},
	"CEB": BankInfo{
		BankCode: "CEB",
		BankName: "光大银行",
		Icon:     "/upload/images/bank/CEB.png",
		IconP:    "/upload/images/bank/CEB_P.png",
		Color:    "#f2f6c7",
	},
	"BOCD": BankInfo{
		BankCode: "BOCD",
		BankName: "成都银行",
		Icon:     "/upload/images/bank/BOCD.png",
		IconP:    "/upload/images/bank/BOCD_P.png",
		Color:    "#f2f6c7",
	},
	"CIB": BankInfo{
		BankCode: "CIB",
		BankName: "兴业银行",
		Icon:     "/upload/images/bank/CIB.png",
		IconP:    "/upload/images/bank/CIB_P.png",
		Color:    "#c3d9ff",
	},
	"SPDB": BankInfo{
		BankCode: "SPDB",
		BankName: "浦发银行",
		Icon:     "/upload/images/bank/SPDB.png",
		IconP:    "/upload/images/bank/SPDB_P.png",
		Color:    "#c3d9ff",
	},
	"CGB": BankInfo{
		BankCode: "CGB",
		BankName: "广发银行",
		Icon:     "/upload/images/bank/CGB.png",
		IconP:    "/upload/images/bank/CGB_P.png",
		Color:    "#ffd8d8",
	},
}
