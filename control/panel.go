package control

import (
	"github.com/go-chassis/go-chassis/v2/core/config/model"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/third_party/forked/afex/hystrix-go/hystrix"
)

var panelPlugin = make(map[string]func(options Options) Panel)

// DefaultPanel get fetch config
var DefaultPanel Panel

const (
	//ScopeAPI is config const
	ScopeAPI = "api"

	//ScopeInstance is config const
	ScopeInstance = "instance"

	//ScopeInstanceAPI is config const
	ScopeInstanceAPI = "instance-api"
)

var (
	DefaultBurst = 10
)

// Panel is a abstraction of pulling configurations from various of systems, and transfer different configuration into standardized model
// you can use different panel implementation to pull different of configs from Istio or Archaius
// TODO able to set configs
type Panel interface {
	GetCircuitBreaker(inv invocation.Invocation, serviceType string) (string, hystrix.CommandConfig)
	GetLoadBalancing(inv invocation.Invocation) LoadBalancingConfig
	GetRateLimiting(inv invocation.Invocation, serviceType string) RateLimitingConfig
	GetFaultInjection(inv invocation.Invocation) model.Fault
	GetEgressRule() []EgressConfig
}

// InstallPlugin install implementation
func InstallPlugin(name string, f func(options Options) Panel) { _ = "STUB: not implemented"; return }

// Init initialize DefaultPanel
func Init(opts Options) error { _ = "STUB: not implemented"; return nil }

// NewCircuitName create circuit command string
// scope means has two choices, service and api
// if you set it to api, a api level command string will be created. like "Consumer.mall.rest./test"
// set to service, a service level command will be created, like "Consumer.mall"
func NewCircuitName(serviceType, scope string, inv invocation.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}
