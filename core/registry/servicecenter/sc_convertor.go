package servicecenter

import (
	scregistry "github.com/go-chassis/cari/discovery"
	"github.com/go-chassis/go-chassis/v2/core/registry"
	"github.com/go-chassis/sc-client"
)

// ToMicroService assign sc micro-service to go chassis micro-service
func ToMicroService(scs *scregistry.MicroService) *registry.MicroService {
	_ = "STUB: not implemented"
	return nil
}

// ToSCService assign go chassis micro-service to the sc micro-service
func ToSCService(cs *registry.MicroService) *scregistry.MicroService {
	_ = "STUB: not implemented"
	return nil
}

// ToMicroServiceInstance assign model micro-service instance parameters to registry micro-service instance parameters
func ToMicroServiceInstance(ins *scregistry.MicroServiceInstance) *registry.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

// ToSCInstance assign registry micro-service instance parameters to model micro-service instance parameters
func ToSCInstance(msi *registry.MicroServiceInstance) *scregistry.MicroServiceInstance {
	_ = "STUB: not implemented"
	return nil
}

// ToSCOptions convert registry opstions into sc client options
func ToSCOptions(options registry.Options) sc.Options {
	_ = "STUB: not implemented"
	return *new(sc.Options)
}
