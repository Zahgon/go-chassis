package metadata

import (
	"sync"
)

// variables of micro-service framework, mutex variable
var (
	msFramework *Framework
	Once        = &sync.Once{}
)

// Framework is for to represents name, version, registration
type Framework struct {
	Name     string
	Version  string
	Register string
}

// SetName is to set the framework name
func (f *Framework) SetName(name string) { _ = "STUB: not implemented"; return }

// SetVersion to set the version of framework
func (f *Framework) SetVersion(version string) { _ = "STUB: not implemented"; return }

// SetRegister to register the framework
func (f *Framework) SetRegister(register string) { _ = "STUB: not implemented"; return }

// NewFramework returns the object of msFramework
func NewFramework() *Framework { _ = "STUB: not implemented"; return nil }
