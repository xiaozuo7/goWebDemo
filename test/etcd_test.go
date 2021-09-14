package test

import (
	"context"
	"fmt"
	"goWebDemo/utils"
	"goWebDemo/utils/etcdctl"
	"testing"
	"time"
)

//  Etcd k-v存储测试
func TestEtcdGet(t *testing.T) {
	utils.LoadConfig()
	etcdctl.InitEtcd()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	putRes, err := etcdctl.EtcdCtl.KV.Put(ctx, "sample_test", "I am res")
	if err != nil {
		t.Errorf("设置值失败, %s\n", err.Error())
	} else {
		fmt.Printf("PutResponse: %v\n", putRes)
	}

	tmp, err := etcdctl.EtcdCtl.KV.Get(ctx, "sample_test")
	if err != nil {
		t.Errorf("获取值失败, %s\n", err.Error())

	} else {
		fmt.Printf("GetResponse: %v\n", tmp.Kvs)
	}
}
