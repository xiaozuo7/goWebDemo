# goWebDemo
基于gin的web demo练习

## 数据库

具体实现参考`model` --> `db.go`

[gorm项目地址](https://github.com/go-gorm/gorm)

[gorm文档](https://gorm.io/docs/)

## 中间件

具体实现参考`middlerware`

+ Jwt  (认证) `middlerware`-->`jwt.go`  [jwt项目地址](https://github.com/dgrijalva/jwt-go)
+ logrus (日志) `middlerware`-->`logger.go` [logrus项目地址](https://github.com/sirupsen/logrus)
+ cors(跨域) `middlerware`-->`cors.go` [cors项目地址](https://github.com/gin-contrib/cors)

## 三方组件

三方组件具体实现参考`utils`

+ Redis -- 传统NoSql  [redigo项目地址](github.com/gomodule/redigo)   `utils`-->`redis_client`

+ Validator -- 表单验证器  [Validator项目地址](github.com/go-playground/validator/v10) `utils`-->`validator`

+ ini -- 配置读取  [ini项目地址](github.com/go-ini/ini) `utils`-->`setting.go`  具体Demo参考后文
+ etcd -- 分布式高可用强一致性k-v数据库  [etcd项目地址](https://github.com/etcd-io/etcd) `utils`-->`etcdctl`
+ machinery -- 基于分布式的异步任务 [machinery项目地址](github.com/RichardKnop/machinery/v2) `utils`-->`machinery`
+ grpc -- 远程过程调用,高性能的、开源的通用的RPC框架  [grpc项目地址](google.golang.org/grpc) `utils`-->`grpc`

## 单元测试

单元测试具体参考`test`

>  注意单元测试读取配置文件时需要改成绝对路径 请在`utils` -->`setting.go`中修改

```go
	// go test时需要用绝对路径F:\project\goWebDemo\config.ini
	file, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("配置文件读取出错,请检查文件路径是否正确", err)
	}
```

+ etcd存储测试 `etcd_test.go`
+ redis存储测试 `redis_test.go`

## 其他

1. bcrypt -- 密码加密算法

```go
// BcryptPw 密码加密
func ScryptPw(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}
```

2. 文件上传 -- 接入的七牛云SDK

实现参考`service`-->`upload_file`

## 配置文件

项目下创建`config.ini`

```ini
[server]
# debug 开发模式，release 生产模式
AppMode = debug
HttpPort = :3000
JwtKey = 

[database]
Db = mysql
DbHost = 127.0.0.1
DbPort = 3306
DbUser = root
DbPassWord = 
DbName = 
MaxOpenConns = 128
MaxIdleConns = 10
ConnMaxLifetime = 10

[redis]
RedisHost = 127.0.0.1
RedisPort = 6379
RedisAuth = 
RedisDb = 1
MaxIdle = 10
MaxActive = 1000
IdleTimeout = 60
ConnFailRetryTimes = 3
ReConnectInterval = 1

[qiniu]
AccessKey = 
SecretKey = 
Bucket = 
EndPoint = 


[machinery]
Broker = 
RedisBackendHost = 127.0.0.1
RedisBackendPort = 6379
RedisBackendAuth = 
RedisBackendDB = 15

[etcd]
EtcdHosts = 127.0.0.1:2379|127.0.0.1:2380
RpcPort = :50051
```

