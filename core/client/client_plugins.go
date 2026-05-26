package client

// NewFunc is function for the client
type NewFunc func(Options) (ProtocolClient, error)

var plugins = make(map[string]NewFunc)

// GetClientNewFunc is to get the client
func GetClientNewFunc(name string) (NewFunc, error) {
	_ = "STUB: not implemented"
	return *new(NewFunc), nil
}

// InstallPlugin is plugin for the new function
func InstallPlugin(protocol string, f NewFunc) { _ = "STUB: not implemented"; return }
