package model

import (
	"fmt"
	"goWebDemo/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

var db *gorm.DB
var err error

func InitDb() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// 日志模式
		Logger: logger.Default.LogMode(logger.Silent),
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务
		SkipDefaultTransaction: true,
		// 表名小写
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数是否正确: ", err)
		os.Exit(1)
	}

	// 迁移数据表
	//_ db.AutoMigrate()

	sqlDB, _ := db.DB()
	// 设置连接池中的最大闲置连接数
	sqlDB.SetConnMaxIdleTime(utils.MaxIdleTime)

	// 设置数据库的最大连接数量
	sqlDB.SetMaxOpenConns(utils.MaxOpenConns)

	// 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(utils.ConnMaxLifetime)
}
