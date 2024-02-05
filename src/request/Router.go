package request

import "github.com/gin-gonic/gin"

// RouterMethod 路由方法
type RouterMethod = map[string]func(*gin.Context)

// AddGETRouters 增加get方法
func AddGETRouters(g *gin.RouterGroup, methods map[string]func(*gin.Context)) {
	for url, f := range methods {
		g.GET(url, f)
	}
}

// AddPOSTRouters 增加post方法
func AddPOSTRouters(g *gin.RouterGroup, methods map[string]func(*gin.Context)) {
	for url, f := range methods {
		g.POST(url, f)
	}
}
