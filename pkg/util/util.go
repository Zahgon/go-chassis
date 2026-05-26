package util

import (
	"errors"
)

var (
	//ErrInvalidPortName happens if your port name is illegal
	ErrInvalidPortName = errors.New("invalid port name, port name must be {protocol}-<suffix>")
	//ErrInvalidURL happens if your utl is illegal
	ErrInvalidURL = errors.New("invalid url, url must be {protocol}://{service-name}:{port-name}")
)

// ParsePortName a port name is composite by protocol-name,like http-admin,http-api,grpc-console,grpc-api
// ParsePortName return two string separately
func ParsePortName(n string) (string, string, error) { _ = "STUB: not implemented"; return "", "", nil }

// ParseServiceAndPort returns service name and port name
func ParseServiceAndPort(n string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// GenProtoEndPoint generate proto and port
func GenProtoEndPoint(proto, port string) string { _ = "STUB: not implemented"; return "" }
