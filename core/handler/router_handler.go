package handler

import (
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// RouterHandler router handler
type RouterHandler struct{}

// Handle is to handle the router related things
func (ph *RouterHandler) Handle(chain *Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

//call next chain

func newRouterHandler() Handler {
	_ = "STUB: not implemented"
	return *

	// Name returns the router string
	new(Handler)
}

func (ph *RouterHandler) Name() string { _ = "STUB: not implemented"; return "" }
