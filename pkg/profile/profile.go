package profile

import (
	"github.com/emicklei/go-restful"
	"github.com/go-chassis/go-chassis/v2/core/config"
	"github.com/go-chassis/go-chassis/v2/core/registry"
)

// const
const (
	msgWriteError = "write to response err: "
)

// Profile contains route rule and discovery
type Profile struct {
	RouteRule map[string][]*config.RouteRule              `json:"routeRule"`
	Discovery map[string][]*registry.MicroServiceInstance `json:"discovery"`
}

// HTTPHandleRouteRuleFunc is a go-restful handler which can expose profile of route rule in http server
func HTTPHandleRouteRuleFunc(req *restful.Request, rep *restful.Response) {
	_ = "STUB: not implemented"
	return
}

// HTTPHandleDiscoveryFunc is a go-restful handler which can expose profile of discovery in http server
func HTTPHandleDiscoveryFunc(req *restful.Request, rep *restful.Response) {
	_ = "STUB: not implemented"
	return
}

// HTTPHandleProfileFunc is a go-restful handler which can expose all profiles in http server
func HTTPHandleProfileFunc(req *restful.Request, rep *restful.Response) {
	_ = "STUB: not implemented"
	return
}

func newProfile() Profile { _ = "STUB: not implemented"; return *new(Profile) }

func listRouteRule() map[string][]*config.RouteRule { _ = "STUB: not implemented"; return nil }

func listMicroServiceInstance() map[string][]*registry.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}
