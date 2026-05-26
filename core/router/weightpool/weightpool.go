package weightpool

import (
	"sync"

	"github.com/go-chassis/go-chassis/v2/core/config"
)

var weightPool *SafePool
var once sync.Once

func init() { once.Do(func() { weightPool = &SafePool{pool: map[string]*Pool{}} }) }

// GetPool returns singleton of weightPool
func GetPool() *SafePool {
	_ = "STUB: not implemented"

	// SafePool is a cache for pool of all destination
	return nil
}

type SafePool struct {
	sync.RWMutex
	pool map[string]*Pool
}

// Get returns specific pool for key
func (s *SafePool) Get(key string) (*Pool, bool) { _ = "STUB: not implemented"; return nil, false }

// Set can set pool to safe cache
func (s *SafePool) Set(key string, value *Pool) { _ = "STUB: not implemented"; return }

// Reset can delete pool for specific key
func (s *SafePool) Reset(key string) { _ = "STUB: not implemented"; return }

// Pool defines sets of weighted tags
type Pool struct {
	tags []config.RouteTag

	mu  sync.RWMutex
	gcd int
	max int
	i   int
	cw  int
	num int
}

// NewPool returns pool for provided tags
func NewPool(routeTags ...*config.RouteTag) *Pool { _ = "STUB: not implemented"; return nil }

// PickOne returns tag according to its weight
func (p *Pool) PickOne() *config.RouteTag { _ = "STUB: not implemented"; return nil }

func (p *Pool) refreshGCD(t *config.RouteTag) { _ = "STUB: not implemented"; return }

func gcd(a, b int) int { _ = "STUB: not implemented"; return 0 }
