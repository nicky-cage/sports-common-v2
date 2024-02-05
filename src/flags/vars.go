package flags

// 默认定义的一些参数
var (
	versionName        = "sports-admin" //版本名称
	versionNameStr     = "2.0.0"        //版本号
	help               bool             //是否显示帮助信息
	isShowVersion      bool             //是否显示版本号
	isShowDesc         bool             //显示该版本的描述信息
	disableSysMaintain bool             //默认不禁用系统维护的标记
	configFilePath     string           //主配置
	extConfigFilePath  string           //微服务自身额外配置
	versionDesc        = "版本描述"         //版本描述
)
