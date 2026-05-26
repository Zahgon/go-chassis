package config

import (
	"time"
)

// constant for default values of abort and delay
const (
	DefaultAbortPercent = 0
	DefaultAbortStatus  = 0
	DefaultDelayPercent = 0
)

// GetAbortPercent get abort percentage
func GetAbortPercent(protocol, microServiceName, schema, operation string) int {
	_ = "STUB: not implemented"
	return 0
}

// GetAbortStatus get abort status
func GetAbortStatus(protocol, microServiceName, schema, operation string) int {
	_ = "STUB: not implemented"
	return 0
}

// GetDelayPercent get delay percentage
func GetDelayPercent(protocol, microServiceName, schema, operation string) int {
	_ = "STUB: not implemented"
	return 0
}

// GetFixedDelay get fixed delay
func GetFixedDelay(protocol, microServiceName, schema, operation string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
