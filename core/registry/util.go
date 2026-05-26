package registry

import (
	"time"

	"crypto/tls"

	"github.com/go-chassis/go-chassis/v2/core/config/model"
)

const (
	protocolSymbol = "://"
)

// GetProtocolMap returns the protocol map
func GetProtocolMap(eps []string) (map[string]*Endpoint, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// GetProtocolList returns the protocol list
func GetProtocolList(m map[string]*Endpoint) []string { _ = "STUB: not implemented"; return nil }

// MakeEndpoints returns the endpoints
func MakeEndpoints(m map[string]model.Protocol) []string { _ = "STUB: not implemented"; return nil }

// MakeEndpointMap returns the endpoints map
func MakeEndpointMap(m map[string]model.Protocol) (map[string]*Endpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FillUnspecifiedIP replace 0.0.0.0 or :: IPv4 and IPv6 unspecified IP address with local NIC IP.
func FillUnspecifiedIP(host string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Microservice2ServiceKeyStr prepares a microservice key
func Microservice2ServiceKeyStr(m *MicroService) string { _ = "STUB: not implemented"; return "" }

const (
	initialInterval = 5 * time.Second
	maxInterval     = 3 * time.Minute
)

func startBackOff(operation func() error) { _ = "STUB: not implemented"; return }

// URIs2Hosts return hosts and scheme
func URIs2Hosts(uris []string) ([]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

//not uri. but still permitted, like zookeeper,file system

func getTLSConfig(scheme, t string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDuration return the time.Duration type value by specified key
func GetDuration(key string, def time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
