package registry

import (
	"sync"

	"github.com/patrickmn/go-cache"
)

// IndexCache return instances by criteria
type IndexCache struct {
	latestV    map[string]string //save every service's latest version number
	muxLatestV sync.RWMutex

	simpleCache *cache.Cache //save service name and correspond instances

	//key must contain service name, cache key includes label key values
	indexedCache *cache.Cache

	CriteriaStore []map[string]string //all criteria need to be saved in here so that we can update indexedCache, during Set process
}

// NewIndexCache create a cache which saves and manage instances
func NewIndexCache() *IndexCache { _ = "STUB: not implemented"; return nil }

// FullCache return all instances
func (ic *IndexCache) FullCache() *cache.Cache { _ = "STUB: not implemented"; return nil }

// Delete remove one service's instances
func (ic *IndexCache) Delete(k string) { _ = "STUB: not implemented"; return }

// Set overwrite instances cache
func (ic *IndexCache) Set(k string, instances []*MicroServiceInstance) {
	_ = "STUB: not implemented"
	return
}

//update latest version number

////TODO update indexed cache
//ic.muxCriteria.RLock()
//for _, criteria := range ic.CriteriaStore {
//	indexKey := ic.GetIndexedCacheKey(k, criteria)
//	result := make([]*MicroServiceInstance, 0)
//	for _, instance := range instances {
//		if instance.Has(criteria) {
//			result = append(result, instance)
//		}
//	}
//	//forcely overwrite indexed cache, that is safe
//	ic.indexedCache.Set(indexKey, result, 0)
//}
//ic.muxCriteria.RUnlock()

// Get return instances cache by criteria
func (ic *IndexCache) Get(k string, tags map[string]string) ([]*MicroServiceInstance, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//if version is latest, then set it to real version

//find from indexed cache first

//no result, then find it and save result

//ic.indexedCache.Set(indexKey, queryResult, 0)
//ic.muxCriteria.Lock()
//ic.CriteriaStore = append(ic.CriteriaStore, tags)
//ic.muxCriteria.Unlock()

func (ic *IndexCache) setTagsBeforeQuery(k string, tags map[string]string) {
	_ = "STUB: not implemented"
	return

	//must set version before query
}

// GetIndexedCacheKey combine keys in order, use sets to return sorted list
func GetIndexedCacheKey(service string, tags map[string]string) (ss string) {
	_ = "STUB: not implemented"
	return ""
}
