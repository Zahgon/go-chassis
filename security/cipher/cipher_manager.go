package cipher

import (
	"github.com/go-chassis/cari/security"
)

const pluginSuffix = ".so"

// CipherPlugins is a map
var cipherPlugins = make(map[string]func() security.Cipher)

// InstallCipherPlugin is a function
func InstallCipherPlugin(name string, f func() security.Cipher) { _ = "STUB: not implemented"; return }

// NewCipher create and return a cipher
func NewCipher(name string) (security.Cipher, error) {
	_ = "STUB: not implemented"
	return *new(security.Cipher), nil
}

// GetCipherNewFunc return a function which is able to create a cipher
func GetCipherNewFunc(name string) (func() security.Cipher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadCipherFromPlugin(name string) (func() security.Cipher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Init() error { _ = "STUB: not implemented"; return nil }
