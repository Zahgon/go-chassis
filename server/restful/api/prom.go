// Package api supply a bunch of common http handler that could be mounted on you rest server and expose ability
// for example: prometheus metrics will be exported on /metrics api
package api

import (
	"github.com/emicklei/go-restful"
)

// PrometheusHandleFunc is a go-restful handler which can expose metrics in http server
func PrometheusHandleFunc(req *restful.Request, rep *restful.Response) {
	_ = "STUB: not implemented"
	return
}
