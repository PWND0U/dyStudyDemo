package main

import (
	"dyStudyDemo/app/config"
	_ "dyStudyDemo/app/config"
	"dyStudyDemo/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// new 一个 Gin Engine 实例
	e := gin.New()
	// 注册中间件
	bootstrap.SetupRouter(e)
	fmt.Println(config.C().App.AppPort)
	err := e.Run("0.0.0.0:" + config.C().App.AppPort)
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
