package register

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type Register struct {
	etcdCli *clientv3.Client
	leaseId clientv3.LeaseID
	ctx     context.Context
	cancel  context.CancelFunc
}

// CreateLease 创建租约
func (r *Register) CreateLease(expire int64) error {
	res, err := r.etcdCli.Grant(r.ctx, expire)
	if err != nil {
		log.Println(err)
		return err
	}
	r.leaseId = res.ID
	return nil
}

// BuildLease 将租约和对应的key, value绑定
func (r *Register) BuildLease(k, v string) error {
	res, err := r.etcdCli.Put(r.ctx, k, v, clientv3.WithLease(r.leaseId))
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(res)
	return nil
}

// KeepAlive 发送心跳包
func (r *Register) KeepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	leaseResCh, err := r.etcdCli.KeepAlive(r.ctx, r.leaseId)
	if err != nil {
		log.Printf("保持心跳失败，%v\n", leaseResCh)
		return nil, err
	}
	return leaseResCh, nil
}

// Watch 监控服务并续约
func (r *Register) Watch(key string, leaseCh <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case lease := <-leaseCh:
			log.Printf("续约成功；%+v\n", lease)
		case <-r.ctx.Done():
			log.Printf("续约失败")
			return
		}
	}
}

// Close 关闭服务
func (r *Register) Close() error {
	r.cancel()
	log.Println("租约关闭")
	// 撤销租约
	_, err := r.etcdCli.Revoke(r.ctx, r.leaseId)
	if err != nil {
		return err
	}
	return r.etcdCli.Close()
}

func NewEtcdReg() (*Register, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("创建服务失败，%v", err)
		return nil, err
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	server := &Register{
		etcdCli: client,
		ctx:     ctx,
		cancel:  cancelFunc,
	}
	return server, nil
}

func (r *Register) RegisterServer(serverName, addr string, expire int64) (err error) {
	err = r.CreateLease(expire)
	if err != nil {
		return
	}
	err = r.BuildLease(serverName, addr)
	if err != nil {
		return
	}
	heartBeat, err := r.KeepAlive()
	if err != nil {
		return
	}
	go r.Watch(serverName, heartBeat)
	return nil
}
