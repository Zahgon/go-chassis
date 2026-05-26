package metricCollector

import (
	"sync"
	"time"

	"github.com/go-chassis/go-chassis/v2/third_party/forked/afex/hystrix-go/hystrix/rolling"
)

// DefaultMetricCollector holds information about the circuit state.
// This implementation of MetricCollector is the canonical source of information about the circuit.
// It is used for for all internal hystrix operations
// including circuit health checks and metrics sent to the hystrix dashboard.
//
// Metric Collectors do not need Mutexes as they are updated by circuits within a locked context.
type DefaultMetricCollector struct {
	mutex       *sync.RWMutex
	name        string
	numRequests *rolling.Number
	errors      *rolling.Number

	successes     *rolling.Number
	failures      *rolling.Number
	rejects       *rolling.Number
	shortCircuits *rolling.Number
	timeouts      *rolling.Number

	fallbackSuccesses *rolling.Number
	fallbackFailures  *rolling.Number
	totalDuration     *rolling.Timing
	runDuration       *rolling.Timing
}

func newDefaultMetricCollector(name string) MetricCollector {
	_ = "STUB: not implemented"
	return *new(MetricCollector)
}

// NumRequests returns the rolling number of requests
func (d *DefaultMetricCollector) NumRequests() *rolling.Number {
	_ = "STUB: not implemented"
	return nil
}

// Errors returns the rolling number of errors
func (d *DefaultMetricCollector) Errors() *rolling.Number { _ = "STUB: not implemented"; return nil }

// Successes returns the rolling number of successes
func (d *DefaultMetricCollector) Successes() *rolling.Number { _ = "STUB: not implemented"; return nil }

// Failures returns the rolling number of failures
func (d *DefaultMetricCollector) Failures() *rolling.Number { _ = "STUB: not implemented"; return nil }

// Rejects returns the rolling number of rejects
func (d *DefaultMetricCollector) Rejects() *rolling.Number { _ = "STUB: not implemented"; return nil }

// ShortCircuits returns the rolling number of short circuits
func (d *DefaultMetricCollector) ShortCircuits() *rolling.Number {
	_ = "STUB: not implemented"
	return nil
}

// Timeouts returns the rolling number of timeouts
func (d *DefaultMetricCollector) Timeouts() *rolling.Number { _ = "STUB: not implemented"; return nil }

// FallbackSuccesses returns the rolling number of fallback successes
func (d *DefaultMetricCollector) FallbackSuccesses() *rolling.Number {
	_ = "STUB: not implemented"
	return nil
}

// FallbackFailures returns the rolling number of fallback failures
func (d *DefaultMetricCollector) FallbackFailures() *rolling.Number {
	_ = "STUB: not implemented"
	return nil
}

// TotalDuration returns the rolling total duration
func (d *DefaultMetricCollector) TotalDuration() *rolling.Timing {
	_ = "STUB: not implemented"
	return nil
}

// RunDuration returns the rolling run duration
func (d *DefaultMetricCollector) RunDuration() *rolling.Timing {
	_ = "STUB: not implemented"
	return nil
}

// IncrementAttempts increments the number of requests seen in the latest time bucket.
func (d *DefaultMetricCollector) IncrementAttempts() { _ = "STUB: not implemented"; return }

// IncrementErrors increments the number of errors seen in the latest time bucket.
// Errors are any result from an attempt that is not a success.
func (d *DefaultMetricCollector) IncrementErrors() { _ = "STUB: not implemented"; return }

// IncrementSuccesses increments the number of successes seen in the latest time bucket.
func (d *DefaultMetricCollector) IncrementSuccesses() { _ = "STUB: not implemented"; return }

// IncrementFailures increments the number of failures seen in the latest time bucket.
func (d *DefaultMetricCollector) IncrementFailures() { _ = "STUB: not implemented"; return }

// IncrementRejects increments the number of rejected requests seen in the latest time bucket.
func (d *DefaultMetricCollector) IncrementRejects() { _ = "STUB: not implemented"; return }

// IncrementShortCircuits increments the number of rejected requests seen in the latest time bucket.
func (d *DefaultMetricCollector) IncrementShortCircuits() { _ = "STUB: not implemented"; return }

// IncrementTimeouts increments the number of requests that timed out in the latest time bucket.
func (d *DefaultMetricCollector) IncrementTimeouts() { _ = "STUB: not implemented"; return }

// IncrementFallbackSuccesses increments the number of successful calls to the fallback function in the latest time bucket.
func (d *DefaultMetricCollector) IncrementFallbackSuccesses() { _ = "STUB: not implemented"; return }

// IncrementFallbackFailures increments the number of failed calls to the fallback function in the latest time bucket.
func (d *DefaultMetricCollector) IncrementFallbackFailures() { _ = "STUB: not implemented"; return }

// UpdateTotalDuration updates the total amount of time this circuit has been running.
func (d *DefaultMetricCollector) UpdateTotalDuration(timeSinceStart time.Duration) {
	_ = "STUB: not implemented"
	return
}

// UpdateRunDuration updates the amount of time the latest request took to complete.
func (d *DefaultMetricCollector) UpdateRunDuration(runDuration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Reset resets all metrics in this collector to 0.
func (d *DefaultMetricCollector) Reset() { _ = "STUB: not implemented"; return }
