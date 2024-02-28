// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/http/middleware"
	"go_learn/routes"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)
	// 注册API路由
	routes.RegisterAPIRoutes(router)
	// 配置404路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	// 初始化  Default 初始化 = New() + 注册中间件
	//r := gin.Default()
	router.Use(
		middleware.Logger(),
		//gin.Logger(), //gin框架自带日志
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	// 处理404
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是html
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回JSON
			c.JSON(http.StatusOK, gin.H{
				"error_code":    404,
				"error_message": "路由未定义， 请确认URL和请求方法是否正确",
			})
		}
	})
}
