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
	MaxOpenConns    int
	MaxIdleTime     time.Duration
	ConnMaxLifetime time.Duration

	RedisHost          string
	RedisPort          string
	RedisAuth          string
	RedisDb            int
	MaxIdle            int
	MaxActive          int
	IdleTimeout        time.Duration
	ConnFailRetryTimes int
	ReConnectInterval  time.Duration

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	// go test时需要用绝对路径F:\project\goWebDemo\config.ini
	file, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("配置文件读取出错,请检查文件路径是否正确", err)
	}
	LoadServer(file)
	LoadDb(file)
	LoadRedis(file)
	LoadQiNiu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("29js10js")

}

func LoadDb(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("xz")
	DbName = file.Section("database").Key("DbName").MustString("goWebDemo")
	MaxOpenConns = file.Section("database").Key("MaxOpenConns").MustInt(100)
	MaxIdleTime = file.Section("database").Key("MaxIdleTime").MustDuration(10)
	ConnMaxLifetime = file.Section("database").Key("ConnMaxLifetime").MustDuration(10 * time.Second)

}

func LoadRedis(file *ini.File) {
	RedisHost = file.Section("redis").Key("RedisHost").MustString("127.0.0.1")
	RedisPort = file.Section("redis").Key("RedisPort").MustString("6379")
	RedisAuth = file.Section("redis").Key("RedisAuth").MustString("")
	RedisDb = file.Section("redis").Key("RedisDb").MustInt(1)
	MaxIdle = file.Section("redis").Key("MaxIdle").MustInt(10)
	MaxActive = file.Section("redis").Key("MaxActive").MustInt(1000)
	IdleTimeout = file.Section("redis").Key("IdleTimeout").MustDuration(60 * time.Second)
	ConnFailRetryTimes = file.Section("redis").Key("ConnFailRetryTimes").MustInt(3)
	ReConnectInterval = file.Section("redis").Key("ReConnectInterval").MustDuration(1)

}

func LoadQiNiu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()

}
