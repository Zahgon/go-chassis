package servicecomb

import (
	"github.com/go-chassis/go-archaius/event"
)

// constants for loadbalancer strategy name, and timeout
const (
	//LoadBalanceKey is variable of type string that matches load balancing events
	LoadBalanceKey = "^cse\\.loadbalance\\."
)

// LoadBalancingEventListener is a struct
type LoadBalancingEventListener struct {
	Key string
}

// Event is a method used to handle a load balancing event
func (e *LoadBalancingEventListener) Event(evt *event.Event) { _ = "STUB: not implemented"; return }
