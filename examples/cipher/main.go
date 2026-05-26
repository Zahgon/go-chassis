package main

import (
	"github.com/go-chassis/cari/security"
	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/middleware/ratelimiter"
	"github.com/go-chassis/go-chassis/v2/security/cipher"
	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
)

type DemoResource struct {
}

func (r *DemoResource) Limit(b *restful.Context) { _ = "STUB: not implemented"; return }

// URLPatterns returns routes
func (r *DemoResource) URLPatterns() []restful.Route { _ = "STUB: not implemented"; return nil }

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/{project_root}/

func main() {
	cipher.InstallCipherPlugin("default", new)
	chassis.RegisterSchema("rest", &DemoResource{})
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}

// DefaultCipher is a struct
type DefaultCipher struct {
}

func new() security.Cipher {
	_ = "STUB: not implemented"
	return *

	// Encrypt is method used for encryption
	new(security.Cipher)
}

func (c *DefaultCipher) Encrypt(src string) (string, error) {
	_ = "STUB: not implemented"

	// Decrypt is method used for decryption
	return "", nil
}

func (c *DefaultCipher) Decrypt(src string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
