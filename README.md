# goWebDemo
基于gin的web implement

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
+ Viper -- 配置读取(具体参考后面文档) `utils`-->`setting.go` [Viper项目地址](https://github.com/spf13/viper)
+ etcd -- 分布式高可用强一致性k-v数据库  `utils`-->`etcdctl`  [etcd项目地址](https://github.com/etcd-io/etcd)
+ Machinery -- 基于分布式的异步任务     `utils`-->`machinery` [Machinery项目地址](github.com/RichardKnop/machinery/v2)
+ GRPC -- 远程过程调用,高性能的、开源的通用的RPC框架 `utils`-->`grpc`  [GRPC项目地址](google.golang.org/grpc)

## 单元测试

单元测试具体参考`test`

>  注意单元测试读取配置文件时需要改成绝对路径 请在`utils` -->`setting.go`中修改

```go
   // go test时需要用绝对路径，比如
    viper.AddConfigPath("F://project//goWebDemo")

```

+ etcd存储测试 `etcd_test.go`
+ redis存储测试 `redis_test.go`

## 其他

1. bcrypt -- 密码加密算法

```go
// BcryptPw 密码加密
func BcryptPw(password string) string {
   const cost = 10
   HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
   if err != nil {
      log.Fatal(err)
   }
   return string(HashPw)
}
```
```go
// 密码校对
PasswordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
```


2. 文件上传 -- 接入的七牛云SDK

实现参考`service`-->`upload_file.go`


## Docker挂载

+ mysql

```shell
sudo docker run --name {$databases_alias} -d -p 3306:3306 -v /home/{$user}/docker_volume/mysql/:/var/lib/mysql -e MYSQL_ROOT_PASSWORD={$pswd} mysql:5.7   # databases_alias为数据库别名、user为本机用户名、pswd为数据库密码 
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

## 配置文件

根目录创建`config.yml`文件, 配置参数如下
```yaml
Server:
  # debug 开发模式，release 生产模式
  AppMode: "debug"
  HttpPort: ":3000"
  JwtKey: ""

Database:
  DbType: "mysql"
  Host: "127.0.0.1"
  Port: 3306
  User: "root"
  PassWord: ""
  Name: ""
  MaxOpenConns: 128
  MaxIdleConns: 10
  ConnMaxLifetime: 10

Redis:
  Host: "127.0.0.1"
  Port: 6379
  Auth: ""
  Db: 1
  MaxIdle: 10
  MaxActive: 1000
  IdleTimeout: 60
  ConnFailRetryTimes: 3
  ReConnectInterval: 1

Qiniu:
  AccessKey: ""
  SecretKey: ""
  Bucket: ""
  EndPoint: ""


Machinery:
  Broker: ""
  RedisBackendHost: "127.0.0.1"
  RedisBackendPort: 6379
  RedisBackendAuth: ""
  RedisBackendDB: 

Etcd:
  EtcdHosts: "127.0.0.1:2379|127.0.0.1:2380"
  RpcPort: ":50051"

```
