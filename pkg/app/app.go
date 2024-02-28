// Package app 应用信息
package app

import "go_learn/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "Local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}
