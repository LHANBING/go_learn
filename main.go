package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_learn/bootstrap"
)

func main() {

	// 注册一个Gin Engine实例
	r := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	// 运行服务，指定8080端口
	err := r.Run(":8000")
	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}
}
