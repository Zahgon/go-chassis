package main

import (
	"sync"

	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/bootstrap"
	_ "github.com/go-chassis/go-chassis/v2/configserver"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/openlog"
)

var wg sync.WaitGroup

// if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/discovery/client/
func main() {
	//chassis operation
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed.")
		return
	}

	n := 10
	wg.Add(n)
	restInvoker := core.NewRestInvoker()
	for m := 0; m < n; m++ {
		go callRest(restInvoker)
	}
	wg.Wait()
}

func callRest(invoker *core.RestInvoker) { _ = "STUB: not implemented"; return }

//use the invoker like http client.
