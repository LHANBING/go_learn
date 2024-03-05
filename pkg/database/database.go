// Package database 数据库操作
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"go_learn/pkg/config"
	"go_learn/pkg/console"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	// 使用gorm.Open链接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取低层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMySQLTables()
	case "sqlite":
		err = delteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}
	return err
}

func delteAllSqliteTables() error {
	tables := []string{}
	// 读取所有数据表
	err := DB.Select(&tables, "SELECT name from sqlite_master where type='table'").Error
	if err != nil {
		return err
	}

	// 删除所有表
	for _, table := range tables {
		err = DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	tables := []string{}

	// 读取所有数据表
	err := DB.Table("information_schema.tables").Where("table_schema = ?", dbname).Pluck("table_name", &tables).Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	DB.Exec("SET foreign_key_checks = 0;")

	//删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
		console.Success("Drop table=" + table)
	}

	// 开启 MYSQL 外键检测
	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}
