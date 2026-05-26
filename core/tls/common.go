package tls

import (
	"crypto/tls"
	"crypto/x509"

	//this import used for plain cipher
	_ "github.com/go-chassis/go-chassis/v2/security/cipher/plugins/plain"
)

// SSLConfig struct stores the necessary info for SSL configuration
type SSLConfig struct {
	CipherPlugin string   `yaml:"cipher_plugin" json:"cipherPlugin"`
	VerifyPeer   bool     `yaml:"verify_peer" json:"verifyPeer"`
	CipherSuites []uint16 `yaml:"cipher_suites" json:"cipherSuits"`
	MinVersion   uint16   `yaml:"min_version" json:"minVersion"`
	MaxVersion   uint16   `yaml:"max_version" json:"maxVersion"`
	CAFile       string   `yaml:"ca_file" json:"caFile"`
	CertFile     string   `yaml:"cert_file" json:"certFile"`
	KeyFile      string   `yaml:"key_file" json:"keyFile"`
	CertPWDFile  string   `yaml:"cert_pwd_file" json:"certPwdFile"`
	ServerName   string   `yaml:"server_name" json:"serverName"`
}

// TLSCipherSuiteMap is a map with key of type string and value of type unsigned integer
var TLSCipherSuiteMap = map[string]uint16{
	"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256": tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384": tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
}

// VersionMap is a map with key of type string and value of type unsigned integer
var VersionMap = map[string]uint16{
	"TLSv1.0": tls.VersionTLS10,
	"TLSv1.1": tls.VersionTLS11,
	"TLSv1.2": tls.VersionTLS12,
	"TLSv1.3": tls.VersionTLS13,
}

// GetX509CACertPool read a certificate file and gets the certificate configuration
func GetX509CACertPool(caCertFile string) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTLSConfig(sslConfig *SSLConfig, role string) (tlsConfig *tls.Config, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ca file is needed when veryPeer is true

// if cert pwd file is set, get the pwd

// certificate is necessary for server, optional for client

// GetClientTLSConfig function gets client side TLS config
func GetClientTLSConfig(sslConfig *SSLConfig) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetServerTLSConfig function gets server side TLD config
func GetServerTLSConfig(sslConfig *SSLConfig) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseSSLCipherSuites function parses cipher suites in to a list
func ParseSSLCipherSuites(ciphers string) ([]uint16, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 配置算法不存在

// ParseSSLProtocol function parses SSL protocols
func ParseSSLProtocol(sprotocol string) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }
