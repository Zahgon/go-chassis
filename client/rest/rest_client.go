package rest

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chassis/go-chassis/v2/core/client"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

const (
	// Name is a constant of type string
	Name = "rest"
	// FailureTypePrefix is a constant of type string
	FailureTypePrefix = "http_"
	//DefaultTimeoutBySecond defines the default timeout for http connections
	DefaultTimeoutBySecond = 60 * time.Second
	//DefaultKeepAliveSecond defines the connection time
	DefaultKeepAliveSecond = 60 * time.Second
	//DefaultMaxConnsPerHost defines the maximum number of concurrent connections
	DefaultMaxConnsPerHost = 512 * 20
	//SchemaHTTP represents the http schema
	SchemaHTTP = "http"
	//SchemaHTTPS represents the https schema
	SchemaHTTPS = "https"
)

var (

	//ErrInvalidResp invalid input
	ErrInvalidResp = errors.New("rest consumer response arg is not *rest.Response type")
)

func init() {
	client.InstallPlugin(Name, NewRestClient)
}

// Client is a struct
type Client struct {
	c    *http.Client
	opts client.Options
}

func (c *Client) Status(rsp interface{}) (status int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// NewRestClient is a function
func NewRestClient(opts client.Options) (client.ProtocolClient, error) {
	_ = "STUB: not implemented"
	return *new(client.ProtocolClient), nil
}

func newTransport(opts client.Options) *http.Transport { _ = "STUB: not implemented"; return nil }

// If a request fails, we generate an error.
func (c *Client) failure2Error(e error, r *http.Response, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// The Failure map defines whether or not a request fail.

// Call is a method which uses client struct object
func (c *Client) Call(ctx context.Context, addr string, inv *invocation.Invocation, rsp interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) String() string { _ = "STUB: not implemented"; return "" }

// Close release the idle connection
func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

// ReloadConfigs  reload configs for timeout and tls
func (c *Client) ReloadConfigs(opts client.Options) { _ = "STUB: not implemented"; return }

// GetOptions method return opts
func (c *Client) GetOptions() client.Options {
	_ = "STUB: not implemented"
	return *new(client.Options)
}

func (c *Client) contextToHeader(ctx context.Context, req *http.Request) {
	_ = "STUB: not implemented"
	return
}
