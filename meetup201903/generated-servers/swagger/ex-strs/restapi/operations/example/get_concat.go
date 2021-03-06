// Code generated by go-swagger; DO NOT EDIT.

package example

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetConcatHandlerFunc turns a function with the right signature into a get concat handler
type GetConcatHandlerFunc func(GetConcatParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConcatHandlerFunc) Handle(params GetConcatParams) middleware.Responder {
	return fn(params)
}

// GetConcatHandler interface for that can handle valid get concat params
type GetConcatHandler interface {
	Handle(GetConcatParams) middleware.Responder
}

// NewGetConcat creates a new http.Handler for the get concat operation
func NewGetConcat(ctx *middleware.Context, handler GetConcatHandler) *GetConcat {
	return &GetConcat{Context: ctx, Handler: handler}
}

/*GetConcat swagger:route GET /concat example getConcat

concat

Returns concatenated string of inputs

*/
type GetConcat struct {
	Context *middleware.Context
	Handler GetConcatHandler
}

func (o *GetConcat) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetConcatParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
