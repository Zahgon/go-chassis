package circuit

import (
	"errors"

	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

const (
	//ReturnNil is build in fallback name
	ReturnNil = "returnnull"
	//ReturnErr is build in fallback name
	ReturnErr = "throwexception"
)

var fallbackFuncMap = make(map[string]Fallback)

// ErrFallbackNotExists happens if fallback implementation does not exist
var ErrFallbackNotExists = errors.New("fallback func does not exist")

// Fallback defines how to return response if remote call fails.
// a implementation should return a closure to handle the error.
// in this closure, if you fallback logic should handle the original error,
// you can return a fallback error to replace the original error
// you can assemble invocation.Response on demand
// in summary the closure defines, "if err happens, how to handle it".
type Fallback func(inv *invocation.Invocation, finish chan *invocation.Response) func(error) error

// Init init functions
func Init() { _ = "STUB: not implemented"; return }

// RegisterFallback register custom logic
func RegisterFallback(name string, f Fallback) { _ = "STUB: not implemented"; return }

// GetFallback return function
func GetFallback(name string) (Fallback, error) {
	_ = "STUB: not implemented"
	return *new(Fallback), nil
}

// FallbackNil return empty response
func FallbackNil(inv *invocation.Invocation, finish chan *invocation.Response) func(error) error {
	_ = "STUB: not implemented"
	return nil
}

// if err is type of hystrix error, return a new response

// isolation happened, so lead to callback

//make sure body is empty

//no need to return error

// call back success

// FallbackErr set err in response
func FallbackErr(inv *invocation.Invocation, finish chan *invocation.Response) func(error) error {
	_ = "STUB: not implemented"
	return nil
}

// if err is type of hystrix error, return a new response

// isolation happened, so lead to callback

// isolation happened, so lead to callback

//do nothing, just give original error

//make sure body is empty

//no need to return error
