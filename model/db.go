package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var Db *gorm.DB
var err error

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func InitDb() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("Database.User"),
		viper.GetString("Database.PassWord"),
		viper.GetString("Database.Host"),
		viper.GetString("Database.Port"),
		viper.GetString("Database.Name"))
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
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
	//_ = Db.AutoMigrate(&User{})

	sqlDB, _ := Db.DB()
	// 设置连接池中的最大闲置连接数
	sqlDB.SetMaxIdleConns(viper.GetInt("Database.MaxIdleConns"))

	// 设置数据库的最大连接数量
	sqlDB.SetMaxOpenConns(viper.GetInt("Database.MaxOpenConns"))

	// 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(viper.GetDuration("Database.ConnMaxLifetime"))
}
