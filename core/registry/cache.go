package registry

import (
	"github.com/patrickmn/go-cache"
)

const (
	//DefaultExpireTime default expiry time is kept as 0
	DefaultExpireTime = 0
)

// MicroserviceInstanceIndex key: ServiceName, value: []instance
var MicroserviceInstanceIndex CacheIndex

// ipIndexedCache is for caching map of instance IP and service information
// key: instance ip, value: SourceInfo
var ipIndexedCache *cache.Cache

// SchemaInterfaceIndexedCache key: schema interface name value: []*microservice
var SchemaInterfaceIndexedCache *cache.Cache

// SchemaServiceIndexedCache key: schema service name value: []*microservice
var SchemaServiceIndexedCache *cache.Cache

// ProvidersMicroServiceCache  key: micro service  name and appId, value: []*MicroService
var ProvidersMicroServiceCache *cache.Cache

func initCache() *cache.Cache { _ = "STUB: not implemented"; return nil }

// EnableRegistryCache init caches
func EnableRegistryCache() { _ = "STUB: not implemented"; return }

// CacheIndex is a unified local instances cache manager
type CacheIndex interface {
	Get(service string, tags map[string]string) ([]*MicroServiceInstance, bool)
	//Set will overwrite all instances correspond to a service name
	Set(service string, instances []*MicroServiceInstance)
	FullCache() *cache.Cache
	Delete(service string)
}

// SetIPIndex save ip index
func SetIPIndex(ip string, si *SourceInfo) { _ = "STUB: not implemented"; return }

// GetIPIndex get ip corresponding source info
func GetIPIndex(ip string) *SourceInfo { _ = "STUB: not implemented"; return nil }

// GetProvidersFromCache get local provider simpleCache
func GetProvidersFromCache() []*MicroService { _ = "STUB: not implemented"; return nil }

// AddProviderToCache refresh provider simpleCache
func AddProviderToCache(serverName, appID string) { _ = "STUB: not implemented"; return }
