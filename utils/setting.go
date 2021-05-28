package utils

import (
	"fmt"
	"github.com/go-ini/ini"
	"time"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost          string
	DbPort          string
	DbUser          string
	DbPassWord      string
	DbName          string
	MaxIdleTime     time.Duration
	MaxOpenConns    int
	ConnMaxLifetime time.Duration

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取出错,请检查文件路径是否正确", err)
	}
	LoadServer(file)
	LoadDb(file)
	LoadQiNiu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("ldjskjdwdq213sjnk")

}

func LoadDb(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("xz")
	DbName = file.Section("database").Key("DbName").MustString("goWebDemo")
	MaxIdleTime = file.Section("database").Key("MaxIdleTime").MustDuration(10)
	MaxOpenConns = file.Section("database").Key("MaxOpenConns").MustInt(100)
	ConnMaxLifetime = file.Section("database").Key("ConnMaxLifetime").MustDuration(10 * time.Second)

}

func LoadQiNiu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()

}
