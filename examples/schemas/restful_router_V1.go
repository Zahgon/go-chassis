package schemas

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
)

// RestFulRouterA is a struct used for implementation of restfull router program
type RestFulRouterA struct {
}

// Equal is method to compare given num and slice sum
func (r *RestFulRouterA) Equal(context *restful.Context) { _ = "STUB: not implemented"; return }

// Say is method to reply version A say some info
func (r *RestFulRouterA) Say(context *restful.Context) { _ = "STUB: not implemented"; return }

// Operation is method to add two num sum
func (r *RestFulRouterA) Operation(context *restful.Context) { _ = "STUB: not implemented"; return }

// Info is a method used to reply version information
func (r *RestFulRouterA) Info(context *restful.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulRouterA) URLPatterns() []restful.Route { _ = "STUB: not implemented"; return nil }
