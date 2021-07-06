# goWebDemo
基于gin的web demo练习

## 项目框架简介

+ api --> 接口
+ middlerware --> 中间件
+ model --> 数据库表结构
+ router --> 路由
+ service --> 服务层crud等
+ test --> 单元测试
+ utils --> 三方组件
  - main.go --> 项目启动窗口

## 数据库

具体实现参考`model` --> `db.go`

[gorm项目地址](https://github.com/go-gorm/gorm)

[gorm文档](https://gorm.io/docs/)

## 中间件

具体实现参考`middlerware`

+ Jwt  (认证) `middlerware`-->`jwt.go`               [jwt项目地址](https://github.com/dgrijalva/jwt-go)
+ logrus (日志) `middlerware`-->`logger.go`    [logrus项目地址](https://github.com/sirupsen/logrus)
+ cors(跨域) `middlerware`-->`cors.go`            [cors项目地址](https://github.com/gin-contrib/cors)

## 三方组件

三方组件具体实现参考`utils`

+ Redis -- 传统NoSql  `utils`-->`redis_client`   [redigo项目地址](github.com/gomodule/redigo)

+ Validator -- 表单验证器  `utils`-->`validator`  [Validator项目地址](github.com/go-playground/validator/v10)

+ ini -- 配置读取(具体Demo参考后文)     `utils`-->`setting.go`  [ini项目地址](github.com/go-ini/ini)
+ etcd -- 分布式高可用强一致性k-v数据库  `utils`-->`etcdctl`  [etcd项目地址](https://github.com/etcd-io/etcd)
+ machinery -- 基于分布式的异步任务     `utils`-->`machinery` [machinery项目地址](github.com/RichardKnop/machinery/v2)
+ grpc -- 远程过程调用,高性能的、开源的通用的RPC框架 `utils`-->`grpc`  [grpc项目地址](google.golang.org/grpc)

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

实现参考`service`-->`upload_file.go`

## 配置文件

项目下创建`config.ini`, 具体配置参数含义请查看`utils`-->`settting.go`

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

## Docker挂载

+ mysql

```shell
sudo docker run --name {$databases_alias} -d -p 3306:3306 -v /home/{$user}/docker_volume/mysql/:/var/lib/mysql -e MYSQL_ROOT_PASSWORD={$pswd}   # databases_alias为数据库别名、user为本机用户名、pswd为数据库密码 
```

+ redis

```shell
sudo docker run --name {$databases_alias} -d -p 6379:6379 -v /home/{$user}/docker_volume/redis:/data redis redis-server --appendonly yes --requirepass {$pswd} # databases_alias为数据库别名 user为用户名 pswd为数据库密码 
```

+ etcd

[etcd_docker启动](https://github.com/etcd-io/etcd/releases/tag/v3.5.0)

```shell
  sudo docker run \
  -p 2379:2379 \
  -p 2380:2380 \
  --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
  --name etcd-gcr-v3.5.0 \
  quay.io/coreos/etcd:v3.5.0 \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr
```





