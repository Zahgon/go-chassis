package main

import (
	"net/http"
	"strings"

	"github.com/go-chassis/go-chassis/v2/middleware/jwt"

	"github.com/go-chassis/go-chassis/v2"
	_ "github.com/go-chassis/go-chassis/v2/middleware/jwt"
	"github.com/go-chassis/go-chassis/v2/security/token"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
)

// if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/server/
func main() {
	chassis.RegisterSchema("rest", &HelloAuth{})

	jwt.Use(&jwt.Auth{
		MustAuth: func(req *http.Request) bool {
			if strings.Contains(req.URL.Path, "/login") {
				return false
			}
			return true
		},
		Realm: "test-realm",
		SecretFunc: func(claims interface{}, method token.SigningMethod) (interface{}, error) {
			return []byte("my_secret"), nil
		},
	})
	//start all server you register in server/schemas.
	if err := chassis.Init(); err != nil {
		openlog.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}

type User struct {
	Name string `json:"name"`
	Pwd  string `json:"password"`
}
type HelloAuth struct {
}

func (r *HelloAuth) Login(b *rf.Context) { _ = "STUB: not implemented"; return }

func (r *HelloAuth) Access(b *rf.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *HelloAuth) URLPatterns() []rf.Route { _ = "STUB: not implemented"; return nil }
