package handler

import (
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// ChainMap just concurrent read
var ChainMap = make(map[string]*Chain)

// Chain struct for service and handlers
type Chain struct {
	ServiceType string
	Name        string
	Handlers    []Handler
}

func (c *Chain) Clone() Chain { _ = "STUB: not implemented"; return *new(Chain) }

// AddHandler chain can add a handler
func (c *Chain) AddHandler(h Handler) { _ = "STUB: not implemented"; return }

// Next is for to handle next handler in the chain
func (c *Chain) Next(i *invocation.Invocation, f invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

// ChainOptions chain options
type ChainOptions struct {
	Name string
}

// ChainOption is a function name
type ChainOption func(*ChainOptions)

// WithChainName returns the name of the chain option
func WithChainName(name string) ChainOption { _ = "STUB: not implemented"; return *new(ChainOption) }

// parseHandlers for parsing the handlers
func parseHandlers(handlerStr string) []string { _ = "STUB: not implemented"; return nil }

//delete empty string

// CreateChains create the chains based on type and handler map
func CreateChains(chainType string, handlerNameMap map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateChain create consumer or provider's chain,the handlers is different
func CreateChain(serviceType string, chainName string, handlerNames ...string) (*Chain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// addHandler add handler
func addHandler(c *Chain, name string) error { _ = "STUB: not implemented"; return nil }

// GetChain is to get chain
func GetChain(serviceType string, name string) (*Chain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
