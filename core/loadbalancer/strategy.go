package loadbalancer

import (
	"math/rand"
	"time"
)

var strategies = make(map[string]func() Strategy)
var i int

func init() {
	rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().Unix())
	i = rand.Int()
}

// InstallStrategy install strategy
func InstallStrategy(name string, s func() Strategy) { _ = "STUB: not implemented"; return }

// GetStrategyPlugin get strategy plugin
func GetStrategyPlugin(name string) (func() Strategy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
