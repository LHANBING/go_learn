package bootstrap

import (
	"fmt"
	"go_learn/pkg/config"
	"go_learn/pkg/redis"
)

// SetupRedis 初始化Redis
func SetupRedis() {
	// 建立连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"))
}
