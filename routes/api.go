// Package routes 注册路由
package routes

import (
	"fmt"
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
				"info": fmt.Sprintf("%v:%v@tcp(%v:%v)%v?charset=%v&parseTime=True&mutiStatements=true=Local",
					config.Get("database.mysql.username"),
					config.Get("database.mysql.password"),
					config.Get("database.mysql.host"),
					config.Get("database.mysql.port"),
					config.Get("database.mysql.database"),
					config.Get("database.mysql.charset")),
			})
		})
	}
}
