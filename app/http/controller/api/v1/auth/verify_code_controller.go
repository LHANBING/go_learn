package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go_learn/app/http/controller/api/v1"
	"go_learn/pkg/captcha"
	"go_learn/pkg/logger"
	"go_learn/pkg/response"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, answer, err := captcha.NewCaptcha().GenerateCaptcha()

	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)
	// 返回用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
		"answer":        answer,
	})
}
