package aes

import (
	"os"

	"github.com/go-chassis/go-chassis/v2/security/cipher"

	"github.com/go-chassis/cari/security"
	"github.com/go-chassis/openlog"
)

const cipherPlugin = "cipher_plugin.so"

// Cipher interface declares Init(), Encrypyt(), Decrypyt() methods
type Cipher interface {
	Init()
	Encrypt(src string) (string, error)
	Decrypt(src string) (string, error)
}

// HWAESCipher is a cipher used in huawei
type HWAESCipher struct {
	gcryptoEngine Cipher
}

func init() {
	if v, exist := os.LookupEnv("CIPHER_ROOT"); exist {
		err := os.Setenv("PAAS_CRYPTO_PATH", v)
		if err != nil {
			openlog.Warn("can not set env for cipher: " + err.Error())
		}
	}
	cipher.InstallCipherPlugin("aes", new)
}

func new() security.Cipher { _ = "STUB: not implemented"; return *new(security.Cipher) }

// Encrypt is method used for encryption
func (ac *HWAESCipher) Encrypt(src string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Decrypt is method used for decryption
func (ac *HWAESCipher) Decrypt(src string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
