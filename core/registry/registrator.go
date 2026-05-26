package registry

import (
	"sync"
	"time"
)

// constant values for registry parameters
const (
	DefaultRegistratorPlugin       = "servicecenter"
	DefaultServiceDiscoveryPlugin  = "servicecenter"
	DefaultContractDiscoveryPlugin = "servicecenter"
	Name                           = "registry"
	SDTag                          = "serviceDiscovery"
	CDTag                          = "contractDiscovery"
	RTag                           = "registrator"
	Auto                           = "auto"
	Manual                         = "manual"
	PersistenceHeartBeat           = "ping-pong"
	NonPersistenceHeartBeat        = "non-keep-alive"
	DefaultInterval                = 30 * time.Second
	MinInterval                    = 10 * time.Second
)

// IsEnabled check enable
var IsEnabled bool
var mu sync.Mutex

// DefaultRegistrator is the client of registry, you can call the method of it to interact with microservice registry
var DefaultRegistrator Registrator

// DefaultAddr default address of service center
var DefaultAddr = "http://127.0.0.1:30100"

// registryFunc registry function
var registryFunc = make(map[string]func(opts Options) Registrator)

// HBService variable of heartbeat service
var HBService = &HeartbeatService{}

// Registrator is the interface for developer to update information in service registry
type Registrator interface {
	//Close destroy connection between the registry client and server
	Close() error
	//RegisterService register a microservice to registry, if it is duplicated in registry, it returns error
	RegisterService(microService *MicroService) (string, error)
	//RegisterServiceInstance register a microservice instance to registry
	RegisterServiceInstance(sid string, instance *MicroServiceInstance) (string, error)
	RegisterServiceAndInstance(microService *MicroService, instance *MicroServiceInstance) (string, string, error)
	Heartbeat(microServiceID, microServiceInstanceID string) (bool, error)
	WSHeartbeat(microServiceID, microServiceInstanceID string, callback func()) (bool, error)
	UnRegisterMicroServiceInstance(microServiceID, microServiceInstanceID string) error
	UpdateMicroServiceInstanceStatus(microServiceID, microServiceInstanceID, status string) error
	UpdateMicroServiceProperties(microServiceID string, properties map[string]string) error
	UpdateMicroServiceInstanceProperties(microServiceID, microServiceInstanceID string, properties map[string]string) error
	AddSchemas(microServiceID, schemaName, schemaInfo string) error
}

func enableRegistrator(opts Options) error { _ = "STUB: not implemented"; return nil }

// InstallRegistrator install registrator plugin
func InstallRegistrator(name string, f func(opts Options) Registrator) {
	_ = "STUB: not implemented"
	return
}

// NewRegistrator return registrator
func NewRegistrator(name string, opts Options) (Registrator, error) {
	_ = "STUB: not implemented"
	return *new(Registrator), nil
}

func getSpecifiedOptions() (oR, oSD, oCD Options, err error) {
	_ = "STUB: not implemented"
	return *new(Options), *new(Options), *new(Options), nil
}

// Enable create DefaultRegistrator
func Enable() (err error) { _ = "STUB: not implemented"; return nil }

// DoRegister for registering micro-service instances
func DoRegister() error { _ = "STUB: not implemented"; return nil }
