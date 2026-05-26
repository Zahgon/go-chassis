package main

import (
	"github.com/go-chassis/openlog"

	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/bootstrap"
)

// if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/fileupload/client/
func main() {
	//Init framework
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}

	// file / form to upload
	uploadfile("file.input")
	uploadform("form.input")
}

func uploadfile(filename string) { _ = "STUB: not implemented"; return }

func uploadform(filename string) {
	_ = "STUB: not implemented"
	// Form part
	return
}
