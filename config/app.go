// Package config 站点配置信息
package config

import "go_learn/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": config.Env("APP_NAME", "learn"),
			// 当前环境
			"env": config.Env("APP_ENV", "production"),
			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),
			// 应用服务端口
			"port": config.Env("APP_PORT", "8000"),
			// JWT
			"key": config.Env("APP_KEY", "12345qwert"),
			// 生成链接
			"url": config.Env("APP_URL", "http://localhost:8000"),
			// 设置时区
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
