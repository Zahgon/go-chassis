package schemas

import (
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

// TracingHello is a struct
type TracingHello struct {
}

// Trace is a method
func (r *TracingHello) Trace(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *TracingHello) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }
