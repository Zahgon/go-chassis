package servicecomb

import (
	"github.com/go-chassis/go-archaius/event"
)

// QPSEventListener is a struct used for Event listener
type QPSEventListener struct {
	//Key []string
	Key string
}

// Event is a method for QPS event listening
func (el *QPSEventListener) Event(e *event.Event) { _ = "STUB: not implemented"; return }
