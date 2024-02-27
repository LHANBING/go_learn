package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// 初始化  Default 初始化 = New() + 注册中间件
	//r := gin.Default()

	// 注册一个Gin Engine实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {
		// 以JSON格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})

	// 404 请求
	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是HTML的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义",
			})
		}
	})

	// 运行服务，指定8080端口
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
