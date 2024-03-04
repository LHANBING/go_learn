package cmd

import (
	"github.com/spf13/cobra"
	"go_learn/pkg/console"
	"go_learn/pkg/redis"
	"time"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	// 存进去 redis 中
	redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	// 从 Redis 里获取
	console.Success(redis.Redis.Get("hello"))
}
