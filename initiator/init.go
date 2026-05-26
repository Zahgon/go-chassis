// Package initiator init necessary module
// before every other package init functions
package initiator

import (
	"github.com/go-chassis/go-chassis/v2/core/lager"
)

// LoggerOptions has the configuration about logging
var LoggerOptions *lager.Options

func init() {
	InitLogger()
}

// InitLogger initiate config file and openlog before other modules
func InitLogger() { _ = "STUB: not implemented"; return }

//initialize log in any case

// ParseLoggerConfig unmarshals the logger configuration file(lager.yaml)
func ParseLoggerConfig(file string) error { _ = "STUB: not implemented"; return nil }

func unmarshalYamlFile(file string, target interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
