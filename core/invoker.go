package core

import (
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// newOptions is for updating options
func newOptions(options ...Option) Options { _ = "STUB: not implemented"; return *new(Options) }

// abstract invoker is a common invoker for RPC
type abstractInvoker struct {
	opts Options
}

func (ri *abstractInvoker) invoke(i *invocation.Invocation) error {
	_ = "STUB: not implemented"
	return nil
}

// add self service name into remote context, this value used in provider rate limiter

// setCookieToCache   set go-chassisLB cookie to cache when use SessionStickiness strategy
func setCookieToCache(inv invocation.Invocation, namespace string) {
	_ = "STUB: not implemented"
	return
}

// getNamespaceFromMetadata get namespace from opts.Metadata
func getNamespaceFromMetadata(metadata map[string]interface{}) string {
	_ = "STUB: not implemented"
	return ""
}
