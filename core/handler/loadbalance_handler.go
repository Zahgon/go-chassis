package handler

import (
	"github.com/go-chassis/go-chassis/v2/control"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/registry"
)

// LBHandler loadbalancer handler struct
type LBHandler struct{}

func (lb *LBHandler) getEndpoint(i *invocation.Invocation, lbConfig control.LoadBalancingConfig) (*registry.Endpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle to handle the load balancing
func (lb *LBHandler) Handle(chain *Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

func (lb *LBHandler) handleWithNoRetry(chain *Chain, i *invocation.Invocation, lbConfig control.LoadBalancingConfig, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

func (lb *LBHandler) handleWithRetry(chain *Chain, i *invocation.Invocation, lbConfig control.LoadBalancingConfig, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

// get retry func

// if get endpoint failed, no need to retry

// if get endpoint failed, no need to retry

// Name returns loadbalancer string
func (lb *LBHandler) Name() string { _ = "STUB: not implemented"; return "" }

func newLBHandler() Handler { _ = "STUB: not implemented"; return *new(Handler) }
