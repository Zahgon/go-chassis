package loadbalancer

import (
	"sync"

	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/registry"
)

var (

	// successiveFailureCount success and failure count
	successiveFailureCount      map[string]int
	successiveFailureCountMutex sync.RWMutex
)

func init() {
	successiveFailureCount = make(map[string]int)
}

// DeleteSuccessiveFailureCount deleting cookie from failure count map
func DeleteSuccessiveFailureCount(cookieValue string) { _ = "STUB: not implemented"; return }

//	successiveFailureCount[ep] = 0

// ResetSuccessiveFailureMap make map again
func ResetSuccessiveFailureMap() { _ = "STUB: not implemented"; return }

// IncreaseSuccessiveFailureCount increase failure count
func IncreaseSuccessiveFailureCount(cookieValue string) { _ = "STUB: not implemented"; return }

// GetSuccessiveFailureCount get failure count
func GetSuccessiveFailureCount(cookieValue string) int { _ = "STUB: not implemented"; return 0 }

// SessionStickinessStrategy is strategy
type SessionStickinessStrategy struct {
	instances []*registry.MicroServiceInstance
	mtx       sync.Mutex
	sessionID string
}

func newSessionStickinessStrategy() Strategy { _ = "STUB: not implemented"; return *new(Strategy) }

// ReceiveData receive data
func (r *SessionStickinessStrategy) ReceiveData(inv *invocation.Invocation, instances []*registry.MicroServiceInstance, serviceName string) {
	_ = "STUB: not implemented"
	return
}

func getNamespace(i *invocation.Invocation) string { _ = "STUB: not implemented"; return "" }

// Pick return instance
func (r *SessionStickinessStrategy) Pick() (*registry.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if micro service instance goes down then related entry in endpoint map will be deleted,
//so instead of sending nil, a new instance can be selected using round robin

func (r *SessionStickinessStrategy) pick() (*registry.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
