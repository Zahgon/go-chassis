package handler

import (
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// constants of bizkeeper fake and loadbalancer fake
const (
	//BIZKEEPERFAKE & LOADBALANCEFAKE are variables of type string
	BIZKEEPERFAKE   = "bizkeeper-fake"
	LOADBALANCEFAKE = "loadbalancer-fake"
)

// BizkeeperFakeHandler fake handler for bizkeeper
type BizkeeperFakeHandler struct{}

// Name 方法 实现
func (bizkeeperfhandler *BizkeeperFakeHandler) Name() string { _ = "STUB: not implemented"; return "" }

// Handle 方法 实现
func (bizkeeperfhandler *BizkeeperFakeHandler) Handle(c *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	// 调用Chain.Next(i *invocation.Invocation, f invocation.ResponseCallBack)方法，
	// 执行Chain的下一个handler
	return
}

// func() Handler方法 实现
func createBizkeeperFakeHandler() handler.Handler {
	_ = "STUB: not implemented"
	return *new(handler.Handler)
}

// LoadbalanceHandlerFake fake handler for loadbalancer
type LoadbalanceHandlerFake struct{}

// Name 方法 实现
func (lbfakehandler *LoadbalanceHandlerFake) Name() string { _ = "STUB: not implemented"; return "" }

// Handle 方法 实现
func (lbfakehandler *LoadbalanceHandlerFake) Handle(c *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	// 调用Chain.Next(i *invocation.Invocation, f invocation.ResponseCallBack)方法，
	// 执行Chain的下一个handler
	return
}

// func() Handler方法 实现
func createLoadbalanceHandler() handler.Handler {
	_ = "STUB: not implemented"
	return *new(handler.Handler)
}

func init() {
	// 注册handler name和对应的func() Handler方法
	handler.RegisterHandler(BIZKEEPERFAKE, createBizkeeperFakeHandler)
	handler.RegisterHandler(LOADBALANCEFAKE, createLoadbalanceHandler)
}
