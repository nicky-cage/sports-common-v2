package flags

import (
	"flag"
	"fmt"
	"os"
	"sports-common/config"
	"sports-common/consts"
	"time"
)

// 初始化相关操作
func InitConfigs() {
	flag.BoolVar(&help, "h", false, "help info")
	flag.BoolVar(&isShowVersion, "v", false, "version 显示代码版本号")
	flag.BoolVar(&isShowDesc, "vd", false, "该版本的描述信息")
	flag.BoolVar(&disableSysMaintain, "m", false, "false系统维护的标记起作用 true表示后台维护的标记有不起作用")
	flag.StringVar(&configFilePath, "config", "", `必填，setting.ini配置文件的绝对路径,如"D:\shares\go-v\src\integrated-game-api\bin\setting.ini"`)
	flag.StringVar(&extConfigFilePath, "ext", "", `必填，setting_ext.ini配置文件的绝对路径,如"D:\shares\go-v\src\integrated-game-api\bin\setting_ext.ini"`)
	flag.Usage = func() { // 重写flag的usage,如果重写，默认的是使用flag.Boovar等中设置帮助信息说明
		_, _ = fmt.Fprintf(os.Stderr, "sports-app version: 2.0.1 beta\n"+
			"Usage: -config=configvalue -ext=extConfigFile")
		flag.PrintDefaults()
	}

	flag.Parse()
	parseParameters()
}

// 加载额外的参数
func InitParameterByString(paramValue *string, paramName string, remark string) {
	flag.StringVar(paramValue, paramName, "", remark)
}

// 加载判断启动参数
func parseParameters() {
	if help {
		flag.Usage()
		os.Exit(0)
	} else if isShowVersion {
		fmt.Println(versionName + versionNameStr)
		os.Exit(0)
	} else if isShowDesc {
		fmt.Println(versionName + versionNameStr + "\n" + versionDesc)
		os.Exit(0)
	}

	LoadSelfExtConfigFile(extConfigFilePath) // 先加载 ext 文件

	consts.CustomDebug = IniGet("platform.custom_debug") == "1" //自定义debug是否开启
	consts.DisableSysMaintain = disableSysMaintain
	consts.AppName = IniGet("app_name")

	versionName = consts.AppName // 修改
	buildVersion := time.Now().Format("20060102_1504")
	versionNameStr += "_build_" + buildVersion
	fmt.Println(versionName + "-" + versionNameStr + "\n" + versionDesc)

	// 初始化自身微服务的全局变量
	config.LoadConfigs(configFilePath) // 再加载 config 文件
}
