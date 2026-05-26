package iputil

import (
	"crypto/tls"
	"net"
	"net/http"
)

// Localhost is a function which returns localhost IP address
func Localhost() string {
	_ = "STUB: not implemented"

	// GetLocalIP 获得本机IP
	return ""
}

func GetLocalIP() string { _ = "STUB: not implemented"; return "" }

// Parse IP

// Check if valid global unicast IPv4 address

// DefaultEndpoint4Protocol : To ensure consistency, we generate default addr for listenAddress and advertiseAddress by one method. To avoid unnecessary port allocation work, we allocate fixed port for user defined protocol.
func DefaultEndpoint4Protocol(proto string) string { _ = "STUB: not implemented"; return "" }

// DefaultPort4Protocol returns the default port for different protocols
func DefaultPort4Protocol(proto string) string { _ = "STUB: not implemented"; return "" }

// URIs2Hosts returns hosts and schema
func URIs2Hosts(uris []string) ([]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

//not uri. but still permitted, like zookeeper,file system

// GetLocalIPv6 Get IPv6 address of NIC.
func GetLocalIPv6() string { _ = "STUB: not implemented"; return "" }

// Parse IP

// Check if valid IPv6 address

// IsIPv6Address check whether the IP is IPv6 address.
func IsIPv6Address(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func NormalizeAddrWithNetwork(addr string) (normalizedAddr string, network string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// IPv4：地址不变，网络类型tcp4

// IPv6：转为[ip]:port格式，网络类型tcp6

// splitIPAndPort 拆分IP和端口（兼容IPv4/IPv6格式）
func splitIPAndPort(addr string) (ip string, port string, err error) {
	_ = "STUB: not implemented"
	// 处理IPv6带方括号的情况（提前兼容，避免拆分错误）
	return "", "", nil
}

// 从后往前找最后一个冒号（区分IPv6多冒号和端口分隔符）

// StartListener start listener with address and tls(if has), returns the listener and the real listened ip/port
func StartListener(listenAddress string, tlsConfig *tls.Config) (listener net.Listener, listenedIP string, port string, err error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), "", "", nil
}

// ClientIP returns client ip
func ClientIP(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// RemoteIP returns remote ip
func RemoteIP(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// ForwardedIPs returns forwarded for ips
func ForwardedIPs(r *http.Request) []string { _ = "STUB: not implemented"; return nil }

// RealIP returns real ip
func RealIP(r *http.Request) string { _ = "STUB: not implemented"; return "" }
