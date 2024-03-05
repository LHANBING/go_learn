package user

import (
	"github.com/gin-gonic/gin"
	"go_learn/pkg/app"
	"go_learn/pkg/database"
	"go_learn/pkg/paginator"
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

// GetByEmail 通过手机号来获取用户
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
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

func ALl() (users []User) {
	database.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage)
	return
}
