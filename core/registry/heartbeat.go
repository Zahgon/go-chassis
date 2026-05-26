package registry

import (
	"time"
)

// DefaultRetryTime default retry time
const DefaultRetryTime = 10 * time.Second

// HeartbeatTask heart beat task struct
type HeartbeatTask struct {
	ServiceID  string
	InstanceID string
	Time       time.Time
	Running    bool
}

// HeartbeatService heartbeat service
type HeartbeatService struct {
	HeartbeatMode string
	shutdown      bool
	Interval      time.Duration
}

// Start start the heartbeat system
func (s *HeartbeatService) Start() { _ = "STUB: not implemented"; return }

// Stop stop the heartbeat system
func (s *HeartbeatService) Stop() {
	_ = "STUB: not implemented"

	// DoHeartBeat do heartbeat for each instance
	return
}

func (s *HeartbeatService) DoHeartBeat(microServiceID, microServiceInstanceID string, instanceHeartbeatMode string) error {
	_ = "STUB: not implemented"
	return nil
}

// run runs the heartbeat system
func (s *HeartbeatService) run() { _ = "STUB: not implemented"; return }

// sendNonPersistenceHeartBeat use http to send heartbeat
func (s *HeartbeatService) sendNonPersistenceHeartBeat() {
	_ = "STUB: not implemented"

	// first, wait for the successful registration, and then send the heartbeat to slow down the pressure of SC
	return
}

// RetryRegister try to register micro-service, and instance
func (s *HeartbeatService) RetryRegister(sid, iid string) { _ = "STUB: not implemented"; return }

// ReRegisterSelfMSandMSI 重新注册微服务和实例
func (s *HeartbeatService) ReRegisterSelfMSandMSI() error { _ = "STUB: not implemented"; return nil }

// reRegisterSelfMSI 只重新注册实例
func reRegisterSelfMSI(sid, iid string) error { _ = "STUB: not implemented"; return nil }
