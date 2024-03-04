// Package auth 用户注册，登录，密码重置
package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go_learn/app/http/controller/api/v1"
	"go_learn/app/models/user"
	"go_learn/app/requests"
	"go_learn/pkg/response"
)

type PasswordController struct {
	v1.BaseAPIController
}

// ResetByPhone 使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	// 1. 表单验证
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	// 2.更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)

	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}

func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	// 1. 表单验证
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}
	// 2. 更新密码
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
