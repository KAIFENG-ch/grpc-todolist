package service

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type Register struct {
	etcdCli *clientv3.Client
	leaseId clientv3.LeaseID
	ctx context.Context
}

func RegisterServer(k,v string, expire int64) (*Register, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err)
		return nil,err
	}
	server := &Register{
		etcdCli: etcdClient,
		ctx: context.Background(),
	}
	err = server.CreateLease(expire)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = server.BuildLease(k, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = server.KeepAlive()
	if err != nil {
		log.Print("租约续期失败", err)
		return nil, err
	}
	return server, nil
}

func (r Register) CreateLease(expire int64) error {
	res, err := r.etcdCli.Grant(r.ctx, expire)
	if err != nil {
		log.Println(err)
		return err
	}
	r.leaseId = res.ID
	return nil
}

func (r Register) BuildLease(k, v string) error {
	res, err := r.etcdCli.Put(r.ctx, k, v, clientv3.WithLease(r.leaseId))
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(res)
	return nil
}

func (r Register) KeepAlive() error {
	leaseResCh, err := r.etcdCli.KeepAlive(r.ctx, r.leaseId)
	if err != nil {
		log.Println(err)
		return err
	}
	go r.Watch(leaseResCh)
	return nil
}

func (r Register) Watch(leaseCh <-chan *clientv3.LeaseKeepAliveResponse) {
	for k := range leaseCh {
		log.Printf("续约成功；%+v", k)
	}
	log.Println("续约租期关闭")
}