// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/http/controller/api/v1/auth"
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
			})
		})

		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断邮箱是否注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("signup/using-phone", suc.SignupUsingPhone)
			authGroup.POST("signup/using-email", suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)

			lgc := new(auth.LoginController)
			// 手机号+验证码登录
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
			authGroup.POST("login/refresh-token", lgc.RefreshToken)
			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", pwc.ResetByEmail)
		}
	}
}
