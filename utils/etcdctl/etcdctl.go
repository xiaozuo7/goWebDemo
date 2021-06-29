package etcdctl

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"goWebDemo/utils"
	"strings"
	"time"
)

var EtcdCtl *clientv3.Client
var err error

func InitEtcd() {
	etcdString := utils.EtcdHosts
	etcdArr := strings.Split(etcdString, "|")
	EtcdCtl, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdArr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic("连接Etcd出错：" + err.Error())
	}
}
