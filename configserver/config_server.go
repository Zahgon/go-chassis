package configserver

import (
	"crypto/tls"
	"errors"
)

const (
	//configServerName is a variable of type string of config server
	configServerName = "configServer"
)

// ErrRefreshMode means config is mis used
var (
	ErrRefreshMode      = errors.New("refreshMode must be 0 or 1")
	ErrRegistryDisabled = errors.New("discovery is disabled")
)

// Init initialize config server
func Init() error { _ = "STUB: not implemented"; return nil }

/*This condition added because member discovery can have multiple ip's with IsHTTPS
having both true and false value.*/

// GetConfigServerEndpoint will read local config server uri first, if there is not,
// it will try to discover config server from registry
func GetConfigServerEndpoint() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getTLSForClient(configServerURL string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initConfigServer(endpoint string, enableSSL bool, tlsConfig *tls.Config, interval int) error {
	_ = "STUB: not implemented"
	return nil
}

func refreshGlobalConfig() error { _ = "STUB: not implemented"; return nil }
