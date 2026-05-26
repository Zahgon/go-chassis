package restful

import (
	"context"
	"net/http"

	"github.com/emicklei/go-restful"
)

// Context is a struct which has both request and response objects
// and request context
type Context struct {
	Ctx  context.Context
	Req  *restful.Request
	Resp *restful.Response
}

// NewBaseServer is a function which return context
func NewBaseServer(ctx context.Context) *Context { _ = "STUB: not implemented"; return nil }

// write is the response writer.
func (bs *Context) Write(body []byte) error { _ = "STUB: not implemented"; return nil }

// WriteHeader is the response head writer
func (bs *Context) WriteHeader(httpStatus int) { _ = "STUB: not implemented"; return }

// AddHeader is a function used to add header to a response
func (bs *Context) AddHeader(header string, value string) { _ = "STUB: not implemented"; return }

// WriteError is a function used to write error into a response
func (bs *Context) WriteError(httpStatus int, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteJSON used to write a JSON file into response
func (bs *Context) WriteJSON(value interface{}, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteHeaderAndJSON used to write head and JSON file in to response
func (bs *Context) WriteHeaderAndJSON(status int, value interface{}, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadEntity is request reader
func (bs *Context) ReadEntity(schema interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ReadHeader is used to read header of request
func (bs *Context) ReadHeader(name string) string { _ = "STUB: not implemented"; return "" }

// ReadPathParameter is used to read path parameter of a request
func (bs *Context) ReadPathParameter(name string) string { _ = "STUB: not implemented"; return "" }

// ReadPathParameters used to read multiple path parameters of a request
func (bs *Context) ReadPathParameters() map[string]string { _ = "STUB: not implemented"; return nil }

// ReadQueryParameter is used to read query parameter of a request
func (bs *Context) ReadQueryParameter(name string) string { _ = "STUB: not implemented"; return "" }

// ReadQueryEntity is used to read query parameters into a specified struct.
// The struct tag should be `form` like:
//
//	type QueryRequest struct {
//	    Name string `form:"name"`
//	    Password string `form:"password"`
//	}
func (bs *Context) ReadQueryEntity(schema interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// ReadBodyParameter used to read body parameter of a request
func (bs *Context) ReadBodyParameter(name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ReadRequest return a native net/http request
func (bs *Context) ReadRequest() *http.Request { _ = "STUB: not implemented"; return nil }

// ReadRestfulRequest return a native  go-restful request
func (bs *Context) ReadRestfulRequest() *restful.Request {
	_ = "STUB: not implemented"

	// ReadResponseWriter return a native net/http ResponseWriter
	return nil
}

func (bs *Context) ReadResponseWriter() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

// ReadRestfulResponse return a native go-restful Response
func (bs *Context) ReadRestfulResponse() *restful.Response { _ = "STUB: not implemented"; return nil }
