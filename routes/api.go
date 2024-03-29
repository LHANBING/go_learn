// Package routes 注册路由
package routes

import (
	controllers "go_learn/app/http/controller/api/v1"
	"go_learn/app/http/controller/api/v1/auth"
	"go_learn/app/http/middleware"
	"go_learn/pkg/config"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("v1")
	}

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middleware.LimitIP("10000-H"))
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middleware.LimitIP("2000-H"))
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

		uc := new(controllers.UsersController)
		// 获取当前用户
		v1.GET("/user", middleware.AuthTWT(), uc.CurrentUser)
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("", uc.Index)
			usersGroup.PUT("", middleware.AuthTWT(), uc.UpdateProfile)
			usersGroup.PUT("/email", middleware.AuthTWT(), uc.UpdateEmail)
			usersGroup.PUT("/phone", middleware.AuthTWT(), uc.UpdatePhone)
			usersGroup.PUT("/password", middleware.AuthTWT(), uc.UpdatePassword)
			usersGroup.PUT("/avatar", middleware.AuthTWT(), uc.UpdateAvatar)
		}

		cgc := new(controllers.CategoriesController)
		cgcGroup := v1.Group("/category")
		{
			cgcGroup.GET("", cgc.Index)
			cgcGroup.POST("", middleware.AuthTWT(), cgc.Store)
			cgcGroup.PUT("/:id", middleware.AuthTWT(), cgc.Update)
			cgcGroup.DELETE("/:id", middleware.AuthTWT(), cgc.Delete)
		}

		tpc := new(controllers.TopicsController)
		tpcGroup := v1.Group("/topic")
		{
			tpcGroup.GET("", tpc.Index)
			tpcGroup.POST("", middleware.AuthTWT(), tpc.Store)
			tpcGroup.PUT("/:id", middleware.AuthTWT(), tpc.Update)
			tpcGroup.DELETE("/:id", middleware.AuthTWT(), tpc.Delete)
			tpcGroup.GET("/:id", tpc.Show)
		}

		lsc := new(controllers.LinksController)
		linksGroup := v1.Group("/link")
		{
			linksGroup.GET("", lsc.Index)
		}
	}
}
