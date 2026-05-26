package servicecenter

import (
	scregistry "github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/go-chassis/v2/core/registry"
	"github.com/go-chassis/sc-client"
)

// parseSchemaContent parse schema content into SchemaContent structure
func parseSchemaContent(content []byte) (registry.SchemaContent, error) {
	_ = "STUB: not implemented"
	return *new(registry.SchemaContent), nil
}

// parseSchemaContent parse schema content into SchemaContent structure
func unmarshalSchemaContent(content []byte) (*registry.SchemaContent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filterInstances filter instances
func filterInstances(providerInstances []*scregistry.MicroServiceInstance) []*registry.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

func closeClient(r *sc.Client) error { _ = "STUB: not implemented"; return nil }

func wrapTagsForServiceCenter(t utiltags.Tags) utiltags.Tags {
	_ = "STUB: not implemented"
	return *new(utiltags.Tags)
}

//if app and version is empty, need to find with latest version in same app

// GetCriteria generate batch find criteria from provider cache
func GetCriteria() []*scregistry.FindService { _ = "STUB: not implemented"; return nil }

// GetCriteriaByService generate batch find criteria from provider cache with same service name and different app
func GetCriteriaByService(sn string) []*scregistry.FindService {
	_ = "STUB: not implemented"
	return nil
}
