package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"sports-common/consts"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// LoggerInfo 日志信息
var LoggerInfo = struct {
	AppName   string
	LogPath   string
	LoggerRus *logrus.Logger
	MongDb    *mongo.Client
}{}

// Logger 日志
var Logger *logrus.Logger

// LogMaxAge 文件最大保存时间
var LogMaxAge = 30 * 24 * time.Hour

// LogRotationTime 日志切割时间间隔
var LogRotationTime = 24 * time.Hour

// Start 初始化
// ```go
// import "sports-common/log"
//	func main() {
//	logger := log.New("api", "./").Init()
//	logger...
// }
// ```
func Start(args ...*mongo.Client) {
	if Logger != nil { // 表示已经初始化过了
		fmt.Printf("log start %v\n", Logger)
		return
	}
	LogPath = consts.LogPath
	//	fmt.Printf("log path %s\n", LogPath)
	l, err := initLogrus(consts.AppName, LogPath)
	if err != nil {
		log.Println("init libs.logrus failed " + err.Error())
		os.Exit(0)
		return
	}

	Logger = l
	LoggerInfo.AppName = consts.AppName
	LoggerInfo.LogPath = LogPath
	LoggerInfo.LoggerRus = l
	if len(args) >= 1 { // 设置mongodb
		LoggerInfo.MongDb = args[0]
	}
}

// Debug 输入调试信息
func Debug(info string, args ...interface{}) {
	fmt.Printf(info, args...)
	fmt.Println("")
}

// Err 输入错误信息
func Err(info string, args ...interface{}) {
	fmt.Printf(info, args...)
	if funcName, file, line, ok := runtime.Caller(1); ok {
		fmt.Printf("|-- 函数: %s\n", runtime.FuncForPC(funcName).Name())
		fmt.Printf("|-- 文件: %s\n|-- 行数: %d\n", file, line)
	}
	// debug.PrintStack()
}

// IsDir 判断所给路径是否为文件夹
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		// panic(err)
		return false
	}
	return s.IsDir()
}

// NewLogger 生成日志
func NewLogger(appName, logPath string) (*logrus.Logger, error) {
	return initLogrus(appName, logPath)
}

// 初始化Logrus
func initLogrus(appName, logPath string) (*logrus.Logger, error) {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	filePrefix := logPath + "/" + appName

	logClient := logrus.New()
	logClient.Out = src
	// logClient.Out = os.Stdout //stdout will output in console
	logClient.SetLevel(logrus.DebugLevel)

	getLogWriter := func(level string) *rotatelogs.RotateLogs {
		fileDir := filePrefix + "/" + time.Now().Format("200601")
		if !isDir(fileDir) {
			err := os.MkdirAll(fileDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
		linkLogFile := fileDir + "/" + appName + "." + level + ".log"
		logWriter, err := rotatelogs.New(
			// filePrefix+"."+level+".%Y%m%d%H.log",
			// fileDir+"/"+appName+"."+level+".%Y%m%d.log",
			fileDir+"/%m%d."+level+"."+appName+".log",    // 因为目录带了年月，所以不要年了，调整下，让日期和level放在前面方便查看
			rotatelogs.WithLinkName(linkLogFile),         // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(LogMaxAge),             //  文件最大保存时间
			rotatelogs.WithRotationTime(LogRotationTime), // 日志切割时间间隔
		)
		if err != nil {
			os.Exit(1)
			return nil
		}
		return logWriter
	}
	infoWriter := getLogWriter("info")
	fatalWriter := getLogWriter("fatal")
	debugWriter := getLogWriter("debug")
	warnWriter := getLogWriter("warn")
	errorWriter := getLogWriter("error")

	writeMap := lfshook.WriterMap{ // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  infoWriter,
		logrus.FatalLevel: fatalWriter,
		logrus.DebugLevel: debugWriter,
		logrus.WarnLevel:  warnWriter,
		logrus.ErrorLevel: errorWriter,
	}

	//formatter := &logrus.JSONFormatter{ //设置日志格式
	//	TimestampFormat: consts.TimeLayoutYmdHis,
	//	//PrettyPrint: true,
	//}

	//formatter := &logrus.TextFormatter{
	//	DisableColors:   true,
	//	TimestampFormat: consts.TimeLayoutYmdHis,
	//	//LogFormat:       "[%lvl%]",
	//}
	//logClient.SetFormatter(&LogTxtFormatter{})
	formatter := &TxtFormatter{}

	logClient.AddHook(NewFileHook())                        // 日志文件 行号
	logClient.AddHook(lfshook.NewHook(writeMap, formatter)) // 日志格式

	return logClient, nil
}

// 定义Hook: 文件名称/行数
type myHook struct {
	FileName string // 输出日志的代码文件名称
	Line     string // 打印日志的行
	Skip     int
	levels   []logrus.Level
}

// Fire 实现 logrus.Hook 接口, 当写入日志时, 可以同时写入数据库/消息队列
func (hook *myHook) Fire(entry *logrus.Entry) error {
	fileName, line := findCaller(hook.Skip)
	entry.Data[hook.FileName] = fileName
	entry.Data[hook.Line] = line
	// 同时写入到数据库当中
	return nil
}

// Levels 实现 logrus.Hook 接口
func (hook *myHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// NewFileHook 自定义hook
func NewFileHook(levels ...logrus.Level) logrus.Hook {
	hook := myHook{
		FileName: "filePath",
		Line:     "line",
		Skip:     5,
		levels:   levels,
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}
	return &hook
}

//
func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	n := 0
	// 获取执行代码的文件名
	for i := len(file) - 1; i > 0; i-- {
		if string(file[i]) == "/" {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}

//
func findCaller(skip int) (string, int) {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		// 文件名不能以logrus开头
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return file, line
}
