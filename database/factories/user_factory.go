package factories

import (
	"github.com/bxcodec/faker/v3"
	"go_learn/app/models/user"
	"go_learn/pkg/hash"
	"go_learn/pkg/helpers"
)

func MakeUser(times int) []user.User {
	var objs []user.User
	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: hash.BcryptHash("123qwe"),
		}
		objs = append(objs, model)
	}
	return objs
}
