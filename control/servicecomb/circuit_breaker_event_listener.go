package servicecomb

import (
	"github.com/go-chassis/go-archaius/event"
)

// constants for consumer isolation, circuit breaker, fallback keys
const (
	// ConsumerIsolationKey is a variable of type string
	ConsumerIsolationKey      = "cse.isolation"
	ConsumerCircuitBreakerKey = "cse.circuitBreaker"
	ConsumerFallbackKey       = "cse.fallback"
	ConsumerFallbackPolicyKey = "cse.fallbackpolicy"
	regex4normal              = "cse\\.(isolation|circuitBreaker|fallback|fallbackpolicy)\\.Consumer\\.(.*)\\.(timeout|timeoutInMilliseconds|maxConcurrentRequests|enabled|forceOpen|forceClosed|sleepWindowInMilliseconds|requestVolumeThreshold|errorThresholdPercentage|enabled|maxConcurrentRequests|policy)\\.(.+)"
)

// CircuitBreakerEventListener is a struct with one string variable
type CircuitBreakerEventListener struct {
	Key string
}

// Event is a method which triggers flush circuit
func (el *CircuitBreakerEventListener) Event(e *event.Event) { _ = "STUB: not implemented"; return }

// FlushCircuitByKey is a function used to flush for a particular key
func FlushCircuitByKey(key string) { _ = "STUB: not implemented"; return }

// GetNames is function
func GetNames(key string) (string, string) { _ = "STUB: not implemented"; return "", "" }

// GetCircuitName is a function used to get circuit names
func GetCircuitName(sourceName, serviceName string) string { _ = "STUB: not implemented"; return "" }
