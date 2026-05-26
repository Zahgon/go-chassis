package hystrix

type executorPool struct {
	Name    string
	Metrics *poolMetrics
	Max     int
	Tickets chan *struct{}
}

const ConcurrentRequestsLimit = 5000

func newExecutorPool(name string) *executorPool { _ = "STUB: not implemented"; return nil }

func (p *executorPool) Return(ticket *struct{}) { _ = "STUB: not implemented"; return }

func (p *executorPool) ActiveCount() int { _ = "STUB: not implemented"; return 0 }
