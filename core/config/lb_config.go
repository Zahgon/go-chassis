package config

import (
	"sync"
)

const (
	lbPrefix                                 = "cse.loadbalance"
	propertySessionStickinessRuleTimeout     = "SessionStickinessRule.sessionTimeoutInSeconds"
	propertySessionStickinessRuleFailedTimes = "SessionStickinessRule.successiveFailedTimes"
	propertyRetryEnabled                     = "retryEnabled"
	propertyRetryOnNext                      = "retryOnNext"
	propertyRetryOnSame                      = "retryOnSame"
	propertyBackoffMinMs                     = "backoff.minMs"
	propertyBackoffMaxMs                     = "backoff.maxMs"

	//DefaultStrategy is default value for strategy
	DefaultStrategy = "RoundRobin"
	//DefaultSessionTimeout is default value for timeout
	DefaultSessionTimeout = 30
	//DefaultFailedTimes is default value for failed times
	DefaultFailedTimes = 5
)

var lbMutex = sync.RWMutex{}

func genKey(s ...string) string { _ = "STUB: not implemented"; return "" }

// GetServerListFilters get server list filters
func GetServerListFilters() (filters []string) { _ = "STUB: not implemented"; return nil }

// GetStrategyName get strategy name
func GetStrategyName(service string) string { _ = "STUB: not implemented"; return "" }

// GetSessionTimeout return session timeout
func GetSessionTimeout(source, service string) int { _ = "STUB: not implemented"; return 0 }

// StrategySuccessiveFailedTimes strategy successive failed times
func StrategySuccessiveFailedTimes(source, service string) int { _ = "STUB: not implemented"; return 0 }

// RetryEnabled retry enabled
func RetryEnabled(source, service string) bool { _ = "STUB: not implemented"; return false }

// GetRetryOnNext return value of GetRetryOnNext
func GetRetryOnNext(source, service string) int { _ = "STUB: not implemented"; return 0 }

// GetRetryOnSame return value of RetryOnSame
func GetRetryOnSame(source, service string) int { _ = "STUB: not implemented"; return 0 }

// BackOffKind get kind
func BackOffKind(service string) string { _ = "STUB: not implemented"; return "" }

// BackOffMinMs get min time
func BackOffMinMs(source, service string) int { _ = "STUB: not implemented"; return 0 }

// BackOffMaxMs get max time
func BackOffMaxMs(source, service string) int { _ = "STUB: not implemented"; return 0 }
