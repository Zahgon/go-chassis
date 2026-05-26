package registry

// const
const (
	SSLEnabledQuery = "sslEnabled=true"
)

// Endpoint struct having full info about micro-service instance endpoint
type Endpoint struct {
	SSLEnabled bool   `json:"sslEnabled"`
	Address    string `json:"address"`
}

// NewEndPoint return a Endpoint object what parse from url
func NewEndPoint(schema string) (*Endpoint, error) {
	_ = "STUB: not implemented"
	return nil,

		// GenEndpoint return the endpoint string which it contain the sslEnabled=true query arg or not
		nil
}

func (e *Endpoint) GenEndpoint() string { _ = "STUB: not implemented"; return "" }

// IsSSLEnable return it is use ssl or not
func (e *Endpoint) IsSSLEnable() bool { _ = "STUB: not implemented"; return false }

// SetSSLEnable set ssl enable or not
func (e *Endpoint) SetSSLEnable(enabled bool) { _ = "STUB: not implemented"; return }

func (e *Endpoint) String() string { _ = "STUB: not implemented"; return "" }

func parseAddress(address string) (*Endpoint, error) { _ = "STUB: not implemented"; return nil, nil }
