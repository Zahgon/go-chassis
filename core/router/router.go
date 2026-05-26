// Package router expose API for user to get or set route rule
package router

import (
	"errors"

	"github.com/go-chassis/go-chassis/v2/core/config"

	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/registry"
)

// Router return route rule, you can also set custom route rule
type Router interface {
	Init(Options) error
	SetRouteRule(map[string][]*config.RouteRule)
	FetchRouteRuleByServiceName(service string) []*config.RouteRule
	ListRouteRule() map[string][]*config.RouteRule
}

// ErrNoExist means if there is no router implementation
var ErrNoExist = errors.New("router not exists")
var routerServices = make(map[string]func() (Router, error))

// DefaultRouter is current router implementation
var DefaultRouter Router

// InstallRouterPlugin install router plugin
func InstallRouterPlugin(name string, f func() (Router, error)) { _ = "STUB: not implemented"; return }

// BuildRouter create a router
func BuildRouter(name string) error { _ = "STUB: not implemented"; return nil }

// Route decide the target service metadata
// it decide based on configuration of route rule
// it will set RouteTag to invocation
func Route(header map[string]string, si *registry.SourceInfo, inv *invocation.Invocation) error {
	_ = "STUB: not implemented"
	return nil
}

// FitRate fit rate
func FitRate(tags []*config.RouteTag, dest string) *config.RouteTag {
	_ = "STUB: not implemented"
	return nil
}

// match check the route rule
func Match(inv *invocation.Invocation, matchConf config.Match, headers map[string]string, source *registry.SourceInfo) bool {
	_ = "STUB: not implemented"
	//validate template first
	return false
}

//matchConf rule is not set

// SourceMatch check the source route
func SourceMatch(match *config.Match, headers map[string]string, source *registry.SourceInfo) bool {
	_ = "STUB: not implemented"
	//source not match
	return false
}

//source tags not match

//source headers not match

// isMatch check the route rule
func isMatch(headers map[string]string, k string, v map[string]string) bool {
	_ = "STUB: not implemented"
	return false
}

func valueToUpper(b, value string) string { _ = "STUB: not implemented"; return "" }

// SortRules sort route rules
func SortRules(name string) []*config.RouteRule { _ = "STUB: not implemented"; return nil }

// QuickSort for sorting the routes it will follow quicksort technique
func QuickSort(left int, right int, rules []*config.RouteRule) (s []*config.RouteRule) {
	_ = "STUB: not implemented"
	return nil
}

//move base to the current position of i&j
