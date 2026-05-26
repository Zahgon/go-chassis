package servicecenter

import (
	scregistry "github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/sc-client"

	"github.com/go-chassis/go-chassis/v2/core/registry"
)

const (
	// ServiceCenter constant string
	ServiceCenter = "servicecenter"
)

// Registrator to represent the object of service center to call the APIs of service center
type Registrator struct {
	Name           string
	registryClient *sc.Client
	opts           sc.Options
}

// RegisterService : 注册微服务
func (r *Registrator) RegisterService(ms *registry.MicroService) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RegisterServiceInstance : 注册微服务
func (r *Registrator) RegisterServiceInstance(sid string, cIns *registry.MicroServiceInstance) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// RegisterServiceAndInstance : 注册微服务
func (r *Registrator) RegisterServiceAndInstance(cMicroService *registry.MicroService, cInstance *registry.MicroServiceInstance) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// UnRegisterMicroServiceInstance : 去注册微服务实例
func (r *Registrator) UnRegisterMicroServiceInstance(microServiceID, microServiceInstanceID string) error {
	_ = "STUB: not implemented"
	return nil
}

// WSHeartbeat : Keep instance heartbeats.
func (r *Registrator) WSHeartbeat(microServiceID string, microServiceInstanceID string, callback func()) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Heartbeat : Keep instance heartbeats.
func (r *Registrator) Heartbeat(microServiceID, microServiceInstanceID string) (bool, error) {
	_ = "STUB: not implemented"
	// use http
	return false, nil
}

// AddSchemas to service center
func (r *Registrator) AddSchemas(microServiceID, schemaName, schemaInfo string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateMicroServiceInstanceStatus : 更新微服务实例状态信息
func (r *Registrator) UpdateMicroServiceInstanceStatus(microServiceID, microServiceInstanceID, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateMicroServiceProperties 更新微服务properties信息
func (r *Registrator) UpdateMicroServiceProperties(microServiceID string, properties map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateMicroServiceInstanceProperties : 更新微服务实例properties信息
func (r *Registrator) UpdateMicroServiceInstanceProperties(microServiceID, microServiceInstanceID string, properties map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Close : Close all client connection.
func (r *Registrator) Close() error { _ = "STUB: not implemented"; return nil }

// ServiceDiscovery to represent the object of service center to call the APIs of service center
type ServiceDiscovery struct {
	Name           string
	registryClient *sc.Client
	opts           sc.Options
}

// GetAllMicroServices : Get all MicroService information.
func (r *ServiceDiscovery) GetAllMicroServices() ([]*registry.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAllApplications : Get all Applications information.
func (r *ServiceDiscovery) GetAllApplications() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMicroService : 根据microServiceID获取对应的微服务信息
func (r *ServiceDiscovery) GetMicroService(microServiceID string) (*registry.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMicroServiceInstances : 获取指定微服务的所有实例
func (r *ServiceDiscovery) GetMicroServiceInstances(consumerID, providerID string) ([]*registry.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindMicroServiceInstances find micro-service instances
func (r *ServiceDiscovery) FindMicroServiceInstances(consumerID, microServiceName string, tags utiltags.Tags) ([]*registry.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	// TODO: wrap default tags for service center
	// because sc need version and appID to generate tags
	return nil, nil
}

// RegroupInstances organize raw data to better format
func RegroupInstances(keys []*scregistry.FindService, response *scregistry.BatchFindInstancesResponse) map[string][]*registry.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

// WatchMicroService : 支持用户自调用主动监听实例变化功能
func (r *ServiceDiscovery) WatchMicroService(selfMicroServiceID string, callback func(*sc.MicroServiceInstanceChangedEvent)) {
	_ = "STUB: not implemented"
	return
}

// AutoSync updating the cache manager
func (r *ServiceDiscovery) AutoSync() { _ = "STUB: not implemented"; return }

// Close : Close all websocket connection.
func (r *ServiceDiscovery) Close() error { _ = "STUB: not implemented"; return nil }

// ContractDiscovery to represent the object of service center to call the APIs of service center
type ContractDiscovery struct {
	Name           string
	registryClient *sc.Client
	opts           sc.Options
}

// GetMicroServicesByInterface get micro-services by interface
func (r *ContractDiscovery) GetMicroServicesByInterface(interfaceName string) (microService []*registry.MicroService) {
	_ = "STUB: not implemented"
	return nil
}

// GetSchemaContentByInterface get schema content by interface
func (r *ContractDiscovery) GetSchemaContentByInterface(interfaceName string) (schemas registry.SchemaContent) {
	_ = "STUB: not implemented"
	return *new(registry.SchemaContent)
}

// GetSchemaContentByServiceName get schema content by service name
func (r *ContractDiscovery) GetSchemaContentByServiceName(svcName, version, appID, env string) (schemas []*registry.SchemaContent) {
	_ = "STUB: not implemented"
	return nil
}

// fillSchemaServiceIndexCache fill schema service index cache
func (r *ContractDiscovery) fillSchemaServiceIndexCache(ms []*scregistry.MicroService, serviceID string) (content []*registry.SchemaContent) {
	_ = "STUB: not implemented"
	return nil
}

// fillCacheAndGetServiceSchemaContent fill cache and get services schema content
func (r *ContractDiscovery) fillCacheAndGetServiceSchemaContent(microServiceList []*scregistry.MicroService, serviceID string) (schemaContent []*registry.SchemaContent) {
	_ = "STUB: not implemented"
	return nil
}

// fillSchemaInterfaceIndexCache fill schema interface index cache
func (r *ContractDiscovery) fillSchemaInterfaceIndexCache(ms []*scregistry.MicroService, interfaceName string) (content registry.SchemaContent) {
	_ = "STUB: not implemented"
	return *new(registry.SchemaContent)
}

// fillCacheAndGetInterfaceSchemaContent fill cache and get interface schema content
func (r *ContractDiscovery) fillCacheAndGetInterfaceSchemaContent(microServiceList []*scregistry.MicroService, interfaceName string) (schemaContent registry.SchemaContent) {
	_ = "STUB: not implemented"
	return *new(registry.SchemaContent)
}

// GetSchema from service center
func (r *ContractDiscovery) GetSchema(microServiceID, schemaName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close close client connection
func (r *ContractDiscovery) Close() error { _ = "STUB: not implemented"; return nil }

// NewRegistrator new Service center registrator
func NewRegistrator(options registry.Options) registry.Registrator {
	_ = "STUB: not implemented"
	return *new(registry.Registrator)
}

// NewServiceDiscovery new service center discovery
func NewServiceDiscovery(options registry.Options) registry.ServiceDiscovery {
	_ = "STUB: not implemented"
	return *new(registry.ServiceDiscovery)
}

func newContractDiscovery(options registry.Options) registry.ContractDiscovery {
	_ = "STUB: not implemented"
	return *new(registry.ContractDiscovery)
}

// init initialize the plugin of service center registry
func init() {
	registry.InstallRegistrator(ServiceCenter, NewRegistrator)
	registry.InstallServiceDiscovery(ServiceCenter, NewServiceDiscovery)
	registry.InstallContractDiscovery(ServiceCenter, newContractDiscovery)

}
