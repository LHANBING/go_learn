// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"go_learn/pkg/config"
	"net/http"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// v1 路由组
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// JSON响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World",
				"port":  config.Get("app.port"),
			})
		})
	}
}
