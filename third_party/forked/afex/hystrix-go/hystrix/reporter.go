package hystrix

import (
	"errors"
)

// Reporter receive a circuit breaker Metrics and sink it to monitoring system
type Reporter func(cb *CircuitBreaker) error

// ErrDuplicated means you can not install reporter with same name
var ErrDuplicated = errors.New("duplicated reporter")
var reporterPlugins = make(map[string]Reporter)

// InstallReporter install reporter implementation
// it receives a circuit breaker and sink its Metrics to monitoring system
func InstallReporter(name string, reporter Reporter) error { _ = "STUB: not implemented"; return nil }

// StartReporter starts reporting to reporters
func StartReporter() { _ = "STUB: not implemented"; return }
