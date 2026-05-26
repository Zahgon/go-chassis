package resource

import (
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

// RestFulHello is a struct used for implementation of restfull hello program
type RestFulHello struct {
}

// Hello
func (r *RestFulHello) Hello(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulHello) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }

// Legacy is a struct
type Legacy struct {
}

// Do
func (r *Legacy) Do(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *Legacy) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }

// Legacy is a struct
type Admin struct {
}

// Do
func (r *Admin) Do(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *Admin) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }
