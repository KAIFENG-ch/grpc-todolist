package loadBalance

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"math/rand"
	"sync"
)

const (
	WEIGHT_LOAD_BALANCE = "weight_load_balance"
	MAX_WEIGHT          = 10
	MIN_WEIGHT          = 1
)

type V2PickerBuild struct {
}

type WeightPicker struct {
	scs []balancer.SubConn
	mu  sync.Mutex
}

type WeightAttributeKey struct {
}

type WeightAddrInfo struct {
	Weight int
}

func (j V2PickerBuild) Build(info base.PickerBuildInfo) balancer.V2Picker {
	// 没有可用的连接
	if len(info.ReadySCs) == 0 {
		return base.NewErrPickerV2(balancer.ErrNoSubConnAvailable)
	}
	scs := make([]balancer.SubConn, 0, len(info.ReadySCs))

	for subConn, subConnInfo := range info.ReadySCs {
		v := subConnInfo.Address.Attributes.Value(WeightAttributeKey{})
		w := v.(WeightAddrInfo).Weight
		// 限制可以设置的最大最小权重，防止设置过大创建连接数太多
		if w < MIN_WEIGHT {
			w = MIN_WEIGHT
		}

		if w > MAX_WEIGHT {
			w = MAX_WEIGHT
		}
		// 根据权重 创建多个重复的连接 权重越高个数越多
		for i := 0; i < w; i++ {
			scs = append(scs, subConn)
		}

	}
	return &WeightPicker{
		scs: scs,
	}
}

func (p *WeightPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {

	// 随机选择一个返回，权重越大，生成的连接个数越多，因此，被选中的概率也越大
	p.mu.Lock()
	index := rand.Intn(len(p.scs))
	sc := p.scs[index]
	p.mu.Unlock()
	return balancer.PickResult{SubConn: sc}, nil
}

func newBuilder() balancer.Builder {
	return base.NewBalancerBuilderV2(WEIGHT_LOAD_BALANCE, &V2PickerBuild{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newBuilder())
}
