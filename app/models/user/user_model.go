// Package user 存放用户model相关逻辑
package user

import (
	"go_learn/app/models"
	"go_learn/pkg/database"
	"go_learn/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"_"`
	Phone    string `json:"_"`
	Password string `json:"_"`

	models.CommonTimestampsField
}

// Create 创建用户
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
