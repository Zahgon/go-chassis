package config

import (
	"sync"

	"time"

	"github.com/go-chassis/go-chassis/v2/core/config/model"
)

// constant for hystrix parameters
const (
	DefaultForceFallback                 = false
	DefaultTimeoutEnabled                = false
	DefaultConsumerCircuitBreakerEnabled = false
	DefaultProviderCircuitBreakerEnabled = false
	DefaultCircuitBreakerForceOpen       = false
	DefaultCircuitBreakerForceClosed     = false
	DefaultFallbackEnable                = true
	DefaultMaxConcurrent                 = 1000
	DefaultSleepWindow                   = 15000
	DefaultTimeout                       = 30000
	DefaultErrorPercentThreshold         = 50
	DefaultRequestVolumeThreshold        = 20
	PolicyNull                           = "returnnull"
	PolicyThrowException                 = "throwexception"
)

var cbMutex = sync.RWMutex{}

// GetFallbackEnabled get fallback enabled
func GetFallbackEnabled(command, t string) bool { _ = "STUB: not implemented"; return false }

// GetCircuitBreakerEnabled get circuit breaker enabled
func GetCircuitBreakerEnabled(command, t string) bool { _ = "STUB: not implemented"; return false }

// GetForceClose get force close
func GetForceClose(service, t string) bool { _ = "STUB: not implemented"; return false }

// GetForceOpen get foce open
func GetForceOpen(service, t string) bool { _ = "STUB: not implemented"; return false }

// GetTimeout get timeout durations
func GetTimeout(service, t string) int { _ = "STUB: not implemented"; return 0 }

// GetTimeoutDuration get timeout durations from cache first, then get from archaius
func GetTimeoutDuration(service, t string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetMaxConcurrentRequests get max concurrent requests
func GetMaxConcurrentRequests(command, t string) int { _ = "STUB: not implemented"; return 0 }

// GetErrorPercentThreshold get error percent threshold
func GetErrorPercentThreshold(command, t string) int { _ = "STUB: not implemented"; return 0 }

// GetRequestVolumeThreshold get request volume threshold
func GetRequestVolumeThreshold(command, t string) int { _ = "STUB: not implemented"; return 0 }

// GetSleepWindow get sleep window
func GetSleepWindow(command, t string) int { _ = "STUB: not implemented"; return 0 }

// GetPolicy get fallback policy
func GetPolicy(service, t string) string { _ = "STUB: not implemented"; return "" }

func getIsolationSpec(command string) *model.IsolationSpec { _ = "STUB: not implemented"; return nil }

func getCircuitBreakerSpec(command string) *model.CircuitBreakerSpec {
	_ = "STUB: not implemented"
	return nil
}

func getFallbackSpec(command string) *model.FallbackSpec { _ = "STUB: not implemented"; return nil }

func getFallbackPolicySpec(command string) *model.FallbackPolicySpec {
	_ = "STUB: not implemented"
	return nil
}

// GetForceFallback get force fallback
func GetForceFallback(service, t string) bool { _ = "STUB: not implemented"; return false }
