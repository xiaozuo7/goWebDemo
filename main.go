package main

import (
	"goWebDemo/model"
	"goWebDemo/router"
	"goWebDemo/utils"
	"goWebDemo/utils/etcdctl"
	"goWebDemo/utils/redis_client"
)

func main() {
	utils.LoadConfig()
	model.InitDb()
	redis_client.InitRedis()
	router.InitRouter()
	etcdctl.InitEtcd()
}
