package config

// GetServiceDiscoveryType returns the Type of SD registry
func GetServiceDiscoveryType() string { _ = "STUB: not implemented"; return "" }

// GetServiceDiscoveryAddress returns the Address of SD registry
func GetServiceDiscoveryAddress() string { _ = "STUB: not implemented"; return "" }

// GetServiceDiscoveryRefreshInterval returns the RefreshInterval of SD registry
func GetServiceDiscoveryRefreshInterval() string { _ = "STUB: not implemented"; return "" }

// GetServiceDiscoveryWatch returns the Watch of SD registry
func GetServiceDiscoveryWatch() bool { _ = "STUB: not implemented"; return false }

// GetServiceDiscoveryAPIVersion returns the APIVersion of SD registry
func GetServiceDiscoveryAPIVersion() string { _ = "STUB: not implemented"; return "" }

// GetServiceDiscoveryDisable returns the Disable of SD registry
func GetServiceDiscoveryDisable() bool { _ = "STUB: not implemented"; return false }

// GetServiceDiscoveryHealthCheck returns the HealthCheck of SD registry
func GetServiceDiscoveryHealthCheck() bool { _ = "STUB: not implemented"; return false }

// GetServiceDiscoveryUploadSchema returns if should register schema of SD registry
func GetServiceDiscoveryUploadSchema() bool { _ = "STUB: not implemented"; return false }

// DefaultConfigPath set the default config path
const DefaultConfigPath = "/etc/.kube/config"

// GetServiceDiscoveryConfigPath returns the configpath of SD registry
func GetServiceDiscoveryConfigPath() string { _ = "STUB: not implemented"; return "" }
