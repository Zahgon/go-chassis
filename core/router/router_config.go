package router

import (
	"crypto/tls"

	"github.com/go-chassis/go-chassis/v2/core/config"
)

// RouterTLS defines tls prefix
const RouterTLS = "router"

// Init initialize router config in local file
// then is create the router component
func Init() error { _ = "STUB: not implemented"; return nil }

// ValidateRule validate the route rules of each service
func ValidateRule(rules map[string][]*config.RouteRule) bool {
	_ = "STUB: not implemented"
	return false
}

// Options defines how to init router and its fetcher
type Options struct {
	Endpoints []string
	EnableSSL bool
	TLSConfig *tls.Config
	Version   string

	//TODO: need timeout for client
	// TimeOut time.Duration
}

func getSpecifiedOptions() (opts Options, err error) {
	_ = "STUB: not implemented"
	return *new(Options), nil
}

// TODO: envoy api v1 or v2
// opts.Version = config.GetRouterAPIVersion()

// routeTagToTags returns tags from a route tag
func routeTagToTags(t *config.RouteTag) utiltags.Tags {
	_ = "STUB: not implemented"
	return *new(utiltags.Tags)
}

// GenWeightPoolKey returns weight pool cache key
func GenWeightPoolKey(dest string, precedence int) string { _ = "STUB: not implemented"; return "" }
