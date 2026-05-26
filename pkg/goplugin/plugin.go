package goplugin

// LookupPlugin lookup plugin
// Caller needs to determine itself whether the plugin file exists
func LookupPlugin(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// firstly search plugin in {ChassisHome}/lib

// secondly search plugin in /usr/lib

// LookUpSymbolFromPlugin looks up symbol from the plugin
func LookUpSymbolFromPlugin(plugName, symName string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
