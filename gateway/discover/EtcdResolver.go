package discover

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

type EtcdResolver struct {
	Client *clientv3.Client
}

func NewEtcdResolver() *EtcdResolver {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("客户获取etcd错误", err)
		panic(err)
	}
	return &EtcdResolver{
		Client: etcdCli,
	}
}

/*
	BuildEtcd 该方法用于创建一个etcd解析器，grpc.Dial()会调用该方法，解析器根据key前缀监听etcd服务列表里的变化并跟新本地
	列表watch,然后注册解析器，创建grpc句柄，使用负载均衡轮询请求服务
*/
func (e *EtcdResolver) BuildEtcd(target resolver.Target, cc resolver.ClientConn) (resolver.Resolver, error) {
	//指定获取前缀的etcd节点值
	prefix := "/" + target.Scheme
	log.Println(prefix)
	res, err := e.Client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		log.Println("get address failed ", err)
		panic(err)
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	es := &EtcdDiscover{
		cancel:     cancelFunc,
		ctx:        ctx,
		clientConn: cc,
		etcdClient: e.Client,
		schema:     target.Scheme,
	}
	log.Printf("etcd res:%v\n", res)
	for _, v := range res.Kvs {
		es.Stores(v.Key, v.Value)
	}
	es.UpdateState()
	go es.Watcher()
	return es, nil
}

func (e *EtcdResolver) Schema() string {
	return "etcd"
}
