package request

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetViewFile 生成视图文件
func GetViewFile(c *gin.Context, formatString string, args ...string) string {
	spec := "_"
	if len(args) > 0 {
		spec = args[0]
	}
	if IsAjax(c) {
		return fmt.Sprintf(formatString, spec)
	}
	return fmt.Sprintf(formatString, "")
}
