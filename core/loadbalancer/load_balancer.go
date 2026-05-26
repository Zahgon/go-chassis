// Package loadbalancer is client side load balancer
package loadbalancer

import (
	"sync"
	"time"

	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/registry"
)

// constant string for zoneaware
const (
	ZoneAware = "zoneaware"
)

// StrategyLatency is name
const StrategyLatency = "WeightedResponse"

// constant strings for load balance variables
const (
	StrategyRoundRobin        = "RoundRobin"
	StrategyRandom            = "Random"
	StrategySessionStickiness = "SessionStickiness"

	OperatorEqual   = "="
	OperatorGreater = ">"
	OperatorSmaller = "<"
	OperatorPattern = "Pattern"
)

var (
	// ErrNoneAvailableInstance is to represent load balance error
	ErrNoneAvailableInstance = LBError{Message: "None available instance"}
)

// LBError load balance error
type LBError struct {
	Message string
}

// Error for to return load balance error message
func (e LBError) Error() string { _ = "STUB: not implemented"; return "" }

// BuildStrategy query instance list and give it to Strategy then return Strategy
func BuildStrategy(i *invocation.Invocation,
	s Strategy) (Strategy, error) {
	_ = "STUB: not implemented"
	return *new(Strategy), nil
}

//append filters in config

// Strategy is load balancer algorithm , call Pick to return one instance
type Strategy interface {
	ReceiveData(inv *invocation.Invocation, instances []*registry.MicroServiceInstance, serviceKey string)
	Pick() (*registry.MicroServiceInstance, error)
}

// Criteria is rule for filter
type Criteria struct {
	Key      string
	Operator string
	Value    string
}

// Filter receive instances and criteria, it will filter instances based on criteria you defined,criteria is optional, you can give nil for it
type Filter func(instances []*registry.MicroServiceInstance, criteria []*Criteria) []*registry.MicroServiceInstance

// Enable function is for to enable load balance strategy
func Enable(strategyName string) error { _ = "STUB: not implemented"; return nil }

// Filters is a map of string and array of *registry.MicroServiceInstance
var Filters = make(map[string]Filter)

// InstallFilter install filter
func InstallFilter(name string, f Filter) { _ = "STUB: not implemented"; return }

// variables for latency map, rest and highway requests count
var (
	//ProtocolStatsMap saves all stats for all service's protocol, one protocol has a lot of instances
	ProtocolStatsMap = make(map[string][]*ProtocolStats)
	//maintain different locks since multiple goroutine access the map
	LatencyMapRWMutex sync.RWMutex
)

// BuildKey return key of stats map
func BuildKey(microServiceName, tags, protocol string) string {
	_ = "STUB: not implemented"
	// TODO add more data
	return ""
}

// SetLatency for a instance, it only save latest 10 stats for instance's protocol
func SetLatency(latency time.Duration, addr, microServiceName string, tags utiltags.Tags, protocol string) {
	_ = "STUB: not implemented"
	return
}
