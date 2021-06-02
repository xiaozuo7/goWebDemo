package test

import (
	"fmt"
	"goWebDemo/utils/redis_client"
	"testing"
)

//  hash 键、值
func TestRedisHashKey(t *testing.T) {
	redis_client.InitRedis()
	redisClient := redis_client.GetOneRedisClient()

	//  hSet 返回 0 1
	res, err := redisClient.Int64(redisClient.Execute("hSet", "hField1", "hKey1", "hash_value_1"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	} else {
		fmt.Println(res)
	}
	//  hash键  get 命令，分为两步：1.执行get命令 2.将结果转为需要的格式
	res2, err := redisClient.String(redisClient.Execute("hGet", "hField1", "hKey1"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	}
	fmt.Println(res2)
	//官方明确说，redis使用完毕，必须释放
	redisClient.ReleaseOneRedisClient()
}
