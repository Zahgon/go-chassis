package client

import (
	"crypto/tls"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/go-chassis/go-chassis/v2/core/config/model"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

var clients = make(map[string]ProtocolClient)
var sl sync.RWMutex

// ErrClientNotExist happens if client do not exist
var ErrClientNotExist = errors.New("client not exists")

// DefaultPoolSize is 500
const DefaultPoolSize = 512

// Options is configs for client creation
type Options struct {
	Service       string
	PoolSize      int
	Timeout       time.Duration
	Endpoint      string
	PoolTTL       time.Duration
	TLSConfig     *tls.Config
	Failure       map[string]bool
	CheckRedirect func(req *http.Request, via []*http.Request) error
}

// GetFailureMap return failure map
func GetFailureMap(p string) map[string]bool { _ = "STUB: not implemented"; return nil }

// GetMaxIdleCon get max idle connection number you defined
// default is 512
func GetMaxIdleCon(p string) int { _ = "STUB: not implemented"; return 0 }

// CreateClient is for to create client based on protocol and the service name
func CreateClient(protocol, service, endpoint string, sslEnable bool, checkRedirect func(req *http.Request, via []*http.Request) error) (ProtocolClient, error) {
	_ = "STUB: not implemented"
	return *new(ProtocolClient), nil
}

//it will set tls config when provider's endpoint has sslEnable=true suffix or
// consumer had set provider tls config

// client verify target micro service's name in mutual tls
// remember to set SAN (Subject Alternative Name) as server's micro service name
// when generating server.csr

func generateKey(protocol, service, endpoint string) string { _ = "STUB: not implemented"; return "" }

// GetClient is to get the client based on protocol, service,endpoint name
func GetClient(i *invocation.Invocation) (ProtocolClient, error) {
	_ = "STUB: not implemented"
	return *new(ProtocolClient), nil
}

// Close close a client conn
func Close(protocol, service, endpoint string) error { _ = "STUB: not implemented"; return nil }

// SetTimeoutToClientCache set timeout to client
func SetTimeoutToClientCache(spec *model.IsolationWrapper) { _ = "STUB: not implemented"; return }

// EqualOpts equal newOpts and oldOpts
func EqualOpts(oldOpts, newOpts Options) Options { _ = "STUB: not implemented"; return *new(Options) }
