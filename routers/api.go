package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRustfulApi(e *gin.Engine) {
	douyinGroup := e.Group("/douyin") //抖音路由组
	{
		douyinGroup.GET("/feed", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "douyin",
			})
		})
	}
	douyinUserGroup := douyinGroup.Group("/user") //抖音用户路由组
	{
		//用户注册
		douyinUserGroup.POST("/register", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "douyin1",
			})
		})
		//用户登录
		douyinUserGroup.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "douyin2",
			})
		})
		//获取用户信息
		douyinUserGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "douyin3",
			})
		})
	}
	douyinPublishGroup := douyinGroup.Group("/publish") //抖音投稿相关接口路由组
	{
		douyinPublishGroup.POST("/action", func(c *gin.Context) {

		})
		douyinPublishGroup.GET("/list", func(c *gin.Context) {

		})
	}
	douyinFavoriteGroup := douyinGroup.Group("/favorite") //抖音关注操作路由组
	{
		douyinFavoriteGroup.GET("/list", func(c *gin.Context) {

		})
		douyinFavoriteGroup.POST("/action", func(c *gin.Context) {

		})
	}

}
