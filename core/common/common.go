package common

import (
	"context"
	"net/http"

	"github.com/go-chassis/go-archaius/source/remote"
)

// constant for provider and consumer
const (
	Provider = "Provider"
	Consumer = "Consumer"
)

const (
	// ScopeFull means service is able to access to another app's service
	ScopeFull = "full"
	// ScopeApp means service is not able to access to another app's service
	ScopeApp = "app"
)

// constant for micro service environment parameters
const (
	EnvCSEEndpoint = "PAAS_CSE_ENDPOINT"
	EnvNodeIP      = "HOSTING_SERVER_IP"
	EnvSchemaRoot  = "SCHEMA_ROOT"
	EnvSCEndpoint  = "PAAS_CSE_SC_ENDPOINT"
	EnvCCEndpoint  = "PAAS_CSE_CC_ENDPOINT"
)

// env connect with "." like servicecomb.service.name and servicecomb.service.version which can not be used in k8s.
// So we can not use archaius to set env.
// To support this declaring constant for service name and version
// constant for service name and version.
const (
	ServiceName = "CAS_COMPONENT_NAME"
	Version     = "CAS_INSTANCE_VERSION"
	App         = "CAS_APPLICATION_NAME"
	Env         = "ENVIRONMENT"
)

// constant for microservice environment
const (
	EnvValueDev  = "development"
	EnvValueProd = "production"
)

// constant for secure socket layer parameters
const (
	SslCipherPluginKey = "cipherPlugin"
	SslVerifyPeerKey   = "verifyPeer"
	SslCipherSuitsKey  = "cipherSuits"
	SslProtocolKey     = "protocol"
	SslCaFileKey       = "caFile"
	SslCertFileKey     = "certFile"
	SslKeyFileKey      = "keyFile"
	SslCertPwdFilePath = "certPwdFile"
	AKSKCustomCipher   = "servicecomb.credentials.akskCustomCipher"
	SslServerNameKey   = "serverName"
)

// constant for protocol types
const (
	ProtocolRest    = "rest"
	ProtocolHighway = "highway"
	LBSessionID     = "go-chassisLB"
)

// configuration placeholders
const (
	PlaceholderInternalIP = "$INTERNAL_IP"
)

// SessionNameSpaceKey metadata session namespace key
const SessionNameSpaceKey = "_Session_Namespace"

// SessionNameSpaceDefaultValue default session namespace value
const SessionNameSpaceDefaultValue = "default"

// DefaultKey default key
const DefaultKey = "default"

// DefaultValue default value
const DefaultValue = "default"

// BuildinTagApp build tag for the application
const BuildinTagApp = "app"

// BuildinTagVersion build tag version
const BuildinTagVersion = "version"

// BuildinLabelVersion build label for version
const BuildinLabelVersion = BuildinTagVersion + ":" + LatestVersion

// CallerKey caller key
const CallerKey = "caller"

// service comb headers
const (
	HeaderSourceName = "x-cse-src-microservice"
	// HeaderXCseContent is constant for header , get some json msg about HeaderSourceName like {"k":"v"}
	HeaderXCseContent = "x-cse-context"

	HeaderMark = "X-Mark"
)

// Rest metadata key for restful protocol
const (
	RestMethod    = "method"
	RestRoutePath = "url_pattern"
)

// constant for default application name and version
const (
	DefaultApp        = "default"
	DefaultVersion    = "0.0.1"
	LatestVersion     = "latest"
	AllVersion        = "0+"
	DefaultStatus     = "UP"
	TESTINGStatus     = "TESTING"
	DefaultLevel      = "BACK"
	DefaultHBInterval = 30
)

// constant used
const (
	HTTP   = "http"
	HTTPS  = "https"
	JSON   = "application/json"
	Create = "CREATE"
	Update = "UPDATE"
	Delete = "DELETE"

	Client           = "client"
	File             = "File"
	DefaultTenant    = "default"
	DefaultChainName = "default"

	FileRegistry      = "File"
	DefaultUserName   = "default"
	DefaultDomainName = "default"
	DefaultProvider   = "default"

	TRUE  = "true"
	FALSE = "false"
)

// const default config for config-server
const (
	DefaultRefreshMode = remote.ModeInterval
)

// ContextHeaderKey is the unified key of header value in context
// all protocol integrated with go chassis must set protocol header into context in this context key
type ContextHeaderKey struct{}

// NewContext transforms a metadata to context object
func NewContext(m map[string]string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// WithContext sets the KV and returns the context object
func WithContext(ctx context.Context, key, val string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FromContext return the headers which should be send to provider
// through transport
func FromContext(ctx context.Context) map[string]string { _ = "STUB: not implemented"; return nil }

// GetXCSEContext  get x-cse-context from req.header
func GetXCSEContext(k string, r *http.Request) string { _ = "STUB: not implemented"; return "" }

// SetXCSEContext  set value into x-cse-context
func SetXCSEContext(vm map[string]string, r *http.Request) { _ = "STUB: not implemented"; return }
