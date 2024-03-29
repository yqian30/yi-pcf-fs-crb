package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetCopyinstancesHandlerFunc turns a function with the right signature into a get copyinstances handler
type GetCopyinstancesHandlerFunc func(GetCopyinstancesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCopyinstancesHandlerFunc) Handle(params GetCopyinstancesParams) middleware.Responder {
	return fn(params)
}

// GetCopyinstancesHandler interface for that can handle valid get copyinstances params
type GetCopyinstancesHandler interface {
	Handle(GetCopyinstancesParams) middleware.Responder
}

// NewGetCopyinstances creates a new http.Handler for the get copyinstances operation
func NewGetCopyinstances(ctx *middleware.Context, handler GetCopyinstancesHandler) *GetCopyinstances {
	return &GetCopyinstances{Context: ctx, Handler: handler}
}

/*GetCopyinstances swagger:route GET /copies crb-web getCopyinstances

List all the current copy instances.

Return copyIds from Metadata-DB

*/
type GetCopyinstances struct {
	Context *middleware.Context
	Handler GetCopyinstancesHandler
}

func (o *GetCopyinstances) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetCopyinstancesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
