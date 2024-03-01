package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go_learn/app/http/controller/api/v1"
	"go_learn/app/requests"
	"go_learn/pkg/auth"
	"go_learn/pkg/jwt"
	"go_learn/pkg/response"
)

type LoginController struct {
	v1.BaseAPIController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}
	// 2.尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，返回错误提示
		response.Error(c, err, "账号不存在")
	} else {
		// 成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
