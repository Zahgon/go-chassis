package ratelimiter

import (
	"github.com/go-chassis/openlog"

	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// names
const (
	Consumer = "ratelimiter-consumer"
	Provider = "ratelimiter-provider"
	Name     = "rate-limiter"
)

// ConsumerRateLimiterHandler consumer rate limiter handler
type ConsumerRateLimiterHandler struct{}

// Handle is handles the consumer rate limiter APIs
func (rl *ConsumerRateLimiterHandler) Handle(chain *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

//qps rate <=0

func newErrResponse(i *invocation.Invocation) *invocation.Response {
	_ = "STUB: not implemented"
	return nil
}

func newConsumerRateLimiterHandler() handler.Handler {
	_ = "STUB: not implemented"
	return *new(handler.Handler)
}

// Name returns name
func (rl *ConsumerRateLimiterHandler) Name() string { _ = "STUB: not implemented"; return "" }

func init() {
	err := handler.RegisterHandler(Consumer, newConsumerRateLimiterHandler)
	if err != nil {
		openlog.Error(err.Error())
	}
	err = handler.RegisterHandler(Provider, newProviderRateLimiterHandler)
	if err != nil {
		openlog.Error(err.Error())
	}
}
