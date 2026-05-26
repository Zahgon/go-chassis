package servicecomb

import (
	"sync"

	"github.com/go-chassis/go-chassis/v2/core/config"
	"github.com/go-chassis/go-chassis/v2/core/router"
)

var cseRouter *Router

// Router is cse router service
type Router struct {
	routeRule map[string][]*config.RouteRule
	lock      sync.RWMutex
}

// SetRouteRule set rules
func (r *Router) SetRouteRule(rr map[string][]*config.RouteRule) { _ = "STUB: not implemented"; return }

// FetchRouteRuleByServiceName get rules for service
func (r *Router) FetchRouteRuleByServiceName(service string) []*config.RouteRule {
	_ = "STUB: not implemented"
	return nil
}

// ListRouteRule get rules for all service
func (r *Router) ListRouteRule() map[string][]*config.RouteRule {
	_ = "STUB: not implemented"
	return nil
}

// Init init router config
func (r *Router) Init(o router.Options) error { _ = "STUB: not implemented"; return nil }

func newRouter() (router.Router, error) { _ = "STUB: not implemented"; return *new(router.Router), nil }

// LoadRules load all the router config
func (r *Router) LoadRules() error { _ = "STUB: not implemented"; return nil }

// SetRouteRuleByKey set route rule by key
func (r *Router) SetRouteRuleByKey(k string, rr []*config.RouteRule) {
	_ = "STUB: not implemented"
	return
}

// DeleteRouteRuleByKey set route rule by key
func (r *Router) DeleteRouteRuleByKey(k string) { _ = "STUB: not implemented"; return }

func init() {
	router.InstallRouterPlugin("cse", newRouter)
}
