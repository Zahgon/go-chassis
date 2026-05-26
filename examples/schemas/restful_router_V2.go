package schemas

import (
	"github.com/go-chassis/go-chassis/v2/server/restful"
)

// RestFulRouterB is a struct used for implementation of restfull router program
type RestFulRouterB struct {
}

// Equal is method to compare given num and slice product
func (r *RestFulRouterB) Equal(context *restful.Context) { _ = "STUB: not implemented"; return }

// Say is method to reply version B say some info
func (r *RestFulRouterB) Say(context *restful.Context) { _ = "STUB: not implemented"; return }

// Operation is method to calculate  two num product
func (r *RestFulRouterB) Operation(context *restful.Context) { _ = "STUB: not implemented"; return }

// Info is a method used to reply version information
func (r *RestFulRouterB) Info(context *restful.Context) { _ = "STUB: not implemented"; return }

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulRouterB) URLPatterns() []restful.Route { _ = "STUB: not implemented"; return nil }
