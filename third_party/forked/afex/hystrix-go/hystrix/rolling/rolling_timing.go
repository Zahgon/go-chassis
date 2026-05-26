package rolling

import (
	"sync"
	"time"
)

// Timing maintains time Durations for each time bucket.
// The Durations are kept in an array to allow for a variety of
// statistics to be calculated from the source data.
type Timing struct {
	Buckets map[int64]*timingBucket
	Mutex   *sync.RWMutex

	CachedSortedDurations []time.Duration
	LastCachedTime        int64
}

type timingBucket struct {
	Durations []time.Duration
}

// NewTiming creates a RollingTiming struct.
func NewTiming() *Timing { _ = "STUB: not implemented"; return nil }

type byDuration []time.Duration

func (c byDuration) Len() int      { _ = "STUB: not implemented"; return 0 }
func (c byDuration) Swap(i, j int) { _ = "STUB: not implemented"; return }
func (c byDuration) Less(i, j int) bool {
	_ = "STUB: not implemented"

	// SortedDurations returns an array of time.Duration sorted from shortest
	// to longest that have occurred in the last 60 seconds.
	return false
}

func (r *Timing) SortedDurations() []time.Duration { _ = "STUB: not implemented"; return nil }

// don't recalculate if current cache is still fresh

// TODO: configurable rolling window

func (r *Timing) getCurrentBucket() *timingBucket { _ = "STUB: not implemented"; return nil }

func (r *Timing) removeOldBuckets() { _ = "STUB: not implemented"; return }

// TODO: configurable rolling window

// Add appends the time.Duration given to the current time bucket.
func (r *Timing) Add(duration time.Duration) { _ = "STUB: not implemented"; return }

// Percentile computes the percentile given with a linear interpolation.
// it returns million seconds
func (r *Timing) Percentile(p float64) uint32 { _ = "STUB: not implemented"; return 0 }

func (r *Timing) ordinal(length int, percentile float64) int64 { _ = "STUB: not implemented"; return 0 }

// Mean computes the average timing in the last 60 seconds.
func (r *Timing) Mean() uint32 { _ = "STUB: not implemented"; return 0 }
