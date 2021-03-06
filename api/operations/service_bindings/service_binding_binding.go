// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ServiceBindingBindingHandlerFunc turns a function with the right signature into a service binding binding handler
type ServiceBindingBindingHandlerFunc func(ServiceBindingBindingParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceBindingBindingHandlerFunc) Handle(params ServiceBindingBindingParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ServiceBindingBindingHandler interface for that can handle valid service binding binding params
type ServiceBindingBindingHandler interface {
	Handle(ServiceBindingBindingParams, interface{}) middleware.Responder
}

// NewServiceBindingBinding creates a new http.Handler for the service binding binding operation
func NewServiceBindingBinding(ctx *middleware.Context, handler ServiceBindingBindingHandler) *ServiceBindingBinding {
	return &ServiceBindingBinding{Context: ctx, Handler: handler}
}

/* ServiceBindingBinding swagger:route PUT /v2/service_instances/{instance_id}/service_bindings/{binding_id} ServiceBindings serviceBindingBinding

generation of a service binding

*/
type ServiceBindingBinding struct {
	Context *middleware.Context
	Handler ServiceBindingBindingHandler
}

func (o *ServiceBindingBinding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewServiceBindingBindingParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
