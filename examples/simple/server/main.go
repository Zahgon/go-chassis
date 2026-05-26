package main

import (
	"github.com/go-chassis/go-chassis/v2"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/server/

type RestFulHello struct {
}

func (r *RestFulHello) Root(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulHello) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }

func main() {
	chassis.RegisterSchema("rest", &RestFulHello{})
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
