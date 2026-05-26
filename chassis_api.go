package chassis

import (
	"os"

	//init logger first
	_ "github.com/go-chassis/go-chassis/v2/initiator"

	// transport handler
	_ "github.com/go-chassis/go-chassis/v2/core/client"

	//load balancing
	_ "github.com/go-chassis/go-chassis/v2/pkg/loadbalancing"

	//protocols
	_ "github.com/go-chassis/go-chassis/v2/client/rest"
	_ "github.com/go-chassis/go-chassis/v2/server/restful"

	//router
	_ "github.com/go-chassis/go-chassis/v2/core/router/servicecomb"
	//control panel
	_ "github.com/go-chassis/go-chassis/v2/control/servicecomb"
	// registry
	_ "github.com/go-chassis/go-chassis/v2/core/registry/servicecenter"
	"github.com/go-chassis/go-chassis/v2/core/server"

	// prometheus reporter for circuit breaker metrics
	_ "github.com/go-chassis/go-chassis/v2/third_party/forked/afex/hystrix-go/hystrix/reporter"
	// aes package handles security related plugins
	_ "github.com/go-chassis/go-chassis/v2/security/cipher/plugins/aes"
	_ "github.com/go-chassis/go-chassis/v2/security/cipher/plugins/plain"

	//config servers
	_ "github.com/go-chassis/go-archaius/source/remote"
	_ "github.com/go-chassis/go-archaius/source/remote/kie"
)

var goChassis *chassis

func init() {
	goChassis = &chassis{}
}

// RegisterSchema Register a API service to specific server by name
// You must register API first before Call Init
func RegisterSchema(serverName string, structPtr interface{}, opts ...server.RegisterOption) {
	_ = "STUB: not implemented"
	return
}

// SetDefaultConsumerChains your custom chain map for Consumer,if there is no config, this default chain will take affect
func SetDefaultConsumerChains(c map[string]string) { _ = "STUB: not implemented"; return }

// SetDefaultProviderChains set your custom chain map for Provider,if there is no config, this default chain will take affect
func SetDefaultProviderChains(c map[string]string) { _ = "STUB: not implemented"; return }

// HijackSignal set signals that want to hijack.
func HijackSignal(sigs ...os.Signal) { _ = "STUB: not implemented"; return }

// InstallPreShutdown instal what you want to achieve before graceful shutdown
func InstallPreShutdown(name string, f func(os.Signal)) {
	_ = "STUB: not implemented"
	// lazy init
	return
}

// InstallPostShutdown instal what you want to achieve after graceful shutdown
func InstallPostShutdown(name string, f func(os.Signal)) {
	_ = "STUB: not implemented"
	// lazy init
	return
}

// HijackGracefulShutdown reset GracefulShutdown
func HijackGracefulShutdown(f func(os.Signal)) { _ = "STUB: not implemented"; return }

// Run bring up the service,it waits for os signal,and shutdown gracefully
// before all protocol server start successfully, it may return error.
func Run(options ...server.RunOption) error { _ = "STUB: not implemented"; return nil }

//Register instance after Server started

func waitingSignal() { _ = "STUB: not implemented"; return }

// GracefulShutdown graceful shut down api
func GracefulShutdown(s os.Signal) { _ = "STUB: not implemented"; return }

// Init prepare the chassis framework runtime
func Init() error { _ = "STUB: not implemented"; return nil }
