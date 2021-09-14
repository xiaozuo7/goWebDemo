package test

import (
	"fmt"
	"goWebDemo/utils"
	"goWebDemo/utils/redis_client"
	"testing"
)

//  hash 键、值
func TestRedisHashKey(t *testing.T) {
	utils.LoadConfig()
	redis_client.InitRedis()
	redisClient := redis_client.GetOneRedisClient()
	defer redisClient.ReleaseOneRedisClient()

	res, err := redisClient.Int64(redisClient.Execute("hSet", "hField1", "hKey1", "hash_value_1"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	} else {
		fmt.Println(res)
	}
	res2, err := redisClient.String(redisClient.Execute("hGet", "hField1", "hKey1"))
	if err != nil {
		t.Errorf("单元测试失败,%s\n", err.Error())
	}
	fmt.Println(res2)
}
