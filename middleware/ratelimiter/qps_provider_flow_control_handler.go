package ratelimiter

import (
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// ProviderRateLimiterHandler provider rate limiter handler
type ProviderRateLimiterHandler struct{}

// Handle is to handle provider rateLimiter things
func (rl *ProviderRateLimiterHandler) Handle(chain *handler.Chain, i *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

//qps rate <=0

func newProviderRateLimiterHandler() handler.Handler {
	_ = "STUB: not implemented"
	return *new(handler.Handler)
}

// Name returns the name providerratelimiter
func (rl *ProviderRateLimiterHandler) Name() string { _ = "STUB: not implemented"; return "" }
