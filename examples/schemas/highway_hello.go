package schemas

import (
	"context"

	"github.com/go-chassis/go-chassis/v2/examples/schemas/helloworld"
)

// HelloServer is a struct
type HelloServer struct {
}

// SayHello is a method used to reply message
func (s *HelloServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
