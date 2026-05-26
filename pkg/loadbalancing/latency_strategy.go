package loadbalancing

import (
	"sync"
	"time"

	"github.com/go-chassis/go-chassis/v2/core/config"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/loadbalancer"
	"github.com/go-chassis/go-chassis/v2/core/registry"
	"github.com/go-chassis/openlog"
)

var i int
var weightedRespMutex sync.Mutex

func init() {
	loadbalancer.InstallStrategy(loadbalancer.StrategyLatency, newWeightedResponseStrategy)
}

// ByDuration is for calculating the duration
type ByDuration []*loadbalancer.ProtocolStats

func (a ByDuration) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a ByDuration) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a ByDuration) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// SortLatency sort instance based on  the average latencies
func SortLatency() { _ = "STUB: not implemented"; return }

// CalculateAvgLatency Calculating the average latency for each instance using the statistics collected,
// key is addr/service/protocol
func CalculateAvgLatency() { _ = "STUB: not implemented"; return }

// WeightedResponseStrategy is a strategy plugin
type WeightedResponseStrategy struct {
	instances   []*registry.MicroServiceInstance
	serviceName string
	protocol    string
	tags        string
}

func init() {
	ticker := time.NewTicker(30 * time.Second)
	//run routine to prepare data
	go func() {
		for range ticker.C {
			if config.GetLoadBalancing() != nil {
				useLatencyAware := false
				for _, v := range config.GetLoadBalancing().AnyService {
					if v.Strategy["name"] == loadbalancer.StrategyLatency {
						useLatencyAware = true
						break
					}
				}
				if config.GetLoadBalancing().Strategy["name"] == loadbalancer.StrategyLatency {
					useLatencyAware = true
				}
				if useLatencyAware {
					CalculateAvgLatency()
					SortLatency()
					openlog.Info("Preparing data for Weighted Response Strategy")
				}
			}

		}
	}()
}
func newWeightedResponseStrategy() loadbalancer.Strategy {
	_ = "STUB: not implemented"
	return *new(loadbalancer.Strategy)
}

// ReceiveData receive data
func (r *WeightedResponseStrategy) ReceiveData(inv *invocation.Invocation, instances []*registry.MicroServiceInstance, serviceKey string) {
	_ = "STUB: not implemented"
	return
}

// Pick return instance
func (r *WeightedResponseStrategy) Pick() (*registry.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//if no instances are selected round robin will be done
