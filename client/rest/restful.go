package rest

import (
	"net/http"
)

// NewRequest is a function which creates new request
func NewRequest(method, urlStr string, body []byte) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewResponse is creating the object of response
func NewResponse() *http.Response { _ = "STUB: not implemented"; return nil }
