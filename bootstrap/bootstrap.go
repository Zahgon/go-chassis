package bootstrap

var bootstrapPlugins = make([]*PluginItem, 0)

// PluginItem include name and plugin implementation
type PluginItem struct {
	Name   string
	Plugin Plugin
}

// Plugin is a interface which declares Init method
type Plugin interface {
	Init() error
}

// Func The Func type is an adapter to allow the use of ordinary functions as bootstrapPlugin.
type Func func() error

// Init is a method
func (b Func) Init() error {
	_ = "STUB: not implemented"

	// InstallPlugin is a function which installs plugin,
	// during initiating of go chassis, plugins will be executed
	return nil
}

func InstallPlugin(name string, plugin Plugin) { _ = "STUB: not implemented"; return }

// Bootstrap will boot plugins in orders
func Bootstrap() { _ = "STUB: not implemented"; return }
