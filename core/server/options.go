package server

import (
	"crypto/tls"
	"time"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/go-chassis/go-chassis/v2/core/provider"
)

// Options is the options for service initiating
type Options struct {
	Address            string
	ProtocolServerName string
	ChainName          string
	Provider           provider.Provider
	TLSConfig          *tls.Config
	BodyLimit          int64
	HeaderLimit        int
	Timeout            time.Duration

	ProfilingEnable bool
	ProfilingAPI    string

	MetricsEnable bool
	MetricsAPI    string
}

// RegisterOptions is options when you register a schema to chassis
type RegisterOptions struct {
	SchemaID   string
	Method     string
	Path       string
	RPCSvcDesc interface{}
}

// RegisterOption is option when you register a schema to chassis
type RegisterOption func(*RegisterOptions)

// WithSchemaID you can specify a unique id for schema
func WithSchemaID(schemaID string) RegisterOption {
	_ = "STUB: not implemented"
	return *new(RegisterOption)
}

// WithPath specify a url pattern
func WithPath(Path string) RegisterOption { _ = "STUB: not implemented"; return *new(RegisterOption) }

// WithMethod specify a method
func WithMethod(Method string) RegisterOption {
	_ = "STUB: not implemented"
	return *new(RegisterOption)
}

// WithRPCServiceDesc you can set rpc service desc, it cloud be *grpc.ServiceDesc
func WithRPCServiceDesc(RPCSvcDesc interface{}) RegisterOption {
	_ = "STUB: not implemented"
	return *new(RegisterOption)
}

type RunOptions struct {
	serverMasks sets.String
}

type RunOption func(*RunOptions)

// WithServerMask you can specify do not start a protocol server
func WithServerMask(serverNames ...string) RunOption {
	_ = "STUB: not implemented"
	return *new(RunOption)
}
