package config

import (
	"errors"

	"github.com/go-chassis/go-chassis/v2/core/config/model"
)

// GlobalDefinition is having the information about region, load balancing, service center, config server,
// protocols, and handlers for the micro service
var GlobalDefinition *model.GlobalCfg
var lbConfig *model.LBWrapper

// MicroserviceDefinition has info about application id, provider info, description of the service,
// and description of the instance
var MicroserviceDefinition *model.ServiceSpec

// MonitorCfgDef has monitor info, including zipkin and apm.
var MonitorCfgDef *model.MonitorCfg

// HystrixConfig is having info about isolation, circuit breaker, fallback properities of the micro service
var HystrixConfig *model.HystrixConfigWrapper

// ErrNoName is used to represent the service name missing error
var ErrNoName = errors.New("micro service name is missing in description file")

// GetConfigServerConf return config server conf
func GetConfigServerConf() model.ConfigClient {
	_ = "STUB: not implemented"
	return *new(model.ConfigClient)
}

// GetTransportConf return transport settings
func GetTransportConf() model.Transport { _ = "STUB: not implemented"; return *new(model.Transport) }

// GetDataCenter return data center info
func GetDataCenter() *model.DataCenterInfo { _ = "STUB: not implemented"; return nil }

// GetAPM return monitor config info
func GetAPM() model.APMStruct { _ = "STUB: not implemented"; return *new(model.APMStruct) }

// readFromArchaius unmarshal configurations to expected pointer
func readFromArchaius() error { _ = "STUB: not implemented"; return nil }

// populateServiceRegistryAddress populate service registry address
func populateServiceRegistryAddress() {
	_ = "STUB: not implemented"
	// Registry Address , higher priority for environment variable
	return
}

// populateConfigServerAddress populate config server address
func populateConfigServerAddress() {
	_ = "STUB: not implemented"
	// config server Address , higher priority for environment variable
	return
}

// readEndpoint
func readEndpoint(env string) string { _ = "STUB: not implemented"; return "" }

// populateServiceEnvironment populate service environment
func populateServiceEnvironment() { _ = "STUB: not implemented"; return }

// populateServiceName populate service name
func populateServiceName() { _ = "STUB: not implemented"; return }

// populateVersion populate version
func populateVersion() { _ = "STUB: not implemented"; return }

func populateApp() { _ = "STUB: not implemented"; return }

// ReadGlobalConfigFromArchaius for to unmarshal the global config file(chassis.yaml) information
func ReadGlobalConfigFromArchaius() error { _ = "STUB: not implemented"; return nil }

// ReadLBFromArchaius for to unmarshal the global config file(chassis.yaml) information
func ReadLBFromArchaius() error { _ = "STUB: not implemented"; return nil }

// ReadMonitorFromArchaius read monitor config from archauis pkg
func ReadMonitorFromArchaius() error { _ = "STUB: not implemented"; return nil }

// ReadHystrixFromArchaius is unmarshal hystrix configuration file(circuit_breaker.yaml)
func ReadHystrixFromArchaius() error { _ = "STUB: not implemented"; return nil }

// GetLoadBalancing return lb config
func GetLoadBalancing() *model.LoadBalancing { _ = "STUB: not implemented"; return nil }

// GetHystrixConfig return cb config
func GetHystrixConfig() *model.HystrixConfig { _ = "STUB: not implemented"; return nil }

// Init is initialize the configuration directory, archaius, route rule, and schema
func Init() error { _ = "STUB: not implemented"; return nil }

//Upload schemas using environment variable SCHEMA_ROOT

//set micro service names
