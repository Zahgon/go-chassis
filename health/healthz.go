package health

import (
	"time"

	"github.com/go-chassis/go-chassis/v2/core/registry"
)

const (
	timeoutToPending = 1 * time.Second
	timeoutToPackage = 100 * time.Millisecond
	chanCapacity     = 1000
)

var defaultHealthChecker = &HealthChecker{}

func init() {
	defaultHealthChecker.Run()
}

// WrapInstance is the struct defines an instance object with appID/serviceName/version
type WrapInstance struct {
	AppID       string
	ServiceName string
	Version     string
	Instance    *registry.MicroServiceInstance
}

// String is the method returns the string type current instance's key value
func (i *WrapInstance) String() string { _ = "STUB: not implemented"; return "" }

// ServiceKey is the method returns the string type current instance's service key value
func (i *WrapInstance) ServiceKey() string { _ = "STUB: not implemented"; return "" }

// HealthChecker is the struct judges the instance health in the removing simpleCache
type HealthChecker struct {
	pendingCh chan *WrapInstance
	delCh     chan map[string]*WrapInstance
}

// Run is the method initializes and starts the health check process
func (hc *HealthChecker) Run() { _ = "STUB: not implemented"; return }

// Add is the method adds a key of the instance simpleCache into pending chan
func (hc *HealthChecker) Add(i *WrapInstance) error { _ = "STUB: not implemented"; return nil }

func (hc *HealthChecker) wait() { _ = "STUB: not implemented"; return }

// chan closed

// HealthCheck is the function adds the instance to HealthChecker
func HealthCheck(service, version, appID string, instance *registry.MicroServiceInstance) error {
	_ = "STUB: not implemented"
	return nil
}

// RefreshCache is the function to filter changes between new pulling instances and simpleCache
func RefreshCache(service string, ups []*registry.MicroServiceInstance, downs map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

// if full new instances or at less one instance, then refresh simpleCache immediately

// case: keep still alive instances

// case: remove instances with the non-up status

// case: keep instances returned HC ok

// case: add new come in instances

//todo remove this when the simpleCache struct can delete the key if the input is an empty slice
