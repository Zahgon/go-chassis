package tracing

import (
	"github.com/opentracing/opentracing-go"
)

// TracerFuncMap saves NewTracer func
// key: impl name
// val: tracer new func
var TracerFuncMap = make(map[string]NewTracer)

// NewTracer is the func to return global tracer
type NewTracer func(o map[string]string) (opentracing.Tracer, error)

// InstallTracer install new opentracing tracer
func InstallTracer(name string, f NewTracer) { _ = "STUB: not implemented"; return }

// GetTracerFunc get NewTracer
func GetTracerFunc(name string) (NewTracer, error) {
	_ = "STUB: not implemented"
	return *new(NewTracer), nil
}

// Init initialize the global tracer
func Init() error { _ = "STUB: not implemented"; return nil }
