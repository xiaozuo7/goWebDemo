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
	MaxOpenConns    int           // 设置数据库的最大连接数量
	MaxIdleConns    int           // 设置连接池中的最大闲置连接数
	ConnMaxLifetime time.Duration // 设置连接的最大可复用时间

	RedisHost          string
	RedisPort          string
	RedisAuth          string
	RedisDb            int
	MaxIdle            int           // 最大空闲数
	MaxActive          int           // 最大活跃数
	IdleTimeout        time.Duration // 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
	ConnFailRetryTimes int           // 从连接池获取连接失败，最大重试次数
	ReConnectInterval  time.Duration // 从连接池获取连接失败，每次重试之间间隔的秒数

	AccessKey string
	SecretKey string
	Bucket    string
	EndPoint  string

	Broker           string
	RedisBackendHost string
	RedisBackendPort string
	RedisBackendAuth string
	RedisBackendDB   int

	EtcdHosts string // etcd主机地址
	RpcPort   string // rpc端口
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
	LoadMachinery(file)
	LoadEtcd(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("")

}

func LoadDb(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("goWebDemo")
	MaxOpenConns = file.Section("database").Key("MaxOpenConns").MustInt(128)
	MaxIdleConns = file.Section("database").Key("MaxIdleConns").MustInt(10)
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
	EndPoint = file.Section("qiniu").Key("EndPoint").String()

}

func LoadMachinery(file *ini.File) {
	Broker = file.Section("machinery").Key("Broker").MustString("amqp://machinery:Cp123456@localhost:5672/go-machinery")
	RedisBackendHost = file.Section("machinery").Key("RedisBackendHost").MustString("127.0.0.1")
	RedisBackendPort = file.Section("machinery").Key("RedisBackendPort").MustString("6379")
	RedisBackendAuth = file.Section("machinery").Key("RedisBackendAuth").MustString("")
	RedisBackendDB = file.Section("machinery").Key("RedisBackendDB").MustInt(15)
}

func LoadEtcd(file *ini.File) {
	EtcdHosts = file.Section("etcd").Key("EtcdHosts").MustString("127.0.0.1:2379|127.0.0.1:2380")
	RpcPort = file.Section("etcd").Key("RpcPort").MustString(":50051")
}
