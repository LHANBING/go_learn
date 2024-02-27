package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_learn/bootstrap"
	btsConfig "go_learn/config"
	"go_learn/pkg/config"
)

func init() {
	// 加载 config 目录下配置文件
	btsConfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 注册一个Gin Engine实例
	router := gin.New()

	// 初始化DB
	bootstrap.SetupDB()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务，指定端口
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}
}
