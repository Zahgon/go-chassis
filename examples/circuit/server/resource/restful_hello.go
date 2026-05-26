package resource

import (
	"sync"

	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

var l sync.Mutex

// RestFulMessage is a struct used to implement restful message
type RestFulMessage struct {
}

// DeadLock is used to simulate deadlock
func (r *RestFulMessage) DeadLock(b *rf.Context) { _ = "STUB: not implemented"; return }

// Sayhi is a method used to reply request user with hello world text
func (r *RestFulMessage) Sayhi(b *rf.Context) { _ = "STUB: not implemented"; return }

// Sayerror is a method used to reply request user with error
func (r *RestFulMessage) Sayerror(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulMessage) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }
