package servicecomb

import (
	"github.com/go-chassis/go-chassis/v2/control"
	"github.com/go-chassis/go-chassis/v2/core/config/model"
)

// SaveToLBCache save configs
func SaveToLBCache(raw *model.LoadBalancing) { _ = "STUB: not implemented"; return }

// if there is no config, none key will be updated

// remove outdated keys

func saveDefaultLB(raw *model.LoadBalancing) string {
	_ = "STUB: not implemented" // return updated key
	return ""
}

func parseFilters(filters string) []string { _ = "STUB: not implemented"; return nil }

//delete empty string

func saveEachLB(k string, raw model.LoadBalancingSpec) string {
	_ = "STUB: not implemented" // return updated key
	return ""
}

func setDefaultLBValue(c *control.LoadBalancingConfig) { _ = "STUB: not implemented"; return }

// SaveToCBCache save configs
func SaveToCBCache(raw *model.HystrixConfig) { _ = "STUB: not implemented"; return }

// if there is no config, none key will be updated

// remove outdated keys

func saveEachCB(serviceName, serviceType string) string {
	_ = "STUB: not implemented" //return updated key
	return ""
}

// GetCBCacheKey generate cache key
func GetCBCacheKey(serviceName, serviceType string) string { _ = "STUB: not implemented"; return "" }

func reloadLBCache(src *model.LoadBalancing) map[string]bool {
	_ = "STUB: not implemented" //return updated keys
	return nil
}

func reloadCBCache(src *model.HystrixConfig) map[string]bool {
	_ = "STUB: not implemented" //return updated keys
	return nil
}

// global level config

// get all services who have configs

// if a service has configurations of IsolationProperties|
// CircuitBreakerProperties|FallbackPolicyProperties|FallbackProperties,
// it's configuration should be added to cache when framework starts

// remove duplicate service names

// service level config

func getServiceNamesByServiceTypeAndAnyService(i interface{}, serviceType string) (services []string, err error) {
	_ = "STUB: not implemented"
	// check type
	return nil, nil
}

// check value

// check type

// check value

//check type

// check value

// get service names
