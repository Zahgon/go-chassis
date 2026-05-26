package core

import (
	"context"

	"net/http"
)

// schemas
const (
	HTTP  = "http"
	HTTPS = "https"
)

// RestInvoker is rest invoker
// one invoker for one microservice
// thread safe
type RestInvoker struct {
	*abstractInvoker
}

// NewRestInvoker is gives the object of rest invoker
func NewRestInvoker(opt ...Option) *RestInvoker { _ = "STUB: not implemented"; return nil }

// ContextDo is for requesting the API
// by default if http status is 5XX, then it will return error
func (ri *RestInvoker) ContextDo(ctx context.Context, req *http.Request, options ...InvocationOption) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set headers to Ctx

//TODO load from openAPI schema
