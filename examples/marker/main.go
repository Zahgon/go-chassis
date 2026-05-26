package main

import (
	"github.com/go-chassis/go-chassis/v2"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"

	_ "github.com/go-chassis/go-chassis/v2/middleware/ratelimiter"
)

type Hello struct{}

// Hello
func (r *Hello) Hello(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *Hello) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }

func main() {
	chassis.RegisterSchema("rest", &Hello{})
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
