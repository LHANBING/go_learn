package main

import (
	"fmt"
	"go_learn/app/cmd"
	"go_learn/bootstrap"
	btsConfig "go_learn/config"
	"go_learn/pkg/config"
	"go_learn/pkg/console"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// 加载 config 目录下配置文件
	btsConfig.Initialize()
}

func main() {
	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   "GoLearn",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v:%s", os.Args, err.Error()))
	}

	// ------- old
	// 配置初始化，依赖命令行 --env参数
	/*var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化logger
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化DB
	bootstrap.SetupDB()

	// 初始化Redis
	bootstrap.SetupRedis()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 测试路由
	//router.GET("/v1/test", middleware.GuestJWT(), func(c *gin.Context) {
	//	c.String(http.StatusOK, "Hello guest !")
	//})

	// 运行服务，指定端口
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}*/
}
