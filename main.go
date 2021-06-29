package main

import (
	"goWebDemo/model"
	"goWebDemo/router"
	"goWebDemo/utils/etcdctl"
	"goWebDemo/utils/redis_client"
)

func main() {
	model.InitDb()
	redis_client.InitRedis()
	router.InitRouter()
	etcdctl.InitEtcd()
}
