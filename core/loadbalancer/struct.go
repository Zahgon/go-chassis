package loadbalancer

import (
	"time"
)

// ProtocolStats store protocol stats
type ProtocolStats struct {
	Latency    []time.Duration
	Addr       string
	AvgLatency time.Duration
}

// CalculateAverageLatency make avg latency
func (ps *ProtocolStats) CalculateAverageLatency() { _ = "STUB: not implemented"; return }

// SaveLatency save latest 10 record
func (ps *ProtocolStats) SaveLatency(l time.Duration) { _ = "STUB: not implemented"; return }

//save latest 10 latencies
