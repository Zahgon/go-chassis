//go:build !go1.10 || !debug
// +build !go1.10 !debug

package goplugin

import "plugin"

// LoadPlugin load plugin
func LoadPlugin(name string) (*plugin.Plugin, error) { _ = "STUB: not implemented"; return nil, nil }

func lookUp(plugName, symName string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
