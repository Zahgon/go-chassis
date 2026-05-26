package schema

import (
	swagger "github.com/go-chassis/go-restful-swagger20"
)

// MicroserviceMeta is the struct for micro service meta
type MicroserviceMeta struct {
	MicroserviceName string
	SchemaIDs        []string
}

// NewMicroserviceMeta gives the object of MicroserviceMeta
func NewMicroserviceMeta(microserviceName string) *MicroserviceMeta {
	_ = "STUB: not implemented"
	return nil
}

// defaultMicroserviceMetaMgr default micro-service meta-manager
var defaultMicroserviceMetaMgr map[string]*MicroserviceMeta

// schemaIDsMap default schema schema IDs map
var schemaIDsMap map[string]string

// defaultMicroServiceNames default micro-service names
var defaultMicroServiceNames = make([]string, 0)

// GetSchemaPath calculate the schema root path and return
func GetSchemaPath(name string) string { _ = "STUB: not implemented"; return "" }

// LoadSchema to load the schema files and micro-service information under the conf directory
// path is the conf path
func LoadSchema(path string) error {
	_ = "STUB: not implemented"
	/*
		conf/
		├── chassis.yaml
		├── microservice1
		│   └── schema
		│       ├── schema1.yaml
	*/return nil
}

// getSchemaNames 目录名为服务名
func getSchemaNames(confDir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// 遍历confDir下的microservice文件夹

// 仅读取负一级目录

// SetMicroServiceNames set micro service names
func SetMicroServiceNames(confDir string) error { _ = "STUB: not implemented"; return nil }

// 仅读取负一级目录

// loadSchemaFileContent load scheme file content
func loadSchemaFileContent(schemaPath string) (*MicroserviceMeta, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetContent get schema content by id
func GetContent(schemaID string) string { _ = "STUB: not implemented"; return "" }

// getFiles get files
func getFiles(fPath string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// schema文件名规则

// 遍历schemaPath下的schema文件

// 仅读取负一级文件

// GetMicroserviceNames get micro-service names
func GetMicroserviceNames() []string { _ = "STUB: not implemented"; return nil }

// GetSchemaIDs get schema IDs
func GetSchemaIDs(microserviceName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// init is for to initialize the defaultMicroserviceMetaMgr, and schemaIDsMap
func init() {
	defaultMicroserviceMetaMgr = make(map[string]*MicroserviceMeta)
	schemaIDsMap = make(map[string]string)
}

// SetSchemaInfo is for fill defaultMicroserviceMetaMgr and schemaIDsMap
func SetSchemaInfo(sws *swagger.SwaggerService) error { _ = "STUB: not implemented"; return nil }

// SetSchemaInfoByMap is for fill defaultMicroserviceMetaMgr and schemaIDsMap
func SetSchemaInfoByMap(schemaMap map[string]string) error { _ = "STUB: not implemented"; return nil }

// already read from conf/ServiceName dir
