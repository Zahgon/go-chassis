package servicecenter

import (
	scregistry "github.com/go-chassis/cari/discovery"

	"time"

	"github.com/go-chassis/go-chassis/v2/core/registry"
	"github.com/go-chassis/sc-client"
	"k8s.io/apimachinery/pkg/util/sets"
)

// constant values for default expiration time, and refresh interval
const (
	DefaultExpireTime      = 0
	DefaultRefreshInterval = time.Second * 30
)

// constant values for checking instance ID status
const (
	InstanceIDIsExist    = "instanceIdIsExist"
	InstanceIDIsNotExist = "instanceIdIsNotExist"
)

// CacheManager cache manager
type CacheManager struct {
	registryClient *sc.Client
}

// AutoSync automatically sync the running instances
func (c *CacheManager) AutoSync() { _ = "STUB: not implemented"; return }

// refreshCache refresh cache
func (c *CacheManager) refreshCache() { _ = "STUB: not implemented"; return }

//connection with sc may lost, reset the revision

// MakeIPIndex make ip index
// if store instance metadata into tags
// it will be used in route management
func (c *CacheManager) MakeIPIndex() error { _ = "STUB: not implemented"; return nil }

//no need to analyze each endpoint, so break

// MakeSchemaIndex make schema index
func (c *CacheManager) MakeSchemaIndex() error { _ = "STUB: not implemented"; return nil }

// This functions checks if the microservices exist in the list passed in argument
func checkIfMicroServiceExistInList(microserviceList []*scregistry.MicroService, serviceID string) bool {
	_ = "STUB: not implemented"
	return false
}

// pullMicroServiceInstance pull micro-service instance
func (c *CacheManager) pullMicroServiceInstance() error {
	_ = "STUB: not implemented"
	// Get Providers
	return nil
}

//fetch remote based on app and service

func (c *CacheManager) compareAndDeleteOutdatedProviders(newProviders sets.String) {
	_ = "STUB: not implemented"
	return
}

//provider is outdated, delete it

// getServiceSet regroup the providers by service name
func getServiceSet(exist []*scregistry.FindService) (sets.String, map[string]sets.String) {
	_ = "STUB: not implemented"
	//get Provider's instances
	return *new(sets.String), nil
}

// key is serviceName
// key is "serviceName" value is app sets

// set app into instance metadata, split instances into ups and downs
// set instance to cache by service name
func filterAndCache(services sets.String, providerInstances map[string][]*registry.MicroServiceInstance) {
	_ = "STUB: not implemented"
	//append instances from different app and same service name into one unified slice
	return
}

//save cache after get all instances of a service name

func setEmptyCache(services sets.String) { _ = "STUB: not implemented"; return }

// watch watching micro-service instance status
func watch(response *sc.MicroServiceInstanceChangedEvent) { _ = "STUB: not implemented"; return }

// createAction added micro-service instance to the cache
func createAction(response *sc.MicroServiceInstanceChangedEvent) { _ = "STUB: not implemented"; return }

// deleteAction delete micro-service instance
func deleteAction(response *sc.MicroServiceInstanceChangedEvent) { _ = "STUB: not implemented"; return }

// updateAction update micro-service instance event
func updateAction(response *sc.MicroServiceInstanceChangedEvent) { _ = "STUB: not implemented"; return }
