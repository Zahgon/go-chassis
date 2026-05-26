package schemas

import (
	"math/rand"

	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

var num = rand.Intn(100)

// RestFulHello is a struct used for implementation of restfull hello program
type RestFulHello struct {
}

// Sayhello is a method used to reply user with hello
func (r *RestFulHello) Root(b *rf.Context) { _ = "STUB: not implemented"; return }

// Sayhello is a method used to reply user with hello
func (r *RestFulHello) Sayhello(b *rf.Context) { _ = "STUB: not implemented"; return }

// Sayhi is a method used to reply user with hello world text
func (r *RestFulHello) Sayhi(b *rf.Context) { _ = "STUB: not implemented"; return }

// SayJSON is a method used to reply user hello in json format
func (r *RestFulHello) SayJSON(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulHello) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }

// RestFulMessage is a struct used to implement restful message
type RestFulMessage struct {
}

// Saymessage is used to reply user with his name
func (r *RestFulMessage) Saymessage(b *rf.Context) { _ = "STUB: not implemented"; return }

// Sayhi is a method used to reply request user with hello world text
func (r *RestFulMessage) Sayhi(b *rf.Context) { _ = "STUB: not implemented"; return }

// Sayerror is a method used to reply request user with error
func (r *RestFulMessage) Sayerror(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulMessage) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }

// Hello is a struct used for implementation of restfull hello program
type Hello struct{}

// Hello
func (r *Hello) Hello(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *Hello) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }

// Legacy is a struct
type Legacy struct{}

// Do
func (r *Legacy) Do(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *Legacy) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }

// Legacy is a struct
type Admin struct{}

// Do
func (r *Admin) Do(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *Admin) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }
