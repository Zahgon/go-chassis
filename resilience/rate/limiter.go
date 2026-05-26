// Package rate supply functionality about QPS
// for example rate limiting
package rate

import (
	"sync"

	"k8s.io/client-go/util/flowcontrol"
)

// constant qps default rate
const (
	DefaultRate = 2147483647
)

// Limiters manages all rate limiters. it is thread safe and singleton.
// it create new limiters and try to limit request.
// each limiter has a unique name.
type Limiters struct {
	sync.RWMutex
	m map[string]flowcontrol.RateLimiter
}

// variables of qps limiter and mutex variable
var (
	once       = new(sync.Once)
	qpsLimiter *Limiters
)

// GetRateLimiters get qps rate limiters
func GetRateLimiters() *Limiters { _ = "STUB: not implemented"; return nil }

// TryAccept try to accept a request. if limiter can not accept a request, it returns false
// name is the limiter name
// qps is not necessary if the limiter already exists
func (qpsL *Limiters) TryAccept(name string, qps, burst int) bool {
	_ = "STUB: not implemented"
	return false
}

//If the name operation is not present in the map, then add the new name operation to the map

// addLimiter create a new limiter and add it to limiter map
func (qpsL *Limiters) addLimiter(name string, qps, burst int) bool {
	_ = "STUB: not implemented"

	// add a limiter object for the newly found operation in the Default Hash map
	// so that the default rate will be applied to subsequent token requests to this new operation
	return false
}

// Create a new bucket for the new operation

// UpdateRateLimit will update the old limiters
func (qpsL *Limiters) UpdateRateLimit(name string, qps, burst int) {
	_ = "STUB: not implemented"
	return
}

// DeleteRateLimiter delete rate limiter
func (qpsL *Limiters) DeleteRateLimiter(name string) { _ = "STUB: not implemented"; return }
