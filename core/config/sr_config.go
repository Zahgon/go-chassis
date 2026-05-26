package config

import (
	"github.com/go-chassis/cari/rbac"
)

// GetRegistratorType returns the Type of service registry
func GetRegistratorType() string { _ = "STUB: not implemented"; return "" }

// GetRegistratorAddress returns the Address of service registry
func GetRegistratorAddress() string { _ = "STUB: not implemented"; return "" }

// GetRegistratorScope returns the Scope of service registry
func GetRegistratorScope() string { _ = "STUB: not implemented"; return "" }

// GetRegistratorAutoRegister returns the AutoRegister of service registry
func GetRegistratorAutoRegister() string { _ = "STUB: not implemented"; return "" }

// GetRegistratorAPIVersion returns the APIVersion of service registry
func GetRegistratorAPIVersion() string { _ = "STUB: not implemented"; return "" }

// GetRegistratorDisable returns the Disable of service registry
func GetRegistratorDisable() bool { _ = "STUB: not implemented"; return false }

// GetRegistratorRbacAccount returns the RbacAccout info of service registry
func GetRegistratorRbacAccount() *rbac.AuthUser { _ = "STUB: not implemented"; return nil }
