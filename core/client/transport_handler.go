package client

import (
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/openlog"
)

// TransportHandler transport handler
type TransportHandler struct{}

// Name returns transport string
func (th *TransportHandler) Name() string { _ = "STUB: not implemented"; return "" }

func errNotNil(err error, cb invocation.ResponseCallBack) { _ = "STUB: not implemented"; return }

// Handle is to handle transport related things
func (th *TransportHandler) Handle(chain *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

//taking the time elapsed to check for latency aware strategy

// ProcessSpecialProtocol handles special logic for protocol
func ProcessSpecialProtocol(inv *invocation.Invocation) { _ = "STUB: not implemented"; return }

// ProcessSuccessiveFailure handles special logic for protocol
func ProcessSuccessiveFailure(i *invocation.Invocation) { _ = "STUB: not implemented"; return }

func newTransportHandler() handler.Handler { _ = "STUB: not implemented"; return *new(handler.Handler) }

func init() {
	err := handler.RegisterHandler(handler.Transport, newTransportHandler)
	if err != nil {
		openlog.Fatal("can not init chassis" + err.Error())
	}
}
