package hystrix

import (
	"sync"

	"github.com/go-chassis/go-chassis/v2/third_party/forked/afex/hystrix-go/hystrix/rolling"
)

type poolMetrics struct {
	Mutex   *sync.RWMutex
	Updates chan poolMetricsUpdate

	Name              string
	MaxActiveRequests *rolling.Number
	Executed          *rolling.Number
}

type poolMetricsUpdate struct {
	activeCount int
}

func newPoolMetrics(name string) *poolMetrics { _ = "STUB: not implemented"; return nil }

func (m *poolMetrics) Reset() { _ = "STUB: not implemented"; return }

func (m *poolMetrics) Monitor() { _ = "STUB: not implemented"; return }
