package registry

var sdFunc = make(map[string]func(opts Options) ServiceDiscovery)

var cdFunc = make(map[string]func(opts Options) ContractDiscovery)

// InstallServiceDiscovery install service discovery client
func InstallServiceDiscovery(name string, f func(opts Options) ServiceDiscovery) {
	_ = "STUB: not implemented"
	return
}

// NewDiscovery create discovery service
func NewDiscovery(name string, opts Options) (ServiceDiscovery, error) {
	_ = "STUB: not implemented"
	return *new(ServiceDiscovery), nil
}

// InstallContractDiscovery install contract service client
func InstallContractDiscovery(name string, f func(opts Options) ContractDiscovery) {
	_ = "STUB: not implemented"
	return
}

// ServiceDiscovery fetch service and instances from remote or local
type ServiceDiscovery interface {
	GetMicroService(microServiceID string) (*MicroService, error)
	FindMicroServiceInstances(consumerID, microServiceName string, tags utiltags.Tags) ([]*MicroServiceInstance, error)
	AutoSync()
	Close() error
}

// DefaultServiceDiscoveryService supplies service discovery
var DefaultServiceDiscoveryService ServiceDiscovery

// DefaultContractDiscoveryService supplies contract discovery
var DefaultContractDiscoveryService ContractDiscovery

// ContractDiscovery fetch schema content from remote or local
type ContractDiscovery interface {
	GetMicroServicesByInterface(interfaceName string) (microservices []*MicroService)
	GetSchemaContentByInterface(interfaceName string) SchemaContent
	GetSchemaContentByServiceName(svcName, version, appID, env string) []*SchemaContent
	Close() error
}

func enableServiceDiscovery(opts Options) error { _ = "STUB: not implemented"; return nil }

func enableContractDiscovery(opts Options) { _ = "STUB: not implemented"; return }
