package storage

// Options is yaml file struct to set db config
type Options struct {
	URI        string `yaml:"uri"`
	PoolSize   int    `yaml:"poolSize"`
	SSLEnabled bool   `yaml:"sslEnabled"`
	RootCA     string `yaml:"rootCAFile"`
	Timeout    string `yaml:"timeout"`
	VerifyPeer bool   `yaml:"verifyPeer"`
	CertFile   string `yaml:"certFile"`
	KeyFile    string `yaml:"keyFile"`
}

type Option func(opt *Options)

func PoolSize(poolSize int) Option { _ = "STUB: not implemented"; return *new(Option) }

func SSLEnabled(sslEnabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func RootCA(rootCAFile string) Option { _ = "STUB: not implemented"; return *new(Option) }

func Timeout(timeout string) Option { _ = "STUB: not implemented"; return *new(Option) }

func VerifyPeer(verifyPeer bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func CertFile(certFile string) Option { _ = "STUB: not implemented"; return *new(Option) }

func KeyFile(keyFile string) Option { _ = "STUB: not implemented"; return *new(Option) }

func NewConfig(uri string, opts ...func(opt *Options)) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}
