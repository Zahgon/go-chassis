package httputil

import (
	"errors"
	"net/http"

	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// ErrInvalidReq invalid input
var ErrInvalidReq = errors.New("rest consumer call arg is not *http.Request type")

// SetURI sets host for the request.
// set http(s)://{domain}/xxx
func SetURI(req *http.Request, url string) { _ = "STUB: not implemented"; return }

// SetBody is a method used for setting body for a request
func SetBody(req *http.Request, body []byte) { _ = "STUB: not implemented"; return }

// SetCookie set key value in request cookie
func SetCookie(req *http.Request, k, v string) { _ = "STUB: not implemented"; return }

// GetCookie is a method which gets cookie from a request
func GetCookie(req *http.Request, key string) string { _ = "STUB: not implemented"; return "" }

// SetContentType is a method used for setting content-type in a request
func SetContentType(req *http.Request, ct string) { _ = "STUB: not implemented"; return }

// GetContentType is a method used for getting content-type in a request
func GetContentType(req *http.Request) string { _ = "STUB: not implemented"; return "" }

// HTTPRequest convert invocation to http request
func HTTPRequest(inv *invocation.Invocation) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadBody read body from the from the response
func ReadBody(resp *http.Response) []byte { _ = "STUB: not implemented"; return nil }

// GetRespCookie returns response Cookie.
func GetRespCookie(resp *http.Response, key string) []byte { _ = "STUB: not implemented"; return nil }

// SetRespCookie sets the cookie.
func SetRespCookie(resp *http.Response, cookie *http.Cookie) { _ = "STUB: not implemented"; return }
