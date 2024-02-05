package colors

import (
	"github.com/gookit/color"
)

// Info 日志
func Info(args ...interface{}) {
	color.Green.Println(args...)
}

// Warn 警告
func Warn(args ...interface{}) {
	color.Warnln(args...)
}

// Error 错误
func Error(args ...interface{}) {
	color.Errorln(args...)
}
