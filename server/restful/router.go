package restful

import (
	"github.com/emicklei/go-restful"
	"github.com/go-chassis/go-chassis/v2/core/server"
)

// const for doc
const (
	Path  = "path"
	Query = "query"
)

// Route describe http route path and swagger specifications for API
type Route struct {
	Method           string                 //Method is one of the following: GET,PUT,POST,DELETE. required
	Path             string                 //Path contains a path pattern. required
	ResourceFunc     func(ctx *Context)     //the func this API calls. you must set this field or ResourceFunc, if you set both, ResourceFunc will be used
	ResourceFuncName string                 //the func this API calls. you must set this field or ResourceFunc
	FuncDesc         string                 //tells what this route is all about. Optional.
	Parameters       []*Parameters          //Parameters is a slice of request parameters for a single endpoint Optional.
	Returns          []*Returns             //what kind of response this API returns. Optional.
	Read             interface{}            //Read tells what resource type will be read from the request payload. Optional.
	Consumes         []string               //Consumes specifies that this WebService can consume one or more MIME types.
	Produces         []string               //Produces specifies that this WebService can produce one or more MIME types.
	Metadata         map[string]interface{} //Metadata adds or updates a key=value pair to api
}

// Returns describe response doc
type Returns struct {
	Code    int // http response code
	Message string
	Model   interface{} // response body structure
	Headers map[string]restful.Header
}

// Parameters describe parameters in url path or query params
type Parameters struct {
	Name      string //parameter name
	DataType  string // string, int etc
	ParamType int    //restful.QueryParameterKind or restful.PathParameterKind
	Desc      string
	Required  bool
}

// Router is to define how route the request
type Router interface {
	//URLPatterns returns route
	URLPatterns() []Route
}

// RouteGroup is to define the route group name
type RouteGroup interface {
	//GroupPath if return non-zero-value, it would be appended to route as prefix
	GroupPath() string
}

// GetRouteGroup is to return a router group path
func GetRouteGroup(schema interface{}) string { _ = "STUB: not implemented"; return "" }

// GetRouteSpecs is to return a rest API specification of a go struct
func GetRouteSpecs(schema interface{}) ([]Route, error) { _ = "STUB: not implemented"; return nil, nil }

// WrapHandlerChain wrap business handler with handler chain
func WrapHandlerChain(route *Route, schema interface{}, schemaName string, opts server.Options) (restful.RouteFunction, error) {
	_ = "STUB: not implemented"
	return *new(restful.RouteFunction), nil
}

//create a new chain for each resource handler

//give inv.Ctx to user handlers, modules may inject headers in handler chain

// GroupRoutePath add group route path to route
func GroupRoutePath(route *Route, schema interface{}) { _ = "STUB: not implemented"; return }

// BuildRouteHandler build handler func from ResourceFunc or ResourceFuncName
func BuildRouteHandler(route *Route, schema interface{}) (func(ctx *Context), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getFunctionName get method name from func
func getFunctionName(i interface{}) string { _ = "STUB: not implemented"; return "" }

// replace suffix "-fm" if function is bounded to struct

// GetTrace get trace
func GetTrace() string { _ = "STUB: not implemented"; return "" }
