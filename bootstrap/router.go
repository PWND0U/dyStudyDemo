package bootstrap

import (
	"dyStudyDemo/routers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//这里是router初始化的地方

func SetupRouter(e *gin.Engine) {
	registerGlobalMiddleWare(e) //注册全局中间件
	//注册路由
	routers.RegisterRustfulApi(e)
	setup404Handler(e) //配置404默认页面
}
func registerGlobalMiddleWare(e *gin.Engine) {
	e.Use(gin.Logger(), gin.Recovery())
}
func setup404Handler(e *gin.Engine) {
	// 处理 404 请求
	e.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
