package fileutil

import (
	"sync"
)

const (
	//ChassisConfDir is constant of type string
	ChassisConfDir = "CHASSIS_CONF_DIR"
	//ChassisHome is constant of type string
	ChassisHome = "CHASSIS_HOME"
	//SchemaDirectory is constant of type string
	SchemaDirectory = "schema"
)

const (
	//Global is a constant of type string
	Global = "chassis.yaml"
	//LoadBalancing is constant of type string
	LoadBalancing = "load_balancing.yaml"
	//RateLimiting is constant of type string
	RateLimiting = "rate_limiting.yaml"
	//Definition is constant of type string
	Definition = "microservice.yaml"
	//Hystric is constant of type string
	Hystric = "circuit_breaker.yaml"
	//PaasLager is constant of type string
	PaasLager = "lager.yaml"
	//TLS is constant of type string
	TLS = "tls.yaml"
	//Monitoring is constant of type string
	Monitoring = "monitoring.yaml"
	//Auth is constant of type string
	Auth = "auth.yaml"
	//Tracing is constant of type string
	Tracing = "tracing.yaml"
	//Router is constant of type string
	Router = "router.yaml"
)

var configDir string
var homeDir string
var once sync.Once

// GetWorkDir is a function used to get the working directory
func GetWorkDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func initDir() { _ = "STUB: not implemented"; return }

// set conf dir, CHASSIS_CONF_DIR has highest priority

// CHASSIS_HOME has second most high priority

// ChassisHomeDir is function used to get the home directory of chassis
func ChassisHomeDir() string { _ = "STUB: not implemented"; return "" }

// GetConfDir is a function used to get the configuration directory
func GetConfDir() string { _ = "STUB: not implemented"; return "" }

// CircuitBreakerConfigPath is a function used to join .yaml file name with configuration path
func CircuitBreakerConfigPath() string { _ = "STUB: not implemented"; return "" }

// GetDefinition is a function used to join .yaml file name with configuration path
func GetDefinition() string { _ = "STUB: not implemented"; return "" }

// LoadBalancingConfigPath is a function used to join .yaml file name with configuration directory
func LoadBalancingConfigPath() string { _ = "STUB: not implemented"; return "" }

// RateLimitingFile is a function used to join .yaml file name with configuration directory
func RateLimitingFile() string { _ = "STUB: not implemented"; return "" }

// TLSConfigPath is a function used to join .yaml file name with configuration directory
func TLSConfigPath() string { _ = "STUB: not implemented"; return "" }

// MonitoringConfigPath is a function used to join .yaml file name with configuration directory
func MonitoringConfigPath() string { _ = "STUB: not implemented"; return "" }

// MicroserviceDefinition is a function used to join .yaml file name with configuration directory
func MicroserviceDefinition(microserviceName string) string { _ = "STUB: not implemented"; return "" }

// MicroServiceConfigPath is a function used to join .yaml file name with configuration directory
func MicroServiceConfigPath() string { _ = "STUB: not implemented"; return "" }

// GlobalConfigPath is a function used to join .yaml file name with configuration directory
func GlobalConfigPath() string { _ = "STUB: not implemented"; return "" }

// LogConfigPath is a function used to join .yaml file name with configuration directory
func LogConfigPath() string { _ = "STUB: not implemented"; return "" }

// RouterConfigPath is a function used to join .yaml file name with configuration directory
func RouterConfigPath() string { _ = "STUB: not implemented"; return "" }

// AuthConfigPath is a function used to join .yaml file name with configuration directory
func AuthConfigPath() string { _ = "STUB: not implemented"; return "" }

// TracingPath is a function used to join .yaml file name with configuration directory
func TracingPath() string { _ = "STUB: not implemented"; return "" }

// SchemaDir is a function used to join .yaml file name with configuration path
func SchemaDir(microserviceName string) string { _ = "STUB: not implemented"; return "" }
