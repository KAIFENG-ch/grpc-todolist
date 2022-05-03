package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	etcdReg, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("connect success")
	defer func(etcdReg *clientv3.Client) {
		err := etcdReg.Close()
		if err != nil {
			return
		}
	}(etcdReg)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = etcdReg.Put(ctx, "lmh", "lmh")
	cancel()
	if err != nil {
		panic(err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := etcdReg.Get(ctx, "lmh")
	cancel()
	if err != nil {
		panic(err)
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("%s,%s\n", kv.Key, kv.Value)
	}
}
