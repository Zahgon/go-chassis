package rolling

import (
	"sync"
	"time"
)

// Number tracks a numberBucket over a bounded number of
// time buckets. Currently the buckets are one second long and only the last 10 seconds are kept.
type Number struct {
	Buckets map[int64]*numberBucket
	Mutex   *sync.RWMutex
}

type numberBucket struct {
	Value float64
}

// NewNumber initializes a RollingNumber struct.
func NewNumber() *Number { _ = "STUB: not implemented"; return nil }

func (r *Number) getCurrentBucket() *numberBucket { _ = "STUB: not implemented"; return nil }

func (r *Number) removeOldBuckets() { _ = "STUB: not implemented"; return }

// TODO: configurable rolling window

// Increment increments the number in current timeBucket.
func (r *Number) Increment(i float64) { _ = "STUB: not implemented"; return }

// UpdateMax updates the maximum value in the current bucket.
func (r *Number) UpdateMax(n float64) { _ = "STUB: not implemented"; return }

// Sum sums the values over the buckets in the last 10 seconds.
func (r *Number) Sum(now time.Time) float64 { _ = "STUB: not implemented"; return 0 }

// TODO: configurable rolling window

// Max returns the maximum value seen in the last 10 seconds.
func (r *Number) Max(now time.Time) float64 { _ = "STUB: not implemented"; return 0 }

// TODO: configurable rolling window

func (r *Number) Avg(now time.Time) float64 { _ = "STUB: not implemented"; return 0 }
