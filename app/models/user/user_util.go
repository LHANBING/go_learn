package user

import (
	"go_learn/pkg/database"
)

// IsEmailExist 判断email是否已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断phone是否已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(LoginID string) (userModel User) {
	database.DB.
		Where("phone = ?", LoginID).
		Or("email = ?", LoginID).
		Or("name = ?", LoginID).
		First(&userModel)
	return
}

func Get(UserID string) (userModel User) {
	database.DB.Where("id", UserID).First(&userModel)
	return
}
