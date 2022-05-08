package discover

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"
	"log"
	"sync"
)

type EtcdDiscover struct {
	ctx        context.Context
	cancel     context.CancelFunc
	ipPool     sync.Map
	clientConn resolver.ClientConn
	etcdClient *clientv3.Client
	schema     string
}

func (e *EtcdDiscover) ResolveNow(options resolver.ResolveNowOptions) {
	log.Println("etcd resolver now")
}

func (e *EtcdDiscover) Close() {
	log.Println("etcd resolver close")
	e.cancel()
}

func (e *EtcdDiscover) Watcher() {
	watchChan := e.etcdClient.Watch(context.Background(), "/"+e.schema, clientv3.WithPrefix())
	for true {
		select {
		case val := <-watchChan:
			for _, event := range val.Events {
				switch event.Type {
				case 0:
					e.Stores(event.Kv.Key, event.Kv.Value)
					log.Println("put:", string(event.Kv.Key))
					e.UpdateState()
				case 1:
					log.Println("del:", string(event.Kv.Key))
					e.Del(event.Kv.Key)
					e.UpdateState()
				}
			}
		}
	}
}

func (e *EtcdDiscover) Stores(k, v []byte) {
	e.ipPool.Store(string(k), string(v))
}

func (e *EtcdDiscover) Del(key []byte) {
	e.ipPool.Delete(string(key))
}

func (e *EtcdDiscover) UpdateState() {
	var addrList resolver.State
	log.Println("etcd升级")
	var i = 1
	e.ipPool.Range(func(key, value interface{}) bool {
		ta, ok := value.(string)
		if !ok {
			return false
		}
		log.Printf("conn.updatestate key[%v],value[%v]\n", key, value)
		addr := resolver.Address{
			Addr:       ta,
			Attributes: attributes.New(),
		}
		addrList.Addresses = append(addrList.Addresses, addr)
		i++
		return true
	})
	e.clientConn.UpdateState(addrList)
}
