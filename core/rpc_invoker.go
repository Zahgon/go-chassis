package core

import (
	"context"
	"sync"
)

// RPCInvoker is rpc invoker
// one invoker for one microservice
// thread safe
type RPCInvoker struct {
	*abstractInvoker
	sync.RWMutex
}

// NewRPCInvoker is gives the object of rpc invoker
func NewRPCInvoker(opt ...Option) *RPCInvoker { _ = "STUB: not implemented"; return nil }

// Invoke is for to invoke the functions during API calls
func (ri *RPCInvoker) Invoke(ctx context.Context, microServiceName, schemaID, operationID string, arg interface{}, reply interface{}, options ...InvocationOption) error {
	_ = "STUB: not implemented"
	return nil
}
