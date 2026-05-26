package resource

import (
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

// MetadataHandler
type MetadataHandler struct {
}

// Handle
func (h *MetadataHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	_ = "STUB: not implemented"
	return
}

// Name
func (h *MetadataHandler) Name() string { _ = "STUB: not implemented"; return "" }

func newMetadataHandler() handler.Handler {
	_ = "STUB: not implemented"
	// call next chain
	return *new(handler.Handler)
}

func init() {
	handler.RegisterHandler("metadata-handler", newMetadataHandler)
}
