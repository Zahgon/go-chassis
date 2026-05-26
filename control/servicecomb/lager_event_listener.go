package servicecomb

import (
	"github.com/go-chassis/go-archaius/event"
)

const (
	//LagerLevelKey is a variable of type string
	LagerLevelKey = "logLevel"
)

// LagerEventListener is a struct used for Event listener
type LagerEventListener struct {
	//Key []string
	Key string
}

// Event is a method for Lager event listening
func (el *LagerEventListener) Event(e *event.Event) { _ = "STUB: not implemented"; return }
