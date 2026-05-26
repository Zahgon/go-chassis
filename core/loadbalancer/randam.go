package loadbalancer

import (
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/registry"

	"sync"
)

// RandomStrategy is strategy
type RandomStrategy struct {
	instances []*registry.MicroServiceInstance
	mtx       sync.Mutex
}

func newRandomStrategy() Strategy {
	_ = "STUB: not implemented"
	return *

	// ReceiveData receive data
	new(Strategy)
}

func (r *RandomStrategy) ReceiveData(inv *invocation.Invocation, instances []*registry.MicroServiceInstance, serviceName string) {
	_ = "STUB: not implemented"
	return

	// Pick return instance
}

func (r *RandomStrategy) Pick() (*registry.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
