// Package user 存放用户model相关逻辑
package user

import "go_learn/app/models"

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"_"`
	Phone    string `json:"_"`
	Password string `json:"_"`

	models.CommonTimestampsField
}
