package loadbalancer

import (
	"sync"

	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/registry"
)

// RoundRobinStrategy is strategy
type RoundRobinStrategy struct {
	instances []*registry.MicroServiceInstance
	key       string
}

func newRoundRobinStrategy() Strategy { _ = "STUB: not implemented"; return *new(Strategy) }

// ReceiveData receive data
func (r *RoundRobinStrategy) ReceiveData(inv *invocation.Invocation, instances []*registry.MicroServiceInstance, serviceKey string) {
	_ = "STUB: not implemented"
	return
}

// Pick return instance
func (r *RoundRobinStrategy) Pick() (*registry.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var rrIdxMap = make(map[string]int)
var mu sync.RWMutex

func pick(key string) int { _ = "STUB: not implemented"; return 0 }
