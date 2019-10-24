package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"188.131.171.145:2379", "188.131.171.145:2379"},
		AutoSyncInterval:     0,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		LogConfig:            nil,
		Context:              nil,
		PermitWithoutStream:  false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("connect ok")

	//err1 := cli.Close()
	//fmt.Println(err1)
	//设置1s超时，访问etcd 有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	_, err = cli.Put(ctx, "/a/b/c", "sample")

	defer cancel()

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)

	res, err := cli.Get(ctx, "/a/b/c")

	defer cancel()

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range res.Kvs {
		fmt.Println(string(v.Key), string(v.Value))
	}

}
