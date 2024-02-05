package log

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// TxtFormatter 文本日志
type TxtFormatter struct{}

// Format 实现logrus.Formatter
func (s *TxtFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	//level := strings.ToUpper(entry.Level.String())
	fileName := entry.Data["filePath"]
	line := entry.Data["line"]
	//msg := fmt.Sprintf("[%s][%s] %s - %s (Line:%d)\n", level, timestamp, entry.Message, fileName, line)
	msg := fmt.Sprintf("[%s] %s - %s (Line:%d)\n", timestamp, entry.Message, fileName, line)
	return []byte(msg), nil
}
