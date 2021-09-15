package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/viper"
	"goWebDemo/utils"
	"log"
)

func main() {
	utils.LoadConfig()
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("Database.User"),
		viper.GetString("Database.Password"),
		viper.GetString("Database.Host"),
		viper.GetString("Database.Port"),
		viper.GetString("Database.Name"))
	db, err := sql.Open("mysql", dns)

	if err != nil {
		log.Fatal("连接数据库失败: ", err)
	}

	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	//err = m.Down()
	if err != nil {
		log.Fatal(err)
	}
	_ = m.Steps(2)

	fmt.Println("迁移数据成功")
}
