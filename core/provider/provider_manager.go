package provider

// plugin name and schemas map
var providerPlugins = make(map[string]func(string) Provider)

// microservice name and schemas map
var providers = make(map[string]Provider)

//TODO locks

// InstallProviderPlugin install provider plugin
func InstallProviderPlugin(pluginName string, newFunc func(string) Provider) {
	_ = "STUB: not implemented"
	return
}

// todo: return error.

// RegisterProvider register provider
func RegisterProvider(pluginName string, microserviceName string) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

// RegisterCustomProvider register customer provider
func RegisterCustomProvider(microserviceName string, p Provider) { _ = "STUB: not implemented"; return }

// GetProvider get provider
func GetProvider(microserviceName string) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

// RegisterSchemaWithName register schema with name
func RegisterSchemaWithName(microserviceName string, schemaID string, schema interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterSchema register schema
func RegisterSchema(microserviceName string, schema interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetOperation get operation
func GetOperation(microserviceName string, schemaID string, operationID string) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}
