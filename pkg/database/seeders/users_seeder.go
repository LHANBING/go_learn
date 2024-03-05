package seeders

import (
	"fmt"
	"go_learn/database/factories"
	"go_learn/pkg/console"
	"go_learn/pkg/logger"
	"go_learn/pkg/seed"
	"gorm.io/gorm"
)

func init() {
	// 添加Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		//创建10个用户对象
		users := factories.MakeUser(10)

		// 批量创建用户(注意批量创建不会调用模型钩)
		result := db.Table("users").Create(&users)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
