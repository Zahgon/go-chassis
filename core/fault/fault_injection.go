package fault

import (
	"github.com/go-chassis/go-chassis/v2/core/config/model"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// constant for default values values of abort and delay percentages
const (
	DefaultAbortPercentage int = 100
	DefaultDelayPercentage int = 100

	MaxPercentage int = 100
	MinPercentage int = 0
)

// ValidateAndApplyFault validate and apply the fault rule
func ValidateAndApplyFault(fault *model.Fault, inv *invocation.Invocation) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateFaultAbort checks that fault injection abort HTTP status and Percentage is valid
func ValidateFaultAbort(fault *model.Fault) error { _ = "STUB: not implemented"; return nil }

// ValidateFaultDelay checks that fault injection delay fixed delay and Percentage is valid
func ValidateFaultDelay(fault *model.Fault) error { _ = "STUB: not implemented"; return nil }

// ApplyFaultInjection abort/delay
func ApplyFaultInjection(fault *model.Fault, inv *invocation.Invocation, configuredPercent int, faultType string) error {
	_ = "STUB: not implemented"
	return nil
}

// injectFault apply fault based on the type
func injectFault(faultType string, fault *model.Fault) error { _ = "STUB: not implemented"; return nil }
