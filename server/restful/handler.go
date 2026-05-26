package restful

import (
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/server"
)

// ResourceHandler wraps go-chassis restful function
type ResourceHandler struct {
	handleFunc func(ctx *Context)
	rc         *Context
	opts       server.Options
}

// Handle is to handle the router related things
func (h *ResourceHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

// check body size

// call real route func

//call next chain

func newHandler(f func(ctx *Context), rc *Context, opts server.Options) handler.Handler {
	_ = "STUB: not implemented"
	return *new(handler.Handler)
}

// Name returns the name string
func (h *ResourceHandler) Name() string { _ = "STUB: not implemented"; return "" }
