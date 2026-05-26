package mock

import (
	"github.com/go-chassis/go-chassis/v2/core/registry"
	"github.com/go-chassis/sc-client"
	"github.com/stretchr/testify/mock"
)

// RegistratorMock struct for register mock
type RegistratorMock struct {
	mock.Mock
}

// Close to close the registry mock
func (m *RegistratorMock) Close() error {
	_ = "STUB: not implemented"

	// RegisterServiceInstance register service center instance
	return nil
}

func (m *RegistratorMock) RegisterServiceInstance(sid string, instance *registry.MicroServiceInstance) (string, error) {
	_ = "STUB: not implemented"

	// RegisterService register service
	return "", nil
}

func (m *RegistratorMock) RegisterService(microservice *registry.MicroService) (string, error) {
	_ = "STUB: not implemented"

	// RegisterServiceAndInstance register service and instance
	return "", nil
}

func (m *RegistratorMock) RegisterServiceAndInstance(microService *registry.MicroService, instance *registry.MicroServiceInstance) (string, string, error) {
	_ = "STUB: not implemented"
	return "",

		// Heartbeat heart beat
		"", nil
}

func (m *RegistratorMock) Heartbeat(microServiceID, microServiceInstanceID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Heartbeat heart beat
func (m *RegistratorMock) WSHeartbeat(microServiceID, microServiceInstanceID string, callback func()) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AddDependencies add dependencies
func (m *RegistratorMock) AddDependencies(request *registry.MicroServiceDependency) error {
	_ = "STUB: not implemented"

	// UpdateMicroServiceInstanceStatus update micro-service instance status
	return nil
}

func (m *RegistratorMock) UpdateMicroServiceInstanceStatus(microServiceID, microServiceInstanceID, status string) error {
	_ = "STUB: not implemented"

	// UpdateMicroServiceProperties update micro-service properties
	return nil
}

func (m *RegistratorMock) UpdateMicroServiceProperties(microServiceID string, properties map[string]string) error {
	_ = "STUB: not implemented"

	// UpdateMicroServiceInstanceProperties update micro-service instance properties
	return nil
}

func (m *RegistratorMock) UpdateMicroServiceInstanceProperties(microServiceID, microServiceInstanceID string, properties map[string]string) error {
	_ = "STUB: not implemented"

	// UnRegisterMicroServiceInstance unregistered micro-service instance
	return nil
}

func (m *RegistratorMock) UnRegisterMicroServiceInstance(microServiceID, microServiceInstanceID string) error {
	_ = "STUB: not implemented"

	// AddSchemas add schemas
	return nil
}

func (m *RegistratorMock) AddSchemas(microServiceID, schemaName, schemaInfo string) error {
	_ = "STUB: not implemented"

	// DiscoveryMock struct for disco mock
	return nil
}

type DiscoveryMock struct {
	mock.Mock
}

// GetMicroServiceID get micro-service id
func (m *DiscoveryMock) GetMicroServiceID(appID, microServiceName, version, env string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetAllMicroServices get all microservices
func (m *DiscoveryMock) GetAllMicroServices() ([]*registry.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAllApplications get all applications
func (m *DiscoveryMock) GetAllApplications() ([]string, error) {
	_ = "STUB: not implemented"
	return nil,

		// GetMicroService get micro service
		nil
}

func (m *DiscoveryMock) GetMicroService(microServiceID string) (*registry.MicroService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMicroServiceInstances get micro-service instances
func (m *DiscoveryMock) GetMicroServiceInstances(consumerID, providerID string) ([]*registry.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindMicroServiceInstances find micro-service instances
func (m *DiscoveryMock) FindMicroServiceInstances(consumerID, microServiceName string, tags utiltags.Tags) ([]*registry.MicroServiceInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WatchMicroService watch micro-service
func (m *DiscoveryMock) WatchMicroService(selfMicroServiceID string, callback func(*sc.MicroServiceInstanceChangedEvent)) {
	_ = "STUB: not implemented"

	// AutoSync auto sync
	return
}

func (m *DiscoveryMock) AutoSync() {
	_ = "STUB: not implemented"

	// Close mock
	return
}

func (m *DiscoveryMock) Close() error { _ = "STUB: not implemented"; return nil }
