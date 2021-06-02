package redis_client

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"goWebDemo/utils"
	"time"
)

var redisPool *redis.Pool

// InitRedis 初始化redis
func InitRedis() *redis.Pool {
	redisPool = &redis.Pool{
		MaxIdle:     utils.MaxIdle,
		MaxActive:   utils.MaxActive,
		IdleTimeout: utils.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", utils.RedisHost+":"+utils.RedisPort)
			if err != nil {
				fmt.Println("初始化redis失败，请检查参数是否正确: ", err)
				return nil, err
			}

			auth := utils.RedisAuth
			if len(auth) >= 1 {
				if _, err := conn.Do("AUTH", auth); err != nil {
					_ = conn.Close()
					fmt.Println("Redis Auth 鉴权失败，密码错误", err)
				}
			}
			_, _ = conn.Do("select", utils.RedisDb)
			return conn, err
		},
	}
	return redisPool
}

// GetOneRedisClient 从连接池拿连接
func GetOneRedisClient() *RedisClient {
	maxRetryTimes := utils.ConnFailRetryTimes
	var oneConn redis.Conn
	for i := 1; i <= maxRetryTimes; i++ {
		oneConn = redisPool.Get()
		if oneConn.Err() != nil {
			if i == maxRetryTimes {
				fmt.Println("Redis 从连接池获取一个连接失败，超过最大重试次数")
				return nil
			}
			time.Sleep(utils.ReConnectInterval)
		} else {
			break
		}
	}
	return &RedisClient{oneConn}
}

// RedisClient redis客户端结构体
type RedisClient struct {
	client redis.Conn
}

// ReleaseOneRedisClient 释放连接到连接池
func (r *RedisClient) ReleaseOneRedisClient() {
	_ = r.client.Close()
}

// Execute 封装统一操作函数入口
func (r *RedisClient) Execute(cmd string, args ...interface{}) (interface{}, error) {
	return r.client.Do(cmd, args...)
}

// Bool 类型转换
func (r *RedisClient) Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}

// String
func (r *RedisClient) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

// Strings 类型转换
func (r *RedisClient) Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

// Float64 类型转换
func (r *RedisClient) Float64(reply interface{}, err error) (float64, error) {
	return redis.Float64(reply, err)
}

// Int 类型转换
func (r *RedisClient) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

// Int64 类型转换
func (r *RedisClient) Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

// Uint64 类型转换
func (r *RedisClient) Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}

// Bytes 类型转换
func (r *RedisClient) Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}
