package retry

import (
	"github.com/cenkalti/backoff"
)

// retry kind
const (
	KindExponential    = "exponential"
	KindConstant       = "constant"
	KindZero           = "zero"
	DefaultBackOffKind = KindExponential
)

// GetBackOff return the the back off policy
// min and max unit is million second
func GetBackOff(kind string, min, max int) backoff.BackOff {
	_ = "STUB: not implemented"
	return *new(backoff.BackOff)
}
