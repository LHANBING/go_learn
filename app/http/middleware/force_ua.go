package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_learn/pkg/response"
)

// ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取User-Agent标头信息
		if len(c.Request.Header["user-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent 标头未找到"), "请求必须附带 User-Agent 标头")
			return
		}
		c.Next()
	}
}
