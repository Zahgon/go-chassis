package handler

import (
	"github.com/go-chassis/go-chassis/v2/core/config/model"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// constant for fault handler name
const (
	FaultHandlerName = "fault-inject"
)

// FaultHandler handler
type FaultHandler struct{}

// newFaultHandler fault handle gives the object of FaultHandler
func newFaultHandler() Handler {
	_ = "STUB: not implemented"
	return *

	// Name function returns fault-inject string
	new(Handler)
}

func (rl *FaultHandler) Name() string { _ = "STUB: not implemented"; return "" }

// Handle is to handle the API
func (rl *FaultHandler) Handle(chain *Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

// GetFaultConfig get faultconfig
func GetFaultConfig(protocol, microServiceName, schemaID, operationID string) model.Fault {
	_ = "STUB: not implemented"
	return *new(model.Fault)
}
