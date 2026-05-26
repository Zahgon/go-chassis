package tls

import (
	"crypto/tls"
	"errors"

	"k8s.io/apimachinery/pkg/util/sets"
)

var errSSLConfigNotExist = errors.New("No SSL config")
var useDefaultSslTag = sets.NewString(
	"registry.Consumer.",
	"configServer.Consumer.",
	"monitor.Consumer.",
	"serviceDiscovery.Consumer.",
	"registrator.Consumer.",
	"contractDiscovery.Consumer.",
	"router.Consumer",
)

func hasDefaultSslTag(tag string) bool { _ = "STUB: not implemented"; return false }

func getDefaultSslConfigMap() map[string]string { _ = "STUB: not implemented"; return nil }

func getSSLConfigMap(tag, protocol, svcType, svcName string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// 使用默认配置

// 若配置了全局配置项，则覆盖默认配置

// consumer如果配置了通配 生效级别为全局配置之上 指定配置之下 且不为自己代理的服务设置证书配置

// 若配置了指定交互方的配置项，则覆盖全局配置

// 未设置ssl 且不提供内部默认ss配置 返回空字典

// use general TLSConfig
func useGeneralTLSConfig(svcType, svcName string) bool { _ = "STUB: not implemented"; return false }

func parseSSLConfig(sslConfigMap map[string]string) (*SSLConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSSLConfigByService get ssl configurations based on service
func GetSSLConfigByService(svcName, protocol, svcType string) (*SSLConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDefaultSSLConfig get default ssl configurations
func GetDefaultSSLConfig() *SSLConfig { _ = "STUB: not implemented"; return nil }

// generateSSLTag generate ssl tag
func generateSSLTag(svcName, protocol, svcType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetTLSConfigByService get tls configurations based on service
func GetTLSConfigByService(svcName, protocol, svcType string) (*tls.Config, *SSLConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// IsSSLConfigNotExist check the status of ssl configurations
func IsSSLConfigNotExist(e error) bool { _ = "STUB: not implemented"; return false }

// GetTLSConfig returns tls config from scheme and type
func GetTLSConfig(scheme, t string) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
