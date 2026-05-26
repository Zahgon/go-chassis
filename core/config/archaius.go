package config

import (
	"time"
)

// InitArchaius initialize the archaius
func InitArchaius() error { _ = "STUB: not implemented"; return nil }

// GetTimeoutDurationFromArchaius get timeout durations from archaius
func GetTimeoutDurationFromArchaius(service, t string) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
