package restful

import (
	"net/http"
	"sync"

	"github.com/emicklei/go-restful"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
	"github.com/go-chassis/go-chassis/v2/core/server"
)

// constants for metric path and name
const (
	//Name is a variable of type string which indicates the protocol being used
	Name = "rest"

	ProfileRouteRuleSubPath = "route-rule"
	ProfileDiscoverySubPath = "discovery"
	MimeFile                = "application/octet-stream"
	MimeMult                = "multipart/form-data"
)

const openTLS = "?sslEnabled=true"

func init() {
	server.InstallPlugin(Name, newRestfulServer)
}

type restfulServer struct {
	container *restful.Container
	ws        *restful.WebService
	opts      server.Options
	mux       sync.RWMutex
	server    *http.Server
}

func newRestfulServer(opts server.Options) server.ProtocolServer {
	_ = "STUB: not implemented"
	return *new(server.ProtocolServer)
}

// Add container filter to respond to OPTIONS

func addProfileRoutes(ws *restful.WebService, opts server.Options) {
	_ = "STUB: not implemented"
	return
}

// HTTPRequest2Invocation convert http request to uniform invocation data format
func HTTPRequest2Invocation(req *restful.Request, schema, operation string, resp *restful.Response) (*invocation.Invocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//set headers to Ctx, then user do not  need to consider about protocol in handlers

func (r *restfulServer) Register(schema interface{}, options ...server.RegisterOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Invocation2HTTPRequest convert invocation back to http request, set down all meta data
func Invocation2HTTPRequest(inv *invocation.Invocation, req *restful.Request) {
	_ = "STUB: not implemented"
	return
}

// Register2GoRestful register http handler to go-restful framework
func Register2GoRestful(routeSpec Route, ws *restful.WebService, handler restful.RouteFunction) error {
	_ = "STUB: not implemented"
	return nil
}

// fillParam is for handle parameter by type
func fillParam(routeSpec Route, rb *restful.RouteBuilder) *restful.RouteBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (r *restfulServer) Start() error { _ = "STUB: not implemented"; return nil }

// create schema

// register to swagger ui,Whether to create a schema, you need to refer to the configuration.
func (r *restfulServer) CreateLocalSchema(opts server.Options) error {
	_ = "STUB: not implemented"
	return nil
}

//set schema information when create local schema file

func (r *restfulServer) Stop() error { _ = "STUB: not implemented"; return nil }

//only golang 1.8 support graceful shutdown.

// failure/timeout shutting down the server gracefully

func (r *restfulServer) String() string { _ = "STUB: not implemented"; return "" }
