package main

import (
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	_, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err)
	}
}
