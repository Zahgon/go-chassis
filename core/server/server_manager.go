package server

import (
	"github.com/go-chassis/go-chassis/v2/core/config/model"
)

// constants for server
const (
	DefaultMetricPath  = "metrics"
	DefaultProfilePath = "profile"
)

// NewFunc returns a ProtocolServer
type NewFunc func(Options) ProtocolServer

var serverPlugins = make(map[string]NewFunc)
var servers = make(map[string]ProtocolServer)

// InstallPlugin For developer
func InstallPlugin(protocol string, newFunc NewFunc) { _ = "STUB: not implemented"; return }

// GetServerFunc returns the server function
func GetServerFunc(protocol string) (NewFunc, error) {
	_ = "STUB: not implemented"
	return *new(NewFunc), nil
}

// GetServer return the server based on protocol
func GetServer(protocol string) (ProtocolServer, error) {
	_ = "STUB: not implemented"
	return *new(ProtocolServer), nil
}

// GetServers returns the map of servers
func GetServers() map[string]ProtocolServer {
	_ = "STUB: not implemented"

	// ErrRuntime is an error channel, if it receive any signal will cause graceful shutdown of go chassis, process will exit
	return nil
}

var ErrRuntime = make(chan error)

// StartServer starting the server
func StartServer(options ...RunOption) error { _ = "STUB: not implemented"; return nil }

// UnRegistrySelfInstances this function removes the self instance
func UnRegistrySelfInstances() error { _ = "STUB: not implemented"; return nil }

// Init initializes
func Init() error { _ = "STUB: not implemented"; return nil }

func initialServer(providerMap map[string]string, p model.Protocol, name string) error {
	_ = "STUB: not implemented"
	return nil
}
