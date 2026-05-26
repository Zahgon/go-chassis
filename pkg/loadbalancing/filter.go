package loadbalancing

import (
	"github.com/go-chassis/go-chassis/v2/core/loadbalancer"
	"github.com/go-chassis/go-chassis/v2/core/registry"
)

func init() {
	loadbalancer.InstallFilter(loadbalancer.ZoneAware, FilterAvailableZoneAffinity)
}

// FilterAvailableZoneAffinity is a region and zone based Select Filter which will Do the selection of instance in the same region and zone, if not Do the selection of instance in any zone in same region , if not Do the selection of instance in any zone of any region
func FilterAvailableZoneAffinity(old []*registry.MicroServiceInstance, c []*loadbalancer.Criteria) []*registry.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

// Either no information or partial data center information specified, return all instances

//out of region (multi region) case

//same region but any available zone case

//same region and same zone case

// getInstancesZoneWise check for the same zone and region
func getInstancesZoneWise(providerInstances []*registry.MicroServiceInstance, region, availableZone string) []*registry.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

// getAvailableInstancesInSameRegion check for available instances in same region
func getAvailableInstancesInSameRegion(providerInstances []*registry.MicroServiceInstance, region string) []*registry.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

// FilterByMetadata filter instances based meta data
func FilterByMetadata(old []*registry.MicroServiceInstance, c []*loadbalancer.Criteria) []*registry.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

//TODO read tags in router.yaml and filter instances based on properties and tags
