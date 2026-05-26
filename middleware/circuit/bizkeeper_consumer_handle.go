package circuit

import (
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/third_party/forked/afex/hystrix-go/hystrix"
	"github.com/go-chassis/openlog"
)

// constant for bizkeeper-consumer
const (
	Name = "bizkeeper-consumer"
)

// BizKeeperConsumerHandler bizkeeper consumer handler
type BizKeeperConsumerHandler struct{}

// Handle function is for to handle the chain
func (bk *BizKeeperConsumerHandler) Handle(chain *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

// means hystrix error occurred

// err is not nil in conditions:
// 1 fallback is nil
//   1.1 chain.Next() fail
//   1.2 hystrix mechanism, retur error as ErrMaxConcurrency / ErrCircuitOpen / ErrForceFallback
// 2 fallback is not nil
//   2.1 fallback failed no matter chain.Next() is executed or not

// GetFallbackFun get fallback function
func GetFallbackFun(cmd, t string, i *invocation.Invocation, finish chan *invocation.Response, isForce bool) (func(error) error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newBizKeeperConsumerHandler new bizkeeper consumer handler
func newBizKeeperConsumerHandler() handler.Handler {
	_ = "STUB: not implemented"
	return *new(handler.Handler)
}

// Name is for to represent the name of bizkeeper handler
func (bk *BizKeeperConsumerHandler) Name() string { _ = "STUB: not implemented"; return "" }

func init() {
	err := handler.RegisterHandler(Name, newBizKeeperConsumerHandler)
	if err != nil {
		openlog.Error(err.Error())
	}
	err = handler.RegisterHandler("bizkeeper-provider", newBizKeeperProviderHandler)
	if err != nil {
		openlog.Error(err.Error())
	}
	Init()
	go hystrix.StartReporter()
}
