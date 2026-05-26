package handler

import (
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// TracingProviderHandler tracing provider handler
type TracingProviderHandler struct{}

// Handle is to handle the provider tracing related things
func (t *TracingProviderHandler) Handle(chain *Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

// extract span context
// header stored in context

// store span in context

// To ensure accuracy, spans should finish immediately once server responds.
// So the best way is that spans finish in the callback func, not after it.
// But server may respond in the callback func too, that we have to remove
// span finishing from callback func's inside to outside.

// Name returns tracing-provider string
func (t *TracingProviderHandler) Name() string { _ = "STUB: not implemented"; return "" }

func newTracingProviderHandler() Handler { _ = "STUB: not implemented"; return *new(Handler) }

// TracingConsumerHandler tracing consumer handler
type TracingConsumerHandler struct{}

// Handle is handle consumer tracing related things
func (t *TracingConsumerHandler) Handle(chain *Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	// the span context is in invocation.Ctx
	// start a new span from context
	return
}

// store span in context

// set span kind to be client

// store span in context

// inject span context into carrier

// header stored in context

// To ensure accuracy, spans should finish immediately once client send req.
// So the best way is that spans finish in the callback func, not after it.
// But client may send req in the callback func too, that we have to remove
// span finishing from callback func's inside to outside.

// Name returns tracing-consumer string
func (t *TracingConsumerHandler) Name() string { _ = "STUB: not implemented"; return "" }

func newTracingConsumerHandler() Handler { _ = "STUB: not implemented"; return *new(Handler) }
