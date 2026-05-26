package hystrix

import (
	"sync"
	"time"

	"github.com/go-chassis/go-chassis/v2/third_party/forked/afex/hystrix-go/hystrix/rolling"
)

type commandExecution struct {
	Types       []string      `json:"types"`
	Start       time.Time     `json:"start_time"`
	RunDuration time.Duration `json:"run_duration"`
}

type metricExchange struct {
	Name    string
	Updates chan *commandExecution
	Mutex   *sync.RWMutex

	metricCollectors []metricCollector.MetricCollector
}

func newMetricExchange(name string, num int) *metricExchange { _ = "STUB: not implemented"; return nil }

// The Default Collector function will panic if collectors are not setup to specification.
func (m *metricExchange) DefaultCollector() *metricCollector.DefaultMetricCollector {
	_ = "STUB: not implemented"
	return nil
}

func (m *metricExchange) Monitor() { _ = "STUB: not implemented"; return }

// we only grab a read lock to make sure Reset() isn't changing the numbers.

func (m *metricExchange) IncrementMetrics(collector metricCollector.MetricCollector, update *commandExecution, totalDuration time.Duration) {
	_ = "STUB: not implemented"
	// granular Metrics
	return
}

// fallback Metrics

func (m *metricExchange) Reset() { _ = "STUB: not implemented"; return }

func (m *metricExchange) Requests() *rolling.Number { _ = "STUB: not implemented"; return nil }

func (m *metricExchange) requestsLocked() *rolling.Number { _ = "STUB: not implemented"; return nil }

func (m *metricExchange) ErrorPercent(now time.Time) int { _ = "STUB: not implemented"; return 0 }

func (m *metricExchange) IsHealthy(now time.Time) bool { _ = "STUB: not implemented"; return false }
