package registry

import (
	"errors"
)

var errEmptyServiceIDFromRegistry = errors.New("got empty serviceID from registry")

// InstanceEndpoints instance endpoints
var InstanceEndpoints = make(map[string]string)

// RegisterService register micro-service
func RegisterService() error { _ = "STUB: not implemented"; return nil }

// from yaml setting

// TODO allows to customize microservice alias

//update metadata

// if the microservice is allowed to be called by consumers with different appId,
// this means that the governance configuration of the consumer side needs to
// support key format with appid, like 'cse.loadbalance.{alias}.strategy.name'.

// RegisterServiceInstances register micro-service instances
func RegisterServiceInstances() error { _ = "STUB: not implemented"; return nil }

// from yaml setting

//Set to runtime
