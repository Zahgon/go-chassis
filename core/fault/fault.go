package fault

import (
	"github.com/go-chassis/go-chassis/v2/core/config/model"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// InjectFault inject fault
type InjectFault func(model.Fault, *invocation.Invocation) error

// Injectors fault injectors
var Injectors = make(map[string]InjectFault)

// Fault fault injection error
type Fault struct {
	Message string
}

func (e Fault) Error() string {
	_ = "STUB: not implemented"

	// InstallFaultInjectionPlugin install fault injection plugin
	return ""
}

func InstallFaultInjectionPlugin(name string, f InjectFault) { _ = "STUB: not implemented"; return }

func init() {
	InstallFaultInjectionPlugin("rest", faultInject)
	InstallFaultInjectionPlugin("dubbo", faultInject)
}

func faultInject(rule model.Fault, inv *invocation.Invocation) error {
	_ = "STUB: not implemented"
	return nil
}
