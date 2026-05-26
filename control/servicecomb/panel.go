package servicecomb

import (
	"github.com/go-chassis/go-chassis/v2/control"
	"github.com/go-chassis/go-chassis/v2/core/config/model"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/third_party/forked/afex/hystrix-go/hystrix"
)

// Panel pull configs from archaius
type Panel struct {
}

func newPanel(options control.Options) control.Panel {
	_ = "STUB: not implemented"
	return *new(control.Panel)
}

// GetCircuitBreaker return command , and circuit breaker settings
func (p *Panel) GetCircuitBreaker(inv invocation.Invocation, serviceType string) (string, hystrix.CommandConfig) {
	_ = "STUB: not implemented"
	return "", *new(hystrix.CommandConfig)
}

// GetLoadBalancing get load balancing config
func (p *Panel) GetLoadBalancing(inv invocation.Invocation) control.LoadBalancingConfig {
	_ = "STUB: not implemented"
	return *new(control.LoadBalancingConfig)
}

// GetRateLimiting get rate limiting config
func (p *Panel) GetRateLimiting(inv invocation.Invocation, serviceType string) control.RateLimitingConfig {
	_ = "STUB: not implemented"
	return *new(control.RateLimitingConfig)
}

// GetFaultInjection get Fault injection config
func (p *Panel) GetFaultInjection(inv invocation.Invocation) model.Fault {
	_ = "STUB: not implemented"
	return *

	// GetEgressRule get egress config
	new(model.Fault)
}

func (p *Panel) GetEgressRule() []control.EgressConfig { _ = "STUB: not implemented"; return nil }

func init() {
	control.InstallPlugin("archaius", newPanel)
}

// GetQPSRateWithPriority get qps rate with priority
func GetQPSRateWithPriority(cmd ...string) (int, string) { _ = "STUB: not implemented"; return 0, "" }

// GetQPSRate get qps rate
func GetQPSRate(rateConfig string) (int, bool) { _ = "STUB: not implemented"; return 0, false }
