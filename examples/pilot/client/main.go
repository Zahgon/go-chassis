package main

import (
	"time"

	"github.com/go-chassis/openlog"

	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/bootstrap"
	"github.com/go-chassis/go-chassis/v2/core"
)

// if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/path/to/conf/folder
func main() {
	//chassis operation
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	restInvoker := core.NewRestInvoker()

	// use the configured chain
	for {
		callRest(restInvoker, 10)
		<-time.After(time.Second)
	}
}

func callRest(invoker *core.RestInvoker, i int) { _ = "STUB: not implemented"; return }

//use the invoker like http client.
