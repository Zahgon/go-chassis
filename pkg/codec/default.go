package codec

import (
	"github.com/go-chassis/cari/codec"
)

// StdJson implement standard json codec
type StdJson struct {
}

func newDefault(opts Options) (codec.Codec, error) {
	_ = "STUB: not implemented"
	return *new(codec.Codec), nil
}

func (s *StdJson) Encode(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *StdJson) Decode(data []byte, v any) error { _ = "STUB: not implemented"; return nil }
