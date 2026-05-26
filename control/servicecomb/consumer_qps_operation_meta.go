package servicecomb

// ConsumerKeys contain consumer keys
type ConsumerKeys struct {
	MicroServiceName       string
	SchemaQualifiedName    string
	OperationQualifiedName string
}

// ProviderKeys contain provider keys
type ProviderKeys struct {
	Global          string
	ServiceOriented string
}

// Prefix is const
const Prefix = "cse.flowcontrol"

// GetConsumerKey get specific key for consumer
func GetConsumerKey(sourceName, serviceName, schemaID, OperationID string) *ConsumerKeys {
	_ = "STUB: not implemented"
	return nil

	//for mesher to govern
}

// GetProviderKey get specific key for provider
func GetProviderKey(sourceServiceName string) *ProviderKeys { _ = "STUB: not implemented"; return nil }

// GetSchemaQualifiedName get schema qualified name
func (op *ConsumerKeys) GetSchemaQualifiedName() string { _ = "STUB: not implemented"; return "" }

// GetMicroServiceSchemaOpQualifiedName get micro-service schema operation qualified name
func (op *ConsumerKeys) GetMicroServiceSchemaOpQualifiedName() string {
	_ = "STUB: not implemented"
	return ""
}

// GetMicroServiceName get micro-service name
func (op *ConsumerKeys) GetMicroServiceName() string { _ = "STUB: not implemented"; return "" }
