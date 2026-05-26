package resource

import (
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

// RestFulHello is a struct used for implementation of restful hello program
type RestFulHello struct {
}

// Health
func (r *RestFulHello) Health(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulHello) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }
